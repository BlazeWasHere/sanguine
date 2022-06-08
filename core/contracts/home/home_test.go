package home_test

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/contracts/home"
	"math/big"
	"time"
)

func (h HomeSuite) TestDispatchTopic() {
	// init the dispatch event
	txContext := h.testBackend.GetTxContext(h.GetTestContext(), nil)

	dispatchSink := make(chan *home.HomeDispatch)
	sub, err := h.homeContract.WatchDispatch(&bind.WatchOpts{Context: h.GetTestContext()}, dispatchSink, [][32]byte{}, []*big.Int{}, []uint64{})
	Nil(h.T(), err)

	tx, err := h.homeContract.Dispatch(txContext.TransactOpts, 1, [32]byte{}, nil)
	Nil(h.T(), err)

	h.testBackend.WaitForConfirmation(h.GetTestContext(), tx)

	watchCtx, cancel := context.WithTimeout(h.GetTestContext(), time.Second*10)
	defer cancel()

	select {
	// check for errors and fail
	case <-watchCtx.Done():
		h.T().Error(h.T(), fmt.Errorf("test context completed %w", h.GetTestContext().Err()))
	case <-sub.Err():
		h.T().Error(h.T(), sub.Err())
	// get dispatch event
	case item := <-dispatchSink:
		parser, err := home.NewParser(h.homeContract.Address())
		Nil(h.T(), err)

		eventType, ok := parser.EventType(item.Raw)
		True(h.T(), ok)
		Equal(h.T(), eventType, home.DispatchEvent)

		break
	}
}

func (h HomeSuite) TestUpdateTopic() {
	h.T().Skip("TODO: test this. Mocker should be able to mock this out")
}