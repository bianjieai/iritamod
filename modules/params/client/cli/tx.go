package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/bianjieai/iritamod/modules/params/client/utils"
	"github.com/bianjieai/iritamod/modules/params/types"
)

func NewTxCmd() *cobra.Command {
	paramsTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Params transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	paramsTxCmd.AddCommand(
		NewUpdateParamsCmd(),
	)

	return paramsTxCmd
}

// NewUpdateParamsCmd implements the update params command handler.
func NewUpdateParamsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update [change-file]",
		Short: "UpdateValidator params",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			changes, _ := utils.ParseParamChange(clientCtx.LegacyAmino, args[0])

			msg := types.NewMsgUpdateParams(
				changes.ToParamChanges(),
				clientCtx.GetFromAddress(),
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}
