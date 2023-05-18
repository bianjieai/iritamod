package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"

	nft "github.com/bianjieai/iritamod/modules/layer2/expected_keeper"
	perm "github.com/bianjieai/iritamod/modules/perm/keeper"
)

type Keeper struct {
	cdc      codec.Codec
	storeKey storetypes.StoreKey
	perm     perm.Keeper
	nft      nft.NFTKeeper
}

func NewKeeper(cdc codec.Codec, storeKey storetypes.StoreKey, perm perm.Keeper, nft nft.NFTKeeper) Keeper {
	return Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		perm:     perm,
		nft:      nft,
	}
}

func (k Keeper) GetNFTKeeper() nft.NFTKeeper {
	return k.nft
}
