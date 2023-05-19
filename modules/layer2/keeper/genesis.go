package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/iritamod/modules/layer2/types"
)

func (k Keeper) InitGenesis(ctx sdk.Context, data types.GenesisState) {
	if err := types.ValidateGenesis(data); err != nil {
		panic(err.Error())
	}

	k.setSpaceSequence(ctx, data.SpaceSequence)
	for _, space := range data.Spaces {
		owner, _ := sdk.AccAddressFromBech32(space.Owner)
		k.setSpace(ctx, space.Id, space)
		k.setSpaceOfOwner(ctx, space.Id, owner)
	}

	for _, blockHeader := range data.L2BlockHeaders {
		k.setL2BlockHeader(ctx, blockHeader.SpaceId, blockHeader.Height, blockHeader.Header)
	}

	for _, class := range data.ClassesWithSpaceForNft {
		classForNFT := types.ClassForNFT{
			Id:                   class.Id,
			Owner:                class.Owner,
			BaseUri:              class.BaseUri,
			Layer1MintRestricted: class.Layer1MintRestricted,
		}
		k.setClassForNFT(ctx, classForNFT)
		k.setSpaceOfClassForNFT(ctx, class.Id, class.ActiveSpace)
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
		SpaceSequence:          0,
		Spaces:                 make([]types.Space, 0),
		L2BlockHeaders:         make([]types.L2BlockHeader, 0),
		ClassesWithSpaceForNft: make([]types.ClassWithSpaceForNFT, 0),
		CollectionsForNft:      make([]types.CollectionForNFT, 0),
	}

	data.SpaceSequence = k.GetSpaceSequence(ctx)
	data.Spaces = k.GetSpaces(ctx)
	data.L2BlockHeaders = k.GetL2BlockHeaders(ctx)
	data.ClassesWithSpaceForNft = k.GetClassesWithSpaceForNFT(ctx)
	data.CollectionsForNft = k.GetCollectionsForNFT(ctx)
	return &data
}
