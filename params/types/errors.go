package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/params/types/proposal"
)

// alias params module sentinel errors
var (
	ErrUnknownSubspace  = proposal.ErrUnknownSubspace
	ErrSettingParameter = proposal.ErrSettingParameter
	ErrEmptyChanges     = proposal.ErrEmptyChanges
	ErrEmptySubspace    = proposal.ErrEmptySubspace
	ErrEmptyKey         = proposal.ErrEmptyKey
	ErrEmptyValue       = proposal.ErrEmptyValue

	// Custom error codes start at 20
	ErrUnknownKey = sdkerrors.Register(ModuleName, 20, "unknown key")
)
