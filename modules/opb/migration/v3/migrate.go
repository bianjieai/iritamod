package v3

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/iritamod/modules/opb/exported"
	"github.com/bianjieai/iritamod/modules/opb/types"
)

const (
	ModuleName = "opb"
)

var ParamsKey = []byte{0x01}

// Migrate migrates the node params from legacy x/params module to node module
func Migrate(
	ctx sdk.Context,
	store sdk.KVStore,
	legacySubspace exported.Subspace,
	cdc codec.BinaryCodec,
) error {
	var params types.Params
	legacySubspace.GetParamSet(ctx, &params)

	if err := params.Validate(); err != nil {
		return err
	}

	bz := cdc.MustMarshal(&params)
	store.Set(ParamsKey, bz)

	return nil
}
