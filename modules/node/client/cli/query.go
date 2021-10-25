package cli

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/bianjieai/iritamod/modules/node/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	nodeQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the node module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	nodeQueryCmd.AddCommand(
		GetCmdQueryValidator(),
		GetCmdQueryValidators(),
		GetCmdQueryNode(),
		GetCmdQueryNodes(),
		GetCmdQueryParams(),
	)

	return nodeQueryCmd
}

// GetCmdQueryValidator implements the validator query command.
func GetCmdQueryValidator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validator [id]",
		Short: "Query a validator",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Validator(context.Background(), &types.QueryValidatorRequest{Id: args[0]})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res.Validator)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryValidators implements the query all validators command.
func GetCmdQueryValidators() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validators",
		Short: "Query for all validators",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := queryClient.Validators(
				context.Background(),
				&types.QueryValidatorsRequest{Pagination: pageReq},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "validators")
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryNode implements the query node command.
func GetCmdQueryNode() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "node [id]",
		Short:   "Query a node in the node whitelist",
		Long:    "Query a node in the node whitelist by id",
		Example: fmt.Sprintf("$ %s query node node <id>", version.AppName),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			id := args[0]
			if err := types.ValidateNodeID(id); err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Node(
				context.Background(),
				&types.QueryNodeRequest{Id: id},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res.Node)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryNodes implements the query nodes command.
func GetCmdQueryNodes() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "nodes",
		Short:   "Query all nodes in the node whitelist",
		Long:    "Query all nodes in the node whitelist",
		Example: fmt.Sprintf("$ %s query node nodes", version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := queryClient.Nodes(
				context.Background(),
				&types.QueryNodesRequest{Pagination: pageReq},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "nodes")

	return cmd
}

// GetCmdQueryParams implements the params query command.
func GetCmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "params",
		Args:    cobra.NoArgs,
		Short:   "Query the current validator parameters information",
		Long:    "Query values set as validator parameters.",
		Example: fmt.Sprintf("$ %s query validator params", version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Params(context.Background(), &types.QueryParamsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(&res.Params)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
