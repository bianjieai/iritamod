package perm

import (
	"github.com/bianjieai/iritamod/modules/perm/keeper"
	"github.com/bianjieai/iritamod/modules/perm/types"
)

const (
	ModuleName              = types.ModuleName
	StoreKey                = types.StoreKey
	QuerierRoute            = types.QuerierRoute
	RouterKey               = types.RouterKey
	EventTypeAssignRoles    = types.EventTypeAssignRoles
	EventTypeUnassignRoles  = types.EventTypeUnassignRoles
	EventTypeBlockAccount   = types.EventTypeBlockAccount
	EventTypeUnblockAccount = types.EventTypeUnblockAccount
	AttributeKeyAccount     = types.AttributeKeyAccount
	AttributeValueCategory  = types.AttributeValueCategory
	RoleRootAdmin           = types.RoleRootAdmin
	RolePermAdmin           = types.RolePermAdmin
	RoleBlacklistAdmin      = types.RoleBlacklistAdmin
	RoleNodeAdmin           = types.RoleNodeAdmin
	RoleParamAdmin          = types.RoleParamAdmin
	RoleIDAdmin             = types.RoleIDAdmin
	RoleBaseM1Admin         = types.RoleBaseM1Admin
	RolePowerUser           = types.RolePowerUser
	RoleRelayerUser         = types.RoleRelayerUser
)

var (
	NewGenesisState             = types.NewGenesisState
	DefaultGenesisState         = types.DefaultGenesisState
	NewQuerier                  = keeper.NewQuerier
	NewKeeper                   = keeper.NewKeeper
	NewAuthDecorator            = keeper.NewAuthDecorator
	GetGenesisStateFromAppState = types.GetGenesisStateFromAppState
)

type (
	MsgAssignRoles    = types.MsgAssignRoles
	MsgUnassignRoles  = types.MsgUnassignRoles
	MsgBlockAccount   = types.MsgBlockAccount
	MsgUnblockAccount = types.MsgUnblockAccount
	Keeper            = keeper.Keeper
	GenesisState      = types.GenesisState
	RoleAccount       = types.RoleAccount
	Role              = types.Role
)
