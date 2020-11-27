package cli

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/bianjieai/iritamod/modules/node/types"
)

// NewTxCmd returns the transaction commands for this module
func NewTxCmd() *cobra.Command {
	nodeTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Node transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	nodeTxCmd.AddCommand(
		NewAddNodeCmd(),
		NewRemoveNodeCmd(),
	)

	return nodeTxCmd
}

// NewAddNodeCmd implements adding a node command
func NewAddNodeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a node to the node whitelist",
		Long:  "Add a node to the node whitelist based on the identity certificate",
		Example: fmt.Sprintf(
			"$ %s tx node add --cert=<certificate-file> --from mykey",
			version.AppName,
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			operator := clientCtx.GetFromAddress()

			certFile := strings.TrimSpace(viper.GetString(FlagCert))
			if len(certFile) == 0 {
				return fmt.Errorf("the certificate file path must not be empty")
			}

			cert, err := ioutil.ReadFile(certFile)
			if err != nil {
				return fmt.Errorf("failed to read the certificate file: %s", err.Error())
			}

			msg := types.NewMsgAddNode(string(cert), operator)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FsAddNode)
	cmd.MarkFlagRequired(FlagCert)

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewRemoveNodeCmd implements removing a node command
func NewRemoveNodeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove [id]",
		Short: "Remove a node from the node whitelist",
		Long:  "Remove a node from the node whitelist by id",
		Example: fmt.Sprintf(
			"$ %s tx node remove <id> --from mykey",
			version.AppName,
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			operator := clientCtx.GetFromAddress()

			id, err := hex.DecodeString(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgRemoveNode(id, operator)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
