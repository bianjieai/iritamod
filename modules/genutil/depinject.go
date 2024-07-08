package genutil

//import (
//	"cosmossdk.io/core/appmodule"
//	"cosmossdk.io/depinject"
//	modulev1 "github.com/bianjieai/iritamod/api/iritamod/genutil/module/v1"
//	"github.com/bianjieai/iritamod/modules/genutil/types"
//	"github.com/cosmos/cosmos-sdk/client"
//	"github.com/cosmos/cosmos-sdk/x/genutil"
//)
//
//func init() {
//	appmodule.Register(&modulev1.Module{},
//		appmodule.Provide(ProvideModule),
//	)
//}
//
//var _ appmodule.AppModule = AppModule{}
//
//func (am AppModule) IsOnePerModuleType() {}
//
//func (am AppModule) IsAppModule() {}
//
//type GenutilInputs struct {
//	depinject.In
//	accountKeeper types.AccountKeeper
//	nodeKeeper    types.NodeKeeper
//	deliverTx     deliverTxfn
//	txConfig      client.TxConfig
//}
//
//type GenutiOutputs struct {
//	depinject.Out
//	Module appmodule.AppModule
//}
//
//func ProvideModule(in GenutilInputs) GenutiOutputs {
//	m := genutil.NewAppModule(in.accountKeeper, in.nodeKeeper, in.deliverTx, in.txConfig)
//	return GenutiOutputs{Module: m}
//}
