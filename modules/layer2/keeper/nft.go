package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/bianjieai/iritamod/modules/layer2/types"
)

// CreateNFTs batch create token mappings for nft
// NOTE: it's allowed to create nft mappings without class mapping existence; think about a user transfer one nft
// from layer1 to layer2, then layer2 to layer1. During this process, the class are never deposited.
func (k Keeper) CreateNFTs(ctx sdk.Context,
	spaceId uint64,
	classId string,
	nfts []types.TokenForNFT,
	sender sdk.AccAddress) error {
	if !k.HasSpace(ctx, spaceId) {
		return sdkerrors.Wrapf(types.ErrInvalidSpace, "space (%d) not exist", spaceId)
	}

	if !k.HasSpaceOfOwner(ctx, sender, spaceId) {
		return sdkerrors.Wrapf(types.ErrNotOwnerOfSpace, "space (%d) is not owned by (%s)", spaceId, sender)
	}

	class, err := k.nft.GetClass(ctx, classId)
	if err != nil {
		return sdkerrors.Wrapf(types.ErrInvalidClassId, "class (%s) not exist on layer one", classId)
	}

	moduleAddr := k.acc.GetModuleAddress(types.ModuleName)
	if class.GetOwner() != moduleAddr.String() {
		return sdkerrors.Wrapf(types.ErrInvalidClassId, "class (%s) is not locked on layer one", classId)
	}

	if _, err := k.GetClassForNFT(ctx, classId); err != nil {
		return sdkerrors.Wrapf(types.ErrInvalidClassId, "class mapping (%s) not exist", classId)
	}

	for _, nft := range nfts {
		owner, err := sdk.AccAddressFromBech32(nft.Owner)
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, err.Error())
		}

		if k.HasTokenForNFT(ctx, spaceId, classId, nft.Id) {
			return sdkerrors.Wrapf(types.ErrInvalidTokenId, "token (%s) already exist", nft.Id)
		}

		// check if the layer 1 nft exist
		if _, err := k.nft.GetNFT(ctx, classId, nft.Id); err == nil {
			return sdkerrors.Wrapf(types.ErrInvalidTokenId, "token (%s) already exist on layer one", nft.Id)
		}

		k.setTokenForNFT(ctx, spaceId, classId, nft.Id, owner)
		k.setTokenOwnerForNFT(ctx, spaceId, classId, nft.Id, owner)
	}

	return nil
}

func (k Keeper) UpdateNFTs(ctx sdk.Context,
	spaceId uint64,
	classId string,
	nfts []types.TokenForNFT,
	sender sdk.AccAddress) error {
	if !k.HasSpace(ctx, spaceId) {
		return sdkerrors.Wrapf(types.ErrInvalidSpace, "space (%d) not exist", spaceId)
	}

	if !k.HasSpaceOfOwner(ctx, sender, spaceId) {
		return sdkerrors.Wrapf(types.ErrNotOwnerOfSpace, "space (%d) is not owned by sender (%s)", spaceId, sender)
	}

	if _, err := k.nft.GetClass(ctx, classId); err != nil {
		return sdkerrors.Wrapf(types.ErrInvalidClassId, "class (%s) not exist on layer one", classId)
	}

	for _, nft := range nfts {
		owner, err := sdk.AccAddressFromBech32(nft.Owner)
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, err.Error())
		}

		if !k.HasTokenForNFT(ctx, spaceId, classId, nft.Id) {
			return sdkerrors.Wrapf(types.ErrInvalidTokenId, "token (%s) not exist under class (%s) in space (%d)", nft.Id, classId, spaceId)
		}

		k.setTokenForNFT(ctx, spaceId, classId, nft.Id, owner)
		k.deleteTokenOwnerForNFT(ctx, spaceId, classId, nft.Id, owner)
		k.setTokenOwnerForNFT(ctx, spaceId, classId, nft.Id, owner)
	}
	return nil
}

// DeleteNFTs delete a token mapping for nft
func (k Keeper) DeleteNFTs(ctx sdk.Context,
	spaceId uint64,
	classId string,
	tokenIds []string,
	sender sdk.AccAddress) error {
	if !k.HasSpace(ctx, spaceId) {
		return sdkerrors.Wrapf(types.ErrInvalidSpace, "space (%d) not exist", spaceId)
	}

	if !k.HasSpaceOfOwner(ctx, sender, spaceId) {
		return sdkerrors.Wrapf(types.ErrNotOwnerOfSpace, "space (%d) is not owned by (%s)", spaceId, sender)
	}

	if _, err := k.nft.GetClass(ctx, classId); err != nil {
		return sdkerrors.Wrapf(types.ErrInvalidClassId, "class (%s) not exist on layer one", classId)
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

		moduleAddr := k.acc.GetModuleAddress(types.ModuleName)
		if err := k.nft.RemoveNFT(ctx, classId, tokenId, moduleAddr); err != nil {
			return sdkerrors.Wrapf(types.ErrFromNftModule, "failed to burn token (%s), error (%s)", tokenId, err.Error())
		}
	}

	return nil
}

