package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidMsgType            = sdkerrors.Register(ModuleName, 2, "invalid message type")
	ErrUnroutableUpdateParamsMsg = sdkerrors.Register(ModuleName, 3, "update params message not recognized by router")
)
