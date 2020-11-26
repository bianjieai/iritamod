package types

import (
	tmbytes "github.com/tendermint/tendermint/libs/bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Node message types and params
const (
	TypeMsgAddNode    = "add_node"    // type for MsgAddNode
	TypeMsgRemoveNode = "remove_node" // type for MsgRemoveNode
)

var (
	_ sdk.Msg = &MsgAddNode{}
	_ sdk.Msg = &MsgRemoveNode{}
)

// NewMsgAddNode creates a new MsgAddNode instance
func NewMsgAddNode(
	cert string,
	operator sdk.AccAddress,
) *MsgAddNode {
	return &MsgAddNode{
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
