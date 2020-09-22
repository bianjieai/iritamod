package keeper

import (
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/encoding"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) ApplyAndReturnValidatorSetUpdates(ctx sdk.Context) (updates []abci.ValidatorUpdate) {
	k.IterateUpdateValidators(
		ctx,
		func(index int64, pubkey string, power int64) bool {
			updates = append(updates, ABCIValidatorUpdate(
				sdk.MustGetPubKeyFromBech32(sdk.Bech32PubKeyTypeConsPub, pubkey),
				power,
			))
			k.DequeueValidatorsUpdate(ctx, pubkey)
			return false
		},
	)
	return
}

func ABCIValidatorUpdate(pubkey crypto.PubKey, power int64) abci.ValidatorUpdate {
	pk, err := encoding.PubKeyToProto(pubkey)
	if err != nil {
		panic(err)
	}

	return abci.ValidatorUpdate{
		PubKey: pk,
		Power:  power,
	}
}
