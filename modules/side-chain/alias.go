package sidechain

import (
	"github.com/bianjieai/iritamod/modules/node/keeper"
	"github.com/bianjieai/iritamod/modules/side-chain/types"
)

const (
	ModuleName = types.ModuleName
	StoreKey   = types.StoreKey
	RouterKey  = types.RouterKey

	EventTypeCreateSpace   = types.EventTypeCreateSpace
	EventTypeTransferSpace = types.EventTypeTransferSpace
	EventTypeCreateRecord  = types.EventTypeCreateRecord

	AttributeKeySender       = types.AttributeKeySender
	AttributeKeyOwner        = types.AttributeKeyOwner
	AttributeKeyRecipient    = types.AttributeKeyRecipient
	AttributeKeySpaceId      = types.AttributeKeySpaceId
	AttributeKeyRecordHeight = types.AttributeKeyRecordHeight
)

var (
	NewKeeper       = keeper.NewKeeper
	ModuleCdc       = types.ModuleCdc
	DefaultGenesis  = types.DefaultGenesisState
	ValidateGenesis = types.ValidateGenesis
	NewGenesisState = types.NewGenesisState
)

type (
	Keeper               = keeper.Keeper
	GenesisState         = types.GenesisState
	MsgCreateSpace       = types.MsgCreateSpace
	MsgTransferSpace     = types.MsgTransferSpace
	MsgCreateBlockHeader = types.MsgCreateBlockHeader
)
