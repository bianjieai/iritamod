package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/bianjieai/iritamod/modules/layer2/types"
)

type ValidateLayer2Decorator struct {
	keeper     Keeper
	permKeeper types.PermKeeper
}

func NewValidateLayer2Decorator(keeper Keeper, permKeeper types.PermKeeper) ValidateLayer2Decorator {
	return ValidateLayer2Decorator{
		keeper:     keeper,
		permKeeper: permKeeper,
	}
}

func (dlt ValidateLayer2Decorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	for _, msg := range tx.GetMsgs() {
		switch msg := msg.(type) {
		case *types.MsgCreateL2Space:
			if err := dlt.validateL2UserRole(ctx, msg.Sender); err != nil {
				return ctx, err
			}
		case *types.MsgTransferL2Space:
			if err := dlt.validateSpaceOwnership(ctx, msg.Sender, msg.SpaceId); err != nil {
				return ctx, err
			}
			if err := dlt.validateL2UserRole(ctx, msg.Sender); err != nil {
				return ctx, err
			}
			if err := dlt.validateL2UserRole(ctx, msg.Recipient); err != nil {
				return ctx, err
			}
		case *types.MsgCreateL2BlockHeader:
			if err := dlt.validateSpaceOwnership(ctx, msg.Sender, msg.SpaceId); err != nil {
				return ctx, err
			}
			if err := dlt.validateL2UserRole(ctx, msg.Sender); err != nil {
				return ctx, err
			}
		case *types.MsgCreateNFTs:
			if err := dlt.validateSpaceOwnership(ctx, msg.Sender, msg.SpaceId); err != nil {
				return ctx, err
			}
			if err := dlt.validateL2UserRole(ctx, msg.Sender); err != nil {
				return ctx, err
			}
		case *types.MsgUpdateNFTs:
			if err := dlt.validateSpaceOwnership(ctx, msg.Sender, msg.SpaceId); err != nil {
				return ctx, err
			}
			if err := dlt.validateL2UserRole(ctx, msg.Sender); err != nil {
				return ctx, err
			}
		case *types.MsgDeleteNFTs:
			if err := dlt.validateSpaceOwnership(ctx, msg.Sender, msg.SpaceId); err != nil {
				return ctx, err
			}
			if err := dlt.validateL2UserRole(ctx, msg.Sender); err != nil {
				return ctx, err
			}
		case *types.MsgUpdateClassesForNFT:
			if err := dlt.validateSpaceOwnership(ctx, msg.Sender, msg.SpaceId); err != nil {
				return ctx, err
			}
			if err := dlt.validateL2UserRole(ctx, msg.Sender); err != nil {
				return ctx, err
			}
		case *types.MsgDepositClassForNFT:
			if err := dlt.validateOnlySpace(ctx, msg.SpaceId); err != nil {
				return ctx, err
			}
			if msg.Sender != msg.Recipient {
				if err := dlt.validateSpaceOwnership(ctx, msg.Sender, msg.SpaceId); err != nil {
					return ctx, sdkerrors.Wrapf(types.ErrInvalidL2User,
						"recipient (%s) must be the same as sender if sender (%s) is not l2 user", msg.Recipient, msg.Sender)
				}
			}
		case *types.MsgWithdrawClassForNFT:
			if err := dlt.validateSpaceOwnership(ctx, msg.Sender, msg.SpaceId); err != nil {
				return ctx, err
			}
			if err := dlt.validateL2UserRole(ctx, msg.Sender); err != nil {
				return ctx, err
			}
		case *types.MsgDepositTokenForNFT:
			if err := dlt.validateOnlySpace(ctx, msg.SpaceId); err != nil {
				return ctx, err
			}
		case *types.MsgWithdrawTokenForNFT:
			if err := dlt.validateSpaceOwnership(ctx, msg.Sender, msg.SpaceId); err != nil {
				return ctx, err
			}
			if err := dlt.validateL2UserRole(ctx, msg.Sender); err != nil {
				return ctx, err
			}
		}
	}

	return next(ctx, tx, simulate)
}

func (dlt ValidateLayer2Decorator) validateOnlySpace(ctx sdk.Context, spaceId uint64) error {
	if !dlt.keeper.HasSpace(ctx, spaceId) {
		return sdkerrors.Wrapf(types.ErrInvalidSpaceId, "space (%d) does not exist", spaceId)
	}
	return nil
}

func (dlt ValidateLayer2Decorator) validateSpaceOwnership(ctx sdk.Context, addr string, spaceId uint64) error {
	accAddr, _ := sdk.AccAddressFromBech32(addr)
	if !dlt.keeper.HasSpace(ctx, spaceId) {
		return sdkerrors.Wrapf(types.ErrInvalidSpaceId, "space (%d) does not exist", spaceId)
	}

	if !dlt.keeper.HasSpaceOfOwner(ctx, accAddr, spaceId) {
		return sdkerrors.Wrapf(types.ErrInvalidSpaceOwner, "space (%d) is not owned by (%s)", spaceId, addr)
	}
	return nil
}

func (dlt ValidateLayer2Decorator) validateL2UserRole(ctx sdk.Context, addr string) error {
	accAddr, _ := sdk.AccAddressFromBech32(addr)
	if !dlt.permKeeper.HasL2UserRole(ctx, accAddr) {
		return sdkerrors.Wrapf(types.ErrInvalidL2User, "account (%s) does not have l2 user role", addr)
	}
	return nil
}
