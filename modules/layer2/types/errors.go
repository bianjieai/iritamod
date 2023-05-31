package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrInvalidSpace    = sdkerrors.Register(ModuleName, 2, "invalid space id")
	ErrNotOwnerOfSpace = sdkerrors.Register(ModuleName, 3, "not space owner")
	ErrBlockHeader     = sdkerrors.Register(ModuleName, 4, "block header error")
	ErrInvalidL2User   = sdkerrors.Register(ModuleName, 5, "invalid l2 user")
	ErrInvalidClassId  = sdkerrors.Register(ModuleName, 6, "invalid class id")
	ErrInvalidTokenId  = sdkerrors.Register(ModuleName, 7, "invalid token id")
	ErrNotClassOwner   = sdkerrors.Register(ModuleName, 8, "not class owner")
	ErrNotTokenOwner   = sdkerrors.Register(ModuleName, 9, "not token owner")
	ErrNotSpaceOfClass = sdkerrors.Register(ModuleName, 10, "the class is not on the space")
	ErrFromNftModule   = sdkerrors.Register(ModuleName, 11, "error from nft module")
)
