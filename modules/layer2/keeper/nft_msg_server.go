package keeper

import (
	"context"

	"github.com/bianjieai/iritamod/modules/layer2/types"
)

var _ types.MsgServer = Keeper{}

func (k Keeper) CreateNFTs(ctx context.Context, msg *types.MsgCreateNFTs) (*types.MsgCreateNFTsResponse, error) {
	panic("implement me")
}

func (k Keeper) UpdateNFTs(ctx context.Context, msg *types.MsgUpdateNFTs) (*types.MsgUpdateNFTsResponse, error) {
	panic("implement me")
}

func (k Keeper) DeleteNFTs(ctx context.Context, msg *types.MsgDeleteNFTs) (*types.MsgDeleteNFTsResponse, error) {
	panic("implement me")
}

func (k Keeper) DepositClass(ctx context.Context, msg *types.MsgDepositClass) (*types.MsgDepositClassResponse, error) {
	panic("implement me")
}

func (k Keeper) WithdrawClass(ctx context.Context, msg *types.MsgWithdrawClass) (*types.MsgWithdrawClassResponse, error) {
	panic("implement me")
}

func (k Keeper) DepositNFT(ctx context.Context, msg *types.MsgDepositNFT) (*types.MsgDepositNFTResponse, error) {
	panic("implement me")
}

func (k Keeper) WithdrawNFT(ctx context.Context, msg *types.MsgWithdrawNFT) (*types.MsgWithdrawNFTResponse, error) {
	panic("implement me")
}

func (k Keeper) CreateL2Space(ctx context.Context, msg *types.MsgCreateL2Space) (*types.MsgCreateL2SpaceResponse, error) {
	panic("implement me")
}

func (k Keeper) TransferL2Space(ctx context.Context, msg *types.MsgTransferL2Space) (*types.MsgTransferL2SpaceResponse, error) {
	panic("implement me")
}
