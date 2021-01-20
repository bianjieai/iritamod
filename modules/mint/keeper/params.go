package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/bianjieai/iritamod/modules/mint/types"
)

// ParamKeyTable for the mint module
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&types.Params{})
}

// MintDenom returns the minting denom
func (k Keeper) MintDenom(ctx sdk.Context) (res string) {
	k.paramSpace.Get(ctx, types.KeyMintDenom, &res)
	return
}

// GetParams gets all parameters
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.MintDenom(ctx),
	)
}

// SetParams sets the params to the store
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}
