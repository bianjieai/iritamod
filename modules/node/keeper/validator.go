package keeper

import (
	"encoding/hex"
	"fmt"

	"cosmossdk.io/math"

	"github.com/cosmos/cosmos-sdk/codec/legacy"
	"github.com/cosmos/cosmos-sdk/types/bech32"

	gogotypes "github.com/gogo/protobuf/types"

	ctmbytes "github.com/cometbft/cometbft/libs/bytes"

	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/bianjieai/iritamod/modules/node/types"
	cautil "github.com/bianjieai/iritamod/utils/ca"
)

// SetHooks sets the validator hooks
func (k *Keeper) SetHooks(sh staking.StakingHooks) *Keeper {
	if k.hooks != nil {
		panic("cannot set validator hooks twice")
	}

	k.hooks = sh

	return k
}

// CreateValidator create a new validator
func (k Keeper) CreateValidator(ctx sdk.Context,
	id ctmbytes.HexBytes,
	name string,
	certificate string,
	pubKey cryptotypes.PubKey,
	power int64,
	description string,
	operator string,
) error {
	if k.HasValidatorName(ctx, name) {
		return types.ErrValidatorNameExists
	}

	if len(certificate) > 0 {
		cert, err := k.VerifyCert(ctx, certificate)
		if err != nil {
			return err
		}

		pk, err := cautil.GetPubkeyFromCert(cert)
		if err != nil {
			return sdkerrors.Wrap(types.ErrInvalidCert, err.Error())
		}

		pubKey, err = cryptocodec.FromTmPubKeyInterface(pk)
		if err != nil {
			return err
		}
	}

	if pubKey == nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "certificate")
	}

	if _, found := k.GetValidatorByConsAddr(ctx, sdk.GetConsAddress(pubKey)); found {
		return types.ErrValidatorPubkeyExists
	}

	operatorAddr, err := sdk.AccAddressFromBech32(operator)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, err.Error())
	}
	validator := types.NewValidator(
		id,
		name,
		description,
		pubKey,
		certificate,
		power,
		operatorAddr,
	)

	k.SetValidator(ctx, validator)
	consAddr, err := validator.GetConsAddr()
	if err != nil {
		return err
	}
	k.SetValidatorConsAddrIndex(ctx, id, consAddr)
	k.EnqueueValidatorsUpdate(ctx, validator, power)

	k.hooks.AfterValidatorCreated(ctx, validator.GetOperator())
	k.hooks.AfterValidatorBonded(ctx, consAddr, validator.GetOperator())
	return nil
}

// UpdateValidator updates an existing validator record
func (k Keeper) UpdateValidator(ctx sdk.Context,
	id ctmbytes.HexBytes,
	name string,
	certificate string,
	power int64,
	description string,
	operator string,
) error {
	if k.HasValidatorName(ctx, name) {
		return types.ErrValidatorNameExists
	}

	validator, found := k.GetValidator(ctx, id)
	if !found {
		return types.ErrUnknownValidator
	}

	if len(certificate) > 0 && certificate != validator.Certificate {
		cert, err := k.VerifyCert(ctx, certificate)
		if err != nil {
			return err
		}
		pk, err := cautil.GetPubkeyFromCert(cert)
		if err != nil {
			return sdkerrors.Wrap(types.ErrInvalidCert, err.Error())
		}
		pubkey, err := cryptocodec.FromTmPubKeyInterface(pk)
		if err != nil {
			return err
		}
		pkStr, err := bech32.ConvertAndEncode(sdk.GetConfig().GetBech32ConsensusPubPrefix(), legacy.Cdc.MustMarshal(pubkey))
		consAddr, err := validator.GetConsAddr()
		if err != nil {
			return err
		}
		// delete pubkey related index
		k.DeleteValidatorConsAddrIndex(ctx, consAddr)
		// delete from tendermint validator set
		k.EnqueueValidatorsUpdate(ctx, validator, 0)

		validator.Pubkey = pkStr
		validator.Certificate = certificate
		newConsAddr, err := validator.GetConsAddr()
		if err != nil {
			return err
		}
		k.SetValidatorConsAddrIndex(ctx, id, newConsAddr)
		k.EnqueueValidatorsUpdate(ctx, validator, validator.Power)
	}
	if power > 0 {
		validator.Power = power
		// override it if already exists
		k.EnqueueValidatorsUpdate(ctx, validator, validator.Power)
	}
	if len(description) > 0 && description != types.DoNotModifyDesc {
		validator.Description = description
	}
	if len(name) > 0 && name != types.DoNotModifyDesc {
		//delete old name
		store := ctx.KVStore(k.storeKey)
		store.Delete(types.GetValidatorNameKey(validator.Name))

		validator.Name = name
	}
	validator.Operator = operator
	k.SetValidator(ctx, validator)
	return nil
}

