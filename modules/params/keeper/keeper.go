package keeper

import (
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"

	"github.com/bianjieai/iritamod/modules/params/types"
)

// Keeper define a slashing keeper
type Keeper struct {
	paramskeeper.Keeper
}

// NewKeeper creates a slashing keeper
func NewKeeper(keeper paramskeeper.Keeper) Keeper {
	return Keeper{
		keeper,
	}
}

// HandleValidatorSignature handles a validator signature, must be called once per validator per block.
// Block all subsequent logic if this validator has been removed.
func (k Keeper) UpdateParams(ctx sdk.Context, params []types.ParamChange) ([]sdk.Attribute, error) {
	var changeEvents []sdk.Attribute
	for _, c := range params {
		ss, ok := k.GetSubspace(c.Subspace)
		if !ok {
			return nil, sdkerrors.Wrap(types.ErrUnknownSubspace, c.Subspace)
		}

		k.Logger(ctx).Info(
			fmt.Sprintf("attempt to set new parameter value; key: %s, value: %s", c.Key, c.Value),
		)

		if !ss.Has(ctx, []byte(c.Key)) {
			return nil, sdkerrors.Wrapf(types.ErrUnknownKey, c.Key)
		}

		if err := ss.Update(ctx, []byte(c.Key), []byte(c.Value)); err != nil {
			return nil, sdkerrors.Wrapf(types.ErrSettingParameter, "key: %s, value: %s, err: %s", c.Key, c.Value, err.Error())
		}

		changeEvents = append(changeEvents, sdk.NewAttribute(types.AttributeKeyParamKey, c.Key))
	}
	return changeEvents, nil
}
