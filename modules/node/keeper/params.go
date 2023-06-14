package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/bianjieai/iritamod/modules/node/types"
)

// Default parameter namespace
const (
	DefaultParamspace = types.ModuleName
)

// HistoricalEntries = number of historical info entries
// to persist in store
func (k Keeper) HistoricalEntries(ctx sdk.Context) (res uint32) {
	return k.GetModuleParams(ctx).HistoricalEntries
}

// GetModuleParams Get all parameteras as types.Params
func (k Keeper) GetModuleParams(ctx sdk.Context) (params types.Params) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.ParamsKey)
	if bz == nil {
		return params
	}

	k.cdc.MustUnmarshal(bz, &params)
	return params
}

// NOTE: as evidence module expect staking keeper to implement GetParams method
// which returns x/staking stakingtypes.Params, we convert node Params to
// staking Params and name the real node GetParmas as GetModuleParams.
func (k Keeper) GetParams(ctx sdk.Context) (params stakingtypes.Params) {
	param := k.GetModuleParams(ctx)
	return stakingtypes.Params{HistoricalEntries: param.HistoricalEntries}
}

// SetParams sets the node module parameters
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) error {
	if err := params.Validate(); err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	bz, err := k.cdc.Marshal(&params)
	if err != nil {
		return err
	}
	store.Set(types.ParamsKey, bz)

	return nil
}

// UnbondingTime returns the unbonding time
func (k Keeper) UnbondingTime(ctx sdk.Context) (res time.Duration) {
	return
}
