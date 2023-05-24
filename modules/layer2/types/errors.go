package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrInvalidSpace    = sdkerrors.Register(ModuleName, 2, "invalid space id")
	ErrSpaceNotExist   = sdkerrors.Register(ModuleName, 3, "space not exist")
	ErrNotOwnerOfSpace = sdkerrors.Register(ModuleName, 4, "not space owner")
	ErrBlockHeader     = sdkerrors.Register(ModuleName, 5, "block header error")
	ErrNotL2UserRole   = sdkerrors.Register(ModuleName, 6, "not l2 user")
	ErrInvalidClassId  = sdkerrors.Register(ModuleName, 7, "invalid class id")
	ErrInvalidTokenId  = sdkerrors.Register(ModuleName, 8, "invalid token id")
	ErrClassNotExist   = sdkerrors.Register(ModuleName, 9, "class not exist")
	ErrTokenNotExist   = sdkerrors.Register(ModuleName, 10, "token not exist")
	ErrNotClassOwner   = sdkerrors.Register(ModuleName, 11, "not class owner")
	ErrNotTokenOwner   = sdkerrors.Register(ModuleName, 12, "not token owner")
	ErrNotSpaceOfClass = sdkerrors.Register(ModuleName, 13, "the class is not on the space")
	ErrFromNftModule   = sdkerrors.Register(ModuleName, 14, "error from nft module")
)
