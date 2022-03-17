package ca

import (
	"errors"

	"github.com/tendermint/tendermint/crypto/gmssl"
)

type GmSSLCert struct {
	*gmssl.Certificate
	*gmssl.PrivateKey
}

func ReadGMSSLCertFromMem(pem []byte) (Cert, error) {
	cert, err := gmssl.NewCertificateFromPEM(string(pem), "")
	return GmSSLCert{cert, nil}, err
}

func (gc GmSSLCert) WritePrivateKeytoMem() ([]byte, error) {
	ret, err := gc.PrivateKey.GetPEM("", "")
	return []byte(ret), err
}

func (gc GmSSLCert) VerifyCertFromRoot(rootCert Cert) error {
	if rc, ok := rootCert.(GmSSLCert); ok {
		return gc.Certificate.CheckSignatureFrom(rc.Certificate)
	}
	return errors.New("can not verify gmssl certificate by other algorithm certificate")
}