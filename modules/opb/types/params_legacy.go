package types

import paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

// Parameter store keys
var (
	KeyBaseTokenDenom            = []byte("BaseTokenDenom")
	KeyPointTokenDenom           = []byte("PointTokenDenom")
	KeyBaseTokenManager          = []byte("BaseTokenManager")
	KeyUnrestrictedTokenTransfer = []byte("UnrestrictedTokenTransfer")
)

// Deprecated: ParamTable for opb module.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// ParamSetPairs implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyBaseTokenDenom, &p.BaseTokenDenom, validateBaseTokenDenom),
		paramtypes.NewParamSetPair(KeyPointTokenDenom, &p.PointTokenDenom, validatePointTokenDenom),
		paramtypes.NewParamSetPair(KeyBaseTokenManager, &p.BaseTokenManager, validateBaseTokenManager),
		paramtypes.NewParamSetPair(KeyUnrestrictedTokenTransfer, &p.UnrestrictedTokenTransfer, validateUnrestrictedTokenTransfer),
	}
}
