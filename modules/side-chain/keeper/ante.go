package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/bianjieai/iritamod/modules/side-chain/types"
)

type ValidateSideChainDecorator struct {
	keeper     Keeper
	permKeeper types.PermKeeper
}

func NewValidateSideChainDecorator(keeper Keeper, permKeeper types.PermKeeper) ValidateSideChainDecorator {
	return ValidateSideChainDecorator{
		keeper:     keeper,
		permKeeper: permKeeper,
	}
}

func (dlt ValidateSideChainDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	for _, msg := range tx.GetMsgs() {
		switch msg := msg.(type) {
		case *types.MsgCreateSpace:
			if err := dlt.validateSideChainUserRole(ctx, msg.Sender); err != nil {
				return ctx, err
			}
		case *types.MsgTransferSpace:
			if err := dlt.validateSpaceOwnership(ctx, msg.Sender, msg.SpaceId); err != nil {
				return ctx, err
			}
			if err := dlt.validateSideChainUserRole(ctx, msg.Sender); err != nil {
				return ctx, err
			}
			if err := dlt.validateSideChainUserRole(ctx, msg.Recipient); err != nil {
				return ctx, err
			}
		case *types.MsgCreateBlockHeader:
			if err := dlt.validateSpaceOwnership(ctx, msg.Sender, msg.SpaceId); err != nil {
				return ctx, err
			}
			if err := dlt.validateSideChainUserRole(ctx, msg.Sender); err != nil {
				return ctx, err
			}
		}
	}

	return next(ctx, tx, simulate)
}

func (dlt ValidateSideChainDecorator) validateOnlySpace(ctx sdk.Context, spaceId uint64) error {
	if !dlt.keeper.HasSpace(ctx, spaceId) {
		return sdkerrors.Wrapf(types.ErrInvalidSpaceId, "space (%d) does not exist", spaceId)
	}
	return nil
}

func (dlt ValidateSideChainDecorator) validateSpaceOwnership(ctx sdk.Context, addr string, spaceId uint64) error {
	accAddr, _ := sdk.AccAddressFromBech32(addr)
	if !dlt.keeper.HasSpace(ctx, spaceId) {
		return sdkerrors.Wrapf(types.ErrInvalidSpaceId, "space (%d) does not exist", spaceId)
	}

	if !dlt.keeper.HasSpaceOfOwner(ctx, accAddr, spaceId) {
		return sdkerrors.Wrapf(types.ErrInvalidSpaceOwner, "space (%d) is not owned by (%s)", spaceId, addr)
	}
	return nil
}

func (dlt ValidateSideChainDecorator) validateSideChainUserRole(ctx sdk.Context, addr string) error {
	accAddr, _ := sdk.AccAddressFromBech32(addr)
	if !dlt.permKeeper.HasSideChainUserRole(ctx, accAddr) {
		return sdkerrors.Wrapf(types.ErrInvalidSideChainUser, "account (%s) does not have l2 user role", addr)
	}
	return nil
}
