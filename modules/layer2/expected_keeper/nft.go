package expected_keeper

import sdk "github.com/cosmos/cosmos-sdk/types"

// Class defines the interface specifications of collection that can be transferred across chains
type Class interface {
	GetID() string
	GetURI() string
	GetData() string
}

// NFT defines the interface specification of nft that can be transferred across chains
type NFT interface {
	GetClassID() string
	GetID() string
	GetURI() string
	GetData() string
}

type NFTKeeper interface {
	CreateOrUpdateClass(ctx sdk.Context, classID, classURI string, classData string) error
	Mint(ctx sdk.Context, classID, tokenID, tokenURI string, tokenData string, receiver sdk.AccAddress) error
	Transfer(ctx sdk.Context, classID string, tokenID string, tokenData string, receiver sdk.AccAddress) error
	Burn(ctx sdk.Context, classID string, tokenID string) error

	GetOwner(ctx sdk.Context, classID string, tokenID string) sdk.AccAddress
	HasClass(ctx sdk.Context, classID string) bool
	GetClass(ctx sdk.Context, classID string) (Class, bool)
	GetNFT(ctx sdk.Context, classID, tokenID string) (NFT, bool)
}
