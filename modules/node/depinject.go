package node

import (
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"
	"github.com/cosmos/cosmos-sdk/codec"
	store "github.com/cosmos/cosmos-sdk/store/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	modulev1 "iritamod.bianjie.ai/api/iritamod/node/module/v1"
	"iritamod.bianjie.ai/modules/node/keeper"
	"iritamod.bianjie.ai/modules/node/types"
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
	Cdc            codec.Codec
	Key            *store.KVStoreKey
	Paramskeeper   paramskeeper.Keeper
	Slashingkeeper slashingkeeper.Keeper
}

type NodeOutputs struct {
	depinject.Out

	NodeKeeper keeper.Keeper
	Module     appmodule.AppModule
}

func ProvideModule(in NodeInputs) NodeOutputs {
	var subspace paramstypes.Subspace
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
	nodekeeper = *nodekeeper.SetHooks(
		stakingtypes.NewMultiStakingHooks(in.Slashingkeeper.Hooks()),
	)
	m := NewAppModule(in.Cdc, nodekeeper)

	return NodeOutputs{NodeKeeper: nodekeeper, Module: m}
}
