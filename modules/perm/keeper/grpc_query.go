package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/iritamod/modules/perm/types"
)

var _ types.QueryServer = Keeper{}

// Roles queries roles of a given address
func (k Keeper) Roles(
	c context.Context,
	req *types.QueryRolesRequest,
) (*types.QueryRolesResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(c)
	auth := k.GetAuth(ctx, addr)

	return &types.QueryRolesResponse{Roles: auth.Roles()}, nil
}

// AccountBlockList queries all blocked accounts
func (k Keeper) AccountBlockList(
	c context.Context,
	req *types.QueryBlockListRequest,
) (*types.QueryBlockListResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryBlockListResponse{Addresses: k.GetAllBlockAccounts(ctx)}, nil
}

// ContractDenyList queries all blocked contract
func (k Keeper) ContractDenyList(
	c context.Context,
	req *types.QueryContractDenyList,
) (*types.QueryContractDenyListResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	return &types.QueryContractDenyListResponse{Addresses: k.GetContractDenyList(ctx)}, nil

}
