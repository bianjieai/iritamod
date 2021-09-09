package types

import (
	fmt "fmt"
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	cautil "github.com/bianjieai/iritamod/utils/ca"
)

const certStr = `-----BEGIN CERTIFICATE-----
MIIBazCCAR0CFGTwvE8oG+N3uNm1gonJBh6pie5TMAUGAytlcDBYMQswCQYDVQQG
EwJDTjENMAsGA1UECAwEcm9vdDENMAsGA1UEBwwEcm9vdDENMAsGA1UECgwEcm9v
dDENMAsGA1UECwwEcm9vdDENMAsGA1UEAwwEcm9vdDAeFw0yMDA2MTkwNzAyMzla
Fw0yMDA3MTkwNzAyMzlaMFgxCzAJBgNVBAYTAkNOMQ0wCwYDVQQIDAR0ZXN0MQ0w
CwYDVQQHDAR0ZXN0MQ0wCwYDVQQKDAR0ZXN0MQ0wCwYDVQQLDAR0ZXN0MQ0wCwYD
VQQDDAR0ZXN0MCowBQYDK2VwAyEA27WvK0goa1sSjsp6eb/xCkgjBEoPC9vfL/6h
f0hqjHYwBQYDK2VwA0EA0fo8y+saUl+8UiyKpKdjv2DsqYWqmqJDz9u3NaioOvrQ
Z0mOxdgj9wfO0t3voldCRUw3hCekjC+GEOoXH5ysDQ==
-----END CERTIFICATE-----`

var (
	cert, _ = cautil.ReadCertificateFromMem([]byte(certStr))
	pk, _   = cautil.GetPubkeyFromCert(cert)
	addr    = pk.Address()
	accAddr = sdk.AccAddress(addr)

	nodeID   = addr
	nodeName = "test_node"

	emptyAddr sdk.AccAddress
	emptyCert = ""
)

