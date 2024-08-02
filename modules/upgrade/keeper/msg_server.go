package keeper

import (
	"context"

	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"iritamod.bianjie.ai/modules/upgrade/types"
)

type msgServer struct {
	*Keeper
}

var _ types.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the bank MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper *Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

func (m msgServer) UpgradeSoftware(goCtx context.Context, msg *types.MsgUpgradeSoftware) (*types.MsgUpgradeSoftwareResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	plan := upgradetypes.Plan{
		Name:   msg.Name,
		Height: msg.Height,
		Info:   msg.Info,
	}
	if err := m.ScheduleUpgrade(ctx, plan); err != nil {
		return nil, err
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Operator),
		),
	)
	return &types.MsgUpgradeSoftwareResponse{}, nil
}

func (m msgServer) CancelUpgrade(goCtx context.Context, msg *types.MsgCancelUpgrade) (*types.MsgCancelUpgradeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	m.ClearUpgradePlan(ctx)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Operator),
		),
	)
	return &types.MsgCancelUpgradeResponse{}, nil
}
