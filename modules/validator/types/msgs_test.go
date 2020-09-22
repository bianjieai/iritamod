package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	cautil "gitlab.bianjie.ai/irita-pro/iritamod/utils/ca"
)

var (
	certStr = `-----BEGIN CERTIFICATE-----
MIIBazCCAR0CFGTwvE8oG+N3uNm1gonJBh6pie5TMAUGAytlcDBYMQswCQYDVQQG
EwJDTjENMAsGA1UECAwEcm9vdDENMAsGA1UEBwwEcm9vdDENMAsGA1UECgwEcm9v
dDENMAsGA1UECwwEcm9vdDENMAsGA1UEAwwEcm9vdDAeFw0yMDA2MTkwNzAyMzla
Fw0yMDA3MTkwNzAyMzlaMFgxCzAJBgNVBAYTAkNOMQ0wCwYDVQQIDAR0ZXN0MQ0w
CwYDVQQHDAR0ZXN0MQ0wCwYDVQQKDAR0ZXN0MQ0wCwYDVQQLDAR0ZXN0MQ0wCwYD
VQQDDAR0ZXN0MCowBQYDK2VwAyEA27WvK0goa1sSjsp6eb/xCkgjBEoPC9vfL/6h
f0hqjHYwBQYDK2VwA0EA0fo8y+saUl+8UiyKpKdjv2DsqYWqmqJDz9u3NaioOvrQ
Z0mOxdgj9wfO0t3voldCRUw3hCekjC+GEOoXH5ysDQ==
-----END CERTIFICATE-----`
	cert, _ = cautil.ReadCertificateFromMem([]byte(certStr))
	pk, _   = cautil.GetPubkeyFromCert(cert)
	addr    = pk.Address()
	accAddr = sdk.AccAddress(addr)

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
