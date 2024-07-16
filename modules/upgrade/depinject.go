package upgrade

import (
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	store "github.com/cosmos/cosmos-sdk/store/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/spf13/cast"
	modulev1 "iritamod.bianjie.ai/api/iritamod/upgrade/module/v1"
	"iritamod.bianjie.ai/modules/upgrade/keeper"
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

type UpgradeInputs struct {
	depinject.In
	Config  *modulev1.Module
	Key     *store.KVStoreKey
	Cdc     codec.Codec
	AppOpts servertypes.AppOptions
}

type UpgradeOutputs struct {
	depinject.Out
	UpgradeKeeper keeper.Keeper
	Module        appmodule.AppModule
}

func ProvideModule(in UpgradeInputs) UpgradeOutputs {
	authority := authtypes.NewModuleAddress(govtypes.ModuleName)
	if in.Config.Authority != "" {
		authority = authtypes.NewModuleAddressOrBech32Address(in.Config.Authority)
	}
	skipUpgradeHeights := make(map[int64]bool)
	keeper := keeper.NewKeeper(skipUpgradeHeights, in.Key, in.Cdc, cast.ToString(in.AppOpts.Get(flags.FlagHome)), nil, authority.String())
	m := NewAppModule(keeper)
	return UpgradeOutputs{UpgradeKeeper: keeper, Module: m}
}
