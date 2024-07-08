package simapp

import (
	"cosmossdk.io/depinject"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/store/streaming"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/cosmos/cosmos-sdk/x/consensus"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/cosmos/cosmos-sdk/x/staking"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	"io"

	"os"
	"path/filepath"

	dbm "github.com/cometbft/cometbft-db"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	testdata_pulsar "github.com/cosmos/cosmos-sdk/testutil/testdata/testpb"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/cosmos/cosmos-sdk/x/capability"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	evidencekeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	feegrantmodule "github.com/cosmos/cosmos-sdk/x/feegrant/module"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	groupmodule "github.com/cosmos/cosmos-sdk/x/group/module"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
)

const appName = "SimApp"

// cslashing
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
		staking.AppModuleBasic{},
		mint.AppModuleBasic{},
		distr.AppModuleBasic{},
		//gov.NewAppModuleBasic(
		//	upgradeclient.ProposalHandler,
		//),
		gov.NewAppModuleBasic(
			[]govclient.ProposalHandler{
				paramsclient.ProposalHandler,
				upgradeclient.LegacyProposalHandler,
				upgradeclient.LegacyCancelProposalHandler,
			},
		),
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		//cslashing.AppModuleBasic{},
		feegrantmodule.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		evidence.AppModuleBasic{},
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		evidence.AppModuleBasic{},
		authzmodule.AppModuleBasic{},
		groupmodule.AppModuleBasic{},
		vesting.AppModuleBasic{},
		consensus.AppModuleBasic{},
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		//node.AppModuleBasic{},
	)

	// module account permissions
	maccPerms = map[string][]string{
		authtypes.FeeCollectorName: nil,
		//gov.ModuleName:                  {authtypes.Burner},
	}

	// module accounts that are allowed to receive tokens
	//allowedReceivingModAcc = map[string]bool{}
)

var (
	_ runtime.AppI            = (*SimApp)(nil)
	_ servertypes.Application = (*SimApp)(nil)
)

// SimApp extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type SimApp struct {
	*runtime.App
	//*baseapp.BaseApp
	legacyAmino       *codec.LegacyAmino
	appCodec          codec.Codec
	txConfig          client.TxConfig
	interfaceRegistry types.InterfaceRegistry

	invCheckPeriod uint

	// keys to access the substores
	keys    map[string]*storetypes.KVStoreKey
	tkeys   map[string]*storetypes.TransientStoreKey
	memKeys map[string]*storetypes.MemoryStoreKey

	// keepers
	AccountKeeper authkeeper.AccountKeeper
	BankKeeper    bankkeeper.Keeper
	//SlashingKeeper slashingkeeper.Keeper
	StakingKeeper *stakingkeeper.Keeper
	//govKeeper        gov.Keeper
	CrisisKeeper  *crisiskeeper.Keeper
	UpgradeKeeper *upgradekeeper.Keeper
	//ParamsKeeper   paramskeeper.Keeper
	EvidenceKeeper evidencekeeper.Keeper
	//NodeKeeper     nodekeeper.Keeper
	FeeGrantKeeper feegrantkeeper.Keeper

	// the module manager
	//mm *module.Manager

	// simulation manager
	sm *module.SimulationManager

	// module configurator
	//configurator module.Configurator
}

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, ".simapp")
}

