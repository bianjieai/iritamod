package types

import (
	"crypto/rand"
	"fmt"
	"strings"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"github.com/tjfoc/gmsm/sm2"

	tmbytes "github.com/tendermint/tendermint/libs/bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	testID    = uuid.NewV4().Bytes()
	testIDStr = tmbytes.HexBytes(testID).String()

	testPrivKeySM2, _ = sm2.GenerateKey(rand.Reader)
	testPubKeySM2     = sm2.Compress(&testPrivKeySM2.PublicKey)
	testPubKeySM2Str  = tmbytes.HexBytes(testPubKeySM2).String()
	testPubKeySM2Info = PubKeyInfo{PubKey: testPubKeySM2Str, Algorithm: SM2}

	testCredentials = "https://kyc.com/user/10001"
	testOwner       = sdk.AccAddress([]byte("test-ownertest-owner"))
	testData        = "test_data"
)

// TestMsgCreateIdentityRoute tests Route for MsgCreateIdentity
func TestMsgCreateIdentityRoute(t *testing.T) {
	msg := NewMsgCreateIdentity(testID, &testPubKeySM2Info, testCertificate, testCredentials, testOwner, testData)

	require.Equal(t, RouterKey, msg.Route())
}

// TestMsgCreateIdentityType tests Type for MsgCreateIdentity
func TestMsgCreateIdentityType(t *testing.T) {
	msg := NewMsgCreateIdentity(testID, &testPubKeySM2Info, testCertificate, testCredentials, testOwner, testData)

	require.Equal(t, "create_identity", msg.Type())
}

// TestMsgCreateIdentityValidation tests ValidateBasic for MsgCreateIdentity
func TestMsgCreateIdentityValidation(t *testing.T) {
	emptyAddress := sdk.AccAddress{}

	invalidID := []byte("invalidID")
	invalidPubKey := PubKeyInfo{PubKey: "invalidPubKey", Algorithm: UnknownPubKeyAlgorithm}
	invalidCertificate := "invalidCertificate"
	invalidCredentials := testCredentials + strings.Repeat("c", MaxURILength)

	testMsgs := []*MsgCreateIdentity{
		NewMsgCreateIdentity(testID, &testPubKeySM2Info, testCertificate, testCredentials, testOwner, testData),    // valid msg
		NewMsgCreateIdentity(testID, nil, "", "", testOwner, testData),                                             // public key, certificate and credentials are allowed to be empty
		NewMsgCreateIdentity(testID, &testPubKeySM2Info, testCertificate, testCredentials, emptyAddress, testData), // missing owner address
		NewMsgCreateIdentity(nil, &testPubKeySM2Info, testCertificate, testCredentials, testOwner, testData),       // missing ID
		NewMsgCreateIdentity(invalidID, &testPubKeySM2Info, testCertificate, testCredentials, testOwner, testData), // invalid ID
		NewMsgCreateIdentity(testID, &invalidPubKey, testCertificate, testCredentials, testOwner, testData),        // invalid public key
		NewMsgCreateIdentity(testID, &testPubKeySM2Info, invalidCertificate, testCredentials, testOwner, testData), // invalid certificate
		NewMsgCreateIdentity(testID, &testPubKeySM2Info, testCertificate, invalidCredentials, testOwner, testData), // invalid credentials
	}

	testCases := []struct {
		msg     *MsgCreateIdentity
		expPass bool
		errMsg  string
	}{
		{testMsgs[0], true, ""},
		{testMsgs[1], true, ""},
		{testMsgs[2], false, "missing owner address"},
		{testMsgs[3], false, "missing ID"},
		{testMsgs[4], false, "invalid ID"},
		{testMsgs[5], false, "invalid public key"},
		{testMsgs[6], false, "invalid certificate"},
		{testMsgs[7], false, "invalid credentials"},
	}

	for i, tc := range testCases {
		err := tc.msg.ValidateBasic()
		if tc.expPass {
			require.NoError(t, err, "Msg %d failed: %v", i, err)
		} else {
			require.Error(t, err, "Invalid Msg %d passed: %s", i, tc.errMsg)
		}
	}
}

