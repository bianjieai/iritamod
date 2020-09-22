package keeper

import (
	"bytes"
	"encoding/binary"
	"strings"

	"github.com/tendermint/tendermint/crypto/tmhash"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"gitlab.bianjie.ai/irita-pro/iritamod/modules/identity/types"
)

// CreateIdentity creates an identity
func (k Keeper) CreateIdentity(
	ctx sdk.Context,
	id tmbytes.HexBytes,
	pubKey *types.PubKeyInfo,
	certificate,
	credentials string,
	owner sdk.AccAddress,
) error {
	if k.HasIdentity(ctx, id) {
		return sdkerrors.Wrap(types.ErrIdentityExists, id.String())
	}

	if pubKey != nil {
		if err := k.AddPubKey(ctx, id, pubKey); err != nil {
			return err
		}
	}

	if len(certificate) > 0 {
		if err := k.AddCertificate(ctx, id, certificate); err != nil {
			return err
		}
	}

	if len(credentials) > 0 {
		k.SetCredentials(ctx, id, credentials)
	}

	k.SetOwner(ctx, id, owner)

	return nil
}

// UpdateIdentity updates the specified identity by adding the given public key
// and certificate or modifying the credentials
func (k Keeper) UpdateIdentity(
	ctx sdk.Context,
	id tmbytes.HexBytes,
	pubKey *types.PubKeyInfo,
	certificate,
	credentials string,
	owner sdk.AccAddress,
) error {
	identityOwner, found := k.GetOwner(ctx, id)
	if !found {
		return sdkerrors.Wrap(types.ErrUnknownIdentity, id.String())
	}

	if !owner.Equals(identityOwner) {
		return sdkerrors.Wrap(types.ErrNotAuthorized, "owner not matching")
	}

	if pubKey != nil {
		if err := k.AddPubKey(ctx, id, pubKey); err != nil {
			return err
		}
	}

	if len(certificate) > 0 {
		if err := k.AddCertificate(ctx, id, certificate); err != nil {
			return err
		}
	}

	if credentials != types.DoNotModifyDesc {
		k.SetCredentials(ctx, id, credentials)
	}

	return nil
}

// AddPubKey adds the given public key for the identity
func (k Keeper) AddPubKey(ctx sdk.Context, identityID tmbytes.HexBytes, pubKey *types.PubKeyInfo) error {
	pubKeyIdentityID, found := k.GetPubKeyIdentity(ctx, pubKey)
	if found {
		if !bytes.Equal(pubKeyIdentityID, identityID) {
			return sdkerrors.Wrap(types.ErrPubKeyExists, "")
		}
	} else {
		k.SetPubKey(ctx, identityID, pubKey)
	}

	return nil
}

// AddCertificate adds the given certificate for the identity
func (k Keeper) AddCertificate(ctx sdk.Context, identityID tmbytes.HexBytes, certificate string) error {
	cert := strings.TrimSpace(certificate)
	certHash := tmhash.Sum([]byte(cert))

	if !k.HasCertificate(ctx, identityID, certHash) {
		k.SetCertificate(ctx, identityID, certHash, cert)

		certPubKey := types.GetPubKeyFromCertificate([]byte(cert))
		return k.AddPubKey(ctx, identityID, certPubKey)
	}

	return nil
}

// SetOwner sets the owner of the given identity
func (k Keeper) SetOwner(ctx sdk.Context, identityID tmbytes.HexBytes, owner sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetOwnerKey(identityID), owner.Bytes())
}

// GetOwner gets the owner of the specified identity
func (k Keeper) GetOwner(ctx sdk.Context, identityID tmbytes.HexBytes) (sdk.AccAddress, bool) {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.GetOwnerKey(identityID))
	if bz == nil {
		return nil, false
	}

	return sdk.AccAddress(bz), true
}

// SetPubKey sets the given public key
func (k Keeper) SetPubKey(ctx sdk.Context, identityID tmbytes.HexBytes, pubKey *types.PubKeyInfo) {
	store := ctx.KVStore(k.storeKey)

	store.Set(types.GetPubKeyInfoKey(identityID, pubKey), []byte{})
	store.Set(types.GetPubKeyIdentityKey(pubKey), identityID)
}

// GetPubKeyIdentity gets the identity ID of the specified public key
func (k Keeper) GetPubKeyIdentity(ctx sdk.Context, pubKey *types.PubKeyInfo) (tmbytes.HexBytes, bool) {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.GetPubKeyIdentityKey(pubKey))
	if bz == nil {
		return nil, false
	}

	return tmbytes.HexBytes(bz), true
}

