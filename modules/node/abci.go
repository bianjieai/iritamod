package node

import (
	abci "github.com/cometbft/cometbft/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"iritamod.bianjie.ai/modules/node/keeper"
)

// BeginBlocker will persist the current header and validator set as a historical entry
// and prune the oldest entry based on the HistoricalEntries parameter
func BeginBlocker(ctx sdk.Context, k *keeper.Keeper) {
	k.TrackHistoricalInfo(ctx)
}

// Called every block, update validator set
func EndBlocker(ctx sdk.Context, k *keeper.Keeper) (updates []abci.ValidatorUpdate) {
	updates, _ = k.ApplyAndReturnValidatorSetUpdates(ctx)
	return updates
}
