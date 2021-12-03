package types

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

const (
	// ModuleName string name of module
	ModuleName = "wevm"

	StoreKey = ModuleName

	// RouterKey uses module name for routing
	RouterKey = ModuleName

	ContractDenyListName = "contract-deny-list"
)

const (
	KeyPrefixContractDenyList = "ContractDenyList"
)

func ContractDenyListPath(contractAddress common.Address) string {
	return fmt.Sprintf("%s/%s", KeyPrefixContractDenyList, contractAddress)
}

// ContractDenyListKey defines the full key under which a contract deny list is stored.
func ContractDenyListKey(contractAddress common.Address) []byte {
	return []byte(ContractDenyListPath(contractAddress))
}
