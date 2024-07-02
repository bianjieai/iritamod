package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"

	"github.com/bianjieai/iritamod/modules/params/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the bank MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper paramskeeper.Keeper) types.MsgServer {
	return &msgServer{Keeper: Keeper{keeper}}
}

var _ types.MsgServer = msgServer{}

func (m msgServer) UpdateParams(goCtx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	changeEvents, err := m.Keeper.UpdateParams(ctx, msg.Changes)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeUpdateParams,
			changeEvents...,
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Operator),
		),
	})
	return &types.MsgUpdateParamsResponse{}, nil
}
