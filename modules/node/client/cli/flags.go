package cli

import (
	flag "github.com/spf13/pflag"

	"github.com/bianjieai/iritamod/modules/node/types"
)

const (
	FlagName        = "name"
	FlagCert        = "cert"
	FlagPower       = "power"
	FlagDescription = "description"

	FlagNodeID = "node-id"
	FlagIP     = "ip"
)

// common flagsets to add to various functions
var (
	FsCreateValidator = flag.NewFlagSet("", flag.ContinueOnError)
	FsUpdateValidator = flag.NewFlagSet("", flag.ContinueOnError)
	FsGrantNode       = flag.NewFlagSet("", flag.ContinueOnError)
)

func init() {
	FsCreateValidator.String(FlagName, "", "The name of the validator")
	FsCreateValidator.String(FlagCert, "", "The certificate file path of the validator identity")
	FsCreateValidator.Int64(FlagPower, 0, "The power of the validator")
	FsCreateValidator.String(FlagDescription, "", "The validator's (optional) details")
	FsCreateValidator.String(FlagNodeID, "", "The node's ID")

	FsUpdateValidator.String(FlagCert, "", "The certificate file path of the validator identity")
	FsUpdateValidator.Int64(FlagPower, 0, "The power of the validator")
	FsUpdateValidator.String(FlagDescription, types.DoNotModifyDesc, "The validator's (optional) details")
	FsUpdateValidator.String(FlagName, types.DoNotModifyDesc, "The alias name of the validator")

	FsGrantNode.String(FlagName, "", "The alias name of the node")
	FsGrantNode.String(FlagCert, "", "The certificate file path of the node identity")
}
