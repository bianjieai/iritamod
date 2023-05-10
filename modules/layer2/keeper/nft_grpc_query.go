package keeper

import (
	"context"
	"github.com/bianjieai/iritamod/modules/layer2/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Class(ctx context.Context, request *types.QueryClassRequest) (*types.QueryClassResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k Keeper) Classes(ctx context.Context, request *types.QueryClassesRequest) (*types.QueryClassesResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k Keeper) NFT(ctx context.Context, request *types.QueryNFTRequest) (*types.QueryNFTResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k Keeper) CollectionOfSpace(ctx context.Context, request *types.QueryCollectionOfSpaceRequest) (*types.QueryCollectionOfSpaceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k Keeper) NFTsOfOwner(ctx context.Context, request *types.QueryNFTsOfOwnerRequest) (*types.QueryNFTsOfOwnerResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k Keeper) BaseUri(ctx context.Context, request *types.QueryBaseUriRequest) (*types.QueryBaseUriResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k Keeper) TokenUri(ctx context.Context, request *types.QueryTokenUriRequest) (*types.QueryTokenUriResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k Keeper) Space(ctx context.Context, request *types.QuerySpaceRequest) (*types.QuerySpaceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k Keeper) SpaceOfOwner(ctx context.Context, request *types.QuerySpaceOfOwnerRequest) (*types.QuerySpaceOfOwnerResponse, error) {
	//TODO implement me
	panic("implement me")
}