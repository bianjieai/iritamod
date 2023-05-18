package cli

//import (
//	"github.com/bianjieai/iritamod/modules/layer2/types"
//	"github.com/cosmos/cosmos-sdk/client"
//	"github.com/spf13/cobra"
//)
//
//func NewTxCmd() *cobra.Command {
//	cmd := &cobra.Command{
//		Use:                        types.ModuleName,
//		Short:                      "Layer2 transaction subcommands",
//		DisableFlagParsing:         true,
//		SuggestionsMinimumDistance: 2,
//		RunE:                       client.ValidateCmd,
//	}
//
//	cmd.AddCommand(
//	)
//
//	return cmd
//}
//
//func GetCmdNftCmd() *cobra.Command {
//	cmd := &cobra.Command{
//		Use:                        "nft",
//		Short:                      "Layer2 NFT transaction subcommands",
//		DisableFlagParsing:         true,
//		SuggestionsMinimumDistance: 2,
//		RunE:                       client.ValidateCmd,
//	}
//}

// irita tx layer2 nft create token [space-id] [class-id] --nfts
// irita tx layer2 nft update tokens [space-id] [class-id] --nfts
// irita tx layer2 nft mint [space-id] [class-id] --nfts
// irita tx layer2 nft update classes [space-id] --classes
//

// irita tx layer2 nft withdraw token [space-id] [class-id] [token-id]
// irita tx layer2 nft deposit token [space-id] [class-id] [token-id]
// irita tx layer2 nft withdraw class [class-id]
// irita tx layer2 nft deposit class  [class-id]

// irita tx layer2 space create
// irita tx layer2 space transfer [space-id] [recipient]
// irita tx layer2 blockheader create [space-id] [height] [header]