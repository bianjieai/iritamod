package upgrade

import (
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"

	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"

	modulev1 "github.com/bianjieai/iritamod/api/iritamod/upgrade/module/v1"
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

	UpgradeKeeper upgradekeeper.Keeper
}

type UpgradeOutputs struct {
	depinject.Out

	Keeper keeper.Keeper
	Module appmodule.AppModule
}

func ProvideModule(in UpgradeInputs) UpgradeOutputs {
	k := keeper.NewKeeper(in.UpgradeKeeper)

	m := NewAppModule(k)

	return UpgradeOutputs{
		Keeper: k,
		Module: m,
	}
}
