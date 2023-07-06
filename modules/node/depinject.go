package node

import (
	"fmt"

	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/cosmos/cosmos-sdk/codec"
	store "github.com/cosmos/cosmos-sdk/store/types"

	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	modulev1 "github.com/bianjieai/iritamod/api/iritamod/node/module/v1"
	"github.com/bianjieai/iritamod/modules/node/exported"
	"github.com/bianjieai/iritamod/modules/node/keeper"
	"github.com/bianjieai/iritamod/modules/node/types"
)

//
// App Wiring Setup
//

func init() {
	appmodule.Register(
		&modulev1.Module{},
		appmodule.Provide(ProvideModule, ProvideKeyTable),
		appmodule.Invoke(InvokeSetStakingHooks),
	)
}

func ProvideKeyTable() paramtypes.KeyTable {
	return types.ParamKeyTable() //nolint:staticcheck
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

	// LegacySubspace is used solely for migration of x/params managed parameters
	LegacySubspace exported.Subspace
}

type NodeOutputs struct {
	depinject.Out

	NodeKeeper *keeper.Keeper
	Module     appmodule.AppModule
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
		NodeKeeper: &k,
		Module:     m,
	}
}

// FIXME： where do staking hooks come from?
func InvokeSetStakingHooks(
	keeper *keeper.Keeper,
	stakingHooks map[string]stakingtypes.StakingHooksWrapper) error {

	if keeper == nil {
		return nil
	}

	var multiHooks stakingtypes.MultiStakingHooks
	for modName := range stakingHooks {
		hook, ok := stakingHooks[modName]
		if !ok {
			return fmt.Errorf("can't find staking hooks for module %s", modName)
		}

		multiHooks = append(multiHooks, hook)
	}

	keeper.SetHooks(multiHooks)

	return nil
}
