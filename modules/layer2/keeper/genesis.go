package keeper

import (
	"github.com/bianjieai/iritamod/modules/layer2/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) InitGenesis(ctx sdk.Context, data types.GenesisState) {
	if err := types.ValidateGenesis(data); err != nil {
		panic(err.Error())
	}

	k.setSpaceId(ctx, data.StartingSpaceId)
	for _, space := range data.Spaces {
		owner, _ := sdk.AccAddressFromBech32(space.Owner)
		k.setSpace(ctx, space.Id, owner)
		k.setSpaceOwner(ctx, space.Id, owner)
	}

	for _, record := range data.Records {
		k.setRecord(ctx, record.SpaceId, record.Height, record.Header)
	}

	for _, class := range data.NftMappings.Classes {
		bz, err := k.cdc.Marshal(&class)
		if err != nil {
			panic(err.Error())
		}

		k.setClassForNFT(ctx, class.Id, bz)
	}

	for _, collection := range data.NftMappings.Collections {
		for _, token := range collection.Tokens {
			owner, _ := sdk.AccAddressFromBech32(token.Owner)
			k.setTokenForNFT(ctx, collection.SpaceId, collection.ClassId, token.Id, owner)
			k.setTokenOwnerForNFT(ctx, collection.SpaceId, collection.ClassId, token.Id, owner)
		}
	}
}

func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	nftMappings := types.MappingsForNFT{
		Classes:     make([]types.ClassForNFT, 0),
		Collections: make([]types.CollectionForNFT, 0),
	}

	data := types.GenesisState{
		StartingSpaceId: 0,
		Spaces:          make([]types.Space, 0),
		Records:         make([]types.Record, 0),
		NftMappings:     &nftMappings,
	}

	data.StartingSpaceId = k.getSpaceId(ctx)
	// TODO： implemtn this
	return &data
}
