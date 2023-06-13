package keeper

import (
	"bytes"
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/bianjieai/iritamod/modules/side-chain/types"
)

// CreateSpace creates a new space
func (k Keeper) CreateSpace(ctx sdk.Context, name, uri string, sender sdk.AccAddress) (uint64, error) {
	// increment the max space id and save it
	k.incrSpaceSequence(ctx)
	spaceId := k.GetSpaceSequence(ctx)

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
func (k Keeper) CreateBlockHeader(ctx sdk.Context, spaceId, height uint64, header string, sender sdk.AccAddress) error {
	if k.HasBlockHeader(ctx, spaceId, height) {
		return sdkerrors.Wrapf(types.ErrBlockHeader, "block header already exists at height (%d) in space (%d)", height, spaceId)
	}

	k.setBlockHeader(ctx, spaceId, height, header)
	return nil
}

// GetSpaceSequence return the current sequence for space
// <0x01> -> <currSpaceSequence>
func (k Keeper) GetSpaceSequence(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	key := types.SpaceSequenceStoreKey()
	bz := store.Get(key)
	return sdk.BigEndianToUint64(bz)
}

// setSpaceSequence set the current sequence for space
func (k Keeper) setSpaceSequence(ctx sdk.Context, spaceId uint64) {
	store := ctx.KVStore(k.storeKey)
	key := types.SpaceSequenceStoreKey()
	store.Set(key, sdk.Uint64ToBigEndian(spaceId))
}

// incrSpaceId increment the space sequence id
// NOTE: this function is called when creating a new space
func (k Keeper) incrSpaceSequence(ctx sdk.Context) {
	store := ctx.KVStore(k.storeKey)
	key := types.SpaceSequenceStoreKey()
	sequence := k.GetSpaceSequence(ctx)
	store.Set(key, sdk.Uint64ToBigEndian(sequence+1))
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
		return types.Space{}, sdkerrors.Wrapf(types.ErrInvalidSpaceId, "space (%d) does not exist", spaceId)
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

func (k Keeper) GetBlockHeaders(ctx sdk.Context) []types.BlockHeader {
	headers := make([]types.BlockHeader, 0)
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefixBlockHeader)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		ret := bytes.Split(iterator.Key(), types.Delimiter)
		if len(ret) != 2 {
			panic("fail to split layer2 block header key")
		}
		var BlockHeader types.BlockHeader
		spaceId, err := strconv.ParseUint(string(ret[0]), 10, 64)
		if err != nil {
			panic("fail to convert spaceId to uint64")
		}
		height, err := strconv.ParseUint(string(ret[1]), 10, 64)
		if err != nil {
			panic("fail to convert height to uint64")
		}

		BlockHeader.SpaceId = spaceId
		BlockHeader.Height = height
		BlockHeader.Header = string(iterator.Value())
		headers = append(headers, BlockHeader)
	}
	return headers
}

func (k Keeper) GetBlockHeader(ctx sdk.Context, spaceId, height uint64) (string, error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.BlockHeaderStoreKey(spaceId, height))
	if bz == nil {
		return "", sdkerrors.Wrapf(types.ErrBlockHeader, "block header does not exist at height (%d) in space (%d)", height, spaceId)
	}
	return string(bz), nil
}

func (k Keeper) HasBlockHeader(ctx sdk.Context, spaceId, height uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.BlockHeaderStoreKey(spaceId, height))
}

func (k Keeper) setBlockHeader(ctx sdk.Context, spaceId, blockHeight uint64, header string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.BlockHeaderStoreKey(spaceId, blockHeight), []byte(header))
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
