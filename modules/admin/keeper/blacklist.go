package keeper

import (
	gogotypes "github.com/gogo/protobuf/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"gitlab.bianjie.ai/irita-pro/iritamod/modules/admin/types"
)

// BlockAccount blocks an account
func (k Keeper) BlockAccount(ctx sdk.Context, address sdk.AccAddress) error {
	if k.IsAdmin(ctx, address) {
		return sdkerrors.Wrap(types.ErrBlockAdminAccount, address.String())
	}
	if k.GetBlockAccount(ctx, address) {
		return sdkerrors.Wrap(types.ErrAlreadyBlockedAccount, address.String())
	}
	k.setBlackAccount(ctx, address)
	return nil
}

// UnblockAccount unblocks an account
func (k Keeper) UnblockAccount(ctx sdk.Context, address sdk.AccAddress) error {
	if !k.GetBlockAccount(ctx, address) {
		return sdkerrors.Wrap(types.ErrUnknownBlockedAccount, address.String())
	}
	k.deleteBlackAccount(ctx, address)
	return nil
}

func (k Keeper) setBlackAccount(ctx sdk.Context, address sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryBare(&gogotypes.BoolValue{Value: true})
	store.Set(types.GetBlackKey(address), bz)
}

// BlockAccount blocks an account
func (k Keeper) GetBlockAccount(ctx sdk.Context, address sdk.AccAddress) bool {
	store := ctx.KVStore(k.storeKey)

	value := store.Get(types.GetBlackKey(address))
	if value == nil {
		return false
	}

	var black gogotypes.BoolValue
	k.cdc.MustUnmarshalBinaryBare(value, &black)

	return true
}

func (k Keeper) deleteBlackAccount(ctx sdk.Context, address sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetBlackKey(address))
}

// GetAllBlackAccounts gets the set of all accounts with no limits, used durng genesis dump
func (k Keeper) GetAllBlackAccounts(ctx sdk.Context) (accounts []sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.BlackKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		accounts = append(accounts, iterator.Key()[1:])
	}

	return accounts
}
