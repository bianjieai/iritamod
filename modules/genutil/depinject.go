package genutil

import (
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/client"
	modulev1 "iritamod.bianjie.ai/api/iritamod/genutil/module/v1"
	"iritamod.bianjie.ai/modules/genutil/types"
)

func init() {
	appmodule.Register(&modulev1.Module{},
		appmodule.Provide(ProvideModule),
	)
}

var _ appmodule.AppModule = AppModule{}

func (am AppModule) IsOnePerModuleType() {}

func (am AppModule) IsAppModule() {}

type GenutilInputs struct {
	depinject.In
	AccountKeeper types.AccountKeeper
	NodeKeeper    types.NodeKeeper
	DeliverTx     func(abci.RequestDeliverTx) abci.ResponseDeliverTx
	TxConfig      client.TxConfig
}

type GenutiOutputs struct {
	depinject.Out
	Module appmodule.AppModule
}

func ProvideModule(in GenutilInputs) appmodule.AppModule {
	m := NewAppModule(in.AccountKeeper, in.NodeKeeper, in.DeliverTx, in.TxConfig)
	return m
}
