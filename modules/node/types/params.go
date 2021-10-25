package types

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// validator params default values
const (
	// DefaultHistorical entries is 100. Apps that don't use IBC can ignore this
	// value by not adding the staking module to the application module manager's
	// SetOrderBeginBlockers.
	DefaultHistoricalEntries uint32 = 100
)

var (
	KeyHistoricalEntries = []byte("HistoricalEntries")
)

var _ paramtypes.ParamSet = (*Params)(nil)

// NewParams creates a new Params instance
func NewParams(historicalEntries uint32) Params {
	return Params{HistoricalEntries: historicalEntries}
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyHistoricalEntries, &p.HistoricalEntries, validateHistoricalEntries),
	}
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return NewParams(DefaultHistoricalEntries)
}

// unmarshal the current staking params value from store key or panic
func MustUnmarshalParams(cdc *codec.LegacyAmino, value []byte) Params {
	params, err := UnmarshalParams(cdc, value)
	if err != nil {
		panic(err)
	}
	return params
}

// unmarshal the current staking params value from store key
func UnmarshalParams(cdc *codec.LegacyAmino, value []byte) (params Params, err error) {
	err = cdc.Unmarshal(value, &params)
	return
}

func validateHistoricalEntries(i interface{}) error {
	if _, ok := i.(uint32); !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return nil
}
