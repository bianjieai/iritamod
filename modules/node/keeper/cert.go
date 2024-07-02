package keeper

import (
	gogotypes "github.com/gogo/protobuf/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/iritamod/modules/node/types"
)

func (k *Keeper) GetRootCert(ctx sdk.Context) (cert string, found bool) {
	store := ctx.KVStore(k.storeKey)

	value := store.Get(types.RootCertKey)
	if value == nil {
		return "", false
	}

	var certStr gogotypes.StringValue
	k.cdc.MustUnmarshal(value, &certStr)
	return certStr.Value, true
}

func (k *Keeper) SetRootCert(ctx sdk.Context, cert string) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.StringValue{Value: cert})
	store.Set(types.RootCertKey, bz)
}
