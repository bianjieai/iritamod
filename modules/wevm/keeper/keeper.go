package keeper

import (
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/bianjieai/iritamod/modules/wevm/types"
)

type Keeper struct {
	storeKey sdk.StoreKey
	cdc      codec.Codec
}

func NewKeeper(cdc codec.Codec, storeKey sdk.StoreKey) Keeper {
	return Keeper{
		storeKey,
		cdc,
	}
}

// GetStore get the store by keeper.storeKey
func (k Keeper) GetStore(ctx sdk.Context) sdk.KVStore {
	store := ctx.KVStore(k.storeKey)
	return store
}

// IteratorContractDanyList iterator the contract addresses in ContractDanyList
func (k Keeper) IteratorContractDanyList(goCtx sdk.Context) ([]string, error) {
	gm := make([]string, 0)
	store := k.GetStore(goCtx)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.KeyPrefixContractDenyList))
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		split := strings.Split(string(key), "/")
		if len(split) < 2 {
			return nil, errors.Wrapf(types.ErrNotFound, "not found any contract address from ContractDanyList")
		}
		gm = append(gm, split[1])
	}
	return gm, nil
}
