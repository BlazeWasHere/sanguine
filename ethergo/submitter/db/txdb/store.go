package txdb

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/imkira/go-interpol"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/submitter/db"
	"github.com/synapsecns/sanguine/ethergo/util"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math/big"
)

// NewTXStore creates a new transaction store.
func NewTXStore(db *gorm.DB, metrics metrics.Handler) *Store {
	return &Store{
		db:      db,
		metrics: metrics,
	}
}

// Store is the sqlite store. It extends the base store for sqlite specific queries.
type Store struct {
	db      *gorm.DB
	metrics metrics.Handler
}

// DBTransaction is a function that can be used to execute a transaction on the database.
func (s *Store) DBTransaction(parentCtx context.Context, f db.TransactionFunc) (err error) {
	ctx, span := s.metrics.Tracer().Start(parentCtx, "db_transaction")
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	//nolint: wrapcheck
	return s.db.Transaction(func(tx *gorm.DB) error {
		txDB := NewTXStore(tx, s.metrics)
		return f(ctx, txDB)
	})
}

// MarkAllBeforeOrAtNonceReplacedOrConfirmed marks all transactions before or at the given nonce as replaced or confirmed.
func (s *Store) MarkAllBeforeOrAtNonceReplacedOrConfirmed(ctx context.Context, signer common.Address, chainID *big.Int, nonce uint64) error {
	dbTX := s.db.WithContext(ctx).Model(&ETHTX{}).
		Where(fmt.Sprintf("%s = ?", chainIDFieldName), chainID).
		Where(fmt.Sprintf("%s <= ?", nonceFieldName), nonce).
		Where(fmt.Sprintf("%s = ?", fromFieldName), signer.String()).
		// just in case we're updating a tx already marked as confirmed
		Where("%s IN ?", statusFieldName, []uint8{db.Submitted.Int(), db.FailedSubmit.Int()}).
		Updates(map[string]interface{}{statusFieldName: db.ReplacedOrConfirmed})

	if dbTX.Error != nil {
		return fmt.Errorf("could not mark all txs before nonce %d as replaced: %w", nonce, dbTX.Error)
	}

	return nil
}

// MaxResultsPerChain is the maximum number of transactions to return per chain id.
// it is exported for testing.
// TODO: this should be an option passed to the GetTXs function.
const MaxResultsPerChain = 50

func statusToArgs(matchStatuses ...db.Status) []int {
	inArgs := make([]int, len(matchStatuses))
	for i := range matchStatuses {
		inArgs[i] = int(matchStatuses[i].Int())
	}
	return inArgs
}

// convertTXS converts a slice of ETHTXes to a slice of db.TXes.
func convertTXS(ethTxes []ETHTX) (txs []db.TX, err error) {
	for _, dbTX := range ethTxes {
		var marshalledTx types.Transaction
		err = marshalledTx.UnmarshalBinary(dbTX.RawTx)
		if err != nil {
			return nil, fmt.Errorf("could not unmarshal tx: %w", err)
		}

		retTX := db.TX{
			Transaction: &marshalledTx,
			Status:      dbTX.Status,
		}

		// this is fine since we're an implementing db package
		retTX.UnsafeSetCreationTime(dbTX.CreatedAt)

		txs = append(txs, retTX)
	}
	return txs, nil
}

