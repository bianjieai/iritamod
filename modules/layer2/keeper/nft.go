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

	moduleAddrStr := types.ModuleAddress.String()
	moduleAddr, err := sdk.AccAddressFromBech32(moduleAddrStr)
	if err != nil {
		return err
	}

	for _, tokenId := range tokenIds {
		owner, ok := k.GetTokenOwnerForNFT(ctx, spaceId, classId, tokenId)
		if !ok {
			return sdkerrors.Wrapf(types.ErrTokenForNFTNotExist, "spaceId: %d, classId: %s, tokenId: %s", spaceId, classId, tokenId)
		}
		k.deleteTokenOwnerForNFT(ctx, spaceId, classId, tokenId, owner)
		k.deleteTokenForNFT(ctx, spaceId, classId, tokenId)

		_, err := k.nft.GetNFT(ctx, classId, tokenId)
		if err != nil {
			continue
		}
		if err := k.nft.RemoveNFT(ctx, classId, tokenId, moduleAddr); err != nil {
			return nil
		}
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

	owner, err := sdk.AccAddressFromBech32(string(ownerbz))
	if err != nil {
		return nil, false
	}
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

	k.setClassForNFT(ctx, classId, bz)
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

	k.setClassForNFT(ctx, classId, bz)
	return nil
}

func (k Keeper) setClassForNFT(ctx sdk.Context, classId string, classBz []byte) {
	store := ctx.KVStore(k.storeKey)
	store.Set(classForNFTStoreKey(classId), classBz)
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

func (k Keeper) DepositL1ClassForNFT(ctx sdk.Context,
	classId,
	baseUri string,
	sender sdk.AccAddress) error {
	// check if the class exists
	class, err := k.nft.GetClass(ctx, classId)
	if err != nil {
		return err
	}
	// check if the denom owned by sender
	// TODO: fix this, not necessary the creator
	if class.GetCreator() != sender.String() {
		return sdkerrors.Wrapf(types.ErrClassNotOwnedByAccount, "class %s is not owned by %s", classId, sender)
	}

	if k.HasClassForNFT(ctx, classId) {
		if err := k.UpdateClassForNFT(ctx, classId, baseUri, sender.String()); err != nil {
			return err
		}
	} else {
		mintRestricted := class.GetMintRestricted()

		if err := k.CreateClassForNFT(ctx, classId, baseUri, sender.String(), mintRestricted); err != nil {
			return err
		}
	}

	moduleAddrStr := types.ModuleAddress.String()
	moduleAddr, err := sdk.AccAddressFromBech32(moduleAddrStr)
	if err != nil {
		return err
	}

	if err := k.nft.TransferClass(ctx, classId, sender, moduleAddr); err != nil {
		return err
	}
	return nil
}

func (k Keeper) WithdrawL2ClassForNFT(ctx sdk.Context,
	classId string,
	sender,
	owner sdk.AccAddress) error {
	// sender must have l2 user
	ok, err := k.HasL2UserRole(ctx, sender)
	if !ok {
		return err
	}

	// check if the class mapping exist
	classForNFT, exist := k.GetClassForNFT(ctx, classId)
	if !exist {
		return sdkerrors.Wrapf(types.ErrClassForNFTNotExist, "class mapping %s not exist", classId)
	}

	// check if the class exists
	class, err := k.nft.GetClass(ctx, classId)
	if err != nil {
		return err
	}

	// check if the class owned by module account
	if class.GetCreator() != types.ModuleAddress.String() {
		return sdkerrors.Wrapf(types.ErrClassNotOwnedByAccount, "class %s is not locked by %s", classId, types.ModuleAddress.String())
	}

	// check if the class mapping owner is msg.owner
	if classForNFT.Owner != owner.String() {
		return sdkerrors.Wrapf(types.ErrClassNotOwnedByAccount, "original owner want %s, got %s", classForNFT.Owner, sender)
	}

	// recover mint_restricted and transfer ownership
	moduleAddrStr := types.ModuleAddress.String()
	moduleAddr, err := sdk.AccAddressFromBech32(moduleAddrStr)
	if err != nil {
		return err
	}

	// recover mint_restricted
	if err := k.nft.UpdateClassMintRestricted(ctx, class.GetID(), classForNFT.Layer1MintRestricted, moduleAddr); err != nil {
		return err
	}

	if err := k.nft.TransferClass(ctx, classId, moduleAddr, owner); err != nil {
		return err
	}
	return nil
}

func (k Keeper) DepositL1TokenForNFT(ctx sdk.Context,
	spaceId uint64,
	classId,
	tokenId string,
	sender sdk.AccAddress) error {
	// token must exist
	nft, err := k.nft.GetNFT(ctx, classId, tokenId)
	if err != nil {
		return err
	}

	if !nft.GetOwner().Equals(sender) {
		return sdkerrors.Wrapf(types.ErrTokenForNFTNotOwnedByAccount, "nft %s is not owned by %s", tokenId, sender)
	}

	if !k.HasSpace(ctx, spaceId) {
		return sdkerrors.Wrapf(types.ErrUnknownSpace, "space %d not exist", spaceId)
	}

	if err := k.createTokenForNFT(ctx, spaceId, classId, tokenId, sender); err != nil {
		return err
	}

	moduleAddrStr := types.ModuleAddress.String()
	moduleAddr, err := sdk.AccAddressFromBech32(moduleAddrStr)
	if err != nil {
		return err
	}
	if err := k.nft.TransferNFT(ctx, classId, tokenId, sender, moduleAddr); err == nil {
		return err
	}
	return nil
}

func (k Keeper) WithdrawL2TokenForNFT(ctx sdk.Context,
	spaceId uint64,
	classId,
	tokenId,
	tokenName,
	tokenUri,
	tokenUriHash,
	tokenData string,
	sender,
	owner sdk.AccAddress,
) error {
	if ok, err := k.HasL2UserRole(ctx, sender); !ok {
		return err
	}

	if !k.HasSpaceByOwner(ctx, sender, spaceId) {
		return sdkerrors.Wrapf(types.ErrNotOwnerOfSpace, "space %d not owned by %s", spaceId, sender)
	}

	tokenOwner, exist := k.GetTokenOwnerForNFT(ctx, spaceId, classId, tokenId)
	if !exist {
		return sdkerrors.Wrapf(types.ErrTokenForNFTNotExist, "token %s not exist", tokenId)
	}
	if !tokenOwner.Equals(owner) {
		return sdkerrors.Wrapf(types.ErrTokenForNFTNotOwnedByAccount, "nft %s is not owned by %s", tokenId, owner)
	}

	_, err := k.nft.GetNFT(ctx, classId, tokenId)
	if err != nil {
		// no such nft, mint it
		if err := k.nft.SaveNFT(ctx, classId, tokenId, tokenName, tokenUri, tokenUriHash, tokenData, owner); err != nil {
			return err
		}
	} else {
		moduleAddrStr := types.ModuleAddress.String()
		moduleAddr, err := sdk.AccAddressFromBech32(moduleAddrStr)
		if err != nil {
			return err
		}

		// nft exist, update and transfer ownership
		if err := k.nft.UpdateNFT(ctx, classId, tokenId, tokenName, tokenUri, tokenUriHash, tokenData, moduleAddr); err != nil {
			return err
		}

		if err := k.nft.TransferNFT(ctx, classId, tokenId, moduleAddr, owner); err != nil {
			return err
		}
	}

	k.deleteTokenOwnerForNFT(ctx, spaceId, classId, tokenId, owner)
	k.deleteTokenForNFT(ctx, spaceId, classId, tokenId)

	return nil
}
