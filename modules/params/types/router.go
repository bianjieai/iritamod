package types

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// ParamsRouter is a wrapper for MsgServiceRouter
type ParamsRouter struct {
	allowed map[string]struct{}
	*baseapp.MsgServiceRouter
}

func NewParamsRouter(msr *baseapp.MsgServiceRouter, msgTypeURLs []string) *ParamsRouter {
	if msr == nil {
		panic("msg service router is nil")
	}

	if len(msgTypeURLs) == 0 {
		panic("no allowed update params")
	}

	allowed := make(map[string]struct{}, len(msgTypeURLs))
	for _, msgTypeURL := range msgTypeURLs {
		if _, ok := allowed[msgTypeURL]; ok {
			panic(fmt.Sprintf("duplicate message type (%s)", msgTypeURL))
		}
		allowed[msgTypeURL] = struct{}{}
	}

	return &ParamsRouter{
		allowed:          allowed,
		MsgServiceRouter: msr,
	}
}

// Execute executes a params message.
func (pr *ParamsRouter) Execute(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
	msgType := sdk.MsgTypeURL(msg)
	if _, ok := pr.allowed[msgType]; !ok {
		return nil, sdkerrors.Wrapf(ErrInvalidMsgType, "%s is not allowed", sdk.MsgTypeURL(msg))
	}

	handler := pr.Handler(msg)
	if handler == nil {
		return nil, sdkerrors.Wrapf(ErrUnroutableUpdateParamsMsg, "%s is not registered", sdk.MsgTypeURL(msg))
	}

	return handler(ctx, msg)
}
