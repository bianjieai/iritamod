package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	"github.com/bianjieai/iritamod/modules/wevm/types"
)

func NewTxCmd() *cobra.Command {
	wevmTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "wevm transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	ContractDenyListTxCmd := &cobra.Command{
		Use:                        types.ContractDenyListName,
		Short:                      "wevm transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	ContractDenyListTxCmd.AddCommand(
		NewAddToContractDenyList(),
		NewRemoveFromContractDenyList(),
	)
	wevmTxCmd.AddCommand(ContractDenyListTxCmd)

	return wevmTxCmd
}

func NewAddToContractDenyList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add [contractAddress] [flags]",
		Args:  cobra.ExactArgs(1),
		Short: "add contract address to contract deny list",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			sender := clientCtx.GetFromAddress().String()
			contractAddr := args[0]
			msg := types.NewMsgAddToContractDenyList(
				contractAddr,
				sender,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewRemoveFromContractDenyList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove [contractAddress] [flags]",
		Args:  cobra.ExactArgs(1),
		Short: "remove contract address from contract deny list",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			sender := clientCtx.GetFromAddress().String()
			contractAddr := args[0]
			msg := types.NewMsgRemoveFromContractDenyList(
				contractAddr,
				sender,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
