package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrInvalidAmount = sdkerrors.Register(ModuleName, 2, "invalid amount")
)
