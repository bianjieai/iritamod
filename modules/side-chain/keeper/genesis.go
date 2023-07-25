package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/iritamod/modules/side-chain/types"
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

	for _, blockHeader := range data.BlockHeaders {
		k.setBlockHeader(ctx, blockHeader.SpaceId, blockHeader.Height, blockHeader.Header)
		if blockHeader.TxHash != "" {
			k.setBlockHeaderTxHashString(ctx, blockHeader.SpaceId, blockHeader.Height, blockHeader.TxHash)
		}
	}

	for _, spaceLatestHeight := range data.SpaceLatestHeights {
		k.setSpaceLatestHeight(ctx, spaceLatestHeight.SpaceId, spaceLatestHeight.Height)
	}
}

func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	data := types.GenesisState{
		SpaceSequence:      0,
		Spaces:             make([]types.Space, 0),
		BlockHeaders:       make([]types.BlockHeader, 0),
		SpaceLatestHeights: make([]types.SpaceLatestHeight, 0),
	}

	data.SpaceSequence = k.GetSpaceSequence(ctx)
	data.Spaces = k.GetSpaces(ctx)
	data.BlockHeaders = k.GetBlockHeaders(ctx)
	data.SpaceLatestHeights = k.GetSpaceLatestHeights(ctx)
	return &data
}
