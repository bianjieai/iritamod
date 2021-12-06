package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/iritamod/modules/wevm/types"
)

// ContractState Check if the contract is in the ContractDenyList
func (k Keeper) ContractState(goCtx context.Context, in *types.QueryContractStateRequest) (*types.QueryContractStateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	state, err := k.GetContractState(ctx, in.Address)
	if err != nil {
		return &types.QueryContractStateResponse{}, err
	}
	return &types.QueryContractStateResponse{Exist: state}, nil

}

// ContractDenyList get the ContractDenyList
func (k Keeper) ContractDenyList(goCtx context.Context, in *types.QueryContractDenyListRequest) (*types.QueryContractDenyListResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	list, err := k.GetContractDenyList(ctx)
	if err != nil {
		return nil, err
	}
	return &types.QueryContractDenyListResponse{ContractAddress: list}, nil
}
