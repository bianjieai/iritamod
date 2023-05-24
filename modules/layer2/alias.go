package layer2

import (
	"github.com/bianjieai/iritamod/modules/layer2/types"
)

type (
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
