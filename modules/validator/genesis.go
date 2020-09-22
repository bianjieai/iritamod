package validator

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"
	tmtypes "github.com/tendermint/tendermint/types"

	cautil "gitlab.bianjie.ai/irita-pro/iritamod/utils/ca"

	"gitlab.bianjie.ai/irita-pro/iritamod/modules/validator/types"
)

// InitGenesis - store genesis validator set
func InitGenesis(ctx sdk.Context, k Keeper, data GenesisState) (res []abci.ValidatorUpdate) {
	if err := ValidateGenesis(data); err != nil {
		panic(err.Error())
	}

	k.SetParams(ctx, data.Params)
	k.SetRootCert(ctx, data.RootCert)

	for _, val := range data.Validators {
		k.SetValidator(ctx, val)
		pk, _ := sdk.GetPubKeyFromBech32(sdk.Bech32PubKeyTypeConsPub, val.Pubkey)

		k.SetValidatorConsAddrIndex(ctx, val.Id, sdk.GetConsAddress(pk))

		res = append(res, ABCIValidatorUpdate(
			sdk.MustGetPubKeyFromBech32(sdk.Bech32PubKeyTypeConsPub, val.Pubkey),
			val.Power,
		))
	}
	return
}

// ExportGenesis - output genesis valiadtor set
func ExportGenesis(ctx sdk.Context, k Keeper) *GenesisState {
	rootCert, _ := k.GetRootCert(ctx)
	return NewGenesisState(rootCert, k.GetParams(ctx), k.GetAllValidators(ctx))
}

// WriteValidators returns a slice of bonded genesis validators.
func WriteValidators(ctx sdk.Context, keeper Keeper) (vals []tmtypes.GenesisValidator) {
	for _, v := range keeper.GetLastValidators(ctx) {
		vals = append(vals, tmtypes.GenesisValidator{
			PubKey: v.GetConsPubKey(),
			Power:  v.GetConsensusPower(),
			Name:   v.GetMoniker(),
		})
	}

	return
}

// ValidateGenesis validates the provided validator genesis state to ensure the
// expected invariants holds. (i.e. no duplicate validators)
func ValidateGenesis(data GenesisState) error {
	if len(data.RootCert) == 0 {
		return errors.New("root certificate is not set in genesis state")
	}
	rootCert, err := cautil.ReadCertificateFromMem([]byte(data.RootCert))
	if err != nil {
		return fmt.Errorf("invalid root certificate in genesis state, %s", err.Error())
	}
	return validateGenesisStateValidators(rootCert, data.Validators)
}

func validateGenesisStateValidators(rootCert cautil.Cert, validators []Validator) error {
	nameMap := make(map[string]bool, len(validators))
	pubkeyMap := make(map[string]bool, len(validators))
	idMap := make(map[string]bool, len(validators))

	for i := 0; i < len(validators); i++ {
		val := validators[i]
		cert, err := cautil.ReadCertificateFromMem([]byte(val.Certificate))
		if err != nil {
			return sdkerrors.Wrap(types.ErrInvalidCert, err.Error())
		}

		err = cautil.VerifyCertFromRoot(cert, rootCert)
		if err != nil {
			return sdkerrors.Wrapf(types.ErrInvalidCert, "cannot be verified by root certificate, err: %s", err.Error())
		}

		pk := sdk.MustGetPubKeyFromBech32(sdk.Bech32PubKeyTypeConsPub, val.Pubkey)
		strPubkey := string(pk.Bytes())

		if _, ok := nameMap[val.Id.String()]; ok {
			return fmt.Errorf("duplicate validator id in genesis state: ID %v, pubkey %v", val.Id, val.Pubkey)
		}

		if _, ok := idMap[val.Name]; ok {
			return fmt.Errorf("duplicate validator name in genesis state: ID %v, pubkey %v", val.Id, val.Pubkey)
		}

		if _, ok := pubkeyMap[strPubkey]; ok {
			return fmt.Errorf("duplicate validator pubkey in genesis state: ID %v, pubkey %v", val.Id, val.Pubkey)
		}

		if val.Jailed {
			return fmt.Errorf("validator is jailed in genesis state: name %v, ID %v", val.Id, val.Pubkey)
		}

		pubkeyMap[strPubkey] = true
		nameMap[val.Name] = true
		idMap[val.Id.String()] = true
	}

	return nil
}
