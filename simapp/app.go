package simapp

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cast"

	dbm "github.com/cometbft/cometbft-db"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/crypto/tmhash"
	"github.com/cometbft/cometbft/libs/log"

	simappparams "cosmossdk.io/simapp/params"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	nodeservice "github.com/cosmos/cosmos-sdk/client/grpc/node"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/store/streaming"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authsims "github.com/cosmos/cosmos-sdk/x/auth/simulation"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/capability"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	"github.com/cosmos/cosmos-sdk/x/consensus"
	consensuskeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	consensustypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	evidencekeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	feegrantmodule "github.com/cosmos/cosmos-sdk/x/feegrant/module"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	"github.com/bianjieai/iritamod/modules/identity"
	identitykeeper "github.com/bianjieai/iritamod/modules/identity/keeper"
	identitytypes "github.com/bianjieai/iritamod/modules/identity/types"
	"github.com/bianjieai/iritamod/modules/node"
	nodekeeper "github.com/bianjieai/iritamod/modules/node/keeper"
	nodetypes "github.com/bianjieai/iritamod/modules/node/types"
	cparams "github.com/bianjieai/iritamod/modules/params"
	cparamskeeper "github.com/bianjieai/iritamod/modules/params/keeper"
	cparamstypes "github.com/bianjieai/iritamod/modules/params/types"
	"github.com/bianjieai/iritamod/modules/perm"
	permkeeper "github.com/bianjieai/iritamod/modules/perm/keeper"
	permtypes "github.com/bianjieai/iritamod/modules/perm/types"
	cslashing "github.com/bianjieai/iritamod/modules/slashing"
)

const appName = "SimApp"

var (
	// DefaultNodeHome default home directories for the application daemon
	DefaultNodeHome string

	// ModuleBasics defines the module BasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	ModuleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
		genutil.AppModuleBasic{},
		bank.AppModuleBasic{},
		capability.AppModuleBasic{},
		params.AppModuleBasic{},
		cparams.AppModuleBasic{},
		crisis.AppModuleBasic{},
		cslashing.AppModuleBasic{},
		feegrantmodule.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		evidence.AppModuleBasic{},
		perm.AppModuleBasic{},
		identity.AppModuleBasic{},
		node.AppModuleBasic{},
		consensus.AppModuleBasic{},
	)

	// module account permissions
	maccPerms = map[string][]string{
		authtypes.FeeCollectorName: nil,
		cparams.ModuleName:         nil,
	}

	// module accounts that are allowed to receive tokens
	allowedReceivingModAcc = map[string]bool{}

	// root admin address in consortium chain
	rootAdmin string

	msgTypeURLs = []string{
		"/cosmos.auth.v1beta1.MsgUpdateParams",
		"/cosmos.bank.v1beta1.MsgUpdateParams",
		"/cosmos.consensus.v1.MsgUpdateParams",
		"/cosmos.crisis.v1beta1.MsgUpdateParams",
		"/slashing.MsgUpdateParams",
		"/node.MsgUpdateParams",
	}
)

var (
	_ runtime.AppI            = (*SimApp)(nil)
	_ servertypes.Application = (*SimApp)(nil)
)

// SimApp extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type SimApp struct {
	*baseapp.BaseApp
	legacyAmino       *codec.LegacyAmino
	appCodec          codec.Codec
	txConfig          client.TxConfig
	interfaceRegistry types.InterfaceRegistry

	// keys to access the substores
	keys    map[string]*storetypes.KVStoreKey
	tkeys   map[string]*storetypes.TransientStoreKey
	memKeys map[string]*storetypes.MemoryStoreKey

	// keepers
	AccountKeeper   authkeeper.AccountKeeper
	BankKeeper      bankkeeper.Keeper
	SlashingKeeper  slashingkeeper.Keeper
	CrisisKeeper    *crisiskeeper.Keeper
	UpgradeKeeper   *upgradekeeper.Keeper
	ParamsKeeper    paramskeeper.Keeper
	EvidenceKeeper  evidencekeeper.Keeper
	PermKeeper      permkeeper.Keeper
	IdentityKeeper  identitykeeper.Keeper
	NodeKeeper      nodekeeper.Keeper
	FeeGrantKeeper  feegrantkeeper.Keeper
	ConsensusKeeper consensuskeeper.Keeper
	CParamsKeeper   cparamskeeper.Keeper

	// the module manager
	ModuleManager *module.Manager

	// simulation manager
	sm *module.SimulationManager

	// module configurator
	configurator module.Configurator
}

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	rootAdmin = sdk.AccAddress(tmhash.SumTruncated([]byte("rootAdmin"))).String()

	DefaultNodeHome = filepath.Join(userHomeDir, ".simapp")
}

