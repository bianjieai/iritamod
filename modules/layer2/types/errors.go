package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrInvalidSpaceId    = sdkerrors.Register(ModuleName, 2, "invalid space id")
	ErrInvalidSpaceOwner = sdkerrors.Register(ModuleName, 3, "invalid space owner")
	ErrBlockHeader       = sdkerrors.Register(ModuleName, 4, "block header error")
	ErrInvalidL2User     = sdkerrors.Register(ModuleName, 5, "invalid l2 user")
	ErrInvalidClassId    = sdkerrors.Register(ModuleName, 6, "invalid class id")
	ErrInvalidTokenId    = sdkerrors.Register(ModuleName, 7, "invalid token id")
	ErrInvalidClassOwner = sdkerrors.Register(ModuleName, 8, "invalid class owner")
	ErrInvalidTokenOwner = sdkerrors.Register(ModuleName, 9, "invalid token owner")
	ErrInvalidClassSpace = sdkerrors.Register(ModuleName, 10, "invalid class space")
	ErrFromNftModule     = sdkerrors.Register(ModuleName, 11, "nft module internal error")
)
