package cli

import (
	"context"
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"

	"github.com/bianjieai/iritamod/modules/side-chain/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                types.ModuleName,
		Short:              "Querying commands for the side-chain module",
		DisableFlagParsing: true,
	}

	cmd.AddCommand(
		GetQuerySpaceCmd(),
		GetCmdQueryBlockHeader(),
	)

	return cmd
}

func GetQuerySpaceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                "space",
		Short:              "Side chain space query subcommands",
		DisableFlagParsing: true,
	}

	cmd.AddCommand(
		GetCmdQuerySpaceInfo(),
		GetCmdQuerySpacesOfOwner(),
	)

	return cmd
}

func GetCmdQuerySpaceInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "space [space-id]",
		Long:    "query the space info of the given space-id",
		Example: fmt.Sprintf("$ %s q layer2 space space [space-id]", version.AppName),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			spaceId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			resp, err := queryClient.Space(
				context.Background(),
				&types.QuerySpaceRequest{
					SpaceId: spaceId,
				})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(resp)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func GetCmdQuerySpacesOfOwner() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "owner [owner]",
		Long:    "query the spaces of the given owner",
		Example: fmt.Sprintf("$ %s q layer2 space owner [owner]", version.AppName),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			if _, err := sdk.AccAddressFromBech32(args[0]); err != nil {
				return err
			}
			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			resp, err := queryClient.SpaceOfOwner(
				context.Background(),
				&types.QuerySpaceOfOwnerRequest{
					Owner:      args[0],
					Pagination: pageReq,
				})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(resp)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "spaces")

	return cmd
}

func GetCmdQueryBlockHeader() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "blockheader [space-id] [height]",
		Long:    "query the side chain block header",
		Example: fmt.Sprintf("$ %s q layer2 blockheader [space-id] [height]", version.AppName),
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			spaceId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			height, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			resp, err := queryClient.BlockHeader(
				context.Background(),
				&types.QueryBlockHeaderRequest{
					SpaceId: spaceId,
					Height:  height,
				})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(resp)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
