package utils

import (
	"encoding/json"
	"io/ioutil"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/bianjieai/iritamod/modules/params/types"
)

type (
	// ParamChangesJSON defines a slice of ParamChangeJSON objects which can be
	// converted to a slice of ParamChange objects.
	ParamChangesJSON []ParamChangeJSON

	// ParamChangeJSON defines a parameter change used in JSON input. This
	// allows values to be specified in raw JSON instead of being string encoded.
	ParamChangeJSON struct {
		Subspace string          `json:"subspace" yaml:"subspace"`
		Key      string          `json:"key" yaml:"key"`
		Value    json.RawMessage `json:"value" yaml:"value"`
	}
)

// ToParamChange converts a ParamChangeJSON object to ParamChange.
func (pcj ParamChangeJSON) ToParamChange() types.ParamChange {
	return types.ParamChange{Subspace: pcj.Subspace, Key: pcj.Key, Value: string(pcj.Value)}
}

// ToParamChanges converts a slice of ParamChangeJSON objects to a slice of
// ParamChange.
func (pcj ParamChangesJSON) ToParamChanges() []types.ParamChange {
	res := make([]types.ParamChange, len(pcj))
	for i, pc := range pcj {
		res[i] = pc.ToParamChange()
	}
	return res
}

// ParseParamChange reads and parses a ParamChangesJSON from file.
func ParseParamChange(cdc *codec.LegacyAmino, proposalFile string) (ParamChangesJSON, error) {
	paramChanges := ParamChangesJSON{}

	contents, err := ioutil.ReadFile(proposalFile)
	if err != nil {
		return paramChanges, err
	}

	if err := cdc.UnmarshalJSON(contents, &paramChanges); err != nil {
		return paramChanges, err
	}

	return paramChanges, nil
}
