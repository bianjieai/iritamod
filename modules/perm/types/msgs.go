package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgAssignRoles    = "assign_roles"    // type for MsgAssignRoles
	TypeMsgUnassignRoles  = "unassign_roles"  // type for MsgUnassignRoles
	TypeMsgBlockAccount   = "block_account"   // type for MsgBlockAccount
	TypeMsgUnblockAccount = "unblock_account" // type for MsgUnblockAccount
)

var (
	_ sdk.Msg = &MsgAssignRoles{}
	_ sdk.Msg = &MsgUnassignRoles{}
	_ sdk.Msg = &MsgBlockAccount{}
	_ sdk.Msg = &MsgUnblockAccount{}
)

// NewMsgAssignRoles creates a new MsgAssignRoles instance.
func NewMsgAssignRoles(roles []Role, address, operator sdk.AccAddress) *MsgAssignRoles {
	return &MsgAssignRoles{
		Address:  address.String(),
		Roles:    roles,
		Operator: operator.String(),
	}
}

// Route returns the RouterKey of MsgAssignRoles
func (m MsgAssignRoles) Route() string {
	return RouterKey
}

// Type returns the type of MsgAssignRoles
func (m MsgAssignRoles) Type() string {
	return TypeMsgAssignRoles
}

// ValidateBasic validates the message MsgAssignRoles
func (m MsgAssignRoles) ValidateBasic() error {
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

// GetSignBytes returns the sign bytes
func (m MsgAssignRoles) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the signers of MsgAssignRoles
func (m MsgAssignRoles) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Operator)
	return []sdk.AccAddress{addr}
}

// NewMsgUnassignRoles creates a new MsgUnassignRoles instance.
func NewMsgUnassignRoles(roles []Role, address, operator sdk.AccAddress) *MsgUnassignRoles {
	return &MsgUnassignRoles{
		Address:  address.String(),
		Roles:    roles,
		Operator: operator.String(),
	}
}

// Route returns the RouterKey of MsgUnassignRoles
func (m MsgUnassignRoles) Route() string {
	return RouterKey
}

// Type returns the type of MsgUnassignRoles
func (m MsgUnassignRoles) Type() string {
	return TypeMsgUnassignRoles
}

// ValidateBasic validates the message MsgUnassignRoles
func (m MsgUnassignRoles) ValidateBasic() error {
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

// GetSignBytes returns the sign bytes
func (m MsgUnassignRoles) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the signers of MsgAssignRoles
func (m MsgUnassignRoles) GetSigners() []sdk.AccAddress {
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

// Route returns the RouterKey of MsgBlockAccount
func (m MsgBlockAccount) Route() string {
	return RouterKey
}

// Type returns the type of MsgBlockAccount
func (m MsgBlockAccount) Type() string {
	return TypeMsgBlockAccount
}

// ValidateBasic validates the message MsgBlockAccount
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

// GetSignBytes returns the sign bytes
func (m MsgBlockAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the signers of MsgBlockAccount
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

// Route returns the RouterKey of MsgUnblockAccount
func (m MsgUnblockAccount) Route() string {
	return RouterKey
}

// Type returns the type of MsgUnblockAccount
func (m MsgUnblockAccount) Type() string {
	return TypeMsgUnblockAccount
}

// ValidateBasic validates the message MsgUnblockAccount
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

// GetSignBytes returns the sign bytes
func (m MsgUnblockAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the signers of MsgUnblockAccount
func (m MsgUnblockAccount) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Operator)
	return []sdk.AccAddress{addr}
}
