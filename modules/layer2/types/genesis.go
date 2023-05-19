package types

import (
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewGenesisState creates a new GenesisState object
func NewGenesisState(spaceSequence uint64,
	spaces []Space,
	blockHeaders []L2BlockHeader,
	classesForNFT []ClassForNFT,
	collectionsForNFT []CollectionForNFT) *GenesisState {
	return &GenesisState{
		SpaceSequence:   spaceSequence,
		Spaces:            spaces,
		L2BlockHeaders:    blockHeaders,
		ClassesForNft:     classesForNFT,
		CollectionsForNft: collectionsForNFT,
	}
}

// DefaultGenesisState creates a default GenesisState object
func DefaultGenesisState() *GenesisState {
	return NewGenesisState(0, []Space{}, []L2BlockHeader{}, []ClassForNFT{}, []CollectionForNFT{})
}

// ValidateGenesis validates the provided genesis state to ensure the
// expected invariants holds.
func ValidateGenesis(data GenesisState) error {
	if len(data.Spaces) != int(data.SpaceSequence) {
		return sdkerrors.Wrapf(ErrInvalidSpace, "space counts not match during space validation, want: %d, got: %d", data.SpaceSequence, len(data.Spaces))
	}

	// validate Spaces
	seenSpaceIds := make(map[uint64]bool)
	for _, space := range data.Spaces {
		if space.Id == 0 {
			return sdkerrors.Wrapf(ErrInvalidSpace, "invalid space id: %d during space validation", space.Id)
		}

		if seenSpaceIds[space.Id] {
			return sdkerrors.Wrapf(ErrDuplicateSpace, "duplicate space id: %d during space validation", space.Id)
		}
		seenSpaceIds[space.Id] = true
	}

	// validate L2BlockHeader
	seenBlockHeaderMap := make(map[string]bool)
	for _, header := range data.L2BlockHeaders {
		if !seenSpaceIds[header.SpaceId] {
			return sdkerrors.Wrapf(ErrInvalidSpace, "unknown space id: %d during record validation", header.SpaceId)
		}

		// space_id/height is unique
		seenBlockHeader := fmt.Sprintf("%d-%d", header.SpaceId, header.Height)
		if seenBlockHeaderMap[seenBlockHeader] {
			return sdkerrors.Wrapf(ErrDuplicateRecord, "duplicate record: %s during record validation", seenBlockHeader)
		}
		seenBlockHeaderMap[seenBlockHeader] = true
	}

	// validate classes from NFT mappings
	seenClassesForNFT := make(map[string]bool)
	for _, class := range data.ClassesForNft {
		if err := ValidateClassIdForNFT(class.Id); err != nil {
			return err
		}

		if seenClassesForNFT[class.Id] {
			return sdkerrors.Wrapf(ErrDuplicateClassIdForNFT, "duplicate class id: %s during class validation", class.Id)
		}
		seenClassesForNFT[class.Id] = true
	}

	for _, collection := range data.CollectionsForNft {
		if !seenSpaceIds[collection.SpaceId] {
			return sdkerrors.Wrapf(ErrInvalidSpace, "unknown space id: %d during collection validation", collection.SpaceId)
		}

		if !seenClassesForNFT[collection.ClassId] {
			return sdkerrors.Wrapf(ErrUnknownClassIdForNFT, "unknown class id: %s during collection validation", collection.ClassId)
		}

		if err := ValidateClassIdForNFT(collection.ClassId); err != nil {
			return err
		}

		for _, token := range collection.Tokens {
			if err := ValidateTokenIdForNFT(token.Id); err != nil {
				return err
			}
		}
	}

	return nil
}
