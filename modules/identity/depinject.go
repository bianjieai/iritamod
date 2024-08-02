package identity

import (
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"

	"github.com/cosmos/cosmos-sdk/codec"
	store "github.com/cosmos/cosmos-sdk/store/types"

	modulev1 "iritamod.bianjie.ai/api/iritamod/identity/module/v1"
	"iritamod.bianjie.ai/modules/identity/keeper"
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

type IdentityInputs struct {
	depinject.In

	Cdc codec.Codec
	Key *store.KVStoreKey
}

type IdentityOutputs struct {
	depinject.Out

	IdentityKeeper keeper.Keeper
	Module         appmodule.AppModule
}

func ProvideModule(in IdentityInputs) IdentityOutputs {
	keeper := keeper.NewKeeper(
		in.Cdc,
		in.Key,
	)
	m := NewAppModule(keeper)

	return IdentityOutputs{IdentityKeeper: keeper, Module: m}
}
