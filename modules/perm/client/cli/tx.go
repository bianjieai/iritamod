package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/iritamod/modules/perm/types"
)

func NewTxCmd() *cobra.Command {
	permTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Perm transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	permTxCmd.AddCommand(
		NewAssignRolesCmd(),
		NewUnassignRolesCmd(),
		NewBlockAccountCmd(),
		NewUnblockAccountCmd(),
	)

	return permTxCmd
}

// NewAssignRolesCmd implements the assign roles command handler.
func NewAssignRolesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "assign-roles [address] [roles]",
		Long: strings.TrimSpace(fmt.Sprintf(
			"Assign roles to an account.\n\n Auth options: %s, %s, %s, %s, %s, %s, %s, %s\n",
			types.RolePermAdmin,
			types.RoleBlacklistAdmin,
			types.RoleNodeAdmin,
			types.RoleParamAdmin,
			types.RoleIDAdmin,
			types.RoleBaseM1Admin,
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

			msg := types.NewMsgAssignRoles(
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

// NewUnassignRolesCmd implements the unassign roles command handler.
func NewUnassignRolesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "unassign-roles [address] [roles]",
		Long: strings.TrimSpace(fmt.Sprintf(
			"Unassign roles from an account.\n\nAuth options: %s, %s, %s, %s, %s, %s, %s, %s",
			types.RolePermAdmin,
			types.RoleBlacklistAdmin,
			types.RoleNodeAdmin,
			types.RoleParamAdmin,
			types.RoleIDAdmin,
			types.RoleBaseM1Admin,
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

			msg := types.NewMsgUnassignRoles(
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
