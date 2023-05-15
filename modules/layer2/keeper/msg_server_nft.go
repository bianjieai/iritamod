package keeper

import (
	"context"
	"github.com/bianjieai/iritamod/modules/layer2/expected_keeper"
	"github.com/bianjieai/iritamod/modules/layer2/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

	// burn these nft locked by module account if they exist
	moduleAddrStr := types.ModuleAddress.String()
	moduleAddr, err := sdk.AccAddressFromBech32(moduleAddrStr)
	if err != nil {
		return nil, err
	}
	for _, nftId := range msg.NftIds {
		_, err := k.nft.GetNFT(ctx, msg.ClassId, nftId)
		if err != nil {
			continue
		}
		if err := k.nft.RemoveNFT(ctx, msg.ClassId, nftId, moduleAddr); err != nil {
			return nil, err
		}
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

	// check if the denom exists
	denomInfo, err := k.nft.GetDenomInfo(ctx, msg.ClassId)
	if err != nil {
		return nil, err
	}
	// check if the denom owned by sender
	if denomInfo.Creator != msg.Sender {
		return nil, sdkerrors.Wrapf(types.ErrClassNotOwnedByAccount, "class %s is not owned by %s", msg.ClassId, msg.Sender)
	}

	if k.HasClassForNFT(ctx, msg.ClassId) {
		if err := k.UpdateClassForNFT(ctx, msg.ClassId, msg.BaseUri, msg.Sender); err != nil {
			return nil, err
		}
	} else {
		if err := k.CreateClassForNFT(ctx, msg.ClassId, msg.BaseUri, msg.Sender, denomInfo.MintRestricted); err != nil {
			return nil, err
		}
	}

	moduleAddrStr := types.ModuleAddress.String()
	moduleAddr, err := sdk.AccAddressFromBech32(moduleAddrStr)
	if err != nil {
		return nil, err
	}

	if err := k.nft.TransferDenomOwner(ctx, msg.ClassId, sender, moduleAddr); err != nil {
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

	// sender must have l2 user
	ok, err := k.HasL2UserRole(ctx, sender)
	if !ok {
		return nil, err
	}

	// check if the class mapping exist
	classForNFT, exist := k.GetClassForNFT(ctx, msg.ClassId)
	if !exist {
		return nil, sdkerrors.Wrapf(types.ErrClassForNFTNotExist, "class mapping %s not exist", msg.ClassId)
	}

	// check if the class exist
	denom, err := k.nft.GetDenomInfo(ctx, msg.ClassId)
	if err != nil {
		return nil, err
	}

	// check if the class owned by module account
	if denom.Creator != types.ModuleAddress.String() {
		return nil, sdkerrors.Wrapf(types.ErrClassNotOwnedByAccount, "class %s is not locked by %s", msg.ClassId, types.ModuleAddress.String())
	}

	// check if the class mapping owner is msg.owner
	if classForNFT.Owner != msg.Owner {
		return nil, sdkerrors.Wrapf(types.ErrClassNotOwnedByAccount, "original owner want %s, got %s", classForNFT.Owner, msg.Sender)
	}

	// recover mint_restricted and transfer ownership
	moduleAddrStr := types.ModuleAddress.String()
	moduleAddr, err := sdk.AccAddressFromBech32(moduleAddrStr)
	if err != nil {
		return nil, err
	}

	// recover mint_restricted
	denomMetadata := &expected_keeper.DenomMetadata{
		Creator:          denom.Creator,
		Schema:           denom.Schema,
		MintRestricted:   denom.MintRestricted,
		UpdateRestricted: classForNFT.Layer1MintRestricted,
		Data:             denom.Data,
	}
	data, err := codectypes.NewAnyWithValue(denomMetadata)
	if err != nil {
		return nil, err
	}
	class := expected_keeper.Class{
		Id:     denom.Id,
		Name:   denom.Name,
		Symbol: denom.Symbol,
		Data:   data,

		Description: denom.Description,
		Uri:         denom.Uri,
		UriHash:     denom.UriHash,
	}

	if err := k.nft.UpdateClass(ctx, class); err != nil {
		return nil, err
	}

	if err := k.nft.TransferDenomOwner(ctx, msg.ClassId, moduleAddr, owner); err != nil {
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

	// token must exist
	nft, err := k.nft.GetNFT(ctx, msg.ClassId, msg.NftId)
	if err != nil {
		return nil, err
	}

	if !nft.GetOwner().Equals(sender) {
		return nil, sdkerrors.Wrapf(types.ErrTokenForNFTNotOwnedByAccount, "nft %s is not owned by %s", msg.NftId, msg.Sender)
	}

	if !k.HasSpace(ctx, msg.SpaceId) {
		return nil, sdkerrors.Wrapf(types.ErrUnknownSpace, "space %d not exist", msg.SpaceId)
	}

	if err := k.createTokenForNFT(ctx, msg.SpaceId, msg.ClassId, msg.NftId, sender); err != nil {
		return nil, err
	}

	moduleAddrStr := types.ModuleAddress.String()
	moduleAddr, err := sdk.AccAddressFromBech32(moduleAddrStr)
	if err != nil {
		return nil, err
	}
	if err := k.nft.Transfer(ctx, msg.ClassId, msg.NftId, moduleAddr); err == nil {
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

	if ok, err := k.HasL2UserRole(ctx, sender); !ok {
		return nil, err
	}

	if !k.HasSpaceByOwner(ctx, sender, msg.SpaceId) {
		return nil, sdkerrors.Wrapf(types.ErrNotOwnerOfSpace, "space %d not owned by %s", msg.SpaceId, msg.Sender)
	}

	tokenOwner, exist := k.GetTokenOwnerForNFT(ctx, msg.SpaceId, msg.ClassId, msg.NftId)
	if !exist {
		return nil, sdkerrors.Wrapf(types.ErrTokenForNFTNotExist, "token %s not exist", msg.NftId)
	}
	if !tokenOwner.Equals(owner) {
		return nil, sdkerrors.Wrapf(types.ErrTokenForNFTNotOwnedByAccount, "nft %s is not owned by %s", msg.NftId, msg.Owner)
	}

	_, err = k.nft.GetNFT(ctx, msg.ClassId, msg.NftId)
	if err != nil {
		// no such nft, mint it
		if err := k.nft.SaveNFT(ctx, msg.ClassId, msg.NftId, msg.Name, msg.Uri, msg.UriHash, msg.Data, owner); err != nil {
			return nil, err
		}
	} else {
		moduleAddrStr := types.ModuleAddress.String()
		moduleAddr, err := sdk.AccAddressFromBech32(moduleAddrStr)
		if err != nil {
			return nil, err
		}

		// nft exist, update and transfer ownership
		if err := k.nft.TransferOwnership(ctx, msg.ClassId, msg.NftId, msg.Name, msg.Uri, msg.UriHash, msg.Data, moduleAddr, owner); err != nil {
			return nil, err
		}
	}

	k.deleteTokenOwnerForNFT(ctx, msg.SpaceId, msg.ClassId, msg.NftId, owner)
	k.deleteTokenForNFT(ctx, msg.SpaceId, msg.ClassId, msg.NftId)

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
