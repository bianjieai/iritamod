package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	layer2types "github.com/bianjieai/iritamod/modules/layer2/types"
	"github.com/bianjieai/iritamod/modules/perm"
	permKeeper "github.com/bianjieai/iritamod/modules/perm/keeper"
)

type ValidateLayer2Decorator struct {
	keeper     Keeper
	permKeeper permKeeper.Keeper
}

func NewValidateLayer2Decorator(keeper Keeper, permKeeper permKeeper.Keeper) ValidateLayer2Decorator {
	return ValidateLayer2Decorator{
		keeper:     keeper,
		permKeeper: permKeeper,
	}
}

func (dlt ValidateLayer2Decorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	for _, msg := range tx.GetMsgs() {
		switch msg := msg.(type) {
		case *layer2types.MsgCreateL2Space:
			if err := dlt.validateL2UserRole(ctx, msg.Sender); err != nil {
				return ctx, err
			}
		case *layer2types.MsgTransferL2Space:
			if err := dlt.validateSpaceOwnership(ctx, msg.Sender, msg.SpaceId); err != nil {
				return ctx, err
			}
			if err := dlt.validateL2UserRole(ctx, msg.Sender); err != nil {
				return ctx, err
			}
			if err := dlt.validateL2UserRole(ctx, msg.Recipient); err != nil {
				return ctx, err
			}
		case *layer2types.MsgCreateL2BlockHeader:
			if err := dlt.validateSpaceOwnership(ctx, msg.Sender, msg.SpaceId); err != nil {
				return ctx, err
			}
			if err := dlt.validateL2UserRole(ctx, msg.Sender); err != nil {
				return ctx, err
			}
		case *layer2types.MsgCreateNFTs:
			if err := dlt.validateSpaceOwnership(ctx, msg.Sender, msg.SpaceId); err != nil {
				return ctx, err
			}
			if err := dlt.validateL2UserRole(ctx, msg.Sender); err != nil {
				return ctx, err
			}
		case *layer2types.MsgUpdateNFTs:
			if err := dlt.validateSpaceOwnership(ctx, msg.Sender, msg.SpaceId); err != nil {
				return ctx, err
			}
			if err := dlt.validateL2UserRole(ctx, msg.Sender); err != nil {
				return ctx, err
			}
		case *layer2types.MsgDeleteNFTs:
			if err := dlt.validateSpaceOwnership(ctx, msg.Sender, msg.SpaceId); err != nil {
				return ctx, err
			}
			if err := dlt.validateL2UserRole(ctx, msg.Sender); err != nil {
				return ctx, err
			}
		case *layer2types.MsgUpdateClassesForNFT:
			if err := dlt.validateSpaceOwnership(ctx, msg.Sender, msg.SpaceId); err != nil {
				return ctx, err
			}
			if err := dlt.validateL2UserRole(ctx, msg.Sender); err != nil {
				return ctx, err
			}
		case *layer2types.MsgDepositClassForNFT:
			if err := dlt.validateOnlySpace(ctx, msg.SpaceId); err != nil {
				return ctx, err
			}
			if msg.Sender != msg.Recipient {
				if err := dlt.validateSpaceOwnership(ctx, msg.Sender, msg.SpaceId); err != nil {
					return ctx, sdkerrors.Wrapf(layer2types.ErrInvalidL2User,
						"recipient (%s) must be the same as sender if it is not l2 user", msg.Recipient)
				}
			}
		case *layer2types.MsgWithdrawClassForNFT:
			if err := dlt.validateSpaceOwnership(ctx, msg.Sender, msg.SpaceId); err != nil {
				return ctx, err
			}
			if err := dlt.validateL2UserRole(ctx, msg.Sender); err != nil {
				return ctx, err
			}
		case *layer2types.MsgDepositTokenForNFT:
			if err := dlt.validateOnlySpace(ctx, msg.SpaceId); err != nil {
				return ctx, err
			}
		case *layer2types.MsgWithdrawTokenForNFT:
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
		return sdkerrors.Wrapf(layer2types.ErrInvalidSpace, "space (%d) not exist", spaceId)
	}
	return nil
}

func (dlt ValidateLayer2Decorator) validateSpaceOwnership(ctx sdk.Context, addr string, spaceId uint64) error {
	accAddr, _ := sdk.AccAddressFromBech32(addr)
	if !dlt.keeper.HasSpace(ctx, spaceId) {
		return sdkerrors.Wrapf(layer2types.ErrInvalidSpace, "space (%d) not exist", spaceId)
	}

	if !dlt.keeper.HasSpaceOfOwner(ctx, accAddr, spaceId) {
		return sdkerrors.Wrapf(layer2types.ErrNotOwnerOfSpace, "space (%d) is not owned by (%s)", spaceId, accAddr)
	}
	return nil
}

func (dlt ValidateLayer2Decorator) validateL2UserRole(ctx sdk.Context, addr string) error {
	accAddr, _ := sdk.AccAddressFromBech32(addr)
	if dlt.permKeeper.IsRootAdmin(ctx, accAddr) {
		return nil
	}

	if err := dlt.permKeeper.Access(ctx, accAddr, perm.RoleLayer2User.Auth()); err != nil {
		return sdkerrors.Wrapf(layer2types.ErrInvalidL2User, "addr (%s) is not l2 user", accAddr)
	}
	return nil
}
