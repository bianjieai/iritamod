package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/iritamod/modules/node/exported"
	"github.com/bianjieai/iritamod/modules/node/migration/v1"
)

type Migrator struct {
	keeper         Keeper
	legacySubspace exported.Subspace
}

func NewMigrator(keeper Keeper, legacySubspace exported.Subspace) Migrator {
	return Migrator{keeper, legacySubspace}
}

func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	return v1.Migrate(ctx, m.keeper, m.legacySubspace)
}