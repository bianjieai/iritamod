package keeper

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/bianjieai/iritamod/modules/node/types"
)

// GetHistoricalInfo gets the historical info at a given height
func (k Keeper) GetHistoricalInfo(ctx sdk.Context, height int64) (stakingtypes.HistoricalInfo, bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetHistoricalInfoKey(height)

	value := store.Get(key)
	if value == nil {
		return stakingtypes.HistoricalInfo{}, false
	}

	var hi types.HistoricalInfo
	k.cdc.MustUnmarshal(value, &hi)

	// convert to staking validator set
	valSet := make([]stakingtypes.Validator, len(hi.Valset))

	for i, v := range hi.Valset {
		pubKey, err := v.ConsPubKey()
		if err != nil {
			panic(err)
		}

		pkAny, err := codectypes.NewAnyWithValue(pubKey)
		if err != nil {
			continue
		}
		valSet[i] = stakingtypes.Validator{
			ConsensusPubkey: pkAny,
			Tokens:          v.GetTokens(),
		}
	}

	return stakingtypes.NewHistoricalInfo(hi.Header, valSet, sdk.DefaultPowerReduction), true
}

// SetHistoricalInfo sets the historical info at a given height
func (k Keeper) SetHistoricalInfo(ctx sdk.Context, height int64, hi types.HistoricalInfo) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetHistoricalInfoKey(height)

	value := k.cdc.MustMarshal(&hi)
	store.Set(key, value)
}

// DeleteHistoricalInfo deletes the historical info at a given height
func (k Keeper) DeleteHistoricalInfo(ctx sdk.Context, height int64) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetHistoricalInfoKey(height)

	store.Delete(key)
}

// IterateHistoricalInfo provides an interator over all stored HistoricalInfo
//  objects. For each HistoricalInfo object, cb will be called. If the cb returns
// true, the iterator will close and stop.
func (k Keeper) IterateHistoricalInfo(ctx sdk.Context, cb func(types.HistoricalInfo) bool) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.HistoricalInfoKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var histInfo types.HistoricalInfo
		k.cdc.MustUnmarshal(iterator.Value(), &histInfo)
		if cb(histInfo) {
			break
		}
	}
}

// GetAllHistoricalInfo returns all stored HistoricalInfo objects.
func (k Keeper) GetAllHistoricalInfo(ctx sdk.Context) []types.HistoricalInfo {
	var infos []types.HistoricalInfo

	k.IterateHistoricalInfo(ctx, func(histInfo types.HistoricalInfo) bool {
		infos = append(infos, histInfo)
		return false
	})

	return infos
}

// TrackHistoricalInfo saves the latest historical-info and deletes the oldest
// heights that are below pruning height
func (k Keeper) TrackHistoricalInfo(ctx sdk.Context) {
	entryNum := k.HistoricalEntries(ctx)

	// Prune store to ensure we only have parameter-defined historical entries.
	// In most cases, this will involve removing a single historical entry.
	// In the rare scenario when the historical entries gets reduced to a lower value k'
	// from the original value k. k - k' entries must be deleted from the store.
	// Since the entries to be deleted are always in a continuous range, we can iterate
	// over the historical entries starting from the most recent version to be pruned
	// and then return at the first empty entry.
	for i := ctx.BlockHeight() - int64(entryNum); i >= 0; i-- {
		if _, found := k.GetHistoricalInfo(ctx, i); found {
			k.DeleteHistoricalInfo(ctx, i)
		} else {
			break
		}
	}

	// if there is no need to persist historicalInfo, return
	if entryNum == 0 {
		return
	}

	// CreateValidator HistoricalInfo struct
	lastVals := k.GetLastValidators(ctx)
	historicalEntry := types.NewHistoricalInfo(ctx.BlockHeader(), lastVals)

	// Set latest HistoricalInfo at current height
	k.SetHistoricalInfo(ctx, ctx.BlockHeight(), historicalEntry)
}
