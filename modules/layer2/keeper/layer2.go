package keeper

import (
	"bytes"
	"math"
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/bianjieai/iritamod/modules/layer2/types"
	"github.com/bianjieai/iritamod/modules/perm"
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

	space, err := k.GetSpace(ctx, spaceId)
	if err != nil {
		return err
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

func (k Keeper) GetSpaces(ctx sdk.Context) []types.Space {
	spaces := make([]types.Space, 0)
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefixSpace)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var space types.Space
		k.cdc.MustUnmarshal(iterator.Value(), &space)
		spaces = append(spaces, space)
	}
	return spaces
}

func (k Keeper) GetSpace(ctx sdk.Context, spaceId uint64) (types.Space, error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.SpaceStoreKey(spaceId))
	if bz == nil {
		return types.Space{}, sdkerrors.Wrapf(types.ErrUnknownSpace, "spaceId: %d", spaceId)
	}

	var space types.Space
	k.cdc.MustUnmarshal(bz, &space)
	return space, nil
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

func (k Keeper) GetL2BlockHeaders(ctx sdk.Context) []types.L2BlockHeader {
	headers := make([]types.L2BlockHeader, 0)
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefixL2BlockHeader)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		ret := bytes.Split(iterator.Key(), types.Delimiter)
		if len(ret) != 2 {
			panic("fail to split layer2 block header key")
		}
		var l2BlockHeader types.L2BlockHeader
		spaceId, err := strconv.ParseUint(string(ret[0]), 10, 64)
		if err != nil {
			panic("fail to convert spaceId to uint64")
		}
		height, err := strconv.ParseUint(string(ret[1]), 10, 64)
		if err != nil {
			panic("fail to convert height to uint64")
		}

		l2BlockHeader.SpaceId = spaceId
		l2BlockHeader.Height = height
		l2BlockHeader.Header = string(iterator.Value())
		headers = append(headers, l2BlockHeader)
	}
	return headers
}

func (k Keeper) GetL2BlockHeader(ctx sdk.Context, spaceId, height uint64) (string, error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.L2BlockHeaderStoreKey(spaceId, height))
	if bz == nil {
		return "", sdkerrors.Wrapf(types.ErrUnknownL2BlockHeader, "spaceId: %d, height: %d", spaceId, height)
	}
	return string(bz), nil
}

func (k Keeper) HasL2BlockHeader(ctx sdk.Context, spaceId, height uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.L2BlockHeaderStoreKey(spaceId, height))
}

func (k Keeper) setL2BlockHeader(ctx sdk.Context, spaceId, blockHeight uint64, header string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.L2BlockHeaderStoreKey(spaceId, blockHeight), []byte(header))
}

func (k Keeper) getSpaceStore(ctx sdk.Context) prefix.Store {
	store := ctx.KVStore(k.storeKey)
	return prefix.NewStore(store, types.KeyPrefixSpace)
}

// getSpaceOfOwnerStore returns a prefix store of <0x02><owner><delimiter>
func (k Keeper) getSpaceOfOwnerStore(ctx sdk.Context, owner sdk.AccAddress) prefix.Store {
	store := ctx.KVStore(k.storeKey)
	key := types.SpaceOfOwnerByOwnerStoreKey(owner)
	return prefix.NewStore(store, key)
}
