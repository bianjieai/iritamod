package admin

import (
	"errors"
	"fmt"

	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"gitlab.bianjie.ai/irita-pro/iritamod/modules/admin/types"
)

// InitGenesis - store genesis account role set
func InitGenesis(ctx sdk.Context, k Keeper, data GenesisState) (res []abci.ValidatorUpdate) {
	if err := ValidateGenesis(data); err != nil {
		panic(err.Error())
	}

	for _, account := range data.RoleAccounts {
		var auth types.Auth
		for _, r := range account.Roles {
			auth = auth ^ r.Auth()
		}
		addr, err := sdk.AccAddressFromBech32(account.Address)
		if err != nil {
			panic(err)
		}
		k.SetAuth(ctx, addr, auth)
	}
	return
}

// ExportGenesis - output genesis account role set
func ExportGenesis(ctx sdk.Context, k Keeper) *GenesisState {
	return NewGenesisState(k.GetRoles(ctx), k.GetAllBlackAccounts(ctx))
}

// ValidateGenesis validates the provided admin genesis state
// must contains root role account
func ValidateGenesis(data GenesisState) error {
	// check root admin
	rootAdminFound := false
	for _, roleAccount := range data.RoleAccounts {
		for _, r := range roleAccount.Roles {
			if r == RoleRootAdmin {
				rootAdminFound = true
			}
		}
	}

	if !rootAdminFound {
		return errors.New("root account is not set in genesis state")
	}

	accountMap := make(map[string]bool, len(data.RoleAccounts))
	for _, roleAccount := range data.RoleAccounts {
		if _, ok := accountMap[roleAccount.Address]; ok {
			return fmt.Errorf("duplicate admin account in genesis state: address %s", roleAccount.Address)
		}
		accountMap[roleAccount.Address] = true
	}

	return nil
}
