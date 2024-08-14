package test

import (
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkupgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/stretchr/testify/suite"
	"iritamod.bianjie.ai/modules/upgrade/keeper"
	"iritamod.bianjie.ai/simapp"
	"testing"
)

var isCheckTx = false

type KeeperTestSuite struct {
	suite.Suite
	ctx    sdk.Context
	keeper *keeper.Keeper
	//encCfg      TestEncodingConfig
	//queryClient upgradetypes.QueryClient
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) SetupTest() {
	depInjectOptions := simapp.DepinjectOptions{
		Config:    AppConfig,
		Providers: []interface{}{},
		Consumers: []interface{}{&suite.keeper},
	}
	app := simapp.Setup(suite.T(), isCheckTx, depInjectOptions)

	suite.ctx = app.BaseApp.NewContext(false, tmproto.Header{})
}

func (suite *KeeperTestSuite) TestUpgrade() {
	plan := sdkupgradetypes.Plan{
		Name:   "all-good",
		Info:   "some text here",
		Height: 123450000,
	}
	err := suite.keeper.ScheduleUpgrade(suite.ctx, plan)
	suite.NoError(err)
	result, has := suite.keeper.GetUpgradePlan(suite.ctx)
	suite.True(has)
	suite.Equal(plan.Name, result.Name)
	suite.Equal(plan.Height, result.Height)
	_, err = suite.keeper.ReadUpgradeInfoFromDisk()
	suite.NoError(err)
	resp, err := suite.keeper.CurrentPlan(suite.ctx, &sdkupgradetypes.QueryCurrentPlanRequest{})
	suite.NoError(err)
	suite.NotNil(resp)
	suite.Equal(plan.Name, resp.Plan.Name)
	suite.Equal(plan.Height, resp.Plan.Height)
	suite.keeper.ClearUpgradePlan(suite.ctx)
	suite.NoError(err)
	_, has = suite.keeper.GetUpgradePlan(suite.ctx)
	suite.False(has)

}
