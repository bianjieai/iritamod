package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/iritamod/modules/layer2/types"
)

type msgServer struct {
	Keeper
}

func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

// CreateL2Space creates a new space for l2 user
func (m msgServer) CreateL2Space(goCtx context.Context, msg *types.MsgCreateL2Space) (*types.MsgCreateL2SpaceResponse, error) {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	spaceId, err := m.Keeper.CreateL2Space(ctx, msg.Name, msg.Uri, sender)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateL2Space,
			sdk.NewAttribute(types.AttributeKeySpaceId, strconv.FormatUint(spaceId, 10)),
			sdk.NewAttribute(types.AttributeKeySender, msg.Sender),
		),
	})

	return &types.MsgCreateL2SpaceResponse{SpaceId: spaceId}, nil
}

// TransferL2Space transfers the ownership of a space
func (m msgServer) TransferL2Space(goCtx context.Context, msg *types.MsgTransferL2Space) (*types.MsgTransferL2SpaceResponse, error) {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}
	recipient, err := sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	err = m.Keeper.TransferL2Space(ctx, msg.SpaceId, sender, recipient)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeTransferL2Space,
			sdk.NewAttribute(types.AttributeKeySpaceId, strconv.FormatUint(msg.SpaceId, 10)),
			sdk.NewAttribute(types.AttributeKeySender, msg.Sender),
			sdk.NewAttribute(types.AttributeKeyRecipient, msg.Recipient),
		),
	})

	return &types.MsgTransferL2SpaceResponse{}, nil
}

// CreateL2BlockHeader creates a layer 2 record
func (m msgServer) CreateL2BlockHeader(goCtx context.Context, msg *types.MsgCreateL2BlockHeader) (*types.MsgCreateL2BlockHeaderResponse, error) {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	err = m.Keeper.CreateL2BlockHeader(ctx, msg.SpaceId, msg.Height, msg.Header, sender)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateL2Record,
			sdk.NewAttribute(types.AttributeKeySpaceId, strconv.FormatUint(msg.SpaceId, 10)),
			sdk.NewAttribute(types.AttributeKeyRecordHeight, strconv.FormatUint(msg.Height, 10)),
		),
	})

	return &types.MsgCreateL2BlockHeaderResponse{}, nil
}
