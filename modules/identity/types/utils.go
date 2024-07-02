package types

import (
	bytes "bytes"
	"crypto/elliptic"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/pem"
	"errors"
	"math/big"
	"time"

	tmbytes "github.com/tendermint/tendermint/libs/bytes"
	"github.com/tjfoc/gmsm/sm2"
	sm2x509 "github.com/tjfoc/gmsm/x509"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	oidPublicKeyRSA     = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 1}
	oidPublicKeyDSA     = asn1.ObjectIdentifier{1, 2, 840, 10040, 4, 1}
	oidPublicKeyECDSA   = asn1.ObjectIdentifier{1, 2, 840, 10045, 2, 1}
	oidPublicKeyEd25519 = asn1.ObjectIdentifier{1, 3, 101, 112}

	oidNamedCurveP224    = asn1.ObjectIdentifier{1, 3, 132, 0, 33}
	oidNamedCurveP256    = asn1.ObjectIdentifier{1, 2, 840, 10045, 3, 1, 7}
	oidNamedCurveP384    = asn1.ObjectIdentifier{1, 3, 132, 0, 34}
	oidNamedCurveP521    = asn1.ObjectIdentifier{1, 3, 132, 0, 35}
	oidNamedCurveP256SM2 = asn1.ObjectIdentifier{1, 2, 156, 10197, 1, 301}
)

// CheckCertificate checks if the given certificate is a PEM-encoded X.509 certificate
// and if the public key algorithm is supported
func CheckCertificate(cert []byte) error {
	certDERBlock, _ := pem.Decode(cert)
	if certDERBlock == nil {
		return sdkerrors.Wrap(ErrInvalidCertificate, "DER block missing")
	}

	asn1Cert, err := parseASN1Certificate(certDERBlock.Bytes)
	if err != nil {
		return err
	}

	pubKeyAlgo, err := getPubKeyAlgorithmFromCert(asn1Cert)
	if err != nil {
		return err
	}

	if pubKeyAlgo == UnknownPubKeyAlgorithm {
		return sdkerrors.Wrap(ErrUnsupportedPubKeyAlgorithm, "the public key algorithm of the certificate is not supported ")
	}

	if _, err := parseCertificate(certDERBlock.Bytes, pubKeyAlgo); err != nil {
		return err
	}

	return nil
}

// GetPubKeyFromCertificate retrieves the public key from the given certificate
// Note: make sure that the certificate is valid
func GetPubKeyFromCertificate(cert []byte) *PubKeyInfo {
	certDERBlock, _ := pem.Decode(cert)
	asn1Cert, _ := parseASN1Certificate(certDERBlock.Bytes)
	pubKeyAlgo, _ := getPubKeyAlgorithmFromCert(asn1Cert)

	var pki PubKeyInfo

	pki.PubKey = tmbytes.HexBytes(getPubKey(asn1Cert, pubKeyAlgo)).String()
	pki.Algorithm = pubKeyAlgo

	return &pki
}

// parseASN1Certificate parses the ASN.1 structured certificate
func parseASN1Certificate(asn1Data []byte) (*certificate, error) {
	var asn1Cert certificate

	rest, err := asn1.Unmarshal(asn1Data, &asn1Cert)
	if err != nil {
		return nil, sdkerrors.Wrap(ErrInvalidCertificate, err.Error())
	}
	if len(rest) > 0 {
		return nil, sdkerrors.Wrap(ErrInvalidCertificate, asn1.SyntaxError{Msg: "trailing data"}.Error())
	}

	return &asn1Cert, nil
}

// parseCertificate parses the X.509 certificate with the given public key algorithm
func parseCertificate(asn1Data []byte, pkAlgo PubKeyAlgorithm) (interface{}, error) {
	if pkAlgo == SM2 {
		x509Cert, err := sm2x509.ParseCertificate(asn1Data)
		if err != nil {
			return nil, sdkerrors.Wrap(ErrInvalidCertificate, err.Error())
		}

		return x509Cert, nil
	}

	x509Cert, err := x509.ParseCertificate(asn1Data)
	if err != nil {
		return nil, sdkerrors.Wrap(ErrInvalidCertificate, err.Error())
	}

	return x509Cert, nil
}

