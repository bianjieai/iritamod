package types

import (
	"fmt"

	"sigs.k8s.io/yaml"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// validator params default values
const (
	// DefaultHistorical entries is 100. Apps that don't use IBC can ignore this
	// value by not adding the staking module to the application module manager's
	// SetOrderBeginBlockers.
	DefaultHistoricalEntries uint32 = 100
)

var _ paramtypes.ParamSet = (*Params)(nil)

// NewParams creates a new Params instance
func NewParams(historicalEntries uint32) Params {
	return Params{HistoricalEntries: historicalEntries}
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return NewParams(DefaultHistoricalEntries)
}

func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func (p Params) Validate() error {
	return validateHistoricalEntries(p.HistoricalEntries)
}

func validateHistoricalEntries(i interface{}) error {
	if _, ok := i.(uint32); !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return nil
}
