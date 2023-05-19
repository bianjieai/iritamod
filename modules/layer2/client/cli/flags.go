package cli

import flag "github.com/spf13/pflag"

const (
	FlagId = "id"
	FlagIds = "ids"
	FlagOwner = "owner"
	FlagOwners = "owners"
	FlagName = "name"
	FlagUri = "uri"
	FlagUriHash = "uri-hash"
	FlagUris = "uris"
	FlagData = "data"
)

var (
	FsSpaceCreate      = flag.NewFlagSet("", flag.ContinueOnError)
	FsNftTokenCreate  = flag.NewFlagSet("", flag.ContinueOnError)
	FsNftTokenUpdate  = flag.NewFlagSet("", flag.ContinueOnError)
	FsNftTokenDelete  = flag.NewFlagSet("", flag.ContinueOnError)
	FsNftTokenWithdraw = flag.NewFlagSet("", flag.ContinueOnError)
	FsNftClassUpdate = flag.NewFlagSet("", flag.ContinueOnError)
	FsNftClassDeposit  = flag.NewFlagSet("", flag.ContinueOnError)
	FsNftClassWithdraw = flag.NewFlagSet("", flag.ContinueOnError)

	FsQueryNftOwner = flag.NewFlagSet("", flag.ContinueOnError)
)

func init() {
	FsSpaceCreate.String(FlagName, "", "name of the space")
	FsSpaceCreate.String(FlagUri, "", "uri of the space")

	FsNftTokenCreate.String(FlagIds, "", "token ids for nft asset")
	FsNftTokenCreate.String(FlagOwners, "", "token owners for nft asset")
	FsNftTokenUpdate.String(FlagIds, "", "token ids for nft asset")
	FsNftTokenUpdate.String(FlagOwners, "", "token owners for nft asset")
	FsNftTokenDelete.String(FlagIds, "", "token ids for nft asset")

	FsNftTokenWithdraw.String(FlagOwner, "", "owner of the nft")
	FsNftTokenWithdraw.String(FlagName, "", "name of the nft")
	FsNftTokenWithdraw.String(FlagUri, "", "uri of the nft")
	FsNftTokenWithdraw.String(FlagUriHash, "", "uri hash of the nft")
	FsNftTokenWithdraw.String(FlagData, "", "data of the nft")

	FsNftClassUpdate.String(FlagIds, "", "class ids for nft asset")
	FsNftClassUpdate.String(FlagUris, "", "class base uris for nft asset")
	FsNftClassUpdate.String(FlagOwners, "", "class owners for nft asset")

	FsNftClassDeposit.String(FlagUri, "", "base uri of the class")
	FsNftClassWithdraw.String(FlagOwner, "", "owner of the class")

	FsQueryNftOwner.String(FlagId, "", "class id of the nft")
}
