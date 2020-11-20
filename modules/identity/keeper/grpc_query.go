package keeper

import (
	"context"
	"encoding/hex"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/bianjieai/iritamod/modules/identity/types"
)

var _ types.QueryServer = Keeper{}

// Identity queries an identity by id
func (k Keeper) Identity(c context.Context, req *types.QueryIdentityRequest) (*types.QueryIdentityResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	id, err := hex.DecodeString(req.Id)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidID, req.Id)
	}
	identity, found := k.GetIdentity(ctx, id)
	if !found {
		return nil, status.Errorf(codes.NotFound, "identity %s not found", req.Id)
	}

	return &types.QueryIdentityResponse{Identity: &identity}, nil
}
