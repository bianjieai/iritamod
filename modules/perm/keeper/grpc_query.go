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
func (k Keeper) Roles(c context.Context, req *types.QueryRolesRequest) (*types.QueryRolesResponse, error) {
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

// Blacklist queries all blocked accounts
func (k Keeper) Blacklist(c context.Context, req *types.QueryBlacklistRequest) (*types.QueryBlacklistResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryBlacklistResponse{Addresses: k.GetAllBlackAccounts(ctx)}, nil
}
