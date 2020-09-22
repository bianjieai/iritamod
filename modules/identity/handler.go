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
	if err := k.CreateIdentity(ctx, msg.Id, msg.PubKey, msg.Certificate, msg.Credentials, msg.Owner); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			EventTypeCreateIdentity,
			sdk.NewAttribute(AttributeKeyID, hex.EncodeToString(msg.Id)),
			sdk.NewAttribute(AttributeKeyOwner, msg.Owner.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Owner.String()),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgUpdateIdentity(ctx sdk.Context, k Keeper, msg *MsgUpdateIdentity) (*sdk.Result, error) {
	if err := k.UpdateIdentity(ctx, msg.Id, msg.PubKey, msg.Certificate, msg.Credentials, msg.Owner); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			EventTypeUpdateIdentity,
			sdk.NewAttribute(AttributeKeyID, hex.EncodeToString(msg.Id)),
			sdk.NewAttribute(AttributeKeyOwner, msg.Owner.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Owner.String()),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
