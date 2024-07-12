package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
)

// RegisterRoutes registers staking-related REST handlers to a router
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
	registerQueryRoutes(clientCtx, r)
}
