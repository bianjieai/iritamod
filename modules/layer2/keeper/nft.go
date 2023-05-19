package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/bianjieai/iritamod/modules/layer2/types"
)

// CreateTokensForNFT batch create token mappings for nft
// NOTE: it's allowed to create nft mappings without class mapping existence; think about a user transfer one nft
// from layer1 to layer2, then layer2 to layer1. During this process, the class are never deposited.
func (k Keeper) CreateTokensForNFT(ctx sdk.Context,
	spaceId uint64,
	classId string,
	nfts []types.TokenForNFT,
	sender sdk.AccAddress) error {
	if !k.HasL2UserRole(ctx, sender) {
		return sdkerrors.Wrapf(types.ErrNotL2UserRole, "address: %s", sender)
	}

	if !k.HasSpaceOfOwner(ctx, sender, spaceId) {
		return sdkerrors.Wrapf(types.ErrNotOwnerOfSpace, "spaceId: %d is not owned by: %s", spaceId, sender)
	}

	for _, nft := range nfts {
		owner, err := sdk.AccAddressFromBech32(nft.Owner)
		if err != nil {
			return err
		}

		if k.HasTokenForNFT(ctx, spaceId, classId, nft.Id) {
			return sdkerrors.Wrapf(types.ErrTokenForNFTAlreadyExist, "spaceId: %d, classId: %s, tokenId: %s", spaceId, classId, nft.Id)
		}

		k.setTokenForNFT(ctx, spaceId, classId, nft.Id, owner)
		k.setTokenOwnerForNFT(ctx, spaceId, classId, nft.Id, owner)
	}

	return nil
}

