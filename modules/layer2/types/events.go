package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"
)

const (
	EventTypeCreateL2Space = "create_l2_space"
	EventTypeTransferL2Space = "transfer_l2_space"

	AttributeKeySender = "sender"
	AttributeKeyOwner = "owner"
	AttributeKeySpaceId = "space_id"
)

func NewCreateL2SpaceEvent(sender sdk.AccAddress, spaceId int64) sdk.Event {
	return sdk.NewEvent(
		EventTypeCreateL2Space,
		sdk.NewAttribute(AttributeKeySender, sender.String()),
		sdk.NewAttribute(AttributeKeySpaceId, strconv.Itoa(int(spaceId))),
	)
}