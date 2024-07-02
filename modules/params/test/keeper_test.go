package test

//
//import (
//	"fmt"
//	"github.com/cosmos/cosmos-sdk/client"
//	"github.com/cosmos/cosmos-sdk/codec"
//	"github.com/cosmos/cosmos-sdk/codec/types"
//	"github.com/cosmos/cosmos-sdk/std"
//	"github.com/cosmos/cosmos-sdk/store"
//	storetypes "github.com/cosmos/cosmos-sdk/store/types"
//	sdk "github.com/cosmos/cosmos-sdk/types"
//	"github.com/cosmos/cosmos-sdk/types/module"
//	"github.com/cosmos/cosmos-sdk/x/auth/tx"
//	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
//	paramstype "github.com/cosmos/cosmos-sdk/x/params/types"
//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/suite"
//	"github.com/tendermint/tendermint/libs/log"
//	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
//	dbm "github.com/tendermint/tm-db"
//	"irita.bianjie.ai/modules/params/client/utils"
//	"irita.bianjie.ai/modules/params/keeper"
//	iritamodparamstype "irita.bianjie.ai/modules/params/types"
//	"testing"
//	"time"
//)
//
//type KeeperTestSuite struct {
//	suite.Suite
//
//	encCfg TestEncodingConfig
//	ctx    sdk.Context
//	keeper keeper.Keeper
//	ss     paramstype.Subspace
//}
//
//var (
//	keyUnbondingTime = []byte("UnbondingTime")
//	keyMaxValidators = []byte("MaxValidators")
//	keyBondDenom     = []byte("BondDenom")
//)
//
//func TestKeeperTestSuite(t *testing.T) {
//	suite.Run(t, new(KeeperTestSuite))
//}
//
//func (suite *KeeperTestSuite) SetupTest() {
//	key := sdk.NewKVStoreKey(iritamodparamstype.StoreKey)
//	tkey := sdk.NewTransientStoreKey("transient_test")
//	testCtx := DefaultContextWithDB(suite.T(), key, tkey)
//	suite.ctx = testCtx.Ctx.WithBlockHeader(tmproto.Header{Time: time.Now()})
//	encCfg := MakeTestEncodingConfig()
//	suite.encCfg = encCfg
//	paramskeeper := paramskeeper.NewKeeper(encCfg.Codec, encCfg.Amino, key, tkey)
//	suite.keeper = keeper.NewKeeper(paramskeeper)
//	ss := suite.keeper.Subspace("testsubspace")
//	suite.ss = ss.WithKeyTable(paramKeyTable())
//}
//
//func (suite *KeeperTestSuite) TestUpdateParams() {
//	t := time.Hour * 48
//	suite.ss.Set(suite.ctx, keyUnbondingTime, t)
//	paramChanges := utils.ParamChangeJSON{}
//	if err := suite.encCfg.Amino.UnmarshalJSON([]byte("{\"subspace\":\"testsubspace\",\"key\":\"UnbondingTime\",\"value\":\"182800000000000\"}"), &paramChanges); err != nil {
//		panic(err)
//	}
//	paramChange := paramChanges.ToParamChange()
//	updateparams := []iritamodparamstype.ParamChange{paramChange}
//	changeEvents, err := suite.keeper.UpdateParams(suite.ctx, updateparams)
//	suite.NoError(err)
//	suite.NotNil(changeEvents)
//
//}
//
//type TestContext struct {
//	Ctx sdk.Context
//	DB  *dbm.MemDB
//	CMS store.CommitMultiStore
//}
//
//func DefaultContextWithDB(t *testing.T, key storetypes.StoreKey, tkey storetypes.StoreKey) TestContext {
//	db := dbm.NewMemDB()
//	cms := store.NewCommitMultiStore(db)
//	cms.MountStoreWithDB(key, storetypes.StoreTypeIAVL, db)
//	cms.MountStoreWithDB(tkey, storetypes.StoreTypeTransient, db)
//	err := cms.LoadLatestVersion()
//	assert.NoError(t, err)
//
//	ctx := sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger())
//
//	return TestContext{ctx, db, cms}
//}
//
//type TestEncodingConfig struct {
//	InterfaceRegistry types.InterfaceRegistry
//	Codec             codec.Codec
//	TxConfig          client.TxConfig
//	Amino             *codec.LegacyAmino
//}
//
//func MakeTestEncodingConfig(modules ...module.AppModuleBasic) TestEncodingConfig {
//	cdc := codec.NewLegacyAmino()
//	interfaceRegistry := types.NewInterfaceRegistry()
//	codec := codec.NewProtoCodec(interfaceRegistry)
//
//	encCfg := TestEncodingConfig{
//		InterfaceRegistry: interfaceRegistry,
//		Codec:             codec,
//		TxConfig:          tx.NewTxConfig(codec, tx.DefaultSignModes),
//		Amino:             cdc,
//	}
//
//	mb := module.NewBasicManager(modules...)
//
//	std.RegisterLegacyAminoCodec(encCfg.Amino)
//	std.RegisterInterfaces(encCfg.InterfaceRegistry)
//	mb.RegisterLegacyAminoCodec(encCfg.Amino)
//	mb.RegisterInterfaces(encCfg.InterfaceRegistry)
//
//	return encCfg
//}
//
//func paramKeyTable() paramstype.KeyTable {
//	return paramstype.NewKeyTable().RegisterParamSet(&params{})
//}
//
//type params struct {
//	UnbondingTime time.Duration `json:"unbonding_time" yaml:"unbonding_time"`
//	MaxValidators uint16        `json:"max_validators" yaml:"max_validators"`
//	BondDenom     string        `json:"bond_denom" yaml:"bond_denom"`
//}
//
//func (p *params) ParamSetPairs() paramstype.ParamSetPairs {
//	return paramstype.ParamSetPairs{
//		{keyUnbondingTime, &p.UnbondingTime, validateUnbondingTime},
//		{keyMaxValidators, &p.MaxValidators, validateMaxValidators},
//		{keyBondDenom, &p.BondDenom, validateBondDenom},
//	}
//}
//
//func validateUnbondingTime(i interface{}) error {
//	v, ok := i.(time.Duration)
//	if !ok {
//		return fmt.Errorf("invalid parameter type: %T", i)
//	}
//
//	if v < (24 * time.Hour) {
//		return fmt.Errorf("unbonding time must be at least one day")
//	}
//
//	return nil
//}
//
//func validateMaxValidators(i interface{}) error {
//	_, ok := i.(uint16)
//	if !ok {
//		return fmt.Errorf("invalid parameter type: %T", i)
//	}
//
//	return nil
//}
//
//func validateBondDenom(i interface{}) error {
//	v, ok := i.(string)
//	if !ok {
//		return fmt.Errorf("invalid parameter type: %T", i)
//	}
//
//	if len(v) == 0 {
//		return fmt.Errorf("denom cannot be empty")
//	}
//
//	return nil
//}
