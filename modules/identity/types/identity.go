package types

import (
	"encoding/json"
	"strings"

	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/crypto/sm2"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewIdentity contructs a new Identity instance
func NewIdentity(
	id tmbytes.HexBytes,
	pubKeys []PubKeyInfo,
	certificates []string,
	credentials string,
	owner sdk.AccAddress,
) Identity {
	return Identity{
		Id:           id,
		PubKeys:      pubKeys,
		Certificates: certificates,
		Credentials:  credentials,
		Owner:        owner,
	}
}

// Validate validates the identity
func (i Identity) Validate() error {
	if i.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "owner missing")
	}

	if len(i.Id) != IDLength {
		return sdkerrors.Wrapf(ErrInvalidID, "size of the ID must be %d in bytes", IDLength)
	}

	for _, pubKey := range i.PubKeys {
		if err := pubKey.Validate(); err != nil {
			return err
		}
	}

	for _, cert := range i.Certificates {
		if err := CheckCertificate([]byte(cert)); err != nil {
			return err
		}
	}

	if len(i.Credentials) > MaxURILength {
		return sdkerrors.Wrapf(ErrInvalidCredentials, "length of the credentials uri must not be greater than %d", MaxURILength)
	}

	return nil
}

// NewPubKeyInfo constructs a new PubKeyInfo instance
func NewPubKeyInfo(pubKey tmbytes.HexBytes, algorithm PubKeyAlgorithm) PubKeyInfo {
	return PubKeyInfo{
		PubKey:    pubKey,
		Algorithm: algorithm,
	}
}

// Validate validates the public key against underlying constraints
// which vary by the algorithm
func (pki PubKeyInfo) Validate() error {
	switch pki.Algorithm {
	case RSA:
		if err := ValidateRSAPubKey(pki.PubKey); err != nil {
			return sdkerrors.Wrapf(ErrInvalidPubKey, err.Error())
		}

	case DSA:
		if err := ValidateDSAPubKey(pki.PubKey); err != nil {
			return sdkerrors.Wrapf(ErrInvalidPubKey, err.Error())
		}

	case ECDSA:
		if len(pki.PubKey) != 33 {
			return sdkerrors.Wrapf(ErrInvalidPubKey, "size of the ECDSA public key must be %d in bytes", 33)
		}

	case ED25519:
		if len(pki.PubKey) != ed25519.PubKeySize {
			return sdkerrors.Wrapf(ErrInvalidPubKey, "size of the ED25519 public key must be %d in bytes", ed25519.PubKeySize)
		}

	case SM2:
		if len(pki.PubKey) != sm2.PubKeySize {
			return sdkerrors.Wrapf(ErrInvalidPubKey, "size of the SM2 public key must be %d in bytes", sm2.PubKeySize)
		}

	default:
		return sdkerrors.Wrap(ErrUnsupportedPubKeyAlgorithm, "")
	}

	return nil
}

// PubKeyAlgorithmFromString converts the given string to PubKeyAlgorithm
func PubKeyAlgorithmFromString(str string) PubKeyAlgorithm {
	if pkAlgo, ok := PubKeyAlgorithm_value[strings.ToUpper(str)]; ok {
		return PubKeyAlgorithm(pkAlgo)
	}

	return UnknownPubKeyAlgorithm
}

// MarshalJSON returns the JSON representation
func (p PubKeyAlgorithm) MarshalJSON() ([]byte, error) {
	return json.Marshal(PubKeyAlgorithm_name[int32(p)])
}

// UnmarshalJSON unmarshals raw JSON bytes into a PubKeyAlgorithm
func (p *PubKeyAlgorithm) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return nil
	}

	*p = PubKeyAlgorithmFromString(s)
	return nil
}

// MarshalYAML returns the YAML representation
func (p PubKeyAlgorithm) MarshalYAML() (interface{}, error) {
	return PubKeyAlgorithm_name[int32(p)], nil
}
