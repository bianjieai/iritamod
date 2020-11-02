package types

import (
	"strings"

	tmbytes "github.com/tendermint/tendermint/libs/bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = &MsgCreateValidator{}
	_ sdk.Msg = &MsgUpdateValidator{}
	_ sdk.Msg = &MsgRemoveValidator{}
)

// NewMsgCreateValidator creates a new MsgCreateValidator instance.
func NewMsgCreateValidator(
	name, description string, cert string, power int64, operator sdk.AccAddress,
) *MsgCreateValidator {
	return &MsgCreateValidator{
		Name:        name,
		Certificate: cert,
		Power:       power,
		Description: description,
		Operator:    operator.String(),
	}
}

// Route implement sdk.Msg
func (m MsgCreateValidator) Route() string {
	return RouterKey
}

// Type implement sdk.Msg
func (m MsgCreateValidator) Type() string {
	return "create_validator"
}

// ValidateBasic implement sdk.Msg
func (m MsgCreateValidator) ValidateBasic() error {
	if len(m.Operator) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "operator missing")
	}

	if _, err := sdk.AccAddressFromBech32(m.Operator); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid operator")
	}

	name := strings.TrimSpace(m.Name)
	if len(name) == 0 || DoNotModifyDesc == name {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "validator name cannot be blank")
	}

	if len(m.Certificate) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "certificate missing")
	}
	if m.Power <= 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "power must be positive")
	}
	return nil
}

// GetSignBytes implement sdk.Msg
func (m MsgCreateValidator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

// GetSigners implement sdk.Msg
func (m MsgCreateValidator) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(m.Operator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

// NewMsgUpdateValidator creates a new MsgUpdateValidator instance.
func NewMsgUpdateValidator(
	id tmbytes.HexBytes, name, description string, cert string, power int64, operator sdk.AccAddress,
) *MsgUpdateValidator {
	return &MsgUpdateValidator{
		Id:          id.String(),
		Name:        name,
		Certificate: cert,
		Power:       power,
		Description: description,
		Operator:    operator.String(),
	}
}

// Route implement sdk.Msg
func (m MsgUpdateValidator) Route() string {
	return RouterKey
}

// Type implement sdk.Msg
func (m MsgUpdateValidator) Type() string {
	return "update_validator"
}

// ValidateBasic implement sdk.Msg
func (m MsgUpdateValidator) ValidateBasic() error {
	if len(m.Operator) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "operator missing")
	}

	if _, err := sdk.AccAddressFromBech32(m.Operator); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid operator")
	}

	if len(m.Id) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "validator id cannot be blank")
	}

	if m.Power < 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "power can not be negative")
	}
	return nil
}

// GetSignBytes implement sdk.Msg
func (m MsgUpdateValidator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

// GetSigners implement sdk.Msg
func (m MsgUpdateValidator) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(m.Operator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

// NewMsgRemoveValidator creates a new MsgRemoveValidator instance.
func NewMsgRemoveValidator(id tmbytes.HexBytes, operator sdk.AccAddress) *MsgRemoveValidator {
	return &MsgRemoveValidator{
		Id:       id.String(),
		Operator: operator.String(),
	}
}

// Route implement sdk.Msg
func (m MsgRemoveValidator) Route() string {
	return RouterKey
}

// Type implement sdk.Msg
func (m MsgRemoveValidator) Type() string {
	return "remove_validator"
}

// ValidateBasic implement sdk.Msg
func (m MsgRemoveValidator) ValidateBasic() error {
	if len(m.Operator) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "operator missing")
	}

	if _, err := sdk.AccAddressFromBech32(m.Operator); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid operator")
	}

	if len(m.Id) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "validator id cannot be blank")
	}
	return nil
}

// GetSignBytes implement sdk.Msg
func (m MsgRemoveValidator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

// GetSigners implement sdk.Msg
func (m MsgRemoveValidator) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(m.Operator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}
