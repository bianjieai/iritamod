package keeper

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"

	"gitlab.bianjie.ai/irita-pro/iritamod/modules/admin/types"
)

const DefaultLimit = 100

// creates a querier for admin REST endpoints
func NewQuerier(keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryRoles:
			return queryRoles(ctx, req, keeper, legacyQuerierCdc)

		case types.QueryBlacklist:
			return queryBlacklist(ctx, req, keeper, legacyQuerierCdc)

		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown %s query endpoint: %s", types.ModuleName, path[0])
		}
	}
}

func queryRoles(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryRolesParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	auth := k.GetAuth(ctx, params.Address)

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, auth.Roles())
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryBlacklist(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryBlacklistParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	blackAccounts := k.GetAllBlackAccounts(ctx)

	start, end := client.Paginate(len(blackAccounts), params.Page, params.Limit, DefaultLimit)
	if start < 0 || end < 0 {
		blackAccounts = []sdk.AccAddress{}
	} else {
		blackAccounts = blackAccounts[start:end]
	}

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, blackAccounts)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}
