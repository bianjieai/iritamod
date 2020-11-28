package types

import (
	"strconv"

	tmbytes "github.com/tendermint/tendermint/libs/bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName is the name of the node module
	ModuleName = "node"

	// StoreKey is the string store representation
	StoreKey = ModuleName

	// QuerierRoute is the querier route for the node module
	QuerierRoute = ModuleName

	// RouterKey is the msg router key for the node module
	RouterKey = ModuleName
)

var (
	// Keys for store prefixes
	RootCertKey              = []byte{0x01} // prefix for root cert certificate
	ValidatorsKey            = []byte{0x02} // prefix for each key to a validator id
	ValidatorsNameKey        = []byte{0x03} // prefix for each key to a validator name
	ValidatorsByConsAddrKey  = []byte{0x04} // prefix for each key to a validator index, by consensus addr
	ValidatorsUpdateQueueKey = []byte{0x05} // prefix for each key of a validator to be updated
	HistoricalInfoKey        = []byte{0x06} // prefix for each key of a validator to be updated
	NodeKey                  = []byte{0x07} // prefix for node
)

// GetValidatorIDKey gets the key for the validator with id
func GetValidatorIDKey(id tmbytes.HexBytes) []byte {
	return append(ValidatorsKey, id.Bytes()...)
}

// GetValidatorNameKey gets the key for the validator with name
func GetValidatorNameKey(name string) []byte {
	return append(ValidatorsNameKey, []byte(name)...)
}

// GetValidatorConsAddrKey gets the key for the validator with cons address
func GetValidatorConsAddrKey(addr sdk.ConsAddress) []byte {
	return append(ValidatorsByConsAddrKey, addr...)
}

// GetValidatorUpdateQueueKey gets the key for the validator update queue
func GetValidatorUpdateQueueKey(pubkey string) []byte {
	return append(ValidatorsUpdateQueueKey, []byte(pubkey)...)
}

// GetHistoricalInfoKey gets the key for the historical info
func GetHistoricalInfoKey(height int64) []byte {
	return append(HistoricalInfoKey, []byte(strconv.FormatInt(height, 10))...)
}

// GetNodeKey gets the key for the node of the specified id
// VALUE: Node
func GetNodeKey(id tmbytes.HexBytes) []byte {
	return append(NodeKey, id...)
}