// getPubKeyAlgorithmFromCert gets the public key algorithm from the given ASN.1 certificate
func getPubKeyAlgorithmFromCert(cert *certificate) (PubKeyAlgorithm, error) {
	pubKeyAlgoOID := cert.TBSCertificate.PublicKey.Algorithm.Algorithm
	pubKeyAlgo := getPubKeyAlgorithmFromOID(pubKeyAlgoOID)

	if pubKeyAlgo == ECDSA {
		namedCurveOID, err := getNamedCurveOID(&cert.TBSCertificate.PublicKey)
		if err != nil {
			return pubKeyAlgo, sdkerrors.Wrap(ErrInvalidCertificate, err.Error())
		}

		if isSM2Curve(namedCurveOID) {
			return SM2, nil
		}
	}

	return pubKeyAlgo, nil
}

// getPubKeyAlgorithmFromOID gets the public key algorithm from the given object identifier
func getPubKeyAlgorithmFromOID(oid asn1.ObjectIdentifier) PubKeyAlgorithm {
	switch {
	case oid.Equal(oidPublicKeyRSA):
		return RSA
	case oid.Equal(oidPublicKeyDSA):
		return DSA
	case oid.Equal(oidPublicKeyECDSA):
		return ECDSA
	case oid.Equal(oidPublicKeyEd25519):
		return ED25519
	default:
		return UnknownPubKeyAlgorithm
	}
}

// getNamedCurveOID gets the named curve for the EC public key
func getNamedCurveOID(keyData *publicKeyInfo) (*asn1.ObjectIdentifier, error) {
	paramsData := keyData.Algorithm.Parameters.FullBytes
	namedCurveOID := new(asn1.ObjectIdentifier)

	rest, err := asn1.Unmarshal(paramsData, namedCurveOID)
	if err != nil {
		return nil, sdkerrors.Wrap(ErrInvalidCertificate, "failed to parse ECDSA parameters as named curve")
	}
	if len(rest) != 0 {
		return nil, sdkerrors.Wrap(ErrInvalidCertificate, "trailing data after ECDSA parameters")
	}

	return namedCurveOID, nil
}

// isSM2Curve returns true if the given object identifier is same as the sm2 object identifier,
// and false otherwise
func isSM2Curve(oid *asn1.ObjectIdentifier) bool {
	return oid.Equal(oidNamedCurveP256SM2)
}

// getPubKey gets the public key of the certificate with the specified algorithm
func getPubKey(asn1Cert *certificate, pkAlgo PubKeyAlgorithm) []byte {
	pubKeyASN1 := asn1Cert.TBSCertificate.PublicKey.PublicKey.RightAlign()

	switch pkAlgo {
	case RSA, DSA:
		return asn1Cert.TBSCertificate.PublicKey.Raw

	case ED25519:
		return pubKeyASN1

	case ECDSA, SM2:
		paramsData := asn1Cert.TBSCertificate.PublicKey.Algorithm.Parameters.FullBytes
		namedCurveOID := new(asn1.ObjectIdentifier)

		_, _ = asn1.Unmarshal(paramsData, namedCurveOID)
		namedCurve := namedCurveFromOID(*namedCurveOID)

		byteLen := (namedCurve.Params().BitSize + 7) >> 3
		x := new(big.Int).SetBytes(pubKeyASN1[1 : 1+byteLen])
		y := new(big.Int).SetBytes(pubKeyASN1[1+byteLen:])

		return compressECPubKey(x, y)

	default:
		return nil
	}
}

// namedCurveFromOID gets the named curve from the given object identifier
func namedCurveFromOID(oid asn1.ObjectIdentifier) elliptic.Curve {
	switch {
	case oid.Equal(oidNamedCurveP224):
		return elliptic.P224()
	case oid.Equal(oidNamedCurveP256):
		return elliptic.P256()
	case oid.Equal(oidNamedCurveP384):
		return elliptic.P384()
	case oid.Equal(oidNamedCurveP521):
		return elliptic.P521()
	case oid.Equal(oidNamedCurveP256SM2):
		return sm2.P256Sm2()
	}

	return nil
}

// compressECPubKey compresses the specified elliptic curve public key
func compressECPubKey(x, y *big.Int) []byte {
	b := make([]byte, 0)
	prefix := byte(0x2)

	if isOdd(y) {
		prefix |= 0x1
	}

	b = append(b, prefix)
	return paddedAppend(32, b, x.Bytes())
}

