package simapp

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	nodetypes "github.com/bianjieai/iritamod/modules/node/types"
	"github.com/cometbft/cometbft/crypto/tmhash"
	ctmbytes "github.com/cometbft/cometbft/libs/bytes"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/testutil/mock"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/stretchr/testify/require"
	"math/rand"
	"strconv"
	"testing"
	"time"

	dbm "github.com/cometbft/cometbft-db"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/log"
	ctmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	ctmtypes "github.com/cometbft/cometbft/types"

	"github.com/bianjieai/iritamod/modules/node"
	"github.com/bianjieai/iritamod/modules/perm"
	bam "github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
)

func setup(withGenesis bool, invCheckPeriod uint) (*SimApp, GenesisState) {
	db := dbm.NewMemDB()

	appOptions := make(simtestutil.AppOptionsMap, 0)
	appOptions[flags.FlagHome] = DefaultNodeHome
	appOptions[server.FlagIndexEvents] = invCheckPeriod

	app := NewSimApp(log.NewNopLogger(), db, nil, true, map[int64]bool{}, EmptyAppOptions{})
	if withGenesis {
		return app, app.DefaultGenesis()
	}

	return app, GenesisState{}
}

// Setup initializes a new SimApp. A Nop logger is set in SimApp.
func Setup(t *testing.T, isCheckTx bool) *SimApp {
	t.Helper()

	privVal := mock.NewPV()
	pubKey, err := privVal.GetPubKey()
	require.NoError(t, err)

	// create validator set with single validator
	validator := ctmtypes.NewValidator(pubKey, 1)
	valSet := ctmtypes.NewValidatorSet([]*ctmtypes.Validator{validator})

	app := SetupWithConsortiumGenesisValSet(t, valSet)

	return app
}

func SetupWithConsortiumGenesisValSet(t *testing.T, valSet *ctmtypes.ValidatorSet) *SimApp {
	t.Helper()
	app, genesisState := setup(true, 5)

	// add root admin
	permGenState := perm.GetGenesisStateFromAppState(app.appCodec, genesisState)
	permGenState.RoleAccounts = append(
		permGenState.RoleAccounts,
		perm.RoleAccount{
			Address: rootAdmin,
			Roles:   []perm.Role{perm.RoleRootAdmin},
		},
	)
	permGenStateBz := app.legacyAmino.MustMarshalJSON(permGenState)
	genesisState[perm.ModuleName] = permGenStateBz

	// create validator
	validators := make([]nodetypes.Validator, 0, len(valSet.Validators))
	for _, val := range valSet.Validators {
		pk, err := cryptocodec.FromTmPubKeyInterface(val.PubKey)
		if err != nil {
			panic("failed to convert pubkey")
		}

		pkStr, err := bech32.ConvertAndEncode(sdk.GetConfig().GetBech32ConsensusPubPrefix(), legacy.Cdc.MustMarshal(pk))
		if err != nil {
			panic("failed to convert pubkey string")
		}

		validator := nodetypes.Validator{
			Id:          ctmbytes.HexBytes(tmhash.Sum([]byte("root"))).String(),
			Name:        "genesisName",
			Pubkey:      pkStr,
			Certificate: rootCert,
			Power:       val.VotingPower,
			Description: "initial validator",
			Jailed:      false,
			Operator:    sdk.ValAddress(val.Address).String(),
		}
		validators = append(validators, validator)
	}

	// add root cert
	validatorGenState := node.GetGenesisStateFromAppState(app.appCodec, genesisState)
	validatorGenState.RootCert = rootCert
	validatorGenState.Validators = validators

	validatorGenStateBz := app.legacyAmino.MustMarshalJSON(validatorGenState)
	genesisState[node.ModuleName] = validatorGenStateBz

	stateBytes, err := codec.MarshalJSONIndent(app.legacyAmino, genesisState)
	require.NoError(t, err)

	// Initialize the chain
	app.InitChain(abci.RequestInitChain{
		Validators:      []abci.ValidatorUpdate{},
		ConsensusParams: simtestutil.DefaultConsensusParams,
		AppStateBytes:   stateBytes,
	})

	return app
}