// RemoveValidator deletes an existing validator record
func (k Keeper) RemoveValidator(ctx sdk.Context,
	id ctmbytes.HexBytes,
	operator string,
) error {
	validator, found := k.GetValidator(ctx, id)
	if !found {
		return types.ErrUnknownValidator
	}

	k.DeleteValidator(ctx, validator)
	consAddr, err := validator.GetConsAddr()
	if err != nil {
		return err
	}
	k.DeleteValidatorConsAddrIndex(ctx, consAddr)
	// delete from tendermint validator set
	k.EnqueueValidatorsUpdate(ctx, validator, 0)

	k.hooks.AfterValidatorRemoved(ctx, consAddr, validator.GetOperator())
	return nil
}

// SetValidator sets the main record holding validator details
func (k Keeper) SetValidator(ctx sdk.Context, validator types.Validator) {
	store := ctx.KVStore(k.storeKey)
	id, _ := hex.DecodeString(validator.Id)
	// set validator by id
	bz := k.cdc.MustMarshal(&validator)
	store.Set(types.GetValidatorIDKey(id), bz)

	bz = k.cdc.MustMarshal(&gogotypes.BytesValue{Value: id})
	store.Set(types.GetValidatorNameKey(validator.Name), bz)
}

// GetValidator returns validator with id
func (k Keeper) GetValidator(ctx sdk.Context, id ctmbytes.HexBytes) (validator types.Validator, found bool) {
	store := ctx.KVStore(k.storeKey)

	value := store.Get(types.GetValidatorIDKey(id))
	if value == nil {
		return validator, false
	}

	k.cdc.MustUnmarshal(value, &validator)
	return validator, true
}

// HasValidatorName returns true or false with name
func (k Keeper) HasValidatorName(ctx sdk.Context, name string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetValidatorNameKey(name))
}

// DeleteValidator deletes the validator with id
func (k Keeper) DeleteValidator(ctx sdk.Context, validator types.Validator) {
	id, _ := hex.DecodeString(validator.Id)
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetValidatorIDKey(id))
	store.Delete(types.GetValidatorNameKey(validator.Name))
}

// SetValidatorConsAddrIndex sets the validator index by pubkey
func (k Keeper) SetValidatorConsAddrIndex(ctx sdk.Context, id ctmbytes.HexBytes, addr sdk.ConsAddress) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.BytesValue{Value: id})
	store.Set(types.GetValidatorConsAddrKey(addr), bz)
}

// DeleteValidatorConsAddrIndex deletes the validator index with pubkey
func (k Keeper) DeleteValidatorConsAddrIndex(ctx sdk.Context, addr sdk.ConsAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetValidatorConsAddrKey(addr))
}

// GetValidatorByConsAddr returns validator with pubkey
func (k Keeper) GetValidatorByConsAddr(ctx sdk.Context, addr sdk.ConsAddress) (validator types.Validator, found bool) {
	store := ctx.KVStore(k.storeKey)

	value := store.Get(types.GetValidatorConsAddrKey(addr))
	if value == nil {
		return validator, false
	}

	var id gogotypes.BytesValue
	k.cdc.MustUnmarshal(value, &id)

	return k.GetValidator(ctx, id.Value)
}

