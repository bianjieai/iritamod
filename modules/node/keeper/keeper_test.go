package keeper_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec/legacy"
	"github.com/cosmos/cosmos-sdk/types/bech32"

	"github.com/stretchr/testify/suite"

	"github.com/tendermint/tendermint/crypto/tmhash"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/codec"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/iritamod/modules/node/keeper"
	"github.com/bianjieai/iritamod/modules/node/types"
	"github.com/bianjieai/iritamod/simapp"
	cautil "github.com/bianjieai/iritamod/utils/ca"
)

type KeeperTestSuite struct {
	suite.Suite

	cdc    codec.Codec
	ctx    sdk.Context
	keeper *keeper.Keeper
	app    *simapp.SimApp
}

var (
	name     = "test_name"
	details  = "test_details"
	power    = int64(1)
	operator = sdk.AccAddress(tmhash.SumTruncated([]byte("test_operator")))
	cert, _  = cautil.ReadCertificateFromMem([]byte(certStr))
	pk, _    = cautil.GetPubkeyFromCert(cert)
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
	suite.ctx = app.BaseApp.NewContext(false, tmproto.Header{})
	suite.app = app
	suite.keeper = &app.NodeKeeper
}

func (suite *KeeperTestSuite) setNode() {
	node := types.NewNode(nodeID, nodeName, certStr)
	suite.keeper.SetNode(suite.ctx, nodeID, node)
}

func (suite *KeeperTestSuite) TestCreateValidator() {
	msg := types.NewMsgCreateValidator(name, details, certStr, power, operator)
	id, err := suite.keeper.CreateValidator(suite.ctx, *msg)
	suite.NoError(err)

	validator, found := suite.keeper.GetValidator(suite.ctx, id)
	suite.True(found)
	suite.Equal(id.String(), validator.Id)
	suite.Equal(msg.Name, validator.Name)
	suite.Equal(msg.Certificate, validator.Certificate)
	suite.Equal(msg.Power, validator.Power)
	suite.Equal(msg.Description, validator.Description)
	suite.Equal(msg.Operator, validator.Operator)
	suite.False(validator.Jailed)

	validator1, found := suite.keeper.GetValidatorByConsAddr(suite.ctx, sdk.GetConsAddress(cospk))
	suite.True(found)
	suite.Equal(validator, validator1)

	validators := suite.keeper.GetAllValidators(suite.ctx)
	suite.Equal(1, len(validators))
	suite.Equal(validator, validators[0])

	suite.keeper.IterateUpdateValidators(
		suite.ctx,
		func(index int64, pubkey string, power int64) bool {
			pkStr, err := bech32.ConvertAndEncode(sdk.GetConfig().GetBech32ConsensusPubPrefix(), legacy.Cdc.MustMarshal(cospk))
			suite.Suite.NoError(err)
			suite.Equal(int64(0), index)
			suite.Equal(pkStr, pubkey)
			suite.Equal(msg.Power, power)
			return false
		},
	)
}

func (suite *KeeperTestSuite) TestUpdateValidator() {
	msg := types.NewMsgCreateValidator(name, details, certStr, power, operator)
	id, err := suite.keeper.CreateValidator(suite.ctx, *msg)
	suite.NoError(err)

	_, found := suite.keeper.GetValidator(suite.ctx, id)
	suite.True(found)

	name1 := "test_name1"
	details1 := "test_details1"
	power1 := int64(2)
	operator1 := sdk.AccAddress("test_operator1")

	cert1, err := cautil.ReadCertificateFromMem([]byte(certStr1))
	suite.NoError(err)
	pk1, err := cautil.GetPubkeyFromCert(cert1)
	suite.NoError(err)
	cospk1, err := cryptocodec.FromTmPubKeyInterface(pk1)
	suite.NoError(err)

	// error name
	msg1 := types.NewMsgUpdateValidator([]byte{0x1}, name1, details1, certStr1, power1, operator1)
	err = suite.keeper.UpdateValidator(suite.ctx, *msg1)
	suite.Error(err)

	msg2 := types.NewMsgUpdateValidator(id, "", details1, certStr1, power1, operator1)
	err = suite.keeper.UpdateValidator(suite.ctx, *msg2)
	suite.NoError(err)

	validator, found := suite.keeper.GetValidator(suite.ctx, id)
	suite.True(found)
	suite.Equal(msg2.Id, validator.Id)
	suite.Equal(msg2.Certificate, validator.Certificate)
	suite.Equal(msg2.Power, validator.Power)
	suite.Equal(msg2.Operator, validator.Operator)
	suite.Equal(msg2.Description, validator.Description)

	// old pubkey index can not be found
	_, found = suite.keeper.GetValidatorByConsAddr(suite.ctx, sdk.GetConsAddress(cospk))
	suite.False(found)

	validator1, found := suite.keeper.GetValidatorByConsAddr(suite.ctx, sdk.GetConsAddress(cospk1))
	suite.True(found)
	suite.Equal(validator, validator1)

	validators := suite.keeper.GetAllValidators(suite.ctx)
	suite.Equal(1, len(validators))
	suite.Equal(validator, validators[0])

	updatesTotal := 0
	suite.keeper.IterateUpdateValidators(suite.ctx, func(index int64, pubkey string, power int64) bool {
		pkStr, err := bech32.ConvertAndEncode(sdk.GetConfig().GetBech32ConsensusPubPrefix(), legacy.Cdc.MustMarshal(cospk))
		suite.Suite.NoError(err)
		pkStr1, err1 := bech32.ConvertAndEncode(sdk.GetConfig().GetBech32ConsensusPubPrefix(), legacy.Cdc.MustMarshal(cospk1))
		suite.Suite.NoError(err1)

		switch pubkey {
		case pkStr:
			updatesTotal++
			suite.Equal(int64(0), power)
		case pkStr1:
			updatesTotal++
			suite.Equal(msg1.Power, power)
		default:
			panic("unexpected case")
		}
		return false
	})
	suite.Equal(2, updatesTotal)
}

