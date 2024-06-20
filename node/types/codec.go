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
	cdc.RegisterConcrete(&MsgCreateValidator{}, "iritamod/validator/MsgCreateValidator", nil)
	cdc.RegisterConcrete(&MsgUpdateValidator{}, "iritamod/validator/MsgUpdateValidator", nil)
	cdc.RegisterConcrete(&MsgRemoveValidator{}, "iritamod/validator/MsgRemoveValidator", nil)
	cdc.RegisterConcrete(&MsgGrantNode{}, "iritamod/node/MsgGrantNode", nil)
	cdc.RegisterConcrete(&MsgRevokeNode{}, "iritamod/node/MsgRevokeNode", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateValidator{},
		&MsgUpdateValidator{},
		&MsgRemoveValidator{},
		&MsgGrantNode{},
		&MsgRevokeNode{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
