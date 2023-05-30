package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewGenesisState creates a new GenesisState object
func NewGenesisState(spaceSequence uint64,
	spaces []Space,
	blockHeaders []L2BlockHeader,
	classesWithSpaceForNFT []ClassWithSpaceForNFT,
	collectionsForNFT []CollectionForNFT) *GenesisState {
	return &GenesisState{
		SpaceSequence:          spaceSequence,
		Spaces:                 spaces,
		L2BlockHeaders:         blockHeaders,
		ClassesWithSpaceForNft: classesWithSpaceForNFT,
		CollectionsForNft:      collectionsForNFT,
	}
}

// DefaultGenesisState creates a default GenesisState object
func DefaultGenesisState() *GenesisState {
	return NewGenesisState(0, []Space{}, []L2BlockHeader{}, []ClassWithSpaceForNFT{}, []CollectionForNFT{})
}

// ValidateGenesis validates the provided genesis state to ensure the
// expected invariants holds.
func ValidateGenesis(data GenesisState) error {
	if uint64(len(data.Spaces)) != data.SpaceSequence {
		return sdkerrors.Wrapf(ErrInvalidSpace, "space counts not match, want (%d) but got (%d)", data.SpaceSequence, len(data.Spaces))
	}

	// validate Spaces
	seenSpaceIds := make(map[uint64]bool)
	for _, space := range data.Spaces {
		if space.Id == 0 {
			return sdkerrors.Wrapf(ErrInvalidSpace, "invalid space (%d) during validation", space.Id)
		}

		if _, err := sdk.AccAddressFromBech32(space.Owner); err != nil {
			return err
		}

		if seenSpaceIds[space.Id] {
			return sdkerrors.Wrapf(ErrInvalidSpace, "duplicate space (%d) during validation", space.Id)
		}
		seenSpaceIds[space.Id] = true
	}

	// validate L2BlockHeader
	seenBlockHeaderMap := make(map[string]bool)
	for _, header := range data.L2BlockHeaders {
		if !seenSpaceIds[header.SpaceId] {
			return sdkerrors.Wrapf(ErrInvalidSpace, "unknown space (%d) during validation", header.SpaceId)
		}

		// space_id/height is unique
		seenBlockHeader := fmt.Sprintf("%d-%d", header.SpaceId, header.Height)
		if seenBlockHeaderMap[seenBlockHeader] {
			return sdkerrors.Wrapf(ErrBlockHeader, "duplicate block header (%s) during validation", seenBlockHeader)
		}
		seenBlockHeaderMap[seenBlockHeader] = true
	}

	// validate classes from NFT mappings
	seenClassesForNFT := make(map[string]bool)
	for _, class := range data.ClassesWithSpaceForNft {
		if err := ValidateClassIdForNFT(class.Id); err != nil {
			return err
		}

		if _, err := sdk.AccAddressFromBech32(class.Owner); err != nil {
			return err
		}

		if seenClassesForNFT[class.Id] {
			return sdkerrors.Wrapf(ErrInvalidClassId, "duplicate class (%s) during validation", class.Id)
		}
		seenClassesForNFT[class.Id] = true
	}

	for _, collection := range data.CollectionsForNft {
		if !seenSpaceIds[collection.SpaceId] {
			return sdkerrors.Wrapf(ErrInvalidSpace, "unknown space (%d) during validation", collection.SpaceId)
		}

		if !seenClassesForNFT[collection.ClassId] {
			return sdkerrors.Wrapf(ErrInvalidClassId, "unknown class (%s) during validation", collection.ClassId)
		}

		if err := ValidateClassIdForNFT(collection.ClassId); err != nil {
			return err
		}

		for _, token := range collection.Tokens {
			if _, err := sdk.AccAddressFromBech32(token.Owner); err != nil {
				return err
			}

			if err := ValidateTokenIdForNFT(token.Id); err != nil {
				return err
			}
		}
	}

	return nil
}
