package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	amino = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateL2Space{},"iritamod/layer2/v1/MsgCreateL2Space",nil)
	cdc.RegisterConcrete(&MsgTransferL2Space{},"iritamod/layer2/v1/MsgTransferL2Space",nil)
	cdc.RegisterConcrete(&MsgCreateL2Record{},"iritamod/layer2/v1/MsgCreateL2Record",nil)
	cdc.RegisterConcrete(&MsgCreateNFTs{},"iritamod/layer2/v1/MsgCreateNFTs",nil)
	cdc.RegisterConcrete(&MsgUpdateNFTs{},"iritamod/layer2/v1/MsgUpdateNFTs",nil)
	cdc.RegisterConcrete(&MsgDeleteNFTs{},"iritamod/layer2/v1/MsgDeleteNFTs",nil)
	cdc.RegisterConcrete(&MsgDepositClassForNFT{},"iritamod/layer2/v1/MsgDepositClassForNFT",nil)
	cdc.RegisterConcrete(&MsgWithdrawClassForNFT{},"iritamod/layer2/v1/MsgWithdrawClassForNFT",nil)
	cdc.RegisterConcrete(&MsgDepositTokenForNFT{},"iritamod/layer2/v1/MsgDepositTokenForNFT",nil)
	cdc.RegisterConcrete(&MsgWithdrawTokenForNFT{},"iritamod/layer2/v1/MsgWithdrawTokenForNFT",nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		 &MsgCreateL2Space{},
		 &MsgTransferL2Space{},
		 &MsgCreateL2Record{},
		 &MsgCreateNFTs{},
		 &MsgUpdateNFTs{},
		 &MsgDeleteNFTs{},
		 &MsgDepositClassForNFT{},
		 &MsgWithdrawClassForNFT{},
		 &MsgDepositTokenForNFT{},
		 &MsgWithdrawTokenForNFT{},
		)
}