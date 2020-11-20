package keeper

import (
	"encoding/hex"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/bianjieai/iritamod/modules/identity/types"
)

// NewQuerier creates a new identity Querier instance
func NewQuerier(keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryIdentity:
			return queryIdentity(ctx, req, keeper, legacyQuerierCdc)

		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown %s query path: %s", types.ModuleName, path[0])
		}
	}
}

func queryIdentity(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryIdentityParams
	if err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	id, err := hex.DecodeString(params.ID)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidID, params.ID)
	}

	identity, found := k.GetIdentity(ctx, id)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrUnknownIdentity, params.ID)
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, identity)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}
