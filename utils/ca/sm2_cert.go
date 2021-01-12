package ca

import (
	"errors"

	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/x509"
)

// Sm2Cert defines sm2 signed X509 certificate
type Sm2Cert struct {
	*x509.Certificate
	*sm2.PrivateKey
}

func ReadSM2CertFromMem(data []byte) (Cert, error) {
	cert, err := x509.ReadCertificateFromPem(data)
	return Sm2Cert{cert, nil}, err
}

func (sm2c Sm2Cert) WritePrivateKeytoMem() ([]byte, error) {
	return x509.WritePrivateKeyToPem(sm2c.PrivateKey, nil)
}

func (sm2c Sm2Cert) VerifyCertFromRoot(rootCert Cert) error {
	if rc, ok := rootCert.(Sm2Cert); ok {
		return sm2c.Certificate.CheckSignatureFrom(rc.Certificate)
	}
	return errors.New("can not verify sm2 certificate by other algorithm certificate")
}
