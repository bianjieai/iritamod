package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrUnknownSpace = sdkerrors.Register(ModuleName, 2, "the space does not exist")
	ErrNotOwnerOfSpace = sdkerrors.Register(ModuleName, 3, "the space is not owned by this address")

	ErrClassForNFTNotExist = sdkerrors.Register(ModuleName, 11, "the class for nft does not exist")
	ErrTokenForNFTAlreadyExist = sdkerrors.Register(ModuleName, 12, "the token for nft already exist")
)