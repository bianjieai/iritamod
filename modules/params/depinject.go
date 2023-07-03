package params

import (
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"

	modulev1 "github.com/bianjieai/iritamod/api/iritamod/params/module/v1"
	"github.com/bianjieai/iritamod/modules/params/keeper"
	"github.com/bianjieai/iritamod/modules/params/types"
)

var _ appmodule.AppModule = AppModule{}

//
// App Wiring Setup
//

func init() {
	appmodule.Register(
		&modulev1.Module{},
		appmodule.Provide(ProvideModule),
	)
}

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (am AppModule) IsOnePerModuleType() {}

// IsAppModule implements the appmodule.AppModule interface.
func (am AppModule) IsAppModule() {}

type ParamsInputs struct {
	depinject.In

	Config *modulev1.Module
	Cdc    codec.Codec

	AuthKeeper types.AccountKeeper

	Router      *baseapp.MsgServiceRouter
	MsgTypeURLs []string
}

type ParamsOutputs struct {
	depinject.Out

	Keeper keeper.Keeper
	Module appmodule.AppModule
}

func ProvideModule(in ParamsInputs) ParamsOutputs {
	k := keeper.NewKeeper(
		in.AuthKeeper,
		in.Router,
		in.MsgTypeURLs)

	m := NewAppModule(
		in.Cdc,
		k)

	return ParamsOutputs{
		Keeper: k,
		Module: m,
	}
}
