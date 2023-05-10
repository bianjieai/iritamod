package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	TypeMsgCreateNFTs = "create_nfts"
	TypeMsgUpdateNFTs = "update_nfts"
	TypeMsgDeleteNFTs = "delete_nfts"
	TypeMsgDepositClass = "deposit_class"
	TypeMsgWithdrawClass = "withdraw_class"
	TypeMsgDepositNFT = "deposit_nft"
	TypeMsgWithdrawNFT = "withdraw_nft"
	TypeMsgCreateL2Space = "create_l2_space"
	TypeMsgTransferL2Space = "transfer_l2_space"
)

var (
	_ sdk.Msg = &MsgCreateNFTs{}
	_ sdk.Msg = &MsgUpdateNFTs{}
	_ sdk.Msg = &MsgDeleteNFTs{}
	_ sdk.Msg = &MsgDepositClass{}
	_ sdk.Msg = &MsgWithdrawClass{}
	_ sdk.Msg = &MsgDepositNFT{}
	_ sdk.Msg = &MsgWithdrawNFT{}
	_ sdk.Msg = &MsgCreateL2Space{}
	_ sdk.Msg = &MsgTransferL2Space{}
)