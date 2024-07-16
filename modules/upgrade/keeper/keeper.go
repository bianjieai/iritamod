package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	xp "github.com/cosmos/cosmos-sdk/x/upgrade/exported"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	"iritamod.bianjie.ai/modules/upgrade/types"
)

type Keeper struct {
	*upgradekeeper.Keeper
}

func NewKeeper(skipUpgradeHeights map[int64]bool, storeKey storetypes.StoreKey, cdc codec.BinaryCodec, homePath string, vs xp.ProtocolVersionSetter, authority string) Keeper {
	return Keeper{
		upgradekeeper.NewKeeper(skipUpgradeHeights, storeKey, cdc, homePath, vs, authority),
	}
}

// ScheduleUpgrade schedules an upgrade based on the specified plan.
func (k Keeper) ScheduleUpgrade(ctx sdk.Context, msg *types.MsgUpgradeSoftware) error {
	p, has := k.GetUpgradePlan(ctx)
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
	return k.Keeper.ScheduleUpgrade(ctx, plan)
}

// ClearUpgradePlan clears currently schedule upgrade
func (k Keeper) ClearUpgradePlan(ctx sdk.Context) error {
	_, has := k.GetUpgradePlan(ctx)
	if !has {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "there is currently no upgrade plan")
	}

	k.Keeper.ClearUpgradePlan(ctx)
	return nil
}

func (k Keeper) UpgradeKeeper() *upgradekeeper.Keeper {
	return k.Keeper
}

// SetUpgradeHandler sets an UpgradeHandler for the upgrade specified by name. This handler will be called when the upgrade
// with this name is applied. In order for an upgrade with the given name to proceed, a handler for this upgrade
// must be set even if it is a no-op function.
//func (k *Keeper) SetUpgradeHandler(name string, upgradeHandler upgradetypes.UpgradeHandler) {
//	k.Keeper.SetUpgradeHandler(name, upgradeHandler)
//}

// ReadUpgradeInfoFromDisk returns the name and height of the upgrade
// which is written to disk by the old binary when panic'ing
// if there's an error in reading the info,
// it assumes that the upgrade info is not available
//func (k Keeper) ReadUpgradeInfoFromDisk() (upgradetypes.Plan, error) {
//	return k.Keeper.ReadUpgradeInfoFromDisk()
//}

// IsSkipHeight checks if the given height is part of skipUpgradeHeights
//func (k Keeper) IsSkipHeight(height int64) bool {
//	return k.Keeper.IsSkipHeight(height)
//}
