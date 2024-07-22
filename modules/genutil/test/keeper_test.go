package test

import (
	abci "github.com/cometbft/cometbft/abci/types"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/stretchr/testify/suite"
	"iritamod.bianjie.ai/modules/genutil"
	"iritamod.bianjie.ai/modules/genutil/types"
	"iritamod.bianjie.ai/simapp"
	"testing"
)

var (
	isCheckTx = false
)

type KeeperTestSuite struct {
	suite.Suite
	ctx           sdk.Context
	accountkeeper authkeeper.AccountKeeper
	app           *simapp.SimApp
	AccountKeeper types.AccountKeeper
	NodeKeeper    types.NodeKeeper
	DeliverTx     func(abci.RequestDeliverTx) abci.ResponseDeliverTx
	TxConfig      client.TxConfig
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) SetupTest() {
	depInjectOptions := simapp.DepinjectOptions{
		Config:    AppConfig,
		Providers: []interface{}{},
		Consumers: []interface{}{&suite.accountkeeper, &suite.AccountKeeper, &suite.NodeKeeper, &suite.DeliverTx, &suite.TxConfig},
	}
	app := simapp.Setup(suite.T(), isCheckTx, depInjectOptions)

	suite.app = app
	suite.ctx = app.BaseApp.NewContext(false, tmproto.Header{})
}

func (suite *KeeperTestSuite) Test_genutils() {
	acc := suite.accountkeeper.GetModuleAccount(suite.ctx, minttypes.ModuleName)
	suite.NotNil(acc)
	modules := genutil.NewAppModule(suite.AccountKeeper, suite.NodeKeeper, suite.DeliverTx, suite.TxConfig)
	defaultjson := modules.DefaultGenesis(suite.app.AppCodec())
	suite.NotNil(defaultjson)
	jsonresult := modules.ExportGenesis(suite.ctx, suite.app.AppCodec())
	suite.NotNil(jsonresult)
}
