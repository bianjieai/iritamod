package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"

	"github.com/bianjieai/iritamod/modules/side-chain/types"
)

type Keeper struct {
	cdc      codec.Codec
	storeKey storetypes.StoreKey
	acc      types.AccountKeeper
}

func NewKeeper(cdc codec.Codec, storeKey storetypes.StoreKey, acc types.AccountKeeper) Keeper {
	return Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		acc:      acc,
	}
}

func (k Keeper) GetAccKeeper() types.AccountKeeper {
	return k.acc
}
