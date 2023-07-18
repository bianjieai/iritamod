package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateSpace{}, "iritamod/side-chain/v1/MsgCreateSpace", nil)
	cdc.RegisterConcrete(&MsgTransferSpace{}, "iritamod/side-chain/v1/MsgTransferSpace", nil)
	cdc.RegisterConcrete(&MsgCreateBlockHeader{}, "iritamod/side-chain/v1/MsgCreateRecord", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateSpace{},
		&MsgTransferSpace{},
		&MsgCreateBlockHeader{},
	)
}
