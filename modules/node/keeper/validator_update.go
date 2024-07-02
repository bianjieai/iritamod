package keeper

import (
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	"github.com/cometbft/cometbft/crypto"
	"github.com/cometbft/cometbft/crypto/encoding"

	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) ApplyAndReturnValidatorSetUpdates(ctx sdk.Context) (updates []abci.ValidatorUpdate, err error) {
	k.IterateUpdateValidators(
		ctx,
		func(index int64, pubkey string, power int64) bool {
			var pk cryptotypes.PubKey
			bz, err := sdk.GetFromBech32(pubkey, sdk.GetConfig().GetBech32ConsensusPubPrefix())
			pk, err = legacy.PubKeyFromBytes(bz)
			if err != nil {
				panic(err)
			}
			tmPubkey, err := cryptocodec.ToTmPubKeyInterface(pk)
			if err != nil {
				panic(err.Error())
			}
			updates = append(
				updates,
				ABCIValidatorUpdate(tmPubkey, power),
			)
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
