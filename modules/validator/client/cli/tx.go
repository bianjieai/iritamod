package cli

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	cfg "github.com/tendermint/tendermint/config"

	"gitlab.bianjie.ai/irita-pro/iritamod/modules/validator/types"
)

var (
	defaultPower = 100
)

func NewTxCmd() *cobra.Command {
	validatorTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Validator transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	validatorTxCmd.AddCommand(
		NewCreateValidatorCmd(),
		NewUpdateValidatorCmd(),
		NewRemoveValidatorCmd(),
	)

	return validatorTxCmd
}

// NewCreateValidatorCmd implements the create validator command handler.
func NewCreateValidatorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create new validator",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			txf := tx.NewFactoryCLI(clientCtx, cmd.Flags()).WithTxConfig(clientCtx.TxConfig).WithAccountRetriever(clientCtx.AccountRetriever)

			txf, msg, err := BuildCreateValidatorMsg(clientCtx, txf)
			if err != nil {
				return err
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FsCreate)
	FsCreate.String(FlagIP, "", fmt.Sprintf("The node's public IP. It takes effect only when used in combination with --%s", flags.FlagGenerateOnly))

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
		Use:   "update [id]",
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

	cmd.Flags().AddFlagSet(FsUpdate)
	flags.AddTxFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}

// NewRemoveValidatorCmd implements the remove validator command handler.
func NewRemoveValidatorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove [id]",
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

// Return the flagset, particular flags, and a description of defaults
// this is anticipated to be used with the gen-tx
func CreateValidatorMsgHelpers(ipDefault string) (fs *flag.FlagSet, pubkeyFlag, powerFlag, defaultsDesc string) {
	fsCreateValidator := flag.NewFlagSet("", flag.ContinueOnError)
	fsCreateValidator.String(FlagIP, ipDefault, "The node's public IP")

	fsCreateValidator.AddFlagSet(FsCreate)

	defaultsDesc = fmt.Sprintf(`
	power:         %d
`, defaultPower)

	return fsCreateValidator, FlagCert, FlagPower, defaultsDesc
}

// prepare flags in config
func PrepareFlagsForTxCreateValidator(
	config *cfg.Config, nodeID, chainID string, cert string,
) {
	ip := viper.GetString(FlagIP)
	if ip == "" {
		_, _ = fmt.Fprintf(os.Stderr, "couldn't retrieve an external IP; "+
			"the tx's memo field will be unset")
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