func (k Keeper) UpdateTokensForNFT(ctx sdk.Context,
	spaceId uint64,
	classId string,
	nfts []types.TokenForNFT,
	sender sdk.AccAddress) error {
	if !k.HasL2UserRole(ctx, sender) {
		return sdkerrors.Wrapf(types.ErrNotL2UserRole, "address: %s", sender)
	}

	if !k.HasSpaceOfOwner(ctx, sender, spaceId) {
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
	if !k.HasL2UserRole(ctx, sender) {
		return sdkerrors.Wrapf(types.ErrNotL2UserRole, "address: %s", sender)
	}

	if !k.HasSpaceOfOwner(ctx, sender, spaceId) {
		return sdkerrors.Wrapf(types.ErrNotOwnerOfSpace, "spaceId: %d is not owned by: %s", spaceId, sender)
	}

	for _, tokenId := range tokenIds {
		owner, err := k.GetTokenOwnerForNFT(ctx, spaceId, classId, tokenId)
		if err != nil {
			return err
		}
		k.deleteTokenOwnerForNFT(ctx, spaceId, classId, tokenId, owner)
		k.deleteTokenForNFT(ctx, spaceId, classId, tokenId)

		// burn the corresponding l1 nft if it exists.
		_, err = k.nft.GetNFT(ctx, classId, tokenId)
		if err != nil {
			continue
		}
		if err := k.nft.RemoveNFT(ctx, classId, tokenId, types.ModuleAccAddress); err != nil {
			return err
		}
	}

	return nil
}

// UpdateL2ClassesForNFT updates class mappings for nft
// TODO： class id corresponding space id
func (k Keeper) UpdateL2ClassesForNFT(ctx sdk.Context,
	classUpdates []types.UpdateClassForNFT,
	sender sdk.AccAddress) error {
	if !k.HasL2UserRole(ctx, sender) {
		return sdkerrors.Wrapf(types.ErrNotL2UserRole, "address: %s", sender)
	}

	for _, classUpdate := range classUpdates {
		_, err := sdk.AccAddressFromBech32(classUpdate.Owner)
		if err != nil {
			return err
		}

		class, err := k.GetClassForNFT(ctx, classUpdate.Id)
		if err != nil {
			return err
		}
		class.Owner = classUpdate.Owner
		class.BaseUri = classUpdate.Uri

		k.setClassForNFT(ctx, class)
	}

	return nil
}

// TODO： class id corresponding space id
func (k Keeper) DepositL1ClassForNFT(ctx sdk.Context,
	spaceId uint64,
	classId,
	baseUri string,
	recipient,
	sender sdk.AccAddress) error {
	if !k.HasSpace(ctx, spaceId) {
		return sdkerrors.Wrapf(types.ErrUnknownSpace, "space %d not exist", spaceId)
	}

	// check if the class exists
	class, err := k.nft.GetClass(ctx, classId)
	if err != nil {
		return err
	}
	// check if the denom owned by sender
	// TODO： => GetOwner()
	if class.GetCreator() != sender.String() {
		return sdkerrors.Wrapf(types.ErrClassNotOwnedByAccount, "class %s is not owned by %s", classId, sender)
	}

	// TODO: fix this
	if !k.HasL2UserRole(ctx, sender) && !sender.Equals(recipient) {
		return sdkerrors.Wrapf(types.ErrClassNotOwnedByAccount, "recipient %s must be sender if not l2 user", sender)
	}

	// TODO: k.HasClassForNFT
	classForNFT, err := k.GetClassForNFT(ctx, classId)
	classForNFT.BaseUri = baseUri
	classForNFT.Owner = recipient.String()
	if err != nil {
		// create class mapping for the first time
		classForNFT.Id = classId
		classForNFT.Layer1MintRestricted = class.GetMintRestricted()
	}
	k.setClassForNFT(ctx, classForNFT)

	if err := k.nft.TransferClass(ctx, classId, sender, types.ModuleAccAddress); err != nil {
		return err
	}
	return nil
}

// TODO: add space id
func (k Keeper) WithdrawL2ClassForNFT(ctx sdk.Context,
	classId string,
	owner,
	sender sdk.AccAddress) error {
	// sender must have l2 user
	if !k.HasL2UserRole(ctx, sender) {
		return sdkerrors.Wrapf(types.ErrNotL2UserRole, "address: %s", sender)
	}

	// check if the class mapping exist
	classForNFT, err := k.GetClassForNFT(ctx, classId)
	if err != nil {
		return err
	}

	// check if the class exists
	class, err := k.nft.GetClass(ctx, classId)
	if err != nil {
		return err
	}

	// check if the class owned by module account
	// TODO: GetOwner()
	if class.GetCreator() != types.ModuleAddress.String() {
		return sdkerrors.Wrapf(types.ErrClassNotOwnedByAccount, "class %s is not locked by %s", classId, types.ModuleAddress.String())
	}

	// check if the class mapping owner is msg.owner
	if classForNFT.Owner != owner.String() {
		return sdkerrors.Wrapf(types.ErrClassNotOwnedByAccount, "original owner want %s, got %s", classForNFT.Owner, sender)
	}

	// recover mint_restricted
	if err := k.nft.UpdateClassMintRestricted(ctx, class.GetID(), classForNFT.Layer1MintRestricted, types.ModuleAccAddress); err != nil {
		return err
	}

	if err := k.nft.TransferClass(ctx, classId, types.ModuleAccAddress, owner); err != nil {
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

	k.setTokenForNFT(ctx, spaceId, classId, tokenId, sender)
	k.setTokenOwnerForNFT(ctx, spaceId, classId, tokenId, sender)

	return  k.nft.TransferNFT(ctx, classId, tokenId, sender, types.ModuleAccAddress)
}

func (k Keeper) WithdrawL2TokenForNFT(ctx sdk.Context,
	spaceId uint64,
	classId,
	tokenId,
	tokenName,
	tokenUri,
	tokenUriHash,
	tokenData string,
	owner,
	sender sdk.AccAddress,
) error {
	if !k.HasL2UserRole(ctx, sender) {
		return sdkerrors.Wrapf(types.ErrNotL2UserRole, "address: %s", sender)
	}

	if !k.HasSpaceOfOwner(ctx, sender, spaceId) {
		return sdkerrors.Wrapf(types.ErrNotOwnerOfSpace, "space %d not owned by %s", spaceId, sender)
	}

	tokenOwner, err := k.GetTokenOwnerForNFT(ctx, spaceId, classId, tokenId)
	if err != nil {
		return sdkerrors.Wrapf(types.ErrTokenForNFTNotExist, "token %s not exist", tokenId)
	}
	if !tokenOwner.Equals(owner) {
		return sdkerrors.Wrapf(types.ErrTokenForNFTNotOwnedByAccount, "nft %s is not owned by %s", tokenId, owner)
	}

	_, err = k.nft.GetNFT(ctx, classId, tokenId)
	if err != nil {
		// no such nft, mint it
		if err := k.nft.SaveNFT(ctx, classId, tokenId, tokenName, tokenUri, tokenUriHash, tokenData, owner); err != nil {
			return err
		}
	} else {
		// nft exist, update and transfer ownership
		if err := k.nft.UpdateNFT(ctx, classId, tokenId, tokenName, tokenUri, tokenUriHash, tokenData, types.ModuleAccAddress); err != nil {
			return err
		}

		if err := k.nft.TransferNFT(ctx, classId, tokenId, types.ModuleAccAddress, owner); err != nil {
			return err
		}
	}

	k.deleteTokenOwnerForNFT(ctx, spaceId, classId, tokenId, owner)
	k.deleteTokenForNFT(ctx, spaceId, classId, tokenId)

	return nil
}

func (k Keeper) GetCollectionsForNFT(ctx sdk.Context) []types.CollectionForNFT {
	collections := make([]types.CollectionForNFT, 0)
	store := ctx.KVStore(k.storeKey)
	spaces := k.GetSpaces(ctx)
	classes := k.GetClassesForNFT(ctx)

	for _, space := range spaces {
		for _, class := range classes {
			var collection types.CollectionForNFT
			tokens := make([]types.TokenForNFT, 0)

			iterator := sdk.KVStorePrefixIterator(store, types.TokenForNFTByCollectionStoreKey(space.Id, class.Id))
			defer iterator.Close()

			for ; iterator.Valid(); iterator.Next() {
				var tokenForNFT types.TokenForNFT
				tokenForNFT.Id = string(iterator.Key())
				tokenForNFT.Owner = sdk.AccAddress(iterator.Value()).String()

				tokens = append(tokens, tokenForNFT)
			}

			if len(tokens) > 0 {
				collection.SpaceId = space.Id
				collection.ClassId = class.Id
				collection.Tokens = tokens
				collections = append(collections, collection)
			}
		}
	}
	return collections
}

func (k Keeper) GetClassesForNFT(ctx sdk.Context) []types.ClassForNFT {
	classes := make([]types.ClassForNFT, 0)
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefixClassForNFT)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var class types.ClassForNFT
		k.cdc.MustUnmarshal(iterator.Value(), &class)
		classes = append(classes, class)
	}
	return classes
}

