package mint

import (
	"github.com/bianjieai/iritamod/modules/mint/keeper"
	"github.com/bianjieai/iritamod/modules/mint/types"
)

const (
	ModuleName             = types.ModuleName
	StoreKey               = types.StoreKey
	QuerierRoute           = types.QuerierRoute
	RouterKey              = types.RouterKey
	EventTypeMint          = types.EventTypeMint
	AttributeValueCategory = types.AttributeValueCategory
)

var (
	NewGenesisState     = types.NewGenesisState
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis
	NewKeeper           = keeper.NewKeeper
)

type (
	MsgMint      = types.MsgMint
	Keeper       = keeper.Keeper
	GenesisState = types.GenesisState
)
