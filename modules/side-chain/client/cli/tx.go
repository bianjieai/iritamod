package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"

	"github.com/bianjieai/iritamod/modules/side-chain/types"
)

func NewTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Side-chain transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetCmdSpaceCmd(),
		GetCmdSpaceTransfer(),
		GetCmdCreateBlockHeader(),
	)

	return cmd
}

func GetCmdSpaceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "space",
		Short:                      "Layer2 space transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetCmdSpaceCreate(),
		GetCmdSpaceTransfer(),
	)

	return cmd
}

func GetCmdSpaceCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "create",
		Long: "create a new layer2 space",
		Example: fmt.Sprintf(
			"$ %s tx layer2 space create "+
				"--name=<name> "+
				"--uri=<uri>",
			version.AppName),
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			spaceName, err := cmd.Flags().GetString(FlagName)
			if err != nil {
				return nil
			}

			spaceUri, err := cmd.Flags().GetString(FlagUri)
			if err != nil {
				return nil
			}

			msg := types.NewMsgCreateSpace(
				spaceName,
				spaceUri,
				clientCtx.GetFromAddress().String(),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FsSpaceCreate)
	_ = cmd.MarkFlagRequired(FlagName)
	_ = cmd.MarkFlagRequired(FlagUri)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCmdSpaceTransfer() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "transfer [space-id] [recipient]",
		Long: "transfer ownership of space from sender to recipient",
		Example: fmt.Sprintf(
			"$ %s tx layer2 space transfer [space-id] [recipient]",
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

			msg := types.NewMsgTransferSpace(
				spaceId,
				args[1],
				clientCtx.GetFromAddress().String(),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCmdCreateBlockHeader() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "create-blockheader [space-id] [height] [header]",
		Long: "create a layer2 block header record",
		Example: fmt.Sprintf(
			"$ %s tx layer2 create-blockheader [space-id] [height] [header]",
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

			height, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateBlockHeader(
				spaceId,
				height,
				args[2],
				clientCtx.GetFromAddress().String(),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
