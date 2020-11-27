package types

import (
	"strings"

	tmbytes "github.com/tendermint/tendermint/libs/bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Node message types and params
const (
	TypeMsgCreateValidator = "create_validator" // type for MsgCreateValidator
	TypeMsgUpdateValidator = "update_validator" // type for MsgUpdateValidator
	TypeMsgRemoveValidator = "remove_validator" // type for MsgRemoveValidator
	TypeMsgAddNode         = "add_node"         // type for MsgAddNode
	TypeMsgRemoveNode      = "remove_node"      // type for MsgRemoveNode
)

var (
	_ sdk.Msg = &MsgCreateValidator{}
	_ sdk.Msg = &MsgUpdateValidator{}
	_ sdk.Msg = &MsgRemoveValidator{}
	_ sdk.Msg = &MsgAddNode{}
	_ sdk.Msg = &MsgRemoveNode{}
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
	return TypeMsgCreateValidator
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
	return TypeMsgUpdateValidator
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
	return TypeMsgRemoveValidator
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

// NewMsgAddNode creates a new MsgAddNode instance
func NewMsgAddNode(
	name string,
	cert string,
	operator sdk.AccAddress,
) *MsgAddNode {
	return &MsgAddNode{
		Name:        name,
		Certificate: cert,
		Operator:    operator.String(),
	}
}

// Route implements Msg
func (msg MsgAddNode) Route() string { return RouterKey }

// Type implements Msg
func (msg MsgAddNode) Type() string { return TypeMsgAddNode }

// ValidateBasic implements Msg
func (msg MsgAddNode) ValidateBasic() error {
	if err := ValidateOperator(msg.Operator); err != nil {
		return err
	}

	if len(strings.TrimSpace(msg.Name)) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "empty node name")
	}

	return ValidateCertificate(msg.Certificate)
}

// GetSignBytes implements Msg
func (msg MsgAddNode) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements Msg
func (msg MsgAddNode) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Operator)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{addr}
}

// NewMsgRemoveNode creates a new MsgRemoveNode instance
func NewMsgRemoveNode(
	id tmbytes.HexBytes,
	operator sdk.AccAddress,
) *MsgRemoveNode {
	return &MsgRemoveNode{
		Id:       id.String(),
		Operator: operator.String(),
	}
}

// Route implements Msg.
func (msg MsgRemoveNode) Route() string { return RouterKey }

// Type implements Msg.
func (msg MsgRemoveNode) Type() string { return TypeMsgRemoveNode }

// GetSignBytes implements Msg.
func (msg MsgRemoveNode) GetSignBytes() []byte {
	b := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(b)
}

// ValidateBasic implements Msg.
func (msg MsgRemoveNode) ValidateBasic() error {
	if err := ValidateOperator(msg.Operator); err != nil {
		return err
	}

	return ValidateNodeID(msg.Id)
}

// GetSigners implements Msg.
func (msg MsgRemoveNode) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Operator)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{addr}
}

// ValidateOperator validates the operator
func ValidateOperator(operator string) error {
	if operator == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "operator missing")
	}

	if _, err := sdk.AccAddressFromBech32(operator); err != nil {
		return err
	}

	return nil
}
