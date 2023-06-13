package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/iritamod/modules/side-chain/types"
)

type msgServer struct {
	Keeper
}

func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

// CreateSpace creates a new space for  user
func (m msgServer) CreateSpace(goCtx context.Context, msg *types.MsgCreateSpace) (*types.MsgCreateSpaceResponse, error) {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	spaceId, err := m.Keeper.CreateSpace(ctx, msg.Name, msg.Uri, sender)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateSpace,
			sdk.NewAttribute(types.AttributeKeySpaceId, strconv.FormatUint(spaceId, 10)),
			sdk.NewAttribute(types.AttributeKeySender, msg.Sender),
		),
	})

	return &types.MsgCreateSpaceResponse{SpaceId: spaceId}, nil
}

// TransferSpace transfers the ownership of a space
func (m msgServer) TransferSpace(goCtx context.Context, msg *types.MsgTransferSpace) (*types.MsgTransferSpaceResponse, error) {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}
	recipient, err := sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	err = m.Keeper.TransferSpace(ctx, msg.SpaceId, sender, recipient)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeTransferSpace,
			sdk.NewAttribute(types.AttributeKeySpaceId, strconv.FormatUint(msg.SpaceId, 10)),
			sdk.NewAttribute(types.AttributeKeySender, msg.Sender),
			sdk.NewAttribute(types.AttributeKeyRecipient, msg.Recipient),
		),
	})

	return &types.MsgTransferSpaceResponse{}, nil
}

// CreateBlockHeader creates a layer 2 record
func (m msgServer) CreateBlockHeader(goCtx context.Context, msg *types.MsgCreateBlockHeader) (*types.MsgCreateBlockHeaderResponse, error) {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	err = m.Keeper.CreateBlockHeader(ctx, msg.SpaceId, msg.Height, msg.Header, sender)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateRecord,
			sdk.NewAttribute(types.AttributeKeySpaceId, strconv.FormatUint(msg.SpaceId, 10)),
			sdk.NewAttribute(types.AttributeKeyRecordHeight, strconv.FormatUint(msg.Height, 10)),
		),
	})

	return &types.MsgCreateBlockHeaderResponse{}, nil
}