// NewSimApp returns a reference to an initialized NewSimApp.
func NewSimApp(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	skipUpgradeHeights map[int64]bool,
	appOpts servertypes.AppOptions,
	baseAppOptions ...func(*baseapp.BaseApp),
) *SimApp {

	encodingConfig := makeEncodingConfig()
	appCodec := encodingConfig.Codec
	legacyAmino := encodingConfig.Amino
	interfaceRegistry := encodingConfig.InterfaceRegistry
	txConfig := encodingConfig.TxConfig

	bApp := baseapp.NewBaseApp(appName, logger, db, txConfig.TxDecoder(), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetVersion(version.Version)
	bApp.SetInterfaceRegistry(interfaceRegistry)
	bApp.SetTxEncoder(txConfig.TxEncoder())

	keys := sdk.NewKVStoreKeys(
		authtypes.StoreKey,
		banktypes.StoreKey,
		crisistypes.StoreKey,
		slashingtypes.StoreKey,
		paramstypes.StoreKey,
		upgradetypes.StoreKey,
		feegrant.StoreKey,
		evidencetypes.StoreKey,
		permtypes.StoreKey,
		identitytypes.StoreKey,
		nodetypes.StoreKey,
		consensustypes.StoreKey,
	)

	tkeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	// load state streaming if enabled
	if _, _, err := streaming.LoadStreamingServices(bApp, appOpts, appCodec, logger, keys); err != nil {
		logger.Error("failed to load state streaming", "err", err)
		os.Exit(1)
	}

	app := &SimApp{
		BaseApp:           bApp,
		legacyAmino:       legacyAmino,
		appCodec:          appCodec,
		txConfig:          txConfig,
		interfaceRegistry: interfaceRegistry,
		keys:              keys,
		tkeys:             tkeys,
		memKeys:           memKeys,
	}

	app.ParamsKeeper = initParamsKeeper(appCodec, legacyAmino, keys[paramstypes.StoreKey], tkeys[paramstypes.TStoreKey])

	// set the BaseApp's consensus parameter store
	app.ConsensusKeeper = consensuskeeper.NewKeeper(
		appCodec,
		keys[consensustypes.StoreKey],
		authtypes.NewModuleAddress(cparamstypes.ModuleName).String())
	bApp.SetParamStore(&app.ConsensusKeeper)

	app.AccountKeeper = authkeeper.NewAccountKeeper(
		appCodec,
		keys[authtypes.StoreKey],
		authtypes.ProtoBaseAccount,
		maccPerms,
		sdk.Bech32MainPrefix,
		authtypes.NewModuleAddress(cparamstypes.ModuleName).String())

	app.BankKeeper = bankkeeper.NewBaseKeeper(
		appCodec,
		keys[banktypes.StoreKey],
		app.AccountKeeper,
		BlockedAddresses(),
		authtypes.NewModuleAddress(cparamstypes.ModuleName).String())

	app.NodeKeeper = nodekeeper.NewKeeper(
		appCodec,
		keys[nodetypes.StoreKey])

	app.SlashingKeeper = slashingkeeper.NewKeeper(
		appCodec,
		legacyAmino,
		keys[slashingtypes.StoreKey],
		&app.NodeKeeper,
		authtypes.NewModuleAddress(cparamstypes.ModuleName).String())

	invCheckPeriod := cast.ToUint(appOpts.Get(server.FlagInvCheckPeriod))
	app.CrisisKeeper = crisiskeeper.NewKeeper(
		appCodec,
		keys[crisistypes.StoreKey],
		invCheckPeriod,
		app.BankKeeper,
		authtypes.FeeCollectorName,
		authtypes.NewModuleAddress(cparamstypes.ModuleName).String())

	app.FeeGrantKeeper = feegrantkeeper.NewKeeper(
		appCodec,
		keys[feegrant.StoreKey],
		app.AccountKeeper)

	homePath := cast.ToString(appOpts.Get(flags.FlagHome))
	app.UpgradeKeeper = upgradekeeper.NewKeeper(
		skipUpgradeHeights,
		keys[upgradetypes.StoreKey],
		appCodec,
		homePath,
		app.BaseApp,
		authtypes.NewModuleAddress(cparamstypes.ModuleName).String())

	evidenceKeeper := evidencekeeper.NewKeeper(
		appCodec,
		keys[evidencetypes.StoreKey],
		&app.NodeKeeper,
		app.SlashingKeeper)
	app.EvidenceKeeper = *evidenceKeeper

	app.NodeKeeper = *app.NodeKeeper.SetHooks(
		stakingtypes.NewMultiStakingHooks(app.SlashingKeeper.Hooks()),
	)

	app.PermKeeper = permkeeper.NewKeeper(
		appCodec,
		keys[permtypes.StoreKey])

	app.IdentityKeeper = identitykeeper.NewKeeper(
		appCodec,
		keys[identitytypes.StoreKey])

	app.CParamsKeeper = cparamskeeper.NewKeeper(
		app.AccountKeeper, app.MsgServiceRouter(), msgTypeURLs)

	/****  Module Options ****/

	// NOTE: we may consider parsing `appOpts` inside module constructors. For the moment
	// we prefer to be more strict in what arguments the modules expect.
	var skipGenesisInvariants = cast.ToBool(appOpts.Get(crisis.FlagSkipGenesisInvariants))

	// NOTE: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.
	app.ModuleManager = module.NewManager(
		genutil.NewAppModule(app.AccountKeeper, app.NodeKeeper, app.BaseApp.DeliverTx, encodingConfig.TxConfig),
		auth.NewAppModule(appCodec, app.AccountKeeper, authsims.RandomGenesisAccounts, app.GetSubspace(authtypes.ModuleName)),
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper, app.GetSubspace(banktypes.ModuleName)),
		crisis.NewAppModule(app.CrisisKeeper, skipGenesisInvariants, app.GetSubspace(crisistypes.ModuleName)),
		feegrantmodule.NewAppModule(appCodec, app.AccountKeeper, app.BankKeeper, app.FeeGrantKeeper, app.interfaceRegistry),
		cslashing.NewAppModule(appCodec, cslashing.NewKeeper(app.SlashingKeeper, app.NodeKeeper), app.AccountKeeper, app.BankKeeper, &app.NodeKeeper, app.GetSubspace(slashingtypes.ModuleName)),
		upgrade.NewAppModule(app.UpgradeKeeper),
		evidence.NewAppModule(app.EvidenceKeeper),
		params.NewAppModule(app.ParamsKeeper),
		perm.NewAppModule(appCodec, app.PermKeeper),
		identity.NewAppModule(app.IdentityKeeper),
		node.NewAppModule(appCodec, app.NodeKeeper, app.GetSubspace(nodetypes.ModuleName)),
		consensus.NewAppModule(appCodec, app.ConsensusKeeper),
		cparams.NewAppModule(appCodec, app.CParamsKeeper),
	)

	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	// NOTE: staking module is required if HistoricalEntries param > 0
	app.ModuleManager.SetOrderBeginBlockers(
		permtypes.ModuleName,
		authtypes.ModuleName,
		nodetypes.ModuleName,
		banktypes.ModuleName,
		slashingtypes.ModuleName,
		crisistypes.ModuleName,
		evidencetypes.ModuleName,
		identitytypes.ModuleName,
		feegrant.ModuleName,
		upgradetypes.ModuleName,
		paramstypes.ModuleName,
		genutiltypes.ModuleName,
		consensustypes.ModuleName,
		cparamstypes.ModuleName,
	)

	app.ModuleManager.SetOrderEndBlockers(
		permtypes.ModuleName,
		authtypes.ModuleName,
		nodetypes.ModuleName,
		banktypes.ModuleName,
		slashingtypes.ModuleName,
		crisistypes.ModuleName,
		evidencetypes.ModuleName,
		identitytypes.ModuleName,
		feegrant.ModuleName,
		upgradetypes.ModuleName,
		paramstypes.ModuleName,
		genutiltypes.ModuleName,
		consensustypes.ModuleName,
		cparamstypes.ModuleName,
	)

	// NOTE: The genutils module must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	// NOTE: Capability module must occur first so that it can initialize any capabilities
	// so that other modules that want to create or claim capabilities afterwards in InitChain
	// can do so safely.
	app.ModuleManager.SetOrderInitGenesis(
		permtypes.ModuleName,
		authtypes.ModuleName,
		nodetypes.ModuleName,
		banktypes.ModuleName,
		slashingtypes.ModuleName,
		crisistypes.ModuleName,
		evidencetypes.ModuleName,
		identitytypes.ModuleName,
		feegrant.ModuleName,
		upgradetypes.ModuleName,
		paramstypes.ModuleName,
		genutiltypes.ModuleName,
		consensustypes.ModuleName,
		cparamstypes.ModuleName,
	)

	app.ModuleManager.SetOrderMigrations(
		permtypes.ModuleName,
		authtypes.ModuleName,
		nodetypes.ModuleName,
		banktypes.ModuleName,
		slashingtypes.ModuleName,
		crisistypes.ModuleName,
		evidencetypes.ModuleName,
		identitytypes.ModuleName,
		feegrant.ModuleName,
		upgradetypes.ModuleName,
		paramstypes.ModuleName,
		genutiltypes.ModuleName,
		consensustypes.ModuleName,
		cparamstypes.ModuleName,
	)

	app.ModuleManager.RegisterInvariants(app.CrisisKeeper)
	app.configurator = module.NewConfigurator(app.appCodec, app.MsgServiceRouter(), app.GRPCQueryRouter())
	app.ModuleManager.RegisterServices(app.configurator)

	// add test gRPC service for testing gRPC queries in isolation
	testdata.RegisterQueryServer(app.GRPCQueryRouter(), testdata.QueryImpl{})

	// create the simulation manager and define the order of the modules for deterministic simulations
	//
	// NOTE: this is not required apps that don't use the simulator for fuzz testing
	// transactions
	app.sm = module.NewSimulationManager(
		auth.NewAppModule(appCodec, app.AccountKeeper, authsims.RandomGenesisAccounts, app.GetSubspace(authtypes.ModuleName)),
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper, app.GetSubspace(banktypes.ModuleName)),
		feegrantmodule.NewAppModule(appCodec, app.AccountKeeper, app.BankKeeper, app.FeeGrantKeeper, app.interfaceRegistry),
		//gov.NewAppModule(appCodec, app.govKeeper, app.AccountKeeper, app.BankKeeper),
		cslashing.NewAppModule(appCodec, cslashing.NewKeeper(app.SlashingKeeper, app.NodeKeeper), app.AccountKeeper, app.BankKeeper, &app.NodeKeeper, app.GetSubspace(slashingtypes.ModuleName)),
		params.NewAppModule(app.ParamsKeeper),
		cparams.NewAppModule(appCodec, app.CParamsKeeper),
		perm.NewAppModule(appCodec, app.PermKeeper),
		identity.NewAppModule(app.IdentityKeeper),
		node.NewAppModule(appCodec, app.NodeKeeper, app.GetSubspace(nodetypes.ModuleName)),
	)

	app.sm.RegisterStoreDecoders()

	// initialize stores
	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)
	app.MountMemoryStores(memKeys)

	// initialize BaseApp
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	anteHandler, err := ante.NewAnteHandler(
		ante.HandlerOptions{
			AccountKeeper:   app.AccountKeeper,
			BankKeeper:      app.BankKeeper,
			SignModeHandler: encodingConfig.TxConfig.SignModeHandler(),
			FeegrantKeeper:  app.FeeGrantKeeper,
			SigGasConsumer:  ante.DefaultSigVerificationGasConsumer,
		},
	)
	if err != nil {
		panic(err)
	}
	app.SetAnteHandler(anteHandler)
	app.SetEndBlocker(app.EndBlocker)

	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			logger.Error("error on loading last version", "err", err)
			os.Exit(1)
		}
	}

	return app
}

