package validator

import (
	"gitlab.bianjie.ai/irita-pro/iritamod/modules/validator/keeper"
	"gitlab.bianjie.ai/irita-pro/iritamod/modules/validator/types"
)

const (
	ModuleName               = types.ModuleName
	StoreKey                 = types.StoreKey
	QuerierRoute             = types.QuerierRoute
	RouterKey                = types.RouterKey
	EventTypeCreateValidator = types.EventTypeCreateValidator
	EventTypeUpdateValidator = types.EventTypeUpdateValidator
	EventTypeRemoveValidator = types.EventTypeRemoveValidator
	AttributeKeyValidator    = types.AttributeKeyValidator
	AttributeKeyPubkey       = types.AttributeKeyPubkey
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
	ABCIValidatorUpdate         = keeper.ABCIValidatorUpdate
	GetGenesisStateFromAppState = types.GetGenesisStateFromAppState
	NewValidator                = types.NewValidator
)

type (
	MsgCreateValidator = types.MsgCreateValidator
	MsgUpdateValidator = types.MsgUpdateValidator
	MsgRemoveValidator = types.MsgRemoveValidator
	GenesisState       = types.GenesisState
	Validator          = types.Validator
	Params             = types.Params
	Keeper             = keeper.Keeper
)