func (suite *KeeperTestSuite) TestRemoveValidator() {
	msg := types.NewMsgCreateValidator(name, details, certStr, power, operator)
	id, err := suite.keeper.CreateValidator(suite.ctx, *msg)
	suite.NoError(err)

	_, found := suite.keeper.GetValidator(suite.ctx, id)
	suite.True(found)

	msg1 := types.NewMsgRemoveValidator(id, operator)
	err = suite.keeper.RemoveValidator(suite.ctx, *msg1)
	suite.NoError(err)

	_, found = suite.keeper.GetValidator(suite.ctx, id)
	suite.False(found)

	_, found = suite.keeper.GetValidatorByConsAddr(suite.ctx, sdk.GetConsAddress(cospk))
	suite.False(found)

	validators := suite.keeper.GetAllValidators(suite.ctx)
	suite.Equal(0, len(validators))

	suite.keeper.IterateUpdateValidators(
		suite.ctx,
		func(index int64, pubkey string, power int64) bool {
			bz, err := sdk.GetFromBech32(pubkey, sdk.GetConfig().GetBech32ConsensusPubPrefix())
			_, err = legacy.PubKeyFromBytes(bz)
			if err != nil {
				panic(err)
			}
			suite.Suite.NoError(err)
			suite.Equal(int64(0), index)
			suite.Equal(int64(0), power)
			return false
		},
	)
}

func (suite *KeeperTestSuite) TestAddNode() {
	id, err := suite.keeper.AddNode(suite.ctx, nodeName, certStr)
	suite.NoError(err)

	node, found := suite.keeper.GetNode(suite.ctx, id)
	suite.True(found)

	suite.Equal(nodeID.String(), node.Id)
	suite.Equal(certStr, node.Certificate)
}

func (suite *KeeperTestSuite) TestRemoveNode() {
	suite.setNode()

	err := suite.keeper.RemoveNode(suite.ctx, nodeID)
	suite.NoError(err)

	_, found := suite.keeper.GetNode(suite.ctx, nodeID)
	suite.False(found)
}

const (
	certStr = `-----BEGIN CERTIFICATE-----
MIIBazCCAR0CFGTwvE8oG+N3uNm1gonJBh6pie5TMAUGAytlcDBYMQswCQYDVQQG
EwJDTjENMAsGA1UECAwEcm9vdDENMAsGA1UEBwwEcm9vdDENMAsGA1UECgwEcm9v
dDENMAsGA1UECwwEcm9vdDENMAsGA1UEAwwEcm9vdDAeFw0yMDA2MTkwNzAyMzla
Fw0yMDA3MTkwNzAyMzlaMFgxCzAJBgNVBAYTAkNOMQ0wCwYDVQQIDAR0ZXN0MQ0w
CwYDVQQHDAR0ZXN0MQ0wCwYDVQQKDAR0ZXN0MQ0wCwYDVQQLDAR0ZXN0MQ0wCwYD
VQQDDAR0ZXN0MCowBQYDK2VwAyEA27WvK0goa1sSjsp6eb/xCkgjBEoPC9vfL/6h
f0hqjHYwBQYDK2VwA0EA0fo8y+saUl+8UiyKpKdjv2DsqYWqmqJDz9u3NaioOvrQ
Z0mOxdgj9wfO0t3voldCRUw3hCekjC+GEOoXH5ysDQ==
-----END CERTIFICATE-----`
	certStr1 = `-----BEGIN CERTIFICATE-----
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
