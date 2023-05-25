package layer2

import (
	"github.com/bianjieai/iritamod/modules/layer2/types"
	"github.com/bianjieai/iritamod/modules/node/keeper"
)

const (
	ModuleName = types.ModuleName
	StoreKey   = types.StoreKey
	RouterKey  = types.RouterKey

	EventTypeCreateL2Space       = types.EventTypeCreateL2Space
	EventTypeTransferL2Space     = types.EventTypeTransferL2Space
	EventTypeCreateL2Record      = types.EventTypeCreateL2Record
	EventTypeDepositClassForNFT  = types.EventTypeDepositClassForNFT
	EventTypeWithdrawClassForNFT = types.EventTypeWithdrawClassForNFT
	EventTypeDepositTokenForNFT  = types.EventTypeDepositTokenForNFT
	EventTypeWithdrawTokenForNFT = types.EventTypeWithdrawTokenForNFT

	AttributeKeySender              = types.AttributeKeySender
	AttributeKeyOwner               = types.AttributeKeyOwner
	AttributeKeyRecipient           = types.AttributeKeyRecipient
	AttributeKeySpaceId             = types.AttributeKeySpaceId
	AttributeKeyRecordHeight        = types.AttributeKeyRecordHeight
	AttributeKeyClassIdForNFT       = types.AttributeKeyClassIdForNFT
	AttributeKeyTokenIdForNFT       = types.AttributeKeyTokenIdForNFT
	AttributeKeyClassesAmountForNFT = types.AttributeKeyClassesAmountForNFT
	AttributeKeyTokenAmountForNFT   = types.AttributeKeyTokenAmountForNFT
)

var (
	NewKeeper       = keeper.NewKeeper
	ModuleCdc       = types.ModuleCdc
	DefaultGenesis  = types.DefaultGenesisState
	ValidateGenesis = types.ValidateGenesis
	NewGenesisState = types.NewGenesisState
)

type (
	Keeper                 = keeper.Keeper
	GenesisState           = types.GenesisState
	MsgCreateL2Space       = types.MsgCreateL2Space
	MsgTransferL2Space     = types.MsgTransferL2Space
	MsgCreateL2BlockHeader = types.MsgCreateL2BlockHeader
	MsgCreateNFTs          = types.MsgCreateNFTs
	MsgUpdateNFTs          = types.MsgUpdateNFTs
	MsgDeleteNFTs          = types.MsgDeleteNFTs
	MsgUpdateClassesForNFT = types.MsgUpdateClassesForNFT
	MsgDepositClassForNFT  = types.MsgDepositClassForNFT
	MsgWithdrawClassForNFT = types.MsgWithdrawClassForNFT
	MsgDepositTokenForNFT  = types.MsgDepositTokenForNFT
	MsgWithdrawTokenForNFT = types.MsgWithdrawTokenForNFT
)
