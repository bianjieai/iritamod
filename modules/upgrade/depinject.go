package upgrade

import (
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/server"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	store "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	sdk "github.com/cosmos/cosmos-sdk/x/upgrade"
	"github.com/spf13/cast"
	modulev1 "iritamod.bianjie.ai/api/iritamod/upgrade/module/v1"
	"iritamod.bianjie.ai/modules/upgrade/keeper"
	"iritamod.bianjie.ai/modules/upgrade/types"
)

// App Wiring Setup
func init() {
	appmodule.Register(&modulev1.Module{},
		appmodule.Provide(ProvideModule),
		appmodule.Invoke(PopulateVersionMap),
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
	GovHandler    govv1beta1.HandlerRoute
	BaseAppOption runtime.BaseAppOption
}

func ProvideModule(in UpgradeInputs) UpgradeOutputs {
	//authority := authtypes.NewModuleAddress(govtypes.ModuleName)
	//if in.Config.Authority != "" {
	//	authority = authtypes.NewModuleAddressOrBech32Address(in.Config.Authority)
	//}
	//skipUpgradeHeights := make(map[int64]bool)
	//keeper := keeper.NewKeeper(skipUpgradeHeights, in.Key, in.Cdc, cast.ToString(in.AppOpts.Get(flags.FlagHome)), nil, authority.String())
	//m := NewAppModule(keeper)
	//return UpgradeOutputs{UpgradeKeeper: keeper, Module: m}
	var (
		homePath           string
		skipUpgradeHeights = make(map[int64]bool)
	)

	if in.AppOpts != nil {
		for _, h := range cast.ToIntSlice(in.AppOpts.Get(server.FlagUnsafeSkipUpgrades)) {
			skipUpgradeHeights[int64(h)] = true
		}

		homePath = cast.ToString(in.AppOpts.Get(flags.FlagHome))
	}

	// default to governance authority if not provided
	authority := authtypes.NewModuleAddress(govtypes.ModuleName)
	if in.Config.Authority != "" {
		authority = authtypes.NewModuleAddressOrBech32Address(in.Config.Authority)
	}

	// set the governance module account as the authority for conducting upgrades
	k := keeper.NewKeeper(skipUpgradeHeights, in.Key, in.Cdc, homePath, nil, authority.String())
	baseappOpt := func(app *baseapp.BaseApp) {
		k.SetVersionSetter(app)
	}
	m := NewAppModule(k)
	gh := govv1beta1.HandlerRoute{RouteKey: types.RouterKey, Handler: sdk.NewSoftwareUpgradeProposalHandler(k.Keeper)}

	return UpgradeOutputs{UpgradeKeeper: k, Module: m, GovHandler: gh, BaseAppOption: baseappOpt}
}
func PopulateVersionMap(upgradeKeeper *keeper.Keeper, modules map[string]appmodule.AppModule) {
	if upgradeKeeper == nil {
		return
	}

	upgradeKeeper.SetInitVersionMap(module.NewManagerFromMap(modules).GetVersionMap())
}
