package types

import (
	"encoding/hex"

	tmbytes "github.com/tendermint/tendermint/libs/bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Identity message types and params
const (
	TypeMsgCreateIdentity = "create_identity" // type for MsgCreateIdentity
	TypeMsgUpdateIdentity = "update_identity" // type for MsgUpdateIdentity

	IDLength     = 16  // size of the ID in bytes
	MaxURILength = 140 // maximum size of the URI

	DoNotModifyDesc = "[do-not-modify]" // description used to indicate not to modify a field
)

var (
	_ sdk.Msg = &MsgCreateIdentity{}
	_ sdk.Msg = &MsgUpdateIdentity{}
)

// NewMsgCreateIdentity creates a new MsgCreateIdentity instance
func NewMsgCreateIdentity(
	id tmbytes.HexBytes,
	pubKey *PubKeyInfo,
	certificate string,
	credentials string,
	owner sdk.AccAddress,
	data string,
) *MsgCreateIdentity {
	return &MsgCreateIdentity{
		Id:          id.String(),
		PubKey:      pubKey,
		Certificate: certificate,
		Credentials: credentials,
		Owner:       owner.String(),
		Data:        data,
	}
}

// Route implements Msg
func (msg MsgCreateIdentity) Route() string { return RouterKey }

// Type implements Msg
func (msg MsgCreateIdentity) Type() string { return TypeMsgCreateIdentity }

// ValidateBasic implements Msg
func (msg MsgCreateIdentity) ValidateBasic() error {
	return ValidateIdentityFields(
		msg.Id,
		msg.PubKey,
		msg.Certificate,
		msg.Credentials,
		msg.Owner,
		msg.Data,
	)
}

// GetSignBytes implements Msg
func (msg MsgCreateIdentity) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements Msg
func (msg MsgCreateIdentity) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// NewMsgUpdateIdentity creates a new MsgUpdateIdentity instance
func NewMsgUpdateIdentity(
	id tmbytes.HexBytes,
	pubKey *PubKeyInfo,
	certificate string,
	credentials string,
	owner sdk.AccAddress,
	data string,
) *MsgUpdateIdentity {
	return &MsgUpdateIdentity{
		Id:          id.String(),
		PubKey:      pubKey,
		Certificate: certificate,
		Credentials: credentials,
		Owner:       owner.String(),
		Data:        data,
	}
}

// Route implements Msg.
func (msg MsgUpdateIdentity) Route() string { return RouterKey }

// Type implements Msg.
func (msg MsgUpdateIdentity) Type() string { return TypeMsgUpdateIdentity }

// GetSignBytes implements Msg.
func (msg MsgUpdateIdentity) GetSignBytes() []byte {
	b := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(b)
}

// ValidateBasic implements Msg.
func (msg MsgUpdateIdentity) ValidateBasic() error {
	return ValidateIdentityFields(
		msg.Id,
		msg.PubKey,
		msg.Certificate,
		msg.Credentials,
		msg.Owner,
		msg.Data,
	)
}

// GetSigners implements Msg.
func (msg MsgUpdateIdentity) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// ValidateIdentityFields validates the given identity fields
func ValidateIdentityFields(
	id string,
	pubKey *PubKeyInfo,
	certificate string,
	credentials string,
	owner string,
	data string,
) error {
	if owner == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "owner missing")
	}

	if _, err := sdk.AccAddressFromBech32(owner); err != nil {
		return err
	}

	if len(id) != IDLength*2 {
		return sdkerrors.Wrapf(ErrInvalidID, "size of the ID must be %d in bytes", IDLength)
	}

	if _, err := hex.DecodeString(id); err != nil {
		return sdkerrors.Wrap(ErrInvalidID, "id not hex encoding")
	}

	if pubKey != nil {
		if err := pubKey.Validate(); err != nil {
			return err
		}
	}

	if len(certificate) > 0 {
		if err := CheckCertificate([]byte(certificate)); err != nil {
			return err
		}
	}

	if len(credentials) > MaxURILength {
		return sdkerrors.Wrapf(ErrInvalidCredentials, "length of the credentials uri must not be greater than %d", MaxURILength)
	}

	return nil
}
