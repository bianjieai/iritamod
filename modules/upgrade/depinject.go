package upgrade

import (
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"
	"github.com/cosmos/cosmos-sdk/codec"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"

	modulev1 "github.com/bianjieai/iritamod/api/iritamod/upgrade/module/v1"
	"github.com/bianjieai/iritamod/modules/upgrade/keeper"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
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

type ParamsInputs struct {
	depinject.In
	Cdc codec.Codec
}

type UpgradeOutputs struct {
	depinject.Out
	UpgradeKeeper keeper.Keeper
}

func ProvideModule(in upgrade.UpgradeInputs) UpgradeOutputs {
	cosmosupgradekeeper := upgradekeeper.NewKeeper(
		in.Cdc,
		in.LegacyAmino,
		in.TransientStoreKey,
		in.KvStoreKey,
	)
	keeper := keeper.NewKeeper(
		cosmosupgradekeeper,
	)
	m := NewAppModule(in.Cdc, cosmosParamsKeeper)
	return UpgradeOutputs{paramsKeeper: keeper, Module: m}
}
