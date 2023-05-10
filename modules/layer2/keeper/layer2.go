package keeper

import (
	"github.com/bianjieai/iritamod/modules/layer2/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) CreateSpace() {

}

func (k Keeper) TransferSpace() {

}

func (k Keeper) OwnerOfSpace(spaceId int64) (sdk.AccAddress, error){
	if !k.HasSpace(spaceId) {
		return nil, types.ErrUnknownSpace
	}

	//FIXME: return owner
	panic("implement me")
}

func (k Keeper) HasSpace(spaceId int64) bool {
	panic("implement me")
}

func (k Keeper) CreateRecord() {

}