package rest

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gorilla/mux"
)

// Rest variable names
// nolint
const (
	RestID = "id"
)

// RegisterRoutes defines routes that get registered by the main application
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
	registerQueryRoutes(clientCtx, r)
	registerTxRoutes(clientCtx, r)
}
