package keeper

import (
	"context"
	"github.com/bianjieai/iritamod/modules/layer2/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"
)

// CreateNFTs create native nft mappings on layer2 module.
func (k Keeper) CreateNFTs(goCtx context.Context, msg *types.MsgCreateNFTs) (*types.MsgCreateNFTsResponse, error) {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.CreateTokensForNFT(ctx, msg.SpaceId, msg.ClassId, msg.Nfts, sender); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateNFTs,
			sdk.NewAttribute(types.AttributeKeySpaceId, strconv.FormatUint(msg.SpaceId, 10)),
			sdk.NewAttribute(types.AttributeKeyClassIdForNFT, msg.ClassId),
			sdk.NewAttribute(types.AttributeKeyTokenAmountForNFT, strconv.Itoa(len(msg.Nfts))),
		),
	})

	return &types.MsgCreateNFTsResponse{}, nil
}

// UpdateNFTs update native nft mappings on layer2 module.
func (k Keeper) UpdateNFTs(goCtx context.Context, msg *types.MsgUpdateNFTs) (*types.MsgUpdateNFTsResponse, error) {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := k.UpdateTokensForNFT(ctx, msg.SpaceId, msg.ClassId, msg.Nfts, sender); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeUpdateNFTs,
			sdk.NewAttribute(types.AttributeKeySpaceId, strconv.FormatUint(msg.SpaceId, 10)),
			sdk.NewAttribute(types.AttributeKeyClassIdForNFT, msg.ClassId),
			sdk.NewAttribute(types.AttributeKeyTokenAmountForNFT, strconv.Itoa(len(msg.Nfts))),
		),
	})

	return &types.MsgUpdateNFTsResponse{}, nil
}

// DeleteNFTs delete native nft mappings on layer2 module.
func (k Keeper) DeleteNFTs(goCtx context.Context, msg *types.MsgDeleteNFTs) (*types.MsgDeleteNFTsResponse, error) {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.DeleteTokensForNFT(ctx, msg.SpaceId, msg.ClassId, msg.NftIds, sender); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeDeleteNFTs,
			sdk.NewAttribute(types.AttributeKeySpaceId, strconv.FormatUint(msg.SpaceId, 10)),
			sdk.NewAttribute(types.AttributeKeyClassIdForNFT, msg.ClassId),
			sdk.NewAttribute(types.AttributeKeyTokenAmountForNFT, strconv.Itoa(len(msg.NftIds))),
		),
	})

	return &types.MsgDeleteNFTsResponse{}, nil
}

// DepositClassForNFT deposit a class from layer1 to layer2
func (k Keeper) DepositClassForNFT(goCtx context.Context, msg *types.MsgDepositClassForNFT) (*types.MsgDepositClassForNFTResponse, error) {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.depositClassForNFT(ctx, msg.ClassId, msg.BaseUri, sender); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeDepositClassForNFT,
			sdk.NewAttribute(types.AttributeKeyClassIdForNFT, msg.ClassId),
			sdk.NewAttribute(types.AttributeKeySender, msg.Sender),
		),
	})

	return &types.MsgDepositClassForNFTResponse{}, nil
}

// WithdrawClassForNFT withdraw a class for nft from layer2 to layer1
func (k Keeper) WithdrawClassForNFT(goCtx context.Context, msg *types.MsgWithdrawClassForNFT) (*types.MsgWithdrawClassForNFTResponse, error) {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.withdrawClassForNFT(ctx, msg.ClassId, sender, owner); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeWithdrawClassForNFT,
			sdk.NewAttribute(types.AttributeKeyClassIdForNFT, msg.ClassId),
			sdk.NewAttribute(types.AttributeKeySender, msg.Owner),
		),
	})

	return &types.MsgWithdrawClassForNFTResponse{}, nil
}

// DepositTokenForNFT deposit a nft from layer2 to layer1
func (k Keeper) DepositTokenForNFT(goCtx context.Context, msg *types.MsgDepositTokenForNFT) (*types.MsgDepositTokenForNFTResponse, error) {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.depositTokenForNFT(ctx, msg.SpaceId, msg.ClassId, msg.NftId, sender); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeDepositTokenForNFT,
			sdk.NewAttribute(types.AttributeKeySpaceId, strconv.FormatUint(msg.SpaceId, 10)),
			sdk.NewAttribute(types.AttributeKeyClassIdForNFT, msg.ClassId),
			sdk.NewAttribute(types.AttributeKeyTokenIdForNFT, msg.NftId),
			sdk.NewAttribute(types.AttributeKeySender, msg.Sender),
		),
	})

	return &types.MsgDepositTokenForNFTResponse{}, nil
}

func (k Keeper) WithdrawTokenForNFT(goCtx context.Context, msg *types.MsgWithdrawTokenForNFT) (*types.MsgWithdrawTokenForNFTResponse, error) {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.withdrawTokenForNFT(ctx, msg.SpaceId, msg.ClassId, msg.NftId, msg.Name, msg.Uri, msg.UriHash, msg.Data, sender, owner); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeWithdrawTokenForNFT,
			sdk.NewAttribute(types.AttributeKeySpaceId, strconv.FormatUint(msg.SpaceId, 10)),
			sdk.NewAttribute(types.AttributeKeyClassIdForNFT, msg.ClassId),
			sdk.NewAttribute(types.AttributeKeyTokenIdForNFT, msg.NftId),
			sdk.NewAttribute(types.AttributeKeyRecipient, msg.Owner),
		),
	})

	return &types.MsgWithdrawTokenForNFTResponse{}, nil
}
