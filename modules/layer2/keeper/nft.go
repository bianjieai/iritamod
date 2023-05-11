package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) CreateTokenForNFT(ctx sdk.Context,
	spaceId int64,
	classId,
	tokenId string,
	owner sdk.AccAddress) error {
	panic("implement me")
}

func (k Keeper) UpdateTokenForNFT(ctx sdk.Context,
	spaceId int64,
	classId,
	tokenId string,
	owner sdk.AccAddress) error {
	panic("implement me")
}

func (k Keeper) DeleteTokenForNFT(ctx sdk.Context,
	spaceId int64,
	classId,
	tokenId string) error {
	panic("implement me")
}

// HasTokenForNFT check if layer2 module has this native nft mapping.
func (k Keeper) HasTokenForNFT(ctx sdk.Context,
	spaceId int64,
	classId,
	tokenId string) bool {
	panic("implement me")
}

func (k Keeper) CreateClassForNFT() {
	panic("implement me")
}

func (k Keeper) UpdateClassForNFT() {
	panic("implement me")
}

func (k Keeper) HasClassForNFT(ctx sdk.Context,
	classId string) bool {
	panic("implement me")
}
