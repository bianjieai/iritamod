package simapp

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/codec"
)

// GenesisState defines a type alias for the irita genesis application state.
type GenesisState map[string]json.RawMessage

// NewDefaultGenesisState generates the default state for the application.
func NewDefaultGenesisState(cdc codec.JSONCodec) GenesisState {
	return ModuleBasics.DefaultGenesis(cdc)
}
