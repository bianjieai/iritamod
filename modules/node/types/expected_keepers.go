package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ValidatorKeeper defines the expected validator keeper (noalias)
type ValidatorKeeper interface {
	GetRootCert(ctx sdk.Context) (string, bool)
}
