package types

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/tendermint/crypto/sm2"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	pk      = sm2.GenPrivKey().PubKey()
	addr    = pk.Address()
	accAddr = sdk.AccAddress(addr)

	emptyAddr sdk.AccAddress
)

// test ValidateBasic for MsgUnjailValidator
func TestMsgUnjail(t *testing.T) {
	testMsgs := []*MsgUnjailValidator{
		NewMsgUnjailValidator([]byte("a"), accAddr),
		NewMsgUnjailValidator([]byte("a"), emptyAddr),
		NewMsgUnjailValidator([]byte(""), accAddr),
	}

	testCases := []struct {
		msg     *MsgUnjailValidator
		expPass bool
		errMsg  string
	}{
		{testMsgs[0], true, ""},
		{testMsgs[1], false, "missing operator"},
		{testMsgs[2], false, "missing name"},
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
