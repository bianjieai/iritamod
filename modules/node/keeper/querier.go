package keeper

import (
	"encoding/hex"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/bianjieai/iritamod/modules/node/types"
)

// NewQuerier creates a new node Querier instance
func NewQuerier(keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryNode:
			return queryNode(ctx, req, keeper, legacyQuerierCdc)

		case types.QueryNodes:
			return queryNodes(ctx, req, keeper, legacyQuerierCdc)

		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown %s query path: %s", types.ModuleName, path[0])
		}
	}
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

	nodes := make([]types.Node, 0)

	k.IterateNodes(ctx, func(node types.Node) (stop bool) {
		nodes = append(nodes, node)
		return false
	})

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, nodes)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}
