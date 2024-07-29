package test

import (
	"cosmossdk.io/math"
	"fmt"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/suite"
	//nodekeeper "iritamod.bianjie.ai/modules/node/keeper"
	//nodeTypes "iritamod.bianjie.ai/modules/node/types"
	"iritamod.bianjie.ai/modules/slashing/keeper"
	slashingtypes "iritamod.bianjie.ai/modules/slashing/types"
	"iritamod.bianjie.ai/simapp"
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

var (
	bondedAcc    = authtypes.NewEmptyModuleAccount(stakingtypes.BondedPoolName)
	notBondedAcc = authtypes.NewEmptyModuleAccount(stakingtypes.NotBondedPoolName)
	PKs          = simtestutil.CreateTestPubKeys(500)
)

type KeeperTestSuite struct {
	suite.Suite

	cdc    codec.Codec
	ctx    sdk.Context
	keeper keeper.Keeper
	//nodeKeeper    nodekeeper.Keeper
	StakingKeeper  stakingkeeper.Keeper
	SlashingKeeper slashingkeeper.Keeper
	//nodekeeper nodekeeper.Keeper
	app *simapp.SimApp
}

//var (
//	isCheckTx   = false
//	name        = "test_name"
//	details     = "test_details"
//	power       = int64(1)
//	cert, _     = ca.ReadCertificateFromMem([]byte(CertStr))
//	pk, _       = ca.GetPubkeyFromCert(cert)
//	cospk, _    = cryptocodec.FromTmPubKeyInterface(pk)
//	nodeID      = pk.Address()
//	operator, _ = sdk.AccAddressFromHexUnsafe(pk.Address().String())
//	nodeName    = "test_node"
//	consAddr    = sdk.ConsAddress(operator)
//)

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) SetupTest() {
	depInjectOptions := simapp.DepinjectOptions{
		Config:    AppConfig,
		Providers: []interface{}{},
		Consumers: []interface{}{&suite.keeper, &suite.StakingKeeper},
	}
	app := simapp.Setup(suite.T(), false, depInjectOptions)
	suite.ctx = app.BaseApp.NewContext(false, tmproto.Header{})
	//suite.nodeKeeper.SetHooks(suite.keeper.Hooks())
}

