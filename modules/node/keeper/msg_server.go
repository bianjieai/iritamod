package keeper

import (
	"context"
	"encoding/hex"

	"github.com/cometbft/cometbft/crypto/tmhash"
	ctmbytes "github.com/cometbft/cometbft/libs/bytes"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/bianjieai/iritamod/modules/node/types"
)

type msgServer struct {
	k Keeper
}

// NewMsgServerImpl returns an implementation of the node MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{k: keeper}
}

var _ types.MsgServer = msgServer{}

func (m msgServer) CreateValidator(goCtx context.Context, msg *types.MsgCreateValidator) (*types.MsgCreateValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	id := ctmbytes.HexBytes(tmhash.Sum(msg.GetSignBytes()))
	if err := m.k.CreateValidator(ctx,
		id,
		msg.Name,
		msg.Certificate,
		nil,
		msg.Power,
		msg.Description,
		msg.Operator,
	); err != nil {
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
	id, err := hex.DecodeString(msg.Id)
	if err != nil {
		return &types.MsgUpdateValidatorResponse{}, types.ErrInvalidValidatorID
	}

	if err := m.k.UpdateValidator(ctx,
		id,
		msg.Name,
		msg.Certificate,
		msg.Power,
		msg.Description,
		msg.Operator,
	); err != nil {
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
	id, err := hex.DecodeString(msg.Id)
	if err != nil {
		return &types.MsgRemoveValidatorResponse{}, types.ErrInvalidValidatorID
	}

	if err := m.k.RemoveValidator(ctx, id, msg.Operator); err != nil {
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

	id, err := m.k.AddNode(ctx, msg.Name, msg.Certificate)
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
	if err := m.k.RemoveNode(ctx, id); err != nil {
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

func (m msgServer) UpdateParams(goCtx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	if m.k.authority != msg.Authority {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrUnauthorized,
			"invalid authority; expected %s, got %s",
			m.k.authority,
			msg.Authority,
		)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := m.k.SetParams(ctx, msg.Params); err != nil {
		return nil, err
	}
	return &types.MsgUpdateParamsResponse{}, nil
}
