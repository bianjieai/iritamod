package params

import (
	"gitlab.bianjie.ai/irita-pro/iritamod/modules/params/types"
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
)
