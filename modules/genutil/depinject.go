package genutil

import (
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"

	"github.com/cosmos/cosmos-sdk/client"

	modulev1 "github.com/bianjieai/iritamod/api/iritamod/genutil/module/v1"
	"github.com/bianjieai/iritamod/modules/genutil/types"
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

type GenesisInputs struct {
	depinject.In

	Config           *modulev1.Module
	AccountKeeper    types.AccountKeeper
	NodeKeeper       types.NodeKeeper
	DeliverTx        deliverTxfn
	TxEncodingConfig client.TxEncodingConfig
}

func ProvideModule(in GenesisInputs) appmodule.AppModule {
	m := NewAppModule(
		in.AccountKeeper,
		in.NodeKeeper,
		in.DeliverTx,
		in.TxEncodingConfig)
	return m
}
