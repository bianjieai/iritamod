package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
)

var (
	_ sdk.Msg = &MsgAddToContractDenyList{}
	_ sdk.Msg = &MsgRemoveFromContractDenyList{}
)

func NewMsgAddToContractDenyList(contractAddr, from string) *MsgAddToContractDenyList {
	return &MsgAddToContractDenyList{
		contractAddr,
		from,
	}
}

func (m MsgAddToContractDenyList) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "string could not be parsed as address: %v", err)
	}

	if !common.IsHexAddress(m.ContractAddress) {
		return sdkerrors.Wrap(ErrInvalidContractAddress, "invalid from address")
	}
	return nil
}

func (m *MsgAddToContractDenyList) GetSigners() []sdk.AccAddress {
	if len(m.From) == 0 {
		panic("do not have signer")
	}
	accAddr, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

func NewMsgRemoveFromContractDenyList(contractAddr, from string) *MsgRemoveFromContractDenyList {
	return &MsgRemoveFromContractDenyList{
		contractAddr,
		from,
	}
}

func (m MsgRemoveFromContractDenyList) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "string could not be parsed as address: %v", err)
	}

	if !common.IsHexAddress(m.ContractAddress) {
		return sdkerrors.Wrap(ErrInvalidContractAddress, "invalid from address")
	}
	return nil
}

func (m *MsgRemoveFromContractDenyList) GetSigners() []sdk.AccAddress {
	if len(m.From) == 0 {
		panic("do not have signer")
	}
	accAddr, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}
