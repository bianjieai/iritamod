package node

import (
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"
	"github.com/cosmos/cosmos-sdk/codec"
	store "github.com/cosmos/cosmos-sdk/store/types"
	cosmosparamstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	paramskeeper "iritamod.bianjie.ai/modules/params/keeper"

	modulev1 "iritamod.bianjie.ai/api/iritamod/node/module/v1"
	"iritamod.bianjie.ai/modules/node/keeper"
	"iritamod.bianjie.ai/modules/node/types"
)

// App Wiring Setup
func init() {
	appmodule.Register(&modulev1.Module{},
		// appmodule.Provide(ProvideModule, ProvideKeyTable),
		appmodule.Provide(ProvideModule),
		appmodule.Invoke(InvokeHooks),
	)
}

var _ appmodule.AppModule = AppModule{}

//func ProvideKeyTable() paramstypes.KeyTable {
//	return keeper.ParamKeyTable()
//}

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (am AppModule) IsOnePerModuleType() {}

// IsAppModule implements the appmodule.AppModule interface.
func (am AppModule) IsAppModule() {}

type HookInputs struct {
	depinject.In

	Hooks  []staking.StakingHooks `optional:"true"`
	Keeper *keeper.Keeper
}

type NodeInputs struct {
	depinject.In
	Cdc          codec.Codec
	Key          *store.KVStoreKey
	Paramskeeper paramskeeper.Keeper
}

type NodeOutputs struct {
	depinject.Out

	NodeKeeper *keeper.Keeper
	Module     appmodule.AppModule
}

func ProvideModule(in NodeInputs) NodeOutputs {
	var subspace cosmosparamstypes.Subspace
	if space, ok := in.Paramskeeper.GetSubspace(types.ModuleName); ok {
		subspace = space
	} else {
		subspace = in.Paramskeeper.Subspace(types.ModuleName)
	}
	nodekeeper := keeper.NewKeeper(
		in.Cdc,
		in.Key,
		subspace,
	)
	//nodekeeper.SetHooks(in.StakingKeeper.Hooks())
	m := NewAppModule(in.Cdc, *nodekeeper)

	return NodeOutputs{NodeKeeper: nodekeeper, Module: m}
}
func InvokeHooks(in HookInputs) {
	in.Keeper.SetHooks(stakingtypes.NewMultiStakingHooks(in.Hooks...))
}