// EnqueueValidatorsUpdate enqueue to the validators update queue
func (k Keeper) EnqueueValidatorsUpdate(ctx sdk.Context, validator types.Validator, power int64) {
	// do not update this validator if already jailed
	if validator.Jailed {
		return
	}
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.Int64Value{Value: power})
	store.Set(types.GetValidatorUpdateQueueKey(validator.Pubkey), bz)
}

// DequeueValidatorsUpdate dequeue from the validators update queue
func (k Keeper) DequeueValidatorsUpdate(ctx sdk.Context, pubkey string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetValidatorUpdateQueueKey(pubkey))
}

// IterateUpdateValidators iterates through the validators update queue
func (k Keeper) IterateUpdateValidators(ctx sdk.Context, fn func(index int64, pubkey string, power int64) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.ValidatorsUpdateQueueKey)
	defer iterator.Close()

	i := int64(0)

	for ; iterator.Valid(); iterator.Next() {
		var power gogotypes.Int64Value
		k.cdc.MustUnmarshal(iterator.Value(), &power)
		stop := fn(i, string(iterator.Key()[1:]), power.Value)

		if stop {
			break
		}
		i++
	}
}

// IterateValidators iterates through the validator set and perform the provided function
func (k Keeper) IterateValidators(ctx sdk.Context, fn func(index int64, validator staking.ValidatorI) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.ValidatorsKey)
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		var validator types.Validator
		k.cdc.MustUnmarshal(iterator.Value(), &validator)
		stop := fn(i, validator)

		if stop {
			break
		}
		i++
	}
}

// GetAllValidators gets the set of all validators with no limits, used during genesis dump
//
// NOTE： add this to impl x/staking keeper interface; never use it.
func (k Keeper) GetAllValidators(ctx sdk.Context) (validators []staking.Validator) {
	return nil
}

// GetAllValidatorsLegacy gets the set of all validators with no limits, used during genesis dump
//
// NOTE： renamed from GetAllValidators since cosmos-sdk v0.47
func (k Keeper) GetAllValidatorsLegacy(ctx sdk.Context) (validators []types.Validator) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.ValidatorsKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var validator types.Validator
		k.cdc.MustUnmarshal(iterator.Value(), &validator)
		validators = append(validators, validator)
	}

	return validators
}

// ValidatorByID return the validator imformation by id
func (k Keeper) ValidatorByID(ctx sdk.Context, id ctmbytes.HexBytes) staking.ValidatorI {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetValidatorIDKey(id))

	var validator types.Validator
	k.cdc.MustUnmarshal(bz, &validator)
	return validator
}

// Validator return the validator imformation by valAddr
func (k Keeper) Validator(ctx sdk.Context, valAddr sdk.ValAddress) staking.ValidatorI {
	return k.ValidatorByConsAddr(ctx, sdk.ConsAddress(valAddr))
}

// ValidatorByConsAddr return the validator imformation by consAddr
func (k Keeper) ValidatorByConsAddr(ctx sdk.Context, consAddr sdk.ConsAddress) staking.ValidatorI {
	validator, found := k.GetValidatorByConsAddr(ctx, consAddr)
	if !found {
		return nil
	}
	return validator
}

// Slash not implement
func (k Keeper) Slash(ctx sdk.Context, consAddr sdk.ConsAddress, i int64, i2 int64, dec sdk.Dec) math.Int {
	return sdk.NewInt(0)
}

// SlashWithInfractionReason not implement
func (k *Keeper) SlashWithInfractionReason(ctx sdk.Context, consAddr sdk.ConsAddress, i int64, i2 int64, dec sdk.Dec, _ staking.Infraction) math.Int {
	return k.Slash(ctx, consAddr, i, i2, dec)
}

