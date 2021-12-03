package keeper_test

import (
	"github.com/stretchr/testify/require"
)

func (suite *KeeperTestSuite) Test_AddToContractDenyList() {
	tests := []struct {
		name    string
		address string
		success bool
	}{
		{"success add ", "0xa8982f13529550ebE01A5daa766804A0E2BfA95E", true},
		{"failed add", "0xa8982f13529550ebE01A5daa766804A0E2BfA95E", false},
	}
	for _, tt := range tests {
		err := suite.Keeper.AddToContractDenyList(suite.Ctx, tt.address)
		if tt.success {
			require.NoError(suite.T(), err)
		} else {
			require.Error(suite.T(), err)
			continue
		}
		state, err := suite.Keeper.GetContractState(suite.Ctx, tt.address)
		require.NoError(suite.T(), err)
		require.NotEmpty(suite.T(), state)
	}
}

func (suite *KeeperTestSuite) Test_RemoveFromContractDenyList() {
	tests := []struct {
		name    string
		address string
		success bool
	}{
		{"failed remove ", "0xa8982f13529550ebE01A5daa766804A0E2BfA95E", false},
		{"success remove", "0xa8982f13529550ebE01A5daa766804A0E2BfA95C", true},
	}
	err := suite.Keeper.AddToContractDenyList(suite.Ctx, tests[1].address)
	require.NoError(suite.T(), err)
	for _, tt := range tests {
		err := suite.Keeper.RemoveFromContractDenyList(suite.Ctx, tt.address)
		if tt.success {
			require.NoError(suite.T(), err)
		} else {
			require.Error(suite.T(), err)
			continue
		}
	}
}