func (k Keeper) GetClassForNFT(ctx sdk.Context, classId string) (types.ClassForNFT, error) {
	store := ctx.KVStore(k.storeKey)
	key := types.ClassForNFTStoreKey(classId)
	bz := store.Get(key)
	if len(bz) == 0 {
		return types.ClassForNFT{}, sdkerrors.Wrapf(types.ErrClassForNFTNotExist, "class mapping %s not exist", classId)
	}

	var classForNFT types.ClassForNFT
	k.cdc.MustUnmarshal(bz, &classForNFT)
	return classForNFT, nil
}

func (k Keeper) HasClassForNFT(ctx sdk.Context, classId string) bool {
	store := ctx.KVStore(k.storeKey)
	key := types.ClassForNFTStoreKey(classId)
	return store.Has(key)
}

func (k Keeper) setClassForNFT(ctx sdk.Context, class types.ClassForNFT) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&class)
	store.Set(types.ClassForNFTStoreKey(class.Id), bz)
}

func (k Keeper) GetTokenForNFT(ctx sdk.Context, spaceId uint64, classId, tokenId string) (sdk.AccAddress, error) {
	store := ctx.KVStore(k.storeKey)
	key := types.TokenForNFTStoreKey(spaceId, classId, tokenId)
	bz := store.Get(key)
	if len(bz) == 0 {
		return nil, sdkerrors.Wrapf(types.ErrTokenForNFTNotExist, "token %s not exist", tokenId)
	}
	return bz, nil
}

