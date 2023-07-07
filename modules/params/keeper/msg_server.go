package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/iritamod/modules/params/types"
)

type msgServer struct {
	k Keeper
}

// NewMsgServerImpl returns an implementation of the bank MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{k: keeper}
}

var _ types.MsgServer = msgServer{}

// UpdateParams implements the MsgServer.UpdateParams method.
func (m msgServer) UpdateParams(
	goCtx context.Context,
	msg *types.MsgUpdateParams,
) (*types.MsgUpdateParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	messages, err := msg.GetMsgs()
	if err != nil {
		return nil, err
	}

	if err := m.k.UpdateParams(ctx, messages); err != nil {
		return nil, err
	}
	return &types.MsgUpdateParamsResponse{}, nil
}
