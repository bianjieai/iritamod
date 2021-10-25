package keeper

import (
	"context"
	"encoding/hex"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/bianjieai/iritamod/modules/node/types"
)

// Querier is used as Keeper will have duplicate methods if used directly, and gRPC names take precedence over keeper
type Querier struct {
	Keeper
}

var _ types.QueryServer = Querier{}

// Validator queries the validator by the given id
func (q Querier) Validator(c context.Context, req *types.QueryValidatorRequest) (*types.QueryValidatorResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	id, err := hex.DecodeString(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid id")
	}

	ctx := sdk.UnwrapSDKContext(c)
	validator, found := q.GetValidator(ctx, id)
	if !found {
		return nil, status.Errorf(codes.NotFound, "validator %s not found", req.Id)
	}

	return &types.QueryValidatorResponse{Validator: &validator}, nil
}

// Validators queries the validators
func (q Querier) Validators(c context.Context, req *types.QueryValidatorsRequest) (*types.QueryValidatorsResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	store := ctx.KVStore(q.storeKey)

	var validators []types.Validator

	validatorStore := prefix.NewStore(store, types.ValidatorsKey)
	pageRes, err := query.Paginate(
		validatorStore,
		req.Pagination,
		func(key []byte, value []byte) error {
			var validator types.Validator
			if err := q.cdc.Unmarshal(value, &validator); err != nil {
				return err
			}

			validators = append(validators, validator)
			return nil
		},
	)
	if err != nil {
		return nil, err
	}

	return &types.QueryValidatorsResponse{
		Validators: validators,
		Pagination: pageRes,
	}, nil
}

// Params queries the parameters of the node module
func (q Querier) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	params := q.GetParams(ctx)

	return &types.QueryParamsResponse{Params: params}, nil
}

// Node queries a node by id
func (q Querier) Node(c context.Context, req *types.QueryNodeRequest) (*types.QueryNodeResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	err := types.ValidateNodeID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	id, _ := hex.DecodeString(req.Id)

	ctx := sdk.UnwrapSDKContext(c)

	node, found := q.GetNode(ctx, id)
	if !found {
		return nil, status.Errorf(codes.NotFound, "node %s not found", req.Id)
	}

	return &types.QueryNodeResponse{Node: &node}, nil
}

// Nodes queries all nodes
func (q Querier) Nodes(c context.Context, req *types.QueryNodesRequest) (*types.QueryNodesResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	nodes := make([]types.Node, 0)
	store := ctx.KVStore(q.storeKey)
	nodeStore := prefix.NewStore(store, types.NodeKey)
	pageRes, err := query.Paginate(nodeStore, req.Pagination, func(key []byte, value []byte) error {
		var node types.Node
		err := q.cdc.Unmarshal(value, &node)
		if err != nil {
			return err
		}
		nodes = append(nodes, node)
		return nil
	})
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "paginate: %v", err)
	}

	return &types.QueryNodesResponse{Nodes: nodes, Pagination: pageRes}, nil
}