// TestMsgCreateIdentityGetSignBytes tests GetSignBytes for MsgCreateIdentity
func TestMsgCreateIdentityGetSignBytes(t *testing.T) {
	msg := NewMsgCreateIdentity(testID, &testPubKeySM2Info, testCertificate, testCredentials, testOwner, testData)
	res := msg.GetSignBytes()

	expected := fmt.Sprintf(`{"type":"iritamod/identity/MsgCreateIdentity","value":{"certificate":"%s","credentials":"https://kyc.com/user/10001","data":"%s","id":"%s","owner":"cosmos1w3jhxapddamkuetjw3jhxapddamkuetjgzplvk","pubkey":{"algorithm":"SM2","pubkey":"%s"}}}`, strings.ReplaceAll(testCertificate, "\n", "\\n"), testData, testIDStr, testPubKeySM2Str)
	require.Equal(t, expected, string(res))
}

// TestMsgCreateIdentityGetSigners tests GetSigners for MsgCreateIdentity
func TestMsgCreateIdentityGetSigners(t *testing.T) {
	msg := NewMsgCreateIdentity(testID, &testPubKeySM2Info, testCertificate, testCredentials, testOwner, testData)
	res := msg.GetSigners()

	expected := "[746573742D6F776E6572746573742D6F776E6572]"
	require.Equal(t, expected, fmt.Sprintf("%v", res))
}

// TestMsgUpdateIdentityRoute tests Route for MsgUpdateIdentity
func TestMsgUpdateIdentityRoute(t *testing.T) {
	msg := NewMsgUpdateIdentity(testID, &testPubKeySM2Info, testCertificate, testCredentials, testOwner, testData)

	require.Equal(t, RouterKey, msg.Route())
}

// TestMsgUpdateIdentityType tests Type for MsgUpdateIdentity
func TestMsgUpdateIdentityType(t *testing.T) {
	msg := NewMsgUpdateIdentity(testID, &testPubKeySM2Info, testCertificate, testCredentials, testOwner, testData)

	require.Equal(t, "update_identity", msg.Type())
}

// TestMsgUpdateIdentityValidation tests ValidateBasic for MsgUpdateIdentity
func TestMsgUpdateIdentityValidation(t *testing.T) {
	emptyAddress := sdk.AccAddress{}

	invalidID := []byte("invalidID")
	invalidPubKey := PubKeyInfo{PubKey: "invalidPubKey", Algorithm: UnknownPubKeyAlgorithm}
	invalidCertificate := "invalidCertificate"
	invalidCredentials := testCredentials + strings.Repeat("c", MaxURILength)

	testMsgs := []*MsgUpdateIdentity{
		NewMsgUpdateIdentity(testID, &testPubKeySM2Info, testCertificate, testCredentials, testOwner, testData), // valid msg
		NewMsgUpdateIdentity(testID, nil, "", "", testOwner, testData),                                          // public key, certificate and credentials are allowed to be empty
		NewMsgUpdateIdentity(testID, nil, "", "", emptyAddress, testData),                                       // missing owner address
		NewMsgUpdateIdentity(nil, nil, "", "", testOwner, testData),                                             // missing ID
		NewMsgUpdateIdentity(invalidID, nil, "", "", testOwner, testData),                                       // invalid ID
		NewMsgUpdateIdentity(testID, &invalidPubKey, "", "", testOwner, testData),                               // invalid public key
		NewMsgUpdateIdentity(testID, nil, invalidCertificate, "", testOwner, testData),                          // invalid certificate
		NewMsgUpdateIdentity(testID, nil, "", invalidCredentials, testOwner, testData),                          // invalid credentials
	}

	testCases := []struct {
		msg     *MsgUpdateIdentity
		expPass bool
		errMsg  string
	}{
		{testMsgs[0], true, ""},
		{testMsgs[1], true, ""},
		{testMsgs[2], false, "missing owner address"},
		{testMsgs[3], false, "missing ID"},
		{testMsgs[4], false, "invalid ID"},
		{testMsgs[5], false, "invalid public key"},
		{testMsgs[6], false, "invalid certificate"},
		{testMsgs[7], false, "invalid credentials"},
	}

	for i, tc := range testCases {
		err := tc.msg.ValidateBasic()
		if tc.expPass {
			require.NoError(t, err, "Msg %d failed: %v", i, err)
		} else {
			require.Error(t, err, "Invalid Msg %d passed: %s", i, tc.errMsg)
		}
	}
}

