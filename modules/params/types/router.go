package types

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ParamsRouter is a wrapper for MsgServiceRouter
type ParamsRouter interface {
	// IsParamsType returns true if the msg type is an update params type.
	IsParamsType(msg sdk.Msg) bool

	// Handler returns true if the router supports a given msg type. Handler
	// must use isParamsType to check if the msg type is an update params type.
	Handler(msg sdk.Msg) (baseapp.MsgServiceHandler, bool)
}
