package keeper

import (
	"context"
	"encoding/hex"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/iritamod/modules/identity/types"
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

func (m msgServer) CreateIdentity(goCtx context.Context, msg *types.MsgCreateIdentity) (*types.MsgCreateIdentityResponse, error) {
	id, _ := hex.DecodeString(msg.Id)
	owner, _ := sdk.AccAddressFromBech32(msg.Owner)

	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := m.Keeper.CreateIdentity(ctx, id, msg.PubKey, msg.Certificate, msg.Credentials, msg.Data, owner); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateIdentity,
			sdk.NewAttribute(types.AttributeKeyID, msg.Id),
			sdk.NewAttribute(types.AttributeKeyOwner, msg.Owner),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Owner),
		),
	})
	return &types.MsgCreateIdentityResponse{}, nil
}

func (m msgServer) UpdateIdentity(goCtx context.Context, msg *types.MsgUpdateIdentity) (*types.MsgUpdateIdentityResponse, error) {
	id, _ := hex.DecodeString(msg.Id)
	owner, _ := sdk.AccAddressFromBech32(msg.Owner)

	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := m.Keeper.UpdateIdentity(ctx, id, msg.PubKey, msg.Certificate, msg.Credentials, msg.Data, owner); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeUpdateIdentity,
			sdk.NewAttribute(types.AttributeKeyID, msg.Id),
			sdk.NewAttribute(types.AttributeKeyOwner, msg.Owner),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Owner),
		),
	})
	return &types.MsgUpdateIdentityResponse{}, nil
}
