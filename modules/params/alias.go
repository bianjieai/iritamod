package params

import (
	"github.com/cosmos/cosmos-sdk/x/params"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"

	"github.com/bianjieai/iritamod/modules/params/types"
)

const (
	ModuleName             = types.ModuleName
	StoreKey               = types.StoreKey
	RouterKey              = types.RouterKey
	QuerierRoute           = types.QuerierRoute
	EventTypeUpdateParams  = types.EventTypeUpdateParams
	AttributeKeyParamKey   = types.AttributeKeyParamKey
	AttributeValueCategory = types.AttributeValueCategory
)

var (
	ErrUnknownSubspace  = types.ErrUnknownSubspace
	ErrSettingParameter = types.ErrSettingParameter
	ErrUnknownKey       = types.ErrUnknownKey

	// NewQuerier alias cosmos-sdk/x/params/keeper.NewQuerier
	NewQuerier = paramskeeper.NewQuerier
)

type (
	// BaseAppModuleBasic alias cosmos-sdk/x/params.AppModuleBasic
	BaseAppModuleBasic = params.AppModuleBasic
	// Keeper alias cosmos-sdk/x/params/keeper.Keeper
	Keeper = paramskeeper.Keeper
)
