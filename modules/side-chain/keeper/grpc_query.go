package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/bianjieai/iritamod/modules/side-chain/types"
)

func (k Keeper) Space(goCtx context.Context, request *types.QuerySpaceRequest) (*types.QuerySpaceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	space, err := k.GetSpace(ctx, request.SpaceId)
	if err != nil {
		return nil, err
	}

	return &types.QuerySpaceResponse{
		Space: &space,
	}, nil
}

func (k Keeper) SpaceOfOwner(goCtx context.Context, req *types.QuerySpaceOfOwnerRequest) (*types.QuerySpaceOfOwnerResponse, error) {

	var spaces []types.Space
	var pageResp *query.PageResponse
	ctx := sdk.UnwrapSDKContext(goCtx)

	owner, err := sdk.AccAddressFromBech32(req.Owner)
	if err != nil {
		return nil, err
	}

	if pageResp, err = query.Paginate(k.getSpaceOfOwnerStore(ctx, owner), req.Pagination, func(key []byte, _ []byte) error {
		// key: 0x02/<owner>/<delimiter>/
		spaceId, err := strconv.ParseUint(string(key), 10, 64)
		if err != nil {
			return err
		}
		space, err := k.GetSpace(ctx, spaceId)
		if err != nil {
			return err
		}
		spaces = append(spaces, space)
		return nil
	}); err != nil {
		return nil, err
	}
	return &types.QuerySpaceOfOwnerResponse{
		Spaces:     spaces,
		Pagination: pageResp,
	}, nil
}

func (k Keeper) BlockHeader(goCtx context.Context, request *types.QueryBlockHeaderRequest) (*types.QueryBlockHeaderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	header, err := k.GetBlockHeader(ctx, request.SpaceId, request.Height)
	if err != nil {
		return nil, err
	}
	return &types.QueryBlockHeaderResponse{Header: header}, nil
}
