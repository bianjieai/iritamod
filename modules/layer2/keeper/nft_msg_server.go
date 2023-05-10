package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"

	"github.com/bianjieai/iritamod/modules/layer2/types"
	"github.com/bianjieai/iritamod/modules/perm"
)

var _ types.MsgServer = Keeper{}

// CreateNFTs create native nft mappings on layer2 module.
func (k Keeper) CreateNFTs(goCtx context.Context, msg *types.MsgCreateNFTs) (*types.MsgCreateNFTsResponse, error) {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// sender must have l2 user role
	if err := k.perm.Access(ctx, sender, perm.RoleLayer2User.Auth()); err != nil {
		return nil, err
	}

	// the space must exist, and the sender must own this space
	spaceId, err := strconv.Atoi(msg.SpaceId)
	if err != nil {
		return nil, err
	}
	owner, err := k.OwnerOfSpace(int64(spaceId))
	if owner == nil || !owner.Equals(sender) {
		return nil, err
	}

	// l2 module must have this native nft class mapping
	if exist := k.HasClassForNFT(ctx, msg.ClassId); !exist {
		return nil, types.ErrClassForNFTNotExist
	}

	for _, nft := range msg.Nfts {
		nftOwner, err := sdk.AccAddressFromBech32(nft.Owner)
		if err != nil {
			return nil, err
		}

		if err := k.CreateTokenForNFT(ctx,
			int64(spaceId),
			msg.ClassId,
			nft.Id,
			nftOwner); err != nil {
			return nil, err
		}
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateNFTs,
			sdk.NewAttribute(types.AttributeKeySpaceId, msg.SpaceId),
			sdk.NewAttribute(types.AttributeKeyClassIdForNFT, msg.ClassId),
			sdk.NewAttribute(types.AttributeKeyTokenAmountForNFT, strconv.Itoa(len(msg.Nfts))),
		),
	})

	return &types.MsgCreateNFTsResponse{}, nil
}

func (k Keeper) UpdateNFTs(goCtx context.Context, msg *types.MsgUpdateNFTs) (*types.MsgUpdateNFTsResponse, error) {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// sender must have l2 user role
	if err := k.perm.Access(ctx, sender, perm.RoleLayer2User.Auth()); err != nil {
		return nil, err
	}

	// the space must exist, and the sender must own this space
	spaceId, err := strconv.Atoi(msg.SpaceId)
	if err != nil {
		return nil, err
	}
	owner, err := k.OwnerOfSpace(int64(spaceId))
	if owner == nil || !owner.Equals(sender) {
		return nil, err
	}

	// l2 module must have this native nft class mapping
	if exist := k.HasClassForNFT(ctx, msg.ClassId); !exist {
		return nil, types.ErrClassForNFTNotExist
	}

	for _, nft := range msg.Nfts {
		nftOwner, err := sdk.AccAddressFromBech32(nft.Owner)
		if err != nil {
			return nil, err
		}

		if err := k.UpdateTokenForNFT(ctx,
			int64(spaceId),
			msg.ClassId,
			nft.Id,
			nftOwner); err != nil {
			return nil, err
		}
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeUpdateNFTs,
			sdk.NewAttribute(types.AttributeKeySpaceId, msg.SpaceId),
			sdk.NewAttribute(types.AttributeKeyClassIdForNFT, msg.ClassId),
			sdk.NewAttribute(types.AttributeKeyTokenAmountForNFT, strconv.Itoa(len(msg.Nfts))),
		),
	})

	return &types.MsgUpdateNFTsResponse{}, nil
}

func (k Keeper) DeleteNFTs(ctx context.Context, msg *types.MsgDeleteNFTs) (*types.MsgDeleteNFTsResponse, error) {
	panic("implement me")
}

func (k Keeper) DepositClassForNFT(ctx context.Context, nft *types.MsgDepositClassForNFT) (*types.MsgDepositClassForNFTResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k Keeper) WithdrawClassForNFT(ctx context.Context, nft *types.MsgWithdrawClassForNFT) (*types.MsgWithdrawClassForNFTResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k Keeper) DepositTokenForNFT(ctx context.Context, nft *types.MsgDepositTokenForNFT) (*types.MsgDepositTokenForNFTResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k Keeper) WithdrawTokenForNFT(ctx context.Context, nft *types.MsgWithdrawTokenForNFT) (*types.MsgWithdrawTokenForNFTResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k Keeper) CreateL2Space(ctx context.Context, space *types.MsgCreateL2Space) (*types.MsgCreateL2SpaceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k Keeper) TransferL2Space(ctx context.Context, space *types.MsgTransferL2Space) (*types.MsgTransferL2SpaceResponse, error) {
	//TODO implement me
	panic("implement me")
}


func (k Keeper) CreateL2Record(ctx context.Context, record *types.MsgCreateL2Record) (*types.MsgCreateL2RecordResponse, error) {
	//TODO implement me
	panic("implement me")
}
