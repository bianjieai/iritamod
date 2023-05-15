package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

const (
	ModuleName = "layer2"

	StoreKey = ModuleName

	RouterKey = ModuleName
)

var ModuleAddress sdk.Address

func init() {
	ModuleAddress = authtypes.NewModuleAddress(ModuleName)
}