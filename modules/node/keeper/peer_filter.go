package keeper

import (
	"encoding/hex"

	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/iritamod/modules/node/types"
)

// FilterNodeByID implements sdk.PeerFilter
func (k Keeper) FilterNodeByID(ctx sdk.Context, nodeID string) abci.ResponseQuery {
	id, err := hex.DecodeString(nodeID)
	if err != nil {
		return abci.ResponseQuery{
			Code: types.ErrInvalidNodeID.ABCICode(),
		}
	}

	if _, found := k.GetNode(ctx, id); !found {
		return abci.ResponseQuery{
			Code: types.ErrUnknownNode.ABCICode(),
		}
	}

	return abci.ResponseQuery{}
}
