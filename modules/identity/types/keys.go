package types

import (
	"encoding/binary"
)

const (
	// ModuleName is the name of the identity module
	ModuleName = "identity"

	// StoreKey is the string store representation
	StoreKey string = ModuleName

	// QuerierRoute is the querier route for the identity module
	QuerierRoute string = ModuleName

	// RouterKey is the msg router key for the identity module
	RouterKey string = ModuleName
)

var (
	// Keys for store prefixes
	OwnerKey          = []byte{0x01} // prefix for identity owner
	PubKeyInfoKey     = []byte{0x02} // prefix for public key
	CertificateKey    = []byte{0x03} // prefix for certificate
	CredentialsKey    = []byte{0x04} // prefix for credentials
	PubKeyIdentityKey = []byte{0x05} // prefix for mapping public key to identity
	DataKey           = []byte{0x06}
)

// GetOwnerKey gets the key for the owner of the specified identity
// VALUE: sdk.AccAddress (owner)
func GetOwnerKey(identityID []byte) []byte {
	return append(OwnerKey, identityID...)
}

// GetPubKeyInfoKey gets the key for the public key with the specified identity
// VALUE: []byte{}
func GetPubKeyInfoKey(identityID []byte, pubKey *PubKeyInfo) []byte {
	algoBz := make([]byte, 4)
	binary.BigEndian.PutUint32(algoBz, uint32(pubKey.Algorithm))

	return append(append(append(PubKeyInfoKey, identityID...), algoBz...), pubKey.PubKeyBytes()...)
}

// GetCertificateKey gets the key for the certificate with the specified identity and certificate hash
// VALUE: []byte (certificate)
func GetCertificateKey(identityID []byte, certHash []byte) []byte {
	return append(append(CertificateKey, identityID...), certHash...)
}

// GetCredentialsKey gets the key for the credentials of the specified identity
// VALUE: []byte (credentials)
func GetCredentialsKey(identityID []byte) []byte {
	return append(CredentialsKey, identityID...)
}

// GetDataKey gets the key for the data of the specified identity
// VALUE: []byte (data)
func GetDataKey(identityID []byte) []byte {
	return append(DataKey, identityID...)
}

// GetPubKeyIdentityKey gets the key for mapping the specified public key to the identity ID
// VALUE: []byte (identity ID)
func GetPubKeyIdentityKey(pubKey *PubKeyInfo) []byte {
	algoBz := make([]byte, 4)
	binary.BigEndian.PutUint32(algoBz, uint32(pubKey.Algorithm))

	return append(append(PubKeyIdentityKey, algoBz...), pubKey.PubKey...)
}

// GetPubKeySubspace gets the key prefix for the public keys of the specified identity
func GetPubKeySubspace(identityID []byte) []byte {
	return append(PubKeyInfoKey, identityID...)
}

// GetCertificateSubspace gets the key prefix for the certificates of the specified identity
func GetCertificateSubspace(identityID []byte) []byte {
	return append(CertificateKey, identityID...)
}
