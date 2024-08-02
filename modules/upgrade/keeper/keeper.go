package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	xp "github.com/cosmos/cosmos-sdk/x/upgrade/exported"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
)

// Keeper of the upgrade module
type Keeper struct {
	*upgradekeeper.Keeper
}

// NewKeeper creates a new upgrade Keeper instance
func NewKeeper(
	skipUpgradeHeights map[int64]bool, 
	storeKey storetypes.StoreKey, 
	cdc codec.BinaryCodec, 
	homePath string, 
	vs xp.ProtocolVersionSetter, 
	authority string,
	) *Keeper {
	return &Keeper{
		upgradekeeper.NewKeeper(skipUpgradeHeights, storeKey, cdc, homePath, vs, authority),
	}
}
