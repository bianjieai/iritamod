package types

import (
	"errors"
	"fmt"
	"strings"

	yaml "gopkg.in/yaml.v2"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

const (
	DefaultMintDenom = sdk.DefaultBondDenom
)

// Parameter store keys
var (
	KeyMintDenom = []byte("MintDenom")
)

// NewParams creates a new Params instance
func NewParams(
	mintDenom string,
) Params {
	return Params{
		MintDenom: mintDenom,
	}
}

// DefaultParams returns the module default parameters
func DefaultParams() Params {
	return NewParams(DefaultMintDenom)
}

// Validate validates the params
func (p Params) Validate() error {
	if err := validateMintDenom(p.MintDenom); err != nil {
		return err
	}

	return nil
}

// String implements Stringer.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// ParamSetPairs implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMintDenom, &p.MintDenom, validateMintDenom),
	}
}

func validateMintDenom(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if strings.TrimSpace(v) == "" {
		return errors.New("mint denom can not be empty")
	}

	if err := sdk.ValidateDenom(v); err != nil {
		return err
	}

	return nil
}
