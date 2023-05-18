package cli

//import (
//	"fmt"
//	"github.com/cosmos/cosmos-sdk/version"
//	"github.com/spf13/cobra"
//)
//
//func GetCmdNftCreateToken() *cobra.Command {
//	cmd := &cobra.Command{
//		Use:   "create token [space-id] [class-id]",
//		Short: "Create a new token",
//		Example: fmt.Sprintf(
//			"$ %s tx layer2 nft create token [space-id] [class-id]" +
//				"--nfts=<nfts.json>" +
//				version.AppName),
//		Args:  cobra.ExactArgs(2),
//		RunE: func(cmd *cobra.Command, args []string) error {
//			return nil
//		},
//	}
//
//	cmd.FLags().AddFlagSet()
//
//	return cmd
//}