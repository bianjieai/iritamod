package keeper_test

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/suite"

	"github.com/tendermint/tendermint/crypto/sm2"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/iritamod/modules/identity/keeper"
	"github.com/bianjieai/iritamod/modules/identity/types"
	"github.com/bianjieai/iritamod/simapp"
)

var (
	testID = tmbytes.HexBytes(uuid.NewV4().Bytes())

	testPubKeySM2     = sm2.GenPrivKey().PubKey().Bytes()
	testPubKeySM2Info = types.PubKeyInfo{PubKey: tmbytes.HexBytes(testPubKeySM2).String(), Algorithm: types.SM2}

	testPrivKeyECDSA, _ = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	testPubKeyECDSA     = elliptic.Marshal(testPrivKeyECDSA.PublicKey.Curve, testPrivKeyECDSA.X, testPrivKeyECDSA.Y)
	testPubKeyECDSAInfo = types.PubKeyInfo{PubKey: tmbytes.HexBytes(testPubKeyECDSA).String(), Algorithm: types.ECDSA}

	testCredentials = "https://kyc.com/user/10001"
	testOwner       = sdk.AccAddress([]byte("test-ownertest-owner"))
	testData        = "test_data"
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
	suite.keeper = &app.IdentityKeeper
}

func (suite *KeeperTestSuite) setIdentity() {
	identity := types.NewIdentity(testID, []types.PubKeyInfo{testPubKeySM2Info}, []string{testCertificate}, testCredentials, testOwner, testData)
	err := suite.keeper.SetIdentity(suite.ctx, identity)
	suite.NoError(err)
}

func (suite *KeeperTestSuite) TestCreateIdentity() {
	err := suite.keeper.CreateIdentity(suite.ctx, testID, &testPubKeySM2Info, testCertificate, testCredentials, testData, testOwner)
	suite.NoError(err)

	identity, found := suite.keeper.GetIdentity(suite.ctx, testID)
	suite.True(found)

	suite.Equal(testID.String(), identity.Id)
	suite.Len(identity.PubKeys, 2)
	suite.Equal(testPubKeySM2Info, identity.PubKeys[1])
	suite.Len(identity.Certificates, 1)
	suite.Equal(testCertificate, identity.Certificates[0])
	suite.Equal(testCredentials, identity.Credentials)
	suite.Equal(testOwner.String(), identity.Owner)
	suite.Equal(testData, identity.Data)
}

func (suite *KeeperTestSuite) TestUpdateIdentity() {
	suite.setIdentity()

	newPubKey := testPubKeyECDSAInfo
	newCredentials := "https://kyc.com/v2/user/10001"

	err := suite.keeper.UpdateIdentity(suite.ctx, testID, &newPubKey, "", newCredentials, testData, testOwner)
	suite.NoError(err)

	identity, found := suite.keeper.GetIdentity(suite.ctx, testID)
	suite.True(found)

	suite.Equal(testID.String(), identity.Id)
	suite.Len(identity.PubKeys, 3)
	suite.Equal(testPubKeyECDSAInfo, identity.PubKeys[1])
	suite.Equal(testPubKeySM2Info, identity.PubKeys[2])
	suite.Len(identity.Certificates, 1)
	suite.Equal(testCertificate, identity.Certificates[0])
	suite.Equal(newCredentials, identity.Credentials)
	suite.Equal(testOwner.String(), identity.Owner)
	suite.Equal(testData, identity.Data)
}

const testCertificate = `-----BEGIN CERTIFICATE-----
MIIDTDCCAjQCCQDvRoz+e/HRpDANBgkqhkiG9w0BAQsFADBoMQswCQYDVQQGEwJj
bjELMAkGA1UECAwCc2gxCzAJBgNVBAcMAnBkMQswCQYDVQQKDAJiajELMAkGA1UE
CwwCYmoxCzAJBgNVBAMMAmJqMRgwFgYJKoZIhvcNAQkBFgliakBiai5jb20wHhcN
MjAwNjEwMDkzNjMxWhcNMjAwNzEwMDkzNjMxWjBoMQswCQYDVQQGEwJjbjELMAkG
A1UECAwCc2gxCzAJBgNVBAcMAnBkMQswCQYDVQQKDAJiajELMAkGA1UECwwCYmox
CzAJBgNVBAMMAmJqMRgwFgYJKoZIhvcNAQkBFgliakBiai5jb20wggEiMA0GCSqG
SIb3DQEBAQUAA4IBDwAwggEKAoIBAQDUEifXes1/CXEjdH8SeSS+1x+ZlhktI8i8
9ncMeOr5oI1Mc7Kd7v85i0hrmjjZzUrHQy0Sdt2ltQjo6dtkq3wDsL4OgIqGO75z
OwG4EB0A1sJ/YTSX+fmWwy5ys19A2O5sTZOJEw3VFgiZHv1TZEiY+GVtpZ5Dti/1
t5ZzNTF+M0rpbICTxLh1GSpdhJs95yci1A8zqmPzPETVkxJwVCOg54WfpRQAiBqM
DKLjVXALuvlDDxVhB0u7kuvKAydZdV/pDs73HuY2srCOiDij3iVS01Ln02JNeMK8
IG9xRSw2eaSDp+fa1jtUXMDMmVNHCJqpQaFv0/1oN/ehUXb/DTMHAgMBAAEwDQYJ
KoZIhvcNAQELBQADggEBAKij8eUTcs+AJFPnzc3aolVZEApwvLum58WRjmoev44A
1528F4dXF7vJhIbqdOvEBy0YNQhNuNUs+JiHIFwuVvhNuAXDgXJNsvymx8fn0E5U
C90iTCiV9WhlL93S6fSelDj65sgD4Gw8Q4bBbNa/SRCu4+oBNS9BPjpcbrGllph9
7AkCGBiaabVLqGNyZJEKZpRQ3kOqdQzHYT/eHRC3hcO/KGf0vCOUTgEhHuYavMy/
JZOeFg1owNP2nZ8cD2TwDKS+T+T1rAG1ovnVp/PV7lbH1o8Kn2rwtj1S42O824Gr
2NyVhhdZkLI/uEX9mdmcFPB+oV6iiPnqEh/r2wswFgw=
-----END CERTIFICATE-----`
