package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrSymbolLenNotValid       = sdkerrors.Register(ModuleName, 1, "Invalid symbol length")
	ErrDescriptionNotValid     = sdkerrors.Register(ModuleName, 2, "Invalid description length")
	ErrTotalSupplyNotValid     = sdkerrors.Register(ModuleName, 3, "Invalid total-supply")
	ErrSymbolIsExists          = sdkerrors.Register(ModuleName, 4, "Determine whether the symbol exists")
	ErrDecimalNotValid         = sdkerrors.Register(ModuleName, 5, "Invalid decimal")
	ErrAccAddressNotPermission = sdkerrors.Register(ModuleName, 6, "Invalid address")
	ErrBurnTokenNotPermission  = sdkerrors.Register(ModuleName, 7, "This symbol coinage is prohibited")
	ErrBurnTokenInsufficient   = sdkerrors.Register(ModuleName, 8, "Invalid account")
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

func ReturnErrTotalSupplyNotValid(total uint64) error {
	return sdkerrors.Wrapf(ErrTotalSupplyNotValid, "The minimum circulation is (%s)", total)
}

func ReturnErrDecimalNotValid(decimal uint64) error {
	return sdkerrors.Wrapf(ErrDecimalNotValid, "The minimum number of decimal places is (%s)", decimal)
}
func ReturnErrAccAddressNotPermission(address string) error {
	return sdkerrors.Wrapf(ErrAccAddressNotPermission, "The current address (%s) does not have the operation authority for this function", address)
}

func ReturnErrBurnTokenInsufficient(addrAccount, burnAccount string) error {
	return sdkerrors.Wrapf(ErrBurnTokenInsufficient, "(%s) is smaller than (%s): insufficient funds", addrAccount, burnAccount)
}
