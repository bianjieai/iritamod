package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateL2Space   = "create_l2_space"
	TypeMsgTransferL2Space = "transfer_l2_space"
	TypeMsgCreateL2Record  = "create_l2_record"

	TypeMsgCreateNFTs          = "create_nfts"
	TypeMsgUpdateNFTs          = "update_nfts"
	TypeMsgDeleteNFTs          = "delete_nfts"
	TypeMsgUpdateClassesForNFT = "update_classes_for_nft"
	TypeMsgDepositClassForNFT  = "deposit_class_for_nft"
	TypeMsgWithdrawClassForNFT = "withdraw_class_for_nft"
	TypeMsgDepositTokenForNFT  = "deposit_token_for_nft"
	TypeMsgWithdrawTokenForNFT = "withdraw_token_for_nft"
)

var (
	_ sdk.Msg = &MsgCreateL2Space{}
	_ sdk.Msg = &MsgTransferL2Space{}
	_ sdk.Msg = &MsgCreateL2BlockHeader{}

	_ sdk.Msg = &MsgCreateNFTs{}
	_ sdk.Msg = &MsgUpdateNFTs{}
	_ sdk.Msg = &MsgDeleteNFTs{}
	_ sdk.Msg = &MsgUpdateClassesForNFT{}
	_ sdk.Msg = &MsgDepositClassForNFT{}
	_ sdk.Msg = &MsgWithdrawClassForNFT{}
	_ sdk.Msg = &MsgDepositTokenForNFT{}
	_ sdk.Msg = &MsgWithdrawTokenForNFT{}
)

// NewMsgCreateL2Space is a constructor function for MsgCreateL2Space
func NewMsgCreateL2Space(name, uri, sender string) *MsgCreateL2Space {
	return &MsgCreateL2Space{
		Name:   name,
		Uri:    uri,
		Sender: sender,
	}
}

func (msg MsgCreateL2Space) Route() string { return RouterKey }

func (msg MsgCreateL2Space) Type() string { return TypeMsgCreateL2Space }

func (msg MsgCreateL2Space) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	return nil
}

func (msg MsgCreateL2Space) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgCreateL2Space) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}

// NewMsgTransferL2Space is a constructor function for MsgTransferL2Space
func NewMsgTransferL2Space(spaceId uint64, recipient string, sender string) *MsgTransferL2Space {
	return &MsgTransferL2Space{
		SpaceId:   spaceId,
		Recipient: recipient,
		Sender:    sender,
	}
}

func (msg MsgTransferL2Space) Route() string { return RouterKey }

func (msg MsgTransferL2Space) Type() string { return TypeMsgTransferL2Space }

func (msg MsgTransferL2Space) ValidateBasic() error {
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

func (msg MsgTransferL2Space) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgTransferL2Space) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}

// NewMsgCreateL2BLockHeader is a constructor function for MsgCreateNFTs
func NewMsgCreateL2BLockHeader(spaceId, height uint64, header string, sender string) *MsgCreateL2BlockHeader {
	return &MsgCreateL2BlockHeader{
		SpaceId: spaceId,
		Height:  height,
		Header:  header,
		Sender:  sender,
	}
}

func (msg MsgCreateL2BlockHeader) Route() string { return RouterKey }

func (msg MsgCreateL2BlockHeader) Type() string { return TypeMsgCreateL2Record }

func (msg MsgCreateL2BlockHeader) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	if msg.Height == 0 {
		return sdkerrors.Wrapf(ErrInvalidRecord, "height cannot be zero")
	}

	if len(msg.Header) == 0 {
		return sdkerrors.Wrapf(ErrInvalidRecord, "header cannot be empty")
	}

	if err := ValidateSpaceId(msg.SpaceId); err != nil {
		return err
	}

	return nil
}

