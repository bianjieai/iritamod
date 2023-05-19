package cli

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"

	"github.com/bianjieai/iritamod/modules/layer2/types"
)

func GetCmdNftCreateTokens() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "create-tokens [space-id] [class-id]",
		Long: "create token mappings for nft asset",
		Example: fmt.Sprintf(
			"$ %s tx layer2 nft create-tokens [space-id] [class-id]" +
				"--nfts=<nfts.json>" +
				version.AppName),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			spaceId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			nftsStr, err := cmd.Flags().GetString(FlagNftTokens)
			if err != nil {
				return err
			}

			var nfts []*types.TokenForNFT
			err = json.Unmarshal([]byte(nftsStr), &nfts)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateNFTs(
				spaceId,
				args[1],
				nfts,
				clientCtx.GetFromAddress().String(),
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FsNftCreateTokens)
	_ = cmd.MarkFlagRequired(FlagNftTokens)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCmdNftUpdateTokens() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "update-tokens [space-id] [class-id]",
		Long: "update token mappings for nft asset",
		Example: fmt.Sprintf(
			"$ %s tx layer2 nft update-tokens [space-id] [class-id]" +
				"--nfts=<nfts.json>" +
				version.AppName),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			spaceId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			nftsStr, err := cmd.Flags().GetString(FlagNftTokens)
			if err != nil {
				return err
			}

			var nfts []*types.TokenForNFT
			err = json.Unmarshal([]byte(nftsStr), &nfts)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateNFTs(
				spaceId,
				args[1],
				nfts,
				clientCtx.GetFromAddress().String(),
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FsNftUpdateTokens)
	_ = cmd.MarkFlagRequired(FlagNftTokens)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCmdNftDeleteTokens() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "delete-tokens [space-id] [class-id]",
		Long: "delete token mappings for nft asset",
		Example: fmt.Sprintf(
			"$ %s tx layer2 nft delete-tokens [space-id] [class-id]" +
				"--nft-ids=<nft-ids.json>" +
				version.AppName),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			spaceId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			nftsStr, err := cmd.Flags().GetString(FlagNftTokenIds)
			if err != nil {
				return err
			}

			var nftIds []string
			err = json.Unmarshal([]byte(nftsStr), &nftIds)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteNFTs(
				spaceId,
				args[1],
				nftIds,
				clientCtx.GetFromAddress().String(),
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FsNftDeleteTokens)
	_ = cmd.MarkFlagRequired(FlagNftTokenIds)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCmdNftUpdateClasses() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "update-classes",
		Long: "update class mappings for nft asset",
		Example: fmt.Sprintf(
			"$ %s tx layer2 nft update-classes" +
				"--class-infos=<class-infos.json>" +
				version.AppName),
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			classInfosStr, err := cmd.Flags().GetString(FlagNftClassInfos)

			var classUpdates []*types.UpdateClassForNFT
			err = json.Unmarshal([]byte(classInfosStr), &classUpdates)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateClassesForNFT(
				classUpdates,
				clientCtx.GetFromAddress().String(),
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FsNftUpdateClasses)
	_ = cmd.MarkFlagRequired(FlagNftClassInfos)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCmdNftDepositToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "deposit-token [space-id] [class-id] [token-id]",
		Long: "deposit an nft from layer1 to layer2",
		Example: fmt.Sprintf(
			"$ %s tx layer2 nft deposit-token [space-id] [class-id] [token-id]" +
				version.AppName),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			spaceId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgDepositTokenForNFT(
				spaceId,
				args[1],
				args[2],
				clientCtx.GetFromAddress().String(),
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

func GetCmdNftWithdrawToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "withdraw token [space-id] [class-id]",
		Long: "withdraw an nft from layer2 to layer1 and update its metadata",
		Example: fmt.Sprintf(
			"$ %s tx layer2 nft withdraw-token [space-id] [class-id] [token-id]" +
				"--owner=<owner>" +
				"--name=<name>" +
				"--uri=<uri>" +
				"--uri-hash=<uri-hash>" +
				"--data=<data>" +
				version.AppName),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			spaceId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			owner, err := cmd.Flags().GetString(FlagNftTokenOwner)
			if err != nil {
				return nil
			}
			name, err := cmd.Flags().GetString(FlagNftTokenName)
			if err != nil {
				return nil
			}
			uri, err := cmd.Flags().GetString(FlagNftTokenUri)
			if err != nil {
				return nil
			}
			uriHash, err := cmd.Flags().GetString(FlagNftTokenUriHash)
			if err != nil {
				return nil
			}
			data, err := cmd.Flags().GetString(FlagNftTokenData)
			if err != nil {
				return nil
			}

			msg := types.NewMsgWithdrawTokenForNFT(
				spaceId,
				args[1],
				args[2],
				owner,
				name,
				uri,
				uriHash,
				data,
				clientCtx.GetFromAddress().String(),
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FsNftWithdrawToken)
	_ = cmd.MarkFlagRequired(FlagNftTokenOwner)
	_ = cmd.MarkFlagRequired(FlagNftTokenName)
	_ = cmd.MarkFlagRequired(FlagNftTokenUri)
	_ = cmd.MarkFlagRequired(FlagNftTokenUriHash)
	_ = cmd.MarkFlagRequired(FlagNftTokenData)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCmdNftDepositClass() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "deposit-class [space-id] [class-id] [recipient]",
		Long: "deposit an nft class from layer1 to layer2",
		Example: fmt.Sprintf(
			"$ %s tx layer2 nft deposit-class [space-id] [class-id] [recipient]" +
				"--base-uri=<base-uri>" +
				version.AppName),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			spaceId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			baseUri, err := cmd.Flags().GetString(FlagNftClassBaseUri)
			if err != nil {
				return nil
			}

			msg := types.NewMsgDepositClassForNFT(
				spaceId,
				args[1],
				args[2],
				baseUri,
				clientCtx.GetFromAddress().String(),
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FsNftDepositClass)
	_ = cmd.MarkFlagRequired(FlagNftClassBaseUri)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCmdNftWithdrawClass() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "withdraw-class [class-id]",
		Long: "withdraw an nft class from layer2 to layer1",
		Example: fmt.Sprintf(
			"$ %s tx layer2 nft withdraw-class [class-id]" +
				"--owner=<owner>" +
				version.AppName),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			owner, err := cmd.Flags().GetString(FlagNftClassOwner)
			if err != nil {
				return nil
			}

			msg := types.NewMsgWithdrawClassForNFT(
				args[0],
				owner,
				clientCtx.GetFromAddress().String(),
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FsNftWithdrawClass)
	_ = cmd.MarkFlagRequired(FlagNftClassOwner)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
