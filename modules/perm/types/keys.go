package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName is the name of the perm module
	ModuleName = "perm"

	// StoreKey is the string store representation
	StoreKey = ModuleName

	// QuerierRoute is the querier route for the perm module
	QuerierRoute = ModuleName

	// RouterKey is the msg router key for the perm module
	RouterKey = ModuleName
)

var (
	// Keys for store prefixes
	AuthKey  = []byte{0x01} // prefix for each key to a account auth
	BlackKey = []byte{0x02} // prefix for each key to a black account
)

// GetAuthKey gets the key for the role with address
func GetAuthKey(addr sdk.AccAddress) []byte {
	return append(AuthKey, addr...)
}

// GetBlackKey gets the key for the black with address
func GetBlackKey(addr sdk.AccAddress) []byte {
	return append(BlackKey, addr...)
}
