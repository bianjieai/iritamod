package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"iritamod.bianjie.ai/modules/identity/types"
)

// Keeper defines the identity keeper
type Keeper struct {
	storeKey storetypes.StoreKey
	cdc      codec.Codec
}

// NewKeeper creates a new identity Keeper instance
func NewKeeper(cdc codec.Codec, key storetypes.StoreKey) Keeper {
	return Keeper{
		storeKey: key,
		cdc:      cdc,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("modules/%s", types.ModuleName))
}
