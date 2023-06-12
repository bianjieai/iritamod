package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"

	"github.com/bianjieai/iritamod/modules/slashing/exported"
)

// SlashingKeeper defines an interface for setting module params
type SlashingKeeper interface {
	SetParams(ctx sdk.Context, params slashingtypes.Params) error
}

// Migrate migrates the coinswap params from legacy x/params module to slashing module
func Migrate(ctx sdk.Context, k SlashingKeeper, legacySubspace exported.Subspace) error {
	var params slashingtypes.Params
	legacySubspace.GetParamSet(ctx, &params)

	if err := params.Validate(); err != nil {
		return err
	}

	return k.SetParams(ctx, params)
}
