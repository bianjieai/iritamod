package keeper

import (
	"context"
	"encoding/hex"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/iritamod/modules/node/types"
)

var _ types.QueryServer = Keeper{}

// Node queries a node by id
func (k Keeper) Node(c context.Context, req *types.QueryNodeRequest) (*types.QueryNodeResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	err := types.ValidateNodeID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	id, _ := hex.DecodeString(req.Id)

	ctx := sdk.UnwrapSDKContext(c)

	node, found := k.GetNode(ctx, id)
	if !found {
		return nil, status.Errorf(codes.NotFound, "node %s not found", req.Id)
	}

	return &types.QueryNodeResponse{Node: &node}, nil
}

// Nodes queries all nodes
func (k Keeper) Nodes(c context.Context, req *types.QueryNodesRequest) (*types.QueryNodesResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	nodes := make([]types.Node, 0)

	k.IterateNodes(ctx, func(node types.Node) (stop bool) {
		nodes = append(nodes, node)
		return false
	})

	return &types.QueryNodesResponse{Nodes: nodes}, nil
}
