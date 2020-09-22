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

// NewMsgCreateValidator creates a new MsgAddRoles instance.
func NewMsgAddRoles(roles []Role, address, operator sdk.AccAddress) *MsgAddRoles {
	return &MsgAddRoles{
		Address:  address,
		Roles:    roles,
		Operator: operator,
	}
}

func (m MsgAddRoles) Route() string {
	return RouterKey
}

func (m MsgAddRoles) Type() string {
	return "add_roles"
}

func (m MsgAddRoles) ValidateBasic() error {
	if m.Address.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "address missing")
	}
	if m.Operator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "operator missing")
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

func (m MsgAddRoles) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

func (m MsgAddRoles) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Operator}
}

// NewMsgRemoveRoles creates a new MsgRemoveRoles instance.
func NewMsgRemoveRoles(roles []Role, address, operator sdk.AccAddress) *MsgRemoveRoles {
	return &MsgRemoveRoles{
		Address:  address,
		Roles:    roles,
		Operator: operator,
	}
}

func (m MsgRemoveRoles) Route() string {
	return RouterKey
}

func (m MsgRemoveRoles) Type() string {
	return "remove_roles"
}

func (m MsgRemoveRoles) ValidateBasic() error {
	if m.Address.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "address missing")
	}
	if m.Operator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "operator missing")
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

func (m MsgRemoveRoles) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

func (m MsgRemoveRoles) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Operator}
}

// NewMsgBlockAccount creates a new MsgBlockAccount instance.
func NewMsgBlockAccount(address, operator sdk.AccAddress) *MsgBlockAccount {
	return &MsgBlockAccount{
		Address:  address,
		Operator: operator,
	}
}

func (m MsgBlockAccount) Route() string {
	return RouterKey
}

func (m MsgBlockAccount) Type() string {
	return "block_account"
}

func (m MsgBlockAccount) ValidateBasic() error {
	if m.Address.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "address missing")
	}
	if m.Operator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "operator missing")
	}
	return nil
}

func (m MsgBlockAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

func (m MsgBlockAccount) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Operator}
}

// NewMsgUnblockAccount creates a new MsgUnblockAccount instance.
func NewMsgUnblockAccount(address, operator sdk.AccAddress) *MsgUnblockAccount {
	return &MsgUnblockAccount{
		Address:  address,
		Operator: operator,
	}
}

func (m MsgUnblockAccount) Route() string {
	return RouterKey
}

func (m MsgUnblockAccount) Type() string {
	return "unblock_account"
}

func (m MsgUnblockAccount) ValidateBasic() error {
	if m.Address.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "address missing")
	}
	if m.Operator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "operator missing")
	}
	return nil
}

func (m MsgUnblockAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

func (m MsgUnblockAccount) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Operator}
}
