package keeper

import (
	"github.com/bianjieai/iritamod/modules/wevm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
)

func (k Keeper) AddToContractDenyList(ctx sdk.Context, contractAddress string) error {
	store := k.GetStore(ctx)
	contractAddr := common.HexToAddress(contractAddress)
	if get := store.Get(types.ContractDenyListKey(contractAddr)); get != nil {
		return errors.Wrap(types.ErrContractAlreadyExist, "contract already in DenyList")
	}
	store.Set(types.ContractDenyListKey(contractAddr), []byte("true"))
	return nil
}

func (k Keeper) RemoveFromContractDenyList(ctx sdk.Context, contractAddress string) error {
	store := k.GetStore(ctx)
	contractAddr := common.HexToAddress(contractAddress)
	get := store.Get(types.ContractDenyListKey(contractAddr))
	if get != nil {
		store.Delete(types.ContractDenyListKey(contractAddr))
	} else {
		return errors.Wrapf(types.ErrNotFound, "the %s is not in contract dany list", contractAddr)
	}
	return nil
}

func (k Keeper) GetContractState(ctx sdk.Context, contractAddress string) (bool, error) {
	store := k.GetStore(ctx)
	contractAddr := common.HexToAddress(contractAddress)
	get := store.Get(types.ContractDenyListKey(contractAddr))
	if get != nil {
		return true, nil
	} else {
		return false, nil
	}
}
func (k Keeper) GetContractDenyList(ctx sdk.Context) ([]string, error) {
	list, err := k.IteratorContractDanyList(ctx)
	if err != nil {
		return nil, err
	}
	return list, nil
}
