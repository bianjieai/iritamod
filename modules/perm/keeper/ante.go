package keeper

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	evmtypes "github.com/tharsis/ethermint/x/evm/types"

	"github.com/bianjieai/iritamod/modules/perm/types"
)

type AuthDecorator struct {
	k Keeper
}

func NewAuthDecorator(k Keeper) AuthDecorator {
	return AuthDecorator{k: k}
}

// AnteHandle returns an AnteHandler that checks the auth to send msg
func (ad AuthDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	for _, msg := range tx.GetMsgs() {
		for _, signer := range msg.GetSigners() {
			if ad.k.GetBlockAccount(ctx, signer) {
				return ctx, sdkerrors.Wrapf(types.ErrUnauthorizedOperation, "The sender %s has been blocked", signer)
			}
			url := sdk.MsgTypeURL(msg)
			if auth, ok := ad.k.AuthMap[url]; ok {
				if err := ad.k.Access(ctx, signer, auth); err != nil {
					return ctx, err
				}
				continue
			}
			route := strings.Split(url, ".")
			if len(route) <= 2 {
				return ctx, sdkerrors.Wrapf(types.ErrInvalidMsgURL, "the url %s is invalid", url)
			}
			if auth, ok := ad.k.AuthMap[route[1]]; ok {
				if err := ad.k.Access(ctx, signer, auth); err != nil {
					return ctx, err
				}
			}
		}
	}
	// continue
	return next(ctx, tx, simulate)
}

type EthCanCallDecorator struct {
	keeper Keeper
}

func NewEthCanCallDecorator(Keeper Keeper) EthCanCallDecorator {
	return EthCanCallDecorator{keeper: Keeper}
}

func (e EthCanCallDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	for _, msg := range tx.GetMsgs() {
		msgEthTx, ok := msg.(*evmtypes.MsgEthereumTx)
		if !ok {
			return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "invalid transaction type %T, expected %T", tx, (*evmtypes.MsgEthereumTx)(nil))
		}
		ethTx := msgEthTx.AsTransaction()
		if ethTx.To() != nil {
			state := e.keeper.GetBlockContract(ctx, *ethTx.To())
			if state {
				return ctx, sdkerrors.Wrapf(types.ErrContractDisable, "the contract %s is in contract deny list ! ", ethTx.To())
			}
		}
	}
	return next(ctx, tx, simulate)
}
