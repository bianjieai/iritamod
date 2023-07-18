package cli

import flag "github.com/spf13/pflag"

const (
	FlagName = "name"
	FlagUri  = "uri"
)

var (
	FsSpaceCreate = flag.NewFlagSet("", flag.ContinueOnError)
)

func init() {
	FsSpaceCreate.String(FlagName, "", "name of the space")
	FsSpaceCreate.String(FlagUri, "", "uri of the space")
}
