package mock

import sdk "github.com/cosmos/cosmos-sdk/types"

type PermKeeper struct {
}

func NewPermKeeper() *PermKeeper {
	return &PermKeeper{}
}

func (pk *PermKeeper) HasL2UserRole(ctx sdk.Context, signer sdk.AccAddress) bool {
	return true
}