// GetTXS returns all transactions for a given address on a given (or any) chain id that match a given status.
// there is a limit of 50 transactions per chain id. The limit does not make any guarantees about the number of nonces per chain.
// the submitter will get only the most recent tx submitted for each chain so this can be used for gas pricing.
func (s *Store) GetTXS(ctx context.Context, fromAddress common.Address, chainID *big.Int, matchStatuses ...db.Status) (txs []db.TX, err error) {
	var dbTXs []ETHTX

	inArgs := statusToArgs(matchStatuses...)

	query := ETHTX{
		From: fromAddress.String(),
	}

	if chainID != nil {
		query.ChainID = chainID.Uint64()
	}

	tableName, err := dbcommon.GetModelName(s.db, &ETHTX{})
	if err != nil {
		return nil, fmt.Errorf("could not get table name: %w", err)
	}

	subQuery := s.DB().Model(&ETHTX{}).
		Select(fmt.Sprintf("MAX(%s) as %s, %s, %s", idFieldName, idFieldName, nonceFieldName, chainIDFieldName)).
		Where(query).
		Where(fmt.Sprintf("%s IN ?", statusFieldName), inArgs).
		Group(fmt.Sprintf("%s, %s", nonceFieldName, chainIDFieldName)).
		Order(fmt.Sprintf("%s asc", nonceFieldName)).
		Limit(MaxResultsPerChain)

	joinQuery, err := interpol.WithMap(
		"INNER JOIN (?) as subquery on `{table}`.`{id}` = `subquery`.`{id}` AND `{table}`.`{chainID}` = `subquery`.`{chainID}`", map[string]string{
			"table":   tableName,
			"id":      idFieldName,
			"chainID": chainIDFieldName,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("could not build join query: %w", err)
	}

	tx := s.DB().WithContext(ctx).
		Model(&ETHTX{}).
		Where(query).
		Joins(joinQuery, subQuery).
		Order(fmt.Sprintf("subquery.%s, subquery.%s, %s desc", chainIDFieldName, nonceFieldName, createdAtFieldName)).
		Find(&dbTXs)

	if tx.Error != nil {
		return nil, fmt.Errorf("could not get txs: %w", tx.Error)
	}

	txs, err = convertTXS(dbTXs)
	if err != nil {
		return nil, fmt.Errorf("could not convert txes: %w", err)
	}

	return txs, nil
}

// GetAllTXAttemptByStatus returns all transactions for a given address on a given (or any) chain id that match a given status.
func (s *Store) GetAllTXAttemptByStatus(ctx context.Context, fromAddress common.Address, chainID *big.Int, matchStatuses ...db.Status) (txs []db.TX, err error) {
	var dbTXs []ETHTX

	query := ETHTX{
		From: fromAddress.String(),
	}

	if chainID != nil {
		query.ChainID = chainID.Uint64()
	}

	tableName, err := dbcommon.GetModelName(s.db, &ETHTX{})
	if err != nil {
		return nil, fmt.Errorf("could not get table name: %w", err)
	}

	subQuery := s.DB().Model(&ETHTX{}).
		Select(fmt.Sprintf("MAX(%s) as %s, %s, %s", idFieldName, idFieldName, nonceFieldName, chainIDFieldName)).
		Where(query).
		Where(fmt.Sprintf("%s IN ?", statusFieldName), statusToArgs(matchStatuses...)).
		Group(fmt.Sprintf("%s, %s", nonceFieldName, chainIDFieldName)).
		Order(fmt.Sprintf("%s asc", nonceFieldName)).
		Limit(MaxResultsPerChain)

	// one consequence of innerjoining on nonce is we can't cap the max results for the whole query. This is a known limitation
	joinQuery, err := interpol.WithMap(
		"INNER JOIN (?) as subquery on `{table}`.`{nonce}` = `subquery`.`{nonce}` AND `{table}`.`{chainID}` = `subquery`.`{chainID}`", map[string]string{
			"table":   tableName,
			"nonce":   nonceFieldName,
			"chainID": chainIDFieldName,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("could not build join query: %w", err)
	}

	tx := s.DB().WithContext(ctx).
		Model(&ETHTX{}).
		Where(query).
		Joins(joinQuery, subQuery).
		Order(fmt.Sprintf("subquery.%s, subquery.%s, %s desc", chainIDFieldName, nonceFieldName, createdAtFieldName)).
		Find(&dbTXs)

	if tx.Error != nil {
		return nil, fmt.Errorf("could not get txes: %w", err)
	}

	txs, err = convertTXS(dbTXs)
	if err != nil {
		return nil, fmt.Errorf("could not convert txes: %w", err)
	}

	return txs, nil
}

// GetNonceForChainID gets the nonce for the given chain id.
func (s *Store) GetNonceForChainID(ctx context.Context, fromAddress common.Address, chainID *big.Int) (nonce uint64, err error) {
	var newNonce sql.NullInt64

	selectMaxNonce := fmt.Sprintf("max(`%s`)", nonceFieldName)

	dbTx := s.DB().WithContext(ctx).Model(&ETHTX{}).Select(selectMaxNonce).Where(ETHTX{
		From:    fromAddress.String(),
		ChainID: chainID.Uint64(),
	}).Scan(&newNonce)

	if dbTx.Error != nil {
		return 0, fmt.Errorf("could not get nonce for chain id: %w", dbTx.Error)
	}

	// if no nonces, return the corresponding error.
	if newNonce.Int64 == 0 {
		// we need to check if any nonces exist first
		var count int64
		dbTx = s.DB().WithContext(ctx).Model(&ETHTX{}).Where(ETHTX{ChainID: chainID.Uint64(), From: fromAddress.String()}).Count(&count)
		if dbTx.Error != nil {
			return 0, fmt.Errorf("error getting count on %T: %w", &ETHTX{}, dbTx.Error)
		}

		if count == 0 {
			return 0, db.ErrNoNonceForChain
		}
	}

	return uint64(newNonce.Int64), nil
}

// PutTXS puts a transaction in the database.
func (s Store) PutTXS(ctx context.Context, txs ...db.TX) error {
	var toInsert []*ETHTX

	for _, tx := range txs {
		marshalledTX, err := tx.MarshalBinary()
		if err != nil {
			return fmt.Errorf("could not marshall tx to json: %w", err)
		}

		msg, err := util.TxToCall(tx)
		if err != nil {
			return fmt.Errorf("could not recover signer from tx: %w", err)
		}

		if tx.To() == nil {
			return errors.New("tx has no to address")
		}

		if !tx.ChainId().IsUint64() {
			return errors.New("chainid is not uint64")
		}

		toInsert = append(toInsert, &ETHTX{
			From:    msg.From.String(),
			ChainID: tx.ChainId().Uint64(),
			Nonce:   tx.Nonce(),
			RawTx:   marshalledTX,
			TXHash:  tx.Hash().String(),
			Status:  tx.Status,
		})
	}

	dbTX := s.DB().WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: txHashFieldName}},
			DoUpdates: clause.AssignmentColumns([]string{
				statusFieldName,
			}),
		}).
		Create(toInsert)

	if dbTX.Error != nil {
		return fmt.Errorf("could not store tx: %w", dbTX.Error)
	}
	return nil
}

// DB gets the database.
func (s Store) DB() *gorm.DB {
	return s.db
}

var _ db.Service = &Store{}
