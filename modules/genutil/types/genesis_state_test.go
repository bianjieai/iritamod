package types_test

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/std"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"cosmossdk.io/simapp"
	simappparams "cosmossdk.io/simapp/params"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/genutil/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

var (
	pk1 = ed25519.GenPrivKey().PubKey()
	pk2 = ed25519.GenPrivKey().PubKey()
)

// GenTxTestSuite is a test suite to be used with gentx tests.
//type GenTxTestSuite struct {
//	suite.Suite
//
//	ctx            sdk.Context
//	app            *simapp.SimApp
//	encodingConfig simappparams.EncodingConfig
//
//	msg1, msg2 *stakingtypes.MsgCreateValidator
//}
//
//func (suite *GenTxTestSuite) SetupTest() {
//	checkTx := false
//	app := simapp.Setup(checkTx)
//	suite.ctx = app.BaseApp.NewContext(checkTx, tmproto.Header{})
//	suite.app = app
//	suite.encodingConfig = simapp.MakeTestEncodingConfig()
//
//}

func TestNewGenesisState(t *testing.T) {
	gen := types.NewGenesisState(nil)
	assert.NotNil(t, gen.GenTxs) // https://github.com/cosmos/cosmos-sdk/issues/5086

	gen = types.NewGenesisState([]json.RawMessage{[]byte(`{"foo":"bar"}`)})
	assert.Equal(t, string(gen.GenTxs[0]), `{"foo":"bar"}`)
}

func TestValidateGenesisMultipleMessages(t *testing.T) {
	desc := stakingtypes.NewDescription("testname", "", "", "", "")
	comm := stakingtypes.CommissionRates{}

	msg1, err := stakingtypes.NewMsgCreateValidator(
		sdk.ValAddress(pk1.Address()), pk1, sdk.NewInt64Coin(sdk.DefaultBondDenom, 50),
		desc, comm, sdk.OneInt(),
	)
	require.NoError(t, err)

	msg2, err := stakingtypes.NewMsgCreateValidator(
		sdk.ValAddress(pk2.Address()), pk2, sdk.NewInt64Coin(sdk.DefaultBondDenom, 50),
		desc, comm, sdk.OneInt(),
	)
	require.NoError(t, err)

	simConfig := simappparams.MakeTestEncodingConfig()
	std.RegisterLegacyAminoCodec(simConfig.Amino)
	std.RegisterInterfaces(simConfig.InterfaceRegistry)
	simapp.ModuleBasics.RegisterLegacyAminoCodec(simConfig.Amino)
	simapp.ModuleBasics.RegisterInterfaces(simConfig.InterfaceRegistry)
	txGen := simConfig.TxConfig

	txBuilder := txGen.NewTxBuilder()
	require.NoError(t, txBuilder.SetMsgs(msg1, msg2))

	tx := txBuilder.GetTx()
	genesisState := types.NewGenesisStateFromTx(txGen.TxJSONEncoder(), []sdk.Tx{tx})

	err = types.ValidateGenesis(genesisState, simappparams.MakeTestEncodingConfig().TxConfig.TxJSONDecoder(), types.DefaultMessageValidator)
	require.Error(t, err)
}

func TestValidateGenesisBadMessage(t *testing.T) {
	desc := stakingtypes.NewDescription("testname", "", "", "", "")

	msg1 := stakingtypes.NewMsgEditValidator(sdk.ValAddress(pk1.Address()), desc, nil, nil)

	simConfig := simappparams.MakeTestEncodingConfig()
	std.RegisterLegacyAminoCodec(simConfig.Amino)
	std.RegisterInterfaces(simConfig.InterfaceRegistry)
	simapp.ModuleBasics.RegisterLegacyAminoCodec(simConfig.Amino)
	simapp.ModuleBasics.RegisterInterfaces(simConfig.InterfaceRegistry)
	txGen := simConfig.TxConfig

	txBuilder := txGen.NewTxBuilder()
	err := txBuilder.SetMsgs(msg1)
	require.NoError(t, err)
	tx := txBuilder.GetTx()
	genesisState := types.NewGenesisStateFromTx(txGen.TxJSONEncoder(), []sdk.Tx{tx})

	err = types.ValidateGenesis(genesisState, simappparams.MakeTestEncodingConfig().TxConfig.TxJSONDecoder(), types.DefaultMessageValidator)
	require.Error(t, err)
}
