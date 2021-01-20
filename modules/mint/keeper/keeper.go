package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/bianjieai/iritamod/modules/mint/types"
)

// Keeper defines the mint keeper
type Keeper struct {
	cdc      codec.Marshaler
	storeKey sdk.StoreKey

	tokenKeeper types.TokenKeeper
	paramSpace  paramstypes.Subspace
}

// NewKeeper creates a new Keeper instance
func NewKeeper(
	cdc codec.Marshaler,
	storeKey sdk.StoreKey,
	tokenKeeper types.TokenKeeper,
	paramSpace paramstypes.Subspace,
) Keeper {
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(ParamKeyTable())
	}

	return Keeper{
		cdc:         cdc,
		storeKey:    storeKey,
		tokenKeeper: tokenKeeper,
		paramSpace:  paramSpace,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("iritamod/%s", types.ModuleName))
}

// Mint mints the base native token by the specified amount
func (k Keeper) Mint(ctx sdk.Context, amount uint64, recipient sdk.AccAddress) error {
	mintDenom := k.MintDenom(ctx)

	return k.tokenKeeper.MintToken(ctx, mintDenom, amount, recipient)
}
