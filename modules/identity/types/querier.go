package types

import (
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
)

const (
	QueryIdentity = "identity" // query identity
)

// QueryIdentityParams defines the params to query an identity
type QueryIdentityParams struct {
	ID tmbytes.HexBytes
}
