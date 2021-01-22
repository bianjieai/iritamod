package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/iritamod/modules/admin/types"
)

func NewTxCmd() *cobra.Command {
	adminTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Admin transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	adminTxCmd.AddCommand(
		NewAddRolesCmd(),
		NewRemoveRolesCmd(),
		NewBlockAccountCmd(),
		NewUnblockAccountCmd(),
	)

	return adminTxCmd
}

// NewAddRolesCmd implements the add roles command handler.
func NewAddRolesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "add-roles [address] [roles]",
		Long: strings.TrimSpace(fmt.Sprintf(
			"Add roles to an account.\n\n Auth options: %s, %s, %s, %s, %s, %s, %s, %s\n",
			types.RolePermAdmin,
			types.RoleBlacklistAdmin,
			types.RoleNodeAdmin,
			types.RoleParamAdmin,
			types.RoleIDAdmin,
			types.RoleMintAdmin,
			types.RolePowerUser,
			types.RoleRelayerUser,
		)),
		Args: cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			addr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			roles, err := types.GetRolesFromStr(args[1:]...)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddRoles(
				roles,
				addr,
				clientCtx.GetFromAddress(),
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}

// NewRemoveRolesCmd implements the remove roles command handler.
func NewRemoveRolesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "remove-roles [address] [roles]",
		Long: strings.TrimSpace(fmt.Sprintf(
			"Remove roles from an account.\n\nAuth options: %s, %s, %s, %s, %s, %s, %s, %s",
			types.RolePermAdmin,
			types.RoleBlacklistAdmin,
			types.RoleNodeAdmin,
			types.RoleParamAdmin,
			types.RoleIDAdmin,
			types.RoleMintAdmin,
			types.RolePowerUser,
			types.RoleRelayerUser,
		)),
		Args: cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			addr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			roles, err := types.GetRolesFromStr(args[1:]...)
			if err != nil {
				return err
			}

			msg := types.NewMsgRemoveRoles(
				roles,
				addr,
				clientCtx.GetFromAddress(),
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}

// NewBlockAccountCmd implements the block account command handler.
func NewBlockAccountCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "block-account [address]",
		Short: "Block an account",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			addr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgBlockAccount(
				addr,
				clientCtx.GetFromAddress(),
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}

// NewUnblockAccountCmd implements the unblock account command handler.
func NewUnblockAccountCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unblock-account [address]",
		Short: "Unblock an account",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			addr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgUnblockAccount(
				addr,
				clientCtx.GetFromAddress(),
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}