// Name returns the name of the App
func (app *SimApp) Name() string { return app.BaseApp.Name() }

// BeginBlocker application updates every begin block
func (app *SimApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.ModuleManager.BeginBlock(ctx, req)
}

// EndBlocker application updates every end block
func (app *SimApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.ModuleManager.EndBlock(ctx, req)
}

// InitChainer application update at chain initialization
func (app *SimApp) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState GenesisState
	app.legacyAmino.MustUnmarshalJSON(req.AppStateBytes, &genesisState)
	return app.ModuleManager.InitGenesis(ctx, app.appCodec, genesisState)
}

// LoadHeight loads a particular height
func (app *SimApp) LoadHeight(height int64) error {
	return app.LoadVersion(height)
}

// ModuleAccountAddrs returns all the app's module account addresses.
func (app *SimApp) ModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

// LegacyAmino returns SimApp's amino codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *SimApp) LegacyAmino() *codec.LegacyAmino {
	return app.legacyAmino
}

// AppCodec returns SimApp's app codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *SimApp) AppCodec() codec.Codec {
	return app.appCodec
}

// InterfaceRegistry returns SimApp's InterfaceRegistry
func (app *SimApp) InterfaceRegistry() types.InterfaceRegistry {
	return app.interfaceRegistry
}

