package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"gitlab.bianjie.ai/irita-pro/iritamod/modules/admin/types"
)

type AuthDecorator struct {
	k Keeper
}

func NewAuthDecorator(k Keeper) AuthDecorator {
	return AuthDecorator{
		k: k,
	}
}

// AnteHandle returns an AnteHandler that checks the auth to send msg
func (ad AuthDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	for _, msg := range tx.GetMsgs() {
		for _, signer := range msg.GetSigners() {
			if ad.k.GetBlockAccount(ctx, signer) {
				return ctx, sdkerrors.Wrapf(types.ErrUnauthorizedOperation,
					"The sender %s has been blocked", signer)
			}
			if auth, ok := ad.k.AuthMap[msg.Type()]; ok {
				err := ad.k.Access(ctx, signer, auth)
				if err != nil {
					return ctx, err
				}
				continue
			}
			// If both msg.Route and msg.Type are registered
			// Only verify msg.Type
			if auth, ok := ad.k.AuthMap[msg.Route()]; ok {
				err := ad.k.Access(ctx, signer, auth)
				if err != nil {
					return ctx, err
				}
			}
		}
	}
	// continue
	return next(ctx, tx, simulate)
}