// SetCertificate sets the given certificate
func (k Keeper) SetCertificate(
	ctx sdk.Context,
	identityID tmbytes.HexBytes,
	certHash []byte,
	certificate string,
) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetCertificateKey(identityID, certHash), []byte(certificate))
}

// SetCredentials sets the given credentials
func (k Keeper) SetCredentials(ctx sdk.Context, identityID tmbytes.HexBytes, credentials string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetCredentialsKey(identityID), []byte(credentials))
}

// GetCredentials retrieves the credentials of the specified identity
func (k Keeper) GetCredentials(ctx sdk.Context, identityID tmbytes.HexBytes) (string, bool) {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.GetCredentialsKey(identityID))
	if bz == nil {
		return "", false
	}

	return string(bz), true
}

// HasIdentity returns true if the specified identity exists, false otherwise
func (k Keeper) HasIdentity(ctx sdk.Context, id tmbytes.HexBytes) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetOwnerKey(id))
}

// HasCertificate returns true if the specified certificate exists for the identity, false otherwise
func (k Keeper) HasCertificate(ctx sdk.Context, id tmbytes.HexBytes, certHash []byte) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetCertificateKey(id, certHash))
}

// SetIdentity sets the given identity
func (k Keeper) SetIdentity(ctx sdk.Context, identity types.Identity) {
	k.SetOwner(ctx, identity.Id, identity.Owner)

	for _, pk := range identity.PubKeys {
		k.SetPubKey(ctx, identity.Id, &pk)
	}

	for _, cert := range identity.Certificates {
		cert = strings.TrimSpace(cert)
		certHash := tmhash.Sum([]byte(cert))
		certPubKey := types.GetPubKeyFromCertificate([]byte(cert))

		k.SetCertificate(ctx, identity.Id, certHash, cert)
		k.SetPubKey(ctx, identity.Id, certPubKey)
	}

	if len(identity.Credentials) > 0 {
		k.SetCredentials(ctx, identity.Id, identity.Credentials)
	}
}

// GetIdentity retrieves the identity of the specified ID
func (k Keeper) GetIdentity(ctx sdk.Context, id tmbytes.HexBytes) (identity types.Identity, found bool) {
	owner, found := k.GetOwner(ctx, id)
	if !found {
		return identity, false
	}

	pubKeys := make([]types.PubKeyInfo, 0)
	certificates := make([]string, 0)

	k.IteratePubKeys(
		ctx, id,
		func(pubKey types.PubKeyInfo) (stop bool) {
			pubKeys = append(pubKeys, pubKey)
			return false
		},
	)

	k.IterateCertificates(
		ctx, id,
		func(cert string) (stop bool) {
			certificates = append(certificates, cert)
			return false
		},
	)

	credentials, _ := k.GetCredentials(ctx, id)

	identity.Id = id
	identity.PubKeys = pubKeys
	identity.Certificates = certificates
	identity.Credentials = credentials
	identity.Owner = owner

	return identity, true
}

// IteratePubKeys iterates through all public keys with the specified identity
func (k Keeper) IteratePubKeys(
	ctx sdk.Context,
	identityID tmbytes.HexBytes,
	op func(pubKey types.PubKeyInfo) (stop bool),
) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.GetPubKeySubspace(identityID))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		pubKeyInfoKey := iterator.Key()[1+types.IDLength:]

		pubKeyAlgo := types.PubKeyAlgorithm(binary.BigEndian.Uint32(pubKeyInfoKey[0:4]))
		pubKey := pubKeyInfoKey[4:]

		pubKeyInfo := types.PubKeyInfo{PubKey: pubKey, Algorithm: pubKeyAlgo}

		if stop := op(pubKeyInfo); stop {
			break
		}
	}
}

// IterateCertificates iterates through all certificates with the specified identity
func (k Keeper) IterateCertificates(
	ctx sdk.Context,
	identityID tmbytes.HexBytes,
	op func(cert string) (stop bool),
) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.GetCertificateSubspace(identityID))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		cert := iterator.Value()

		if stop := op(string(cert)); stop {
			break
		}
	}
}

// IterateIdentities iterates through all identities
func (k Keeper) IterateIdentities(
	ctx sdk.Context,
	op func(identity types.Identity) (stop bool),
) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.OwnerKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		id := iterator.Key()[1:]
		identity, _ := k.GetIdentity(ctx, id)

		if stop := op(identity); stop {
			break
		}
	}
}
