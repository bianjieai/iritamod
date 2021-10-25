package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrUnauthorizedOperation = sdkerrors.Register(ModuleName, 2, "unauthorized operation")
	ErrAlreadyBlockedAccount = sdkerrors.Register(ModuleName, 3, "account already blocked")
	ErrBlockAdminAccount     = sdkerrors.Register(ModuleName, 4, "can not block admin account")
	ErrUnknownBlockedAccount = sdkerrors.Register(ModuleName, 5, "unknown blocked account")
	ErrAddRootAdmin          = sdkerrors.Register(ModuleName, 6, "can not add root account")
	ErrRemoveRootAdmin       = sdkerrors.Register(ModuleName, 7, "can not remove root account")
	ErrRemoveUnknownRole     = sdkerrors.Register(ModuleName, 8, "the account does not have this role")
	ErrInvalidMsgURL         = sdkerrors.Register(ModuleName, 9, "invalid url")

	ErrOperateRootAdmin = sdkerrors.Wrap(ErrUnauthorizedOperation, "can not operate root admin")
	ErrOperatePermAdmin = sdkerrors.Wrap(ErrUnauthorizedOperation, "can not operate another permission admin")
)
