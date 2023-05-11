package layer2

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bianjieai/iritamod/modules/layer2/keeper"
	"github.com/bianjieai/iritamod/modules/layer2/types"
	"github.com/bianjieai/iritamod/modules/opb/client/cli"
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

// AppModuleBasic defines the basic application module used by layer2 module.
type AppModuleBasic struct {
	cdc codec.Codec
}

// Name returns the layer2 module's name.
func (AppModuleBasic) Name() string { return types.ModuleName }

// RegisterLegacyAminoCodec registers the layer2 module's types on the LegacyAmino codec.
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	types.RegisterLegacyAminoCodec(cdc)
}

// RegisterInterfaces registers the layer2 module's interface types
func (AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	types.RegisterInterfaces(registry)
}

// DefaultGenesis returns default genesis state as raw bytes for the layer2 module.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesisState())
}

// ValidateGenesis performs genesis state validation for the layer2 module.
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, _ client.TxEncodingConfig, bz json.RawMessage) error {
	var data types.GenesisState
	if err := cdc.UnmarshalJSON(bz, &data); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
	}
	return types.ValidateGenesis(data)
}

// RegisterRESTRoutes registers the REST routes for the layer2 module.
func (AppModuleBasic) RegisterRESTRoutes(clientCtx client.Context, rtr *mux.Router) {}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the layer2 module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	_ = types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx))
}

// GetTxCmd returns the root tx command for the layer2 module.
func (AppModuleBasic) GetTxCmd() *cobra.Command {
    return cli.NewTxCmd()
}

// GetQueryCmd returns no root query command for the layer2 module.
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd()
}

// AppModule Implementation

type AppModule struct {
	AppModuleBasic

	keeper keeper.Keeper
}

// NewAppModule creates a new AppModule object
func NewAppModule(cdc codec.Codec, keeper keeper.Keeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{cdc: cdc},
		keeper:         keeper,
	}
}

// RegisterInvariants registers the layer2 module invariants.
func (AppModule) RegisterInvariants(sdk.InvariantRegistry) {
	//TODO: add what invariants for layer2 module?
}

// Route returns the message routing key for the layer2 module.
func (am AppModule) Route() sdk.Route {
	return sdk.Route{}
}

// NewHandler returns the layer2 module's querier route name.
func (AppModule) QuerierRoute() string {
	return ""
}

// LegacyQuerierHandler returns the layer2 module's querier r.
func (AppModule) LegacyQuerierHandler(*codec.LegacyAmino) sdk.Querier {
    return nil
}

func (am AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterMsgServer(cfg.MsgServer(), am.keeper)
	types.RegisterQueryServer(cfg.QueryServer(), am.keeper)
}

// ConsensusVersion implements AppModule/ConsensusVersion.
func (AppModule) ConsensusVersion() uint64 {
	return 1
}

// BeginBlock returns the begin blocker for the layer2 module.
func (AppModule) BeginBlock(sdk.Context, abci.RequestBeginBlock) {}

// EndBlock returns the end blocker for the layer2 module.
func (AppModule) EndBlock(sdk.Context, abci.RequestEndBlock) []abci.ValidatorUpdate {
	return nil
}

// InitGenesis performs genesis initialization for the layer2 module. It returns no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) []abci.ValidatorUpdate {
	var genesisState types.GenesisState
	cdc.MustUnmarshalJSON(data, &genesisState)
	am.keeper.InitGenesis(ctx, &genesisState)

	return []abci.ValidatorUpdate{}
}

func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	gs := am.keeper.ExportGenesis(ctx)
	return cdc.MustMarshalJSON(gs)
}

// AppModuleSimulation Implementation
// NOTE: we haven't made a concrete implementation for simulation of the layer2 module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {}

func (AppModule) ProposalContents(simState module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

func (AppModule) RandomizedParams(r *rand.Rand) []simtypes.ParamChange {
	return nil
}

func (AppModule) RegisterStoreDecoder(sdr sdk.StoreDecoderRegistry) {}

func (AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	return nil
}
