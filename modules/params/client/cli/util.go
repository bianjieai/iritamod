package cli

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type updateParams struct {
	Messages []json.RawMessage `json:"messages,omitempty"`
}

func parseUpdateParamsFile(cdc codec.Codec, path string) ([]sdk.Msg, error) {
	var updateParams updateParams

	contents, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, &updateParams)
	if err != nil {
		return nil, err
	}

	msgs := make([]sdk.Msg, len(updateParams.Messages))
	for i, anyJSON := range updateParams.Messages {
		var msg sdk.Msg
		err := cdc.UnmarshalInterfaceJSON(anyJSON, &msg)
		if err != nil {
			return nil, err
		}
		msgs[i] = msg
	}

	if len(msgs) == 0 {
		return nil, errors.New("no messages found in update params file")
	}

	return msgs, nil
}
