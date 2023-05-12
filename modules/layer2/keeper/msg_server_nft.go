package keeper

import (
	"context"
	"github.com/bianjieai/iritamod/modules/layer2/types"
)

// CreateNFTs create native nft mappings on layer2 module.
func (k Keeper) CreateNFTs(goCtx context.Context, msg *types.MsgCreateNFTs) (*types.MsgCreateNFTsResponse, error) {
	//TODO implement me
	panic("implement me")
}

// UpdateNFTs update native nft mappings on layer2 module.
func (k Keeper) UpdateNFTs(goCtx context.Context, msg *types.MsgUpdateNFTs) (*types.MsgUpdateNFTsResponse, error) {
	//TODO implement me
	panic("implement me")
}

// DeleteNFTs delete native nft mappings on layer2 module.
func (k Keeper) DeleteNFTs(goCtx context.Context, msg *types.MsgDeleteNFTs) (*types.MsgDeleteNFTsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k Keeper) DepositClassForNFT(goCtx context.Context, msg *types.MsgDepositClassForNFT) (*types.MsgDepositClassForNFTResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k Keeper) WithdrawClassForNFT(goCtx context.Context, msg *types.MsgWithdrawClassForNFT) (*types.MsgWithdrawClassForNFTResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k Keeper) DepositTokenForNFT(goCtx context.Context, msg *types.MsgDepositTokenForNFT) (*types.MsgDepositTokenForNFTResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k Keeper) WithdrawTokenForNFT(goCtx context.Context, msg *types.MsgWithdrawTokenForNFT) (*types.MsgWithdrawTokenForNFTResponse, error) {
	//TODO implement me
	panic("implement me")
}