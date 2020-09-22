package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewMsgUpdateParams creates a new MsgUpdateParams instance.
func NewMsgUpdateParams(changes []ParamChange, operator sdk.AccAddress) *MsgUpdateParams {
	return &MsgUpdateParams{
		Changes:  changes,
		Operator: operator,
	}
}

func (m MsgUpdateParams) Route() string {
	return RouterKey
}

func (m MsgUpdateParams) Type() string {
	return "update_params"
}

func (m MsgUpdateParams) ValidateBasic() error {
	if m.Operator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "operator missing")
	}
	return ValidateChanges(m.Changes)
}

func (m MsgUpdateParams) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

func (m MsgUpdateParams) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Operator}
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
