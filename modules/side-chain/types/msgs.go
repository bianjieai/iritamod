package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateSpace   = "create_space"
	TypeMsgTransferSpace = "transfer_space"
	TypeMsgCreateRecord  = "create_record"
)

var (
	_ sdk.Msg = &MsgCreateSpace{}
	_ sdk.Msg = &MsgTransferSpace{}
	_ sdk.Msg = &MsgCreateBlockHeader{}
)

// NewMsgCreateSpace is a constructor function for MsgCreateSpace
func NewMsgCreateSpace(name, uri, sender string) *MsgCreateSpace {
	return &MsgCreateSpace{
		Name:   name,
		Uri:    uri,
		Sender: sender,
	}
}

func (msg MsgCreateSpace) Route() string { return RouterKey }

func (msg MsgCreateSpace) Type() string { return TypeMsgCreateSpace }

func (msg MsgCreateSpace) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	return nil
}

func (msg MsgCreateSpace) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgCreateSpace) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}

// NewMsgTransferSpace is a constructor function for MsgTransferSpace
func NewMsgTransferSpace(spaceId uint64, recipient string, sender string) *MsgTransferSpace {
	return &MsgTransferSpace{
		SpaceId:   spaceId,
		Recipient: recipient,
		Sender:    sender,
	}
}

func (msg MsgTransferSpace) Route() string { return RouterKey }

func (msg MsgTransferSpace) Type() string { return TypeMsgTransferSpace }

func (msg MsgTransferSpace) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	if _, err := sdk.AccAddressFromBech32(msg.Recipient); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid recipient address (%s)", err)
	}

	if err := ValidateSpaceId(msg.SpaceId); err != nil {
		return err
	}

	return nil
}

func (msg MsgTransferSpace) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgTransferSpace) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}

// NewMsgCreateBlockHeader is a constructor function for MsgCreateNFTs
func NewMsgCreateBlockHeader(spaceId, height uint64, header string, sender string) *MsgCreateBlockHeader {
	return &MsgCreateBlockHeader{
		SpaceId: spaceId,
		Height:  height,
		Header:  header,
		Sender:  sender,
	}
}

func (msg MsgCreateBlockHeader) Route() string { return RouterKey }

func (msg MsgCreateBlockHeader) Type() string { return TypeMsgCreateRecord }

func (msg MsgCreateBlockHeader) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	if msg.Height == 0 {
		return sdkerrors.Wrapf(ErrBlockHeader, "height cannot be zero")
	}

	if len(msg.Header) == 0 {
		return sdkerrors.Wrapf(ErrBlockHeader, "header cannot be empty string")
	}

	if err := ValidateSpaceId(msg.SpaceId); err != nil {
		return err
	}

	return nil
}

func (msg MsgCreateBlockHeader) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgCreateBlockHeader) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}
