package node

import (
	"encoding/hex"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis - store genesis parameters
func InitGenesis(ctx sdk.Context, k Keeper, data GenesisState) {
	if err := ValidateGenesis(data); err != nil {
		panic(err.Error())
	}

	for _, node := range data.Nodes {
		id, _ := hex.DecodeString(node.Id)
		k.SetNode(ctx, id, node)
	}
}

// ExportGenesis - output genesis parameters
func ExportGenesis(ctx sdk.Context, k Keeper) *GenesisState {
	nodes := make([]Node, 0)

	k.IterateNodes(
		ctx,
		func(n Node) bool {
			nodes = append(nodes, n)
			return false
		},
	)

	return NewGenesisState(nodes)
}
