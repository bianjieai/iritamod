package opb

import (
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"

	"github.com/cosmos/cosmos-sdk/codec"
	store "github.com/cosmos/cosmos-sdk/store/types"

	modulev1 "github.com/bianjieai/iritamod/api/iritamod/opb/module/v1"
	"github.com/bianjieai/iritamod/modules/opb/exported"
	"github.com/bianjieai/iritamod/modules/opb/keeper"
	"github.com/bianjieai/iritamod/modules/opb/types"
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

type OpbInputs struct {
	depinject.In

	Config *modulev1.Module
	Key    *store.KVStoreKey
	Cdc    codec.Codec

	AccountKeeper types.AccountKeeper
	BankKeeper    types.BankKeeper
	TokenKeeper   types.TokenKeeper
	PermKeeper    types.PermKeeper

	// LegacySubspace is used solely for migration of x/params managed parameters
	LegacySubspace exported.Subspace
}

type OpbOutputs struct {
	depinject.Out

	Keeper keeper.Keeper
	Module appmodule.AppModule
}

func ProvideModule(in OpbInputs) OpbOutputs {
	k := keeper.NewKeeper(
		in.Cdc,
		in.Key,
		in.AccountKeeper,
		in.BankKeeper,
		in.TokenKeeper,
		in.PermKeeper)

	m := NewAppModule(
		in.Cdc,
		k,
		in.LegacySubspace)

	return OpbOutputs{
		Keeper: k,
		Module: m,
	}
}
