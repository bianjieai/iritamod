package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	perm "github.com/bianjieai/iritamod/modules/perm/keeper"
	nft "github.com/bianjieai/iritamod/modules/layer2/expected_keeper"
)

type Keeper struct {
	cdc codec.Codec
	storeKey storetypes.StoreKey
	bank bank.Keeper
	perm perm.Keeper
	nft nft.NFTKeeper
}

func NewKeeper(cdc codec.Codec, storeKey storetypes.StoreKey, perm perm.Keeper, nft nft.NFTKeeper) Keeper {
	panic("implement me")
}

