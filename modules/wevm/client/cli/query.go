package cli

import (
	"context"

	"github.com/bianjieai/iritamod/modules/wevm/types"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	wevmQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "wevm transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	ContractDenyListQueryCmd := &cobra.Command{
		Use:                        types.ContractDenyListName,
		Short:                      "wevm transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	ContractDenyListQueryCmd.AddCommand(
		NewGetContractDenyList(),
		NewGetContractState(),
	)
	wevmQueryCmd.AddCommand(ContractDenyListQueryCmd)
	return wevmQueryCmd
}

func NewGetContractDenyList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deny-list  [flags]",
		Short: "get all contract state",
		Long:  "get all contract state",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryContractDenyListRequest{}
			res, err := queryClient.ContractDenyList(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	return cmd
}
func NewGetContractState() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "contract [contractAddress] [flags]",
		Short: "get contract state",
		Long:  "get contract state by contract address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			contractAddr := args[0]
			req := &types.QueryContractStateRequest{
				Address: contractAddr,
			}
			res, err := queryClient.ContractState(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	return cmd
}
