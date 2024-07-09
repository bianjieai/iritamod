package params

import (
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"
	coamosparams "github.com/cosmos/cosmos-sdk/x/params"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	modulev1 "iritamod.bianjie.ai/api/iritamod/params/module/v1"
	"iritamod.bianjie.ai/modules/params/keeper"
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

//type ParamsInputs struct {
//	depinject.In
//	Cosparams paramskeeper.Keeper
//	Cdc       codec.Codec
//}

type ParamsOutputs struct {
	depinject.Out
	ParamsKeeper keeper.Keeper
	Module       appmodule.AppModule
}

func ProvideModule(in coamosparams.ParamsInputs) ParamsOutputs {
	cosmosParamsKeeper := paramskeeper.NewKeeper(
		in.Cdc,
		in.LegacyAmino,
		in.TransientStoreKey,
		in.KvStoreKey,
	)
	keeper := keeper.NewKeeper(cosmosParamsKeeper)
	m := NewAppModule(in.Cdc, cosmosParamsKeeper)
	return ParamsOutputs{ParamsKeeper: keeper, Module: m}
}
