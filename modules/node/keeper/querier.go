package keeper

import (
	"encoding/hex"
	"fmt"
	"strconv"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/bianjieai/iritamod/modules/node/types"
)

const DefaultLimit = 100

// creates a querier for validator REST endpoints
func NewQuerier(keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryValidators:
			return queryValidators(ctx, req, keeper, legacyQuerierCdc)

		case types.QueryValidator:
			return queryValidator(ctx, req, keeper, legacyQuerierCdc)

		case types.QueryNode:
			return queryNode(ctx, req, keeper, legacyQuerierCdc)

		case types.QueryNodes:
			return queryNodes(ctx, req, keeper, legacyQuerierCdc)

		case types.QueryParameters:
			return queryParameters(ctx, keeper, legacyQuerierCdc)

		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown %s query endpoint: %s", types.ModuleName, path[0])
		}
	}
}

func queryValidators(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryValidatorsParams

	if err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	validators := k.GetAllValidators(ctx)
	var filterValidators []types.Validator

	if len(params.Jailed) > 0 {
		jailed, err := strconv.ParseBool(params.Jailed)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
		}

		i := 0
		for _, val := range validators {
			if jailed == val.Jailed {
				validators[i] = val
				i++
			}
		}
		filterValidators = validators[:i]
	} else {
		filterValidators = validators
	}

	start, end := client.Paginate(len(filterValidators), params.Page, params.Limit, DefaultLimit)
	if start < 0 || end < 0 {
		filterValidators = []types.Validator{}
	} else {
		filterValidators = filterValidators[start:end]
	}

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, filterValidators)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryValidator(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryValidatorParams

	if err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	id, err := hex.DecodeString(params.ID)
	if err != nil {
		return nil, fmt.Errorf("invalid validator id:%s", params.ID)
	}

	validator, found := k.GetValidator(ctx, id)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrUnknownValidator, params.ID)
	}

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, validator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryNode(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryNodeParams
	if err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	id, err := hex.DecodeString(params.ID)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidNodeID, params.ID)
	}

	node, found := k.GetNode(ctx, id)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrUnknownNode, params.ID)
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, node)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func queryNodes(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryNodesParams
	if err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	nodes := k.GetNodes(ctx)

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, nodes)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func queryParameters(ctx sdk.Context, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	params := k.GetParams(ctx)

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}
