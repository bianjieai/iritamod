package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/bianjieai/iritamod/modules/node/types"
)

// Default parameter namespace
const (
	DefaultParamspace = types.ModuleName
)

// ParamTable for node module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&types.Params{})
}

// HistoricalEntries = number of historical info entries
// to persist in store
func (k Keeper) HistoricalEntries(ctx sdk.Context) (res uint32) {
	k.paramstore.Get(ctx, types.KeyHistoricalEntries, &res)
	return
}

// GetParamsLegacy Get all parameteras as types.Params
func (k Keeper) GetParamsLegacy(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.HistoricalEntries(ctx),
	)
}

// NOTE: implement expected staking keeper for evidence module
func (k Keeper) GetParams(ctx sdk.Context) (params stakingtypes.Params) {
	param := k.GetParamsLegacy(ctx)
	return stakingtypes.Params{HistoricalEntries: param.HistoricalEntries}
}

// set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// UnbondingTime
func (k Keeper) UnbondingTime(ctx sdk.Context) (res time.Duration) {
	return
}
