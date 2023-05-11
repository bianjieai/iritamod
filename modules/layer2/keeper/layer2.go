package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) CreateSpace() (int64, error) {
	panic("implement me")
}

func (k Keeper) TransferSpace(spaceId int64, from, to sdk.AccAddress) error {
	panic("implement me")
}

func (k Keeper) OwnerOfSpace(spaceId int64) (sdk.AccAddress, error){
	panic("implement me")
}

func (k Keeper) HasSpace(spaceId int64) bool {
	panic("implement me")
}

func (k Keeper) CreateRecord() {
	panic("implement me")
}