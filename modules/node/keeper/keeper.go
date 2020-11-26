package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/iritamod/modules/node/types"
)

// Keeper defines the node keeper
type Keeper struct {
	storeKey sdk.StoreKey
	cdc      codec.Marshaler

	validatorKeeper types.ValidatorKeeper
}

// NewKeeper creates a new node Keeper instance
func NewKeeper(
	cdc codec.Marshaler,
	key sdk.StoreKey,
	validatorKeeper types.ValidatorKeeper,
) Keeper {
	return Keeper{
		storeKey:        key,
		cdc:             cdc,
		validatorKeeper: validatorKeeper,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("modules/%s", types.ModuleName))
}
