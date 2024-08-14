package types

import (
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	NewParamSetPair = paramtypes.NewParamSetPair
	NewKeyTable     = paramtypes.NewKeyTable
)

type (
	ParamSet      = paramtypes.ParamSet
	ParamSetPairs = paramtypes.ParamSetPairs
	KeyTable      = paramtypes.KeyTable

	// Subspace defines an interface that implements the legacy x/params Subspace
	// type.
	//
	// NOTE: This is used solely for migration of x/params managed parameters.
	Subspace interface {
		paramtypes.Subspace
	}
)
