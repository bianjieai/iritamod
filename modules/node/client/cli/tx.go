package cli

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"

	cfg "github.com/tendermint/tendermint/config"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/bianjieai/iritamod/modules/node/types"
)

var (
	defaultPower = 100
)

func NewTxCmd() *cobra.Command {
	nodeTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Node transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	nodeTxCmd.AddCommand(
		NewCreateValidatorCmd(),
		NewUpdateValidatorCmd(),
		NewRemoveValidatorCmd(),
		NewAddNodeCmd(),
		NewRemoveNodeCmd(),
	)

	return nodeTxCmd
}

// NewCreateValidatorCmd implements the create validator command handler.
func NewCreateValidatorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-validator",
		Short: "Create a new validator",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			txf := tx.NewFactoryCLI(clientCtx, cmd.Flags()).WithTxConfig(clientCtx.TxConfig).WithAccountRetriever(clientCtx.AccountRetriever)

			_, msg, err := BuildCreateValidatorMsg(clientCtx, txf)
			if err != nil {
				return err
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	FsCreateValidator.String(FlagIP, "", fmt.Sprintf("The node's public IP. It takes effect only when used in combination with --%s", flags.FlagGenerateOnly))
	cmd.Flags().AddFlagSet(FsCreateValidator)

	flags.AddTxFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(flags.FlagFrom)
	_ = cmd.MarkFlagRequired(FlagName)
	_ = cmd.MarkFlagRequired(FlagCert)
	_ = cmd.MarkFlagRequired(FlagPower)

	return cmd
}

// NewUpdateValidatorCmd implements the update validator command handler.
func NewUpdateValidatorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-validator [id]",
		Args:  cobra.ExactArgs(1),
		Short: "Update an existing validator",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			certPath := viper.GetString(FlagCert)
			data, _ := ioutil.ReadFile(certPath)

			id, err := hex.DecodeString(args[0])
			if err != nil {
				return fmt.Errorf("invalid validator id: %s", args[0])
			}

			msg := types.NewMsgUpdateValidator(
				id,
				viper.GetString(FlagName),
				viper.GetString(FlagDescription),
				string(data),
				viper.GetInt64(FlagPower),
				clientCtx.GetFromAddress(),
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FsUpdateValidator)
	flags.AddTxFlagsToCmd(cmd)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}

// NewRemoveValidatorCmd implements the remove validator command handler.
func NewRemoveValidatorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-validator [id]",
		Short: "Remove an existing validator",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			id, err := hex.DecodeString(args[0])
			if err != nil {
				return fmt.Errorf("invalid validator id:%s", args[0])
			}

			msg := types.NewMsgRemoveValidator(
				id,
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

// NewAddNodeCmd implements adding a node command
func NewAddNodeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a node to the node whitelist",
		Long:  "Add a node to the node whitelist based on the identity certificate",
		Example: fmt.Sprintf(
			"$ %s tx node add --name=<name> --cert=<certificate-file> --from mykey",
			version.AppName,
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			operator := clientCtx.GetFromAddress()

			name := strings.TrimSpace(viper.GetString(FlagName))

			certFile := strings.TrimSpace(viper.GetString(FlagCert))
			if len(certFile) == 0 {
				return fmt.Errorf("the certificate file path must not be empty")
			}

			cert, err := ioutil.ReadFile(certFile)
			if err != nil {
				return fmt.Errorf("failed to read the certificate file: %s", err.Error())
			}

			msg := types.NewMsgAddNode(name, string(cert), operator)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FsAddNode)
	cmd.MarkFlagRequired(FlagCert)

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewRemoveNodeCmd implements removing a node command
func NewRemoveNodeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove [id]",
		Short: "Remove a node from the node whitelist",
		Long:  "Remove a node from the node whitelist by id",
		Example: fmt.Sprintf(
			"$ %s tx node remove <id> --from mykey",
			version.AppName,
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			operator := clientCtx.GetFromAddress()

			id, err := hex.DecodeString(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgRemoveNode(id, operator)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// CreateValidatorMsgHelpers Return the flagset, particular flags, and a description of defaults
// this is anticipated to be used with the gen-tx
func CreateValidatorMsgHelpers(ipDefault string) (fs *flag.FlagSet, pubkeyFlag, powerFlag, defaultsDesc string) {
	fsCreateValidator := flag.NewFlagSet("", flag.ContinueOnError)
	fsCreateValidator.String(FlagIP, ipDefault, "The node's public IP")

	fsCreateValidator.AddFlagSet(fsCreateValidator)

	defaultsDesc = fmt.Sprintf("\n	power:		%d\n", defaultPower)

	return fsCreateValidator, FlagCert, FlagPower, defaultsDesc
}

// PrepareFlagsForTxCreateValidator prepare flags in config
func PrepareFlagsForTxCreateValidator(config *cfg.Config, nodeID, chainID string, cert string) {
	ip := viper.GetString(FlagIP)
	if ip == "" {
		_, _ = fmt.Fprintf(os.Stderr, "couldn't retrieve an external IP; the tx's memo field will be unset")
	}

	if len(strings.TrimSpace(viper.GetString(FlagNodeID))) > 0 {
		nodeID = strings.TrimSpace(viper.GetString(FlagNodeID))
	}

	details := viper.GetString(FlagDescription)

	viper.Set(flags.FlagChainID, chainID)
	viper.Set(flags.FlagFrom, viper.GetString(flags.FlagFrom))
	viper.Set(FlagCert, cert)
	viper.Set(FlagDescription, details)
	viper.Set(FlagNodeID, nodeID)
	viper.Set(FlagIP, ip)

	if viper.GetInt64(FlagPower) == 0 {
		viper.Set(FlagPower, defaultPower)
	}
}

// BuildCreateValidatorMsg makes a new MsgCreateValidator.
func BuildCreateValidatorMsg(clientCtx client.Context, txBldr tx.Factory) (tx.Factory, sdk.Msg, error) {
	certPath := viper.GetString(FlagCert)

	data, err := ioutil.ReadFile(certPath)
	if err != nil {
		return txBldr, nil, err
	}

	msg := types.NewMsgCreateValidator(
		viper.GetString(FlagName),
		viper.GetString(FlagDescription),
		string(data),
		viper.GetInt64(FlagPower),
		clientCtx.GetFromAddress(),
	)

	if viper.GetBool(flags.FlagGenerateOnly) {
		ip := viper.GetString(FlagIP)
		nodeID := viper.GetString(FlagNodeID)

		if nodeID != "" && ip != "" {
			txBldr = txBldr.WithMemo(fmt.Sprintf("%s@%s:26656", nodeID, ip))
		}
	}
	return txBldr, msg, nil
}
