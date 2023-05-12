package keeper

import (
	"github.com/bianjieai/iritamod/modules/layer2/types"
	"github.com/bianjieai/iritamod/modules/perm"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"math"
)

// CreateSpace creates a new space
func (k Keeper) CreateSpace(ctx sdk.Context, addr sdk.AccAddress) (uint64, error) {
	ok, err := k.HasL2UserRole(ctx, addr)
	if !ok {
		return 0, err
	}

	k.incrSpaceId(ctx)
	spaceId := k.getSpaceId(ctx)

	k.setSpace(ctx, spaceId, addr)
	k.setSpaceOwner(ctx, spaceId, addr)
	return spaceId, nil
}

// UpdateSpace updates the space info
func (k Keeper) UpdateSpace(ctx sdk.Context, spaceId uint64, from, to sdk.AccAddress) error {
	ok, err := k.HasL2UserRole(ctx, from)
	if !ok {
		return err
	}
	ok, err = k.HasL2UserRole(ctx, to)
	if !ok {
		return err
	}

	if ok := k.HasSpaceByOwner(ctx, from, spaceId); !ok {
		return sdkerrors.Wrapf(types.ErrNotOwnerOfSpace, "spaceId: %d is not owned by: %s", spaceId, from)
	}

	k.setSpace(ctx, spaceId, to)
	k.deleteSpaceOwner(ctx, spaceId, from)
	k.setSpaceOwner(ctx, spaceId, to)

	return nil
}

func (k Keeper) HasSpace(ctx sdk.Context, spaceId uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(spaceStoreKey(spaceId))
}

// GetSpaceOwner returns the space owner if space exists
func (k Keeper) GetSpaceOwner(ctx sdk.Context, spaceId uint64) sdk.AccAddress {
	if !k.HasSpace(ctx, spaceId) {
		return nil
	}
	store := ctx.KVStore(k.storeKey)
	ownerbz := store.Get(spaceStoreKey(spaceId))
	owner, _ := sdk.AccAddressFromBech32(string(ownerbz))
	return owner
}

// HasSpaceByOwner return ture if the owner has the space
func (k Keeper) HasSpaceByOwner(ctx sdk.Context,owner sdk.AccAddress, spaceId uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(spaceOfL2UserStoreKey(owner, spaceId))
}

// CreateRecord creates a layer2 block header record
func (k Keeper) CreateRecord(ctx sdk.Context, spaceId, height uint64, header string, addr sdk.AccAddress) error {
	ok, err := k.HasL2UserRole(ctx, addr)
	if !ok {
		return err
	}

	if k.HasRecord(ctx, spaceId, height) {
		return sdkerrors.Wrapf(types.ErrRecordAlreadyExist,"space: d%, height: d%", spaceId, height)
	}

	k.setRecord(ctx, spaceId, height, header)
	return nil
}

func (k Keeper) HasRecord(ctx sdk.Context, spaceId, height uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(recordStoreKey(spaceId, height))
}

// HasL2UserRole checks if an account has the l2 user role
func (k Keeper) HasL2UserRole(ctx sdk.Context, address sdk.AccAddress) (bool, error) {
	if err := k.perm.Access(ctx, address, perm.RoleLayer2User.Auth()); err != nil {
		return false, err
	}
	return true, nil
}

// getSpaceId return the current max id for space
// we save the current max id into Space Store
// <0x01><math.maxUint64> -> <currMaxSpaceId>
func (k Keeper) getSpaceId(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	spaceCountKey := spaceStoreKey(math.MaxUint64)
	bz := store.Get(spaceCountKey)
	return sdk.BigEndianToUint64(bz)
}

// incrSpaceId increment the space unique id
func (k Keeper) incrSpaceId(ctx sdk.Context) {
	store := ctx.KVStore(k.storeKey)
	currSpaceId := k.getSpaceId(ctx)
	spaceCountKey := spaceStoreKey(math.MaxUint64)
	store.Set(spaceCountKey, sdk.Uint64ToBigEndian(currSpaceId + 1))
}

// setSpace set the space owner
func (k Keeper) setSpace(ctx sdk.Context, spaceId uint64, owner sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Set(spaceStoreKey(spaceId), owner.Bytes())
	k.incrSpaceId(ctx)
}

func (k Keeper) setSpaceOwner(ctx sdk.Context, spaceId uint64, owner sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Set(spaceOfL2UserStoreKey(owner, spaceId), Placeholder)
}

func (k Keeper) deleteSpaceOwner(ctx sdk.Context, spaceId uint64, owner sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(spaceOfL2UserStoreKey(owner, spaceId))
}

func (k Keeper) setRecord(ctx sdk.Context, spaceId, blockHeight uint64, header string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(recordStoreKey(spaceId, blockHeight), []byte(header))
}