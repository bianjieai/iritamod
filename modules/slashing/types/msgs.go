package types

import (
	ctmbytes "github.com/cometbft/cometbft/libs/bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
)

var (
	_ sdk.Msg = &MsgUnjailValidator{}
	_ sdk.Msg = &MsgUpdateParams{}
)

// NewMsgUnjailValidator creates a new MsgUnjailValidator instance.
func NewMsgUnjailValidator(id ctmbytes.HexBytes, operator sdk.AccAddress) *MsgUnjailValidator {
	return &MsgUnjailValidator{
		Id:       id.String(),
		Operator: operator.String(),
	}
}

// Route implement sdk.Msg
func (m MsgUnjailValidator) Route() string {
	return slashingtypes.RouterKey
}

// Type implement sdk.Msg
func (m MsgUnjailValidator) Type() string {
	return "unjail_validator"
}

// ValidateBasic implement sdk.Msg
func (m MsgUnjailValidator) ValidateBasic() error {
	if len(m.Operator) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "operator missing")
	}
	if len(m.Id) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "validator id cannot be blank")
	}
	return nil
}

// GetSignBytes implement sdk.Msg
func (m MsgUnjailValidator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

// GetSigners implement sdk.Msg
func (m MsgUnjailValidator) GetSigners() []sdk.AccAddress {
	singer, err := sdk.AccAddressFromBech32(m.Operator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{singer}
}

// GetSignBytes implements the LegacyMsg interface.
func (m MsgUpdateParams) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic executes sanity validation on the provided data
func (m *MsgUpdateParams) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return sdkerrors.Wrap(err, "invalid authority address")
	}

	params := slashingtypes.Params{
		SignedBlocksWindow:      m.Params.SignedBlocksWindow,
		MinSignedPerWindow:      m.Params.MinSignedPerWindow,
		DowntimeJailDuration:    m.Params.DowntimeJailDuration,
		SlashFractionDoubleSign: m.Params.SlashFractionDoubleSign,
		SlashFractionDowntime:   m.Params.SlashFractionDowntime,
	}

	return params.Validate()
}

// GetSigners returns the expected signers for a MsgUpdateParams message
func (m *MsgUpdateParams) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}
