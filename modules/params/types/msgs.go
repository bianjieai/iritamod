package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	sdktx "github.com/cosmos/cosmos-sdk/types/tx"
)

// NewMsgUpdateParams creates a new MsgUpdateParams instance.
func NewMsgUpdateParams(messages []sdk.Msg, authority string) (*MsgUpdateParams, error) {
	anys, err := sdktx.SetMsgs(messages)
	if err != nil {
		return nil, err
	}

	return &MsgUpdateParams{
		Messages:  anys,
		Authority: authority,
	}, nil
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
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid authority address: %s", err)
	}

	msgs, err := m.GetMsgs()
	if err != nil {
		return err
	}

	for _, msg := range msgs {
		if err := msg.ValidateBasic(); err != nil {
			return err
		}
	}
	return nil
}

// GetSignBytes implements Msg
func (m MsgUpdateParams) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements Msg
func (m MsgUpdateParams) GetSigners() []sdk.AccAddress {
	singer, err := sdk.AccAddressFromBech32(m.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{singer}
}

// GetMsgs unpacks m.Messages Any's into sdk.Msg's
func (m MsgUpdateParams) GetMsgs() ([]sdk.Msg, error) {
	return sdktx.GetMsgs(m.Messages, "iritamod.MsgUpdateParams")
}
