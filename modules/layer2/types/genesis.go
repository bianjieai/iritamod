package types

import (
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewGenesisState creates a new GenesisState object
func NewGenesisState(startingSpaceId uint64,
	spaces []Space,
	records []Record,
	nftMappings *MappingsForNFT) *GenesisState {
	return &GenesisState{
		StartingSpaceId: startingSpaceId,
		Spaces:          spaces,
		Records:         records,
		NftMappings:     nftMappings,
	}
}

// DefaultGenesisState creates a default GenesisState object
func DefaultGenesisState() *GenesisState {
	mappingsForNFT := &MappingsForNFT{
		Classes: []ClassForNFT{},
		Collections: []CollectionForNFT{},
	}

	return NewGenesisState(1, []Space{}, []Record{}, mappingsForNFT)
}

// ValidateGenesis validates the provided genesis state to ensure the
// expected invariants holds.
func ValidateGenesis(data GenesisState) error {
	if data.StartingSpaceId == 0 {
		return sdkerrors.Wrapf(ErrInvalidSpace, "invalid starting space id: %d", data.StartingSpaceId)
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

	// validate Records
	seenRecordMap := make(map[string]bool)
	for _, record := range data.Records {
		if !seenSpaceIds[record.SpaceId] {
			return sdkerrors.Wrapf(ErrInvalidSpace, "unknown space id: %d during record validation", record.SpaceId)
		}

		// space_id/height is unique
		seenRecord := fmt.Sprintf("%d-%d", record.SpaceId, record.Height)
		if seenRecordMap[seenRecord] {
			return sdkerrors.Wrapf(ErrDuplicateRecord, "duplicate record: %s during record validation", seenRecord)
		}
		seenRecordMap[seenRecord] = true
	}

	// validate classes from NFT mappings
	seenClassesForNFT := make(map[string]bool)
	for _, class := range data.NftMappings.Classes {
		if err := ValidateClassIdForNFT(class.Id); err != nil {
			return err
		}

		if seenClassesForNFT[class.Id] {
			return sdkerrors.Wrapf(ErrDuplicateClassIdForNFT, "duplicate class id: %s during class validation", class.Id)
		}
		seenClassesForNFT[class.Id] = true
	}

	for _, collection := range data.NftMappings.Collections {
		if !seenSpaceIds[collection.SpaceId] {
			return sdkerrors.Wrapf(ErrInvalidSpace, "unknown space id: %d during collection validation", collection.SpaceId)
		}

		if !seenClassesForNFT[collection.ClassId] {
			return sdkerrors.Wrapf(ErrUnkownClassIdForNFT, "unknown class id: %s during collection validation", collection.ClassId)
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
