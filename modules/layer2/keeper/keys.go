package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"

	"strconv"
)

var (
	// Space storekey prefix
	KeyPrefixSpace = []byte{0x01}
	KeyPrefixSpaceOfUser = []byte{0x02}
	// Record storekey prefix
	KeyPrefixRecord = []byte{0x03}
	// NFT storekey prefix
	KeyPrefixClassForNFT = []byte{0x04}
	KeyPrefixTokenForNFT = []byte{0x05}
	KeyPrefixNFTsOfOwner = []byte{0x06}

	Delimiter = []byte{0x00}
	Placeholder = []byte{0x01}
)

// space store key

// spaceStoreKey returns the byte representation of the space key
// Items are stored with the following key: values
// <0x01><space_id>
func spaceStoreKey(spaceId uint64) []byte {
	spaceIdStr := strconv.FormatUint(spaceId, 10)
	key := make([]byte, len(KeyPrefixSpace)+len(spaceIdStr))
	copy(key, KeyPrefixSpace)
	copy(key[len(KeyPrefixSpace):], spaceIdStr)
	return key
}

// spaceOfL2UserStoreKey returns the byte representation of the space of l2 user key
// Items are stored with the following key: values
// <0x02><owner><delimiter><space_id>
func spaceOfL2UserStoreKey(owner sdk.AccAddress, spaceId uint64) []byte {
	owner = address.MustLengthPrefix(owner)
	spaceIdStr := strconv.FormatUint(spaceId, 10)
	key := make([]byte, len(KeyPrefixSpaceOfUser)+len(owner)+len(Delimiter)+len(spaceIdStr))
	copy(key, KeyPrefixSpaceOfUser)
	copy(key[len(KeyPrefixSpaceOfUser):], owner)
	copy(key[len(KeyPrefixSpaceOfUser)+len(owner):], Delimiter)
	copy(key[len(KeyPrefixSpaceOfUser)+len(owner)+len(Delimiter):], spaceIdStr)
	return key
}

// record mappings store key

// recordStoreKey returns the byte representation of the record key
// Items are stored with the following key: values
// <0x03><space_id><delimiter><block_height>
func recordStoreKey(spaceId, blockHeight uint64) []byte {
	spaceIdStr := strconv.FormatUint(spaceId, 10)
	blockHeightStr := strconv.FormatUint(blockHeight, 10)
	key := make([]byte, len(KeyPrefixRecord)+len(spaceIdStr)+len(Delimiter)+len(blockHeightStr))
	copy(key, KeyPrefixRecord)
	copy(key[len(KeyPrefixRecord):], spaceIdStr)
	copy(key[len(KeyPrefixRecord)+len(spaceIdStr):], Delimiter)
	copy(key[len(KeyPrefixRecord)+len(spaceIdStr)+len(Delimiter):], blockHeightStr)
	return key
}

// nft mappings store key

// classForNFTStoreKey returns the byte representation of the class key of nft mappings
// Items are stored with the following key: values
// <0x04><space_id>
func classForNFTStoreKey(classId string) []byte {
	key := make([]byte,len(KeyPrefixClassForNFT)+len(classId))
	copy(key, KeyPrefixClassForNFT)
	copy(key[len(KeyPrefixClassForNFT):], classId)
	return key
}

// tokenForNFTStoreKey returns the byte representation of the nft key of nft mappings
// Items are stored with the following key: values
// <0x05><space_id><delimiter><class_id><delimiter><token_id>
func tokenForNFTStoreKey(spaceId uint64, classId, tokenId string) []byte {
	spaceIdStr := strconv.FormatUint(spaceId, 10)

	key := make([]byte, len(KeyPrefixTokenForNFT)+
		len(spaceIdStr)+len(Delimiter)+
		len(classId)+len(Delimiter)+
		len(tokenId))

	copy(key, KeyPrefixTokenForNFT)
	copy(key[len(KeyPrefixTokenForNFT):], spaceIdStr)
	copy(key[len(KeyPrefixTokenForNFT)+len(spaceIdStr):], Delimiter)
	copy(key[len(KeyPrefixTokenForNFT)+len(spaceIdStr)+len(Delimiter):], classId)
	copy(key[len(KeyPrefixTokenForNFT)+len(spaceIdStr)+len(Delimiter)+len(classId):], Delimiter)
	copy(key[len(KeyPrefixTokenForNFT)+len(spaceIdStr)+len(Delimiter)+len(classId)+len(Delimiter):], tokenId)
	return key
}

// nftsOfOwnerStoreKey returns the byte representation of the owner key of nft mappings
// Items are stored with the following key: values
// <0x06><owner><delimiter><space_id><delimiter><class_id><delimiter><nft_id>
func nftsOfOwnerStoreKey(owner sdk.AccAddress, spaceId uint64, classId, tokenId string) []byte{
	owner = address.MustLengthPrefix(owner)
	spaceIdStr := strconv.FormatUint(spaceId, 10)

	key := make([]byte, len(KeyPrefixNFTsOfOwner)+
		len(owner)+len(Delimiter)+
		len(spaceIdStr)+len(Delimiter)+
		len(classId)+len(Delimiter)+
		len(tokenId))

	copy(key, KeyPrefixNFTsOfOwner)
	copy(key[len(KeyPrefixNFTsOfOwner):], owner)
	copy(key[len(KeyPrefixNFTsOfOwner)+len(owner):], Delimiter)
	copy(key[len(KeyPrefixNFTsOfOwner)+len(owner)+len(Delimiter):], spaceIdStr)
	copy(key[len(KeyPrefixNFTsOfOwner)+len(owner)+len(Delimiter)+len(spaceIdStr):], Delimiter)
	copy(key[len(KeyPrefixNFTsOfOwner)+len(owner)+len(Delimiter)+len(spaceIdStr)+len(Delimiter):], classId)
	copy(key[len(KeyPrefixNFTsOfOwner)+len(owner)+len(Delimiter)+len(spaceIdStr)+len(Delimiter)+len(classId):], Delimiter)
	copy(key[len(KeyPrefixNFTsOfOwner)+len(owner)+len(Delimiter)+len(spaceIdStr)+len(Delimiter)+len(classId)+len(Delimiter):], tokenId)
	return key
}