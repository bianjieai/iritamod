package test

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"github.com/bianjieai/iritamod/identity/keeper"
	identityType "github.com/bianjieai/iritamod/identity/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/crypto/sm2"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"
	"testing"
	"time"
)

var (
	testID = tmbytes.HexBytes(uuid.NewV4().Bytes())

	testPubKeySM2     = sm2.GenPrivKey().PubKey().Bytes()
	testPubKeySM2Info = identityType.PubKeyInfo{PubKey: tmbytes.HexBytes(testPubKeySM2).String(), Algorithm: identityType.SM2}

	testPrivKeyECDSA, _ = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	testPubKeyECDSA     = elliptic.Marshal(testPrivKeyECDSA.PublicKey.Curve, testPrivKeyECDSA.X, testPrivKeyECDSA.Y)
	testPubKeyECDSAInfo = identityType.PubKeyInfo{PubKey: tmbytes.HexBytes(testPubKeyECDSA).String(), Algorithm: identityType.ECDSA}

	testCredentials = "https://kyc.com/user/10001"
	testOwner       = sdk.AccAddress([]byte("test-ownertest-owner"))
	testData        = "test_data"
	rootCert        = `-----BEGIN CERTIFICATE-----

MIIBxTCCAXegAwIBAgIUHMPutrm+7FT7fIFf2fEgyQnIg8kwBQYDK2VwMFgxCzAJ
BgNVBAYTAkNOMQ0wCwYDVQQIDARyb290MQ0wCwYDVQQHDARyb290MQ0wCwYDVQQK
DARyb290MQ0wCwYDVQQLDARyb290MQ0wCwYDVQQDDARyb290MB4XDTIwMDYxOTA3
MDExMVoXDTIxMDYxOTA3MDExMVowWDELMAkGA1UEBhMCQ04xDTALBgNVBAgMBHJv
b3QxDTALBgNVBAcMBHJvb3QxDTALBgNVBAoMBHJvb3QxDTALBgNVBAsMBHJvb3Qx
DTALBgNVBAMMBHJvb3QwKjAFBgMrZXADIQDdzGFcck4I7Wa1vRj4JsdQ9RjVgH92
7iOhXJ8mFLwQKaNTMFEwHQYDVR0OBBYEFPrjTGR+/g4RUduZ9E8JSXNyI4mzMB8G
A1UdIwQYMBaAFPrjTGR+/g4RUduZ9E8JSXNyI4mzMA8GA1UdEwEB/wQFMAMBAf8w
BQYDK2VwA0EAT8EG5nGxwCAP4ZlfQvAhrnJI+SojlsOoE3rZ8W6/knZsrnVb6RI8
QAVleeE0pMY+MtENXcQ2wH0QRXs+wO0XCw==
-----END CERTIFICATE-----`
	testCertificate = `-----BEGIN CERTIFICATE-----
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
)

type KeeperTestSuite struct {
	suite.Suite
	ctx    sdk.Context
	keeper keeper.Keeper
	encCfg TestEncodingConfig
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) SetupTest() {
	key := sdk.NewKVStoreKey(identityType.StoreKey)
	testCtx := DefaultContextWithDB(suite.T(), key, sdk.NewTransientStoreKey("transient_test"))
	suite.ctx = testCtx.Ctx.WithBlockHeader(tmproto.Header{Time: time.Now()})
	encCfg := MakeTestEncodingConfig()
	suite.keeper = keeper.NewKeeper(encCfg.Codec, key)

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

func (suite *KeeperTestSuite) setIdentity() {
	identity := identityType.NewIdentity(testID, []identityType.PubKeyInfo{testPubKeySM2Info}, []string{testCertificate}, testCredentials, testOwner, testData)
	err := suite.keeper.SetIdentity(suite.ctx, identity)
	suite.NoError(err)
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

type TestContext struct {
	Ctx sdk.Context
	DB  *dbm.MemDB
	CMS store.CommitMultiStore
}

func DefaultContextWithDB(t *testing.T, key storetypes.StoreKey, tkey storetypes.StoreKey) TestContext {
	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)
	cms.MountStoreWithDB(key, storetypes.StoreTypeIAVL, db)
	cms.MountStoreWithDB(tkey, storetypes.StoreTypeTransient, db)
	err := cms.LoadLatestVersion()
	assert.NoError(t, err)

	ctx := sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger())

	return TestContext{ctx, db, cms}
}

type TestEncodingConfig struct {
	InterfaceRegistry types.InterfaceRegistry
	Codec             codec.Codec
	TxConfig          client.TxConfig
	Amino             *codec.LegacyAmino
}

func MakeTestEncodingConfig(modules ...module.AppModuleBasic) TestEncodingConfig {
	cdc := codec.NewLegacyAmino()
	interfaceRegistry := types.NewInterfaceRegistry()
	codec := codec.NewProtoCodec(interfaceRegistry)

	encCfg := TestEncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Codec:             codec,
		TxConfig:          tx.NewTxConfig(codec, tx.DefaultSignModes),
		Amino:             cdc,
	}

	mb := module.NewBasicManager(modules...)

	std.RegisterLegacyAminoCodec(encCfg.Amino)
	std.RegisterInterfaces(encCfg.InterfaceRegistry)
	mb.RegisterLegacyAminoCodec(encCfg.Amino)
	mb.RegisterInterfaces(encCfg.InterfaceRegistry)

	return encCfg
}
