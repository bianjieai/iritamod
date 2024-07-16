package test

import (
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkupgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/stretchr/testify/suite"
	"iritamod.bianjie.ai/modules/upgrade/keeper"
	uktype "iritamod.bianjie.ai/modules/upgrade/types"
	"iritamod.bianjie.ai/simapp"
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
	plan, has := suite.keeper.GetUpgradePlan(suite.ctx)
	suite.True(has)
	suite.Equal(msg.Name, plan.Name)
	suite.Equal(msg.Height, plan.Height)
	plan, err = suite.keeper.ReadUpgradeInfoFromDisk()
	suite.NoError(err)
	resp, err := suite.keeper.CurrentPlan(suite.ctx, &sdkupgradetypes.QueryCurrentPlanRequest{})
	suite.NoError(err)
	suite.NotNil(resp)
	suite.Equal(msg.Name, resp.Plan.Name)
	suite.Equal(msg.Height, resp.Plan.Height)
	err = suite.keeper.ClearUpgradePlan(suite.ctx)
	suite.NoError(err)

}
