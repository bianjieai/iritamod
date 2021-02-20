package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/tendermint/tendermint/crypto/tmhash"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/iritamod/modules/perm/keeper"
	"github.com/bianjieai/iritamod/modules/perm/types"
	"github.com/bianjieai/iritamod/simapp"
)

type KeeperTestSuite struct {
	suite.Suite

	cdc    *codec.LegacyAmino
	ctx    sdk.Context
	keeper *keeper.Keeper
	app    *simapp.SimApp
}

var (
	rootAdmin   sdk.AccAddress
	account     = sdk.AccAddress("test_account")
	account1    = sdk.AccAddress("test_account1")
	addRoles    = []types.Role{types.RolePermAdmin, types.RoleBlacklistAdmin, types.RoleNodeAdmin}
	removeRoles = []types.Role{types.RoleParamAdmin, types.RolePowerUser}
)

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) SetupTest() {
	app := simapp.Setup(false)

	suite.cdc = app.LegacyAmino()
	suite.ctx = app.BaseApp.NewContext(false, tmproto.Header{})
	suite.app = app

	rootAdmin = sdk.AccAddress(tmhash.SumTruncated([]byte("rootAdmin")))

	suite.keeper = &app.PermKeeper
}

func (suite *KeeperTestSuite) TestRootAdminCheck() {
	auth := suite.keeper.GetAuth(suite.ctx, rootAdmin)
	suite.Equal(types.RoleRootAdmin.Auth(), auth)

	roleAccounts := suite.keeper.GetRoles(suite.ctx)
	suite.Equal(1, len(roleAccounts))
}

func (suite *KeeperTestSuite) TestRoleAuth() {
	// can not operate root admin
	err := suite.keeper.Authorize(suite.ctx, rootAdmin, rootAdmin, addRoles...)
	suite.Error(err)

	// can not add root admin
	err = suite.keeper.Authorize(suite.ctx, account, rootAdmin, types.RoleRootAdmin)
	suite.Error(err)

	err = suite.keeper.Authorize(suite.ctx, account, rootAdmin, addRoles...)
	suite.NoError(err)

	err = suite.keeper.Authorize(suite.ctx, account1, rootAdmin, addRoles...)
	suite.NoError(err)

	// perm admin can not operate another perm admin
	err = suite.keeper.Authorize(suite.ctx, account1, account, addRoles...)
	suite.Error(err)

	err = suite.keeper.Authorize(suite.ctx, account1, rootAdmin, addRoles...)
	suite.NoError(err)

	// perm admin can not operate another perm admin
	err = suite.keeper.Unauthorize(suite.ctx, account1, account, removeRoles...)
	suite.Error(err)

	// can not remove root admin
	err = suite.keeper.Unauthorize(suite.ctx, account, rootAdmin, types.RoleRootAdmin)
	suite.Error(err)

	// can not operate root admin
	err = suite.keeper.Unauthorize(suite.ctx, rootAdmin, rootAdmin, addRoles...)
	suite.Error(err)

	err = suite.keeper.Unauthorize(suite.ctx, account, rootAdmin, addRoles...)
	suite.NoError(err)

	err = suite.keeper.Unauthorize(suite.ctx, account1, rootAdmin, addRoles...)
	suite.NoError(err)

	suite.TestRootAdminCheck()
}

func (suite *KeeperTestSuite) TestAddRoles() {
	err := suite.keeper.Authorize(suite.ctx, account, rootAdmin, addRoles...)
	suite.NoError(err)

	roles := suite.keeper.GetAuth(suite.ctx, account).Roles()
	suite.Equal(addRoles, roles)

	roleAccounts := suite.keeper.GetRoles(suite.ctx)
	suite.Equal(2, len(roleAccounts))
	for _, ra := range roleAccounts {
		if ra.Address == rootAdmin.String() {
			continue
		}
		suite.Equal(account.String(), ra.Address)
		suite.Equal(addRoles, ra.Roles)
	}

	// permission admin can not add other permission admin
	err = suite.keeper.Authorize(suite.ctx, account1, account, types.RolePermAdmin)
	suite.Error(err)
}

func (suite *KeeperTestSuite) TestRemoveRoles() {
	err := suite.keeper.Authorize(suite.ctx, account, rootAdmin, types.RolePermAdmin,
		types.RoleBlacklistAdmin, types.RoleNodeAdmin, types.RoleParamAdmin, types.RolePowerUser)
	suite.NoError(err)

	err = suite.keeper.Unauthorize(suite.ctx, account, rootAdmin, removeRoles...)
	suite.NoError(err)

	existingRoles := addRoles
	roles := suite.keeper.GetAuth(suite.ctx, account).Roles()
	suite.Equal(existingRoles, roles)

	roleAccounts := suite.keeper.GetRoles(suite.ctx)
	suite.Equal(2, len(roleAccounts))
	for _, ra := range roleAccounts {
		if ra.Address == rootAdmin.String() {
			continue
		}
		suite.Equal(account.String(), ra.Address)
		suite.Equal(existingRoles, ra.Roles)
	}

	// remove all roles from this account
	err = suite.keeper.Unauthorize(suite.ctx, account, rootAdmin, existingRoles...)
	suite.NoError(err)

	roles = suite.keeper.GetAuth(suite.ctx, account).Roles()
	suite.Empty(roles)
	roleAccounts = suite.keeper.GetRoles(suite.ctx)
	suite.Equal(1, len(roleAccounts))
	suite.Equal(roleAccounts[0].Address, rootAdmin.String())
	suite.Equal(roleAccounts[0].Roles, []types.Role{types.RoleRootAdmin})
}

func (suite *KeeperTestSuite) TestBlockAccount() {
	err := suite.keeper.Authorize(suite.ctx, account, rootAdmin, addRoles...)
	suite.NoError(err)

	// can not block admin account
	err = suite.keeper.Block(suite.ctx, account)
	suite.Error(err)

	err = suite.keeper.Unauthorize(suite.ctx, account, rootAdmin, addRoles...)
	suite.NoError(err)

	err = suite.keeper.Block(suite.ctx, account)
	suite.NoError(err)

	// already blocked
	err = suite.keeper.Block(suite.ctx, account)
	suite.Error(err)

	blackList := suite.keeper.GetAllBlackAccounts(suite.ctx)
	suite.Equal(1, len(blackList))
	suite.Equal(account.String(), blackList[0])

	err = suite.keeper.Unblock(suite.ctx, account)
	suite.NoError(err)

	err = suite.keeper.Unblock(suite.ctx, account)
	suite.Error(err)

	blackList = suite.keeper.GetAllBlackAccounts(suite.ctx)
	suite.Empty(blackList)
}
