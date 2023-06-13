package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/iritamod/modules/node/exported"
	"github.com/bianjieai/iritamod/modules/node/types"
)

// NodeKeeper defines am interface for SetParams function
type NodeKeeper interface {
	SetParams(ctx sdk.Context, params types.Params) error
}

// Migrate migrates the node params from legacy x/params module to node module
func Migrate(ctx sdk.Context, k NodeKeeper, legacySubspace exported.Subspace) error {
	var params types.Params
	legacySubspace.GetParamSet(ctx, &params)
	return k.SetParams(ctx, params)
}
