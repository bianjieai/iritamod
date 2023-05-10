package layer2

import (
	"encoding/json"
	"fmt"
	"context"
	"github.com/bianjieai/iritamod/modules/layer2/types"
	"github.com/bianjieai/iritamod/modules/node/keeper"
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
