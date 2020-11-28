package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/libs/tempfile"

	"github.com/cosmos/cosmos-sdk/server"

	"github.com/bianjieai/iritamod/modules/genutil"
)

var (
	FlagType    = "type"
	FlagOutFile = "out-file"
)

// Genkey returns a command that generates the key from priv_validator_key.json or node_key.json
// Will used to generate CA request.
func GenKey(defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "genkey",
		Short:   "generate key from the key file",
		Args:    cobra.NoArgs,
		PreRunE: preCheckCmd,
		RunE: func(cmd *cobra.Command, args []string) error {
			serverCtx := server.GetServerContextFromCmd(cmd)
			config := serverCtx.Config

			nodeKey, filePv, err := genutil.InitializeNodeValidatorFiles(config)
			if err != nil {
				return err
			}

			var privKey crypto.PrivKey

			keyType := strings.TrimSpace(viper.GetString(FlagType))
			if keyType == "node" {
				privKey = nodeKey.PrivKey
			} else {
				privKey = filePv.Key.PrivKey
			}

			key, err := genutil.Genkey(privKey)
			if err != nil {
				return err
			}

			return tempfile.WriteFileAtomic(viper.GetString(FlagOutFile), key, 0600)
		},
	}

	cmd.Flags().String(FlagType, "validator", "key type (node|validator)")
	cmd.Flags().String(cli.HomeFlag, defaultNodeHome, "node's home directory")
	cmd.Flags().String(FlagOutFile, "priv.pem", "private key file path")

	return cmd
}

func preCheckCmd(cmd *cobra.Command, _ []string) error {
	flags := cmd.Flags()

	if flags.Changed(FlagType) {
		keyType := strings.TrimSpace(viper.GetString(FlagType))
		if keyType != "node" && keyType != "validator" {
			return fmt.Errorf("key type must be node or validator")
		}
	}

	return nil
}
