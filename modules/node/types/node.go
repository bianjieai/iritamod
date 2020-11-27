package types

import (
	"encoding/hex"

	tmbytes "github.com/tendermint/tendermint/libs/bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	cautils "github.com/bianjieai/iritamod/utils/ca"
)

// NewNode contructs a new Node instance
func NewNode(
	id tmbytes.HexBytes,
	cert string,
) Node {
	return Node{
		Id:          id.String(),
		Certificate: cert,
	}
}

// Validate validates the node
func (n Node) Validate() error {
	if err := ValidateNodeID(n.Id); err != nil {
		return err
	}

	return ValidateCertificate(n.Certificate)
}

// ValidateNodeID validates the node ID
func ValidateNodeID(id string) error {
	if len(id) == 0 {
		return sdkerrors.Wrap(ErrInvalidNodeID, "empty node ID")
	}

	bz, err := hex.DecodeString(id)
	if err != nil {
		return sdkerrors.Wrapf(ErrInvalidNodeID, "%s: %s", id, err)
	}

	err = sdk.VerifyAddressFormat(bz)
	if err != nil {
		return sdkerrors.Wrapf(ErrInvalidNodeID, "%s: %s", id, err)
	}

	return nil
}

// ValidateCertificate validates the node certificate
func ValidateCertificate(cert string) error {
	_, err := cautils.ReadX509CertFromMem([]byte(cert))
	if err != nil {
		return sdkerrors.Wrap(ErrInvalidCertificate, err.Error())
	}

	return nil
}
