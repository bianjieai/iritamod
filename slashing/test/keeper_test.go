package test

import (
	nodekeeper "github.com/bianjieai/iritamod/node/keeper"
	"github.com/bianjieai/iritamod/node/types"
	"github.com/bianjieai/iritamod/node/utils/ca"
	"github.com/bianjieai/iritamod/simapp"
	"github.com/bianjieai/iritamod/slashing/keeper"
	slashingtype "github.com/bianjieai/iritamod/slashing/types"
	"github.com/cosmos/cosmos-sdk/codec"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/crypto/tmhash"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	"testing"
	"time"
)

const (
	CertStr = `-----BEGIN CERTIFICATE-----
MIIBazCCAR0CFGTwvE8oG+N3uNm1gonJBh6pie5TMAUGAytlcDBYMQswCQYDVQQG
EwJDTjENMAsGA1UECAwEcm9vdDENMAsGA1UEBwwEcm9vdDENMAsGA1UECgwEcm9v
dDENMAsGA1UECwwEcm9vdDENMAsGA1UEAwwEcm9vdDAeFw0yMDA2MTkwNzAyMzla
Fw0yMDA3MTkwNzAyMzlaMFgxCzAJBgNVBAYTAkNOMQ0wCwYDVQQIDAR0ZXN0MQ0w
CwYDVQQHDAR0ZXN0MQ0wCwYDVQQKDAR0ZXN0MQ0wCwYDVQQLDAR0ZXN0MQ0wCwYD
VQQDDAR0ZXN0MCowBQYDK2VwAyEA27WvK0goa1sSjsp6eb/xCkgjBEoPC9vfL/6h
f0hqjHYwBQYDK2VwA0EA0fo8y+saUl+8UiyKpKdjv2DsqYWqmqJDz9u3NaioOvrQ
Z0mOxdgj9wfO0t3voldCRUw3hCekjC+GEOoXH5ysDQ==
-----END CERTIFICATE-----`
	CertStr1 = `-----BEGIN CERTIFICATE-----
MIIBazCCAR0CFGTwvE8oG+N3uNm1gonJBh6pie5UMAUGAytlcDBYMQswCQYDVQQG
EwJDTjENMAsGA1UECAwEcm9vdDENMAsGA1UEBwwEcm9vdDENMAsGA1UECgwEcm9v
dDENMAsGA1UECwwEcm9vdDENMAsGA1UEAwwEcm9vdDAeFw0yMDA2MTkwNzA2NTFa
Fw0yMDA3MTkwNzA2NTFaMFgxCzAJBgNVBAYTAkNOMQ0wCwYDVQQIDAR0ZXN0MQ0w
CwYDVQQHDAR0ZXN0MQ0wCwYDVQQKDAR0ZXN0MQ0wCwYDVQQLDAR0ZXN0MQ0wCwYD
VQQDDAR0ZXN0MCowBQYDK2VwAyEABowXNYsnvLHjFzk93HY7+OOaQAiso8f30dw/
9jgdUIIwBQYDK2VwA0EA6dDMI3rp7VFbzoIesTy+qcISlymfR5RHuV2vl51hSapa
MlygOgSe/qRes/xvFG6ilC/v81ZuS6ll99tkEm+ZDA==
-----END CERTIFICATE-----`
)

type KeeperTestSuite struct {
	suite.Suite

	cdc        codec.Codec
	ctx        sdk.Context
	keeper     keeper.Keeper
	nodekeeper nodekeeper.Keeper
	app        *simapp.SimApp
}

var (
	name     = "test_name"
	details  = "test_details"
	power    = int64(1)
	operator = sdk.AccAddress(tmhash.SumTruncated([]byte("test_operator")))
	cert, _  = ca.ReadCertificateFromMem([]byte(CertStr))
	pk, _    = ca.GetPubkeyFromCert(cert)
	cospk, _ = cryptocodec.FromTmPubKeyInterface(pk)
	nodeID   = pk.Address()
	nodeName = "test_node"
)

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) SetupTest() {
	app := simapp.Setup(false)

	suite.cdc = app.AppCodec()
	suite.ctx = app.BaseApp.NewContext(false, tmproto.Header{
		Time: time.Now(),
	})
	suite.app = app
	suite.nodekeeper = app.NodeKeeper
	suite.keeper = keeper.NewKeeper(app.SlashingKeeper, suite.nodekeeper)
}

func (suite *KeeperTestSuite) TestSlashing() {
	msg := types.NewMsgCreateValidator(name, details, CertStr, power, operator)
	validator := tmbytes.HexBytes(tmhash.Sum(msg.GetSignBytes()))
	err := suite.nodekeeper.CreateValidator(suite.ctx,
		validator,
		msg.Name,
		msg.Certificate,
		nil,
		msg.Power,
		msg.Description,
		msg.Operator,
	)
	suite.NoError(err)
	validator1, found := suite.nodekeeper.GetValidator(suite.ctx, validator)
	suite.True(found)
	suite.Equal(validator.String(), validator1.Id)

	conAddr, err := validator1.GetConsAddr()
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
	validator1, found = suite.nodekeeper.GetValidator(suite.ctx, validator)
	suite.True(found)
	suite.True(validator1.Jailed)
	unjailmsg := slashingtype.MsgUnjailValidator{
		Id:       validator1.Id,
		Operator: validator1.Operator,
	}
	suite.ctx = suite.ctx.WithBlockHeight(int64(5000))
	suite.ctx = suite.ctx.WithBlockHeader(tmproto.Header{
		Time: time.Now().AddDate(0, 0, 1),
	})
	signInfo, found := suite.keeper.GetValidatorSigningInfo(suite.ctx, conAddr)
	suite.True(found)
	suite.NotNil(signInfo)

	err = suite.keeper.HandleUnjail(suite.ctx, unjailmsg)
	suite.NoError(err)
	validator1, found = suite.nodekeeper.GetValidator(suite.ctx, validator)
	suite.True(found)
	suite.False(validator1.Jailed)

}
