package node

import (
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/cosmos/cosmos-sdk/codec"
	store "github.com/cosmos/cosmos-sdk/store/types"

	modulev1 "github.com/bianjieai/iritamod/api/iritamod/node/module/v1"
	"github.com/bianjieai/iritamod/modules/node/keeper"
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

type NodeInputs struct {
	depinject.In

	Cdc        codec.Codec
	Key        *store.KVStoreKey
	Paramstore paramtypes.Subspace
}

type NodeOutputs struct {
	depinject.Out

	NodeKeeper keeper.Keeper
	Module     appmodule.AppModule
}

func ProvideModule(in NodeInputs) NodeOutputs {
	keeper := keeper.NewKeeper(
		in.Cdc,
		in.Key,
		in.Paramstore,
	)
	m := NewAppModule(in.Cdc, keeper)

	return NodeOutputs{NodeKeeper: keeper, Module: m}
}
