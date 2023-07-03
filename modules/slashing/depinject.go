package slashing

import (
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"
	"github.com/cosmos/cosmos-sdk/codec"

	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"

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
		appmodule.Provide(ProvideModule),
	)
}

var _ appmodule.AppModule = AppModule{}

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (am AppModule) IsOnePerModuleType() {}

// IsAppModule implements the appmodule.AppModule interface.
func (am AppModule) IsAppModule() {}

type SlashingInputs struct {
	depinject.In

	Config *modulev1.Module
	Cdc    codec.Codec

	AccountKeeper AccountKeeper
	BankKeeper    BankKeeper
	SlashKeeper   slashingkeeper.Keeper
	NodeKeeper    types.NodeKeeper
	StakingKeeper StakingKeeper

	// LegacySubspace is used solely for migration of x/params managed parameters
	LegacySubspace exported.Subspace
}

type SlashingOutputs struct {
	depinject.Out

	Keeper keeper.Keeper
	Module appmodule.AppModule
}

func ProvideModule(in SlashingInputs) SlashingOutputs {
	k := keeper.NewKeeper(
		in.SlashKeeper,
		in.NodeKeeper)

	m := NewAppModule(
		in.Cdc,
		k,
		in.AccountKeeper,
		in.BankKeeper,
		in.StakingKeeper,
		in.LegacySubspace)

	return SlashingOutputs{
		Keeper: k,
		Module: m,
	}
}
