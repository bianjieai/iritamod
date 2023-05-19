package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/bianjieai/iritamod/modules/layer2/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) ClassForNFT(goCtx context.Context, req *types.QueryClassForNFTRequest) (*types.QueryClassForNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	class, err := k.GetClassForNFT(ctx, req.ClassId)
	if err != nil {
		return nil, err
	}
	return &types.QueryClassForNFTResponse{
		Class: &class,
	}, nil
}

func (k Keeper) ClassesForNFT(goCtx context.Context, req *types.QueryClassesForNFTRequest) (*types.QueryClassesForNFTResponse, error) {
	var classes []*types.ClassForNFT
	ctx := sdk.UnwrapSDKContext(goCtx)

	pageResp, err := query.Paginate(k.getClassStore(ctx), req.Pagination, func(key []byte, value []byte) error {
		class, err := k.GetClassForNFT(ctx, string(key))
		if err != nil {
			return err
		}

		classes = append(classes, &class)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &types.QueryClassesForNFTResponse{
		Classes:    classes,
		Pagination: pageResp,
	}, nil
}

func (k Keeper) TokenForNFT(goCtx context.Context, req *types.QueryTokenForNFTRequest) (*types.QueryTokenForNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	owner, err := k.GetTokenForNFT(ctx, req.SpaceId, req.ClassId, req.NftId)
	if err != nil {
		return nil, err
	}
	return &types.QueryTokenForNFTResponse{
		Owner: owner.String(),
	}, nil
}

func (k Keeper) CollectionForNFT(goCtx context.Context, req *types.QueryCollectionForNFTRequest) (*types.QueryCollectionForNFTResponse, error) {
	var nfts []*types.TokenForNFT
	ctx := sdk.UnwrapSDKContext(goCtx)

	pageResp, err := query.Paginate(k.getCollectionStore(ctx, req.SpaceId, req.ClassId), req.Pagination, func(key []byte, value []byte) error {
		tokenId := string(key)
		owner := sdk.AccAddress(value)
		nft := types.TokenForNFT{
			Id:    tokenId,
			Owner: owner.String(),
		}
		nfts = append(nfts, &nft)

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &types.QueryCollectionForNFTResponse{
		ClassId:    req.ClassId,
		Nfts:       nfts,
		Pagination: pageResp,
	}, nil
}

func (k Keeper) TokensOfOwnerForNFT(goCtx context.Context, req *types.QueryTokensOfOwnerForNFTRequest) (*types.QueryTokensOfOwnerForNFTResponse, error) {
	owner, err := sdk.AccAddressFromBech32(req.Owner)
	if err != nil {
		return nil, err
	}

	var nfts []*types.TokenForNFTByOwner
	var pageRes *query.PageResponse
	ctx := sdk.UnwrapSDKContext(goCtx)

	switch {
	case req.SpaceId == 0 && len(req.ClassId) == 0:
		if pageRes, err = query.Paginate(k.getNFTsOfOwnerStore(ctx, owner), req.Pagination, func(key []byte, _ []byte) error {
			spaceId, classId, tokenId := types.ParseNFTsOfOwnerAllStoreKey(key)

			nft := types.TokenForNFTByOwner{
				SpaceId: spaceId,
				ClassId: classId,
				TokenId: tokenId,
			}
			nfts = append(nfts, &nft)
			return nil
		}); err != nil {
			return nil, err
		}
	case req.SpaceId > 0 && len(req.ClassId) == 0:
		if pageRes, err = query.Paginate(k.getNFTsOfOwnerBySpaceStore(ctx, owner, req.SpaceId), req.Pagination, func(key []byte, _ []byte) error {
			classId, tokenId := types.ParseNFTsOfOwnerBySpaceStoreKey(key)

			nft := types.TokenForNFTByOwner{
				SpaceId: req.SpaceId,
				ClassId: classId,
				TokenId: tokenId,
			}
			nfts = append(nfts, &nft)
			return nil
		}); err != nil {
			return nil, err
		}
	case req.SpaceId > 0 && len(req.ClassId) > 0:
		if pageRes, err = query.Paginate(k.getNFTsOfOwnerBySpaceAndClassStore(ctx, owner, req.SpaceId, req.ClassId), req.Pagination, func(key []byte, _ []byte) error {
			nft := types.TokenForNFTByOwner{
				SpaceId: req.SpaceId,
				ClassId: req.ClassId,
				TokenId: string(key),
			}
			nfts = append(nfts, &nft)
			return nil
		}); err != nil {
			return nil, err
		}
	}

	return &types.QueryTokensOfOwnerForNFTResponse{
		Tokens:     nfts,
		Pagination: pageRes,
	}, nil
}

func (k Keeper) BaseUriForNFT(goCtx context.Context, req *types.QueryBaseUriForNFTRequest) (*types.QueryBaseUriForNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	class, err := k.GetClassForNFT(ctx, req.ClassId)
	if err != nil {
		return nil, err
	}

	return &types.QueryBaseUriForNFTResponse{
		BaseUri: class.BaseUri,
	}, nil
}

func (k Keeper) TokenUriForNFT(goCtx context.Context, req *types.QueryTokenUriForNFTRequest) (*types.QueryTokenUriForNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	class, err := k.GetClassForNFT(ctx, req.ClassId)
	if err != nil {
		return nil, err
	}

	_, err = k.GetTokenForNFT(ctx, req.SpaceId, req.ClassId, req.TokenId)
	if err != nil {
		return nil, err
	}

	return &types.QueryTokenUriForNFTResponse{
		TokenUri: class.BaseUri + "/" + req.TokenId,
	}, nil
}