// TxConfig returns SimApp's TxConfig
func (app *SimApp) TxConfig() client.TxConfig {
	return app.txConfig
}

// DefaultGenesis returns a default genesis from the registered AppModuleBasic's.
func (a *SimApp) DefaultGenesis() map[string]json.RawMessage {
	return ModuleBasics.DefaultGenesis(a.appCodec)
}

// GetKey returns the KVStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *SimApp) GetKey(storeKey string) *storetypes.KVStoreKey {
	return app.keys[storeKey]
}

// GetTKey returns the TransientStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *SimApp) GetTKey(storeKey string) *storetypes.TransientStoreKey {
	return app.tkeys[storeKey]
}

// GetMemKey returns the MemStoreKey for the provided mem key.
//
// NOTE: This is solely used for testing purposes.
func (app *SimApp) GetMemKey(storeKey string) *storetypes.MemoryStoreKey {
	return app.memKeys[storeKey]
}

// GetSubspace returns a param subspace for a given module name.
//
// NOTE: This is solely to be used for testing purposes.
func (app *SimApp) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := app.ParamsKeeper.GetSubspace(moduleName)
	return subspace
}

// SimulationManager implements the SimulationApp interface
func (app *SimApp) SimulationManager() *module.SimulationManager {
	return app.sm
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *SimApp) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	clientCtx := apiSvr.ClientCtx
	// Register new tx routes from grpc-gateway.
	authtx.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Register new tendermint queries routes from grpc-gateway.
	tmservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Register node gRPC service for grpc-gateway.
	nodeservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Register grpc-gateway routes for all modules.
	ModuleBasics.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// register swagger API from root so that other applications can override easily
	if err := server.RegisterSwaggerAPI(apiSvr.ClientCtx, apiSvr.Router, apiConfig.Swagger); err != nil {
		panic(err)
	}
}

