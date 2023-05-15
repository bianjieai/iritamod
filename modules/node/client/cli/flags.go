package cli

import (
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"

	"github.com/tendermint/tendermint/crypto/algo"

	"github.com/bianjieai/iritamod/modules/node/types"
)

const (
	FlagName        = "name"
	FlagCertType    = "cert-type"
	FlagCert        = "cert"
	FlagPower       = "power"
	FlagDescription = "description"

	FlagNodeID = "node-id"
	FlagIP     = "ip"
)

const (
	FlagNodeAlgo = "node-algo"
)

// common flagsets to add to various functions
var (
	FsCreateValidator = flag.NewFlagSet("", flag.ContinueOnError)
	FsUpdateValidator = flag.NewFlagSet("", flag.ContinueOnError)
	FsGrantNode       = flag.NewFlagSet("", flag.ContinueOnError)
)

func init() {
	FsCreateValidator.String(FlagName, "", "The name of the validator")
	FsCreateValidator.String(FlagCertType, "", "The certificate type of the validator identity (e.g. ed25519)")
	FsCreateValidator.String(FlagCert, "", "The certificate file path of the validator identity")
	FsCreateValidator.Int64(FlagPower, 0, "The power of the validator")
	FsCreateValidator.String(FlagDescription, "", "The validator's (optional) details")
	FsCreateValidator.String(FlagNodeID, "", "The node's ID")

	FsUpdateValidator.String(FlagName, types.DoNotModifyDesc, "The alias name of the validator")
	FsUpdateValidator.String(FlagCertType, "", "The certificate type of the validator identity (e.g. ed25519)")
	FsUpdateValidator.String(FlagCert, "", "The certificate file path of the validator identity")
	FsUpdateValidator.Int64(FlagPower, 0, "The power of the validator")
	FsUpdateValidator.String(FlagDescription, types.DoNotModifyDesc, "The validator's (optional) details")

	FsGrantNode.String(FlagName, "", "The alias name of the node")
	FsGrantNode.String(FlagCertType, "", "The certificate type of the node identity (e.g. ed25519)")
	FsGrantNode.String(FlagCert, "", "The certificate file path of the node identity")
}

func SetGlobeFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().String(FlagNodeAlgo, algo.Algo, "type of this node (sm2|ed25519)")
}
