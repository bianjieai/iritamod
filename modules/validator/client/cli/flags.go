package cli

import (
	flag "github.com/spf13/pflag"

	"gitlab.bianjie.ai/irita-pro/iritamod/modules/validator/types"
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
	FsCreate = flag.NewFlagSet("", flag.ContinueOnError)
	FsUpdate = flag.NewFlagSet("", flag.ContinueOnError)
)

func init() {
	FsCreate.String(FlagName, "", "The name of the validator")
	FsCreate.String(FlagCert, "", "The certificate file path of the validator identity")
	FsCreate.Int64(FlagPower, 0, "The power of the validator")
	FsCreate.String(FlagDescription, "", "The validator's (optional) details")
	FsCreate.String(FlagNodeID, "", "The node's ID")

	FsUpdate.String(FlagCert, "", "The certificate file path of the validator identity")
	FsUpdate.Int64(FlagPower, 0, "The power of the validator")
	FsUpdate.String(FlagDescription, types.DoNotModifyDesc, "The validator's (optional) details")
	FsUpdate.String(FlagName, types.DoNotModifyDesc, "The alias name of the validator")
}
