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

// NewMsgServerImpl returns an implementation of the node MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (m msgServer) CreateValidator(goCtx context.Context, msg *types.MsgCreateValidator) (*types.MsgCreateValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	id, err := m.Keeper.CreateValidator(ctx, *msg)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateValidator,
			sdk.NewAttribute(types.AttributeKeyValidator, id.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Operator),
		),
	})

	return &types.MsgCreateValidatorResponse{
		Id: id.String(),
	}, nil
}

func (m msgServer) UpdateValidator(goCtx context.Context, msg *types.MsgUpdateValidator) (*types.MsgUpdateValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := m.Keeper.UpdateValidator(ctx, *msg); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeUpdateValidator,
			sdk.NewAttribute(types.AttributeKeyValidator, msg.Id),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Operator),
		),
	})
	return &types.MsgUpdateValidatorResponse{}, nil
}

func (m msgServer) RemoveValidator(goCtx context.Context, msg *types.MsgRemoveValidator) (*types.MsgRemoveValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := m.Keeper.RemoveValidator(ctx, *msg); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeUpdateValidator,
			sdk.NewAttribute(types.AttributeKeyValidator, msg.Id),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Operator),
		),
	})
	return &types.MsgRemoveValidatorResponse{}, nil
}

func (m msgServer) GrantNode(goCtx context.Context, msg *types.MsgGrantNode) (*types.MsgGrantNodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	id, err := m.Keeper.AddNode(ctx, msg.Name, msg.Certificate)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeGrantNode,
			sdk.NewAttribute(types.AttributeKeyID, id.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Operator),
		),
	})

	return &types.MsgGrantNodeResponse{}, nil
}

func (m msgServer) RevokeNode(goCtx context.Context, msg *types.MsgRevokeNode) (*types.MsgRevokeNodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	id, _ := hex.DecodeString(msg.Id)
	if err := m.Keeper.RemoveNode(ctx, id); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeRevokeNode,
			sdk.NewAttribute(types.AttributeKeyID, msg.Id),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Operator),
		),
	})

	return &types.MsgRevokeNodeResponse{}, nil
}
