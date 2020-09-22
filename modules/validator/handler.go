package validator

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"gitlab.bianjie.ai/irita-pro/iritamod/modules/validator/keeper"
)

func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *MsgCreateValidator:
			return handleMsgCreateValidator(ctx, msg, k)

		case *MsgUpdateValidator:
			return handleMsgUpdateValidator(ctx, msg, k)

		case *MsgRemoveValidator:
			return handleMsgRemoveValidator(ctx, msg, k)

		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s message type: %T", ModuleName, msg)
		}
	}
}

func handleMsgCreateValidator(ctx sdk.Context, msg *MsgCreateValidator, k keeper.Keeper) (*sdk.Result, error) {
	validatorID, err := k.CreateValidator(ctx, *msg)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			EventTypeCreateValidator,
			sdk.NewAttribute(AttributeKeyValidator, validatorID.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Operator.String()),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgUpdateValidator(ctx sdk.Context, msg *MsgUpdateValidator, k keeper.Keeper) (*sdk.Result, error) {
	err := k.UpdateValidator(ctx, *msg)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			EventTypeUpdateValidator,
			sdk.NewAttribute(AttributeKeyValidator, msg.Id.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Operator.String()),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgRemoveValidator(ctx sdk.Context, msg *MsgRemoveValidator, k keeper.Keeper) (*sdk.Result, error) {
	err := k.RemoveValidator(ctx, *msg)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			EventTypeRemoveValidator,
			sdk.NewAttribute(AttributeKeyValidator, msg.Id.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Operator.String()),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
