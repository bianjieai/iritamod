package params

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"

	"gitlab.bianjie.ai/irita-pro/iritamod/modules/params/types"
)

func NewHandler(k paramskeeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgUpdateParams:
			return handleParametersUpdate(ctx, k, msg)

		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s message type: %T", ModuleName, msg)
		}
	}
}

func handleParametersUpdate(ctx sdk.Context, k paramskeeper.Keeper, msg *types.MsgUpdateParams) (*sdk.Result, error) {
	var changeEvents []sdk.Attribute
	for _, c := range msg.Changes {
		ss, ok := k.GetSubspace(c.Subspace)
		if !ok {
			return nil, sdkerrors.Wrap(ErrUnknownSubspace, c.Subspace)
		}

		k.Logger(ctx).Info(
			fmt.Sprintf("attempt to set new parameter value; key: %s, value: %s", c.Key, c.Value),
		)

		if !ss.Has(ctx, []byte(c.Key)) {
			return nil, sdkerrors.Wrapf(ErrUnknownKey, c.Key)
		}

		if err := ss.Update(ctx, []byte(c.Key), []byte(c.Value)); err != nil {
			return nil, sdkerrors.Wrapf(ErrSettingParameter, "key: %s, value: %s, err: %s", c.Key, c.Value, err.Error())
		}

		changeEvents = append(changeEvents, sdk.NewAttribute(AttributeKeyParamKey, c.Key))
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			EventTypeUpdateParams,
			changeEvents...,
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Operator.String()),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
