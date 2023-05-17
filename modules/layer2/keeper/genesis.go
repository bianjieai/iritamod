package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/iritamod/modules/layer2/types"
)

func (k Keeper) InitGenesis(ctx sdk.Context, data types.GenesisState) {
	if err := types.ValidateGenesis(data); err != nil {
		panic(err.Error())
	}

	k.setSpaceId(ctx, data.StartingSpaceId)
	for _, space := range data.Spaces {
		owner, err := sdk.AccAddressFromBech32(space.Owner)
		if err != nil {
			panic(err)
		}
		k.setSpace(ctx, space.Id, space)
		k.setSpaceOfOwner(ctx, space.Id, owner)
	}

	for _, blockHeader := range data.L2BlockHeaders {
		k.setL2BlockHeader(ctx, blockHeader.SpaceId, blockHeader.Height, blockHeader.Header)
	}

	for _, class := range data.ClassesForNft {
		k.setClassForNFT(ctx, class)
	}

	for _, collection := range data.CollectionsForNft {
		for _, token := range collection.Tokens {
			owner, _ := sdk.AccAddressFromBech32(token.Owner)
			k.setTokenForNFT(ctx, collection.SpaceId, collection.ClassId, token.Id, owner)
			k.setTokenOwnerForNFT(ctx, collection.SpaceId, collection.ClassId, token.Id, owner)
		}
	}
}

func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	data := types.GenesisState{
		StartingSpaceId:   0,
		Spaces:            make([]types.Space, 0),
		L2BlockHeaders:    make([]types.L2BlockHeader, 0),
		ClassesForNft:     make([]types.ClassForNFT, 0),
		CollectionsForNft: make([]types.CollectionForNFT, 0),
	}

	data.StartingSpaceId = k.getSpaceId(ctx)
	data.Spaces = k.GetSpaces(ctx)
	data.L2BlockHeaders = k.GetL2BlockHeaders(ctx)
	data.ClassesForNft = k.GetClassesForNFT(ctx)
	data.CollectionsForNft = k.GetCollectionsForNFT(ctx)
	return &data
}
