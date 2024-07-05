package slashing

import (
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"
	"github.com/bianjieai/iritamod/modules/slashing/types"
	"github.com/cosmos/cosmos-sdk/codec"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	modulev1 "github.com/bianjieai/iritamod/api/iritamod/params/module/v1"
	"github.com/bianjieai/iritamod/modules/slashing/keeper"
	coamosslashing "github.com/cosmos/cosmos-sdk/x/slashing"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
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

type ParamsInputs struct {
	depinject.In
	coamosslashing.SlashingInputs
	NodeKeeper types.NodeKeeper
	Cdc        codec.Codec
}

type ParamsOutputs struct {
	depinject.Out
	SlashingKeeper keeper.Keeper
	Module         appmodule.AppModule
}

func ProvideModule(in ParamsInputs) ParamsOutputs {
	authority := authtypes.NewModuleAddress(govtypes.ModuleName)
	if in.Config.Authority != "" {
		authority = authtypes.NewModuleAddressOrBech32Address(in.Config.Authority)
	}
	cosmosSlashingKeeper := slashingkeeper.NewKeeper(in.Cdc, in.LegacyAmino, in.Key, in.StakingKeeper, authority.String())
	keeper := keeper.NewKeeper(
		cosmosSlashingKeeper,
		in.NodeKeeper,
	)
	m := NewAppModule(in.Cdc, keeper, in.AccountKeeper, in.BankKeeper, in.StakingKeeper)
	return ParamsOutputs{SlashingKeeper: keeper, Module: m}
}