// NewSimApp returns a reference to an initialized NewSimApp.
func NewSimApp(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	depInjectOptions DepinjectOptions,
	appOpts servertypes.AppOptions,
	baseAppOptions ...func(*baseapp.BaseApp),
) *SimApp {

	var (
		app        = &SimApp{}
		appBuilder *runtime.AppBuilder

		providers = append(depInjectOptions.Providers[:], appOpts)
		// merge the AppConfig and other configuration in one config
		appConfig = depinject.Configs(
			depInjectOptions.Config,
			depinject.Supply(
				providers...,

			// ADVANCED CONFIGURATION

			//
			// AUTH
			//
			// For providing a custom function required in auth to generate custom account types
			// add it below. By default the auth module uses simulation.RandomGenesisAccounts.
			//
			// authtypes.RandomGenesisAccountsFn(simulation.RandomGenesisAccounts),

			// For providing a custom a base account type add it below.
			// By default the auth module uses authtypes.ProtoBaseAccount().
			//
			// func() authtypes.AccountI { return authtypes.ProtoBaseAccount() },

			//
			// MINT
			//

			// For providing a custom inflation function for x/mint add here your
			// custom function that implements the minttypes.InflationCalculationFn
			// interface.

			// For providing a mock evm function for token module
			// mocks.ProvideEVMKeeper(),
			// mocks.ProvideICS20Keeper(),
			),
		)
	)

	consumer := append(depInjectOptions.Consumers[:],
		&appBuilder,
		&app.appCodec,
		&app.legacyAmino,
		&app.interfaceRegistry,
		&app.txConfig,
		&app.AccountKeeper,
		&app.BankKeeper,
		//&app.StakingKeeper,
		//&app.SlashingKeeper,
		//&app.MintKeeper,
		//&app.DistrKeeper,
		//&app.GovKeeper,
		&app.CrisisKeeper,
		//&app.CrisisKeeper,
		//&app.UpgradeKeeper,
		//&app.ParamsKeeper,
		//&app.AuthzKeeper,
		&app.EvidenceKeeper,
		&app.FeeGrantKeeper,
	)

	if err := depinject.Inject(appConfig, consumer...); err != nil {
		panic(err)
	}
	app.App = appBuilder.Build(logger, db, traceStore, baseAppOptions...)
	// load state streaming if enabled
	if _, _, err := streaming.LoadStreamingServices(app.App.BaseApp, appOpts, app.appCodec, logger, app.kvStoreKeys()); err != nil {
		logger.Error("failed to load state streaming", "err", err)
		os.Exit(1)
	}

	// initParamsKeeper(app.ParamsKeeper)

	/****  Module Options ****/

	app.ModuleManager.RegisterInvariants(app.CrisisKeeper)

	// RegisterUpgradeHandlers is used for registering any on-chain upgrades.
	// app.RegisterUpgradeHandlers()

	// add test gRPC service for testing gRPC queries in isolation
	testdata_pulsar.RegisterQueryServer(app.GRPCQueryRouter(), testdata_pulsar.QueryImpl{})

	// create the simulation manager and define the order of the modules for deterministic simulations
	//
	// NOTE: this is not required apps that don't use the simulator for fuzz testing
	// transactions
	//overrideModules := map[string]module.AppModuleSimulation{
	//	authtypes.ModuleName: auth.NewAppModule(
	//		app.appCodec,
	//		app.AccountKeeper,
	//		authsims.RandomGenesisAccounts,
	//		app.GetSubspace(authtypes.ModuleName),
	//	),
	//}
	app.sm = module.NewSimulationManagerFromAppModules(app.ModuleManager.Modules, nil)

	app.sm.RegisterStoreDecoders()
	app.SetInitChainer(app.InitChainer)

	// A custom InitChainer can be set if extra pre-init-genesis logic is required.
	// By default, when using app wiring enabled module, this is not required.
	// For instance, the upgrade module will set automatically the module version map in its init genesis thanks to app wiring.
	// However, when registering a module manually (i.e. that does not support app wiring), the module version map
	// must be set manually as follow. The upgrade module will de-duplicate the module version map.
	//
	// app.SetInitChainer(func(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	// 	app.UpgradeKeeper.SetModuleVersionMap(ctx, app.ModuleManager.GetVersionMap())
	// 	return app.App.InitChainer(ctx, req)
	// })

	if err := app.Load(loadLatest); err != nil {
		panic(err)
	}

	// TODO: RemoveValidator cdc in favor of appCodec once all modules are migrated.
	//appCodec := encodingConfig.Marshaler
	//cdc := encodingConfig.Amino
	//interfaceRegistry := encodingConfig.InterfaceRegistry
	//
	//bApp := baseapp.NewBaseApp(appName, logger, db, encodingConfig.TxConfig.TxDecoder(), baseAppOptions...)
	//bApp.SetCommitMultiStoreTracer(traceStore)
	//bApp.SetVersion(version.Version)
	//bApp.SetInterfaceRegistry(interfaceRegistry)
	//
	//keys := sdk.NewKVStoreKeys(
	//	authtypes.StoreKey,
	//	banktypes.StoreKey,
	//	slashingtypes.StoreKey,
	//	paramstypes.StoreKey,
	//	//gov.StoreKey,
	//	upgradetypes.StoreKey,
	//	feegrant.StoreKey,
	//	evidencetypes.StoreKey,
	//	//nodetypes.StoreKey,
	//)
	//tkeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
	//memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)
	//
	//app := &SimApp{
	//	BaseApp:           bApp,
	//	cdc:               cdc,
	//	appCodec:          appCodec,
	//	interfaceRegistry: interfaceRegistry,
	//	invCheckPeriod:    invCheckPeriod,
	//	keys:              keys,
	//	tkeys:             tkeys,
	//	memKeys:           memKeys,
	//}
	//
	//app.ParamsKeeper = initParamsKeeper(appCodec, cdc, keys[paramstypes.StoreKey], tkeys[paramstypes.TStoreKey])
	//
	//// set the BaseApp's parameter store
	//bApp.SetParamStore(app.ParamsKeeper.Subspace(baseapp.Paramspace).WithKeyTable(paramskeeper.ConsensusParamsKeyTable()))
	//
	//// add keepers
	//app.AccountKeeper = authkeeper.NewAccountKeeper(
	//	appCodec, keys[authtypes.StoreKey], app.GetSubspace(authtypes.ModuleName), authtypes.ProtoBaseAccount, maccPerms,
	//)
	//app.BankKeeper = bankkeeper.NewBaseKeeper(
	//	appCodec, keys[banktypes.StoreKey], app.AccountKeeper, app.GetSubspace(banktypes.ModuleName), app.ModuleAccountAddrs(),
	//)
	////app.NodeKeeper = nodekeeper.NewKeeper(appCodec, keys[nodetypes.StoreKey], app.GetSubspace(nodetypes.ModuleName))
	//app.SlashingKeeper = slashingkeeper.NewKeeper(
	//	appCodec, keys[slashingtypes.StoreKey], &app.NodeKeeper, app.GetSubspace(slashingtypes.ModuleName),
	//)
	//app.CrisisKeeper = crisiskeeper.NewKeeper(
	//	app.GetSubspace(crisistypes.ModuleName), invCheckPeriod, app.BankKeeper, authtypes.FeeCollectorName,
	//)
	//app.FeeGrantKeeper = feegrantkeeper.NewKeeper(appCodec, keys[feegrant.StoreKey], app.AccountKeeper)
	//app.UpgradeKeeper = upgradekeeper.NewKeeper(skipUpgradeHeights, keys[upgradetypes.StoreKey], appCodec, homePath, app.BaseApp)
	//
	//// create evidence keeper with router
	//EvidenceKeeper := evidencekeeper.NewKeeper(
	//	appCodec, keys[evidencetypes.StoreKey], &app.NodeKeeper, app.SlashingKeeper,
	//)
	//// If evidence needs to be handled for the app, set routes in router here and seal
	//app.EvidenceKeeper = *EvidenceKeeper
	//
	///*app.NodeKeeper = *app.NodeKeeper.SetHooks(
	//	stakingtypes.NewMultiStakingHooks(app.SlashingKeeper.Hooks()),
	//)*/
	//
	///****  Module Options ****/
	//
	//// NOTE: we may consider parsing `appOpts` inside module constructors. For the moment
	//// we prefer to be more strict in what arguments the modules expect.
	//var skipGenesisInvariants = cast.ToBool(appOpts.Get(crisis.FlagSkipGenesisInvariants))
	//
	//// NOTE: Any module instantiated in the module manager that is later modified
	//// must be passed by reference here.
	//app.mm = module.NewManager(
	//	//genutil.NewAppModule(app.AccountKeeper, app.NodeKeeper, app.BaseApp.DeliverTx, encodingConfig.TxConfig),
	//	auth.NewAppModule(appCodec, app.AccountKeeper, authsims.RandomGenesisAccounts),
	//	bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper),
	//	crisis.NewAppModule(&app.CrisisKeeper, skipGenesisInvariants),
	//	feegrantmodule.NewAppModule(appCodec, app.AccountKeeper, app.BankKeeper, app.FeeGrantKeeper, app.interfaceRegistry),
	//	//gov.NewAppModule(appCodec, app.govKeeper, app.AccountKeeper, app.BankKeeper),
	//	//cslashing.NewAppModule(appCodec, cslashing.NewKeeper(app.SlashingKeeper, app.NodeKeeper), app.AccountKeeper, app.BankKeeper, app.NodeKeeper),
	//	upgrade.NewAppModule(app.UpgradeKeeper),
	//	evidence.NewAppModule(app.EvidenceKeeper),
	//	params.NewAppModule(app.ParamsKeeper),
	//	//node.NewAppModule(appCodec, app.NodeKeeper),
	//)
	//
	//// During begin block slashing happens after distr.BeginBlocker so that
	//// there is nothing left over in the validator fee pool, so as to keep the
	//// CanWithdrawInvariant invariant.
	//// NOTE: staking module is required if HistoricalEntries param > 0
	//app.mm.SetOrderBeginBlockers(
	//	authtypes.ModuleName,
	//	//nodetypes.ModuleName,
	//	banktypes.ModuleName,
	//	slashingtypes.ModuleName,
	//	crisistypes.ModuleName,
	//	evidencetypes.ModuleName,
	//	feegrant.ModuleName,
	//	upgradetypes.ModuleName,
	//	paramstypes.ModuleName,
	//	genutiltypes.ModuleName,
	//)
	//
	//app.mm.SetOrderEndBlockers(
	//	authtypes.ModuleName,
	//	//nodetypes.ModuleName,
	//	banktypes.ModuleName,
	//	slashingtypes.ModuleName,
	//	crisistypes.ModuleName,
	//	evidencetypes.ModuleName,
	//	feegrant.ModuleName,
	//	upgradetypes.ModuleName,
	//	paramstypes.ModuleName,
	//	genutiltypes.ModuleName,
	//)
	//
	//// NOTE: The genutils module must occur after staking so that pools are
	//// properly initialized with tokens from genesis accounts.
	//// NOTE: Capability module must occur first so that it can initialize any capabilities
	//// so that other modules that want to create or claim capabilities afterwards in InitChain
	//// can do so safely.
	//app.mm.SetOrderInitGenesis(
	//	authtypes.ModuleName,
	//	//types.ModuleName,
	//	banktypes.ModuleName,
	//	slashingtypes.ModuleName,
	//	crisistypes.ModuleName,
	//	evidencetypes.ModuleName,
	//	feegrant.ModuleName,
	//	upgradetypes.ModuleName,
	//	paramstypes.ModuleName,
	//	genutiltypes.ModuleName,
	//)
	//
	//app.mm.SetOrderMigrations(
	//	authtypes.ModuleName,
	//	//nodetypes.ModuleName,
	//	banktypes.ModuleName,
	//	slashingtypes.ModuleName,
	//	crisistypes.ModuleName,
	//	evidencetypes.ModuleName,
	//	feegrant.ModuleName,
	//	upgradetypes.ModuleName,
	//	paramstypes.ModuleName,
	//	genutiltypes.ModuleName,
	//)
	//
	//app.mm.RegisterInvariants(&app.CrisisKeeper)
	//app.mm.RegisterRoutes(app.Router(), app.QueryRouter(), encodingConfig.Amino)
	//app.configurator = module.NewConfigurator(app.appCodec, app.MsgServiceRouter(), app.GRPCQueryRouter())
	//app.mm.RegisterServices(app.configurator)
	//
	//// add test gRPC service for testing gRPC queries in isolation
	//testdata.RegisterQueryServer(app.GRPCQueryRouter(), testdata.QueryImpl{})
	//
	//// create the simulation manager and define the order of the modules for deterministic simulations
	////
	//// NOTE: this is not required apps that don't use the simulator for fuzz testing
	//// transactions
	//app.sm = module.NewSimulationManager(
	//	auth.NewAppModule(appCodec, app.AccountKeeper, authsims.RandomGenesisAccounts),
	//	bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper),
	//	feegrantmodule.NewAppModule(appCodec, app.AccountKeeper, app.BankKeeper, app.FeeGrantKeeper, app.interfaceRegistry),
	//	//gov.NewAppModule(appCodec, app.govKeeper, app.AccountKeeper, app.BankKeeper),
	//	//cslashing.NewAppModule(appCodec, cslashing.NewKeeper(app.SlashingKeeper, app.NodeKeeper), app.AccountKeeper, app.BankKeeper, app.NodeKeeper),
	//	params.NewAppModule(app.ParamsKeeper),
	//	//node.NewAppModule(appCodec, app.NodeKeeper),
	//)
	//
	//app.sm.RegisterStoreDecoders()
	//
	//// initialize stores
	//app.MountKVStores(keys)
	//app.MountTransientStores(tkeys)
	//app.MountMemoryStores(memKeys)
	//
	//// initialize BaseApp
	//app.SetInitChainer(app.InitChainer)
	//app.SetBeginBlocker(app.BeginBlocker)
	//anteHandler, err := ante.NewAnteHandler(
	//	ante.HandlerOptions{
	//		AccountKeeper:   app.AccountKeeper,
	//		BankKeeper:      app.BankKeeper,
	//		SignModeHandler: encodingConfig.TxConfig.SignModeHandler(),
	//		FeegrantKeeper:  app.FeeGrantKeeper,
	//		SigGasConsumer:  ante.DefaultSigVerificationGasConsumer,
	//	},
	//)
	//if err != nil {
	//	panic(err)
	//}
	//app.SetAnteHandler(anteHandler)
	//app.SetEndBlocker(app.EndBlocker)
	//
	//if loadLatest {
	//	if err := app.LoadLatestVersion(); err != nil {
	//		tmos.Exit(err.Error())
	//	}
	//}

	return app
}

