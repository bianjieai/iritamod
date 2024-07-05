package test

import (
	"github.com/bianjieai/iritamod/modules/simapp"
	"github.com/bianjieai/iritamod/modules/simapp/utils/ca"
	"github.com/bianjieai/iritamod/modules/slashing/keeper"
	slashingtype "github.com/bianjieai/iritamod/modules/slashing/types"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/codec"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

const CertStr = `-----BEGIN CERTIFICATE-----
MIIBazCCAR0CFGTwvE8oG+N3uNm1gonJBh6pie5TMAUGAytlcDBYMQswCQYDVQQG
EwJDTjENMAsGA1UECAwEcm9vdDENMAsGA1UEBwwEcm9vdDENMAsGA1UECgwEcm9v
dDENMAsGA1UECwwEcm9vdDENMAsGA1UEAwwEcm9vdDAeFw0yMDA2MTkwNzAyMzla
Fw0yMDA3MTkwNzAyMzlaMFgxCzAJBgNVBAYTAkNOMQ0wCwYDVQQIDAR0ZXN0MQ0w
CwYDVQQHDAR0ZXN0MQ0wCwYDVQQKDAR0ZXN0MQ0wCwYDVQQLDAR0ZXN0MQ0wCwYD
VQQDDAR0ZXN0MCowBQYDK2VwAyEA27WvK0goa1sSjsp6eb/xCkgjBEoPC9vfL/6h
f0hqjHYwBQYDK2VwA0EA0fo8y+saUl+8UiyKpKdjv2DsqYWqmqJDz9u3NaioOvrQ
Z0mOxdgj9wfO0t3voldCRUw3hCekjC+GEOoXH5ysDQ==
-----END CERTIFICATE-----`

type KeeperTestSuite struct {
	suite.Suite

	cdc    codec.Codec
	ctx    sdk.Context
	keeper keeper.Keeper
	//nodekeeper nodekeeper.Keeper
	app *simapp.SimApp
}

var (
	isCheckTx   = false
	name        = "test_name"
	details     = "test_details"
	power       = int64(1)
	cert, _     = ca.ReadCertificateFromMem([]byte(CertStr))
	pk, _       = ca.GetPubkeyFromCert(cert)
	cospk, _    = cryptocodec.FromTmPubKeyInterface(pk)
	nodeID      = pk.Address()
	operator, _ = sdk.AccAddressFromHexUnsafe(pk.Address().String())
	nodeName    = "test_node"
	consAddr    = sdk.ConsAddress(operator)
)

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) SetupTest() {
	//app := simapp.Setup(false)
	//
	//suite.cdc = app.AppCodec()
	//suite.ctx = app.BaseApp.NewContext(false, tmproto.Header{
	//	Time: time.Now(),
	//})
	//suite.app = app
	//suite.nodekeeper = app.NodeKeeper
	//suite.keeper = keeper.NewKeeper(app.SlashingKeeper, suite.nodekeeper)
	depInjectOptions := simapp.DepinjectOptions{
		Config:    AppConfig,
		Providers: []interface{}{},
		Consumers: []interface{}{&suite.keeper},
	}
	app := simapp.Setup(suite.T(), isCheckTx, depInjectOptions)

	suite.ctx = app.BaseApp.NewContext(false, tmproto.Header{})
}

func (suite *KeeperTestSuite) TestSlashing() {
	//msg := types.NewMsgCreateValidator(name, details, CertStr, power, operator)
	//validator := tmbytes.HexBytes(tmhash.Sum(msg.GetSignBytes()))
	//err := suite.nodekeeper.CreateValidator(suite.ctx,
	//	validator,
	//	msg.Name,
	//	msg.Certificate,
	//	nil,
	//	msg.Power,
	//	msg.Description,
	//	msg.Operator,
	//)
	//suite.NoError(err)
	//validator1, found := suite.nodekeeper.GetValidator(suite.ctx, validator)
	//suite.True(found)
	//suite.Equal(validator.String(), validator1.Id)
	//
	//conAddr, err := validator1.GetConsAddr()
	//suite.NoError(err)

	height := int64(0)
	target := int64(100)
	for ; height < target; height++ {
		suite.ctx = suite.ctx.WithBlockHeight(height)
		suite.keeper.HandleValidatorSignature(suite.ctx, nodeID, int64(1), true)
	}

	for ; height < target+500; height++ {
		suite.ctx = suite.ctx.WithBlockHeight(height)
		suite.keeper.HandleValidatorSignature(suite.ctx, nodeID, int64(1), false)
	}
	//validator1, found = suite.nodekeeper.GetValidator(suite.ctx, validator)
	//suite.True(found)
	//suite.True(validator1.Jailed)
	unjailmsg := slashingtype.MsgUnjailValidator{
		Id:       "1",
		Operator: operator.String(),
	}
	suite.ctx = suite.ctx.WithBlockHeight(int64(5000))
	suite.ctx = suite.ctx.WithBlockHeader(tmproto.Header{
		Time: time.Now().AddDate(0, 0, 1),
	})
	signInfo, found := suite.keeper.GetValidatorSigningInfo(suite.ctx, consAddr)
	suite.True(found)
	suite.NotNil(signInfo)

	err = suite.keeper.HandleUnjail(suite.ctx, unjailmsg)
	suite.NoError(err)
	validator1, found = suite.nodekeeper.GetValidator(suite.ctx, validator)
	suite.True(found)
	suite.False(validator1.Jailed)

}
