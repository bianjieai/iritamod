package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// node module sentinel errors
var (
	ErrInvalidNodeID      = sdkerrors.Register(ModuleName, 2, "invalid node ID")
	ErrInvalidCertificate = sdkerrors.Register(ModuleName, 3, "invalid certificate")
	ErrNodeExists         = sdkerrors.Register(ModuleName, 4, "node already exists")
	ErrUnknownNode        = sdkerrors.Register(ModuleName, 5, "unknown node")
)
