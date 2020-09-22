package types

import (
	tmbytes "github.com/tendermint/tendermint/libs/bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
)

var (
	_ sdk.Msg = &MsgUnjailValidator{}
)

// NewMsgUnjailValidator creates a new MsgUnjailValidator instance.
func NewMsgUnjailValidator(id tmbytes.HexBytes, operator sdk.AccAddress) *MsgUnjailValidator {
	return &MsgUnjailValidator{
		Id:       id,
		Operator: operator,
	}
}

func (m MsgUnjailValidator) Route() string {
	return slashingtypes.RouterKey
}

func (m MsgUnjailValidator) Type() string {
	return "unjail_validator"
}

func (m MsgUnjailValidator) ValidateBasic() error {
	if m.Operator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "operator missing")
	}
	if len(m.Id) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "validator id cannot be blank")
	}
	return nil
}

func (m MsgUnjailValidator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

func (m MsgUnjailValidator) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Operator}
}
