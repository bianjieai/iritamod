package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/iritamod/modules/identity/types"
)

// Keeper defines the identity keeper
type Keeper struct {
	storeKey sdk.StoreKey
	cdc      codec.Marshaler
}

// NewKeeper creates a new identity Keeper instance
func NewKeeper(cdc codec.Marshaler, key sdk.StoreKey) Keeper {
	return Keeper{
		storeKey: key,
		cdc:      cdc,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("modules/%s", types.ModuleName))
}
