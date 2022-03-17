package ca_test

import (
	"fmt"
	"testing"
	cautil "github.com/bianjieai/iritamod/utils/ca"

)

func TestSm2Cert(t *testing.T) {
	/* Certificate */
	certpem := `-----BEGIN CERTIFICATE-----
MIICAjCCAaigAwIBAgIBATAKBggqgRzPVQGDdTBSMQswCQYDVQQGEwJDTjELMAkG
A1UECAwCQkoxCzAJBgNVBAcMAkJKMQwwCgYDVQQKDANQS1UxCzAJBgNVBAsMAkNB
MQ4wDAYDVQQDDAVQS1VDQTAeFw0xNzA2MDEwMDAwMDBaFw0yMDA2MDEwMDAwMDBa
MEYxCzAJBgNVBAYTAkNOMQswCQYDVQQIDAJCSjEMMAoGA1UECgwDUEtVMQswCQYD
VQQLDAJDQTEPMA0GA1UEAwwGYW50c3NzMFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0D
QgAEHpXtrYNlwesl7IyPuaHKKHqn4rHBk+tCU0l0T+zuBNMHAOJzKNDbobno6gOI
EQlVfC9q9uk9lO174GJsMLWJJqN7MHkwCQYDVR0TBAIwADAsBglghkgBhvhCAQ0E
HxYdT3BlblNTTCBHZW5lcmF0ZWQgQ2VydGlmaWNhdGUwHQYDVR0OBBYEFJsrRYOA
J8gpNq0KK6yuh/Dv9SjaMB8GA1UdIwQYMBaAFH1Dhf9CqQQYHF/8euzcPROIzn0r
MAoGCCqBHM9VAYN1A0gAMEUCIQCjrQ2nyiPqod/gZdj5X1+WW4fGtyqXvXLL3lOF
31nA/gIgZOpHLnvkyggY9VFfEQVp+8t6kewSfxb4eOImSu+dZcE=
-----END CERTIFICATE-----`
	sm2cert, err := cautil.ReadSM2CertFromMem([]byte(certpem))
	if err != nil {
		t.Fatalf("Read SM2 Certificate Error")
	}
	sm2pk, err := cautil.GetPubkeyFromCert(sm2cert)
	if err != nil {
		t.Fatalf("Get SM2 Public Key Error")
	}
	fmt.Println(sm2pk.Bytes())

	gmsslcert, err := cautil.ReadGMSSLCertFromMem([]byte(certpem))
	if err != nil {
		t.Fatalf("Read SM2 Certificate Error")
	}
	gmsslpk, err := cautil.GetPubkeyFromCert(gmsslcert)
	if err != nil {
		t.Fatalf("Get SM2 Public Key Error")
	}
	fmt.Println(gmsslpk.Bytes())
}