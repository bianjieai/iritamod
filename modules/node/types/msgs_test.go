package types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	cautils "github.com/bianjieai/iritamod/utils/ca"
)

var (
	cert, _   = cautils.ReadCertificateFromMem([]byte(testCertificate))
	pubKey, _ = cautils.GetPubkeyFromCert(cert)

	testID       = pubKey.Address()
	testOperator = sdk.AccAddress([]byte("test-operatoraddress"))
)

// TestMsgAddNodeRoute tests Route for MsgAddNode
func TestMsgAddNodeRoute(t *testing.T) {
	msg := NewMsgAddNode(testCertificate, testOperator)

	require.Equal(t, RouterKey, msg.Route())
}

// TestMsgAddNode tests Type for MsgAddNode
func TestMsgAddNodeType(t *testing.T) {
	msg := NewMsgAddNode(testCertificate, testOperator)

	require.Equal(t, "add_node", msg.Type())
}

// TestMsgAddNodeValidation tests ValidateBasic for MsgAddNode
func TestMsgAddNodeValidation(t *testing.T) {
	emptyAddress := sdk.AccAddress{}
	invalidCertificate := "invalidCertificate"

	testMsgs := []*MsgAddNode{
		NewMsgAddNode(testCertificate, testOperator),    // valid msg
		NewMsgAddNode(testCertificate, emptyAddress),    // missing operator address
		NewMsgAddNode("", testOperator),                 // missing certificate
		NewMsgAddNode(invalidCertificate, testOperator), // invalid certificate
	}

	testCases := []struct {
		msg     *MsgAddNode
		expPass bool
		errMsg  string
	}{
		{testMsgs[0], true, ""},
		{testMsgs[1], false, "missing operator address"},
		{testMsgs[2], false, "missing certificate"},
		{testMsgs[3], false, "invalid certificate"},
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

// TestMsgAddNodeGetSignBytes tests GetSignBytes for MsgAddNode
func TestMsgAddNodeGetSignBytes(t *testing.T) {
	msg := NewMsgAddNode(testCertificate, testOperator)
	res := msg.GetSignBytes()

	expected := fmt.Sprintf(`{"type":"iritamod/node/MsgAddNode","value":{"certificate":"-----BEGIN CERTIFICATE-----\nMIIBazCCAR0CFGTwvE8oG+N3uNm1gonJBh6pie5TMAUGAytlcDBYMQswCQYDVQQG\nEwJDTjENMAsGA1UECAwEcm9vdDENMAsGA1UEBwwEcm9vdDENMAsGA1UECgwEcm9v\ndDENMAsGA1UECwwEcm9vdDENMAsGA1UEAwwEcm9vdDAeFw0yMDA2MTkwNzAyMzla\nFw0yMDA3MTkwNzAyMzlaMFgxCzAJBgNVBAYTAkNOMQ0wCwYDVQQIDAR0ZXN0MQ0w\nCwYDVQQHDAR0ZXN0MQ0wCwYDVQQKDAR0ZXN0MQ0wCwYDVQQLDAR0ZXN0MQ0wCwYD\nVQQDDAR0ZXN0MCowBQYDK2VwAyEA27WvK0goa1sSjsp6eb/xCkgjBEoPC9vfL/6h\nf0hqjHYwBQYDK2VwA0EA0fo8y+saUl+8UiyKpKdjv2DsqYWqmqJDz9u3NaioOvrQ\nZ0mOxdgj9wfO0t3voldCRUw3hCekjC+GEOoXH5ysDQ==\n-----END CERTIFICATE-----","operator":"cosmos1w3jhxapddacx2unpw3hhyctyv3ex2umnxw5yuk"}}`)
	require.Equal(t, expected, string(res))
}

// TestMsgAddNodeGetSigners tests GetSigners for MsgAddNode
func TestMsgAddNodeGetSigners(t *testing.T) {
	msg := NewMsgAddNode(testCertificate, testOperator)
	res := msg.GetSigners()

	expected := "[746573742D6F70657261746F7261646472657373]"
	require.Equal(t, expected, fmt.Sprintf("%v", res))
}

// TestMsgRemoveNodeRoute tests Route for MsgRemoveNode
func TestMsgRemoveNodeRoute(t *testing.T) {
	msg := NewMsgRemoveNode(testID, testOperator)

	require.Equal(t, RouterKey, msg.Route())
}

// TestMsgRemoveNodeType tests Type for MsgRemoveNode
func TestMsgRemoveNodeType(t *testing.T) {
	msg := NewMsgRemoveNode(testID, testOperator)

	require.Equal(t, "remove_node", msg.Type())
}

// TestMsgRemoveNodeValidation tests ValidateBasic for MsgRemoveNode
func TestMsgRemoveNodeValidation(t *testing.T) {
	emptyAddress := sdk.AccAddress{}
	invalidID := []byte("invalidID")

	testMsgs := []*MsgRemoveNode{
		NewMsgRemoveNode(testID, testOperator),    // valid msg
		NewMsgRemoveNode(testID, emptyAddress),    // missing operator address
		NewMsgRemoveNode(nil, testOperator),       // missing node ID
		NewMsgRemoveNode(invalidID, testOperator), // invalid node ID
	}

	testCases := []struct {
		msg     *MsgRemoveNode
		expPass bool
		errMsg  string
	}{
		{testMsgs[0], true, ""},
		{testMsgs[1], false, "missing operator address"},
		{testMsgs[2], false, "missing node ID"},
		{testMsgs[3], false, "invalid node ID"},
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

// TestMsgRemoveNodeGetSignBytes tests GetSignBytes for MsgRemoveNode
func TestMsgRemoveNodeGetSignBytes(t *testing.T) {
	msg := NewMsgRemoveNode(testID, testOperator)
	res := msg.GetSignBytes()

	expected := fmt.Sprintf(`{"type":"iritamod/node/MsgRemoveNode","value":{"id":"13EED7BA57FC3EA94A31AC4A1093DC5606CFB880","operator":"cosmos1w3jhxapddacx2unpw3hhyctyv3ex2umnxw5yuk"}}`)
	require.Equal(t, expected, string(res))
}

// TestMsgRemoveNodeGetSigners tests GetSigners for MsgRemoveNode
func TestMsgRemoveNodeGetSigners(t *testing.T) {
	msg := NewMsgRemoveNode(testID, testOperator)
	res := msg.GetSigners()

	expected := "[746573742D6F70657261746F7261646472657373]"
	require.Equal(t, expected, fmt.Sprintf("%v", res))
}

const testCertificate = `-----BEGIN CERTIFICATE-----
MIIBazCCAR0CFGTwvE8oG+N3uNm1gonJBh6pie5TMAUGAytlcDBYMQswCQYDVQQG
EwJDTjENMAsGA1UECAwEcm9vdDENMAsGA1UEBwwEcm9vdDENMAsGA1UECgwEcm9v
dDENMAsGA1UECwwEcm9vdDENMAsGA1UEAwwEcm9vdDAeFw0yMDA2MTkwNzAyMzla
Fw0yMDA3MTkwNzAyMzlaMFgxCzAJBgNVBAYTAkNOMQ0wCwYDVQQIDAR0ZXN0MQ0w
CwYDVQQHDAR0ZXN0MQ0wCwYDVQQKDAR0ZXN0MQ0wCwYDVQQLDAR0ZXN0MQ0wCwYD
VQQDDAR0ZXN0MCowBQYDK2VwAyEA27WvK0goa1sSjsp6eb/xCkgjBEoPC9vfL/6h
f0hqjHYwBQYDK2VwA0EA0fo8y+saUl+8UiyKpKdjv2DsqYWqmqJDz9u3NaioOvrQ
Z0mOxdgj9wfO0t3voldCRUw3hCekjC+GEOoXH5ysDQ==
-----END CERTIFICATE-----`
