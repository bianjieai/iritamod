package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	upgradecli "github.com/cosmos/cosmos-sdk/x/upgrade/client/cli"

	"github.com/bianjieai/iritamod/modules/upgrade/types"
)

func GetTxCmd() *cobra.Command {
	upgradeTxCmd := &cobra.Command{
		Use:                        "upgrade",
		Short:                      "upgrade transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	upgradeTxCmd.AddCommand(
		GetCmdUpgradeSoftware(),
		GetCmdCancelSoftwareUpgrade(),
	)

	return upgradeTxCmd
}

// GetCmdUpgradeSoftware implements the upgrade software command handler.
func GetCmdUpgradeSoftware() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [name]",
		Short: "create a software upgrade plan",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			planHeight, err := cmd.Flags().GetInt64(upgradecli.FlagUpgradeHeight)
			if err != nil {
				return err
			}

			info, err := cmd.Flags().GetString(upgradecli.FlagUpgradeInfo)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpgradeSoftware(
				args[0],
				planHeight,
				info,
				clientCtx.GetFromAddress(),
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().Int64(upgradecli.FlagUpgradeHeight, 0, "The height at which the upgrade must happen (not to be used together with --upgrade-time)")
	cmd.Flags().String(upgradecli.FlagUpgradeInfo, "", "Optional info for the planned upgrade such as commit hash, etc.")
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetCmdCancelSoftwareUpgrade implements the cancel software upgrade command handler.
func GetCmdCancelSoftwareUpgrade() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cancel",
		Short: "cancel current upgrade plan",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCancelUpgrade(
				clientCtx.GetFromAddress(),
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	_ = cmd.MarkFlagRequired(flags.FlagFrom)
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
