package slashing

// nolint

import (
	"gitlab.bianjie.ai/irita-pro/iritamod/modules/slashing/keeper"
	"gitlab.bianjie.ai/irita-pro/iritamod/modules/slashing/types"
)

var (
	// functions aliases
	NewKeeper             = keeper.NewKeeper
	NewMsgUnjailValidator = types.NewMsgUnjailValidator
)

type (
	MsgUnjailValidator = types.MsgUnjailValidator
	Keeper             = keeper.Keeper
)
