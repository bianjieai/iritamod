package keeper

import (
	store "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	"github.com/bianjieai/iritamod/modules/upgrade/types"
)

type Keeper struct {
	uk upgradekeeper.Keeper
}

func NewKeeper(uk upgradekeeper.Keeper) Keeper {
	return Keeper{
		uk: uk,
	}
}

// ScheduleUpgrade schedules an upgrade based on the specified plan.
func (k Keeper) ScheduleUpgrade(ctx sdk.Context, msg *types.MsgUpgradeSoftware) error {
	p, has := k.uk.GetUpgradePlan(ctx)
	if has {
		return sdkerrors.Wrapf(
			sdkerrors.ErrInvalidRequest,
			"there is already an upgrade plan: %s. if you want to submit a new plan, please cancel the current plan",
			p.Name,
		)
	}

	plan := upgradetypes.Plan{
		Name:   msg.Name,
		Height: msg.Height,
		Info:   msg.Info,
	}
	return k.uk.ScheduleUpgrade(ctx, plan)
}

// ClearUpgradePlan clears currently schedule upgrade
func (k Keeper) ClearUpgradePlan(ctx sdk.Context) error {
	_, has := k.uk.GetUpgradePlan(ctx)
	if !has {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "there is currently no upgrade plan")
	}

	k.uk.ClearUpgradePlan(ctx)
	return nil
}

func (k Keeper) UpgradeKeeper() upgradekeeper.Keeper {
	return k.uk
}

// SetUpgradeHandler sets an UpgradeHandler for the upgrade specified by name. This handler will be called when the upgrade
// with this name is applied. In order for an upgrade with the given name to proceed, a handler for this upgrade
// must be set even if it is a no-op function.
func (k *Keeper) SetUpgradeHandler(name string, upgradeHandler upgradetypes.UpgradeHandler) {
	k.uk.SetUpgradeHandler(name, upgradeHandler)
}

// ReadUpgradeInfoFromDisk returns the name and height of the upgrade
// which is written to disk by the old binary when panic'ing
// if there's an error in reading the info,
// it assumes that the upgrade info is not available
func (k Keeper) ReadUpgradeInfoFromDisk() (store.UpgradeInfo, error) {
	return k.uk.ReadUpgradeInfoFromDisk()
}

// IsSkipHeight checks if the given height is part of skipUpgradeHeights
func (k Keeper) IsSkipHeight(height int64) bool {
	return k.uk.IsSkipHeight(height)
}
