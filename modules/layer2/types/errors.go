package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

// TODO: error names need to be refactored
var (
	ErrInvalidSpace         = sdkerrors.Register(ModuleName, 2, "invalid space id")
	ErrUnknownSpace         = sdkerrors.Register(ModuleName, 3, "unknown space id")
	ErrNotOwnerOfSpace      = sdkerrors.Register(ModuleName, 4, "the space is not owned by this address")
	ErrDuplicateSpace       = sdkerrors.Register(ModuleName, 5, "duplicate space")
	ErrInvalidRecord        = sdkerrors.Register(ModuleName, 6, "invalid record")
	ErrRecordAlreadyExist   = sdkerrors.Register(ModuleName, 7, "record already exist")
	ErrDuplicateRecord      = sdkerrors.Register(ModuleName, 8, "duplicate record")
	ErrUnknownL2BlockHeader = sdkerrors.Register(ModuleName, 9, "unknown l2 block record")
	ErrNotL2UserRole = sdkerrors.Register(ModuleName, 10, "the account is not a l2 user")
	ErrInvalidClassIdForNFT         = sdkerrors.Register(ModuleName, 11, "invalid class id for nft mapping")
	ErrInvalidTokenIdForNFT         = sdkerrors.Register(ModuleName, 12, "invalid class id for nft mapping")
	ErrClassForNFTNotExist          = sdkerrors.Register(ModuleName, 13, "the class for nft mapping does not exist")
	ErrTokenForNFTNotExist          = sdkerrors.Register(ModuleName, 14, "the token for nft mapping does not exist")
	ErrTokenForNFTAlreadyExist      = sdkerrors.Register(ModuleName, 15, "the token for nft mapping already exist")
	ErrDuplicateClassIdForNFT       = sdkerrors.Register(ModuleName, 16, "duplicate class id for nft mapping")
	ErrDuplicateTokenIdsForNFT      = sdkerrors.Register(ModuleName, 17, "duplicate token ids for nft mapping")
	ErrUnknownClassIdForNFT         = sdkerrors.Register(ModuleName, 18, "unknown class id for nft mapping")
	ErrClassNotOwnedByAccount       = sdkerrors.Register(ModuleName, 19, "the class is not owned by this address")
	ErrTokenForNFTNotOwnedByAccount = sdkerrors.Register(ModuleName, 20, "the token for nft is not owned by this address")
	ErrL1ClassNotFoundForNFT        = sdkerrors.Register(ModuleName, 21, "the class for nft is not found")
	ErrL1TokenNotFoundForNFT        = sdkerrors.Register(ModuleName, 22, "the token for nft is not found")
	ErrClassNotOnSpace = sdkerrors.Register(ModuleName, 23, "the class is not on the space")
)
