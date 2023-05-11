package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	TypeMsgCreateNFTs = "create_nfts"
	TypeMsgUpdateNFTs = "update_nfts"
	TypeMsgDeleteNFTs = "delete_nfts"
	TypeMsgDepositClassForNFT = "deposit_class_for_nft"
	TypeMsgWithdrawClassForNFT = "withdraw_class_for_nft"
	TypeMsgDepositTokenForNFT = "deposit_token_for_nft"
	TypeMsgWithdrawTokenForNFT = "withdraw_token_for_nft"
	TypeMsgCreateL2Space = "create_l2_space"
	TypeMsgTransferL2Space = "transfer_l2_space"
	TypeMsgCreateL2Record = "create_l2_record"
)

var (
	_ sdk.Msg = &MsgCreateNFTs{}
	_ sdk.Msg = &MsgUpdateNFTs{}
	_ sdk.Msg = &MsgDeleteNFTs{}
	_ sdk.Msg = &MsgDepositClassForNFT{}
	_ sdk.Msg = &MsgWithdrawClassForNFT{}
	_ sdk.Msg = &MsgDepositTokenForNFT{}
	_ sdk.Msg = &MsgWithdrawTokenForNFT{}
	_ sdk.Msg = &MsgCreateL2Space{}
	_ sdk.Msg = &MsgTransferL2Space{}
	_ sdk.Msg = &MsgCreateL2Record{}
)