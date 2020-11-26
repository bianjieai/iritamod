package keeper

import (
	"context"
	"encoding/hex"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/iritamod/modules/node/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the bank MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (m msgServer) AddNode(goCtx context.Context, msg *types.MsgAddNode) (*types.MsgAddNodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	id, err := m.Keeper.AddNode(ctx, msg.Certificate)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeAddNode,
			sdk.NewAttribute(types.AttributeKeyID, id.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Operator),
		),
	})

	return &types.MsgAddNodeResponse{}, nil
}

func (m msgServer) RemoveNode(goCtx context.Context, msg *types.MsgRemoveNode) (*types.MsgRemoveNodeResponse, error) {
	id, _ := hex.DecodeString(msg.Id)
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := m.Keeper.RemoveNode(ctx, id); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeRemoveNode,
			sdk.NewAttribute(types.AttributeKeyID, msg.Id),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Operator),
		),
	})

	return &types.MsgRemoveNodeResponse{}, nil
}
