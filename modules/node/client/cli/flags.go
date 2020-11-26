package cli

import (
	flag "github.com/spf13/pflag"
)

const (
	FlagCert = "cert"
)

// common flagsets to add to various functions
var (
	FsAddNode = flag.NewFlagSet("", flag.ContinueOnError)
)

func init() {
	FsAddNode.String(FlagCert, "", "file path to the X.509 certificate of the node identity")
}