// UpdateClassesForNFT updates class mappings for nft
func (k Keeper) UpdateClassesForNFT(ctx sdk.Context,
	spaceId uint64,
	classUpdates []types.UpdateClassForNFT,
	sender sdk.AccAddress) error {
	if !k.HasSpace(ctx, spaceId) {
		return sdkerrors.Wrapf(types.ErrInvalidSpace, "space (%d) not exist", spaceId)
	}

	if !k.HasSpaceOfOwner(ctx, sender, spaceId) {
		return sdkerrors.Wrapf(types.ErrNotOwnerOfSpace, "space (%d) is not owned by (%s)", spaceId, sender)
	}

	for _, classUpdate := range classUpdates {
		_, err := sdk.AccAddressFromBech32(classUpdate.Owner)
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, err.Error())
		}

		class, err := k.GetClassForNFT(ctx, classUpdate.Id)
		if err != nil {
			return err
		}

		currSpaceId, err := k.GetSpaceOfClassForNFT(ctx, classUpdate.Id)
		if err != nil {
			return err
		}
		if spaceId != currSpaceId {
			return sdkerrors.Wrapf(types.ErrNotSpaceOfClass, "class not active on space (%d)", spaceId)
		}

		class.Owner = classUpdate.Owner
		class.BaseUri = classUpdate.Uri

		k.setClassForNFT(ctx, class)
	}

	return nil
}

func (k Keeper) DepositClassForNFT(ctx sdk.Context,
	spaceId uint64,
	classId,
	baseUri string,
	recipient,
	sender sdk.AccAddress) error {
	if !k.HasSpace(ctx, spaceId) {
		return sdkerrors.Wrapf(types.ErrInvalidSpace, "space (%d) not exist", spaceId)
	}

	// check if the class exists
	class, err := k.nft.GetClass(ctx, classId)
	if err != nil {
		return sdkerrors.Wrapf(types.ErrInvalidClassId, "class (%s) not exist on layer one", classId)
	}
	// check if the denom owned by sender
	if class.GetOwner() != sender.String() {
		return sdkerrors.Wrapf(types.ErrNotClassOwner, "class (%s) is not owned by (%s)", classId, sender)
	}

	if !k.GetPermKeeper().HasL2UserRole(ctx, sender) && !sender.Equals(recipient) {
		return sdkerrors.Wrapf(types.ErrInvalidL2User, "recipient (%s) must be sender if it is not l2 user", recipient)
	}

	classForNFT := types.ClassForNFT{
		Id:                   classId,
		Owner:                recipient.String(),
		BaseUri:              baseUri,
		Layer1MintRestricted: class.GetMintRestricted(),
	}

	k.setClassForNFT(ctx, classForNFT)
	k.setSpaceOfClassForNFT(ctx, classId, spaceId)

	moduleAddr := k.acc.GetModuleAddress(types.ModuleName)
	if err := k.nft.UpdateClassMintRestricted(ctx, classId, true, sender); err != nil {
		return sdkerrors.Wrapf(types.ErrFromNftModule, "failed to update class (%s) mint restricted, error (%s)", classId, err.Error())
	}

	return k.nft.TransferClass(ctx, classId, sender, moduleAddr)
}

