package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"

	"strconv"
)

const (
	ModuleName = "layer2"

	StoreKey = ModuleName

	RouterKey = ModuleName
)

var (
	SpaceMainKey = []byte{0x01}
	RecordMainKey = []byte{0x02}
	NFTMainKey = []byte{0x03}

	// Space Subkey
	Spacekey = []byte{0x01}
	SpaceOfUserKey = []byte{0x02}

	// NFT Subkey
	ClassKeyForNft = []byte{0x01}
	TokenKeyForNft = []byte{0x02}
	NFTsOfOwnerKey = []byte{0x03}

	Delimiter = []byte{0x00}
	Placeholder = []byte{0x01}
)

// space store key

// spaceStoreKey returns the byte representation of the space key
// Items are stored with the following key: values
// <0x01><0x01><space_id>
func spaceStoreKey(spaceId uint64) []byte {
	spaceIdStr := strconv.FormatUint(spaceId, 10)
	key := make([]byte, len(SpaceMainKey)+len(Spacekey)+len(spaceIdStr))
	copy(key, SpaceMainKey)
	copy(key[len(SpaceMainKey):], Spacekey)
	copy(key[len(SpaceMainKey)+len(Spacekey):], spaceIdStr)
	return key
}

// spaceOfL2UserStoreKey returns the byte representation of the space of l2 user key
// Items are stored with the following key: values
// <0x01><0x02><owner><delimiter><space_id>
func spaceOfL2UserStoreKey(owner sdk.AccAddress, spaceId uint64) []byte {
	owner = address.MustLengthPrefix(owner)
	spaceIdStr := strconv.FormatUint(spaceId, 10)
	key := make([]byte, len(SpaceMainKey)+len(SpaceOfUserKey)+len(owner)+len(Delimiter)+len(spaceIdStr))
	copy(key, SpaceMainKey)
	copy(key[len(SpaceMainKey):], SpaceOfUserKey)
	copy(key[len(SpaceMainKey)+len(SpaceOfUserKey):], owner)
	copy(key[len(SpaceMainKey)+len(SpaceOfUserKey)+len(owner):], Delimiter)
	copy(key[len(SpaceMainKey)+len(SpaceOfUserKey)+len(owner)+len(Delimiter):], spaceIdStr)
	return key
}

// record mappings store key

// recordStoreKey returns the byte representation of the record key
// Items are stored with the following key: values
// <0x02><space_id><delimiter><block_height>
func recordStoreKey(spaceId, blockHeight uint64) []byte {
	spaceIdStr := strconv.FormatUint(spaceId, 10)
	blockHeightStr := strconv.FormatUint(blockHeight, 10)
	key := make([]byte, len(RecordMainKey)+len(spaceIdStr)+len(Delimiter)+len(blockHeightStr))
	copy(key, RecordMainKey)
	copy(key[len(RecordMainKey):], spaceIdStr)
	copy(key[len(RecordMainKey)+len(spaceIdStr):], Delimiter)
	copy(key[len(RecordMainKey)+len(spaceIdStr)+len(Delimiter):], blockHeightStr)
	return key
}

// nft mappings store key

// classStoreKeyForNft returns the byte representation of the class key of nft mappings
// Items are stored with the following key: values
// <0x03><0x01><space_id>
func classStoreKeyForNft(classId string) []byte {
	key := make([]byte, len(NFTMainKey)+len(ClassKeyForNft)+len(classId))
	copy(key, NFTMainKey)
	copy(key[len(NFTMainKey):], ClassKeyForNft)
	copy(key[len(NFTMainKey)+len(ClassKeyForNft):], classId)
	return key
}

// nftStoreKeyForNft returns the byte representation of the nft key of nft mappings
// Items are stored with the following key: values
// <0x03><0x02><space_id><delimiter><class_id><delimiter><token_id>
func nftStoreKeyForNft(spaceId uint64, classId, tokenId string) []byte {
	spaceIdStr := strconv.FormatUint(spaceId, 10)

	key := make([]byte,
		len(NFTMainKey)+len(TokenKeyForNft)+
		len(spaceIdStr)+len(Delimiter)+
		len(classId)+len(Delimiter)+
		len(tokenId))

	copy(key, NFTMainKey)
	copy(key[len(NFTMainKey):], TokenKeyForNft)
	copy(key[len(NFTMainKey)+len(TokenKeyForNft):], spaceIdStr)
	copy(key[len(NFTMainKey)+len(TokenKeyForNft)+len(spaceIdStr):], Delimiter)
	copy(key[len(NFTMainKey)+len(TokenKeyForNft)+len(spaceIdStr)+len(Delimiter):], classId)
	copy(key[len(NFTMainKey)+len(TokenKeyForNft)+len(spaceIdStr)+len(Delimiter)+len(classId):], Delimiter)
	copy(key[len(NFTMainKey)+len(TokenKeyForNft)+len(spaceIdStr)+len(Delimiter)+len(classId)+len(Delimiter):], tokenId)
	return key
}

// nftsOfOwnerStoreKeyForNft returns the byte representation of the owner key of nft mappings
// Items are stored with the following key: values
// <0x03><0x03><owner><delimiter><space_id><delimiter><class_id><delimiter><nft_id>
func nftsOfOwnerStoreKeyForNft(owner sdk.AccAddress, spaceId uint64, classId, tokenId string) []byte{
	owner = address.MustLengthPrefix(owner)
	spaceIdStr := strconv.FormatUint(spaceId, 10)

	key := make([]byte,
		len(NFTMainKey)+len(NFTsOfOwnerKey)+
		len(owner)+len(Delimiter)+
		len(spaceIdStr)+len(Delimiter)+
		len(classId)+len(Delimiter)+
		len(tokenId))

	copy(key, NFTMainKey)
	copy(key[len(NFTMainKey):], NFTsOfOwnerKey)
	copy(key[len(NFTMainKey)+len(NFTsOfOwnerKey):], owner)
	copy(key[len(NFTMainKey)+len(NFTsOfOwnerKey)+len(owner):], Delimiter)
	copy(key[len(NFTMainKey)+len(NFTsOfOwnerKey)+len(owner)+len(Delimiter):], spaceIdStr)
	copy(key[len(NFTMainKey)+len(NFTsOfOwnerKey)+len(owner)+len(Delimiter)+len(spaceIdStr):], Delimiter)
	copy(key[len(NFTMainKey)+len(NFTsOfOwnerKey)+len(owner)+len(Delimiter)+len(spaceIdStr)+len(Delimiter):], classId)
	copy(key[len(NFTMainKey)+len(NFTsOfOwnerKey)+len(owner)+len(Delimiter)+len(spaceIdStr)+len(Delimiter)+len(classId):], Delimiter)
	copy(key[len(NFTMainKey)+len(NFTsOfOwnerKey)+len(owner)+len(Delimiter)+len(spaceIdStr)+len(Delimiter)+len(classId)+len(Delimiter):], tokenId)
	return key
}