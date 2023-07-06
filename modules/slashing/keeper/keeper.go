package keeper

import (
	"encoding/hex"
	"fmt"

	"github.com/cometbft/cometbft/crypto"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"

	"github.com/bianjieai/iritamod/modules/slashing/types"
)

// Keeper define a slashing keeper
type Keeper struct {
	slashingkeeper.Keeper
	nodeKeeper types.NodeKeeper
}

// NewKeeper creates a slashing keeper
func NewKeeper(
	cdc codec.BinaryCodec,
	legacyAmino *codec.LegacyAmino,
	key storetypes.StoreKey,
	nodeKeeper types.NodeKeeper,
	authority string,
) Keeper {
	return Keeper{
		Keeper:     slashingkeeper.NewKeeper(cdc, legacyAmino, key, nodeKeeper, authority),
		nodeKeeper: nodeKeeper,
	}
}

// HandleValidatorSignature handles a validator signature, must be called once per validator per block.
// Block all subsequent logic if this validator has been removed.
func (k Keeper) HandleValidatorSignature(
	ctx sdk.Context,
	addr crypto.Address,
	power int64,
	signed bool,
) {
	logger := k.Logger(ctx)

	// fetch the validator public key
	consAddr := sdk.ConsAddress(addr)
	if _, err := k.GetPubkey(ctx, addr); err != nil {
		logger.Info(fmt.Sprintf("Validator consensus-address %s not found", consAddr))
		return
	}

	// fetch signing info
	if _, found := k.GetValidatorSigningInfo(ctx, consAddr); !found {
		logger.Info(fmt.Sprintf("Expected signing info for validator %s but not found", consAddr))
		return
	}

	k.Keeper.HandleValidatorSignature(ctx, addr, power, signed)
}

// HandleUnjail handles ths unjail msg
func (k Keeper) HandleUnjail(ctx sdk.Context, msg types.MsgUnjailValidator) error {
	id, err := hex.DecodeString(msg.Id)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid validator id : %s", msg.Id)
	}
	validator := k.nodeKeeper.ValidatorByID(ctx, id)
	if validator == nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "unknown validator: %s", msg.Id)
	}
	return k.Unjail(ctx, validator.GetOperator())
}