// TestMsgUpdateIdentityGetSignBytes tests GetSignBytes for MsgUpdateIdentity
func TestMsgUpdateIdentityGetSignBytes(t *testing.T) {
	msg := NewMsgUpdateIdentity(testID, &testPubKeySM2Info, testCertificate, testCredentials, testOwner, testData)
	res := msg.GetSignBytes()

	expected := fmt.Sprintf(`{"type":"iritamod/identity/MsgUpdateIdentity","value":{"certificate":"%s","credentials":"https://kyc.com/user/10001","data":"%s","id":"%s","owner":"cosmos1w3jhxapddamkuetjw3jhxapddamkuetjgzplvk","pubkey":{"algorithm":"SM2","pubkey":"%s"}}}`, strings.ReplaceAll(testCertificate, "\n", "\\n"), testData, testIDStr, testPubKeySM2Str)
	require.Equal(t, expected, string(res))
}

// TestMsgUpdateIdentityGetSigners tests GetSigners for MsgUpdateIdentity
func TestMsgUpdateIdentityGetSigners(t *testing.T) {
	msg := NewMsgUpdateIdentity(testID, &testPubKeySM2Info, testCertificate, testCredentials, testOwner, testData)
	res := msg.GetSigners()

	expected := "[746573742D6F776E6572746573742D6F776E6572]"
	require.Equal(t, expected, fmt.Sprintf("%v", res))
}

const testCertificate = `-----BEGIN CERTIFICATE-----
MIIDTDCCAjQCCQDvRoz+e/HRpDANBgkqhkiG9w0BAQsFADBoMQswCQYDVQQGEwJj
bjELMAkGA1UECAwCc2gxCzAJBgNVBAcMAnBkMQswCQYDVQQKDAJiajELMAkGA1UE
CwwCYmoxCzAJBgNVBAMMAmJqMRgwFgYJKoZIhvcNAQkBFgliakBiai5jb20wHhcN
MjAwNjEwMDkzNjMxWhcNMjAwNzEwMDkzNjMxWjBoMQswCQYDVQQGEwJjbjELMAkG
A1UECAwCc2gxCzAJBgNVBAcMAnBkMQswCQYDVQQKDAJiajELMAkGA1UECwwCYmox
CzAJBgNVBAMMAmJqMRgwFgYJKoZIhvcNAQkBFgliakBiai5jb20wggEiMA0GCSqG
SIb3DQEBAQUAA4IBDwAwggEKAoIBAQDUEifXes1/CXEjdH8SeSS+1x+ZlhktI8i8
9ncMeOr5oI1Mc7Kd7v85i0hrmjjZzUrHQy0Sdt2ltQjo6dtkq3wDsL4OgIqGO75z
OwG4EB0A1sJ/YTSX+fmWwy5ys19A2O5sTZOJEw3VFgiZHv1TZEiY+GVtpZ5Dti/1
t5ZzNTF+M0rpbICTxLh1GSpdhJs95yci1A8zqmPzPETVkxJwVCOg54WfpRQAiBqM
DKLjVXALuvlDDxVhB0u7kuvKAydZdV/pDs73HuY2srCOiDij3iVS01Ln02JNeMK8
IG9xRSw2eaSDp+fa1jtUXMDMmVNHCJqpQaFv0/1oN/ehUXb/DTMHAgMBAAEwDQYJ
KoZIhvcNAQELBQADggEBAKij8eUTcs+AJFPnzc3aolVZEApwvLum58WRjmoev44A
1528F4dXF7vJhIbqdOvEBy0YNQhNuNUs+JiHIFwuVvhNuAXDgXJNsvymx8fn0E5U
C90iTCiV9WhlL93S6fSelDj65sgD4Gw8Q4bBbNa/SRCu4+oBNS9BPjpcbrGllph9
7AkCGBiaabVLqGNyZJEKZpRQ3kOqdQzHYT/eHRC3hcO/KGf0vCOUTgEhHuYavMy/
JZOeFg1owNP2nZ8cD2TwDKS+T+T1rAG1ovnVp/PV7lbH1o8Kn2rwtj1S42O824Gr
2NyVhhdZkLI/uEX9mdmcFPB+oV6iiPnqEh/r2wswFgw=
-----END CERTIFICATE-----`

// TestMsgUpdateIdentityGetSigners tests GetSigners for MsgUpdateIdentity
func TestValidateGenesis(t *testing.T) {
	id := Identity{
		Id: testIDStr,
		PubKeys: []PubKeyInfo{
			{PubKey: testPubKeySM2Str, Algorithm: SM2},
		},
		Certificates: []string{testCertificate},
		Credentials:  testCredentials,
		Owner:        testOwner.String(),
	}
	err := ValidateGenesis(GenesisState{[]Identity{id}})
	require.NoError(t, err)
}
