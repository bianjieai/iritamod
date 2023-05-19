package types

import (
	"fmt"
	"regexp"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	nftClassIdString = `[a-z][a-zA-Z0-9/]{2,100}`
	nftTokenIdString = `[a-zA-Z0-9/]{1,100}`

	regexpNftClassId = regexp.MustCompile(fmt.Sprintf(`^%s$`, nftClassIdString)).MatchString
	regexpNftTokenId = regexp.MustCompile(fmt.Sprintf(`^%s$`, nftTokenIdString)).MatchString
)

func ValidateSpaceId(spaceId uint64) error {
	if spaceId == 0 {
		return sdkerrors.Wrapf(ErrInvalidSpace, "space id cannot be zero")
	}
	return nil
}

func ValidateClassIdForNFT(classId string) error {
	if !regexpNftClassId(classId) {
		return sdkerrors.Wrapf(ErrInvalidClassIdForNFT, "class id can only accept characters that match the regular expression: (%s),but got (%s)", nftClassIdString, classId)
	}
	return nil
}

func ValidateTokenIdForNFT(tokenId string) error {
	if !regexpNftClassId(tokenId) {
		return sdkerrors.Wrapf(ErrInvalidTokenIdForNFT, "token id can only accept characters that match the regular expression: (%s),but got (%s)", nftClassIdString, tokenId)
	}
	return nil
}

func ValidateTokensForNFT(nfts []*TokenForNFT) error {
	seenIDs := make(map[string]bool)

	for _, token := range nfts {
		if seenIDs[token.Id] {
			return sdkerrors.Wrapf(ErrDuplicateTokenIdsForNFT, "token id %s is duplicated", token.Id)
		}
		seenIDs[token.Id] = true

		if err := ValidateTokenIdForNFT(token.Id); err != nil {
			return err
		}

		if _, err := sdk.AccAddressFromBech32(token.Owner); err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s) with err (%s)", token.Owner, err.Error())
		}
	}

	return nil
}

func ValidateClassUpdatesForNFT(classes []*UpdateClassForNFT) error {
	seenIDs := make(map[string]bool)

	for _, class := range classes {
		if seenIDs[class.Id] {
			return sdkerrors.Wrapf(ErrDuplicateClassIdForNFT, "class id %s is duplicated", class.Id)
		}
		seenIDs[class.Id] = true

		if err := ValidateClassIdForNFT(class.Id); err != nil {
			return err
		}

		if _, err := sdk.AccAddressFromBech32(class.Owner); err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s) with error (%s)", class.Owner, err.Error())
		}
	}

	return nil
}
