package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"gitlab.bianjie.ai/irita-pro/iritamod/modules/identity/types"
)

var _ types.QueryServer = Keeper{}

// Identity queries an identity by id
func (k Keeper) Identity(c context.Context, req *types.QueryIdentityRequest) (*types.QueryIdentityResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	identity, found := k.GetIdentity(ctx, req.Id)
	if !found {
		return nil, status.Errorf(codes.NotFound, "identity %s not found", req.Id)
	}

	return &types.QueryIdentityResponse{Identity: &identity}, nil
}
