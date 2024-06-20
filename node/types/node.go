package types

import (
	"encoding/hex"
	"strings"

	tmbytes "github.com/tendermint/tendermint/libs/bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	cautils "github.com/bianjieai/iritamod/utils/ca"
)

// NewNode contructs a new Node instance
func NewNode(
	id tmbytes.HexBytes,
	name string,
	cert string,
) Node {
	return Node{
		Id:          id.String(),
		Name:        name,
		Certificate: cert,
	}
}

// Validate validates the node
func (n Node) Validate() error {
	if err := ValidateNodeID(n.Id); err != nil {
		return err
	}

	if len(strings.TrimSpace(n.Name)) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "empty node name")
	}

	// TODO: ValidateCertificate(n.Certificate)
	// here is a workaround for genesis validation
	return nil
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
	_, err := cautils.ReadCertificateFromMem([]byte(cert))
	if err != nil {
		return sdkerrors.Wrap(ErrInvalidCert, err.Error())
	}

	return nil
}