func (msg MsgCreateL2BlockHeader) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgCreateL2BlockHeader) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// NewMsgCreateNFTs is a constructor function for MsgCreateNFTs
func NewMsgCreateNFTs(spaceId uint64, classId string, tokens []TokenForNFT, sender string) *MsgCreateNFTs {
	return &MsgCreateNFTs{
		SpaceId: spaceId,
		ClassId: classId,
		Tokens:    tokens,
		Sender:  sender,
	}
}

func (msg MsgCreateNFTs) Route() string { return RouterKey }

func (msg MsgCreateNFTs) Type() string { return TypeMsgCreateNFTs }

func (msg MsgCreateNFTs) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	if err := ValidateSpaceId(msg.SpaceId); err != nil {
		return err
	}

	if err := ValidateClassIdForNFT(msg.ClassId); err != nil {
		return err
	}

	if err := ValidateTokensForNFT(msg.Tokens); err != nil {
		return err
	}

	return nil
}

func (msg MsgCreateNFTs) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgCreateNFTs) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// NewMsgUpdateNFTs is a constructor function for MsgUpdateNFTs
func NewMsgUpdateNFTs(spaceId uint64, classId string, tokens []TokenForNFT, sender string) *MsgUpdateNFTs {
	return &MsgUpdateNFTs{
		SpaceId: spaceId,
		ClassId: classId,
		Tokens:  tokens,
		Sender:  sender,
	}
}

func (msg MsgUpdateNFTs) Route() string { return RouterKey }

func (msg MsgUpdateNFTs) Type() string { return TypeMsgUpdateNFTs }

func (msg MsgUpdateNFTs) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	if err := ValidateSpaceId(msg.SpaceId); err != nil {
		return err
	}

	if err := ValidateClassIdForNFT(msg.ClassId); err != nil {
		return err
	}

	if err := ValidateTokensForNFT(msg.Tokens); err != nil {
		return err
	}

	return nil
}

func (msg MsgUpdateNFTs) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgUpdateNFTs) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// NewMsgDeleteNFTs is a constructor function for MsgDeleteNFTs
func NewMsgDeleteNFTs(spaceId uint64, classId string, tokenIds []string, sender string) *MsgDeleteNFTs {
	return &MsgDeleteNFTs{
		SpaceId: spaceId,
		ClassId: classId,
		TokenIds:  tokenIds,
		Sender:  sender,
	}
}

func (msg MsgDeleteNFTs) Route() string { return RouterKey }

func (msg MsgDeleteNFTs) Type() string { return TypeMsgDeleteNFTs }

func (msg MsgDeleteNFTs) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	if err := ValidateSpaceId(msg.SpaceId); err != nil {
		return err
	}

	if err := ValidateClassIdForNFT(msg.ClassId); err != nil {
		return err
	}

	// validate that the tokenId is not duplicated.
	seenIDs := make(map[string]bool)
	for _, tokenId := range msg.TokenIds {
		if seenIDs[tokenId] {
			return sdkerrors.Wrapf(ErrDuplicateTokenIdsForNFT, "token id %s is duplicated", tokenId)
		}
		seenIDs[tokenId] = true
	}

	return nil
}

func (msg MsgDeleteNFTs) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgDeleteNFTs) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// NewMsgUpdateClassesForNFT is a constructor function for MsgUpdateClassesForNFT
func NewMsgUpdateClassesForNFT(spaceId uint64, updateClasses []UpdateClassForNFT, sender string) *MsgUpdateClassesForNFT {
	return &MsgUpdateClassesForNFT{
		SpaceId: spaceId,
		ClassUpdatesForNft: updateClasses,
		Sender:             sender,
	}
}

func (msg MsgUpdateClassesForNFT) Route() string { return RouterKey }

func (msg MsgUpdateClassesForNFT) Type() string { return TypeMsgUpdateClassesForNFT }

func (msg MsgUpdateClassesForNFT) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	if err := ValidateClassUpdatesForNFT(msg.ClassUpdatesForNft); err != nil {
		return err
	}

	return nil
}

