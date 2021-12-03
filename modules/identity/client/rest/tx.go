package rest

import (
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/bianjieai/iritamod/modules/identity/types"
)

func registerTxRoutes(clientCtx client.Context, r *mux.Router) {
	// create an identity
	r.HandleFunc("/identity/identities", createIdentityHandlerFn(clientCtx)).Methods("POST")
	// update an identity
	r.HandleFunc(fmt.Sprintf("/identity/identities/{%s}", RestID), updateIdentityHandlerFn(clientCtx)).Methods("PUT")
}

// CreateIdentityReq defines the properties of an identity creation request's body.
type CreateIdentityReq struct {
	BaseReq     rest.BaseReq   `json:"base_req" yaml:"base_req"`
	ID          string         `json:"id" yaml:"id"`
	PubKey      string         `json:"pubkey" yaml:"pubkey"`
	PubKeyAlgo  string         `json:"pubkey_algo" yaml:"pubkey_algo"`
	Certificate string         `json:"certificate" yaml:"certificate"`
	Credentials string         `json:"credentials" yaml:"credentials"`
	Data        string         `json:"data" yaml:"data"`
	Owner       sdk.AccAddress `json:"owner" yaml:"owner"`
}

// UpdateIdentityReq defines the properties of an identity update request's body.
type UpdateIdentityReq struct {
	BaseReq     rest.BaseReq   `json:"base_req" yaml:"base_req"`
	PubKey      string         `json:"pubkey" yaml:"pubkey"`
	PubKeyAlgo  string         `json:"pubkey_algo" yaml:"pubkey_algo"`
	Certificate string         `json:"certificate" yaml:"certificate"`
	Credentials string         `json:"credentials" yaml:"credentials"`
	Data        string         `json:"data" yaml:"data"`
	Owner       sdk.AccAddress `json:"owner" yaml:"owner"`
}

func createIdentityHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateIdentityReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		var id []byte
		var err error

		if len(req.ID) > 0 {
			id, err = hex.DecodeString(req.ID)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}
		} else {
			id = uuid.NewV4().Bytes()
		}

		var pubKeyInfo *types.PubKeyInfo

		if len(req.PubKey) > 0 {
			pubKeyInfo = new(types.PubKeyInfo)
			pubKeyInfo.PubKey = req.PubKey
			pubKeyInfo.Algorithm = types.PubKeyAlgorithmFromString(req.PubKeyAlgo)
		}

		msg := types.NewMsgCreateIdentity(id, pubKeyInfo, req.Certificate, req.Credentials, req.Owner, req.Data)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

func updateIdentityHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := hex.DecodeString(vars[RestID])
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		var req UpdateIdentityReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		var pubKeyInfo *types.PubKeyInfo

		if len(req.PubKey) > 0 {
			pubKeyInfo = new(types.PubKeyInfo)
			pubKeyInfo.PubKey = req.PubKey
			pubKeyInfo.Algorithm = types.PubKeyAlgorithmFromString(req.PubKeyAlgo)
		}

		msg := types.NewMsgUpdateIdentity(id, pubKeyInfo, req.Certificate, req.Credentials, req.Owner, req.Data)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
