package slashing

// nolint

import (
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"

	"github.com/bianjieai/iritamod/modules/slashing/keeper"
	"github.com/bianjieai/iritamod/modules/slashing/types"
)

var (
	// NewKeeper alias keeper/NewKeeper
	NewKeeper = keeper.NewKeeper
	// NewMsgUnjailValidator alias types/NewMsgUnjailValidator
	NewMsgUnjailValidator = types.NewMsgUnjailValidator

	// RouterKey alias cosmos-sdk/x/slashing RouterKey
	RouterKey = slashingtypes.RouterKey
	// ModuleName alias cosmos-sdk/x/slashing ModuleName
	ModuleName = slashingtypes.ModuleName
	// QuerierRoute alias cosmos-sdk/x/slashing QuerierRoute
	QuerierRoute = slashingtypes.QuerierRoute
	// DefaultGenesisState alias cosmos-sdk/x/slashing DefaultGenesisState
	DefaultGenesisState = slashingtypes.DefaultGenesisState
	// ValidateGenesis alias cosmos-sdk/x/slashing ValidateGenesis
	ValidateGenesis = slashingtypes.ValidateGenesis
	// RegisterQueryHandlerClient alias cosmos-sdk/x/slashing RegisterQueryHandlerClient
	RegisterQueryHandlerClient = slashingtypes.RegisterQueryHandlerClient
	// NewQueryClient alias cosmos-sdk/x/slashing NewQueryClient
	NewQueryClient = slashingtypes.NewQueryClient
	// RegisterQueryServer alias cosmos-sdk/x/slashing RegisterQueryServer
	RegisterQueryServer = slashingtypes.RegisterQueryServer
)

type (
	// MsgUnjailValidator alias types/MsgUnjailValidator
	MsgUnjailValidator = types.MsgUnjailValidator
	// Keeper alias keeper/Keeper
	Keeper = keeper.Keeper

	// GenesisState alias cosmos-sdk/x/slashing GenesisState
	GenesisState = slashingtypes.GenesisState
	// AccountKeeper alias cosmos-sdk/x/slashing AccountKeeper
	AccountKeeper = slashingtypes.AccountKeeper
	// BankKeeper alias cosmos-sdk/x/slashing BankKeeper
	BankKeeper = slashingtypes.BankKeeper
	// StakingKeeper alias cosmos-sdk/x/slashing StakingKeeper
	StakingKeeper = slashingtypes.StakingKeeper
)
