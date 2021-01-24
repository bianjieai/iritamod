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
	EventTypeAddRoles       = types.EventTypeAddRoles
	EventTypeRemoveRoles    = types.EventTypeRemoveRoles
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
	RoleMintAdmin           = types.RoleMintAdmin
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
	MsgAddRoles       = types.MsgAddRoles
	MsgRemoveRoles    = types.MsgRemoveRoles
	MsgBlockAccount   = types.MsgBlockAccount
	MsgUnblockAccount = types.MsgUnblockAccount
	Keeper            = keeper.Keeper
	GenesisState      = types.GenesisState
	RoleAccount       = types.RoleAccount
	Role              = types.Role
)