func (k Keeper) WithdrawClassForNFT(ctx sdk.Context,
	spaceId uint64,
	classId string,
	owner,
	sender sdk.AccAddress) error {
	if !k.HasSpace(ctx, spaceId) {
		return sdkerrors.Wrapf(types.ErrInvalidSpace, "space (%d) not exist", spaceId)
	}

	// sender must have this space
	if !k.HasSpaceOfOwner(ctx, sender, spaceId) {
		return sdkerrors.Wrapf(types.ErrNotOwnerOfSpace, "space (%d) is not owned by (%s)", spaceId, sender)
	}

	// check if the class mapping exist
	classForNFT, err := k.GetClassForNFT(ctx, classId)
	if err != nil {
		return err
	}

	// class must on corresponding space
	currSpaceId, err := k.GetSpaceOfClassForNFT(ctx, classId)
	if err != nil {
		return err
	}
	if spaceId != currSpaceId {
		return sdkerrors.Wrapf(types.ErrNotSpaceOfClass, "class no active on space (%d)", spaceId)
	}

	// check if the class exists
	class, err := k.nft.GetClass(ctx, classId)
	if err != nil {
		return sdkerrors.Wrapf(types.ErrInvalidClassId, "class (%s) not exist on layer one", classId)
	}

	// check if the class owned by module account
	moduleAddr := k.acc.GetModuleAddress(types.ModuleName)
	if class.GetOwner() != moduleAddr.String() {
		return sdkerrors.Wrapf(types.ErrNotClassOwner, "class (%s) is not locked by (%s)", classId, moduleAddr.String())
	}

	// check if the class mapping owner is msg.owner
	if classForNFT.Owner != owner.String() {
		return sdkerrors.Wrapf(types.ErrNotClassOwner, "original owner want (%s) but got (%s)", classForNFT.Owner, sender)
	}

	// remove class corresponding space
	k.deleteSpaceOfClassForNFT(ctx, classId)

	// recover mint_restricted
	if err := k.nft.UpdateClassMintRestricted(ctx, class.GetID(), classForNFT.Layer1MintRestricted, moduleAddr); err != nil {
		return sdkerrors.Wrapf(types.ErrFromNftModule, "failed to update class (%s) mint restricted, error (%s)", classId, err.Error())
	}

	return k.nft.TransferClass(ctx, classId, moduleAddr, owner)
}

func (k Keeper) DepositTokenForNFT(ctx sdk.Context,
	spaceId uint64,
	classId,
	tokenId string,
	sender sdk.AccAddress) error {
	if !k.HasSpace(ctx, spaceId) {
		return sdkerrors.Wrapf(types.ErrInvalidSpace, "space (%d) not exist", spaceId)
	}

	// token must exist
	nft, err := k.nft.GetNFT(ctx, classId, tokenId)
	if err != nil {
		return sdkerrors.Wrapf(types.ErrInvalidTokenId, "token (%s) already exist on layer one", tokenId)
	}

	if !nft.GetOwner().Equals(sender) {
		return sdkerrors.Wrapf(types.ErrNotTokenOwner, "nft (%s) is not owned by (%s)", tokenId, sender)
	}

	k.setTokenForNFT(ctx, spaceId, classId, tokenId, sender)
	k.setTokenOwnerForNFT(ctx, spaceId, classId, tokenId, sender)

	moduleAddr := k.acc.GetModuleAddress(types.ModuleName)
	return k.nft.TransferNFT(ctx, classId, tokenId, sender, moduleAddr)
}

func (k Keeper) WithdrawTokenForNFT(ctx sdk.Context,
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
	if !k.HasSpace(ctx, spaceId) {
		return sdkerrors.Wrapf(types.ErrInvalidSpace, "space (%d) not exist", spaceId)
	}

	if !k.HasSpaceOfOwner(ctx, sender, spaceId) {
		return sdkerrors.Wrapf(types.ErrNotOwnerOfSpace, "space (%d) is not owned by (%s)", spaceId, sender)
	}

	if _, err := k.nft.GetClass(ctx, classId); err != nil {
		return sdkerrors.Wrapf(types.ErrInvalidClassId, "class (%s) not exist", classId)
	}

	tokenOwner, err := k.GetTokenOwnerForNFT(ctx, spaceId, classId, tokenId)
	if err != nil {
		return sdkerrors.Wrapf(types.ErrInvalidTokenId, "token (%s) not exist under class (%s) in space (%d)", tokenId, classId, spaceId)
	}
	if !tokenOwner.Equals(owner) {
		return sdkerrors.Wrapf(types.ErrNotTokenOwner, "nft (%s) is not owned by (%s)", tokenId, owner)
	}

	_, err = k.nft.GetNFT(ctx, classId, tokenId)
	if err != nil {
		// no such nft, mint it
		if err := k.nft.SaveNFT(ctx, classId, tokenId, tokenName, tokenUri, tokenUriHash, tokenData, owner); err != nil {
			return sdkerrors.Wrapf(types.ErrFromNftModule, "failed to mint token (%s), error (%s)", tokenId, err.Error())
		}
	} else {
		// nft exist, update and transfer ownership
		moduleAddr := k.acc.GetModuleAddress(types.ModuleName)
		if err := k.nft.UpdateNFT(ctx, classId, tokenId, tokenName, tokenUri, tokenUriHash, tokenData, moduleAddr); err != nil {
			return sdkerrors.Wrapf(types.ErrFromNftModule, "failed to edit token (%s), error (%s)", tokenId, err.Error())
		}

		if err := k.nft.TransferNFT(ctx, classId, tokenId, moduleAddr, owner); err != nil {
			return sdkerrors.Wrapf(types.ErrFromNftModule, "failed to transfer token (%s), error (%s)", tokenId, err.Error())
		}
	}

	k.deleteTokenOwnerForNFT(ctx, spaceId, classId, tokenId, owner)
	k.deleteTokenForNFT(ctx, spaceId, classId, tokenId)

	return nil
}

