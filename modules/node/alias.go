package node

import (
	"github.com/bianjieai/iritamod/modules/node/keeper"
	"github.com/bianjieai/iritamod/modules/node/types"
)

const (
	ModuleName               = types.ModuleName
	StoreKey                 = types.StoreKey
	QuerierRoute             = types.QuerierRoute
	RouterKey                = types.RouterKey
	EventTypeCreateValidator = types.EventTypeCreateValidator
	EventTypeUpdateValidator = types.EventTypeUpdateValidator
	EventTypeRemoveValidator = types.EventTypeRemoveValidator
	EventTypeAddNode         = types.EventTypeAddNode
	EventTypeRemoveNode      = types.EventTypeRemoveNode
	AttributeKeyValidator    = types.AttributeKeyValidator
	AttributeKeyPubkey       = types.AttributeKeyPubkey
	AttributeKeyID           = types.AttributeKeyID
	AttributeValueCategory   = types.AttributeValueCategory
	DefaultParamspace        = keeper.DefaultParamspace
)

var (
	NewGenesisState             = types.NewGenesisState
	DefaultGenesisState         = types.DefaultGenesisState
	NewQuerier                  = keeper.NewQuerier
	NewKeeper                   = keeper.NewKeeper
	NewMsgCreateValidator       = types.NewMsgCreateValidator
	NewMsgUpdateValidator       = types.NewMsgUpdateValidator
	NewMsgRemoveValidator       = types.NewMsgRemoveValidator
	NewMsgAddNode               = types.NewMsgAddNode
	NewMsgRemoveNode            = types.NewMsgRemoveNode
	ABCIValidatorUpdate         = keeper.ABCIValidatorUpdate
	GetGenesisStateFromAppState = types.GetGenesisStateFromAppState
	NewValidator                = types.NewValidator
)

type (
	MsgCreateValidator = types.MsgCreateValidator
	MsgUpdateValidator = types.MsgUpdateValidator
	MsgRemoveValidator = types.MsgRemoveValidator
	MsgAddNode         = types.MsgAddNode
	MsgRemoveNode      = types.MsgRemoveNode
	GenesisState       = types.GenesisState
	Validator          = types.Validator
	Node               = types.Node
	Params             = types.Params
	Keeper             = keeper.Keeper
)
