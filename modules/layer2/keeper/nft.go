package keeper

import (
	"github.com/bianjieai/iritamod/modules/layer2/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// CreateTokensForNFT batch create token mappings for nft
// NOTE: it's allowed to create nft mappings without class mapping existence; think about a user transfer one nft
// from layer1 to layer2, then layer2 to layer1. During this process, the class are never deposited.
func (k Keeper) CreateTokensForNFT(ctx sdk.Context,
	spaceId uint64,
	classId string,
	nfts []*types.TokenForNFT,
	sender sdk.AccAddress) error {
	ok, err := k.HasL2UserRole(ctx, sender)
	if !ok {
		return err
	}

	if !k.HasSpaceByOwner(ctx, sender, spaceId) {
		return sdkerrors.Wrapf(types.ErrNotOwnerOfSpace, "spaceId: %d is not owned by: %s", spaceId, sender)
	}

	for _, nft := range nfts {
		owner, err := sdk.AccAddressFromBech32(nft.Owner)
		if err != nil {
			return err
		}

		if err := k.createTokenForNFT(ctx, spaceId, classId, nft.Id, owner); err != nil {
			return err
		}
	}

	return nil
}

// createTokenForNFT create a token mapping for nft.
// NOTE: examine the existence of spaceId before calling this function.
func (k Keeper) createTokenForNFT(ctx sdk.Context,
	spaceId uint64,
	classId,
	tokenId string,
	owner sdk.AccAddress) error {
	if k.HasTokenForNFT(ctx, spaceId, classId, tokenId) {
		return sdkerrors.Wrapf(types.ErrTokenForNFTAlreadyExist, "spaceId: %d, classId: %s, tokenId: %s", spaceId, classId, tokenId)
	}

	k.setTokenForNFT(ctx, spaceId, classId, tokenId, owner)
	k.setTokenOwnerForNFT(ctx, spaceId, classId, tokenId, owner)
	return nil
}

func (k Keeper) UpdateTokensForNFT(ctx sdk.Context,
	spaceId uint64,
	classId string,
	nfts []*types.TokenForNFT,
	sender sdk.AccAddress) error {
	ok, err := k.HasL2UserRole(ctx, sender)
	if !ok {
		return err
	}

	if !k.HasSpaceByOwner(ctx, sender, spaceId) {
		return sdkerrors.Wrapf(types.ErrNotOwnerOfSpace, "spaceId: %d is not owned by: %s", spaceId, sender)
	}

	for _, nft := range nfts {
		owner, err := sdk.AccAddressFromBech32(nft.Owner)
		if err != nil {
			return err
		}

		if !k.HasTokenForNFT(ctx, spaceId, classId, nft.Id) {
			return sdkerrors.Wrapf(types.ErrTokenForNFTNotExist, "spaceId: %d, classId: %s, tokenId: %s", spaceId, classId, nft.Id)
		}

		k.setTokenForNFT(ctx, spaceId, classId, nft.Id, owner)
		k.deleteTokenOwnerForNFT(ctx, spaceId, classId, nft.Id, owner)
		k.setTokenOwnerForNFT(ctx, spaceId, classId, nft.Id, owner)
	}
	return nil
}

// DeleteTokensForNFT delete a token mapping for nft
func (k Keeper) DeleteTokensForNFT(ctx sdk.Context,
	spaceId uint64,
	classId string,
	tokenIds []string,
	sender sdk.AccAddress) error {
	ok, err := k.HasL2UserRole(ctx, sender)
	if !ok {
		return err
	}

	if !k.HasSpaceByOwner(ctx, sender, spaceId) {
		return sdkerrors.Wrapf(types.ErrNotOwnerOfSpace, "spaceId: %d is not owned by: %s", spaceId, sender)
	}

	for _, tokenId := range tokenIds {
		owner, ok := k.GetTokenOwnerForNFT(ctx, spaceId, classId, tokenId)
		if !ok {
			return sdkerrors.Wrapf(types.ErrTokenForNFTNotExist, "spaceId: %d, classId: %s, tokenId: %s", spaceId, classId, tokenId)
		}
		k.deleteTokenOwnerForNFT(ctx, spaceId, classId, tokenId, owner)
		k.deleteTokenForNFT(ctx, spaceId, classId, tokenId)
	}

	return nil
}

func (k Keeper) GetTokenOwnerForNFT(ctx sdk.Context,
	spaceId uint64,
	classId,
	tokenId string) (sdk.AccAddress, bool) {
	store := ctx.KVStore(k.storeKey)
	tokenKey := tokenForNFTStoreKey(spaceId, classId, tokenId)
	if !store.Has(tokenKey) {
		return nil, false
	}
	ownerbz := store.Get(tokenKey)
	owner, _ := sdk.AccAddressFromBech32(string(ownerbz))
	return owner, true
}

// HasTokenForNFT check if layer2 module has this native nft mapping.
func (k Keeper) HasTokenForNFT(ctx sdk.Context,
	spaceId uint64,
	classId,
	tokenId string) bool {
	store := ctx.KVStore(k.storeKey)
	tokenKey := tokenForNFTStoreKey(spaceId, classId, tokenId)
	return store.Has(tokenKey)
}

// CreateClassForNFT creates  a class mapping for nft.
// NOTE: examine the existence of class mapping and ownership of layer 1 class before calling.
func (k Keeper) CreateClassForNFT(ctx sdk.Context,
	classId,
	baseUri,
	owner string,
	mintRestricted bool) error {

	class := types.ClassForNFT{
		Id:                   classId,
		Owner:                owner,
		BaseUri:              baseUri,
		Layer1MintRestricted: mintRestricted,
	}

	bz, err := k.cdc.Marshal(&class)
	if err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	store.Set(classForNFTStoreKey(classId), bz)
	return nil
}

func (k Keeper) UpdateClassForNFT(ctx sdk.Context,
	classId,
	baseUri,
	owner string) error {
	class, ok := k.GetClassForNFT(ctx, classId)
	if !ok {
		return sdkerrors.Wrapf(types.ErrClassForNFTNotExist, "classId: %s", classId)
	}

	class.BaseUri = baseUri
	class.Owner = owner

	bz, err := k.cdc.Marshal(&class)
	if err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	store.Set(classForNFTStoreKey(classId), bz)
	return nil
}

func (k Keeper) HasClassForNFT(ctx sdk.Context,
	classId string) bool {
	store := ctx.KVStore(k.storeKey)
	classKey := classForNFTStoreKey(classId)
	return store.Has(classKey)
}

func (k Keeper) GetClassForNFT(ctx sdk.Context,
	classId string) (types.ClassForNFT, bool) {
	store := ctx.KVStore(k.storeKey)
	classKey := classForNFTStoreKey(classId)
	var classForNFT types.ClassForNFT
	if !store.Has(classKey) {
		return classForNFT, false
	}

	classBz := store.Get(classKey)
	k.cdc.MustUnmarshal(classBz, &classForNFT)
	return classForNFT, true
}

func (k Keeper) setTokenForNFT(ctx sdk.Context,
	spaceId uint64,
	classId,
	tokenId string,
	owner sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	tokenKey := tokenForNFTStoreKey(spaceId, classId, tokenId)
	store.Set(tokenKey, owner.Bytes())
}

func (k Keeper) deleteTokenForNFT(ctx sdk.Context,
	spaceId uint64,
	classId,
	tokenId string) {
	store := ctx.KVStore(k.storeKey)
	tokenKey := tokenForNFTStoreKey(spaceId, classId, tokenId)
	store.Delete(tokenKey)
}

func (k Keeper) setTokenOwnerForNFT(ctx sdk.Context,
	spaceId uint64,
	classId,
	tokenId string,
	owner sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	nftsOwnerKey := nftsOfOwnerStoreKey(owner, spaceId, classId, tokenId)
	store.Set(nftsOwnerKey, Placeholder)
}

func (k Keeper) deleteTokenOwnerForNFT(ctx sdk.Context,
	spaceId uint64,
	classId,
	tokenId string,
	owner sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	nftsOwnerKey := nftsOfOwnerStoreKey(owner, spaceId, classId, tokenId)
	store.Delete(nftsOwnerKey)
}
