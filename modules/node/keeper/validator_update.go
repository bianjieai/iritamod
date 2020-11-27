package keeper

import (
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/encoding"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) ApplyAndReturnValidatorSetUpdates(ctx sdk.Context) (updates []abci.ValidatorUpdate, err error) {
	k.IterateUpdateValidators(
		ctx,
		func(index int64, pubkey string, power int64) bool {
			pk := sdk.MustGetPubKeyFromBech32(sdk.Bech32PubKeyTypeConsPub, pubkey)
			intoTmPk, ok := pk.(cryptotypes.IntoTmPubKey)
			if !ok {
				panic("invalid public key type")
			}
			updates = append(updates, ABCIValidatorUpdate(
				intoTmPk.AsTmPubKey(),
				power,
			))
			k.DequeueValidatorsUpdate(ctx, pubkey)
			return false
		},
	)
	return updates, err
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
