package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"

	"github.com/bianjieai/iritamod/modules/slashing/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	k Keeper
}

// NewMsgServerImpl returns an implementation of the bank MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{k: keeper}
}

func (m msgServer) UnjailValidator(
	goCtx context.Context,
	msg *types.MsgUnjailValidator,
) (*types.MsgUnjailValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := m.k.HandleUnjail(ctx, *msg); err != nil {
		return nil, err
	}
	return &types.MsgUnjailValidatorResponse{}, nil
}

// UpdateParams updates the slashing params.
// WARN： must register perm access control for this method.
func (m msgServer) UpdateParams(
	goCtx context.Context,
	msg *types.MsgUpdateParams,
) (*types.MsgUpdateParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	params := slashingtypes.Params{
		SignedBlocksWindow:      msg.Params.SignedBlocksWindow,
		MinSignedPerWindow:      msg.Params.MinSignedPerWindow,
		DowntimeJailDuration:    msg.Params.DowntimeJailDuration,
		SlashFractionDoubleSign: msg.Params.SlashFractionDoubleSign,
		SlashFractionDowntime:   msg.Params.SlashFractionDowntime,
	}

	if err := m.k.SetParams(ctx, params); err != nil {
		return nil, err
	}
	return &types.MsgUpdateParamsResponse{}, nil
}
