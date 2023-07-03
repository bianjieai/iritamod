package node

import (
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"

	"github.com/cosmos/cosmos-sdk/codec"
	store "github.com/cosmos/cosmos-sdk/store/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"

	modulev1 "github.com/bianjieai/iritamod/api/iritamod/node/module/v1"
	"github.com/bianjieai/iritamod/modules/node/exported"
	"github.com/bianjieai/iritamod/modules/node/keeper"
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

type NodeInputs struct {
	depinject.In

	Config *modulev1.Module
	Key    *store.KVStoreKey
	Cdc    codec.Codec

	Hooks staking.StakingHooksWrapper

	// LegacySubspace is used solely for migration of x/params managed parameters
	LegacySubspace exported.Subspace
}

type NodeOutputs struct {
	depinject.Out

	Keeper keeper.Keeper
	Module appmodule.AppModule
}

func ProvideModule(in NodeInputs) NodeOutputs {
	k := keeper.NewKeeper(
		in.Cdc,
		in.Key,
	)

	m := NewAppModule(
		in.Cdc,
		k,
		in.LegacySubspace,
	)

	return NodeOutputs{
		Keeper: k,
		Module: m,
	}
}
