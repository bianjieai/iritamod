package v044

import (
	"encoding/hex"
	"fmt"

	gogotypes "github.com/gogo/protobuf/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/iritamod/modules/node/types"
	cautil "github.com/bianjieai/iritamod/utils/ca"
)

// NOTE：If a node module needs to be migrated, assign this variable in the upgrade plan.
var (
	AddNewRootCertValue = ""
	AddNewRootCertType  = ""

	OldRootCertType = ""
)

// MigrateStore performs in-place store migrations from v1 to v2. The
// migration includes:
//
// - The `Certificate` type is added.
// - Change the root certificate type to array `Certificate` type.
// - Change the certificate type of the validator and node.
func MigrateStore(ctx sdk.Context, storeKey sdk.StoreKey, cdc codec.BinaryCodec, historicalEntries uint32) error {
	store := ctx.KVStore(storeKey)

	if AddNewRootCertValue == "" || AddNewRootCertType == "" || OldRootCertType == "" {
		return fmt.Errorf("missing certificate parameter")
	}

	// overwrite the root certificate
	err := migrateRootCert(ctx, store, cdc)
	if err != nil {
		return err
	}

	// updated the certificate structure in kv database, certificate of node is empty
	err = migrateCertStruct(ctx, store, cdc, historicalEntries)
	if err != nil {
		return err
	}

	return nil
}

func migrateRootCert(ctx sdk.Context, store sdk.KVStore, cdc codec.BinaryCodec) error {
	value := store.Get(types.RootCertKey)
	if value == nil {
		return fmt.Errorf("get old root certificate is empty")
	}

	var certStr gogotypes.StringValue
	cdc.MustUnmarshal(value, &certStr)
	oldRootCertStr := certStr.Value

	newRootCert := []types.Certificate{{
		Key:   AddNewRootCertType,
		Value: AddNewRootCertValue,
	}}
	newRootCert = append(newRootCert, types.Certificate{
		Key:   OldRootCertType,
		Value: oldRootCertStr,
	})

	// set root cert
	for _, cert := range newRootCert {
		var RootCertKey, bz []byte

		if _, err := cautil.IsSupportedAlgorithms(cert.Key); err != nil {
			panic(err)
		}
		RootCertKey = types.GetRootCertKey(cert.Key)

		bz = cdc.MustMarshal(&cert)
		store.Set(RootCertKey, bz)
	}

	return nil
}

func migrateCertStruct(ctx sdk.Context, store sdk.KVStore, cdc codec.BinaryCodec, historicalEntries uint32) error {
	// node
	nodeIterator := sdk.KVStorePrefixIterator(store, types.NodeKey)
	defer nodeIterator.Close()

	for ; nodeIterator.Valid(); nodeIterator.Next() {
		var oldNode OldNode
		cdc.MustUnmarshal(nodeIterator.Value(), &oldNode)

		newNode := types.Node{
			Id:          oldNode.Id,
			Name:        oldNode.Name,
			Certificate: &types.Certificate{},
		}

		bz := cdc.MustMarshal(&newNode)
		id, _ := hex.DecodeString(oldNode.Id)
		store.Set(types.GetNodeKey(id), bz)
	}

	// validator
	validatorIterator := sdk.KVStorePrefixIterator(store, types.ValidatorsKey)
	defer validatorIterator.Close()

	for ; validatorIterator.Valid(); validatorIterator.Next() {
		var oldValidator OldValidator
		cdc.MustUnmarshal(validatorIterator.Value(), &oldValidator)

		newValidator := types.Validator{
			Id:          oldValidator.Id,
			Name:        oldValidator.Name,
			Pubkey:      oldValidator.Pubkey,
			Power:       oldValidator.Power,
			Description: oldValidator.Description,
			Jailed:      oldValidator.Jailed,
			Operator:    oldValidator.Operator,
			Certificate: &types.Certificate{
				Key:   OldRootCertType,
				Value: oldValidator.Certificate,
			},
		}

		bz := cdc.MustMarshal(&newValidator)
		id, _ := hex.DecodeString(oldValidator.Id)
		store.Set(types.GetValidatorIDKey(id), bz)
	}

	// historical info
	for height := ctx.BlockHeight() - int64(historicalEntries); height <= ctx.BlockHeight(); height++ {
		key := types.GetHistoricalInfoKey(height)
		value := store.Get(key)
		if value == nil {
			continue
		}

		var oldHiInfo OldHistoricalInfo
		var newValidator []types.Validator
		cdc.MustUnmarshal(value, &oldHiInfo)

		for _, valset := range oldHiInfo.Valset {
			newValidator = append(newValidator, types.Validator{
				Id:          valset.Id,
				Name:        valset.Name,
				Pubkey:      valset.Pubkey,
				Power:       valset.Power,
				Description: valset.Description,
				Jailed:      valset.Jailed,
				Operator:    valset.Operator,
				Certificate: &types.Certificate{
					Key:   OldRootCertType,
					Value: valset.Certificate,
				},
			})
		}

		newHiInfo := types.HistoricalInfo{
			Header: oldHiInfo.Header,
			Valset: newValidator,
		}
		store.Set(key, cdc.MustMarshal(&newHiInfo))
	}

	return nil
}

func getOldCertValue(store sdk.KVStore, cdc codec.BinaryCodec, certKey []byte) (string, error) {
	value := store.Get(certKey)
	if value == nil {
		return "", fmt.Errorf("get old root certificate is empty")
	}

	var certStr gogotypes.StringValue
	cdc.MustUnmarshal(value, &certStr)
	return certStr.Value, nil
}
