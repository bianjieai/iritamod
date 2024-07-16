package upgrade

import (
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"
	"github.com/cosmos/cosmos-sdk/codec"
	store "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
	Key       *store.KVStoreKey
	Cdc       codec.Codec
	Authority sdk.AccAddress
}

type UpgradeOutputs struct {
	depinject.Out
	UpgradeKeeper keeper.Keeper
	Module        appmodule.AppModule
}

func ProvideModule(in UpgradeInputs) UpgradeOutputs {
	skipUpgradeHeights := make(map[int64]bool)
	keeper := keeper.NewKeeper(skipUpgradeHeights, in.Key, in.Cdc, "/", nil, in.Authority.String())
	m := NewAppModule(keeper)
	return UpgradeOutputs{UpgradeKeeper: keeper, Module: m}
}
