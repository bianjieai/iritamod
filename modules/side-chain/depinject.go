package sidechain

import (
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"

	"github.com/cosmos/cosmos-sdk/codec"
	store "github.com/cosmos/cosmos-sdk/store/types"

	modulev1 "github.com/bianjieai/iritamod/api/iritamod/side-chain/module/v1"
	"github.com/bianjieai/iritamod/modules/side-chain/keeper"
	"github.com/bianjieai/iritamod/modules/side-chain/types"
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

type SideChainInputs struct {
	depinject.In

	Config *modulev1.Module
	Key    *store.KVStoreKey
	Cdc    codec.Codec

	AccountKeeper types.AccountKeeper
}

type SideChainOutputs struct {
	depinject.Out

	Keeper keeper.Keeper
	Module appmodule.AppModule
}

func ProvideModule(in SideChainInputs) SideChainOutputs {
	k := keeper.NewKeeper(
		in.Cdc,
		in.Key,
		in.AccountKeeper,
	)

	m := NewAppModule(
		in.Cdc,
		k)

	return SideChainOutputs{
		Keeper: k,
		Module: m,
	}
}
