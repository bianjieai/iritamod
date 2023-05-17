package keeper

import (
	"github.com/bianjieai/iritamod/modules/layer2/types"
	"github.com/bianjieai/iritamod/modules/perm"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"math"
)

// CreateSpace creates a new space
func (k Keeper) CreateSpace(ctx sdk.Context, name, uri string, sender sdk.AccAddress) (uint64, error) {
	ok, err := k.HasL2UserRole(ctx, sender)
	if !ok {
		return 0, err
	}

	// increment the max space id and save it
	k.incrSpaceId(ctx)
	spaceId := k.getSpaceId(ctx)

	space := types.Space{
		Id:    spaceId,
		Name:  name,
		Uri:   uri,
		Owner: sender.String(),
	}

	k.setSpace(ctx, spaceId, space)
	k.setSpaceOfOwner(ctx, spaceId, sender)
	return spaceId, nil
}

// TransferSpace transfer the space ownership
func (k Keeper) TransferSpace(ctx sdk.Context, spaceId uint64, from, to sdk.AccAddress) error {
	ok, err := k.HasL2UserRole(ctx, from)
	if !ok {
		return err
	}
	ok, err = k.HasL2UserRole(ctx, to)
	if !ok {
		return err
	}

	if ok := k.HasSpaceOfOwner(ctx, from, spaceId); !ok {
		return sdkerrors.Wrapf(types.ErrNotOwnerOfSpace, "spaceId: %d is not owned by: %s", spaceId, from)
	}

	space, ok := k.GetSpace(ctx, spaceId)
	if !ok {
		return sdkerrors.Wrapf(types.ErrUnknownSpace, "fail to get space: %d", spaceId)
	}
	space.Owner = to.String()

	k.setSpace(ctx, spaceId, space)
	k.deleteSpaceOfOwner(ctx, spaceId, from)
	k.setSpaceOfOwner(ctx, spaceId, to)

	return nil
}

// CreateBlockHeader creates a layer2 block header record
func (k Keeper) CreateBlockHeader(ctx sdk.Context, spaceId, height uint64, header string, addr sdk.AccAddress) error {
	ok, err := k.HasL2UserRole(ctx, addr)
	if !ok {
		return err
	}

	if k.HasL2BlockHeader(ctx, spaceId, height) {
		return sdkerrors.Wrapf(types.ErrRecordAlreadyExist, "space: %d, height: %d", spaceId, height)
	}

	k.setL2BlockHeader(ctx, spaceId, height, header)
	return nil
}

// HasL2UserRole checks if an account has the l2 user role
func (k Keeper) HasL2UserRole(ctx sdk.Context, address sdk.AccAddress) (bool, error) {
	if err := k.perm.Access(ctx, address, perm.RoleLayer2User.Auth()); err != nil {
		return false, err
	}
	return true, nil
}

// GetSpaceId return the current max id for space
// we save the current max id into Space Store
// <0x01><math.maxUint64> -> <currMaxSpaceId>
func (k Keeper) getSpaceId(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	spaceCountKey := types.SpaceStoreKey(math.MaxUint64)
	bz := store.Get(spaceCountKey)
	return sdk.BigEndianToUint64(bz)
}

// setSpaceId set the current max id for space
func (k Keeper) setSpaceId(ctx sdk.Context, spaceId uint64) {
	store := ctx.KVStore(k.storeKey)
	spaceCountKey := types.SpaceStoreKey(math.MaxUint64)
	store.Set(spaceCountKey, sdk.Uint64ToBigEndian(spaceId))
}

// incrSpaceId increment the space unique id
func (k Keeper) incrSpaceId(ctx sdk.Context) {
	store := ctx.KVStore(k.storeKey)
	currSpaceId := k.getSpaceId(ctx)
	spaceCountKey := types.SpaceStoreKey(math.MaxUint64)
	store.Set(spaceCountKey, sdk.Uint64ToBigEndian(currSpaceId+1))
}

func (k Keeper) GetSpace(ctx sdk.Context, spaceId uint64) (types.Space, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.SpaceStoreKey(spaceId))
	if bz == nil {
		return types.Space{}, false
	}

	var space types.Space
	k.cdc.MustUnmarshal(bz, &space)
	return space, true
}

func (k Keeper) HasSpace(ctx sdk.Context, spaceId uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.SpaceStoreKey(spaceId))
}

func (k Keeper) setSpace(ctx sdk.Context, spaceId uint64, space types.Space) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&space)
	store.Set(types.SpaceStoreKey(spaceId), bz)
	k.incrSpaceId(ctx)
}

func (k Keeper) HasSpaceOfOwner(ctx sdk.Context, owner sdk.AccAddress, spaceId uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.SpaceOfOwnerStoreKey(owner, spaceId))
}

func (k Keeper) setSpaceOfOwner(ctx sdk.Context, spaceId uint64, owner sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.SpaceOfOwnerStoreKey(owner, spaceId), types.Placeholder)
}

func (k Keeper) deleteSpaceOfOwner(ctx sdk.Context, spaceId uint64, owner sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.SpaceOfOwnerStoreKey(owner, spaceId))
}

func (k Keeper) HasL2BlockHeader(ctx sdk.Context, spaceId, height uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.L2BlockHeaderStoreKey(spaceId, height))
}

func (k Keeper) setL2BlockHeader(ctx sdk.Context, spaceId, blockHeight uint64, header string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.L2BlockHeaderStoreKey(spaceId, blockHeight), []byte(header))
}
