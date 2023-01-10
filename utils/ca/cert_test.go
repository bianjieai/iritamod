package ca_test

import (
	"fmt"
	cautil "github.com/bianjieai/iritamod/utils/ca"
	"testing"
)

var rootCert = `-----BEGIN CERTIFICATE-----
MIICBTCCAaugAwIBAgIUYuQTIaBJbRbm6UDlcwdD7Ti2p+QwCgYIKoEcz1UBg3Uw
WDELMAkGA1UEBhMCQ04xDTALBgNVBAgMBHJvb3QxDTALBgNVBAcMBHJvb3QxDTAL
BgNVBAoMBHJvb3QxDTALBgNVBAsMBHJvb3QxDTALBgNVBAMMBHJvb3QwHhcNMjMw
MTEwMDIxNTA0WhcNMjQwMTEwMDIxNTA0WjBYMQswCQYDVQQGEwJDTjENMAsGA1UE
CAwEcm9vdDENMAsGA1UEBwwEcm9vdDENMAsGA1UECgwEcm9vdDENMAsGA1UECwwE
cm9vdDENMAsGA1UEAwwEcm9vdDBZMBMGByqGSM49AgEGCCqBHM9VAYItA0IABO8q
GZPMFdeHMV8Ptr/JPr5z/Rn9o3yDxmNBCyvqeEBzlPbUaYLrtyX3UpSt5CTaVWwR
W4s1PDZ+2TEO+K8YMXCjUzBRMB0GA1UdDgQWBBQWa18ifHI2LfsnWI6ywIYed3qf
fDAfBgNVHSMEGDAWgBQWa18ifHI2LfsnWI6ywIYed3qffDAPBgNVHRMBAf8EBTAD
AQH/MAoGCCqBHM9VAYN1A0gAMEUCICblwu1sCNSI/FUdSCrNexS4Krkkv0aPwxPy
9snHn+/kAiEAi+NVOjuXaJdPb6he6tXbtJ+3AAVIcYh/qTSd7kZ06wA=
-----END CERTIFICATE-----
`

var certpem = `-----BEGIN CERTIFICATE-----
MIIBqzCCAVECFAv2fNTDZI+TSvp0Bq0fStRL4vf6MAoGCCqBHM9VAYN1MFgxCzAJ
BgNVBAYTAkNOMQ0wCwYDVQQIDARyb290MQ0wCwYDVQQHDARyb290MQ0wCwYDVQQK
DARyb290MQ0wCwYDVQQLDARyb290MQ0wCwYDVQQDDARyb290MB4XDTIzMDExMDAy
Mjg0NloXDTI0MDExMDAyMjg0NlowWDELMAkGA1UEBhMCQ04xDTALBgNVBAgMBHJv
b3QxDTALBgNVBAcMBHJvb3QxDTALBgNVBAoMBHJvb3QxDTALBgNVBAsMBHJvb3Qx
DTALBgNVBAMMBHJvb3QwWTATBgcqhkjOPQIBBggqgRzPVQGCLQNCAATvKhmTzBXX
hzFfD7a/yT6+c/0Z/aN8g8ZjQQsr6nhAc5T21GmC67cl91KUreQk2lVsEVuLNTw2
ftkxDvivGDFwMAoGCCqBHM9VAYN1A0gAMEUCIQDCwnr9jzMUBEo42nVHO3Mw82GM
Ih/5snbg3ZzXnfcd3AIgLG1swYwYEhLmKzd/FuMK0P9SQ0KZWWD04nMxU5V3Qi0=
-----END CERTIFICATE-----
`

func TestSm2Cert(t *testing.T) {
	rootC, err := cautil.ReadSM2CertFromMem([]byte(rootCert))
	if err != nil {
		t.Fatalf("Read SM2 Certificate Error: %s", err)
	}
	sm2cert, err := cautil.ReadSM2CertFromMem([]byte(certpem))
	if err != nil {
		t.Fatalf("Read SM2 Certificate Error: %s", err)
	}

	err = sm2cert.VerifyCertFromRoot(rootC)
	if err != nil {
		t.Fatalf("verify cert error: %s", err)
	}

	sm2pk, err := cautil.GetPubkeyFromCert(sm2cert)
	if err != nil {
		t.Fatalf("Get SM2 Public Key Error: %s", err)
	}
	fmt.Println(sm2pk.Bytes())

}

func TestGMSSLSm2Cert(t *testing.T) {
	/* Certificate */

	gmsslrootcert, err := cautil.ReadGMSSLCertFromMem([]byte(rootCert))
	if err != nil {
		t.Fatalf("Read GMSSLSM2 Certificate Error: %s", err)
	}

	gmsslcert, err := cautil.ReadGMSSLCertFromMem([]byte(certpem))
	if err != nil {
		t.Fatalf("Read GMSSLSM2 Certificate Error: %s", err)
	}

	err = gmsslcert.VerifyCertFromRoot(gmsslrootcert)
	if err != nil {
		t.Fatalf("Verify GMSSLSM2 Certificate Error: %s", err)
	}
	gmsslpk, err := cautil.GetPubkeyFromCert(gmsslcert)
	if err != nil {
		t.Fatalf("Get GMSSLSM2 Public Key Error: %s", err)
	}
	fmt.Println(gmsslpk.Bytes())
}
