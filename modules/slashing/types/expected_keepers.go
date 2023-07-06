// noalias
// DONTCOVER
package types

import (
	"cosmossdk.io/math"
	ctmbytes "github.com/cometbft/cometbft/libs/bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// NodeKeeper defines the expected node keeper
type NodeKeeper interface {
	ValidatorByID(ctx sdk.Context, id ctmbytes.HexBytes) stakingtypes.ValidatorI

	// iterate through validators by operator address, execute func for each validator
	IterateValidators(sdk.Context,
		func(index int64, validator stakingtypes.ValidatorI) (stop bool))

	Validator(sdk.Context, sdk.ValAddress) stakingtypes.ValidatorI            // get a particular validator by operator address
	ValidatorByConsAddr(sdk.Context, sdk.ConsAddress) stakingtypes.ValidatorI // get a particular validator by consensus address

	// slash the validator and delegators of the validator, specifying offence height, offence power, and slash fraction
	Slash(sdk.Context, sdk.ConsAddress, int64, int64, sdk.Dec) math.Int
	SlashWithInfractionReason(sdk.Context, sdk.ConsAddress, int64, int64, sdk.Dec, stakingtypes.Infraction) math.Int
	Jail(sdk.Context, sdk.ConsAddress)   // jail a validator
	Unjail(sdk.Context, sdk.ConsAddress) // unjail a validator

	// Delegation allows for getting a particular delegation for a given validator
	// and delegator outside the scope of the staking module.
	Delegation(sdk.Context, sdk.AccAddress, sdk.ValAddress) stakingtypes.DelegationI
	GetAllValidators(ctx sdk.Context) (validators []stakingtypes.Validator)

	// MaxValidators returns the maximum amount of bonded validators
	MaxValidators(sdk.Context) uint32

	// IsValidatorJailed returns if the validator is jailed.
	IsValidatorJailed(ctx sdk.Context, addr sdk.ConsAddress) bool
}