// SetupWithGenesisValSet initializes a new SimApp with a validator set and genesis accounts
// that also act as delegators. For simplicity, each validator is bonded with a delegation
// of one consensus engine unit (10^6) in the default token of the simapp from first genesis
// account. A Nop logger is set in SimApp.
func SetupWithGenesisValSet(t *testing.T, valSet *ctmtypes.ValidatorSet, genAccs []authtypes.GenesisAccount, balances ...banktypes.Balance) *SimApp {
	t.Helper()

	app, genesisState := setup(true, 0)
	genesisState, err := simtestutil.GenesisStateWithValSet(app.AppCodec(), genesisState, valSet, genAccs, balances...)
	require.NoError(t, err)

	stateBytes, err := json.MarshalIndent(genesisState, "", " ")
	require.NoError(t, err)

	// init chain will set the validator set and initialize the genesis accounts
	app.InitChain(
		abci.RequestInitChain{
			Validators:      []abci.ValidatorUpdate{},
			ConsensusParams: simtestutil.DefaultConsensusParams,
			AppStateBytes:   stateBytes,
		},
	)

	// commit genesis changes
	app.Commit()
	app.BeginBlock(abci.RequestBeginBlock{Header: ctmproto.Header{
		Height:             app.LastBlockHeight() + 1,
		AppHash:            app.LastCommitID().Hash,
		ValidatorsHash:     valSet.Hash(),
		NextValidatorsHash: valSet.Hash(),
	}})

	return app
}

type GenerateAccountStrategy func(int) []sdk.AccAddress

// createRandomAccounts is a strategy used by addTestAddrs() in order to generated addresses in random order.
func createRandomAccounts(accNum int) []sdk.AccAddress {
	testAddrs := make([]sdk.AccAddress, accNum)
	for i := 0; i < accNum; i++ {
		pk := ed25519.GenPrivKey().PubKey()
		testAddrs[i] = sdk.AccAddress(pk.Address())
	}

	return testAddrs
}

// createIncrementalAccounts is a strategy used by addTestAddrs() in order to generated addresses in ascending order.
func createIncrementalAccounts(accNum int) []sdk.AccAddress {
	var addresses []sdk.AccAddress
	var buffer bytes.Buffer

	// start at 100 so we can make up to 999 test addresses with valid test addresses
	for i := 100; i < (accNum + 100); i++ {
		numString := strconv.Itoa(i)
		buffer.WriteString("A58856F0FD53BF058B4909A21AEC019107BA6") //base address string

		buffer.WriteString(numString) //adding on final two digits to make addresses unique
		res, _ := sdk.AccAddressFromHexUnsafe(buffer.String())
		bech := res.String()
		addr, _ := TestAddr(buffer.String(), bech)

		addresses = append(addresses, addr)
		buffer.Reset()
	}

	return addresses
}

// AddTestAddrs constructs and returns accNum amount of accounts with an
// initial balance of accAmt in random order
func AddTestAddrs(app *SimApp, ctx sdk.Context, accNum int, accAmt sdk.Int) []sdk.AccAddress {
	return addTestAddrs(app, ctx, accNum, accAmt, createRandomAccounts)
}

// AddTestAddrs constructs and returns accNum amount of accounts with an
// initial balance of accAmt in random order
func AddTestAddrsIncremental(app *SimApp, ctx sdk.Context, accNum int, accAmt sdk.Int) []sdk.AccAddress {
	return addTestAddrs(app, ctx, accNum, accAmt, createIncrementalAccounts)
}

