package keeper

import (
	"context"
	"github.com/bianjieai/iritamod/modules/layer2/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) ClassForNFT(ctx context.Context, request *types.QueryClassForNFTRequest) (*types.QueryClassForNFTResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k Keeper) ClassesForNFT(ctx context.Context, request *types.QueryClassesForNFTRequest) (*types.QueryClassesForNFTResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k Keeper) TokenForNFT(ctx context.Context, request *types.QueryTokenForNFTRequest) (*types.QueryTokenForNFTResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k Keeper) CollectionForNFT(ctx context.Context, request *types.QueryCollectionForNFTRequest) (*types.QueryCollectionForNFTResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k Keeper) TokensOfOwnerForNFT(ctx context.Context, request *types.QueryTokensOfOwnerForNFTRequest) (*types.QueryTokensOfOwnerForNFTResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k Keeper) BaseUriForNFT(ctx context.Context, request *types.QueryBaseUriForNFTRequest) (*types.QueryBaseUriForNFTResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k Keeper) TokenUriForNFT(ctx context.Context, request *types.QueryTokenUriForNFTRequest) (*types.QueryTokenUriForNFTResponse, error) {
	//TODO implement me
	panic("implement me")
}
