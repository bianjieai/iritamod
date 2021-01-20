package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgMint = "mint" // type for MsgMint
)

var (
	_ sdk.Msg = &MsgMint{}
)

// NewMsgMint creates a new MsgMint instance.
func NewMsgMint(amount uint64, operator sdk.AccAddress) *MsgMint {
	return &MsgMint{
		Amount:   amount,
		Operator: operator.String(),
	}
}

// Route implements Msg.
func (m MsgMint) Route() string {
	return RouterKey
}

// Type implements Msg.
func (m MsgMint) Type() string {
	return TypeMsgMint
}

// ValidateBasic implements Msg.
func (m MsgMint) ValidateBasic() error {
	if len(m.Operator) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "operator missing")
	}

	if _, err := sdk.AccAddressFromBech32(m.Operator); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid operator")
	}

	if m.Amount == 0 {
		return sdkerrors.Wrap(ErrInvalidAmount, "amount should be greater than 0")
	}

	return nil
}

// GetSignBytes implements Msg.
func (m MsgMint) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements Msg.
func (m MsgMint) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Operator)
	return []sdk.AccAddress{addr}
}
