package layer2

import (
	"encoding/json"
	"github.com/bianjieai/iritamod/modules/node/keeper"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
	abci "github.com/tendermint/tendermint/abci/types"
	"math/rand"
)

var (
	_ module.AppModule           = AppModule{}
	_ module.AppModuleBasic      = AppModuleBasic{}
	_ module.AppModuleSimulation = AppModule{}
)

// AppModuleBasic defines the basic application module used by this module.
type AppModuleBasic struct {
	cdc codec.Codec
}

func (AppModuleBasic) Name() string {
	panic("implement me")
}

func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	panic("implement me")
}

func (AppModuleBasic) RegisterInterfaces(codectypes.InterfaceRegistry) {
	panic("implement me")
}

func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	panic("implement me")
}

func (AppModuleBasic) ValidateGenesis(codec.JSONCodec, client.TxEncodingConfig, json.RawMessage) error {
	panic("implement me")
}

func (AppModuleBasic) RegisterRESTRoutes(clientCtx client.Context, rtr *mux.Router) {
	panic("implement me")
}

func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	panic("implement me")
}

func (AppModuleBasic) GetTxCmd() *cobra.Command {
	panic("implement me")
}

func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	panic("implement me")
}

// AppModule Implementation

type AppModule struct {
	AppModuleBasic

	keeper keeper.Keeper
}

func NewAppModule() AppModule {
	panic("implement me")
}

func (AppModule) RegisterInvariants(sdk.InvariantRegistry) {
	panic("implement me")
}

func (AppModule) Route() sdk.Route {
	panic("implement me")
}

func (AppModule) QuerierRoute() string {
	panic("implement me")
}

func (AppModule) LegacyQuerierHandler(*codec.LegacyAmino) sdk.Querier {
	panic("implement me")
}

func (AppModule) RegisterServices(module.Configurator) {
	panic("implement me")
}

func (AppModule) ConsensusVersion() uint64 {
	panic("implement me")
}

func (AppModule) BeginBlock(sdk.Context, abci.RequestBeginBlock) {
	panic("implement me")
}

func (AppModule) EndBlock(sdk.Context, abci.RequestEndBlock) []abci.ValidatorUpdate {
	panic("implement me")
}

func (AppModule) InitGenesis(sdk.Context, codec.JSONCodec, json.RawMessage) []abci.ValidatorUpdate {
	panic("implement me")
}

func (AppModule) ExportGenesis(ctx sdk.Context, jsonCodec codec.JSONCodec) json.RawMessage {
	panic("implement me")
}

// AppModuleSimulation Implementation

func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	panic("implement me")
}

func (AppModule) ProposalContents(simState module.SimulationState) []simtypes.WeightedProposalContent {
	panic("implement me")
}

func (AppModule) RandomizedParams(r *rand.Rand) []simtypes.ParamChange {
	panic("implement me")
}

func (AppModule) RegisterStoreDecoder(sdr sdk.StoreDecoderRegistry) {
	panic("implement me")
}

func (AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	panic("implement me")
}
