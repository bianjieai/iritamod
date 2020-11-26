package types

import (
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
)

const (
	// ModuleName is the name of the node module
	ModuleName = "node"

	// StoreKey is the string store representation
	StoreKey string = ModuleName

	// QuerierRoute is the querier route for the node module
	QuerierRoute string = ModuleName

	// RouterKey is the msg router key for the node module
	RouterKey string = ModuleName
)

var (
	// Keys for store prefixes
	NodeKey = []byte{0x01} // prefix for node
)

// GetNodeKey gets the key for the node of the specified id
// VALUE: Node
func GetNodeKey(id tmbytes.HexBytes) []byte {
	return append(NodeKey, id...)
}
