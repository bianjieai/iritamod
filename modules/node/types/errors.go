package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidRootCert       = sdkerrors.Register(ModuleName, 2, "invalid root certificate")
	ErrValidatorNameExists   = sdkerrors.Register(ModuleName, 3, "validator already exist for this validator name; must use new validator name")
	ErrValidatorPubkeyExists = sdkerrors.Register(ModuleName, 4, "validator already exist for this validator pubkey; must use new validator pubkey")
	ErrUnknownValidator      = sdkerrors.Register(ModuleName, 5, "unknown validator")
	ErrInvalidCert           = sdkerrors.Register(ModuleName, 6, "invalid certificate")
	ErrInvalidNodeID         = sdkerrors.Register(ModuleName, 7, "invalid node ID")
	ErrNodeExists            = sdkerrors.Register(ModuleName, 8, "node already exists")
	ErrUnknownNode           = sdkerrors.Register(ModuleName, 9, "unknown node")
)
