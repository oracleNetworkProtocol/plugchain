package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrSymbolLenNotValid   = sdkerrors.Register(ModuleName, 1, "Invalid symbol length")
	ErrDescriptionNotValid = sdkerrors.Register(ModuleName, 2, "Invalid description length")
	ErrTotalSupplyNotValid = sdkerrors.Register(ModuleName, 3, "Invalid total-supply")
	ErrSymbolIsExists      = sdkerrors.Register(ModuleName, 4, "Determine whether the symbol exists")
)

func ReturnErrSymbolLenNotValid(maxLen int) error {
	return sdkerrors.Wrapf(ErrSymbolLenNotValid, "The symbol length exceeds the maximum limit length, the maximum length is (%s)", maxLen)
}

func ReturnErrDescriptionNotValid(maxBytes int) error {
	return sdkerrors.Wrapf(ErrDescriptionNotValid, "The description length exceeds the maximum length of (%s) bytes", maxBytes)
}

func ReturnErrSymbolIsExists(symbol string) error {
	return sdkerrors.Wrapf(ErrSymbolIsExists, "The symbol （%s） already exists and cannot be created", symbol)
}
