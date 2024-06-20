package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

var (
	_ sdk.Msg = &MsgUpgradeSoftware{}
	_ sdk.Msg = &MsgCancelUpgrade{}
)

// NewMsgUpgradeSoftware creates a new MsgUpgradeSoftware instance.
func NewMsgUpgradeSoftware(name string, height int64, info string, operator sdk.AccAddress) *MsgUpgradeSoftware {
	return &MsgUpgradeSoftware{
		Name:     name,
		Height:   height,
		Info:     info,
		Operator: operator.String(),
	}
}

func (m MsgUpgradeSoftware) Route() string {
	return RouterKey
}

func (m MsgUpgradeSoftware) Type() string {
	return "upgrade_software"
}

func (m MsgUpgradeSoftware) ValidateBasic() error {
	if len(m.Operator) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "signer address cannot be empty")
	}

	if _, err := sdk.AccAddressFromBech32(m.Operator); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid signer")
	}

	plan := upgradetypes.Plan{
		Name:   m.Name,
		Height: m.Height,
		Info:   m.Info,
	}
	return plan.ValidateBasic()
}

func (m MsgUpgradeSoftware) GetSignBytes() []byte {
	return sdk.MustSortJSON(SubModuleCdc.MustMarshalJSON(&m))
}

func (m MsgUpgradeSoftware) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(m.Operator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

// NewMsgCancelUpgrade creates a new MsgCancelUpgrade instance.
func NewMsgCancelUpgrade(operator sdk.AccAddress) *MsgCancelUpgrade {
	return &MsgCancelUpgrade{
		Operator: operator.String(),
	}
}

func (m MsgCancelUpgrade) Route() string {
	return RouterKey
}

func (m MsgCancelUpgrade) Type() string {
	return "cancel_upgrade"
}

func (m MsgCancelUpgrade) ValidateBasic() error {
	if len(m.Operator) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "signer address cannot be empty")
	}

	if _, err := sdk.AccAddressFromBech32(m.Operator); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid signer")
	}
	return nil
}

func (m MsgCancelUpgrade) GetSignBytes() []byte {
	return sdk.MustSortJSON(SubModuleCdc.MustMarshalJSON(&m))
}

func (m MsgCancelUpgrade) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(m.Operator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}
