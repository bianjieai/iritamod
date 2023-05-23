package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

type NFT interface {
	GetOwner() sdk.AccAddress
}

type Class interface {
	GetID() string
	GetOwner() string
	GetMintRestricted() bool
}

type NFTKeeper interface {
	SaveNFT(ctx sdk.Context, classID, tokenID, tokenNm, tokenURI, tokenUriHash, tokenData string, receiver sdk.AccAddress) error
	UpdateNFT(ctx sdk.Context, classID, tokenID, tokenNm, tokenURI, tokenUriHash, tokenData string, owner sdk.AccAddress) error
	RemoveNFT(ctx sdk.Context, classID, tokenID string, owner sdk.AccAddress) error
	TransferNFT(ctx sdk.Context, classID, tokenID string, srcOwner, dstOwner sdk.AccAddress) error
	TransferClass(ctx sdk.Context, classID string, srcOwner, dstOwner sdk.AccAddress) error
	UpdateClassMintRestricted(ctx sdk.Context, classID string, mintRestricted bool, owner sdk.AccAddress) error

	GetClass(ctx sdk.Context, classID string) (Class, error)
	GetNFT(ctx sdk.Context, classID, tokenID string) (NFT, error)
}

type PermKeeper interface {
	HasL2UserRole(ctx sdk.Context, signer sdk.AccAddress) bool
}

type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI
	GetModuleAddress(name string) sdk.AccAddress
	GetModuleAccount(ctx sdk.Context, name string) authtypes.ModuleAccountI
}