// test ValidateBasic for MsgCreateValidator
func TestMsgCreateValidator(t *testing.T) {
	testMsgs := []*MsgCreateValidator{
		NewMsgCreateValidator("a", "b", certStr, 1, accAddr),
		NewMsgCreateValidator("a", "b", certStr, 1, emptyAddr),
		NewMsgCreateValidator("", "b", certStr, 1, accAddr),
		NewMsgCreateValidator("  ", "b", certStr, 1, accAddr),
		NewMsgCreateValidator("a", "b", emptyCert, 1, accAddr),
		NewMsgCreateValidator("a", "b", certStr, -1, accAddr),
		NewMsgCreateValidator("a", "b", certStr, 0, accAddr),
	}

	testCases := []struct {
		msg     *MsgCreateValidator
		expPass bool
		errMsg  string
	}{
		{testMsgs[0], true, ""},
		{testMsgs[1], false, "missing operator"},
		{testMsgs[2], false, "missing name"},
		{testMsgs[3], false, "blank name"},
		{testMsgs[4], false, "missing pubkey"},
		{testMsgs[5], false, "negative power"},
		{testMsgs[6], false, "zero power"},
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

// test ValidateBasic for MsgUpdateValidator
func TestMsgUpdateValidator(t *testing.T) {
	testMsgs := []*MsgUpdateValidator{
		NewMsgUpdateValidator([]byte("a"), "b", "b", certStr, 1, accAddr),
		NewMsgUpdateValidator([]byte("a"), "b", "b", certStr, 1, emptyAddr),
		NewMsgUpdateValidator([]byte{}, "b", "b", certStr, 1, accAddr),
		NewMsgUpdateValidator([]byte("a"), "b", "b", certStr, -1, accAddr),
	}

	testCases := []struct {
		msg     *MsgUpdateValidator
		expPass bool
		errMsg  string
	}{
		{testMsgs[0], true, ""},
		{testMsgs[1], false, "missing operator"},
		{testMsgs[2], false, "missing name"},
		{testMsgs[3], false, "negative power"},
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

// test ValidateBasic for MsgRemoveValidator
func TestMsgRemoveValidator(t *testing.T) {
	testMsgs := []*MsgRemoveValidator{
		NewMsgRemoveValidator([]byte("a"), accAddr),
		NewMsgRemoveValidator([]byte("a"), emptyAddr),
		NewMsgRemoveValidator([]byte{}, accAddr),
	}

	testCases := []struct {
		msg     *MsgRemoveValidator
		expPass bool
		errMsg  string
	}{
		{testMsgs[0], true, ""},
		{testMsgs[1], false, "missing operator"},
		{testMsgs[2], false, "missing id"},
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

// TestMsgGrantNodeRoute tests Route for MsgGrantNode
func TestMsgGrantNodeRoute(t *testing.T) {
	msg := NewMsgGrantNode(nodeName, certStr, accAddr)

	require.Equal(t, RouterKey, msg.Route())
}

// TestMsgGrantNode tests Type for MsgGrantNode
func TestMsgGrantNodeType(t *testing.T) {
	msg := NewMsgGrantNode(nodeName, certStr, accAddr)

	require.Equal(t, "grant_node", msg.Type())
}

// TestMsgGrantNodeValidation tests ValidateBasic for MsgGrantNode
func TestMsgGrantNodeValidation(t *testing.T) {
	invalidCertificate := "invalidCertificate"

	testMsgs := []*MsgGrantNode{
		NewMsgGrantNode(nodeName, certStr, accAddr),            // valid msg
		NewMsgGrantNode(nodeName, certStr, emptyAddr),          // missing operator address
		NewMsgGrantNode("", certStr, accAddr),                  // name can not be empty
		NewMsgGrantNode(nodeName, "", accAddr),                 // missing certificate
		NewMsgGrantNode(nodeName, invalidCertificate, accAddr), // invalid certificate
	}

	testCases := []struct {
		msg     *MsgGrantNode
		expPass bool
		errMsg  string
	}{
		{testMsgs[0], true, ""},
		{testMsgs[1], false, "missing operator address"},
		{testMsgs[2], false, "name can not be empty"},
		{testMsgs[3], false, "missing certificate"},
		{testMsgs[4], false, "invalid certificate"},
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

// TestMsgGrantNodeGetSignBytes tests GetSignBytes for MsgGrantNode
func TestMsgGrantNodeGetSignBytes(t *testing.T) {
	msg := NewMsgGrantNode(nodeName, certStr, accAddr)
	res := msg.GetSignBytes()

	expected := `{"type":"iritamod/node/MsgGrantNode","value":{"certificate":"-----BEGIN CERTIFICATE-----\nMIIBazCCAR0CFGTwvE8oG+N3uNm1gonJBh6pie5TMAUGAytlcDBYMQswCQYDVQQG\nEwJDTjENMAsGA1UECAwEcm9vdDENMAsGA1UEBwwEcm9vdDENMAsGA1UECgwEcm9v\ndDENMAsGA1UECwwEcm9vdDENMAsGA1UEAwwEcm9vdDAeFw0yMDA2MTkwNzAyMzla\nFw0yMDA3MTkwNzAyMzlaMFgxCzAJBgNVBAYTAkNOMQ0wCwYDVQQIDAR0ZXN0MQ0w\nCwYDVQQHDAR0ZXN0MQ0wCwYDVQQKDAR0ZXN0MQ0wCwYDVQQLDAR0ZXN0MQ0wCwYD\nVQQDDAR0ZXN0MCowBQYDK2VwAyEA27WvK0goa1sSjsp6eb/xCkgjBEoPC9vfL/6h\nf0hqjHYwBQYDK2VwA0EA0fo8y+saUl+8UiyKpKdjv2DsqYWqmqJDz9u3NaioOvrQ\nZ0mOxdgj9wfO0t3voldCRUw3hCekjC+GEOoXH5ysDQ==\n-----END CERTIFICATE-----","name":"test_node","operator":"cosmos1z0hd0wjhlsl2jj33439ppy7u2crvlwyq8qedsm"}}`
	require.Equal(t, expected, string(res))
}

// TestMsgGrantNodeGetSigners tests GetSigners for MsgGrantNode
func TestMsgGrantNodeGetSigners(t *testing.T) {
	msg := NewMsgGrantNode(nodeName, certStr, accAddr)
	res := msg.GetSigners()

	expected := "[13EED7BA57FC3EA94A31AC4A1093DC5606CFB880]"
	require.Equal(t, expected, fmt.Sprintf("%v", res))
}

// TestMsgRevokeNodeRoute tests Route for MsgRevokeNode
func TestMsgRevokeNodeRoute(t *testing.T) {
	msg := NewMsgRevokeNode(nodeID, accAddr)

	require.Equal(t, RouterKey, msg.Route())
}

// TestMsgRevokeNodeType tests Type for MsgRevokeNode
func TestMsgRevokeNodeType(t *testing.T) {
	msg := NewMsgRevokeNode(nodeID, accAddr)

	require.Equal(t, "revoke_node", msg.Type())
}

// TestMsgRevokeNodeValidation tests ValidateBasic for MsgRevokeNode
func TestMsgRevokeNodeValidation(t *testing.T) {
	testMsgs := []*MsgRevokeNode{
		NewMsgRevokeNode(nodeID, accAddr),   // valid msg
		NewMsgRevokeNode(nodeID, emptyAddr), // missing operator address
		NewMsgRevokeNode(nil, accAddr),      // missing node ID
	}

	testCases := []struct {
		msg     *MsgRevokeNode
		expPass bool
		errMsg  string
	}{
		{testMsgs[0], true, ""},
		{testMsgs[1], false, "missing operator address"},
		{testMsgs[2], false, "missing node ID"},
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

// TestMsgRevokeNodeGetSignBytes tests GetSignBytes for MsgRevokeNode
func TestMsgRevokeNodeGetSignBytes(t *testing.T) {
	msg := NewMsgRevokeNode(nodeID, accAddr)
	res := msg.GetSignBytes()

	expected := `{"type":"iritamod/node/MsgRevokeNode","value":{"id":"13EED7BA57FC3EA94A31AC4A1093DC5606CFB880","operator":"cosmos1z0hd0wjhlsl2jj33439ppy7u2crvlwyq8qedsm"}}`
	require.Equal(t, expected, string(res))
}

// TestMsgRevokeNodeGetSigners tests GetSigners for MsgRevokeNode
func TestMsgRevokeNodeGetSigners(t *testing.T) {
	msg := NewMsgRevokeNode(nodeID, accAddr)
	res := msg.GetSigners()

	expected := "[13EED7BA57FC3EA94A31AC4A1093DC5606CFB880]"
	require.Equal(t, expected, fmt.Sprintf("%v", res))
}
