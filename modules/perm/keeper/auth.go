package keeper

import (
	gogotypes "github.com/gogo/protobuf/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/iritamod/modules/perm/types"
)

// SetAuth sets the auth for an address
func (k Keeper) SetAuth(ctx sdk.Context, address sdk.AccAddress, auth types.Auth) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.Int32Value{Value: int32(auth)})
	store.Set(types.GetAuthKey(address), bz)
}

// DeleteAuth deletes the auth for an address
func (k Keeper) DeleteAuth(ctx sdk.Context, address sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetAuthKey(address))
}

// GetAuth gets the auth for an address
func (k Keeper) GetAuth(ctx sdk.Context, address sdk.AccAddress) types.Auth {
	store := ctx.KVStore(k.storeKey)

	value := store.Get(types.GetAuthKey(address))
	if value == nil {
		return 0
	}

	var role gogotypes.Int32Value
	k.cdc.MustUnmarshal(value, &role)

	return types.Auth(role.Value)
}

func (k Keeper) IsRootAdmin(ctx sdk.Context, address sdk.AccAddress) bool {
	auth := k.GetAuth(ctx, address)
	return (auth & types.RoleRootAdmin.Auth()) > 0
}

func (k Keeper) IsPermAdmin(ctx sdk.Context, address sdk.AccAddress) bool {
	auth := k.GetAuth(ctx, address)
	return (auth & types.RolePermAdmin.Auth()) > 0
}

func (k Keeper) IsBaseM1Admin(ctx sdk.Context, address sdk.AccAddress) bool {
	auth := k.GetAuth(ctx, address)
	return (auth & types.RoleBaseM1Admin.Auth()) > 0
}

func (k Keeper) IsAdmin(ctx sdk.Context, address sdk.AccAddress) bool {
	auth := k.GetAuth(ctx, address)
	return (auth&types.RoleRootAdmin.Auth()) > 0 ||
		(auth&types.RolePermAdmin.Auth()) > 0 ||
		(auth&types.RoleBlacklistAdmin.Auth()) > 0 ||
		(auth&types.RoleNodeAdmin.Auth()) > 0 ||
		(auth&types.RoleParamAdmin.Auth()) > 0
}
