package rest

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/bianjieai/iritamod/modules/identity/types"
)

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	// query an identity
	r.HandleFunc(fmt.Sprintf("/identity/identities/{%s}", RestID), queryIdentityHandlerFn(clientCtx)).Methods("GET")
}

func queryIdentityHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		clientCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, clientCtx, r)
		if !ok {
			return
		}

		params := types.QueryIdentityParams{
			ID: vars[RestID],
		}

		bz, err := clientCtx.LegacyAmino.MarshalJSON(params)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryIdentity)
		res, height, err := clientCtx.QueryWithData(route, bz)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		clientCtx = clientCtx.WithHeight(height)
		rest.PostProcessResponse(w, clientCtx, res)
	}
}
