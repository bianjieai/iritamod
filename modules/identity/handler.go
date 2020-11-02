package identity

import (
	"encoding/hex"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler creates an sdk.Handler for all the identity type messages
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *MsgCreateIdentity:
			return handleMsgCreateIdentity(ctx, k, msg)

		case *MsgUpdateIdentity:
			return handleMsgUpdateIdentity(ctx, k, msg)

		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s message type: %T", ModuleName, msg)
		}
	}
}

func handleMsgCreateIdentity(ctx sdk.Context, k Keeper, msg *MsgCreateIdentity) (*sdk.Result, error) {
	id, _ := hex.DecodeString(msg.Id)
	owner, _ := sdk.AccAddressFromBech32(msg.Owner)

	if err := k.CreateIdentity(ctx, id, msg.PubKey, msg.Certificate, msg.Credentials, owner); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			EventTypeCreateIdentity,
			sdk.NewAttribute(AttributeKeyID, msg.Id),
			sdk.NewAttribute(AttributeKeyOwner, msg.Owner),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Owner),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgUpdateIdentity(ctx sdk.Context, k Keeper, msg *MsgUpdateIdentity) (*sdk.Result, error) {
	id, _ := hex.DecodeString(msg.Id)
	owner, _ := sdk.AccAddressFromBech32(msg.Owner)

	if err := k.UpdateIdentity(ctx, id, msg.PubKey, msg.Certificate, msg.Credentials, owner); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			EventTypeUpdateIdentity,
			sdk.NewAttribute(AttributeKeyID, msg.Id),
			sdk.NewAttribute(AttributeKeyOwner, msg.Owner),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Owner),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
