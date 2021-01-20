package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// TokenKeeper defines the expected token keeper (noalias)
type TokenKeeper interface {
	MintToken(ctx sdk.Context, denom string, amount uint64, recipient sdk.AccAddress, owner sdk.AccAddress) error
}
