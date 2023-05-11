package types

const (
	ModuleName = "layer2"

	StoreKey = ModuleName

	RouterKey = ModuleName
)

var (
	SpaceMainKey = []byte{0x01}
	BlockMainKey = []byte{0x02}
	NFTMainKey = []byte{0x03}

	// Space Subkey
	Spacekey = []byte{0x01}
	SpaceOfUserKey = []byte{0x02}

	// BlockHeader Subkey
	// TODO

	// NFT Subkey
	ClassKey = []byte{0x01}
	NFTKey = []byte{0x02}
	NFTOfClassOfSpaceByOwnerKey = []byte{0x03}

	Delimiter = []byte{0x00}
	Placeholder = []byte{0x01}
)
