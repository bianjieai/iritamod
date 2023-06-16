package keeper

import (
	"github.com/bianjieai/iritamod/modules/params/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Keeper define a slashing keeper
type Keeper struct {
	authKeeper types.AccountKeeper

	router *baseapp.MsgServiceRouter
}

// NewKeeper creates a slashing keeper
func NewKeeper(ak types.AccountKeeper, router *baseapp.MsgServiceRouter) Keeper {
	return Keeper{
		authKeeper: ak,
		router:     router,
	}
}

// UpdateParams updates the params
func (k Keeper) UpdateParams(ctx sdk.Context, messages []sdk.Msg) error {
	var (
		err    error
		events sdk.Events
	)

	cacheCtx, writeCache := ctx.CacheContext()
	for _, msg := range messages {
		handler := k.router.Handler(msg)
		if handler == nil {
			return sdkerrors.Wrap(types.ErrUnroutableUpdateParamsMsg, sdk.MsgTypeURL(msg))
		}

		var res *sdk.Result

		res, err = handler(cacheCtx, msg)
		if err != nil {
			break
		}

		events = append(events, res.GetEvents()...)
		ctx.EventManager().EmitEvents(events)
	}

	if err != nil {
		return err
	}
	writeCache()

	return nil
}
