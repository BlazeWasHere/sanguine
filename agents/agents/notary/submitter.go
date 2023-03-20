package notary

import (
	"context"
	"time"

	"github.com/synapsecns/sanguine/agents/db"
	"github.com/synapsecns/sanguine/agents/domains"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
)

// AttestationSubmitter submits updates continuously.
type AttestationSubmitter struct {
	// domain allows access to the origin contract
	domain domains.DomainClient
	// db contains the transaction queue legacyDB
	db db.SynapseDB
	// signer is the signer
	signer signer.Signer
	// interval is count in seconds
	interval time.Duration
}

// NewAttestationSubmitter creates an update producer.
func NewAttestationSubmitter(domain domains.DomainClient, db db.SynapseDB, signer signer.Signer, interval time.Duration) AttestationSubmitter {
	return AttestationSubmitter{
		domain:   domain,
		db:       db,
		signer:   signer,
		interval: interval,
	}
}

// Start runs the update submitter.
func (u AttestationSubmitter) Start(ctx context.Context) error {
	// TODO (joe): this needs to be refactored after we do the GlobalRegistry stuff
	return nil
	/*committedNonce, err := u.domain.AttestationCollector().GetLatestNonce(ctx, u.domain.Config().DomainID, u.signer)
	if err != nil {
		return fmt.Errorf("could not get committed root: %w", err)
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(u.interval):
			nonce, err := u.db.RetrieveSignedAttestationByNonce(ctx, u.domain.Config().DomainID, committedNonce+1)
			if errors.Is(err, db.ErrNotFound) {
				logger.Infof("No produced attestation to submit for nonce: %d", nonce)
				continue
			} else if err != nil {
				return fmt.Errorf("could not retrieve produced update: %w", err)
			}

			err = u.domain.AttestationCollector().SubmitAttestation(ctx, u.signer, nonce)
			if err != nil {
				return fmt.Errorf("could not produce update: %w", err)
			}

			committedNonce++
		}
	}*/
}