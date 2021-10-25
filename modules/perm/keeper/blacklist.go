package keeper

import (
	gogotypes "github.com/gogo/protobuf/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/bianjieai/iritamod/modules/perm/types"
)

// BlockAccount blocks an account
func (k Keeper) Block(ctx sdk.Context, address sdk.AccAddress) error {
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
func (k Keeper) Unblock(ctx sdk.Context, address sdk.AccAddress) error {
	if !k.GetBlockAccount(ctx, address) {
		return sdkerrors.Wrap(types.ErrUnknownBlockedAccount, address.String())
	}
	k.deleteBlackAccount(ctx, address)
	return nil
}

func (k Keeper) setBlackAccount(ctx sdk.Context, address sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.BoolValue{Value: true})
	store.Set(types.GetBlackKey(address), bz)
}

// GetBlockAccount return an account blocked
func (k Keeper) GetBlockAccount(ctx sdk.Context, address sdk.AccAddress) bool {
	store := ctx.KVStore(k.storeKey)

	value := store.Get(types.GetBlackKey(address))
	if value == nil {
		return false
	}

	var black gogotypes.BoolValue
	k.cdc.MustUnmarshal(value, &black)

	return true
}

func (k Keeper) deleteBlackAccount(ctx sdk.Context, address sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetBlackKey(address))
}

// GetAllBlackAccounts gets the set of all accounts with no limits, used durng genesis dump
func (k Keeper) GetAllBlackAccounts(ctx sdk.Context) (accounts []string) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.BlackKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		account := sdk.AccAddress(iterator.Key()[1:])
		accounts = append(accounts, account.String())
	}

	return accounts
}
