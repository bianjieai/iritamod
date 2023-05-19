package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/iritamod/modules/layer2/types"
)

// CreateNFTs create native nft mappings on layer2 module.
func (m msgServer) CreateNFTs(goCtx context.Context, msg *types.MsgCreateNFTs) (*types.MsgCreateNFTsResponse, error) {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := m.Keeper.CreateNFTs(ctx, msg.SpaceId, msg.ClassId, msg.Tokens, sender); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateNFTs,
			sdk.NewAttribute(types.AttributeKeySpaceId, strconv.FormatUint(msg.SpaceId, 10)),
			sdk.NewAttribute(types.AttributeKeyClassIdForNFT, msg.ClassId),
			sdk.NewAttribute(types.AttributeKeyTokenAmountForNFT, strconv.Itoa(len(msg.Tokens))),
		),
	})

	return &types.MsgCreateNFTsResponse{}, nil
}

// UpdateNFTs update native nft mappings on layer2 module.
func (m msgServer) UpdateNFTs(goCtx context.Context, msg *types.MsgUpdateNFTs) (*types.MsgUpdateNFTsResponse, error) {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := m.Keeper.UpdateNFTs(ctx, msg.SpaceId, msg.ClassId, msg.Tokens, sender); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeUpdateNFTs,
			sdk.NewAttribute(types.AttributeKeySpaceId, strconv.FormatUint(msg.SpaceId, 10)),
			sdk.NewAttribute(types.AttributeKeyClassIdForNFT, msg.ClassId),
			sdk.NewAttribute(types.AttributeKeyTokenAmountForNFT, strconv.Itoa(len(msg.Tokens))),
		),
	})

	return &types.MsgUpdateNFTsResponse{}, nil
}

// DeleteNFTs delete native nft mappings on layer2 module.
// NOTE: this service is called as nfts on layer2 are burnt
func (m msgServer) DeleteNFTs(goCtx context.Context, msg *types.MsgDeleteNFTs) (*types.MsgDeleteNFTsResponse, error) {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := m.Keeper.DeleteNFTs(ctx, msg.SpaceId, msg.ClassId, msg.TokenIds, sender); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeDeleteNFTs,
			sdk.NewAttribute(types.AttributeKeySpaceId, strconv.FormatUint(msg.SpaceId, 10)),
			sdk.NewAttribute(types.AttributeKeyClassIdForNFT, msg.ClassId),
			sdk.NewAttribute(types.AttributeKeyTokenAmountForNFT, strconv.Itoa(len(msg.TokenIds))),
		),
	})

	return &types.MsgDeleteNFTsResponse{}, nil
}

// UpdateClassesForNFT update class mappings for nft
func (m msgServer) UpdateClassesForNFT(goCtx context.Context, msg *types.MsgUpdateClassesForNFT) (*types.MsgUpdateClassesForNFTResponse, error) {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := m.Keeper.UpdateClassesForNFT(ctx, msg.ClassUpdatesForNft, sender); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeUpdateClassesForNFT,
			sdk.NewAttribute(types.AttributeKeyClassesAmountForNFT, strconv.Itoa(len(msg.ClassUpdatesForNft))),
		),
	})
	return &types.MsgUpdateClassesForNFTResponse{}, nil
}

// DepositClassForNFT deposit a class from layer1 to layer2
func (m msgServer) DepositClassForNFT(goCtx context.Context, msg *types.MsgDepositClassForNFT) (*types.MsgDepositClassForNFTResponse, error) {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	recipient, err := sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := m.Keeper.DepositClassForNFT(ctx, msg.SpaceId, msg.ClassId, msg.BaseUri, recipient, sender); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeDepositClassForNFT,
			sdk.NewAttribute(types.AttributeKeySpaceId, strconv.FormatUint(msg.SpaceId, 10)),
			sdk.NewAttribute(types.AttributeKeyClassIdForNFT, msg.ClassId),
			sdk.NewAttribute(types.AttributeKeySender, msg.Sender),
			sdk.NewAttribute(types.AttributeKeyRecipient, msg.Recipient),
		),
	})

	return &types.MsgDepositClassForNFTResponse{}, nil
}

// WithdrawClassForNFT withdraw a class for nft from layer2 to layer1
// TODO： add space id
func (m msgServer) WithdrawClassForNFT(goCtx context.Context, msg *types.MsgWithdrawClassForNFT) (*types.MsgWithdrawClassForNFTResponse, error) {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := m.Keeper.WithdrawClassForNFT(ctx, msg.ClassId, owner, sender); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeWithdrawClassForNFT,
			sdk.NewAttribute(types.AttributeKeyClassIdForNFT, msg.ClassId),
			sdk.NewAttribute(types.AttributeKeySender, msg.Sender),
			sdk.NewAttribute(types.AttributeKeyRecipient, msg.Owner),
		),
	})

	return &types.MsgWithdrawClassForNFTResponse{}, nil
}

// DepositTokenForNFT deposit a nft from layer2 to layer1
func (m msgServer) DepositTokenForNFT(goCtx context.Context, msg *types.MsgDepositTokenForNFT) (*types.MsgDepositTokenForNFTResponse, error) {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := m.Keeper.DepositTokenForNFT(ctx, msg.SpaceId, msg.ClassId, msg.TokenId, sender); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeDepositTokenForNFT,
			sdk.NewAttribute(types.AttributeKeySpaceId, strconv.FormatUint(msg.SpaceId, 10)),
			sdk.NewAttribute(types.AttributeKeyClassIdForNFT, msg.ClassId),
			sdk.NewAttribute(types.AttributeKeyTokenIdForNFT, msg.TokenId),
			sdk.NewAttribute(types.AttributeKeySender, msg.Sender),
		),
	})

	return &types.MsgDepositTokenForNFTResponse{}, nil
}

// TODO： suggestion: token metadata
func (m msgServer) WithdrawTokenForNFT(goCtx context.Context, msg *types.MsgWithdrawTokenForNFT) (*types.MsgWithdrawTokenForNFTResponse, error) {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := m.Keeper.WithdrawTokenForNFT(ctx, msg.SpaceId, msg.ClassId, msg.TokenId, msg.Name, msg.Uri, msg.UriHash, msg.Data, owner, sender); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeWithdrawTokenForNFT,
			sdk.NewAttribute(types.AttributeKeySpaceId, strconv.FormatUint(msg.SpaceId, 10)),
			sdk.NewAttribute(types.AttributeKeyClassIdForNFT, msg.ClassId),
			sdk.NewAttribute(types.AttributeKeyTokenIdForNFT, msg.TokenId),
			sdk.NewAttribute(types.AttributeKeyRecipient, msg.Owner),
		),
	})

	return &types.MsgWithdrawTokenForNFTResponse{}, nil
}
