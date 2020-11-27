package node

import (
	"encoding/hex"
	"errors"
	"fmt"

	abci "github.com/tendermint/tendermint/abci/types"
	tmtypes "github.com/tendermint/tendermint/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/bianjieai/iritamod/modules/node/types"
	cautil "github.com/bianjieai/iritamod/utils/ca"
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

		id, err := hex.DecodeString(val.Id)
		if err != nil {
			panic(err)
		}

		k.SetValidatorConsAddrIndex(ctx, id, sdk.GetConsAddress(pk))

		tmPubKey, err := val.TmConsPubKey()
		if err != nil {
			panic(err)
		}

		res = append(res, ABCIValidatorUpdate(
			tmPubKey,
			val.Power,
		))
	}

	for _, node := range data.Nodes {
		id, _ := hex.DecodeString(node.Id)
		k.SetNode(ctx, id, node)
	}

	return
}

// ExportGenesis - output genesis valiadtor set
func ExportGenesis(ctx sdk.Context, k Keeper) *GenesisState {
	rootCert, _ := k.GetRootCert(ctx)
	return NewGenesisState(rootCert, k.GetParams(ctx), k.GetAllValidators(ctx), k.GetNodes(ctx))
}

// WriteValidators returns a slice of bonded genesis validators.
func WriteValidators(ctx sdk.Context, keeper Keeper) (vals []tmtypes.GenesisValidator) {
	for _, v := range keeper.GetLastValidators(ctx) {
		consPk, err := v.TmConsPubKey()
		if err != nil {
			continue
		}
		vals = append(vals, tmtypes.GenesisValidator{
			PubKey: consPk,
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

	err = validateGenesisStateValidators(rootCert, data.Validators)
	if err != nil {
		return err
	}

	return validateNodes(data.Nodes)
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

		if err = cautil.VerifyCertFromRoot(cert, rootCert); err != nil {
			return sdkerrors.Wrapf(types.ErrInvalidCert, "cannot be verified by root certificate, err: %s", err.Error())
		}

		pk := sdk.MustGetPubKeyFromBech32(sdk.Bech32PubKeyTypeConsPub, val.Pubkey)
		strPubkey := string(pk.Bytes())

		if _, ok := nameMap[val.Id]; ok {
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
		idMap[val.Id] = true
	}

	return nil
}

// validateNodes validates the nodes in genesis state
func validateNodes(nodes []types.Node) error {
	for _, node := range nodes {
		if err := node.Validate(); err != nil {
			return err
		}
	}

	return nil
}
