package expected_keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type NFT interface {
	GetID() string
	GetName() string
	GetOwner() sdk.AccAddress
	GetURI() string
	GetURIHash() string
	GetData() string
}

type Class interface {
	GetID() string
	GetCreator() string
	GetMintRestricted() bool
}

type NFTKeeper interface {
	SaveNFT(ctx sdk.Context, classID, tokenID, tokenNm, tokenURI, tokenUriHash, tokenData string, receiver sdk.AccAddress) error
	UpdateNFT(ctx sdk.Context, classID, tokenID, tokenNm, tokenURI, tokenUriHash, tokenData string, owner sdk.AccAddress) error
	RemoveNFT(ctx sdk.Context, classID, tokenID string, owner sdk.AccAddress) error
	TransferNFT(ctx sdk.Context, classID, tokenID string, srcOwner, dstOwner sdk.AccAddress) error
	TransferClass(ctx sdk.Context, classID string, srcOwner, dstOwner sdk.AccAddress) error
	UpdateClassMintRestricted(ctx sdk.Context, classID Class, mintRestricted bool, owner sdk.AccAddress) error

	GetClass(ctx sdk.Context, classID string) (Class, error)
	GetNFT(ctx sdk.Context, classID, tokenID string) (NFT, error)
}
