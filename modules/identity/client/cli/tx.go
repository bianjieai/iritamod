package cli

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"

	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/bianjieai/iritamod/modules/identity/types"
)

// NewTxCmd returns the transaction commands for this module
func NewTxCmd() *cobra.Command {
	identityTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Identity transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	identityTxCmd.AddCommand(
		NewCreateIdentityCmd(),
		NewUpdateIdentityCmd(),
	)

	return identityTxCmd
}

// NewCreateIdentityCmd implements creating an identity command
func NewCreateIdentityCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "CreateValidator an identity",
		Long:  "CreateValidator a new identity based on the given params.",
		Example: fmt.Sprintf(
			"$ %s tx identity create "+
				"--id=<id> "+
				"--pubkey=<public-key> "+
				"--pubkey-algo=<pubkey-algorithm> "+
				"--cert-file=<certificate-file> "+
				"--credentials=<credentials-uri> "+
				"--data=<data> "+
				"--from=<key>",
			version.AppName,
		),
		PreRunE: preCheckCmd,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			owner := clientCtx.GetFromAddress()

			var id []byte

			idStr := viper.GetString(FlagID)
			if len(idStr) > 0 {
				id, err = hex.DecodeString(idStr)
				if err != nil {
					return err
				}
			} else {
				id = uuid.NewV4().Bytes()
			}

			var pubKeyInfo *types.PubKeyInfo

			pubKeyStr := viper.GetString(FlagPubKey)
			if len(pubKeyStr) > 0 {
				pubKeyInfo = new(types.PubKeyInfo)
				pubKeyInfo.PubKey = pubKeyStr
				pubKeyInfo.Algorithm = types.PubKeyAlgorithmFromString(viper.GetString(FlagPubKeyAlgo))
			}

			certFile := viper.GetString(FlagCertificateFile)
			var cert []byte

			if len(certFile) > 0 {
				cert, err = ioutil.ReadFile(certFile)
				if err != nil {
					return fmt.Errorf("failed to read the certificate file: %s", err.Error())
				}
			}

			credentials := viper.GetString(FlagCredentials)

			data := viper.GetString(FlagData)
			msg := types.NewMsgCreateIdentity(id, pubKeyInfo, string(cert), credentials, owner, data)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FsCreateIdentity)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewUpdateIdentityCmd implements updating an identity command
func NewUpdateIdentityCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update [id]",
		Short: "UpdateValidator an identity",
		Long:  "UpdateValidator an existing identity.",
		Example: fmt.Sprintf(
			"$ %s tx identity update <id> "+
				"--pubkey=<public-key> "+
				"--pubkey-algo=<pubkey-algorithm> "+
				"--cert-file=<certificate-file> "+
				"--credentials=<credentials-uri> "+
				"--data=<data> "+
				"--from mykey",
			version.AppName,
		),
		Args:    cobra.ExactArgs(1),
		PreRunE: preCheckCmd,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			owner := clientCtx.GetFromAddress()

			id, err := hex.DecodeString(args[0])
			if err != nil {
				return err
			}

			var pubKeyInfo *types.PubKeyInfo

			pubKeyStr := viper.GetString(FlagPubKey)
			if len(pubKeyStr) > 0 {
				pubKeyInfo = new(types.PubKeyInfo)
				pubKeyInfo.PubKey = pubKeyStr
				pubKeyInfo.Algorithm = types.PubKeyAlgorithmFromString(viper.GetString(FlagPubKeyAlgo))
			}

			certFile := viper.GetString(FlagCertificateFile)
			var cert []byte

			if len(certFile) > 0 {
				cert, err = ioutil.ReadFile(certFile)
				if err != nil {
					return fmt.Errorf("failed to read the certificate file: %s", err.Error())
				}
			}

			credentials := viper.GetString(FlagCredentials)
			data := viper.GetString(FlagData)
			msg := types.NewMsgUpdateIdentity(id, pubKeyInfo, string(cert), credentials, owner, data)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FsUpdateIdentity)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func preCheckCmd(cmd *cobra.Command, _ []string) error {
	flags := cmd.Flags()

	if flags.Changed(FlagPubKey) {
		if !flags.Changed(FlagPubKeyAlgo) {
			return fmt.Errorf("public key algorithm must be provided when the public key set")
		}
	} else if flags.Changed(FlagPubKeyAlgo) {
		return fmt.Errorf("public key algorithm should not be provided when the public key not set")
	}

	return nil
}