func (k Keeper) HasTokenForNFT(ctx sdk.Context, spaceId uint64, classId, tokenId string) bool {
	store := ctx.KVStore(k.storeKey)
	tokenKey := types.TokenForNFTStoreKey(spaceId, classId, tokenId)
	return store.Has(tokenKey)
}

func (k Keeper) setTokenForNFT(ctx sdk.Context,
	spaceId uint64,
	classId,
	tokenId string,
	owner sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	key := types.TokenForNFTStoreKey(spaceId, classId, tokenId)
	store.Set(key, owner.Bytes())
}

func (k Keeper) deleteTokenForNFT(ctx sdk.Context,
	spaceId uint64,
	classId,
	tokenId string) {
	store := ctx.KVStore(k.storeKey)
	tokenKey := types.TokenForNFTStoreKey(spaceId, classId, tokenId)
	store.Delete(tokenKey)
}

func (k Keeper) GetTokenOwnerForNFT(ctx sdk.Context, spaceId uint64, classId, tokenId string) (sdk.AccAddress, error) {
	store := ctx.KVStore(k.storeKey)
	key := types.TokenForNFTStoreKey(spaceId, classId, tokenId)
	bz := store.Get(key)
	if len(bz) == 0 {
		return nil, sdkerrors.Wrapf(types.ErrTokenForNFTNotExist, "token %s not exist", tokenId)
	}

	return bz, nil
}

func (k Keeper) setTokenOwnerForNFT(ctx sdk.Context,
	spaceId uint64,
	classId,
	tokenId string,
	owner sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	key := types.NFTsOfOwnerStoreKey(owner, spaceId, classId, tokenId)
	store.Set(key, types.Placeholder)
}

func (k Keeper) deleteTokenOwnerForNFT(ctx sdk.Context,
	spaceId uint64,
	classId,
	tokenId string,
	owner sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	key := types.NFTsOfOwnerStoreKey(owner, spaceId, classId, tokenId)
	store.Delete(key)
}

// prefix store of <0x05>
func (k Keeper) getClassStore(ctx sdk.Context) prefix.Store {
	store := ctx.KVStore(k.storeKey)
	return prefix.NewStore(store, types.KeyPrefixClassForNFT)
}

// prefix store of  <0x06><space_id><delimiter><class_id><delimiter>
func (k Keeper) getCollectionStore(ctx sdk.Context, spaceId uint64, classId string) prefix.Store {
	store := ctx.KVStore(k.storeKey)
	return prefix.NewStore(store, types.TokenForNFTByCollectionStoreKey(spaceId, classId))
}

func (k Keeper) getNFTsOfOwnerStore(ctx sdk.Context, owner sdk.AccAddress) prefix.Store {
	store := ctx.KVStore(k.storeKey)
	return prefix.NewStore(store, types.NFTsOfOwnerAllStoreKey(owner))
}

func (k Keeper) getNFTsOfOwnerBySpaceStore(ctx sdk.Context, owner sdk.AccAddress, spaceId uint64) prefix.Store {
	store := ctx.KVStore(k.storeKey)
	return prefix.NewStore(store, types.NFTsOfOwnerBySpaceStoreKey(owner, spaceId))
}

func (k Keeper) getNFTsOfOwnerBySpaceAndClassStore(ctx sdk.Context, owner sdk.AccAddress, spaceId uint64, classId string) prefix.Store {
	store := ctx.KVStore(k.storeKey)
	return prefix.NewStore(store, types.NFTsOfOwnerBySpaceAndClassStoreKey(owner, spaceId, classId))
}