// MakeCodecs constructs the *std.Codec and *codec.LegacyAmino instances used by
// SimApp. It is useful for tests and clients who do not want to construct the
// full SimApp
//func MakeCodecs() (codec.Codec, *codec.LegacyAmino) {
//	encodingConfig := MakeEncodingConfig()
//	return encodingConfig.Marshaler, encodingConfig.Amino
//}

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

func (app *SimApp) TxConfig() client.TxConfig {
	return app.txConfig
}

// InterfaceRegistry returns SimApp's InterfaceRegistry
func (app *SimApp) InterfaceRegistry() types.InterfaceRegistry {
	return app.interfaceRegistry
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
func (app *SimApp) kvStoreKeys() map[string]*storetypes.KVStoreKey {
	keys := make(map[string]*storetypes.KVStoreKey)
	for _, k := range app.GetStoreKeys() {
		if kv, ok := k.(*storetypes.KVStoreKey); ok {
			keys[kv.Name()] = kv
		}
	}

	return keys
}

// GetSubspace returns a param subspace for a given module name.
//
// NOTE: This is solely to be used for testing purposes.
/*func (app *SimApp) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := app.ParamsKeeper.GetSubspace(moduleName)
	return subspace
}*/

// SimulationManager implements the SimulationApp interface
func (app *SimApp) SimulationManager() *module.SimulationManager {
	return app.sm
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *SimApp) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	app.App.RegisterAPIRoutes(apiSvr, apiConfig)
	// register swagger API in app.go so that other applications can override easily
	if err := server.RegisterSwaggerAPI(apiSvr.ClientCtx, apiSvr.Router, apiConfig.Swagger); err != nil {
		panic(err)
	}
}