func addTestAddrs(app *SimApp, ctx sdk.Context, accNum int, accAmt sdk.Int, strategy GenerateAccountStrategy) []sdk.AccAddress {
	testAddrs := strategy(accNum)

	initCoins := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, accAmt))

	for _, addr := range testAddrs {
		initAccountWithCoins(app, ctx, addr, initCoins)
	}

	return testAddrs
}

func initAccountWithCoins(app *SimApp, ctx sdk.Context, addr sdk.AccAddress, coins sdk.Coins) {
	if err := app.BankKeeper.MintCoins(ctx, minttypes.ModuleName, coins); err != nil {
		panic(err)
	}

	if err := app.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, addr, coins); err != nil {
		panic(err)
	}
}

// ConvertAddrsToValAddrs converts the provided addresses to ValAddress.
func ConvertAddrsToValAddrs(addrs []sdk.AccAddress) []sdk.ValAddress {
	valAddrs := make([]sdk.ValAddress, len(addrs))

	for i, addr := range addrs {
		valAddrs[i] = sdk.ValAddress(addr)
	}

	return valAddrs
}

func TestAddr(addr string, bech string) (sdk.AccAddress, error) {
	res, err := sdk.AccAddressFromHexUnsafe(addr)
	if err != nil {
		return nil, err
	}
	bechexpected := res.String()
	if bech != bechexpected {
		return nil, fmt.Errorf("bech encoding doesn't match reference")
	}

	bechres, err := sdk.AccAddressFromBech32(bech)
	if err != nil {
		return nil, err
	}
	if !bytes.Equal(bechres, res) {
		return nil, err
	}

	return res, nil
}

// CheckBalance checks the balance of an account.
func CheckBalance(t *testing.T, app *SimApp, addr sdk.AccAddress, balances sdk.Coins) {
	ctxCheck := app.BaseApp.NewContext(true, ctmproto.Header{})
	require.True(t, balances.IsEqual(app.BankKeeper.GetAllBalances(ctxCheck, addr)))
}

// SignCheckDeliver checks a generated signed transaction and simulates a
// block commitment with the given transaction. A test assertion is made using
// the parameter 'expPass' against the result. A corresponding result is
// returned.
func SignCheckDeliver(
	t *testing.T, txCfg client.TxConfig, app *bam.BaseApp, header ctmproto.Header, msgs []sdk.Msg,
	chainID string, accNums, accSeqs []uint64, expSimPass, expPass bool, priv ...cryptotypes.PrivKey,
) (sdk.GasInfo, *sdk.Result, error) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	tx, err := simtestutil.GenSignedMockTx(
		r,
		txCfg,
		msgs,
		sdk.Coins{sdk.NewInt64Coin(sdk.DefaultBondDenom, 0)},
		simtestutil.DefaultGenTxGas,
		chainID,
		accNums,
		accSeqs,
		priv...,
	)
	require.NoError(t, err)
	txBytes, err := txCfg.TxEncoder()(tx)
	require.Nil(t, err)

	// Must simulate now as CheckTx doesn't run Msgs anymore
	_, res, err := app.Simulate(txBytes)

	if expSimPass {
		require.NoError(t, err)
		require.NotNil(t, res)
	} else {
		require.Error(t, err)
		require.Nil(t, res)
	}

	// Simulate a sending a transaction and committing a block
	app.BeginBlock(abci.RequestBeginBlock{Header: header})
	gInfo, res, err := app.SimDeliver(txCfg.TxEncoder(), tx)

	if expPass {
		require.NoError(t, err)
		require.NotNil(t, res)
	} else {
		require.Error(t, err)
		require.Nil(t, res)
	}

	app.EndBlock(abci.RequestEndBlock{})
	app.Commit()

	return gInfo, res, err
}

