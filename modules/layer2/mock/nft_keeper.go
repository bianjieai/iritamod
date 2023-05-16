package mock

import (
	_ "embed"
	"encoding/json"
	"errors"
	"github.com/bianjieai/iritamod/modules/layer2/expected_keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

//go:embed nfts.json
var rawNFTsData []byte

type MockNFTKeeper struct {
	store map[string]*MockClass
}

func NewMockNFTKeeper() *MockNFTKeeper {
	var nfts map[string]*MockClass
	json.Unmarshal(rawNFTsData, &nfts)

	return &MockNFTKeeper{
		store: nfts,
	}
}

type MockNFT struct {
	TokenId      string `json:"token_id,omitempty"`
	TokenName    string `json:"token_name,omitempty"`
	TokenUri     string `json:"token_uri,omitempty"`
	TokenUriHash string `json:"token_uri_hash,omitempty"`
	TokenData    string `json:"token_data,omitempty"`
	Owner        string `json:"owner,omitempty"`
}

type MockClass struct {
	ClassId        string              `json:"class_id,omitempty"`
	TokenIds       map[string]*MockNFT `json:"token_ids,omitempty"`
	Creator        string              `json:"creator,omitempty"`
	MintRestricted bool                `json:"mint_restricted,omitempty"`
}

func (mk *MockNFTKeeper) SaveNFT(_ sdk.Context, classID, tokenID, tokenNm, tokenURI, tokenUriHash, tokenData string, receiver sdk.AccAddress) error {
	class, ok := mk.store[classID]
	if !ok {
		return errors.New("class not found")
	}

	if _, ok := class.TokenIds[tokenID]; ok {
		return errors.New("token already exists")
	}

	nft := MockNFT{
		TokenId:      tokenID,
		TokenName:    tokenNm,
		TokenUri:     tokenURI,
		TokenUriHash: tokenUriHash,
		TokenData:    tokenData,
		Owner:        receiver.String(),
	}

	class.TokenIds[tokenID] = &nft
	return nil
}

func (mk *MockNFTKeeper) UpdateNFT(_ sdk.Context, classID, tokenID, tokenNm, tokenURI, tokenUriHash, tokenData string, owner sdk.AccAddress) error {
	class, ok := mk.store[classID]
	if !ok {
		return errors.New("class not found")
	}

	nft, ok := class.TokenIds[tokenID]
	if !ok {
		return errors.New("token not exists")
	}

	if nft.Owner != owner.String() {
		return errors.New("token not belongs to owner")
	}

	nft.TokenName = tokenNm
	nft.TokenUri = tokenURI
	nft.TokenUriHash = tokenUriHash
	nft.TokenData = tokenData

	return nil
}

func (mk *MockNFTKeeper) RemoveNFT(_ sdk.Context, classID, tokenID string, owner sdk.AccAddress) error {
	class, ok := mk.store[classID]
	if !ok {
		return errors.New("class not found")
	}

	nft, ok := class.TokenIds[tokenID]
	if !ok {
		return errors.New("token not exists")
	}

	if nft.Owner != owner.String() {
		return errors.New("token not belongs to owner")
	}

	delete(class.TokenIds, tokenID)
	return nil
}

func (mk *MockNFTKeeper) TransferNFT(_ sdk.Context, classID, tokenID string, srcOwner, dstOwner sdk.AccAddress) error {
	class, ok := mk.store[classID]
	if !ok {
		return errors.New("class not found")
	}

	nft, ok := class.TokenIds[tokenID]
	if !ok {
		return errors.New("token not exists")
	}

	if nft.Owner != srcOwner.String() {
		return errors.New("token not owned by this owner")
	}

	nft.Owner = dstOwner.String()
	return nil
}

func (mk *MockNFTKeeper) TransferClass(_ sdk.Context, classID string, srcOwner, dstOwner sdk.AccAddress) error {
	class, ok := mk.store[classID]
	if !ok {
		return errors.New("class not found")
	}

	if class.Creator != srcOwner.String() {
		return errors.New("class not owned by this owner")
	}

	class.Creator = dstOwner.String()
	return nil
}

func (mk *MockNFTKeeper) UpdateClassMintRestricted(_ sdk.Context, classID string, mintRestricted bool, owner sdk.AccAddress) error {
	class, ok := mk.store[classID]
	if !ok {
		return errors.New("class not found")
	}

	if class.Creator != owner.String() {
		return errors.New("class not owned by this owner")
	}

	class.MintRestricted = mintRestricted
	return nil
}

func (mk *MockNFTKeeper) GetClass(_ sdk.Context, classID string) (expected_keeper.Class, error) {
	class, ok := mk.store[classID]
	if !ok {
		return nil, errors.New("class not found")
	}

	return class, nil
}

func (mk *MockNFTKeeper) GetNFT(_ sdk.Context, classID, tokenID string) (expected_keeper.NFT, error) {
	class, ok := mk.store[classID]
	if !ok {
		return nil, errors.New("class not found")
	}

	nft, ok := class.TokenIds[tokenID]
	if !ok {
		return nil, errors.New("token not found")
	}

	return nft, nil
}

func (mc *MockClass) GetID() string {
	return mc.ClassId
}

func (mc *MockClass) GetCreator() string {
	return mc.Creator
}

func (mc *MockClass) GetMintRestricted() bool {
	return mc.MintRestricted
}

func (mn *MockNFT) GetOwner() sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(mn.Owner)
	return addr
}
