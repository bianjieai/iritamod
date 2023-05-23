package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"

	"github.com/bianjieai/iritamod/modules/layer2/types"
)

func GetCmdNftCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "nft",
		Short:                      "Layer2 NFT transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetCmdNftTokenCmd(),
		GetCmdNftClassCmd(),
	)
	return cmd
}

func GetCmdNftTokenCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "token",
		Short:                      "Layer2 NFT token transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetCmdNftTokenCreate(),
		GetCmdNftTokenUpdate(),
		GetCmdNftTokenDelete(),
		GetCmdNftTokenDeposit(),
		GetCmdNftTokenWithdraw(),
	)
	return cmd
}

func GetCmdNftClassCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "class",
		Short:                      "Layer2 NFT class transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetCmdNftClassUpdate(),
		GetCmdNftClassDeposit(),
		GetCmdNftClassWithdraw(),
	)
	return cmd
}

func GetCmdNftTokenCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "create [space-id] [class-id]",
		Long: "create token mappings for layer2 nft asset",
		Example: fmt.Sprintf(
			"$ %s tx layer2 nft token create [space-id] [class-id] "+
				"--ids=token1,token2,token3 "+
				"--owners=owner1,owner2,owner3",
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

			ids, err := cmd.Flags().GetString(FlagIds)
			if err != nil {
				return err
			}

			owners, err := cmd.Flags().GetString(FlagOwners)
			if err != nil {
				return err
			}

			ids = strings.TrimSpace(ids)
			idArray := strings.Split(ids, ",")

			owners = strings.TrimSpace(owners)
			ownerArray := strings.Split(owners, ",")

			if len(idArray) != len(ownerArray) {
				return fmt.Errorf("ids and owners length not match")
			}

			var tokens []types.TokenForNFT
			for i := 0; i < len(idArray); i++ {
				tokens = append(tokens, types.TokenForNFT{
					Id:    idArray[i],
					Owner: ownerArray[i],
				})
			}

			msg := types.NewMsgCreateNFTs(
				spaceId,
				args[1],
				tokens,
				clientCtx.GetFromAddress().String(),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FsNftTokenCreate)
	_ = cmd.MarkFlagRequired(FlagIds)
	_ = cmd.MarkFlagRequired(FlagOwners)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCmdNftTokenUpdate() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "update [space-id] [class-id]",
		Long: "update token mappings for layer2 nft asset",
		Example: fmt.Sprintf(
			"$ %s tx layer2 nft token update [space-id] [class-id] "+
				"--ids=token1,token2,token3 "+
				"--owners=owner1,owner2,owner3",
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

			ids, err := cmd.Flags().GetString(FlagIds)
			if err != nil {
				return err
			}

			owners, err := cmd.Flags().GetString(FlagOwners)
			if err != nil {
				return err
			}

			ids = strings.TrimSpace(ids)
			idArray := strings.Split(ids, ",")

			owners = strings.TrimSpace(owners)
			ownerArray := strings.Split(owners, ",")

			if len(idArray) != len(ownerArray) {
				return fmt.Errorf("ids and owners length not match")
			}

			var tokens []types.TokenForNFT
			for i := 0; i < len(idArray); i++ {
				tokens = append(tokens, types.TokenForNFT{
					Id:    idArray[i],
					Owner: ownerArray[i],
				})
			}

			msg := types.NewMsgUpdateNFTs(
				spaceId,
				args[1],
				tokens,
				clientCtx.GetFromAddress().String(),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FsNftTokenUpdate)
	_ = cmd.MarkFlagRequired(FlagIds)
	_ = cmd.MarkFlagRequired(FlagOwners)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCmdNftTokenDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "delete [space-id] [class-id]",
		Long: "delete token mappings for layer2 nft asset",
		Example: fmt.Sprintf(
			"$ %s tx layer2 nft token delete [space-id] [class-id] "+
				"--ids=token1,token2,token3",
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

			ids, err := cmd.Flags().GetString(FlagIds)
			if err != nil {
				return err
			}

			ids = strings.TrimSpace(ids)
			idArray := strings.Split(ids, ",")

			msg := types.NewMsgDeleteNFTs(
				spaceId,
				args[1],
				idArray,
				clientCtx.GetFromAddress().String(),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FsNftTokenDelete)
	_ = cmd.MarkFlagRequired(FlagIds)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCmdNftTokenDeposit() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "deposit [space-id] [class-id] [token-id]",
		Long: "deposit an nft from layer1 to layer2",
		Example: fmt.Sprintf(
			"$ %s tx layer2 nft token deposit [space-id] [class-id] [token-id]",
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

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCmdNftTokenWithdraw() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "withdraw [space-id] [class-id] [token-id] ",
		Long: "withdraw an nft from layer2 to layer1 and update its metadata",
		Example: fmt.Sprintf(
			"$ %s tx layer2 nft token withdraw [space-id] [class-id] [token-id] "+
				"--owner=<owner> "+
				"--name=<name> "+
				"--uri=<uri> "+
				"--uri-hash=<uri-hash> "+
				"--data=<data>",
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

			owner, err := cmd.Flags().GetString(FlagOwner)
			if err != nil {
				return nil
			}

			name, err := cmd.Flags().GetString(FlagName)
			if err != nil {
				return nil
			}

			uri, err := cmd.Flags().GetString(FlagUri)
			if err != nil {
				return nil
			}

			uriHash, err := cmd.Flags().GetString(FlagUriHash)
			if err != nil {
				return nil
			}

			data, err := cmd.Flags().GetString(FlagData)
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

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FsNftTokenWithdraw)
	_ = cmd.MarkFlagRequired(FlagOwner)
	_ = cmd.MarkFlagRequired(FlagName)
	_ = cmd.MarkFlagRequired(FlagUri)
	_ = cmd.MarkFlagRequired(FlagUriHash)
	_ = cmd.MarkFlagRequired(FlagData)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCmdNftClassUpdate() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "update [space-id]",
		Long: "update class mappings for nft asset",
		Example: fmt.Sprintf(
			"$ %s tx layer2 nft class update "+
				"--ids=token1,token2,token3 "+
				"--uris=uri1,uri2,uri3 "+
				"--owners=owner1,owner2,owner3",
			version.AppName),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			spaceId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			ids, err := cmd.Flags().GetString(FlagIds)
			if err != nil {
				return err
			}

			uris, err := cmd.Flags().GetString(FlagUris)
			if err != nil {
				return err
			}

			owners, err := cmd.Flags().GetString(FlagOwners)
			if err != nil {
				return err
			}

			ids = strings.TrimSpace(ids)
			idArray := strings.Split(ids, ",")

			uris = strings.TrimSpace(uris)
			uriArray := strings.Split(uris, ",")

			owners = strings.TrimSpace(owners)
			ownerArray := strings.Split(owners, ",")

			if len(idArray) != len(ownerArray) || len(idArray) != len(uriArray) {
				return fmt.Errorf("ids and owners length not match")
			}

			var classUpdates []types.UpdateClassForNFT
			for i := 0; i < len(idArray); i++ {
				classUpdates = append(classUpdates, types.UpdateClassForNFT{
					Id:    idArray[i],
					Uri:   uriArray[i],
					Owner: ownerArray[i],
				})
			}

			msg := types.NewMsgUpdateClassesForNFT(
				spaceId,
				classUpdates,
				clientCtx.GetFromAddress().String(),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FsNftClassUpdate)
	_ = cmd.MarkFlagRequired(FlagIds)
	_ = cmd.MarkFlagRequired(FlagUris)
	_ = cmd.MarkFlagRequired(FlagOwners)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCmdNftClassDeposit() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "deposit [space-id] [class-id] [recipient]",
		Long: "deposit an nft class from layer1 to layer2",
		Example: fmt.Sprintf(
			"$ %s tx layer2 nft class deposit [space-id] [class-id] [recipient] "+
				"--uri=<uri>",
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

			baseUri, err := cmd.Flags().GetString(FlagUri)
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

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FsNftClassDeposit)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCmdNftClassWithdraw() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "withdraw [space-id] [class-id]",
		Long: "withdraw an nft class from layer2 to layer1",
		Example: fmt.Sprintf(
			"$ %s tx layer2 nft withdraw-class [class-id]"+
				"--owner=<owner>",
			version.AppName),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			spaceId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			owner, err := cmd.Flags().GetString(FlagOwner)
			if err != nil {
				return nil
			}

			msg := types.NewMsgWithdrawClassForNFT(
				spaceId,
				args[0],
				owner,
				clientCtx.GetFromAddress().String(),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FsNftClassWithdraw)
	_ = cmd.MarkFlagRequired(FlagOwner)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
