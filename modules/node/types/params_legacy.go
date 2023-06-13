package types

import paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

var (
	KeyHistoricalEntries = []byte("HistoricalEntries")
)

// Implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(
			KeyHistoricalEntries,
			&p.HistoricalEntries,
			validateHistoricalEntries,
		),
	}
}

func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}
