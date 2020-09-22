package types

import (
	"sort"
	"strings"

	tmbytes "github.com/tendermint/tendermint/libs/bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingexported "github.com/cosmos/cosmos-sdk/x/staking/exported"
	"github.com/tendermint/tendermint/crypto"
)

var _ stakingexported.ValidatorI = Validator{}

// constant used in flags to indicate that description field should not be updated
const DoNotModifyDesc = "[do-not-modify]"

// NewValidator creates a new MsgCreateValidator instance.
func NewValidator(id tmbytes.HexBytes,
	name, description string,
	pubKey crypto.PubKey,
	cert string,
	power int64,
	operator sdk.AccAddress) Validator {
	var pkStr string
	if pubKey != nil {
		pkStr = sdk.MustBech32ifyPubKey(sdk.Bech32PubKeyTypeConsPub, pubKey)
	}

	return Validator{
		Id:          id,
		Name:        name,
		Pubkey:      pkStr,
		Certificate: cert,
		Power:       power,
		Description: description,
		Operator:    operator,
	}
}

func (v Validator) IsJailed() bool {
	return v.Jailed
}

func (v Validator) GetMoniker() string {
	return v.Name
}

func (v Validator) GetStatus() sdk.BondStatus {
	if v.Jailed {
		return sdk.Unbonded
	} else {
		return sdk.Bonded
	}
}

func (v Validator) IsBonded() bool {
	return !v.Jailed
}

func (v Validator) IsUnbonded() bool {
	return v.Jailed
}

func (v Validator) IsUnbonding() bool {
	return false
}

func (v Validator) GetOperator() sdk.ValAddress {
	return sdk.ValAddress(v.GetConsPubKey().Address())
}

func (v Validator) GetConsPubKey() crypto.PubKey {
	return sdk.MustGetPubKeyFromBech32(sdk.Bech32PubKeyTypeConsPub, v.Pubkey)
}

func (v Validator) GetConsAddr() sdk.ConsAddress {
	return sdk.ConsAddress(v.GetConsPubKey().Address())
}

func (v Validator) GetTokens() sdk.Int {
	return sdk.TokensFromConsensusPower(v.Power)
}

func (v Validator) GetBondedTokens() sdk.Int {
	if v.Jailed {
		return sdk.NewInt(0)
	}
	return sdk.TokensFromConsensusPower(v.Power)
}

func (v Validator) GetConsensusPower() int64 {
	return v.Power
}

func (v Validator) GetCommission() sdk.Dec {
	return sdk.NewDec(0)
}

func (v Validator) GetMinSelfDelegation() sdk.Int {
	return sdk.NewInt(0)
}

func (v Validator) GetDelegatorShares() sdk.Dec {
	return sdk.NewDec(0)
}

func (v Validator) TokensFromShares(dec sdk.Dec) sdk.Dec {
	return sdk.NewDec(0)
}

func (v Validator) TokensFromSharesTruncated(dec sdk.Dec) sdk.Dec {
	return sdk.NewDec(0)
}

func (v Validator) TokensFromSharesRoundUp(dec sdk.Dec) sdk.Dec {
	return sdk.NewDec(0)
}

func (v Validator) SharesFromTokens(amt sdk.Int) (sdk.Dec, error) {
	return sdk.NewDec(0), nil
}

func (v Validator) SharesFromTokensTruncated(amt sdk.Int) (sdk.Dec, error) {
	return sdk.NewDec(0), nil

}

// Validators is a collection of Validator
type Validators []Validator

// Sort Validators sorts validator array in ascending operator address order
func (v Validators) Sort() {
	sort.Sort(v)
}

// Implements sort interface
func (v Validators) Len() int {
	return len(v)
}

// Implements sort interface
func (v Validators) Less(i, j int) bool {
	return strings.Compare(v[i].Name, v[j].Name) == -1
}

// Implements sort interface
func (v Validators) Swap(i, j int) {
	it := v[i]
	v[i] = v[j]
	v[j] = it
}
