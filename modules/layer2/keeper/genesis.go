package keeper

import (
	"github.com/bianjieai/iritamod/modules/layer2/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) InitGenesis(ctx sdk.Context, data *types.GenesisState) {
	panic("implement me")
}

func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	panic("implement me")
}