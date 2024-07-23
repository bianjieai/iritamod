package slashing

import (
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"
	"github.com/cosmos/cosmos-sdk/codec"
	store "github.com/cosmos/cosmos-sdk/store/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	"iritamod.bianjie.ai/modules/slashing/types"

	modulev1 "iritamod.bianjie.ai/api/iritamod/slashing/module/v1"
	"iritamod.bianjie.ai/modules/slashing/keeper"
)

// App Wiring Setup
func init() {
	appmodule.Register(&modulev1.Module{},
		appmodule.Provide(ProvideModule),
	)
}

var _ appmodule.AppModule = AppModule{}

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (am AppModule) IsOnePerModuleType() {}

// IsAppModule implements the appmodule.AppModule interface.
func (am AppModule) IsAppModule() {}

type SlashingInputs struct {
	depinject.In
	NodeKeeper    types.NodeKeeper
	Config        *modulev1.Module
	Key           *store.KVStoreKey
	Cdc           codec.Codec
	LegacyAmino   *codec.LegacyAmino
	AccountKeeper AccountKeeper
	BankKeeper    BankKeeper
	StakingKeeper StakingKeeper
	//Slashingkeeper slashingkeeper.Keeper
}

type SlashingOutputs struct {
	depinject.Out
	SlashingKeeper keeper.Keeper
	Module         appmodule.AppModule
}

func ProvideModule(in SlashingInputs) SlashingOutputs {
	authority := authtypes.NewModuleAddress(govtypes.ModuleName)
	if in.Config.Authority != "" {
		authority = authtypes.NewModuleAddressOrBech32Address(in.Config.Authority)
	}
	cosmosSlashingKeeper := slashingkeeper.NewKeeper(in.Cdc, in.LegacyAmino, in.Key, in.StakingKeeper, authority.String())
	keeper := keeper.NewKeeper(
		cosmosSlashingKeeper,
		in.NodeKeeper,
	)
	m := NewAppModule(in.Cdc, keeper, in.AccountKeeper, in.BankKeeper, in.StakingKeeper)
	return SlashingOutputs{SlashingKeeper: keeper, Module: m}
}
