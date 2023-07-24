package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewGenesisState creates a new GenesisState object
func NewGenesisState(spaceSequence uint64,
	spaces []Space,
	blockHeaders []BlockHeader,
	spaceLatestHeights []SpaceLatestHeight) *GenesisState {
	return &GenesisState{
		SpaceSequence:      spaceSequence,
		Spaces:             spaces,
		BlockHeaders:       blockHeaders,
		SpaceLatestHeights: spaceLatestHeights,
	}
}

// DefaultGenesisState creates a default GenesisState object
func DefaultGenesisState() *GenesisState {
	return NewGenesisState(0, []Space{}, []BlockHeader{}, []SpaceLatestHeight{})
}

// ValidateGenesis validates the provided genesis state to ensure the
// expected invariants holds.
func ValidateGenesis(data GenesisState) error {
	if uint64(len(data.Spaces)) != data.SpaceSequence {
		return sdkerrors.Wrapf(ErrInvalidSpaceId, "space count mismatch, wanted (%d) but got (%d)", data.SpaceSequence, len(data.Spaces))
	}

	// validate Spaces
	seenSpaceIds := make(map[uint64]bool)
	for _, space := range data.Spaces {
		if space.Id == 0 {
			return sdkerrors.Wrapf(ErrInvalidSpaceId, "space id must not be zero")
		}

		if _, err := sdk.AccAddressFromBech32(space.Owner); err != nil {
			return err
		}

		if seenSpaceIds[space.Id] {
			return sdkerrors.Wrapf(ErrInvalidSpaceId, "duplicate space (%d) during validation", space.Id)
		}
		seenSpaceIds[space.Id] = true
	}

	// validate BlockHeader
	seenBlockHeaderMap := make(map[string]bool)
	for _, header := range data.BlockHeaders {
		if !seenSpaceIds[header.SpaceId] {
			return sdkerrors.Wrapf(ErrInvalidSpaceId, "unknown space (%d) during validation", header.SpaceId)
		}

		// space_id/height is unique
		seenBlockHeader := fmt.Sprintf("%d-%d", header.SpaceId, header.Height)
		if seenBlockHeaderMap[seenBlockHeader] {
			return sdkerrors.Wrapf(ErrBlockHeader, "duplicate block header (%s) during validation", seenBlockHeader)
		}
		seenBlockHeaderMap[seenBlockHeader] = true
	}

	// validate SpaceLatestHeight
	for _, latestHeight := range data.SpaceLatestHeights {
		if !seenSpaceIds[latestHeight.SpaceId] {
			return sdkerrors.Wrapf(ErrInvalidSpaceId, "unknown space (%d) during validation", latestHeight.SpaceId)
		}

		seenBlockHeader := fmt.Sprintf("%d-%d", latestHeight.SpaceId, latestHeight.Height)
		if !seenBlockHeaderMap[seenBlockHeader] {
			return sdkerrors.Wrapf(ErrBlockHeader, "unknown block header (%s) during validation", seenBlockHeader)
		}
	}

	return nil
}
