package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/bianjieai/iritamod/modules/wevm/keeper"
	"github.com/bianjieai/iritamod/simapp"

	"github.com/stretchr/testify/suite"
)

var (
	testID = "irita_10086-1"
)

type KeeperTestSuite struct {
	suite.Suite

	Ctx    sdk.Context
	Keeper keeper.Keeper
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) SetupTest() {
	app := simapp.Setup(false)

	suite.Keeper = app.WevmKeeper
	suite.Ctx = app.BaseApp.NewContext(false, tmproto.Header{})
}
