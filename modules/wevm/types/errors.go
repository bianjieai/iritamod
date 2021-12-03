package types

import (
	"errors"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	codeErrInvalidState = uint32(iota) + 2 // NOTE: code 1 is reserved for internal errors
	codeErrNotFound
	codeErrInvalidContractAddress
	codeErrContractDisable
	codeErrContractAlreadyExist
)

var ErrPostTxProcessing = errors.New("failed to execute post processing")

var (
	// ErrInvalidState returns an error resulting from an invalid Storage State.
	ErrInvalidState           = sdkerrors.Register(ModuleName, codeErrInvalidState, "invalid storage state")
	ErrNotFound               = sdkerrors.Register(ModuleName, codeErrNotFound, "not found")
	ErrInvalidContractAddress = sdkerrors.Register(ModuleName, codeErrInvalidContractAddress, "contract address is invalid")
	ErrContractDisable        = sdkerrors.Register(ModuleName, codeErrContractDisable, "contract is disable")
	ErrContractAlreadyExist   = sdkerrors.Register(ModuleName, codeErrContractAlreadyExist, "contract already exist")
)
