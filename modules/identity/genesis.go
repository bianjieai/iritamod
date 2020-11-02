package identity

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis - store genesis parameters
func InitGenesis(ctx sdk.Context, k Keeper, data GenesisState) {
	if err := ValidateGenesis(data); err != nil {
		panic(err.Error())
	}

	for _, identity := range data.Identities {
		if err := k.SetIdentity(ctx, identity); err != nil {
			panic(err.Error())
		}
	}
}

// ExportGenesis - output genesis parameters
func ExportGenesis(ctx sdk.Context, k Keeper) *GenesisState {
	identities := make([]Identity, 0)

	k.IterateIdentities(
		ctx,
		func(identity Identity) bool {
			identities = append(identities, identity)
			return false
		},
	)

	return NewGenesisState(identities)
}
