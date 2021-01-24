package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = &MsgAddRoles{}
	_ sdk.Msg = &MsgRemoveRoles{}
	_ sdk.Msg = &MsgBlockAccount{}
	_ sdk.Msg = &MsgUnblockAccount{}
)

// NewMsgAddRoles creates a new MsgAddRoles instance.
func NewMsgAddRoles(roles []Role, address, operator sdk.AccAddress) *MsgAddRoles {
	return &MsgAddRoles{
		Address:  address.String(),
		Roles:    roles,
		Operator: operator.String(),
	}
}

// Route return the RouterKey of MsgAddRoles
func (m MsgAddRoles) Route() string {
	return RouterKey
}

// Type return the Type of MsgAddRoles
func (m MsgAddRoles) Type() string {
	return "add_roles"
}

// ValidateBasic validate the message MsgAddRoles
func (m MsgAddRoles) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Address)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid address")
	}

	_, err = sdk.AccAddressFromBech32(m.Operator)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid address")
	}

	if len(m.Roles) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "roles missing")
	}
	for _, r := range m.Roles {
		if !ValidRole(r) {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid role %s", r.String())
		}
		if r == RoleRootAdmin {
			return ErrAddRootAdmin
		}
	}
	return nil
}

// GetSignBytes return the sign bytes
func (m MsgAddRoles) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

// GetSigners return the signers of MsgAddRoles
func (m MsgAddRoles) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Operator)
	return []sdk.AccAddress{addr}
}

// NewMsgRemoveRoles creates a new MsgRemoveRoles instance.
func NewMsgRemoveRoles(roles []Role, address, operator sdk.AccAddress) *MsgRemoveRoles {
	return &MsgRemoveRoles{
		Address:  address.String(),
		Roles:    roles,
		Operator: operator.String(),
	}
}

// Route return the RouterKey of MsgRemoveRoles
func (m MsgRemoveRoles) Route() string {
	return RouterKey
}

// Type return the Type of MsgRemoveRoles
func (m MsgRemoveRoles) Type() string {
	return "remove_roles"
}

// ValidateBasic validate the message MsgRemoveRoles
func (m MsgRemoveRoles) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Address)
	if err != nil {
		return err
	}

	_, err = sdk.AccAddressFromBech32(m.Operator)
	if err != nil {
		return err
	}

	if len(m.Roles) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "roles missing")
	}
	for _, r := range m.Roles {
		if !ValidRole(r) {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid role %s", r.String())
		}
		if r == RoleRootAdmin {
			return ErrRemoveRootAdmin
		}
	}
	return nil
}

// GetSignBytes return the sign bytes
func (m MsgRemoveRoles) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

// GetSigners return the signers of MsgAddRoles
func (m MsgRemoveRoles) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Operator)
	return []sdk.AccAddress{addr}
}

// NewMsgBlockAccount creates a new MsgBlockAccount instance.
func NewMsgBlockAccount(address, operator sdk.AccAddress) *MsgBlockAccount {
	return &MsgBlockAccount{
		Address:  address.String(),
		Operator: operator.String(),
	}
}

// Route return the RouterKey of MsgBlockAccount
func (m MsgBlockAccount) Route() string {
	return RouterKey
}

// Type return the Type of MsgBlockAccount
func (m MsgBlockAccount) Type() string {
	return "block_account"
}

// ValidateBasic validate the message MsgBlockAccount
func (m MsgBlockAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Address)
	if err != nil {
		return err
	}

	_, err = sdk.AccAddressFromBech32(m.Operator)
	if err != nil {
		return err
	}
	return nil
}

// GetSignBytes return the sign bytes
func (m MsgBlockAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

// GetSigners return the signers of MsgBlockAccount
func (m MsgBlockAccount) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Operator)
	return []sdk.AccAddress{addr}
}

// NewMsgUnblockAccount creates a new MsgUnblockAccount instance.
func NewMsgUnblockAccount(address, operator sdk.AccAddress) *MsgUnblockAccount {
	return &MsgUnblockAccount{
		Address:  address.String(),
		Operator: operator.String(),
	}
}

// Route return the RouterKey of MsgUnblockAccount
func (m MsgUnblockAccount) Route() string {
	return RouterKey
}

// Type return the Type of MsgUnblockAccount
func (m MsgUnblockAccount) Type() string {
	return "unblock_account"
}

// ValidateBasic validate the message MsgUnblockAccount
func (m MsgUnblockAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Address)
	if err != nil {
		return err
	}

	_, err = sdk.AccAddressFromBech32(m.Operator)
	if err != nil {
		return err
	}
	return nil
}

// GetSignBytes return the sign bytes
func (m MsgUnblockAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

// GetSigners return the signers of MsgUnblockAccount
func (m MsgUnblockAccount) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Operator)
	return []sdk.AccAddress{addr}
}
