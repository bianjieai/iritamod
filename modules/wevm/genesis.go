package wevm

import (
	"github.com/bianjieai/iritamod/modules/wevm/keeper"
	"github.com/bianjieai/iritamod/modules/wevm/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// InitGenesis initializes genesis state based on exported genesis
func InitGenesis(
	ctx sdk.Context,
	k keeper.Keeper,
	data types.GenesisState,
) []abci.ValidatorUpdate {

	for _, contractAddr := range data.ContractAddress {
		err := k.AddToContractDenyList(ctx, contractAddr)
		if err != nil {
			panic(err)
		}
		state, err := k.GetContractState(ctx, contractAddr)
		if err != nil {
			panic(err)
		}
		if !state {
			panic("add to contract deny list failed")
		}
	}
	return []abci.ValidatorUpdate{}
}

// ExportGenesis exports genesis state of the EVM module
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	list, err := k.IteratorContractDanyList(ctx)
	if err != nil {
		panic(err)
	}
	return &types.GenesisState{
		ContractAddress: list,
	}
}
