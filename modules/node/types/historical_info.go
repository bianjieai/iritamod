package types

import (
	"sort"

	ctmproto "github.com/cometbft/cometbft/proto/tendermint/types"
)

// NewHistoricalInfo will create a historical information struct from header and valset
// it will first sort valset before inclusion into historical info
func NewHistoricalInfo(header ctmproto.Header, valSet Validators) HistoricalInfo {
	sort.Sort(valSet)

	return HistoricalInfo{
		Header: header,
		Valset: valSet,
	}
}
