package v2_test

// import (
// 	"testing"

// 	"github.com/stretchr/testify/require"

// 	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"

// 	v2 "github.com/bianjieai/iritamod/modules/node/migration/v2"
// 	nodetypes "github.com/bianjieai/iritamod/modules/node/types"
// 	"github.com/bianjieai/iritamod/simapp"
// )

// func TestMigrate(t *testing.T) {
// 	app := simapp.Setup(t, false)
// 	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

// 	legacySubSpace := app.GetSubspace(nodetypes.ModuleName)
// 	params := nodetypes.DefaultParams()
// 	legacySubSpace.SetParamSet(ctx, &params)

// 	err := v2.Migrate(ctx, app.NodeKeeper, legacySubSpace)
// 	require.NoError(t, err)

// 	expParams := app.NodeKeeper.GetModuleParams(ctx)
// 	require.Equal(t, params, expParams)
// }
