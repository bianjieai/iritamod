package keeper

import (
	"fmt"

	gogotypes "github.com/gogo/protobuf/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/bianjieai/iritamod/modules/perm/types"
)

// keeper of the perm store
type Keeper struct {
	cdc      codec.Codec
	storeKey sdk.StoreKey

	AuthMap map[string]types.Auth
}

func NewKeeper(cdc codec.Codec, storeKey sdk.StoreKey) Keeper {
	return Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		AuthMap:  make(map[string]types.Auth),
	}
}

// RegisterMsgAuth registers the auth to send the msg.
// Each role gets the access control
func (k Keeper) RegisterMsgAuth(msg sdk.Msg, roles ...types.Role) {
	if _, ok := k.AuthMap[sdk.MsgTypeURL(msg)]; ok {
		panic(fmt.Sprintf("msg type or module name %s has already been initialized", sdk.MsgTypeURL(msg)))
	}
	auth := types.AuthDefault
	for _, r := range roles {
		auth = auth | r.Auth()
	}
	k.AuthMap[sdk.MsgTypeURL(msg)] = auth
}

// RegisterModuleAuth registers the auth to send the module related msgs.
// Each role gets the access control
func (k *Keeper) RegisterModuleAuth(module string, roles ...types.Role) {
	if _, ok := k.AuthMap[module]; ok {
		panic(fmt.Sprintf("msg type or module name %s has already been initialized", module))
	}
	auth := types.AuthDefault
	for _, r := range roles {
		auth = auth | r.Auth()
	}
	k.AuthMap[module] = auth
}

// Authorize assigns the specified roles to an address
func (k *Keeper) Authorize(ctx sdk.Context, address, operator sdk.AccAddress, rs ...types.Role) error {
	if k.IsRootAdmin(ctx, address) {
		return types.ErrOperateRootAdmin
	}
	if k.IsPermAdmin(ctx, address) &&
		(!k.IsRootAdmin(ctx, operator) || operator.Equals(address)) {
		return types.ErrOperatePermAdmin
	}
	auth := k.GetAuth(ctx, address)
	for _, r := range rs {
		if r == types.RoleRootAdmin {
			return types.ErrAddRootAdmin
		}
		if r == types.RolePermAdmin && !k.IsRootAdmin(ctx, operator) {
			return sdkerrors.Wrap(types.ErrUnauthorizedOperation, "can not add permission admin role")
		}
		auth = auth | r.Auth()
	}
	k.SetAuth(ctx, address, auth)
	return nil
}

// Unauthorize unassigns the specified roles from an address
func (k Keeper) Unauthorize(ctx sdk.Context, address, operator sdk.AccAddress, roles ...types.Role) error {
	if k.IsRootAdmin(ctx, address) {
		return types.ErrOperateRootAdmin
	}
	if k.IsPermAdmin(ctx, address) &&
		(!k.IsRootAdmin(ctx, operator) || operator.Equals(address)) {
		return types.ErrOperatePermAdmin
	}

	auth := k.GetAuth(ctx, address)
	for _, r := range roles {
		if r == types.RoleRootAdmin {
			return types.ErrRemoveRootAdmin
		}
		if !auth.Access(r.Auth()) {
			return sdkerrors.Wrapf(types.ErrRemoveUnknownRole, "%s", r)
		}
		auth = auth & (auth ^ r.Auth())
	}
	if auth == types.AuthDefault {
		k.DeleteAuth(ctx, address)
	} else {
		k.SetAuth(ctx, address, auth)
	}
	return nil
}

// GetRoles gets the role set for all accounts
func (k Keeper) GetRoles(ctx sdk.Context) (roleAccounts []types.RoleAccount) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.AuthKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var role gogotypes.Int32Value
		k.cdc.MustUnmarshal(iterator.Value(), &role)

		account := sdk.AccAddress(iterator.Key()[len(types.AuthKey):])
		roleAccounts = append(roleAccounts, types.RoleAccount{
			Address: account.String(),
			Roles:   types.Auth(role.Value).Roles(),
		})
	}

	return roleAccounts
}

// Access checks the signer auth
func (k Keeper) Access(ctx sdk.Context, signer sdk.AccAddress, auth types.Auth) error {
	signerAuth := k.GetAuth(ctx, signer)
	if !auth.Access(signerAuth) {
		return sdkerrors.Wrapf(
			types.ErrUnauthorizedOperation,
			"Required roles: %s; sender roles: %s. ",
			auth.Roles(), signerAuth.Roles(),
		)
	}
	return nil
}
