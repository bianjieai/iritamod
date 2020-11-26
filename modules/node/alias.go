package node

// nolint

import (
	"github.com/bianjieai/iritamod/modules/node/keeper"
	"github.com/bianjieai/iritamod/modules/node/types"
)

const (
	ModuleName             = types.ModuleName
	StoreKey               = types.StoreKey
	QuerierRoute           = types.QuerierRoute
	RouterKey              = types.RouterKey
	QueryNode              = types.QueryNode
	QueryNodes             = types.QueryNodes
	EventTypeAddNode       = types.EventTypeAddNode
	EventTypeRemoveNode    = types.EventTypeRemoveNode
	AttributeValueCategory = types.AttributeValueCategory
	AttributeKeyID         = types.AttributeKeyID
)

var (
	NewKeeper           = keeper.NewKeeper
	NewQuerier          = keeper.NewQuerier
	ModuleCdc           = types.ModuleCdc
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis
	NewGenesisState     = types.NewGenesisState
)

type (
	Keeper           = keeper.Keeper
	Node             = types.Node
	GenesisState     = types.GenesisState
	MsgAddNode       = types.MsgAddNode
	MsgRemoveNode    = types.MsgRemoveNode
	QueryNodeParams  = types.QueryNodeParams
	QueryNodesParams = types.QueryNodesParams
)
