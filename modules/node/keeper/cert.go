package keeper

import (
	"github.com/tendermint/tendermint/crypto/algo"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/iritamod/modules/node/types"
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

	switch certType {
	case algo.SM2:
		RootCertKey = types.GetRootCertKey(algo.SM2)
	case algo.ED25519:
		RootCertKey = types.GetRootCertKey(algo.ED25519)
	default:
		return types.Certificate{}, false
	}

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

		switch cert.Key {
		case algo.SM2, "SM2":
			RootCertKey = types.GetRootCertKey(algo.SM2)
		case algo.ED25519, "ED25519":
			RootCertKey = types.GetRootCertKey(algo.ED25519)
		}

		bz = k.cdc.MustMarshal(&cert)
		store.Set(RootCertKey, bz)
	}
}
