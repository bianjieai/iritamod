package cli

import flag "github.com/spf13/pflag"

const (
	FlagSpaceName = "name"
	FlagSpaceUri  = "uri"
	FlagSpaceId   = "space-id"

	FlagNftTokens       = "nfts"
	FlagNftTokenId      = "token-id"
	FlagNftTokenIds     = "nft-ids"
	FlagNftTokenOwner   = "owner"
	FlagNftTokenName    = "name"
	FlagNftTokenUri     = "uri"
	FlagNftTokenUriHash = "uri-hash"
	FlagNftTokenData    = "data"

	FlagNftClassInfos   = "class-infos"
	FlagNftClassBaseUri = "base-uri"
	FlagNftClassOwner   = "owner"
	FlagNftClassId      = "class-id"
)

var (
	FsCreateSpace      = flag.NewFlagSet("", flag.ContinueOnError)
	FsNftCreateTokens  = flag.NewFlagSet("", flag.ContinueOnError)
	FsNftUpdateTokens  = flag.NewFlagSet("", flag.ContinueOnError)
	FsNftDeleteTokens  = flag.NewFlagSet("", flag.ContinueOnError)
	FsNftUpdateClasses = flag.NewFlagSet("", flag.ContinueOnError)
	FsNftWithdrawToken = flag.NewFlagSet("", flag.ContinueOnError)
	FsNftDepositClass  = flag.NewFlagSet("", flag.ContinueOnError)
	FsNftWithdrawClass = flag.NewFlagSet("", flag.ContinueOnError)

	FsQueryNftOwner = flag.NewFlagSet("", flag.ContinueOnError)
	FsQueryNftUri   = flag.NewFlagSet("", flag.ContinueOnError)
)

func init() {
	FsCreateSpace.String(FlagSpaceName, "", "name of the space")
	FsCreateSpace.String(FlagSpaceUri, "", "uri of the space")

	FsNftCreateTokens.String(FlagNftTokens, "", "nft mappings to create")
	FsNftUpdateTokens.String(FlagNftTokens, "", "nft mappings to update")
	FsNftDeleteTokens.String(FlagNftTokenIds, "", "nft ids to delete")
	FsNftUpdateClasses.String(FlagNftClassInfos, "", "nft class infos to update")

	FsNftWithdrawToken.String(FlagNftTokenOwner, "", "owner of the nft token")
	FsNftWithdrawToken.String(FlagNftTokenName, "", "name of the nft token")
	FsNftWithdrawToken.String(FlagNftTokenUri, "", "uri of the nft token")
	FsNftWithdrawToken.String(FlagNftTokenUriHash, "", "uri hash of the nft token")
	FsNftWithdrawToken.String(FlagNftTokenData, "", "data of the nft token")

	FsNftDepositClass.String(FlagNftClassBaseUri, "", "base uri of the nft class")
	FsNftWithdrawClass.String(FlagNftClassOwner, "", "owner of the nft class")

	FsQueryNftOwner.String(FlagNftClassId, "", "class id of the nft")
	FsQueryNftUri.String(FlagNftTokenId, "", "token id of the nft")
	FsQueryNftUri.String(FlagSpaceId, "", "space id")
}
