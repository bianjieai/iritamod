package mock

import (
	_ "embed"
	"encoding/json"
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/iritamod/modules/layer2/types"
)

//go:embed mock_data.json
var badKidsRawData []byte //nolint: golint

type NFTKeeper struct {
	store map[string]*Class
}

func NewNFTKeeper() *NFTKeeper {
	var badKids Class
	json.Unmarshal(badKidsRawData, &badKids)

	store := make(map[string]*Class)
	store[badKids.ClassId] = &badKids

	return &NFTKeeper{
		store: store,
	}
}

type NFT struct {
	TokenId      string `json:"token_id,omitempty"`
	TokenName    string `json:"token_name,omitempty"`
	TokenUri     string `json:"token_uri,omitempty"`
	TokenUriHash string `json:"token_uri_hash,omitempty"`
	TokenData    string `json:"token_data,omitempty"`
	Owner        string `json:"owner,omitempty"`
}

type Class struct {
	ClassId        string          `json:"class_id,omitempty"`
	TokenIds       map[string]*NFT `json:"token_ids,omitempty"`
	Owner          string          `json:"owner,omitempty"`
	MintRestricted bool            `json:"mint_restricted,omitempty"`
}

func (mk *NFTKeeper) SaveNFT(_ sdk.Context, classID, tokenID, tokenNm, tokenURI, tokenUriHash, tokenData string, receiver sdk.AccAddress) error {
	class, ok := mk.store[classID]
	if !ok {
		return errors.New("class not found")
	}

	if _, ok := class.TokenIds[tokenID]; ok {
		return errors.New("token already exists")
	}

	nft := NFT{
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

func (mk *NFTKeeper) UpdateNFT(_ sdk.Context, classID, tokenID, tokenNm, tokenURI, tokenUriHash, tokenData string, owner sdk.AccAddress) error {
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

func (mk *NFTKeeper) RemoveNFT(_ sdk.Context, classID, tokenID string, owner sdk.AccAddress) error {
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

func (mk *NFTKeeper) TransferNFT(_ sdk.Context, classID, tokenID string, srcOwner, dstOwner sdk.AccAddress) error {
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

func (mk *NFTKeeper) TransferClass(_ sdk.Context, classID string, srcOwner, dstOwner sdk.AccAddress) error {
	class, ok := mk.store[classID]
	if !ok {
		return errors.New("class not found")
	}

	if class.Owner != srcOwner.String() {
		return errors.New("class not owned by this owner")
	}

	class.Owner = dstOwner.String()
	return nil
}

func (mk *NFTKeeper) UpdateClassMintRestricted(_ sdk.Context, classID string, mintRestricted bool, owner sdk.AccAddress) error {
	class, ok := mk.store[classID]
	if !ok {
		return errors.New("class not found")
	}

	if class.Owner != owner.String() {
		return errors.New("class not owned by this owner")
	}

	class.MintRestricted = mintRestricted
	return nil
}

func (mk *NFTKeeper) GetClass(_ sdk.Context, classID string) (types.Class, error) {
	class, ok := mk.store[classID]
	if !ok {
		return nil, errors.New("class not found")
	}

	return class, nil
}

func (mk *NFTKeeper) GetNFT(_ sdk.Context, classID, tokenID string) (types.NFT, error) {
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

func (mc *Class) GetID() string {
	return mc.ClassId
}

func (mc *Class) GetOwner() string {
	return mc.Owner
}

func (mc *Class) GetMintRestricted() bool {
	return mc.MintRestricted
}

func (mn *NFT) GetOwner() sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(mn.Owner)
	return addr
}