func (suite *KeeperTestSuite) TestSlashing() {
	//suite.keeper.Jail()
	//suite.StakingKeeper.SetValidator()
	valAddr := sdk.ValAddress(PKs[0].Address().Bytes())
	consAddr := sdk.ConsAddress(PKs[0].Address())
	validator, err := NewValidator(valAddr, PKs[0], stakingtypes.Description{})
	suite.NoError(err)
	//msg := nodeTypes.NewMsgCreateValidator(name, details, CertStr, power, operator)
	//validator := tmbytes.HexBytes(tmhash.Sum(msg.GetSignBytes()))
	//err := suite.nodeKeeper.CreateValidator(suite.ctx,
	//	validator,
	//	msg.Name,
	//	msg.Certificate,
	//	nil,
	//	msg.Power,
	//	msg.Description,
	//	msg.Operator,
	//)
	//suite.NoError(err)
	//validator1, found := suite.nodeKeeper.GetValidator(suite.ctx, validator)
	//suite.True(found)
	//suite.Equal(validator.String(), validator1.Id)
	suite.StakingKeeper.SetValidator(suite.ctx, validator)
	vali, has := suite.StakingKeeper.GetValidator(suite.ctx, valAddr)
	suite.True(has)
	suite.NotNil(vali)
	suite.Equal(validator.OperatorAddress, vali.OperatorAddress)
	conAddr, err := vali.GetConsAddr()
	fmt.Println(conAddr)
	suite.NoError(err)

	height := int64(0)
	target := int64(100)
	for ; height < target; height++ {
		suite.ctx = suite.ctx.WithBlockHeight(height)
		//suite.keeper.HandleValidatorSignature(suite.ctx, consAddr, int64(1), true)
	}

	for ; height < target+500; height++ {
		suite.ctx = suite.ctx.WithBlockHeight(height)
		//suite.keeper.HandleValidatorSignature(suite.ctx, valAddr, int64(1), false)
	}
	vali1, found := suite.StakingKeeper.GetValidator(suite.ctx, valAddr)
	suite.True(found)
	suite.True(vali1.Jailed)
	unjailmsg := slashingtypes.MsgUnjailValidator{
		Id:       "1",
		Operator: valAddr.String(),
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
	vali1, found = suite.StakingKeeper.GetValidator(suite.ctx, valAddr)
	suite.True(found)
	suite.False(vali1.Jailed)

}

func NewValidator(operator sdk.ValAddress, pubKey cryptotypes.PubKey, description stakingtypes.Description) (stakingtypes.Validator, error) {
	pkAny, err := codectypes.NewAnyWithValue(pubKey)
	if err != nil {
		return stakingtypes.Validator{}, err
	}

	return stakingtypes.Validator{
		OperatorAddress:         operator.String(),
		ConsensusPubkey:         pkAny,
		Jailed:                  false,
		Status:                  stakingtypes.Unbonded,
		Tokens:                  math.ZeroInt(),
		DelegatorShares:         math.LegacyZeroDec(),
		Description:             description,
		UnbondingHeight:         int64(0),
		UnbondingTime:           time.Unix(0, 0).UTC(),
		Commission:              stakingtypes.NewCommission(math.LegacyZeroDec(), math.LegacyZeroDec(), math.LegacyZeroDec()),
		MinSelfDelegation:       math.OneInt(),
		UnbondingOnHoldRefCount: 0,
	}, nil
}

//var consAddr = sdk.ConsAddress(sdk.AccAddress([]byte("addr1_______________")))
//
//type KeeperTestSuite struct {
//	suite.Suite
//
//	ctx            sdk.Context
//	stakingKeeper  stakingkeeper.Keeper
//	slashingKeeper slashingkeeper.Keeper
//	//queryClient    slashingtypes.QueryClient
//	msgServer      slashingtypes.MsgServer
//}
//func (s *KeeperTestSuite) SetupTest() {
//	depInjectOptions := simapp.DepinjectOptions{
//		Config:    AppConfig,
//		Providers: []interface{}{},
//		Consumers: []interface{}{&suite.keeper, &suite.nodeKeeper, &suite.StakingKeeper},
//	}
//	app := simapp.Setup(suite.T(), isCheckTx, depInjectOptions)
//	suite.ctx = app.BaseApp.NewContext(false, tmproto.Header{})
//	suite.nodeKeeper.SetHooks(suite.keeper.Hooks())
//}
//	//key := sdk.NewKVStoreKey(slashingtypes.StoreKey)
//	//testCtx := sdktestutil.DefaultContextWithDB(s.T(), key, sdk.NewTransientStoreKey("transient_test"))
//	//ctx := testCtx.Ctx.WithBlockHeader(tmproto.Header{Time: tmtime.Now()})
//	//encCfg := moduletestutil.MakeTestEncodingConfig()
//
//	// gomock initializations
//	//ctrl := gomock.NewController(s.T())
//	//s.stakingKeeper = slashingtestutil.NewMockStakingKeeper(ctrl)
//
//	s.ctx = ctx
//	//s.slashingKeeper = slashingkeeper.NewKeeper(
//	//	encCfg.Codec,
//	//	encCfg.Amino,
//	//	key,
//	//	s.stakingKeeper,
//	//	authtypes.NewModuleAddress(govtypes.ModuleName).String(),
//	//)
//	// set test params
//	s.slashingKeeper.SetParams(ctx, slashingtestutil.TestParams())
//
//	slashingtypes.RegisterInterfaces(encCfg.InterfaceRegistry)
//	queryHelper := baseapp.NewQueryServerTestHelper(ctx, encCfg.InterfaceRegistry)
//	slashingtypes.RegisterQueryServer(queryHelper, s.slashingKeeper)
//
//	s.queryClient = slashingtypes.NewQueryClient(queryHelper)
//	s.msgServer = slashingkeeper.NewMsgServerImpl(s.slashingKeeper)
//}
//
//func (s *KeeperTestSuite) TestPubkey() {
//	ctx, keeper := s.ctx, s.slashingKeeper
//	require := s.Require()
//
//	_, pubKey, addr := testdata.KeyTestPubAddr()
//	require.NoError(keeper.AddPubkey(ctx, pubKey))
//
//	expectedPubKey, err := keeper.GetPubkey(ctx, addr.Bytes())
//	require.NoError(err)
//	require.Equal(pubKey, expectedPubKey)
//}
//
//func (s *KeeperTestSuite) TestJailAndSlash() {
//	s.stakingKeeper.EXPECT().SlashWithInfractionReason(s.ctx,
//		consAddr,
//		s.ctx.BlockHeight(),
//		sdk.TokensToConsensusPower(sdk.NewInt(1), sdk.DefaultPowerReduction),
//		s.slashingKeeper.SlashFractionDoubleSign(s.ctx),
//		stakingtypes.Infraction_INFRACTION_UNSPECIFIED,
//	).Return(sdk.NewInt(0))
//
//	s.slashingKeeper.Slash(
//		s.ctx,
//		consAddr,
//		s.slashingKeeper.SlashFractionDoubleSign(s.ctx),
//		sdk.TokensToConsensusPower(sdk.NewInt(1), sdk.DefaultPowerReduction),
//		s.ctx.BlockHeight(),
//	)
//
//	s.stakingKeeper.EXPECT().Jail(s.ctx, consAddr).Return()
//	s.slashingKeeper.Jail(s.ctx, consAddr)
//}
//
//func (s *KeeperTestSuite) TestJailAndSlashWithInfractionReason() {
//	s.stakingKeeper.EXPECT().SlashWithInfractionReason(s.ctx,
//		consAddr,
//		s.ctx.BlockHeight(),
//		sdk.TokensToConsensusPower(sdk.NewInt(1), sdk.DefaultPowerReduction),
//		s.slashingKeeper.SlashFractionDoubleSign(s.ctx),
//		stakingtypes.Infraction_INFRACTION_DOUBLE_SIGN,
//	).Return(sdk.NewInt(0))
//
//	s.slashingKeeper.SlashWithInfractionReason(
//		s.ctx,
//		consAddr,
//		s.slashingKeeper.SlashFractionDoubleSign(s.ctx),
//		sdk.TokensToConsensusPower(sdk.NewInt(1), sdk.DefaultPowerReduction),
//		s.ctx.BlockHeight(),
//		stakingtypes.Infraction_INFRACTION_DOUBLE_SIGN,
//	)
//
//	s.stakingKeeper.EXPECT().Jail(s.ctx, consAddr).Return()
//	s.slashingKeeper.Jail(s.ctx, consAddr)
//}
//
//func TestKeeperTestSuite(t *testing.T) {
//	suite.Run(t, new(KeeperTestSuite))
