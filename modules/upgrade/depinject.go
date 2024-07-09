package upgrade

import (
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"
	modulev1 "github.com/bianjieai/iritamod/api/iritamod/upgrade/module/v1"
	"github.com/bianjieai/iritamod/modules/upgrade/keeper"
	"github.com/cosmos/cosmos-sdk/codec"
	store "github.com/cosmos/cosmos-sdk/store/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
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
	//upgrade.UpgradeInputs
	Key *store.KVStoreKey
	Cdc codec.Codec
}

type UpgradeOutputs struct {
	depinject.Out
	UpgradeKeeper keeper.Keeper
	Module        appmodule.AppModule
}

func ProvideModule(in UpgradeInputs) UpgradeOutputs {
	skipUpgradeHeights := make(map[int64]bool)
	authority := authtypes.NewModuleAddress(govtypes.ModuleName)
	cosmosupgradekeeper := upgradekeeper.NewKeeper(skipUpgradeHeights, in.Key, in.Cdc, "/", nil, authority.String())
	keeper := keeper.NewKeeper(
		cosmosupgradekeeper,
	)
	m := NewAppModule(keeper)
	return UpgradeOutputs{UpgradeKeeper: keeper, Module: m}
}
