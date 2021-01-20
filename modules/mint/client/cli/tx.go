package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/bianjieai/iritamod/modules/mint/types"
)

func NewTxCmd() *cobra.Command {
	mintTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Mint transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	mintTxCmd.AddCommand(
		NewMintCmd(),
	)

	return mintTxCmd
}

// NewMintCmd implements the mint command.
func NewMintCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint [amount]",
		Short: "Mint the base native token",
		Long: strings.TrimSpace(fmt.Sprintf(
			"Mint the base native token by the given amount in main unit",
		)),
		Example: fmt.Sprintf(
			"$ %s tx %s mint <amount> --from mykey",
			version.AppName, types.ModuleName,
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			operator := clientCtx.GetFromAddress()

			amount, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgMint(amount, operator)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