func (k Keeper) GetCollectionsForNFT(ctx sdk.Context) []types.CollectionForNFT {
	collections := make([]types.CollectionForNFT, 0)
	spaces := k.GetSpaces(ctx)
	classes := k.GetClassesForNFT(ctx)

	for _, space := range spaces {
		for _, class := range classes {
			var collection types.CollectionForNFT
			tokens := k.GetTokensForNFT(ctx, space.Id, class.Id)

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

func (k Keeper) GetClassesWithSpaceForNFT(ctx sdk.Context) []types.ClassWithSpaceForNFT {
	classesWithSpaceForNFT := make([]types.ClassWithSpaceForNFT, 0)
	classesForNFT := k.GetClassesForNFT(ctx)
	for _, classForNFT := range classesForNFT {
		classWithSpaceForNFT := types.ClassWithSpaceForNFT{
			Id:                   classForNFT.Id,
			Owner:                classForNFT.Owner,
			BaseUri:              classForNFT.BaseUri,
			Layer1MintRestricted: classForNFT.Layer1MintRestricted,
			ActiveSpace:          0,
		}

		spaceId, err := k.GetSpaceOfClassForNFT(ctx, classForNFT.Id)
		if err == nil {
			classWithSpaceForNFT.ActiveSpace = spaceId
		}

		classesWithSpaceForNFT = append(classesWithSpaceForNFT, classWithSpaceForNFT)
	}
	return classesWithSpaceForNFT
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
		return types.ClassForNFT{}, sdkerrors.Wrapf(types.ErrInvalidClassId, "unable to get class (%s) mapping ", classId)
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

func (k Keeper) HasSpaceOfClassForNFT(ctx sdk.Context, class string) bool {
	store := ctx.KVStore(k.storeKey)
	key := types.SpaceOfClassForNFTStoreKey(class)
	return store.Has(key)
}

func (k Keeper) GetSpaceOfClassForNFT(ctx sdk.Context, class string) (uint64, error) {
	store := ctx.KVStore(k.storeKey)
	key := types.SpaceOfClassForNFTStoreKey(class)
	bz := store.Get(key)
	if len(bz) == 0 {
		return 0, sdkerrors.Wrapf(types.ErrNotSpaceOfClass, "class (%s) not active on layer2", class)
	}
	return sdk.BigEndianToUint64(bz), nil
}

func (k Keeper) setSpaceOfClassForNFT(ctx sdk.Context, class string, spaceId uint64) {
	store := ctx.KVStore(k.storeKey)
	key := types.SpaceOfClassForNFTStoreKey(class)
	store.Set(key, sdk.Uint64ToBigEndian(spaceId))
}

func (k Keeper) deleteSpaceOfClassForNFT(ctx sdk.Context, class string) {
	store := ctx.KVStore(k.storeKey)
	key := types.SpaceOfClassForNFTStoreKey(class)
	store.Delete(key)
}

func (k Keeper) GetTokensForNFT(ctx sdk.Context, spaceId uint64, classId string) []types.TokenForNFT {
	store := k.getCollectionStore(ctx, spaceId, classId)
	tokens := make([]types.TokenForNFT, 0)

	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var tokenForNFT types.TokenForNFT
		tokenForNFT.Id = string(iterator.Key())
		tokenForNFT.Owner = sdk.AccAddress(iterator.Value()).String()

		tokens = append(tokens, tokenForNFT)
	}

	return tokens
}

func (k Keeper) GetTokenForNFT(ctx sdk.Context, spaceId uint64, classId, tokenId string) (sdk.AccAddress, error) {
	store := ctx.KVStore(k.storeKey)
	key := types.TokenForNFTStoreKey(spaceId, classId, tokenId)
	bz := store.Get(key)
	if len(bz) == 0 {
		return nil, sdkerrors.Wrapf(types.ErrInvalidTokenId, "unable to get token (%s) mapping", tokenId)
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
		return nil, sdkerrors.Wrapf(types.ErrInvalidTokenId, "unable to get token (%s) mapping", tokenId)
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

// prefix store of  <0x07><space_id><delimiter><class_id><delimiter>
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
