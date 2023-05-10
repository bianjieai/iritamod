package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"

	"github.com/bianjieai/iritamod/modules/layer2/types"
	perm "github.com/bianjieai/iritamod/modules/perm/keeper"
)

type Keeper struct {
	cdc codec.Codec
	storeKey storetypes.StoreKey
	perm perm.Keeper
	nft types.NFTKeeper
}

func NewKeeper(cdc codec.Codec, storeKey storetypes.StoreKey, perm perm.Keeper, nft types.NFTKeeper) Keeper {
	panic("implement me")
}

