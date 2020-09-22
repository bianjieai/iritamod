package admin

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *MsgAddRoles:
			return handleMsgAddRoles(ctx, msg, k)

		case *MsgRemoveRoles:
			return handleMsgRemoveRoles(ctx, msg, k)

		case *MsgBlockAccount:
			return handleMsgBlockAccount(ctx, msg, k)

		case *MsgUnblockAccount:
			return handleMsgUnblockAccount(ctx, msg, k)

		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s message type: %T", ModuleName, msg)
		}
	}
}

func handleMsgAddRoles(ctx sdk.Context, msg *MsgAddRoles, k Keeper) (*sdk.Result, error) {
	err := k.AddRoles(ctx, msg.Address, msg.Operator, msg.Roles...)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			EventTypeAddRoles,
			sdk.NewAttribute(AttributeKeyAccount, msg.Address.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Operator.String()),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgRemoveRoles(ctx sdk.Context, msg *MsgRemoveRoles, k Keeper) (*sdk.Result, error) {
	err := k.RemoveRoles(ctx, msg.Address, msg.Operator, msg.Roles...)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			EventTypeRemoveRoles,
			sdk.NewAttribute(AttributeKeyAccount, msg.Address.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Operator.String()),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgBlockAccount(ctx sdk.Context, msg *MsgBlockAccount, k Keeper) (*sdk.Result, error) {
	err := k.BlockAccount(ctx, msg.Address)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			EventTypeBlockAccount,
			sdk.NewAttribute(AttributeKeyAccount, msg.Address.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Operator.String()),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgUnblockAccount(ctx sdk.Context, msg *MsgUnblockAccount, k Keeper) (*sdk.Result, error) {
	err := k.UnblockAccount(ctx, msg.Address)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			EventTypeUnblockAccount,
			sdk.NewAttribute(AttributeKeyAccount, msg.Address.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Operator.String()),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
