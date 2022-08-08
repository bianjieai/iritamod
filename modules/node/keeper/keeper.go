package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/bianjieai/iritamod/modules/node/types"
	cautil "github.com/bianjieai/iritamod/utils/ca"
)

// keeper of the node store
type Keeper struct {
	cdc      codec.Codec
	storeKey sdk.StoreKey

	paramstore   paramtypes.Subspace
	hooks        staking.StakingHooks
	verifyCertFn func(sdk.Context, string) (cert cautil.Cert, err error)
}

func NewKeeper(cdc codec.Codec, storeKey sdk.StoreKey, ps paramtypes.Subspace) Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(ParamKeyTable())
	}

	k := Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		paramstore: ps,
	}
	return k.SetVerifyCertFn(k.VerifyCert)
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("modules/%s", types.ModuleName))
}

func (k Keeper) SetVerifyCertFn(fn func(sdk.Context, string) (cert cautil.Cert, err error)) Keeper {
	k.verifyCertFn = fn
	return k
}
