package types

import (
	"sort"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec/legacy"
	"github.com/cosmos/cosmos-sdk/types/bech32"

	tmbytes "github.com/tendermint/tendermint/libs/bytes"
	tmprotocrypto "github.com/tendermint/tendermint/proto/tendermint/crypto"

	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

var _ stakingtypes.ValidatorI = Validator{}

// DoNotModifyDesc used in flags to indicate that description field should not be updated
const DoNotModifyDesc = "[do-not-modify]"

// NewValidator creates a new MsgCreateValidator instance.
func NewValidator(
	id tmbytes.HexBytes,
	name string,
	description string,
	pubKey cryptotypes.PubKey,
	cert string,
	power int64,
	operator sdk.AccAddress,
) Validator {
	var pkStr string
	var err error
	if pubKey != nil {
		pkStr, err = bech32.ConvertAndEncode(sdk.GetConfig().GetBech32ConsensusPubPrefix(), legacy.Cdc.MustMarshal(pubKey))
	}
	if err != nil {
		panic(err)
	}
	return Validator{
		Id:          id.String(),
		Name:        name,
		Pubkey:      pkStr,
		Certificate: cert,
		Power:       power,
		Description: description,
		Operator:    operator.String(),
	}
}

// IsJailed implement ValidatorI
func (v Validator) IsJailed() bool {
	return v.Jailed
}

// GetMoniker implement ValidatorI
func (v Validator) GetMoniker() string {
	return v.Name
}

// GetStatus implement ValidatorI
func (v Validator) GetStatus() stakingtypes.BondStatus {
	if v.Jailed {
		return stakingtypes.Unbonded
	} else {
		return stakingtypes.Bonded
	}
}

// IsBonded implement ValidatorI
func (v Validator) IsBonded() bool {
	return !v.Jailed
}

// IsUnbonded implement ValidatorI
func (v Validator) IsUnbonded() bool {
	return v.Jailed
}

// IsUnbonding implement ValidatorI
func (v Validator) IsUnbonding() bool {
	return false
}

// GetOperator implement ValidatorI
func (v Validator) GetOperator() sdk.ValAddress {
	pubkey, err := v.ConsPubKey()
	if err != nil {
		panic(err)
	}
	return sdk.ValAddress(pubkey.Address())
}

// ConsPubKey returns the validator PubKey as a cryptotypes.PubKey.
func (v Validator) ConsPubKey() (pk cryptotypes.PubKey, err error) {
	bz, err := sdk.GetFromBech32(v.Pubkey, sdk.GetConfig().GetBech32ConsensusPubPrefix())
	return legacy.PubKeyFromBytes(bz)
}

// TmConsPublicKey casts Validator.ConsensusPubkey to tmprotocrypto.PubKey.
func (v Validator) TmConsPublicKey() (tmprotocrypto.PublicKey, error) {
	pk, err := v.ConsPubKey()
	if err != nil {
		return tmprotocrypto.PublicKey{}, err
	}

	tmPk, err := cryptocodec.ToTmProtoPublicKey(pk)
	if err != nil {
		return tmprotocrypto.PublicKey{}, err
	}

	return tmPk, nil
}

// GetConsAddr implement ValidatorI
func (v Validator) GetConsAddr() (sdk.ConsAddress, error) {
	pk, err := v.ConsPubKey()
	if err != nil {
		return sdk.ConsAddress{}, err
	}

	return sdk.ConsAddress(pk.Address()), nil
}

// GetTokens implement ValidatorI
func (v Validator) GetTokens() sdk.Int {
	return sdk.TokensFromConsensusPower(v.Power, sdk.DefaultPowerReduction)
}

// GetBondedTokens implement ValidatorI
func (v Validator) GetBondedTokens() sdk.Int {
	if v.Jailed {
		return sdk.NewInt(0)
	}
	return sdk.TokensFromConsensusPower(v.Power, sdk.DefaultPowerReduction)
}

// GetConsensusPower implement ValidatorI
func (v Validator) GetConsensusPower(_ sdk.Int) int64 {
	return v.Power
}

// GetCommission implement ValidatorI
func (v Validator) GetCommission() sdk.Dec {
	return sdk.NewDec(0)
}

// GetMinSelfDelegation implement ValidatorI
func (v Validator) GetMinSelfDelegation() sdk.Int {
	return sdk.NewInt(0)
}

// GetDelegatorShares implement ValidatorI
func (v Validator) GetDelegatorShares() sdk.Dec {
	return sdk.NewDec(0)
}

// TokensFromShares implement ValidatorI
func (v Validator) TokensFromShares(dec sdk.Dec) sdk.Dec {
	return sdk.NewDec(0)
}

// TokensFromSharesTruncated implement ValidatorI
func (v Validator) TokensFromSharesTruncated(dec sdk.Dec) sdk.Dec {
	return sdk.NewDec(0)
}

// TokensFromSharesRoundUp implement ValidatorI
func (v Validator) TokensFromSharesRoundUp(dec sdk.Dec) sdk.Dec {
	return sdk.NewDec(0)
}

// SharesFromTokens implement ValidatorI
func (v Validator) SharesFromTokens(amt sdk.Int) (sdk.Dec, error) {
	return sdk.NewDec(0), nil
}

// SharesFromTokensTruncated implement ValidatorI
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