// isOdd returns true if the given number is odd, and false otherwise
func isOdd(a *big.Int) bool {
	return a.Bit(0) == 1
}

// paddedAppend appends the src byte slice to dst, returning the new slice.
// If the length of the source is smaller than the passed size, leading zero
// bytes are appended to the dst slice before appending src.
func paddedAppend(size uint, dst, src []byte) []byte {
	for i := 0; i < int(size)-len(src); i++ {
		dst = append(dst, 0)
	}
	return append(dst, src...)
}

// ValidateRSAPubKey validates the given RSA public key which is DER-encoded
func ValidateRSAPubKey(pubKey []byte) error {
	pk := new(publicKeyInfoNoRaw)
	rest, err := asn1.Unmarshal(pubKey, pk)
	if err != nil {
		return err
	}
	if len(rest) != 0 {
		return errors.New("trailing data")
	}

	if !bytes.Equal(pk.Algorithm.Parameters.FullBytes, asn1.NullBytes) {
		return errors.New("RSA key missing NULL parameters")
	}

	p := new(pkcs1PublicKey)
	rest, err = asn1.Unmarshal(pk.PublicKey.RightAlign(), p)
	if err != nil {
		return err
	}
	if len(rest) != 0 {
		return errors.New("trailing data after RSA public key")
	}

	if p.N.Sign() <= 0 {
		return errors.New("RSA modulus is not a positive number")
	}
	if p.E <= 0 {
		return errors.New("RSA public exponent is not a positive number")
	}

	return nil
}

// ValidateDSAPubKey validates the given DSA public key which is DER-encoded
func ValidateDSAPubKey(pubKey []byte) error {
	pk := new(publicKeyInfoNoRaw)
	rest, err := asn1.Unmarshal(pubKey, pk)
	if err != nil {
		return err
	}
	if len(rest) != 0 {
		return errors.New("trailing data")
	}

	var p *big.Int
	rest, err = asn1.Unmarshal(pk.PublicKey.RightAlign(), &p)
	if err != nil {
		return err
	}
	if len(rest) != 0 {
		return errors.New("trailing data after DSA public key")
	}

	paramsData := pk.Algorithm.Parameters.FullBytes
	params := new(dsaAlgorithmParameters)
	rest, err = asn1.Unmarshal(paramsData, params)
	if err != nil {
		return err
	}
	if len(rest) != 0 {
		return errors.New("trailing data after DSA parameters")
	}
	if p.Sign() <= 0 || params.P.Sign() <= 0 || params.Q.Sign() <= 0 || params.G.Sign() <= 0 {
		return errors.New("zero or negative DSA parameter")
	}

	return nil
}

// These structures reflect the ASN.1 structure of X.509 certificates.:

type certificate struct {
	Raw                asn1.RawContent
	TBSCertificate     tbsCertificate
	SignatureAlgorithm pkix.AlgorithmIdentifier
	SignatureValue     asn1.BitString
}

type tbsCertificate struct {
	Raw                asn1.RawContent
	Version            int `asn1:"optional,explicit,default:0,tag:0"`
	SerialNumber       *big.Int
	SignatureAlgorithm pkix.AlgorithmIdentifier
	Issuer             asn1.RawValue
	Validity           validity
	Subject            asn1.RawValue
	PublicKey          publicKeyInfo
	UniqueId           asn1.BitString   `asn1:"optional,tag:1"`
	SubjectUniqueId    asn1.BitString   `asn1:"optional,tag:2"`
	Extensions         []pkix.Extension `asn1:"optional,explicit,tag:3"`
}

type validity struct {
	NotBefore, NotAfter time.Time
}

type publicKeyInfo struct {
	Raw       asn1.RawContent
	Algorithm pkix.AlgorithmIdentifier
	PublicKey asn1.BitString
}

// pkcs1PublicKey reflects the ASN.1 structure of a PKCS#1 public key.
type pkcs1PublicKey struct {
	N *big.Int
	E int
}

// publicKeyInfoNoRaw is the ASN.1 public key
type publicKeyInfoNoRaw struct {
	Algorithm pkix.AlgorithmIdentifier
	PublicKey asn1.BitString
}

// dsaAlgorithmParameters is the parameters struct of the DSA public key
type dsaAlgorithmParameters struct {
	P, Q, G *big.Int
}
