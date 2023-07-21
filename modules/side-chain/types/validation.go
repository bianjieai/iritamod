package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func ValidateSpaceId(spaceId uint64) error {
	if spaceId == 0 {
		return sdkerrors.Wrapf(ErrInvalidSpaceId, "space id cannot be zero")
	}
	return nil
}
