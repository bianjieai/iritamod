package upgrade

import (
	"github.com/spf13/cast"

	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/server"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	modulev1 "github.com/bianjieai/iritamod/api/iritamod/upgrade/module/v1"
	cparamtypes "github.com/bianjieai/iritamod/modules/params/types"
	"github.com/bianjieai/iritamod/modules/upgrade/keeper"
)

//
// App Wiring Setup
//

func init() {
	appmodule.Register(
		&modulev1.Module{},
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

	Key *storetypes.KVStoreKey
	Cdc codec.Codec

	AppOpts servertypes.AppOptions `optional:"true"`
}

type UpgradeOutputs struct {
	depinject.Out

	UpgradeKeeper *keeper.Keeper
	Module        appmodule.AppModule
	BaseAppOption runtime.BaseAppOption
}

func ProvideModule(in UpgradeInputs) UpgradeOutputs {
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
	authority := authtypes.NewModuleAddress(cparamtypes.ModuleName)
	// set the governance module account as the authority for conducting upgrades
	k := keeper.NewKeeper(skipUpgradeHeights, in.Key, in.Cdc, homePath, nil, authority.String())
	baseappOpt := func(app *baseapp.BaseApp) {
		k.SetVersionSetter(app)
	}
	m := NewAppModule(k)

	return UpgradeOutputs{UpgradeKeeper: k, Module: m, BaseAppOption: baseappOpt}

}
