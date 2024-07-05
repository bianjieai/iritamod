package test

import (
	"github.com/bianjieai/iritamod/modules/simapp"
	"github.com/bianjieai/iritamod/modules/upgrade/keeper"
	uktype "github.com/bianjieai/iritamod/modules/upgrade/types"
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
	//req := &upgradetypes.QueryCurrentPlanRequest{}
	//result, err := suite.queryClient.CurrentPlan(gocontext.Background(), req)
	//suite.Equal(msg.Height, result.Plan.Height)
	err = suite.keeper.ClearUpgradePlan(suite.ctx)
	suite.NoError(err)
	//result, err = suite.queryClient.CurrentPlan(gocontext.Background(), req)
	//suite.NoError(err)
	//suite.Nil(result.Plan)
}

//type TestEncodingConfig struct {
//	InterfaceRegistry types.InterfaceRegistry
//	Codec             codec.Codec
//	TxConfig          client.TxConfig
//	Amino             *codec.LegacyAmino
//}
//
//func MakeTestEncodingConfig(modules ...module.AppModuleBasic) TestEncodingConfig {
//	cdc := codec.NewLegacyAmino()
//	interfaceRegistry := types.NewInterfaceRegistry()
//	codec := codec.NewProtoCodec(interfaceRegistry)
//
//	encCfg := TestEncodingConfig{
//		InterfaceRegistry: interfaceRegistry,
//		Codec:             codec,
//		TxConfig:          tx.NewTxConfig(codec, tx.DefaultSignModes),
//		Amino:             cdc,
//	}
//
//	mb := module.NewBasicManager(modules...)
//
//	std.RegisterLegacyAminoCodec(encCfg.Amino)
//	std.RegisterInterfaces(encCfg.InterfaceRegistry)
//	mb.RegisterLegacyAminoCodec(encCfg.Amino)
//	mb.RegisterInterfaces(encCfg.InterfaceRegistry)
//
//	return encCfg
//}
//
//type TestContext struct {
//	Ctx sdk.Context
//	DB  *dbm.MemDB
//	CMS store.CommitMultiStore
//}
//
//func DefaultContextWithDB(t *testing.T, key storetypes.StoreKey, tkey storetypes.StoreKey) TestContext {
//	db := dbm.NewMemDB()
//	cms := store.NewCommitMultiStore(db)
//	cms.MountStoreWithDB(key, storetypes.StoreTypeIAVL, db)
//	cms.MountStoreWithDB(tkey, storetypes.StoreTypeTransient, db)
//	err := cms.LoadLatestVersion()
//	assert.NoError(t, err)
//
//	ctx := sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger())
//
//	return TestContext{ctx, db, cms}
//}
