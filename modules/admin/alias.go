package admin

import (
	"github.com/bianjieai/iritamod/modules/admin/keeper"
	"github.com/bianjieai/iritamod/modules/admin/types"
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
	RolePowerUser           = types.RolePowerUser
	RoleRelayerUser         = types.RoleRelayerUser
	RoleIDAdmin             = types.RoleIDAdmin
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
