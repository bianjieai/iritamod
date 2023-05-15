package expected_keeper

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
)

type NFT interface {
	GetID() string
	GetName() string
	GetOwner() sdk.AccAddress
	GetURI() string
	GetURIHash() string
	GetData() string
}

type NFTKeeper interface {
	SaveNFT(ctx sdk.Context, denomID, tokenID, tokenNm, tokenURI, tokenUriHash, tokenData string, receiver sdk.AccAddress) error
	Transfer(ctx sdk.Context, classID string, tokenID string, receiver sdk.AccAddress) error
	TransferOwnership(ctx sdk.Context, denomID, tokenID, tokenNm, tokenURI, tokenURIHash, tokenData string, srcOwner, dstOwner sdk.AccAddress) error
	RemoveNFT(ctx sdk.Context, denomID, tokenID string, owner sdk.AccAddress) error
	GetNFT(ctx sdk.Context, classID, tokenID string) (NFT, error)

	UpdateClass(ctx sdk.Context, class Class) error
	GetDenomInfo(ctx sdk.Context, denomID string) (*Denom, error)
	TransferDenomOwner(ctx sdk.Context, denomID string, srcOwner, dstOwner sdk.AccAddress) error
}

type Denom struct {
	Id               string
	Name             string
	Schema           string
	Creator          string
	Symbol           string
	MintRestricted   bool
	UpdateRestricted bool
	Description      string
	Uri              string
	UriHash          string
	Data             string
}

type Class struct {
	Id          string
	Name        string
	Symbol      string
	Description string
	Uri         string
	UriHash     string
	Data        *types.Any
}

type DenomMetadata struct {
	Creator          string
	Schema           string
	MintRestricted   bool
	UpdateRestricted bool
	Data             string
}

func (m *DenomMetadata) Reset()         { *m = DenomMetadata{} }
func (m *DenomMetadata) String() string { return proto.CompactTextString(m) }
func (*DenomMetadata) ProtoMessage()    {}
