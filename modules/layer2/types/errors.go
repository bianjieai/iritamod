package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrInvalidSpace = sdkerrors.Register(ModuleName, 2, "invalid space id")
	ErrUnknownSpace = sdkerrors.Register(ModuleName, 3, "unknown space id")
	ErrNotOwnerOfSpace = sdkerrors.Register(ModuleName, 4, "the space is not owned by this address")
	ErrInvalidRecord = sdkerrors.Register(ModuleName, 5, "invalid record")

	ErrInvalidClassIdForNFT = sdkerrors.Register(ModuleName, 11, "invalid class id for nft mapping")
	ErrInvalidTokenIdForNFT = sdkerrors.Register(ModuleName, 12, "invalid class id for nft mapping")
	ErrClassForNFTNotExist = sdkerrors.Register(ModuleName, 13, "the class for nft mapping does not exist")
	ErrTokenForNFTAlreadyExist = sdkerrors.Register(ModuleName, 14, "the token for nft mapping already exist")
	ErrDuplicateTokenIdsForNFT = sdkerrors.Register(ModuleName, 15, "duplicate token ids for nft mapping")
)