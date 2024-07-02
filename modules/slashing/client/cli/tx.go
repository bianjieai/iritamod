package cli

import (
	"encoding/hex"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"

	"github.com/bianjieai/iritamod/modules/slashing/types"
)

func NewTxCmd() *cobra.Command {
	slashingTxCmd := &cobra.Command{
		Use:                        slashingtypes.ModuleName,
		Short:                      "Slashing transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	slashingTxCmd.AddCommand(
		NewUnjailValidatorCmd(),
	)

	return slashingTxCmd
}

// NewUnjailValidatorCmd implements the unjail validator command handler.
func NewUnjailValidatorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unjail [id]",
		Short: "Unjail an jailed validator",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := hex.DecodeString(args[0])
			if err != nil {
				return fmt.Errorf("invalid validator id:%s", args[0])
			}

			msg := types.NewMsgUnjailValidator(
				id,
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
