package cli

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/bianjieai/iritamod/modules/perm/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	permQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the perm module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	permQueryCmd.AddCommand(
		GetCmdQueryRoles(),
		GetCmdQueryAccountBlackList(),
		GetCmdQueryContractBlockList(),
	)

	return permQueryCmd
}

// GetCmdQueryRoles implements the roles query command.
func GetCmdQueryRoles() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "roles [account]",
		Short: "Query a account roles",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Roles(context.Background(), &types.QueryRolesRequest{Address: args[0]})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryAccountBlackList implements the black list query command.
func GetCmdQueryAccountBlackList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "block-list-account",
		Short: "Query account blockList",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Blacklist(context.Background(), &types.QueryBlacklistRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetCmdQueryContractBlockList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "block-list-contract  [flags]",
		Short: "Query contract blockList ",
		Long:  "Query contract blockList",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryContractDenyList{}
			res, err := queryClient.ContractDenyList(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	return cmd
}
