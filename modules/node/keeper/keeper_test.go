package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/iritamod/modules/node/keeper"
	"github.com/bianjieai/iritamod/modules/node/types"
	"github.com/bianjieai/iritamod/simapp"
	cautils "github.com/bianjieai/iritamod/utils/ca"
)

var (
	cert, _   = cautils.ReadCertificateFromMem([]byte(testCertificate))
	pubKey, _ = cautils.GetPubkeyFromCert(cert)

	testID = pubKey.Address()
)

type KeeperTestSuite struct {
	suite.Suite

	ctx    sdk.Context
	keeper *keeper.Keeper
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) SetupTest() {
	app := simapp.Setup(false)

	suite.ctx = app.BaseApp.NewContext(false, tmproto.Header{})
	suite.keeper = &app.NodeKeeper
}

func (suite *KeeperTestSuite) setNode() {
	node := types.NewNode(testID, testCertificate)
	suite.keeper.SetNode(suite.ctx, testID, node)
}

func (suite *KeeperTestSuite) TestAddNode() {
	id, err := suite.keeper.AddNode(suite.ctx, testCertificate)
	suite.NoError(err)

	node, found := suite.keeper.GetNode(suite.ctx, id)
	suite.True(found)

	suite.Equal(testID.String(), node.Id)
	suite.Equal(testCertificate, node.Certificate)
}

func (suite *KeeperTestSuite) TestRemoveNode() {
	suite.setNode()

	err := suite.keeper.RemoveNode(suite.ctx, testID)
	suite.NoError(err)

	_, found := suite.keeper.GetNode(suite.ctx, testID)
	suite.False(found)
}

const testCertificate = `-----BEGIN CERTIFICATE-----
MIIBazCCAR0CFGTwvE8oG+N3uNm1gonJBh6pie5TMAUGAytlcDBYMQswCQYDVQQG
EwJDTjENMAsGA1UECAwEcm9vdDENMAsGA1UEBwwEcm9vdDENMAsGA1UECgwEcm9v
dDENMAsGA1UECwwEcm9vdDENMAsGA1UEAwwEcm9vdDAeFw0yMDA2MTkwNzAyMzla
Fw0yMDA3MTkwNzAyMzlaMFgxCzAJBgNVBAYTAkNOMQ0wCwYDVQQIDAR0ZXN0MQ0w
CwYDVQQHDAR0ZXN0MQ0wCwYDVQQKDAR0ZXN0MQ0wCwYDVQQLDAR0ZXN0MQ0wCwYD
VQQDDAR0ZXN0MCowBQYDK2VwAyEA27WvK0goa1sSjsp6eb/xCkgjBEoPC9vfL/6h
f0hqjHYwBQYDK2VwA0EA0fo8y+saUl+8UiyKpKdjv2DsqYWqmqJDz9u3NaioOvrQ
Z0mOxdgj9wfO0t3voldCRUw3hCekjC+GEOoXH5ysDQ==
-----END CERTIFICATE-----`
