package types

import (
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

const (
	// ModuleName is the name of the params module
	ModuleName = paramtypes.ModuleName

	// StoreKey is the string store representation
	StoreKey = ModuleName

	// QuerierRoute is the querier route for the params module
	QuerierRoute = ModuleName

	// RouterKey is the msg router key for the params module
	RouterKey = ModuleName
)
