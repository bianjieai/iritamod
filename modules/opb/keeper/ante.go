package keeper

import (
	"github.com/bianjieai/iritamod/modules/opb/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

type ValidateFn func(ctx sdk.Context, msg sdk.Msg) error

// ValidateTokenTransferDecorator checks if the token transfer satisfies the underlying constraint
type ValidateTokenTransferDecorator struct {
	keeper      Keeper
	tokenKeeper types.TokenKeeper
	permKeeper  types.PermKeeper
	validators  map[string]ValidateFn
}

// NewValidateTokenTransferDecorator constructs a new ValidateTokenTransferDecorator instance
func NewValidateTokenTransferDecorator(
	keeper Keeper,
	tokenKeeper types.TokenKeeper,
	permKeeper types.PermKeeper,
) *ValidateTokenTransferDecorator {
	return &ValidateTokenTransferDecorator{
		keeper:      keeper,
		tokenKeeper: tokenKeeper,
		permKeeper:  permKeeper,
		validators:  make(map[string]ValidateFn),
	}
}

func (vtd *ValidateTokenTransferDecorator) DefaultValidateFn() *ValidateTokenTransferDecorator {
	return vtd.Append(&banktypes.MsgSend{}, vtd.validateMsgSend).
		Append(&banktypes.MsgMultiSend{}, vtd.validateMsgMultiSend)
}

func (vtd *ValidateTokenTransferDecorator) Append(
	m sdk.Msg,
	fn ValidateFn,
) *ValidateTokenTransferDecorator {
	vtd.validators[sdk.MsgTypeURL(m)] = fn
	return vtd
}

// AnteHandle implements AnteHandler
func (vtd ValidateTokenTransferDecorator) AnteHandle(
	ctx sdk.Context,
	tx sdk.Tx,
	simulate bool,
	next sdk.AnteHandler,
) (newCtx sdk.Context, err error) {
	restrictionEnabled := !vtd.keeper.UnrestrictedTokenTransfer(ctx)

	// check only if the transfer restriction is enabled
	if restrictionEnabled {
		for _, msg := range tx.GetMsgs() {
			validateFn, ok := vtd.validators[sdk.MsgTypeURL(msg)]
			if !ok {
				continue
			}
			if err := validateFn(ctx, msg); err != nil {
				return ctx, err
			}
		}
	}
	return next(ctx, tx, simulate)
}

// validateMsgSend validates the MsgSend msg
func (vtd ValidateTokenTransferDecorator) validateMsgSend(ctx sdk.Context, m sdk.Msg) error {
	msg, ok := m.(*banktypes.MsgSend)
	if !ok {
		return nil
	}
	for _, coin := range msg.Amount {
		owner, err := vtd.getOwner(ctx, coin.Denom)
		if err != nil {
			continue
		}
		fromAddress, err := sdk.AccAddressFromBech32(msg.FromAddress)
		if err != nil {
			continue
		}

		toAddress, err := sdk.AccAddressFromBech32(msg.ToAddress)
		if err != nil {
			continue
		}

		// If sender have platform user permissions, you can transfer token
		if vtd.hasPlatformUserPerm(ctx, fromAddress) ||
			vtd.hasPlatformUserPerm(ctx, toAddress) {
			return nil
		}
		// If sender have not platform user permissions,
		// determine whether the recipient is the owner

		if msg.FromAddress != owner && msg.ToAddress != owner {
			return sdkerrors.Wrapf(
				types.ErrUnauthorized,
				"either the sender or recipient must be the owner %s for token %s",
				owner, coin.Denom,
			)
		}
	}

	return nil
}

// validateMsgMultiSend validates the MsgMultiSend msg
func (vtd ValidateTokenTransferDecorator) validateMsgMultiSend(ctx sdk.Context, m sdk.Msg) error {
	msg, ok := m.(*banktypes.MsgMultiSend)
	if !ok {
		return nil
	}
	inputMap := getInputMap(msg.Inputs)
	outputMap := getOutputMap(msg.Outputs)

	for denom, addresses := range inputMap {
		owner, err := vtd.getOwner(ctx, denom)
		if err != nil {
			continue
		}
		// If sender have platform user permissions, you can transfer token
		if vtd.hasPlatformUserPermFromArr(ctx, addresses) ||
			vtd.hasPlatformUserPermFromArr(ctx, outputMap[denom]) {
			return nil
		}

		if !owned(owner, addresses) && !owned(owner, outputMap[denom]) {
			return sdkerrors.Wrapf(
				types.ErrUnauthorized,
				"either the sender or recipient must be the owner %s for token %s",
				owner, denom,
			)
		}
	}

	return nil
}

// getOwner gets the owner of the specified denom
func (vtd ValidateTokenTransferDecorator) getOwner(
	ctx sdk.Context,
	denom string,
) (owner string, err error) {
	baseTokenDenom := vtd.keeper.BaseTokenDenom(ctx)

	if denom == baseTokenDenom {
		owner = vtd.keeper.BaseTokenManager(ctx)
	} else {
		ownerAddr, err := vtd.tokenKeeper.GetOwner(ctx, denom)
		if err == nil {
			owner = ownerAddr.String()
		}
	}

	return
}

// hasPlatformUserPermFromArr determine whether the account is a platform user from addresses
func (vtd ValidateTokenTransferDecorator) hasPlatformUserPermFromArr(
	ctx sdk.Context,
	addresses []string,
) bool {
	for _, addr := range addresses {
		fromAddress, err := sdk.AccAddressFromBech32(addr)
		if err != nil {
			return false
		}
		if !vtd.hasPlatformUserPerm(ctx, fromAddress) {
			return false
		}
	}

	return true
}

// hasPlatformUserPerm determine whether the account is a platform user
func (vtd ValidateTokenTransferDecorator) hasPlatformUserPerm(
	ctx sdk.Context,
	address sdk.AccAddress,
) bool {
	return vtd.permKeeper.IsRootAdmin(ctx, address) || vtd.permKeeper.IsBaseM1Admin(ctx, address) ||
		vtd.permKeeper.IsPlatformUser(ctx, address)
}

// owned returns false if any address is not the owner of the denom among the given non-empty addresses
// True otherwise
func owned(owner string, addresses []string) bool {
	for _, addr := range addresses {
		if addr != owner {
			return false
		}
	}

	return true
}

// getInputMap maps input denoms to addresses
func getInputMap(inputs []banktypes.Input) map[string][]string {
	inputMap := make(map[string][]string)

	for _, input := range inputs {
		for _, coin := range input.Coins {
			inputMap[coin.Denom] = append(inputMap[coin.Denom], input.Address)
		}
	}

	return inputMap
}

// getOutputMap maps output denoms to addresses
func getOutputMap(outputs []banktypes.Output) map[string][]string {
	outputMap := make(map[string][]string)

	for _, output := range outputs {
		for _, coin := range output.Coins {
			outputMap[coin.Denom] = append(outputMap[coin.Denom], output.Address)
		}
	}

	return outputMap
}
