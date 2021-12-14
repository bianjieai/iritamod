package keeper

import (
	"github.com/ethereum/go-ethereum/common"
	gogotypes "github.com/gogo/protobuf/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/bianjieai/iritamod/modules/perm/types"
)

// Block blocks an account
func (k Keeper) Block(ctx sdk.Context, address sdk.AccAddress) error {
	if k.IsAdmin(ctx, address) {
		return sdkerrors.Wrap(types.ErrBlockAdminAccount, address.String())
	}
	if k.GetBlockAccount(ctx, address) {
		return sdkerrors.Wrap(types.ErrAlreadyBlockedAccount, address.String())
	}
	k.setBlockAccount(ctx, address)
	return nil
}

// Unblock unblocks an account
func (k Keeper) Unblock(ctx sdk.Context, address sdk.AccAddress) error {
	if !k.GetBlockAccount(ctx, address) {
		return sdkerrors.Wrap(types.ErrUnknownBlockedAccount, address.String())
	}
	k.deleteBlockAccount(ctx, address)
	return nil
}

func (k Keeper) setBlockAccount(ctx sdk.Context, address sdk.AccAddress) {
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

func (k Keeper) deleteBlockAccount(ctx sdk.Context, address sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetBlackKey(address))
}

// GetAllBlockAccounts gets the set of all accounts with no limits, used durng genesis dump
func (k Keeper) GetAllBlockAccounts(ctx sdk.Context) (accounts []string) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.BlackKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		account := sdk.AccAddress(iterator.Key()[1:])
		accounts = append(accounts, account.String())
	}

	return accounts
}

// BlockContract blocks a contract
func (k Keeper) BlockContract(ctx sdk.Context, contractAddress string) error {
	contractAddr := common.HexToAddress(contractAddress)
	if k.GetBlockContract(ctx, contractAddr) {
		return sdkerrors.Wrap(types.ErrAlreadyBlockedAccount, contractAddr.String())
	}
	k.setContractDenyList(ctx, contractAddr)
	return nil
}

// UnblockContract unblocks a contract
func (k Keeper) UnblockContract(ctx sdk.Context, contractAddress string) error {
	address := common.HexToAddress(contractAddress)
	if !k.GetBlockContract(ctx, address) {
		return sdkerrors.Wrap(types.ErrUnknownBlockedAccount, address.String())
	}
	k.deleteContractDenyList(ctx, address)
	return nil
}

// GetContractDenyList gets the set of all contract with no limits, used durng genesis dump
func (k Keeper) GetContractDenyList(ctx sdk.Context) (accounts []string) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.ContractDenyListKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		account := common.BytesToAddress(iterator.Key()[1:])
		accounts = append(accounts, account.String())
	}

	return accounts
}

// GetBlockContract return a contract blocked
func (k Keeper) GetBlockContract(ctx sdk.Context, address common.Address) bool {
	store := ctx.KVStore(k.storeKey)

	value := store.Get(types.GetContractDenyListKey(address))
	if value == nil {
		return false
	}

	var black gogotypes.BoolValue
	k.cdc.MustUnmarshal(value, &black)

	return true
}

func (k Keeper) deleteContractDenyList(ctx sdk.Context, address common.Address) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetContractDenyListKey(address))
}

func (k Keeper) setContractDenyList(ctx sdk.Context, address common.Address) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.BoolValue{Value: true})
	store.Set(types.GetContractDenyListKey(address), bz)
}
