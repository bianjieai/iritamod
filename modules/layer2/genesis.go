package nftl2

import (
	"github.com/bianjieai/iritamod/modules/nftl2/keeper"
	"github.com/bianjieai/iritamod/modules/nftl2/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func InitGenesis(ctx sdk.Context, keeper keeper.Keeper, data *types.GenesisState) {
	panic("implement me")
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper keeper.Keeper) *types.GenesisState {
	panic("implement me")
}
