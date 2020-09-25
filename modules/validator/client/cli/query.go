package cli

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"

	"gitlab.bianjie.ai/irita-pro/iritamod/modules/validator/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	validatorQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the validator module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	validatorQueryCmd.AddCommand(
		GetCmdQueryValidator(),
		GetCmdQueryValidators(),
		GetCmdQueryParams(),
	)

	return validatorQueryCmd
}

// GetCmdQueryValidator implements the validator query command.
func GetCmdQueryValidator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validator [id]",
		Short: "Query a validator",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := hex.DecodeString(args[0])
			if err != nil {
				return err
			}

			clientCtx := client.GetClientContextFromCmd(cmd)
			if clientCtx, err = client.ReadQueryCommandFlags(clientCtx, cmd.Flags()); err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Validator(context.Background(), &types.QueryValidatorRequest{Id: id})
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res.Validator)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryValidators implements the query all validators command.
func GetCmdQueryValidators() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Query for all validators",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
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

			return clientCtx.PrintOutput(res)
		},
	}
	flags.AddPaginationFlagsToCmd(cmd, "validators")
	flags.AddQueryFlagsToCmd(cmd)
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
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Params(context.Background(), &types.QueryParamsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(&res.Params)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
