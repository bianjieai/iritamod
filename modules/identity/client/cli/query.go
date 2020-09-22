package cli

import (
	"context"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/version"

	"gitlab.bianjie.ai/irita-pro/iritamod/modules/identity/types"
)

// GetQueryCmd returns the cli query commands for the module.
func GetQueryCmd() *cobra.Command {
	identityQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the identity module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	identityQueryCmd.AddCommand(
		GetCmdQueryIdentity(),
	)

	return identityQueryCmd
}

// GetCmdQueryIdentity implements the query identity command.
func GetCmdQueryIdentity() *cobra.Command {
	return &cobra.Command{
		Use:   "identity [id]",
		Short: "Query an identity",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query details of an identity with the specified ID.

Example:
$ %s query identity identity <id>
`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := hex.DecodeString(args[0])
			if err != nil {
				return err
			}

			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err = client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Identity(context.Background(), &types.QueryIdentityRequest{Id: id})
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res.Identity)
		},
	}
}
