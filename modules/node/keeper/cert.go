package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/iritamod/modules/node/types"
	cautil "github.com/bianjieai/iritamod/utils/ca"
)

func (k *Keeper) GetRootCert(ctx sdk.Context) (certs []types.Certificate, found bool) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.GetRootCertKey(""))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var certStr types.Certificate
		k.cdc.MustUnmarshal(iterator.Value(), &certStr)
		certs = append(certs, certStr)
	}
	return certs, true
}

func (k *Keeper) GetRootCertByType(ctx sdk.Context, certType string) (cert types.Certificate, found bool) {
	store := ctx.KVStore(k.storeKey)
	var RootCertKey []byte

	if _, err := cautil.IsSupportedAlgorithms(certType); err != nil {
		return types.Certificate{}, false
	}
	RootCertKey = types.GetRootCertKey(certType)

	value := store.Get(RootCertKey)
	if value == nil {
		return types.Certificate{}, false
	}

	var certStr types.Certificate
	k.cdc.MustUnmarshal(value, &certStr)
	return certStr, true
}

func (k *Keeper) SetRootCert(ctx sdk.Context, certs []types.Certificate) {
	store := ctx.KVStore(k.storeKey)

	for _, cert := range certs {
		var RootCertKey, bz []byte

		if _, err := cautil.IsSupportedAlgorithms(cert.Key); err != nil {
			panic(err)
		}
		RootCertKey = types.GetRootCertKey(cert.Key)

		bz = k.cdc.MustMarshal(&cert)
		store.Set(RootCertKey, bz)
	}
}
