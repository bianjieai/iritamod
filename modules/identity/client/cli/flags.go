// nolint
package cli

import (
	flag "github.com/spf13/pflag"

	"github.com/bianjieai/iritamod/modules/identity/types"
)

const (
	FlagID              = "id"
	FlagPubKey          = "pubkey"
	FlagPubKeyAlgo      = "pubkey-algo"
	FlagCertificateFile = "cert-file"
	FlagCredentials     = "credentials"
	FlagData            = "data"
)

// common flagsets to add to various functions
var (
	FsCreateIdentity = flag.NewFlagSet("", flag.ContinueOnError)
	FsUpdateIdentity = flag.NewFlagSet("", flag.ContinueOnError)
)

func init() {
	FsCreateIdentity.BytesHex(FlagID, nil, "ID of the identity with 32 hex characters, automatically generated if not provided")
	FsCreateIdentity.BytesHex(FlagPubKey, nil, "hex encoded public key")
	FsCreateIdentity.String(FlagPubKeyAlgo, "", "algorithm of the public key (rsa|dsa|ecdsa|ed25519|sm2)")
	FsCreateIdentity.String(FlagCertificateFile, "", "X.509 certificate file path")
	FsCreateIdentity.String(FlagCredentials, "", "uri pointing to credentials of the identity")
	FsCreateIdentity.String(FlagData, "", "custom data of the identity")

	FsUpdateIdentity.BytesHex(FlagPubKey, nil, "hex encoded public key to be added")
	FsUpdateIdentity.String(FlagPubKeyAlgo, "", "algorithm of the public key (rsa|dsa|ecdsa|ed25519|sm2)")
	FsUpdateIdentity.String(FlagCertificateFile, "", "file path of the X.509 certificate to be added")
	FsUpdateIdentity.String(FlagCredentials, types.DoNotModifyDesc, "uri pointing to credentials of the identity")
	FsUpdateIdentity.String(FlagData, types.DoNotModifyDesc, "custom data of the identity")
}
