package test

import (
	"fmt"
	"github.com/cometbft/cometbft/crypto/tmhash"
	tmbytes "github.com/cometbft/cometbft/libs/bytes"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/codec"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	nodekeeper "iritamod.bianjie.ai/modules/node/keeper"
	nodeTypes "iritamod.bianjie.ai/modules/node/types"
	"iritamod.bianjie.ai/modules/slashing/keeper"
	slashingtypes "iritamod.bianjie.ai/modules/slashing/types"
	"iritamod.bianjie.ai/simapp"
	"iritamod.bianjie.ai/simapp/utils/ca"
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

	cdc        codec.Codec
	ctx        sdk.Context
	keeper     keeper.Keeper
	nodeKeeper nodekeeper.Keeper
	//slashingKeeper cosmosslashingkeeper.Keeper
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
	depInjectOptions := simapp.DepinjectOptions{
		Config:    AppConfig,
		Providers: []interface{}{},
		Consumers: []interface{}{&suite.keeper, &suite.nodeKeeper},
	}
	app := simapp.Setup(suite.T(), isCheckTx, depInjectOptions)
	suite.ctx = app.BaseApp.NewContext(false, tmproto.Header{})
	suite.nodeKeeper.SetHooks(suite.keeper.Hooks())
}

func (suite *KeeperTestSuite) TestSlashing() {
	msg := nodeTypes.NewMsgCreateValidator(name, details, CertStr, power, operator)
	validator := tmbytes.HexBytes(tmhash.Sum(msg.GetSignBytes()))
	err := suite.nodeKeeper.CreateValidator(suite.ctx,
		validator,
		msg.Name,
		msg.Certificate,
		nil,
		msg.Power,
		msg.Description,
		msg.Operator,
	)
	suite.NoError(err)
	validator1, found := suite.nodeKeeper.GetValidator(suite.ctx, validator)
	suite.True(found)
	suite.Equal(validator.String(), validator1.Id)

	conAddr, err := validator1.GetConsAddr()
	fmt.Println(conAddr)
	suite.NoError(err)

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
	validator1, found = suite.nodeKeeper.GetValidator(suite.ctx, validator)
	suite.True(found)
	suite.True(validator1.Jailed)
	unjailmsg := slashingtypes.MsgUnjailValidator{
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
	validator1, found = suite.nodeKeeper.GetValidator(suite.ctx, validator)
	suite.True(found)
	suite.False(validator1.Jailed)

}
