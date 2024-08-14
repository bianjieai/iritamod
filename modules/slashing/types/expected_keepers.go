// noalias
// DONTCOVER
package types

import (
	tmbytes "github.com/cometbft/cometbft/libs/bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingexported "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// NodeKeeper defines the expected node keeper
type NodeKeeper interface {
	ValidatorByID(ctx sdk.Context, id tmbytes.HexBytes) stakingexported.ValidatorI
}
