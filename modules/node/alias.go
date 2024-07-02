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
	EventTypeGrantNode       = types.EventTypeGrantNode
	EventTypeRevokeNode      = types.EventTypeRevokeNode
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
	NewMsgGrantNode             = types.NewMsgGrantNode
	NewMsgRevokeNode            = types.NewMsgRevokeNode
	ABCIValidatorUpdate         = keeper.ABCIValidatorUpdate
	GetGenesisStateFromAppState = types.GetGenesisStateFromAppState
	NewValidator                = types.NewValidator
)

type (
	MsgCreateValidator = types.MsgCreateValidator
	MsgUpdateValidator = types.MsgUpdateValidator
	MsgRemoveValidator = types.MsgRemoveValidator
	MsgGrantNode       = types.MsgGrantNode
	MsgRevokeNode      = types.MsgRevokeNode
	GenesisState       = types.GenesisState
	Validator          = types.Validator
	Node               = types.Node
	Params             = types.Params
	Keeper             = keeper.Keeper
)
