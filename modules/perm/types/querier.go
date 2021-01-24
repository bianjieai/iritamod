package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// query endpoints supported by the validator Querier
const (
	QueryRoles     = "roles"
	QueryBlacklist = "blacklist"
)

type QueryRolesParams struct {
	Address sdk.AccAddress
}

func NewQueryRolesParams(address sdk.AccAddress) QueryRolesParams {
	return QueryRolesParams{
		Address: address,
	}
}

type QueryBlacklistParams struct {
	Page  int
	Limit int
}

func NewQueryBlacklistParams(page, limit int) QueryBlacklistParams {
	return QueryBlacklistParams{page, limit}
}
