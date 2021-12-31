package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}

// RegisterLegacyAminoCodec registers the necessary interfaces and concrete types
// on the provided Amino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgAssignRoles{}, "iritamod/perm/MsgAssignRoles", nil)
	cdc.RegisterConcrete(&MsgUnassignRoles{}, "iritamod/perm/MsgUnassignRoles", nil)
	cdc.RegisterConcrete(&MsgBlockAccount{}, "iritamod/perm/MsgBlockAccount", nil)
	cdc.RegisterConcrete(&MsgUnblockAccount{}, "iritamod/perm/MsgUnblockAccount", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAssignRoles{},
		&MsgUnassignRoles{},
		&MsgBlockAccount{},
		&MsgUnblockAccount{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}