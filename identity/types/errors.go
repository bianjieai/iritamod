package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// identity module sentinel errors
var (
	ErrInvalidID                  = sdkerrors.Register(ModuleName, 2, "invalid ID")
	ErrInvalidPubKey              = sdkerrors.Register(ModuleName, 3, "invalid public key")
	ErrPubKeyExists               = sdkerrors.Register(ModuleName, 4, "public key already exists")
	ErrInvalidCertificate         = sdkerrors.Register(ModuleName, 5, "invalid certificate")
	ErrCertificateExists          = sdkerrors.Register(ModuleName, 6, "certificate already exists")
	ErrInvalidCredentials         = sdkerrors.Register(ModuleName, 7, "invalid credentials uri")
	ErrIdentityExists             = sdkerrors.Register(ModuleName, 8, "identity already exists")
	ErrUnknownIdentity            = sdkerrors.Register(ModuleName, 9, "unknown identity")
	ErrUnsupportedPubKeyAlgorithm = sdkerrors.Register(ModuleName, 10, "unsupported public key algorithm; only RSA, DSA, ECDSA, ED25519 and SM2 supported")
	ErrNotAuthorized              = sdkerrors.Register(ModuleName, 11, "owner not matching")
)
