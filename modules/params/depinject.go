package params

import (
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"
	coamosparams "github.com/cosmos/cosmos-sdk/x/params"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	modulev1 "iritamod.bianjie.ai/api/iritamod/params/module/v1"
	"iritamod.bianjie.ai/modules/params/keeper"

	coamosparamstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// App Wiring Setup
func init() {
	appmodule.Register(&modulev1.Module{},
		appmodule.Provide(ProvideModule, ProvideSubspace),
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

type SubspaceInputs struct {
	depinject.In

	Key       depinject.ModuleKey
	Keeper    keeper.Keeper
	KeyTables map[string]coamosparamstypes.KeyTable
}

func ProvideSubspace(in SubspaceInputs) coamosparamstypes.Subspace {
	moduleName := in.Key.Name()
	kt, exists := in.KeyTables[moduleName]
	if !exists {
		return in.Keeper.Subspace(moduleName)
	} else {
		return in.Keeper.Subspace(moduleName).WithKeyTable(kt)
	}
}
