package types_test

import (
	"context"
	"crypto/rand"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/testutil"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
	"math/big"
	"testing"
	"time"
)

func TestEncodeTipsParity(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	testBackend := simulated.NewSimulatedBackend(ctx, t)
	deployManager := testutil.NewDeployManager(t)

	_, handle := deployManager.GetTipsHarness(ctx, testBackend)

	// make sure constants match
	tipsVersion, err := handle.TipsVersion(&bind.CallOpts{Context: ctx})
	Nil(t, err)
	Equal(t, tipsVersion, types.TipsVersion)

	notaryOffset, err := handle.OffsetNotary(&bind.CallOpts{Context: ctx})
	Nil(t, err)
	Equal(t, notaryOffset, big.NewInt(types.OffsetNotary))

	relayerOffset, err := handle.OffsetBroadcaster(&bind.CallOpts{Context: ctx})
	Nil(t, err)
	Equal(t, relayerOffset, big.NewInt(types.OffsetBroadcaster))

	proverOffset, err := handle.OffsetProver(&bind.CallOpts{Context: ctx})
	Nil(t, err)
	Equal(t, proverOffset, big.NewInt(types.OffsetProver))

	processorOffset, err := handle.OffsetExecutor(&bind.CallOpts{Context: ctx})
	Nil(t, err)
	Equal(t, processorOffset, big.NewInt(types.OffsetExecutor))

	// we want to make sure we can deal w/ overflows
	notaryTip := randomUint96BigInt(t)
	broadcasterTip := randomUint96BigInt(t)
	proverTip := randomUint96BigInt(t)
	executorTip := randomUint96BigInt(t)

	solidityFormattedTips, err := handle.FormatTips(&bind.CallOpts{Context: ctx}, notaryTip, broadcasterTip, proverTip, executorTip)
	Nil(t, err)

	goTips, err := types.EncodeTips(types.NewTips(notaryTip, broadcasterTip, proverTip, executorTip))
	Nil(t, err)

	Equal(t, goTips, solidityFormattedTips)
}

// randomUint96BigInt is a helper method for generating random uint96 values
// see:  https://stackoverflow.com/a/45428754
func randomUint96BigInt(tb testing.TB) *big.Int {
	tb.Helper()

	// Max random value, a 130-bits integer, i.e 2^96 - 1
	max := new(big.Int)
	max.Exp(big.NewInt(2), big.NewInt(96), nil).Sub(max, big.NewInt(1))

	// Generate cryptographically strong pseudo-random between 0 - max
	n, err := rand.Int(rand.Reader, max)
	Nil(tb, err)

	return n
}

func TestEncodeAttestationParity(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	testBackend := simulated.NewSimulatedBackend(ctx, t)
	deployManager := testutil.NewDeployManager(t)

	domain := gofakeit.Uint32()
	nonce := gofakeit.Uint32()
	root := common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64()))

	_, attesationContract := deployManager.GetAttestationHarness(ctx, testBackend)

	contractData, err := attesationContract.FormatAttestationData(&bind.CallOpts{Context: ctx}, domain, nonce, root)
	Nil(t, err)

	goFormattedData, err := types.EncodeAttestation(types.NewAttestation(domain, nonce, root))
	Nil(t, err)
	Equal(t, contractData, goFormattedData)
}

func TestEncodeSignedAttestationParity(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	testBackend := simulated.NewSimulatedBackend(ctx, t)
	deployManager := testutil.NewDeployManager(t)

	_, attesationContract := deployManager.GetAttestationHarness(ctx, testBackend)

	domain := gofakeit.Uint32()
	nonce := gofakeit.Uint32()
	root := common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64()))

	sig := types.NewSignature(new(big.Int).SetUint64(uint64(gofakeit.Uint8())), new(big.Int).SetUint64(gofakeit.Uint64()), new(big.Int).SetUint64(gofakeit.Uint64()))

	signedAttestation := types.NewSignedAttestation(
		types.NewAttestation(domain, nonce, root),
		sig,
	)

	encodedSignature, err := types.EncodeSignature(sig)
	Nil(t, err)

	signedContractAttestation, err := attesationContract.FormatAttestation(&bind.CallOpts{Context: ctx}, domain, nonce, root, encodedSignature)
	Nil(t, err)

	goData, err := types.EncodeSignedAttestation(signedAttestation)
	Nil(t, err)

	Equal(t, signedContractAttestation, goData)
}

func TestMessageEncodeParity(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	testBackend := simulated.NewSimulatedBackend(ctx, t)
	deployManager := testutil.NewDeployManager(t)
	_, messageContract := deployManager.GetMessageHarness(ctx, testBackend)

	// check constant parity
	version, err := messageContract.MessageVersion(&bind.CallOpts{Context: ctx})
	Nil(t, err)
	Equal(t, version, types.MessageVersion)

	headerOffset, err := messageContract.OffsetHeader(&bind.CallOpts{Context: ctx})
	Nil(t, err)
	Equal(t, headerOffset, big.NewInt(int64(types.HeaderOffset)))

	// generate some fake data
	origin := gofakeit.Uint32()
	sender := common.BigToHash(big.NewInt(gofakeit.Int64()))
	nonce := gofakeit.Uint32()
	destination := gofakeit.Uint32()
	recipient := common.BigToHash(big.NewInt(gofakeit.Int64()))
	body := []byte(gofakeit.Sentence(gofakeit.Number(5, 15)))
	optimisticSeconds := gofakeit.Uint32()

	notaryTip := randomUint96BigInt(t)
	broadcasterTip := randomUint96BigInt(t)
	proverTip := randomUint96BigInt(t)
	executorTip := randomUint96BigInt(t)

	formattedMessage, err := messageContract.FormatMessage0(&bind.CallOpts{Context: ctx}, origin, sender, nonce, destination, recipient, optimisticSeconds, notaryTip, broadcasterTip, proverTip, executorTip, body)
	Nil(t, err)

	decodedMessage, err := types.DecodeMessage(formattedMessage)
	Nil(t, err)

	Equal(t, decodedMessage.OriginDomain(), origin)
	Equal(t, decodedMessage.Sender(), sender)
	Equal(t, decodedMessage.Nonce(), nonce)
	Equal(t, decodedMessage.DestinationDomain(), destination)
	Equal(t, decodedMessage.Body(), body)
}

func TestHeaderEncodeParity(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()

	testBackend := simulated.NewSimulatedBackend(ctx, t)
	deployManager := testutil.NewDeployManager(t)
	_, headerHarnessContract := deployManager.GetHeaderHarness(ctx, testBackend)

	domain := gofakeit.Uint32()
	sender := common.BigToHash(big.NewInt(gofakeit.Int64()))
	nonce := gofakeit.Uint32()
	destination := gofakeit.Uint32()
	recipient := common.BigToHash(big.NewInt(gofakeit.Int64()))
	optimisticSeconds := gofakeit.Uint32()

	solHeader, err := headerHarnessContract.FormatHeader(&bind.CallOpts{Context: ctx},
		domain,
		sender,
		nonce,
		destination,
		recipient,
		optimisticSeconds,
	)
	Nil(t, err)

	goHeader, err := types.EncodeHeader(types.NewHeader(domain, sender, nonce, destination, recipient, optimisticSeconds))
	Nil(t, err)

	Equal(t, goHeader, solHeader)

	headerVersion, err := headerHarnessContract.HeaderVersion(&bind.CallOpts{Context: ctx})
	Nil(t, err)

	Equal(t, headerVersion, types.HeaderVersion)
}