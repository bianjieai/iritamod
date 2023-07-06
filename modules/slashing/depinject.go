package slashing

import (
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"

	modulev1 "github.com/bianjieai/iritamod/api/iritamod/slashing/module/v1"
	"github.com/bianjieai/iritamod/modules/slashing/exported"
	"github.com/bianjieai/iritamod/modules/slashing/keeper"
	"github.com/bianjieai/iritamod/modules/slashing/types"
)

//
// App Wiring Setup
//

func init() {
	appmodule.Register(
		&modulev1.Module{},
		appmodule.Provide(ProvideModule, ProvideKeyTable),
	)
}

func ProvideKeyTable() paramtypes.KeyTable {
	return slashingtypes.ParamKeyTable() //nolint:staticcheck
}

var _ appmodule.AppModule = AppModule{}

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (am AppModule) IsOnePerModuleType() {}

// IsAppModule implements the appmodule.AppModule interface.
func (am AppModule) IsAppModule() {}

type SlashingInputs struct {
	depinject.In

	Config      *modulev1.Module
	Cdc         codec.Codec
	Key         *storetypes.KVStoreKey
	LegacyAmino *codec.LegacyAmino

	NodeKeeper types.NodeKeeper

	// LegacySubspace is used solely for migration of x/params managed parameters
	LegacySubspace exported.Subspace
}

type SlashingOutputs struct {
	depinject.Out

	SlashingKeeper keeper.Keeper
	Module         appmodule.AppModule
}

func ProvideModule(in SlashingInputs) SlashingOutputs {
	// default to governance authority if not provided
	authority := authtypes.NewModuleAddress(govtypes.ModuleName)

	k := keeper.NewKeeper(
		in.Cdc,
		in.LegacyAmino,
		in.Key,
		in.NodeKeeper,
		authority.String(),
	)

	m := NewAppModule(
		in.Cdc,
		k,
		in.NodeKeeper,
		in.LegacySubspace,
	)

	return SlashingOutputs{
		SlashingKeeper: k,
		Module:         m,
	}
}
