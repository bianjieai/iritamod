package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	ParamsMsgTypeURL = "/iritamod.params.MsgUpdateParams"
	ParamsMsgSuffix  = "UpdateParams"
)

// validateMsgType validates the messages must have update param type but not that of params module
func validateMsgType(messages []sdk.Msg) error {
	for _, msg := range messages {
		msgURL := sdk.MsgTypeURL(msg)
		if msgURL == ParamsMsgTypeURL {
			return sdkerrors.Wrapf(ErrInvalidMsgType, "message %s must not be params messages of param module", msgURL)
		}
		if !strings.HasSuffix(msgURL, ParamsMsgSuffix) {
			return sdkerrors.Wrapf(ErrInvalidMsgType, "%s is not params message type", msgURL)
		}
	}
	return nil
}
