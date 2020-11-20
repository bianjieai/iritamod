package identity

// nolint

import (
	"github.com/bianjieai/iritamod/modules/identity/keeper"
	"github.com/bianjieai/iritamod/modules/identity/types"
)

const (
	ModuleName              = types.ModuleName
	StoreKey                = types.StoreKey
	QuerierRoute            = types.QuerierRoute
	RouterKey               = types.RouterKey
	QueryIdentity           = types.QueryIdentity
	EventTypeCreateIdentity = types.EventTypeCreateIdentity
	EventTypeUpdateIdentity = types.EventTypeUpdateIdentity
	AttributeValueCategory  = types.AttributeValueCategory
	AttributeKeyID          = types.AttributeKeyID
	AttributeKeyOwner       = types.AttributeKeyOwner
	DoNotModifyDesc         = types.DoNotModifyDesc
)

var (
	NewKeeper           = keeper.NewKeeper
	NewQuerier          = keeper.NewQuerier
	ModuleCdc           = types.ModuleCdc
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis
	NewGenesisState     = types.NewGenesisState
)

type (
	Keeper              = keeper.Keeper
	Identity            = types.Identity
	GenesisState        = types.GenesisState
	MsgCreateIdentity   = types.MsgCreateIdentity
	MsgUpdateIdentity   = types.MsgUpdateIdentity
	QueryIdentityParams = types.QueryIdentityParams
)
