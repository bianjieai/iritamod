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

func (k Keeper) DepositClassForNFT(goCtx context.Context, msg *types.MsgDepositClassForNFT) (*types.MsgDepositClassForNFTResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k Keeper) WithdrawClassForNFT(goCtx context.Context, msg *types.MsgWithdrawClassForNFT) (*types.MsgWithdrawClassForNFTResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k Keeper) DepositTokenForNFT(goCtx context.Context, msg *types.MsgDepositTokenForNFT) (*types.MsgDepositTokenForNFTResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k Keeper) WithdrawTokenForNFT(goCtx context.Context, msg *types.MsgWithdrawTokenForNFT) (*types.MsgWithdrawTokenForNFTResponse, error) {
	//TODO implement me
	panic("implement me")
}