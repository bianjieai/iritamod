// noalias
// DONTCOVER
package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingexported "github.com/cosmos/cosmos-sdk/x/staking/exported"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
)

// AccountKeeper expected account keeper
type ValidatorKeeper interface {
	ValidatorByID(ctx sdk.Context, id tmbytes.HexBytes) stakingexported.ValidatorI
}
