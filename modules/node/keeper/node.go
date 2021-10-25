package keeper

import (
	"github.com/tendermint/tendermint/crypto"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/bianjieai/iritamod/modules/node/types"
	cautils "github.com/bianjieai/iritamod/utils/ca"
)

// AddNode adds a node
func (k Keeper) AddNode(ctx sdk.Context, name string, cert string) (id tmbytes.HexBytes, err error) {
	pubKey, err := k.VerifyCertificate(ctx, cert)
	if err != nil {
		return nil, err
	}

	id = pubKey.Address()

	if k.HasNode(ctx, id) {
		return nil, sdkerrors.Wrap(types.ErrNodeExists, id.String())
	}

	node := types.NewNode(id, name, cert)
	k.SetNode(ctx, id, node)

	return id, nil
}

// RemoveNode removes the specified node
func (k Keeper) RemoveNode(ctx sdk.Context, id tmbytes.HexBytes) error {
	if !k.HasNode(ctx, id) {
		return sdkerrors.Wrap(types.ErrUnknownNode, id.String())
	}

	k.DeleteNode(ctx, id)

	return nil
}

// SetNode sets the given node
func (k Keeper) SetNode(ctx sdk.Context, id tmbytes.HexBytes, node types.Node) {
	store := ctx.KVStore(k.storeKey)

	bz := k.cdc.MustMarshal(&node)
	store.Set(types.GetNodeKey(id), bz)
}

// DeleteNode deletes the given node
func (k Keeper) DeleteNode(ctx sdk.Context, id tmbytes.HexBytes) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetNodeKey(id))
}

// HasNode returns true if the specified node exists, false otherwise
func (k Keeper) HasNode(ctx sdk.Context, id tmbytes.HexBytes) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetNodeKey(id))
}

// GetNode retrieves the node of the specified ID
func (k Keeper) GetNode(ctx sdk.Context, id tmbytes.HexBytes) (node types.Node, found bool) {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.GetNodeKey(id))
	if bz == nil {
		return node, false
	}

	k.cdc.MustUnmarshal(bz, &node)
	return node, true
}

// GetNodes gets all nodes
func (k Keeper) GetNodes(ctx sdk.Context) []types.Node {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.NodeKey)
	defer iterator.Close()

	nodes := make([]types.Node, 0)

	for ; iterator.Valid(); iterator.Next() {
		var node types.Node

		bz := iterator.Value()
		k.cdc.MustUnmarshal(bz, &node)

		nodes = append(nodes, node)
	}

	return nodes
}

// VerifyCertificate verifies the given certificate against the root certificate
// Ensure that the given certificate is a valid X.509 format
func (k Keeper) VerifyCertificate(ctx sdk.Context, certificate string) (crypto.PubKey, error) {
	cert, _ := cautils.ReadCertificateFromMem([]byte(certificate))

	rootCertStr, _ := k.GetRootCert(ctx)
	rootCert, _ := cautils.ReadCertificateFromMem([]byte(rootCertStr))

	if err := cautils.VerifyCertFromRoot(cert, rootCert); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidCert, "verification failed: %s", err)
	}

	pubKey, _ := cautils.GetPubkeyFromCert(cert)
	return pubKey, nil
}
