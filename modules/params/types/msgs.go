package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewMsgUpdateParams creates a new MsgUpdateParams instance.
func NewMsgUpdateParams(changes []ParamChange, operator sdk.AccAddress) *MsgUpdateParams {
	return &MsgUpdateParams{
		Changes:  changes,
		Operator: operator.String(),
	}
}

// Route implements Msg
func (m MsgUpdateParams) Route() string {
	return RouterKey
}

// Type implements Msg
func (m MsgUpdateParams) Type() string {
	return "update_params"
}

// ValidateBasic implements Msg
func (m MsgUpdateParams) ValidateBasic() error {
	if len(m.Operator) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "operator missing")
	}
	if _, err := sdk.AccAddressFromBech32(m.Operator); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid operator")
	}
	return ValidateChanges(m.Changes)
}

// GetSignBytes implements Msg
func (m MsgUpdateParams) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements Msg
func (m MsgUpdateParams) GetSigners() []sdk.AccAddress {
	singer, err := sdk.AccAddressFromBech32(m.Operator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{singer}
}

// ValidateChanges performs basic validation checks over a set of ParamChange. It
// returns an error if any ParamChange is invalid.
func ValidateChanges(changes []ParamChange) error {
	if len(changes) == 0 {
		return ErrEmptyChanges
	}

	for _, pc := range changes {
		if len(pc.Subspace) == 0 {
			return ErrEmptySubspace
		}
		if len(pc.Key) == 0 {
			return ErrEmptyKey
		}
		if len(pc.Value) == 0 {
			return ErrEmptyValue
		}
	}

	return nil
}
