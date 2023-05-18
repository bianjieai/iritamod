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

	"github.com/bianjieai/iritamod/modules/layer2/types"
)

func GetNftQueryNftClassCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "class [class-id]",
		Long:    "query the class mapping info",
		Example: fmt.Sprintf("$ %s q layer2 nft class [class-id]", version.AppName),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			resp, err := queryClient.ClassForNFT(
				context.Background(), &types.QueryClassForNFTRequest{
					ClassId: args[0],
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

func GetNftQueryNftClassesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "classes [class-id]",
		Long:    "query all of class mapping info",
		Example: fmt.Sprintf("$ %s q layer2 nft classes", version.AppName),
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)
			resp, err := queryClient.ClassesForNFT(
				context.Background(), &types.QueryClassesForNFTRequest{
					Pagination: pageReq,
				})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(resp)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "classes")

	return cmd
}

func GetNftQueryNftCollectionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "collection [space-id] [class-id]",
		Long:    "query the collection mapping info under a space",
		Example: fmt.Sprintf("$ %s q layer2 nft collection [space-id] [class-id]", version.AppName),
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			spaceId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)
			resp, err := queryClient.CollectionForNFT(
				context.Background(), &types.QueryCollectionForNFTRequest{
					SpaceId:    spaceId,
					ClassId:    args[1],
					Pagination: pageReq,
				})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(resp)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "collection")

	return cmd
}

func GetNftQueryNftTokenCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "token [space-id] [class-id] [token-id]",
		Long:    "query the nft token mapping",
		Example: fmt.Sprintf("$ %s q layer2 nft token [space-id] [class-id] [token-id]", version.AppName),
		Args:    cobra.ExactArgs(0),
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
			resp, err := queryClient.TokenForNFT(
				context.Background(), &types.QueryTokenForNFTRequest{
					SpaceId: spaceId,
					ClassId: args[1],
					NftId:   args[2],
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

func GetNftQueryNftOwnerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "owner [owner] [space-id]",
		Long: "query the nft token mappings of an owner. If space-id is 0, query nft across spaces owned by this owner",
		Example: fmt.Sprintf("$ %s q layer2 nft owner [owner] [space-id] "+
			"--class-id=<class-id>", version.AppName),
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			if _, err := sdk.AccAddressFromBech32(args[0]); err != nil {
				return err
			}

			spaceId, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}
			classId, err := cmd.Flags().GetString(FlagNftClassId)
			if err != nil {
				return err
			}
			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			resp, err := queryClient.TokensOfOwnerForNFT(
				context.Background(), &types.QueryTokensOfOwnerForNFTRequest{
					SpaceId:    spaceId,
					ClassId:    classId,
					Owner:      args[0],
					Pagination: pageReq,
				})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(resp)
		},
	}

	cmd.Flags().AddFlagSet(FsQueryNftOwner)
	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "nfts")

	return cmd
}

func GetNftQueryNftUriCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "uri [class-id]",
		Long: "query the base uri of a class or token uri of an nft. Must provide space-id and token-id to query token uri",
		Example: fmt.Sprintf("$ %s q layer2 nft uri [class-id] "+
			"--space-id=<space-id> "+
			"--token-id=<token-id> ", version.AppName),
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			if _, err := sdk.AccAddressFromBech32(args[0]); err != nil {
				return err
			}

			spaceIdStr, err := cmd.Flags().GetString(FlagSpaceId)
			if err != nil {
				return err
			}
			tokenId, err := cmd.Flags().GetString(FlagNftTokenId)
			if err != nil {
				return err
			}

			spaceId, err := strconv.ParseUint(spaceIdStr, 10, 64)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			// if spaceId is given
			if len(spaceIdStr) > 0 {
				resp, err := queryClient.TokenUriForNFT(
					context.Background(), &types.QueryTokenUriForNFTRequest{
						SpaceId: spaceId,
						ClassId: args[0],
						TokenId: tokenId,
					})
				if err != nil {
					return err
				}
				return clientCtx.PrintProto(resp)
			}

			resp, err := queryClient.BaseUriForNFT(
				context.Background(), &types.QueryBaseUriForNFTRequest{
					ClassId: args[0],
				})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(resp)
		},
	}

	cmd.Flags().AddFlagSet(FsQueryNftUri)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