func (msg MsgUpdateClassesForNFT) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgUpdateClassesForNFT) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// NewMsgDepositClassForNFT is a constructor function for NewMsgDepositClassForNFT
func NewMsgDepositClassForNFT(spaceId uint64, classId, baseUri, recipient, sender string) *MsgDepositClassForNFT {
	return &MsgDepositClassForNFT{
		SpaceId:   spaceId,
		ClassId:   classId,
		BaseUri:   baseUri,
		Recipient: recipient,
		Sender:    sender,
	}
}

func (msg MsgDepositClassForNFT) Route() string { return RouterKey }

func (msg MsgDepositClassForNFT) Type() string { return TypeMsgDepositClassForNFT }

func (msg MsgDepositClassForNFT) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	if err := ValidateClassIdForNFT(msg.ClassId); err != nil {
		return err
	}

	return nil
}

func (msg MsgDepositClassForNFT) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgDepositClassForNFT) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// NewMsgWithdrawClassForNFT is a constructor function for NewMsgWithdrawClassForNFT
func NewMsgWithdrawClassForNFT(spaceId uint64, classId, owner, sender string) *MsgWithdrawClassForNFT {
	return &MsgWithdrawClassForNFT{
		ClassId: classId,
		Owner:   owner,
		Sender:  sender,
	}
}

func (msg MsgWithdrawClassForNFT) Route() string { return RouterKey }

func (msg MsgWithdrawClassForNFT) Type() string { return TypeMsgWithdrawClassForNFT }

func (msg MsgWithdrawClassForNFT) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	if _, err := sdk.AccAddressFromBech32(msg.Owner); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}

	if err := ValidateClassIdForNFT(msg.ClassId); err != nil {
		return err
	}

	return nil
}

func (msg MsgWithdrawClassForNFT) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgWithdrawClassForNFT) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// NewMsgDepositTokenForNFT is a constructor function for NewMsgDepositTokenForNFT
func NewMsgDepositTokenForNFT(spaceId uint64, classId, tokenId string, sender string) *MsgDepositTokenForNFT {
	return &MsgDepositTokenForNFT{
		SpaceId: spaceId,
		ClassId: classId,
		TokenId:   tokenId,
		Sender:  sender,
	}
}

func (msg MsgDepositTokenForNFT) Route() string { return RouterKey }

func (msg MsgDepositTokenForNFT) Type() string { return TypeMsgDepositTokenForNFT }

func (msg MsgDepositTokenForNFT) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	if err := ValidateSpaceId(msg.SpaceId); err != nil {
		return err
	}

	if err := ValidateClassIdForNFT(msg.ClassId); err != nil {
		return err
	}

	if err := ValidateTokenIdForNFT(msg.TokenId); err != nil {
		return err
	}

	return nil
}

func (msg MsgDepositTokenForNFT) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgDepositTokenForNFT) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func NewMsgWithdrawTokenForNFT(spaceId uint64, classId, tokenId, owner, name, uri, uriHash, data, sender string) *MsgWithdrawTokenForNFT {
	return &MsgWithdrawTokenForNFT{
		SpaceId: spaceId,
		ClassId: classId,
		TokenId:   tokenId,
		Owner:   owner,
		Name:    name,
		Uri:     uri,
		UriHash: uriHash,
		Data:    data,
		Sender:  sender,
	}
}

func (msg MsgWithdrawTokenForNFT) Route() string { return RouterKey }

func (msg MsgWithdrawTokenForNFT) Type() string { return TypeMsgWithdrawTokenForNFT }

func (msg MsgWithdrawTokenForNFT) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	if _, err := sdk.AccAddressFromBech32(msg.Owner); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}

	if err := ValidateSpaceId(msg.SpaceId); err != nil {
		return err
	}

	if err := ValidateClassIdForNFT(msg.ClassId); err != nil {
		return err
	}

	if err := ValidateTokenIdForNFT(msg.TokenId); err != nil {
		return err
	}
	return nil
}

func (msg MsgWithdrawTokenForNFT) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgWithdrawTokenForNFT) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}
