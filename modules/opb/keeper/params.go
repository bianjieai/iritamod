package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/bianjieai/iritamod/modules/opb/types"
)

// ParamKeyTable for the OPB module
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&types.Params{})
}

// BaseTokenDenom returns the base token denom
func (k Keeper) BaseTokenDenom(ctx sdk.Context) (res string) {
	return k.GetParams(ctx).BaseTokenDenom
}

// PointTokenDenom returns the point token denom
func (k Keeper) PointTokenDenom(ctx sdk.Context) (res string) {
	return k.GetParams(ctx).PointTokenDenom
}

// BaseTokenManager returns the base token manager
func (k Keeper) BaseTokenManager(ctx sdk.Context) (res string) {
	return k.GetParams(ctx).BaseTokenManager
}

// UnrestrictedTokenTransfer returns the boolean value which indicates if the token transfer is restricted
func (k Keeper) UnrestrictedTokenTransfer(ctx sdk.Context) (res bool) {
	return k.GetParams(ctx).UnrestrictedTokenTransfer
}

// GetParams gets all parameters
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.ParamsKey)
	if bz == nil {
		return
	}

	k.cdc.MustUnmarshal(bz, &params)
	return params
}

// SetParams sets the params to the store
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
