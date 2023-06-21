package keeper

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/iritamod/modules/params/types"
)

// Keeper define a params keeper
type Keeper struct {
	authKeeper types.AccountKeeper

	router *types.ParamsRouter
}

// NewKeeper creates a params keeper
func NewKeeper(ak types.AccountKeeper, msr *baseapp.MsgServiceRouter, msgTypeURLs []string) Keeper {
	return Keeper{
		authKeeper: ak,
		router:     types.NewParamsRouter(msr, msgTypeURLs),
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
		var res *sdk.Result
		res, err = k.router.Execute(cacheCtx, msg)
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