// RegisterTxService implements the Application.RegisterTxService method.
func (app *SimApp) RegisterTxService(clientCtx client.Context) {
	authtx.RegisterTxService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.BaseApp.Simulate, app.interfaceRegistry)
}

// RegisterSwaggerAPI registers swagger route with API Server
//func RegisterSwaggerAPI(ctx client.Context, rtr *mux.Router) {
//	statikFS, err := fs.New()
//	if err != nil {
//		panic(err)
//	}
//
//	staticServer := http.FileServer(statikFS)
//	rtr.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", staticServer))
//}

// GetMaccPerms returns a copy of the module account permissions
//func GetMaccPerms() map[string][]string {
//	dupMaccPerms := make(map[string][]string)
//	for k, v := range maccPerms {
//		dupMaccPerms[k] = v
//	}
//	return dupMaccPerms
//}

// initParamsKeeper init params keeper and its subspaces
//func initParamsKeeper(appCodec codec.BinaryCodec, legacyAmino *codec.LegacyAmino, key, tkey sdk.StoreKey) paramskeeper.Keeper {
//	ParamsKeeper := paramskeeper.NewKeeper(appCodec, legacyAmino, key, tkey)
//
//	ParamsKeeper.Subspace(authtypes.ModuleName)
//	ParamsKeeper.Subspace(banktypes.ModuleName)
//	//ParamsKeeper.Subspace(nodetypes.ModuleName)
//	ParamsKeeper.Subspace(slashingtypes.ModuleName)
//	//ParamsKeeper.Subspace(crisistypes.ModuleName)
//
//	return ParamsKeeper
//}