// RegisterTxService implements the Application.RegisterTxService method.
func (app *SimApp) RegisterTxService(clientCtx client.Context) {
	authtx.RegisterTxService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.BaseApp.Simulate, app.interfaceRegistry)
}

// RegisterTendermintService implements the Application.RegisterTendermintService method.
func (app *SimApp) RegisterTendermintService(clientCtx client.Context) {
	tmservice.RegisterTendermintService(
		clientCtx,
		app.BaseApp.GRPCQueryRouter(),
		app.interfaceRegistry,
		app.Query,
	)
}

func (app *SimApp) RegisterNodeService(clientCtx client.Context) {
	nodeservice.RegisterNodeService(clientCtx, app.GRPCQueryRouter())
}

// GetMaccPerms returns a copy of the module account permissions
//
// NOTE: This is solely to be used for testing purposes.
func GetMaccPerms() map[string][]string {
	dupMaccPerms := make(map[string][]string)
	for k, v := range maccPerms {
		dupMaccPerms[k] = v
	}
	return dupMaccPerms
}

// BlockedAddresses returns all the app's blocked account addresses.
func BlockedAddresses() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range GetMaccPerms() {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	// allow the following addresses to receive funds
	// NOTE: there's no gov module in consortium chain
	// delete(modAccAddrs, authtypes.NewModuleAddress(govtypes.ModuleName).String())

	return modAccAddrs
}

// initParamsKeeper init params keeper and its subspaces
func initParamsKeeper(appCodec codec.BinaryCodec, legacyAmino *codec.LegacyAmino, key, tkey storetypes.StoreKey) paramskeeper.Keeper {
	ParamsKeeper := paramskeeper.NewKeeper(appCodec, legacyAmino, key, tkey)

	ParamsKeeper.Subspace(authtypes.ModuleName)
	ParamsKeeper.Subspace(banktypes.ModuleName)
	ParamsKeeper.Subspace(nodetypes.ModuleName).WithKeyTable(nodetypes.ParamKeyTable())
	ParamsKeeper.Subspace(slashingtypes.ModuleName).WithKeyTable(slashingtypes.ParamKeyTable())
	ParamsKeeper.Subspace(crisistypes.ModuleName)

	return ParamsKeeper
}

func makeEncodingConfig() simappparams.EncodingConfig {
	encodingConfig := simappparams.MakeTestEncodingConfig()
	std.RegisterLegacyAminoCodec(encodingConfig.Amino)
	std.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	ModuleBasics.RegisterLegacyAminoCodec(encodingConfig.Amino)
	ModuleBasics.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	return encodingConfig
}
