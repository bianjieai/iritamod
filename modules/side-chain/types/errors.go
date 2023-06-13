package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrInvalidSpaceId       = sdkerrors.Register(ModuleName, 2, "invalid space id")
	ErrInvalidSpaceOwner    = sdkerrors.Register(ModuleName, 3, "invalid space owner")
	ErrBlockHeader          = sdkerrors.Register(ModuleName, 4, "block header error")
	ErrInvalidSideChainUser = sdkerrors.Register(ModuleName, 5, "invalid side chain user")
)