// GenSequenceOfTxs generates a set of signed transactions of messages, such
// that they differ only by having the sequence numbers incremented between
// every transaction.
func GenSequenceOfTxs(
	txGen client.TxConfig, msgs []sdk.Msg, accNums []uint64,
	initSeqNums []uint64, numToGenerate int, priv ...cryptotypes.PrivKey,
) ([]sdk.Tx, error) {
	txs := make([]sdk.Tx, numToGenerate)
	var err error
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < numToGenerate; i++ {
		txs[i], err = simtestutil.GenSignedMockTx(
			r,
			txGen,
			msgs,
			sdk.Coins{sdk.NewInt64Coin(sdk.DefaultBondDenom, 0)},
			simtestutil.DefaultGenTxGas,
			"",
			accNums,
			initSeqNums,
			priv...,
		)
		if err != nil {
			break
		}
		incrementAllSequenceNumbers(initSeqNums)
	}

	return txs, err
}

func incrementAllSequenceNumbers(initSeqNums []uint64) {
	for i := 0; i < len(initSeqNums); i++ {
		initSeqNums[i]++
	}
}

// CreateTestPubKeys returns a total of numPubKeys public keys in ascending order.
func CreateTestPubKeys(numPubKeys int) []cryptotypes.PubKey {
	var publicKeys []cryptotypes.PubKey
	var buffer bytes.Buffer

	// start at 10 to avoid changing 1 to 01, 2 to 02, etc
	for i := 100; i < (numPubKeys + 100); i++ {
		numString := strconv.Itoa(i)
		buffer.WriteString("0B485CFC0EECC619440448436F8FC9DF40566F2369E72400281454CB552AF") // base pubkey string
		buffer.WriteString(numString)                                                       // adding on final two digits to make pubkeys unique
		publicKeys = append(publicKeys, NewPubKeyFromHex(buffer.String()))
		buffer.Reset()
	}

	return publicKeys
}

// NewPubKeyFromHex returns a PubKey from a hex string.
func NewPubKeyFromHex(pk string) (res cryptotypes.PubKey) {
	pkBytes, err := hex.DecodeString(pk)
	if err != nil {
		panic(err)
	}
	if len(pkBytes) != ed25519.PubKeySize {
		panic(errors.Wrap(errors.ErrInvalidPubKey, "invalid pubkey size"))
	}
	return &ed25519.PubKey{Key: pkBytes}
}

// EmptyAppOptions is a stub implementing AppOptions
type EmptyAppOptions struct{}

// Get implements AppOptions
func (ao EmptyAppOptions) Get(o string) interface{} {
	return nil
}

const rootCert = `-----BEGIN CERTIFICATE-----
MIIBxTCCAXegAwIBAgIUHMPutrm+7FT7fIFf2fEgyQnIg8kwBQYDK2VwMFgxCzAJ
BgNVBAYTAkNOMQ0wCwYDVQQIDARyb290MQ0wCwYDVQQHDARyb290MQ0wCwYDVQQK
DARyb290MQ0wCwYDVQQLDARyb290MQ0wCwYDVQQDDARyb290MB4XDTIwMDYxOTA3
MDExMVoXDTIxMDYxOTA3MDExMVowWDELMAkGA1UEBhMCQ04xDTALBgNVBAgMBHJv
b3QxDTALBgNVBAcMBHJvb3QxDTALBgNVBAoMBHJvb3QxDTALBgNVBAsMBHJvb3Qx
DTALBgNVBAMMBHJvb3QwKjAFBgMrZXADIQDdzGFcck4I7Wa1vRj4JsdQ9RjVgH92
7iOhXJ8mFLwQKaNTMFEwHQYDVR0OBBYEFPrjTGR+/g4RUduZ9E8JSXNyI4mzMB8G
A1UdIwQYMBaAFPrjTGR+/g4RUduZ9E8JSXNyI4mzMA8GA1UdEwEB/wQFMAMBAf8w
BQYDK2VwA0EAT8EG5nGxwCAP4ZlfQvAhrnJI+SojlsOoE3rZ8W6/knZsrnVb6RI8
QAVleeE0pMY+MtENXcQ2wH0QRXs+wO0XCw==
-----END CERTIFICATE-----`
