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
	// Space storekey prefix
	SpaceKey = []byte{0x01}
	SpaceOfUserKey = []byte{0x02}
	// Record storekey prefix
	RecordKey = []byte{0x03}
	// NFT storekey prefix
	NftClassKey = []byte{0x04}
	NftTokenKey = []byte{0x05}
	NftsOfOwnerKey = []byte{0x06}

	Delimiter = []byte{0x00}
	Placeholder = []byte{0x01}
)

// space store key

// spaceStoreKey returns the byte representation of the space key
// Items are stored with the following key: values
// <0x01><space_id>
func spaceStoreKey(spaceId uint64) []byte {
	spaceIdStr := strconv.FormatUint(spaceId, 10)
	key := make([]byte, len(SpaceKey)+len(spaceIdStr))
	copy(key, SpaceKey)
	copy(key[len(SpaceKey):], spaceIdStr)
	return key
}

// spaceOfL2UserStoreKey returns the byte representation of the space of l2 user key
// Items are stored with the following key: values
// <0x02><owner><delimiter><space_id>
func spaceOfL2UserStoreKey(owner sdk.AccAddress, spaceId uint64) []byte {
	owner = address.MustLengthPrefix(owner)
	spaceIdStr := strconv.FormatUint(spaceId, 10)
	key := make([]byte, len(SpaceOfUserKey)+len(owner)+len(Delimiter)+len(spaceIdStr))
	copy(key, SpaceOfUserKey)
	copy(key[len(SpaceOfUserKey):], owner)
	copy(key[len(SpaceOfUserKey)+len(owner):], Delimiter)
	copy(key[len(SpaceOfUserKey)+len(owner)+len(Delimiter):], spaceIdStr)
	return key
}

// record mappings store key

// recordStoreKey returns the byte representation of the record key
// Items are stored with the following key: values
// <0x03><space_id><delimiter><block_height>
func recordStoreKey(spaceId, blockHeight uint64) []byte {
	spaceIdStr := strconv.FormatUint(spaceId, 10)
	blockHeightStr := strconv.FormatUint(blockHeight, 10)
	key := make([]byte, len(RecordKey)+len(spaceIdStr)+len(Delimiter)+len(blockHeightStr))
	copy(key, RecordKey)
	copy(key[len(RecordKey):], spaceIdStr)
	copy(key[len(RecordKey)+len(spaceIdStr):], Delimiter)
	copy(key[len(RecordKey)+len(spaceIdStr)+len(Delimiter):], blockHeightStr)
	return key
}

// nft mappings store key

// nftClassStoreKey returns the byte representation of the class key of nft mappings
// Items are stored with the following key: values
// <0x04><space_id>
func nftClassStoreKey(classId string) []byte {
	key := make([]byte,len(NftClassKey)+len(classId))
	copy(key, NftClassKey)
	copy(key[len(NftClassKey):], classId)
	return key
}

// nftTokenStoreKey returns the byte representation of the nft key of nft mappings
// Items are stored with the following key: values
// <0x05><space_id><delimiter><class_id><delimiter><token_id>
func nftTokenStoreKey(spaceId uint64, classId, tokenId string) []byte {
	spaceIdStr := strconv.FormatUint(spaceId, 10)

	key := make([]byte, len(NftTokenKey)+
		len(spaceIdStr)+len(Delimiter)+
		len(classId)+len(Delimiter)+
		len(tokenId))

	copy(key, NftTokenKey)
	copy(key[len(NftTokenKey):], spaceIdStr)
	copy(key[len(NftTokenKey)+len(spaceIdStr):], Delimiter)
	copy(key[len(NftTokenKey)+len(spaceIdStr)+len(Delimiter):], classId)
	copy(key[len(NftTokenKey)+len(spaceIdStr)+len(Delimiter)+len(classId):], Delimiter)
	copy(key[len(NftTokenKey)+len(spaceIdStr)+len(Delimiter)+len(classId)+len(Delimiter):], tokenId)
	return key
}

// nftsOfOwnerStoreKey returns the byte representation of the owner key of nft mappings
// Items are stored with the following key: values
// <0x06><owner><delimiter><space_id><delimiter><class_id><delimiter><nft_id>
func nftsOfOwnerStoreKey(owner sdk.AccAddress, spaceId uint64, classId, tokenId string) []byte{
	owner = address.MustLengthPrefix(owner)
	spaceIdStr := strconv.FormatUint(spaceId, 10)

	key := make([]byte, len(NftsOfOwnerKey)+
		len(owner)+len(Delimiter)+
		len(spaceIdStr)+len(Delimiter)+
		len(classId)+len(Delimiter)+
		len(tokenId))

	copy(key, NftsOfOwnerKey)
	copy(key[len(NftsOfOwnerKey):], owner)
	copy(key[len(NftsOfOwnerKey)+len(owner):], Delimiter)
	copy(key[len(NftsOfOwnerKey)+len(owner)+len(Delimiter):], spaceIdStr)
	copy(key[len(NftsOfOwnerKey)+len(owner)+len(Delimiter)+len(spaceIdStr):], Delimiter)
	copy(key[len(NftsOfOwnerKey)+len(owner)+len(Delimiter)+len(spaceIdStr)+len(Delimiter):], classId)
	copy(key[len(NftsOfOwnerKey)+len(owner)+len(Delimiter)+len(spaceIdStr)+len(Delimiter)+len(classId):], Delimiter)
	copy(key[len(NftsOfOwnerKey)+len(owner)+len(Delimiter)+len(spaceIdStr)+len(Delimiter)+len(classId)+len(Delimiter):], tokenId)
	return key
}