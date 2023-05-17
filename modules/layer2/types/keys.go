package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"strconv"
)

const (
	ModuleName = "layer2"

	StoreKey = ModuleName

	RouterKey = ModuleName
)

var (
	ModuleAddress sdk.Address
	ModuleAddressStr string
	ModuleAccAddress sdk.AccAddress
)

func init() {
	ModuleAddress = authtypes.NewModuleAddress(ModuleName)
	ModuleAddressStr = ModuleAddress.String()
	acc, err := sdk.AccAddressFromBech32(ModuleAddressStr)
	if err != nil {
		panic(err.Error())
	}
	ModuleAccAddress = acc
}

var (
	// Space storekey prefix
	KeyPrefixSpace       = []byte{0x01}
	KeyPrefixSpaceOfOwner = []byte{0x02}
	// Record storekey prefix
	KeyPrefixRecord = []byte{0x03}
	// NFT storekey prefix
	KeyPrefixClassForNFT = []byte{0x04}
	KeyPrefixTokenForNFT = []byte{0x05}
	KeyPrefixNFTsOfOwner = []byte{0x06}

	Delimiter   = []byte{0x00}
	Placeholder = []byte{0x01}
)

// space store key

// SpaceStoreKey returns the byte representation of the space key
// Items are stored with the following key: values
// <0x01><space_id>
func SpaceStoreKey(spaceId uint64) []byte {
	spaceIdStr := strconv.FormatUint(spaceId, 10)
	key := make([]byte, len(KeyPrefixSpace)+len(spaceIdStr))
	copy(key, KeyPrefixSpace)
	copy(key[len(KeyPrefixSpace):], spaceIdStr)
	return key
}

// SpaceOfOwnerStoreKey returns the byte representation of the space of l2 user key
// Items are stored with the following key: values
// <0x02><owner><delimiter><space_id>
func SpaceOfOwnerStoreKey(owner sdk.AccAddress, spaceId uint64) []byte {
	owner = address.MustLengthPrefix(owner)
	spaceIdStr := strconv.FormatUint(spaceId, 10)
	key := make([]byte, len(KeyPrefixSpaceOfOwner)+len(owner)+len(Delimiter)+len(spaceIdStr))
	copy(key, KeyPrefixSpaceOfOwner)
	copy(key[len(KeyPrefixSpaceOfOwner):], owner)
	copy(key[len(KeyPrefixSpaceOfOwner)+len(owner):], Delimiter)
	copy(key[len(KeyPrefixSpaceOfOwner)+len(owner)+len(Delimiter):], spaceIdStr)
	return key
}

// record mappings store key

// L2BlockHeaderStoreKey returns the byte representation of the record key
// Items are stored with the following key: values
// <0x03><space_id><delimiter><block_height>
func L2BlockHeaderStoreKey(spaceId, blockHeight uint64) []byte {
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

// ClassForNFTStoreKey returns the byte representation of the class key of nft mappings
// Items are stored with the following key: values
// <0x04><class_id>
func ClassForNFTStoreKey(classId string) []byte {
	key := make([]byte, len(KeyPrefixClassForNFT)+len(classId))
	copy(key, KeyPrefixClassForNFT)
	copy(key[len(KeyPrefixClassForNFT):], classId)
	return key
}

// TokenForNFTStoreKey returns the byte representation of the nft key of nft mappings
// Items are stored with the following key: values
// <0x05><space_id><delimiter><class_id><delimiter><token_id>
func TokenForNFTStoreKey(spaceId uint64, classId, tokenId string) []byte {
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

// NFTsOfOwnerStoreKey returns the byte representation of the owner key of nft mappings
// Items are stored with the following key: values
// <0x06><owner><delimiter><space_id><delimiter><class_id><delimiter><nft_id>
func NFTsOfOwnerStoreKey(owner sdk.AccAddress, spaceId uint64, classId, tokenId string) []byte {
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