// Jail disable the validator
func (k Keeper) Jail(ctx sdk.Context, consAddr sdk.ConsAddress) {
	validator, found := k.GetValidatorByConsAddr(ctx, consAddr)
	if !found {
		panic(fmt.Errorf("validator with consensus-Address %s not found", consAddr))
	}
	if validator.Jailed {
		panic(fmt.Sprintf("cannot jail already jailed validator, validator: %v\n", validator))
	}

	k.EnqueueValidatorsUpdate(ctx, validator, 0)
	validator.Jailed = true
	k.SetValidator(ctx, validator)
}

// Unjail enable the validator
func (k Keeper) Unjail(ctx sdk.Context, consAddr sdk.ConsAddress) {
	validator, found := k.GetValidatorByConsAddr(ctx, consAddr)
	if !found {
		panic(fmt.Errorf("validator with consensus-Address %s not found", consAddr))
	}
	if !validator.Jailed {
		panic(fmt.Sprintf("cannot unjail already unjailed validator, validator: %v\n", validator))
	}
	validator.Jailed = false
	k.SetValidator(ctx, validator)
	k.EnqueueValidatorsUpdate(ctx, validator, validator.Power)
}

func (k Keeper) Delegation(context sdk.Context, accAddr sdk.AccAddress, consAddr sdk.ValAddress) staking.DelegationI {
	return staking.Delegation{}
}

func (k Keeper) MaxValidators(context sdk.Context) uint32 {
	return 0
}

// get the group of the bonded validators
func (k Keeper) GetLastValidators(ctx sdk.Context) (validators []types.Validator) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.ValidatorsKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var validator types.Validator
		k.cdc.MustUnmarshal(iterator.Value(), &validator)
		if !validator.Jailed {
			validators = append(validators, validator)
		}
	}

	return validators
}

func (k *Keeper) IterateBondedValidatorsByPower(ctx sdk.Context, fn func(index int64, validator staking.ValidatorI) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.ValidatorsKey)
	defer iterator.Close()

	i := int64(0)

	for ; iterator.Valid(); iterator.Next() {
		var validator types.Validator
		k.cdc.MustUnmarshal(iterator.Value(), &validator)
		if !validator.Jailed {
			stop := fn(i, validator)
			if stop {
				break
			}
			i++
		}
	}
}

func (k *Keeper) TotalBondedTokens(ctx sdk.Context) sdk.Int {
	total := sdk.NewInt(0)
	k.IterateValidators(ctx,
		func(index int64, validator staking.ValidatorI) bool {
			total = total.Add(validator.GetBondedTokens())
			return false
		},
	)
	return total
}

func (k *Keeper) IterateDelegations(
	ctx sdk.Context, delegator sdk.AccAddress,
	fn func(index int64, delegation staking.DelegationI) (stop bool),
) {
}

func (k *Keeper) VerifyCert(ctx sdk.Context, certStr string) (cert cautil.Cert, err error) {
	rootCertStr, _ := k.GetRootCert(ctx)
	rootCert, err := cautil.ReadCertificateFromMem([]byte(rootCertStr))
	if err != nil {
		return cert, sdkerrors.Wrap(types.ErrInvalidRootCert, err.Error())
	}

	cert, err = cautil.ReadCertificateFromMem([]byte(certStr))
	if err != nil {
		return cert, sdkerrors.Wrap(types.ErrInvalidCert, err.Error())
	}

	if err = cert.VerifyCertFromRoot(rootCert); err != nil {
		return cert, sdkerrors.Wrapf(types.ErrInvalidCert, "cannot be verified by root certificate, err: %s", err.Error())
	}

	return cert, nil
}

// IsValidatorJailed returns if the validator is jailed.
func (k *Keeper) IsValidatorJailed(ctx sdk.Context, addr sdk.ConsAddress) bool {
	v, ok := k.GetValidatorByConsAddr(ctx, addr)
	if !ok {
		return false
	}

	return v.Jailed
}
