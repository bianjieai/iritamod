package types

import (
	"bytes"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

const (
	ModuleName = "layer2"

	StoreKey = ModuleName

	RouterKey = ModuleName
)

var (
	ModuleAddress    sdk.Address
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
	KeyPrefixSpace        = []byte{0x01}
	KeyPrefixSpaceOfOwner = []byte{0x02}
	// BlockHeader storekey prefix
	KeyPrefixL2BlockHeader = []byte{0x03}
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

// <0x02><owner><delimiter>
func SpaceOfOwnerByOwnerStoreKey(owner sdk.AccAddress) []byte {
	owner = address.MustLengthPrefix(owner)
	key := make([]byte, len(KeyPrefixSpaceOfOwner)+len(owner)+len(Delimiter))
	copy(key, KeyPrefixSpaceOfOwner)
	copy(key[len(KeyPrefixSpaceOfOwner):], owner)
	copy(key[len(KeyPrefixSpaceOfOwner)+len(owner):], Delimiter)
	return key
}

// record mappings store key

// L2BlockHeaderStoreKey returns the byte representation of the record key
// Items are stored with the following key: values
// <0x03><space_id><delimiter><block_height>
func L2BlockHeaderStoreKey(spaceId, blockHeight uint64) []byte {
	spaceIdStr := strconv.FormatUint(spaceId, 10)
	blockHeightStr := strconv.FormatUint(blockHeight, 10)
	key := make([]byte, len(KeyPrefixL2BlockHeader)+len(spaceIdStr)+len(Delimiter)+len(blockHeightStr))
	copy(key, KeyPrefixL2BlockHeader)
	copy(key[len(KeyPrefixL2BlockHeader):], spaceIdStr)
	copy(key[len(KeyPrefixL2BlockHeader)+len(spaceIdStr):], Delimiter)
	copy(key[len(KeyPrefixL2BlockHeader)+len(spaceIdStr)+len(Delimiter):], blockHeightStr)
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

// <0x05><space_id><delimiter><class_id><delimiter>
func TokenForNFTByCollectionStoreKey(spaceId uint64, classId string) []byte {
	spaceIdStr := strconv.FormatUint(spaceId, 10)

	key := make([]byte, len(KeyPrefixTokenForNFT)+
		len(spaceIdStr)+len(Delimiter)+
		len(classId)+len(Delimiter))

	copy(key, KeyPrefixTokenForNFT)
	copy(key[len(KeyPrefixTokenForNFT):], spaceIdStr)
	copy(key[len(KeyPrefixTokenForNFT)+len(spaceIdStr):], Delimiter)
	copy(key[len(KeyPrefixTokenForNFT)+len(spaceIdStr)+len(Delimiter):], classId)
	copy(key[len(KeyPrefixTokenForNFT)+len(spaceIdStr)+len(Delimiter)+len(classId):], Delimiter)
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

func NFTsOfOwnerAllStoreKey(owner sdk.AccAddress) []byte {
	owner = address.MustLengthPrefix(owner)

	key := make([]byte, len(KeyPrefixNFTsOfOwner)+
		len(owner)+len(Delimiter))

	copy(key, KeyPrefixNFTsOfOwner)
	copy(key[len(KeyPrefixNFTsOfOwner):], owner)
	copy(key[len(KeyPrefixNFTsOfOwner)+len(owner):], Delimiter)
	return key
}

// ret =  <space_id> + <class_id> + <nft_id>
func ParseNFTsOfOwnerAllStoreKey(key []byte) (spaceId uint64, classId, tokenId string) {
	ret := bytes.Split(key, Delimiter)
	if len(ret) != 3 {
		panic("invalid NFTsOfOwnerAllStoreKey")
	}
	spaceId, err := strconv.ParseUint(string(ret[0]), 10, 64)
	if err != nil {
		panic("fail to convert spaceId from string to uint64")
	}
	classId = string(ret[1])
	tokenId = string(ret[2])
	return
}

func NFTsOfOwnerBySpaceStoreKey(owner sdk.AccAddress, spaceId uint64) []byte {
	owner = address.MustLengthPrefix(owner)
	spaceIdStr := strconv.FormatUint(spaceId, 10)

	key := make([]byte, len(KeyPrefixNFTsOfOwner)+
		len(owner)+len(Delimiter)+
		len(spaceIdStr)+len(Delimiter))

	copy(key, KeyPrefixNFTsOfOwner)
	copy(key[len(KeyPrefixNFTsOfOwner):], owner)
	copy(key[len(KeyPrefixNFTsOfOwner)+len(owner):], Delimiter)
	copy(key[len(KeyPrefixNFTsOfOwner)+len(owner)+len(Delimiter):], spaceIdStr)
	copy(key[len(KeyPrefixNFTsOfOwner)+len(owner)+len(Delimiter)+len(spaceIdStr):], Delimiter)
	return key
}

// ret =   <class_id> + <nft_id>
func ParseNFTsOfOwnerBySpaceStoreKey(key []byte) (classId, tokenId string) {
	ret := bytes.Split(key, Delimiter)
	if len(ret) != 2 {
		panic("invalid NFTsOfOwnerAllStoreKey")
	}

	classId = string(ret[0])
	tokenId = string(ret[1])
	return
}

func NFTsOfOwnerBySpaceAndClassStoreKey(owner sdk.AccAddress, spaceId uint64, classId string) []byte {
	owner = address.MustLengthPrefix(owner)
	spaceIdStr := strconv.FormatUint(spaceId, 10)

	key := make([]byte, len(KeyPrefixNFTsOfOwner)+
		len(owner)+len(Delimiter)+
		len(spaceIdStr)+len(Delimiter)+
		len(classId)+len(Delimiter))

	copy(key, KeyPrefixNFTsOfOwner)
	copy(key[len(KeyPrefixNFTsOfOwner):], owner)
	copy(key[len(KeyPrefixNFTsOfOwner)+len(owner):], Delimiter)
	copy(key[len(KeyPrefixNFTsOfOwner)+len(owner)+len(Delimiter):], spaceIdStr)
	copy(key[len(KeyPrefixNFTsOfOwner)+len(owner)+len(Delimiter)+len(spaceIdStr):], Delimiter)
	copy(key[len(KeyPrefixNFTsOfOwner)+len(owner)+len(Delimiter)+len(spaceIdStr)+len(Delimiter):], classId)
	copy(key[len(KeyPrefixNFTsOfOwner)+len(owner)+len(Delimiter)+len(spaceIdStr)+len(Delimiter)+len(classId):], Delimiter)
	return key
}
