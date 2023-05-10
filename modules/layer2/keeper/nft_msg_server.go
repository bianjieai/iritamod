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

func (k Keeper) DepositClassForNFT(ctx context.Context, nft *types.MsgDepositClassForNFT) (*types.MsgDepositClassForNFTResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k Keeper) WithdrawClassForNFT(ctx context.Context, nft *types.MsgWithdrawClassForNFT) (*types.MsgWithdrawClassForNFTResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k Keeper) DepositTokenForNFT(ctx context.Context, nft *types.MsgDepositTokenForNFT) (*types.MsgDepositTokenForNFTResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k Keeper) WithdrawTokenForNFT(ctx context.Context, nft *types.MsgWithdrawTokenForNFT) (*types.MsgWithdrawTokenForNFTResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k Keeper) CreateL2Space(ctx context.Context, space *types.MsgCreateL2Space) (*types.MsgCreateL2SpaceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k Keeper) TransferL2Space(ctx context.Context, space *types.MsgTransferL2Space) (*types.MsgTransferL2SpaceResponse, error) {
	//TODO implement me
	panic("implement me")
}


func (k Keeper) CreateL2Record(ctx context.Context, record *types.MsgCreateL2Record) (*types.MsgCreateL2RecordResponse, error) {
	//TODO implement me
	panic("implement me")
}
