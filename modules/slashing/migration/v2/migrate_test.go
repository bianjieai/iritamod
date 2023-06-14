package v2_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"

	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"

	v2 "github.com/bianjieai/iritamod/modules/slashing/migration/v2"
	"github.com/bianjieai/iritamod/simapp"
)

func TestMigrate(t *testing.T) {
	app := simapp.Setup(t, false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	legacySubspace := app.GetSubspace(slashingtypes.ModuleName)

	params := slashingtypes.DefaultParams()
	legacySubspace.SetParamSet(ctx, &params)

	err := v2.Migrate(
		ctx,
		app.SlashingKeeper,
		legacySubspace,
	)
	require.NoError(t, err)

	expParams := app.SlashingKeeper.GetParams(ctx)
	require.Equal(t, expParams, params, "v1.Migrate failed")

}
