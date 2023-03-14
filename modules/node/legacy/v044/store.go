package v044

import (
	"encoding/hex"
	"fmt"

	gogotypes "github.com/gogo/protobuf/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/tendermint/crypto/algo"

	"github.com/bianjieai/iritamod/modules/node/types"
	cautil "github.com/bianjieai/iritamod/utils/ca"
)

// TODO 生成的 ed25519 根证书放置此处
// NOTE：Place the root certificate and certificate type that you want to add here.
var (
	addNewRootCertValue = ``
	addNewRootCertType  = ""

	oldRootCertType = algo.SM2
)

// MigrateStore performs in-place store migrations from v1 to v2. The
// migration includes:
//
// - The `Certificate` type is added.
// - Change the root certificate type to array `Certificate` type.
// - Change the certificate type of the validator and node.
func MigrateStore(ctx sdk.Context, storeKey sdk.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)

	// overwrite the root certificate
	err := migrateRootCert(ctx, store, cdc)
	if err != nil {
		return err
	}

	// updated the certificate structure in kv database
	err = migrateCertStruct(ctx, store, cdc)
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
		Key:   addNewRootCertType,
		Value: addNewRootCertValue,
	}}
	newRootCert = append(newRootCert, types.Certificate{
		Key:   oldRootCertStr,
		Value: oldRootCertType,
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

func migrateCertStruct(ctx sdk.Context, store sdk.KVStore, cdc codec.BinaryCodec) error {
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
				Key:   oldRootCertType,
				Value: oldValidator.Certificate,
			},
		}

		bz := cdc.MustMarshal(&newValidator)
		id, _ := hex.DecodeString(oldValidator.Id)
		store.Set(types.GetValidatorIDKey(id), bz)
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
