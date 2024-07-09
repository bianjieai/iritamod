package test

import (
	"github.com/bianjieai/iritamod/modules/upgrade/keeper"
	uktype "github.com/bianjieai/iritamod/modules/upgrade/types"
	"github.com/bianjieai/iritamod/simapp"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	"testing"
)

var isCheckTx = false

type KeeperTestSuite struct {
	suite.Suite
	ctx    sdk.Context
	keeper keeper.Keeper
	//encCfg      TestEncodingConfig
	//queryClient upgradetypes.QueryClient
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) SetupTest() {
	/*homeDir := filepath.Join(suite.T().TempDir(), "x_upgrade_keeper_test")
	key := sdk.NewKVStoreKey(upgradetypes.StoreKey)
	testCtx := DefaultContextWithDB(suite.T(), key, sdk.NewTransientStoreKey("transient_test"))
	suite.ctx = testCtx.Ctx.WithBlockHeader(tmproto.Header{Time: time.Now()})
	encCfg := MakeTestEncodingConfig()
	uk := upgradekeeper.NewKeeper(make(map[int64]bool), key, encCfg.Codec, homeDir, nil)
	suite.keeper = keeper.NewKeeper(uk)
	queryHelper := baseapp.NewQueryServerTestHelper(suite.ctx, suite.encCfg.InterfaceRegistry)
	uktype.RegisterQueryServer(queryHelper, suite.keeper.UpgradeKeeper())
	suite.queryClient = uktype.NewQueryClient(queryHelper)*/
	depInjectOptions := simapp.DepinjectOptions{
		Config:    AppConfig,
		Providers: []interface{}{},
		Consumers: []interface{}{&suite.keeper},
	}
	app := simapp.Setup(suite.T(), isCheckTx, depInjectOptions)

	suite.ctx = app.BaseApp.NewContext(false, tmproto.Header{})
}

func (suite *KeeperTestSuite) TestUpgrade() {
	msg := &uktype.MsgUpgradeSoftware{
		Name:   "all-good",
		Info:   "some text here",
		Height: 123450000,
	}
	err := suite.keeper.ScheduleUpgrade(suite.ctx, msg)
	suite.NoError(err)
	err = suite.keeper.ClearUpgradePlan(suite.ctx)
	suite.NoError(err)
}
