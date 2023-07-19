package types

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

const (
	ModuleName = "sidechain"

	StoreKey = ModuleName

	RouterKey = ModuleName
)

var (
	// Space storekey prefix
	KeyPrefixSpaceSequence = []byte{0x01}
	KeyPrefixSpace         = []byte{0x02}
	KeyPrefixSpaceOfOwner  = []byte{0x03}

	// BlockHeader storekey prefix
	KeyPrefixBlockHeader             = []byte{0x04}
	KeyPrefixBlockHeaderTxHash       = []byte{0x05}
	KeyPrefixBlockHeaderLatestHeight = []byte{0x06}

	Delimiter   = []byte{0x00}
	Placeholder = []byte{0x01}
)

// space store key

// SpaceSequenceStoreKey returns the byte representation of the space sequence key
func SpaceSequenceStoreKey() []byte {
	return KeyPrefixSpaceSequence
}

// SpaceStoreKey returns the byte representation of the space key
// Items are stored with the following key: values
// <0x02><space_id>
func SpaceStoreKey(spaceId uint64) []byte {
	spaceIdStr := strconv.FormatUint(spaceId, 10)
	key := make([]byte, len(KeyPrefixSpace)+len(spaceIdStr))
	copy(key, KeyPrefixSpace)
	copy(key[len(KeyPrefixSpace):], spaceIdStr)
	return key
}

// SpaceOfOwnerStoreKey returns the byte representation of the space of user key
// Items are stored with the following key: values
// <0x03><owner><delimiter><space_id>
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

// <0x03><owner><delimiter>
func SpaceOfOwnerByOwnerStoreKey(owner sdk.AccAddress) []byte {
	owner = address.MustLengthPrefix(owner)
	key := make([]byte, len(KeyPrefixSpaceOfOwner)+len(owner)+len(Delimiter))
	copy(key, KeyPrefixSpaceOfOwner)
	copy(key[len(KeyPrefixSpaceOfOwner):], owner)
	copy(key[len(KeyPrefixSpaceOfOwner)+len(owner):], Delimiter)
	return key
}

// record mappings store key

// BlockHeaderStoreKey returns the byte representation of the block header key
// Items are stored with the following key: values
// <0x04><space_id><delimiter><block_height>
func BlockHeaderStoreKey(spaceId, blockHeight uint64) []byte {
	spaceIdStr := strconv.FormatUint(spaceId, 10)
	blockHeightStr := strconv.FormatUint(blockHeight, 10)
	key := make([]byte, len(KeyPrefixBlockHeader)+len(spaceIdStr)+len(Delimiter)+len(blockHeightStr))
	copy(key, KeyPrefixBlockHeader)
	copy(key[len(KeyPrefixBlockHeader):], spaceIdStr)
	copy(key[len(KeyPrefixBlockHeader)+len(spaceIdStr):], Delimiter)
	copy(key[len(KeyPrefixBlockHeader)+len(spaceIdStr)+len(Delimiter):], blockHeightStr)
	return key
}

// BlockHeaderTxHashStoreKey returns the byte representation of the block header tx hash key
// Items are stored with the following key: values
// <0x05><space_id><delimiter><block_height>
func BlockHeaderTxHashStoreKey(spaceId, blockHeight uint64) []byte {
	spaceIdStr := strconv.FormatUint(spaceId, 10)
	blockHeightStr := strconv.FormatUint(blockHeight, 10)
	key := make([]byte, len(KeyPrefixBlockHeaderTxHash)+len(spaceIdStr)+len(Delimiter)+len(blockHeightStr))
	copy(key, KeyPrefixBlockHeaderTxHash)
	copy(key[len(KeyPrefixBlockHeaderTxHash):], spaceIdStr)
	copy(key[len(KeyPrefixBlockHeaderTxHash)+len(spaceIdStr):], Delimiter)
	copy(key[len(KeyPrefixBlockHeaderTxHash)+len(spaceIdStr)+len(Delimiter):], blockHeightStr)
	return key
}

// BlockHeaderLatestHeightStoreKey returns the byte representation of the block header latest height key
// Items are stored with the following key: values
// <0x06><space_id>
func BlockHeaderLatestHeightStoreKey(spaceId uint64) []byte {
	spaceIdStr := strconv.FormatUint(spaceId, 10)
	key := make([]byte, len(KeyPrefixBlockHeaderLatestHeight)+len(spaceIdStr))
	copy(key, KeyPrefixBlockHeaderLatestHeight)
	copy(key[len(KeyPrefixBlockHeaderLatestHeight):], spaceIdStr)
	return key
}
