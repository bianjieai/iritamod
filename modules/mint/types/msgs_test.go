package types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/tendermint/crypto/tmhash"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	testAddress = sdk.AccAddress(tmhash.SumTruncated([]byte("test-address")))
	testAmount  = uint64(1000)

	emptyAddress = sdk.AccAddress{}
)

// TestMsgMintRoute tests Route for MsgMint
func TestMsgMintRoute(t *testing.T) {
	msg := NewMsgMint(testAmount, testAddress)
	require.Equal(t, RouterKey, msg.Route())
}

// TestMsgMintType tests Type for MsgMint
func TestMsgMintType(t *testing.T) {
	msg := NewMsgMint(testAmount, testAddress)
	require.Equal(t, TypeMsgMint, msg.Type())
}

// TestMsgMintValidation tests ValidateBasic for MsgMint
func TestMsgMintValidation(t *testing.T) {
	testMsgs := []*MsgMint{
		NewMsgMint(testAmount, testAddress),  // valid msg
		NewMsgMint(0, testAddress),           // amount must be greater than 0
		NewMsgMint(testAmount, emptyAddress), // missing operator address
	}

	testCases := []struct {
		msg     *MsgMint
		expPass bool
		errMsg  string
	}{
		{testMsgs[0], true, ""},
		{testMsgs[1], false, "amount must be greater than 0"},
		{testMsgs[2], false, "missing operator address"},
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

// TestMsgMintGetSignBytes tests GetSignBytes for MsgMint
func TestMsgMintGetSignBytes(t *testing.T) {
	msg := NewMsgMint(testAmount, testAddress)
	res := msg.GetSignBytes()

	expected := `{"type":"iritamod/mint/MsgMint","value":{"amount":"1000","operator":"cosmos1hjppmlx4fgtnpsya0pzqyg7el9qrq5lw58dd9x"}}`
	require.Equal(t, expected, string(res))
}

// TestMsgMintGetSigners tests GetSigners for MsgMint
func TestMsgMintGetSigners(t *testing.T) {
	msg := NewMsgMint(testAmount, testAddress)
	res := msg.GetSigners()

	expected := "[BC821DFCD54A1730C09D78440223D9F9403053EE]"
	require.Equal(t, expected, fmt.Sprintf("%v", res))
}
