package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/iritamod/modules/slashing/types"
)

type msgServer struct {
	Keeper
}

func (m msgServer) UnjailValidator(goCtx context.Context, msg *types.MsgUnjailValidator) (*types.MsgUnjailValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := m.Keeper.HandleUnjail(ctx, *msg); err != nil {
		return nil, err
	}
	// NOTE: comment this because these event will be emitted directly at message execution.
	// TODO: remove this snippet.
	//ctx.EventManager().EmitEvent(
	//	sdk.NewEvent(
	//		sdk.EventTypeMessage,
	//		sdk.NewAttribute(sdk.AttributeKeyModule, slashingtypes.ModuleName),
	//		sdk.NewAttribute(sdk.AttributeKeySender, msg.Operator),
	//	),
	//)
	return &types.MsgUnjailValidatorResponse{}, nil
}

// NewMsgServerImpl returns an implementation of the bank MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}
