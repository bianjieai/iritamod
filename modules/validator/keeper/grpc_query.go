package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"gitlab.bianjie.ai/irita-pro/iritamod/modules/validator/types"
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

	ctx := sdk.UnwrapSDKContext(c)

	validator, found := q.GetValidator(ctx, req.Id)
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
	pageRes, err := query.Paginate(validatorStore, req.Pagination, func(key []byte, value []byte) error {
		var validator types.Validator
		err := q.cdc.UnmarshalBinaryBare(value, &validator)
		if err != nil {
			return err
		}

		validators = append(validators, validator)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &types.QueryValidatorsResponse{Validators: validators, Pagination: pageRes}, nil
}

// Params queries the parameters of the validator module
func (q Querier) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	params := q.GetParams(ctx)

	return &types.QueryParamsResponse{Params: params}, nil
}
