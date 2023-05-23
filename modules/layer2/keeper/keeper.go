package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"

	"github.com/bianjieai/iritamod/modules/layer2/types"
)

type Keeper struct {
	cdc      codec.Codec
	storeKey storetypes.StoreKey
	perm     types.PermKeeper
	nft      types.NFTKeeper
}

func NewKeeper(cdc codec.Codec, storeKey storetypes.StoreKey, perm types.PermKeeper, nft types.NFTKeeper) Keeper {
	return Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		perm:     perm,
		nft:      nft,
	}
}

func (k Keeper) GetNFTKeeper() types.NFTKeeper {
	return k.nft
}

func (k Keeper) GetPermKeeper() types.PermKeeper {
	return k.perm
}
