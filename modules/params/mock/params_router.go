package mock

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"

	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"

	nodetypes "github.com/bianjieai/iritamod/modules/node/types"
	opbtypes "github.com/bianjieai/iritamod/modules/opb/types"
	paramstypes "github.com/bianjieai/iritamod/modules/params/types"
	slashingtypes "github.com/bianjieai/iritamod/modules/slashing/types"
)

var _ paramstypes.ParamsRouter = (*ParamsRouter)(nil)

// ParamsRouter is a mock implementation of ParamsRouter interface
type ParamsRouter struct {
	router *baseapp.MsgServiceRouter
}

func NewParamsRouter(msr *baseapp.MsgServiceRouter) *ParamsRouter {
	return &ParamsRouter{
		router: msr,
	}
}

func (pr *ParamsRouter) IsParamsType(msg sdk.Msg) bool {
	switch msg.(type) {
	case *opbtypes.MsgUpdateParams:
		return true
	case *nodetypes.MsgUpdateParams:
		return true
	case *slashingtypes.MsgUpdateParams:
		return true
	case *minttypes.MsgUpdateParams:
		return true
	}
	return false
}

func (pr *ParamsRouter) Handler(msg sdk.Msg) (baseapp.MsgServiceHandler, bool) {
	if !pr.IsParamsType(msg) {
		return nil, false
	}

	handler := pr.router.Handler(msg)
	return handler, true
}
