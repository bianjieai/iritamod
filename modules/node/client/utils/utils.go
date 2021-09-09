package utils

import (
	"encoding/hex"
	"fmt"

	gogotypes "github.com/gogo/protobuf/types"

	tmbytes "github.com/tendermint/tendermint/libs/bytes"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/bianjieai/iritamod/modules/node/types"
)

func QueryValidator(clientCtx client.Context, param string) (validator types.Validator, height int64, err error) {
	queryValidator := func(id tmbytes.HexBytes) (validator types.Validator, height int64, err error) {
		res, height, err := clientCtx.QueryStore(types.GetValidatorIDKey(id), types.StoreKey)
		if err != nil {
			return
		}

		err = types.ModuleCdc.Unmarshal(res, &validator)
		return
	}

	if res, _, err := clientCtx.QueryStore(types.GetValidatorNameKey(param), types.StoreKey); err == nil && len(res) > 0 {
		var id gogotypes.BytesValue
		if err := types.ModuleCdc.Unmarshal(res, &id); err != nil {
			return validator, height, fmt.Errorf("no validator found %s", param)
		}
		return queryValidator(id.Value)
	}

	id, err := hex.DecodeString(param)
	if err != nil {
		return validator, height, fmt.Errorf("invalid validator id:%s", param)
	}

	return queryValidator(id)
}
