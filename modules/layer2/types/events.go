package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"
)

const (
	EventTypeCreateL2Space = "create_l2_space"
	EventTypeTransferL2Space = "transfer_l2_space"

	AttributeKeySender = "sender"
	AttributeKeySpaceId = "space_id"

	EventTypeCreateNFTs = "create_nfts"
	EventTypeUpdateNFTs = "update_nfts"
	EventTypeDeleteNFTs = "delete_nfts"

	AttributeKeyClassIdForNFT = "class_id_for_nft"
	AttributeKeyTokenIdForNFT = "token_id_for_nft"
	AttributeKeyTokenAmountForNFT = "token_amount_for_nft"

)

func NewCreateL2SpaceEvent(sender sdk.AccAddress, spaceId int64) sdk.Event {
	return sdk.NewEvent(
		EventTypeCreateL2Space,
		sdk.NewAttribute(AttributeKeySender, sender.String()),
		sdk.NewAttribute(AttributeKeySpaceId, strconv.Itoa(int(spaceId))),
	)
}