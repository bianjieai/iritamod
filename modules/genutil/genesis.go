package genutil

import (
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/iritamod/modules/genutil/types"
)

// InitGenesis - initialize accounts and deliver genesis transactions
func InitGenesis(
	ctx sdk.Context, nodeKeeper types.NodeKeeper,
	deliverTx deliverTxfn, genesisState types.GenesisState,
	txEncodingConfig client.TxEncodingConfig,
) []abci.ValidatorUpdate {

	var validators []abci.ValidatorUpdate
	if len(genesisState.GenTxs) > 0 {
		validators = DeliverGenTxs(ctx, genesisState.GenTxs, nodeKeeper, deliverTx, txEncodingConfig)
	}

	return validators
}
