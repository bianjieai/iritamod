package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/iritamod/modules/perm/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the perm MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (m msgServer) AssignRoles(goCtx context.Context, msg *types.MsgAssignRoles) (*types.MsgAssignRolesResponse, error) {
	addr, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, err
	}

	operator, err := sdk.AccAddressFromBech32(msg.Operator)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := m.Keeper.Authorize(ctx, addr, operator, msg.Roles...); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeAssignRoles,
			sdk.NewAttribute(types.AttributeKeyAccount, msg.Address),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Operator),
		),
	})
	return &types.MsgAssignRolesResponse{}, nil
}

func (m msgServer) UnassignRoles(goCtx context.Context, msg *types.MsgUnassignRoles) (*types.MsgUnassignRolesResponse, error) {
	addr, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, err
	}

	operator, err := sdk.AccAddressFromBech32(msg.Operator)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := m.Keeper.Unauthorize(ctx, addr, operator, msg.Roles...); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeUnassignRoles,
			sdk.NewAttribute(types.AttributeKeyAccount, msg.Address),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Operator),
		),
	})
	return &types.MsgUnassignRolesResponse{}, nil
}

func (m msgServer) BlockAccount(goCtx context.Context, msg *types.MsgBlockAccount) (*types.MsgBlockAccountResponse, error) {
	addr, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := m.Keeper.Block(ctx, addr); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeBlockAccount,
			sdk.NewAttribute(types.AttributeKeyAccount, msg.Address),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Operator),
		),
	})
	return &types.MsgBlockAccountResponse{}, nil
}

func (m msgServer) UnblockAccount(goCtx context.Context, msg *types.MsgUnblockAccount) (*types.MsgUnblockAccountResponse, error) {
	addr, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := m.Keeper.Unblock(ctx, addr); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeUnblockAccount,
			sdk.NewAttribute(types.AttributeKeyAccount, msg.Address),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Operator),
		),
	})
	return &types.MsgUnblockAccountResponse{}, nil
}
