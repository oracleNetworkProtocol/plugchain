package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/prc10 module sentinel errors
var (
	ErrInvalidName         = sdkerrors.Register(ModuleName, 2, "invalid token name")
	ErrInvalidMinUnit      = sdkerrors.Register(ModuleName, 3, "invalid token min unit")
	ErrInvalidSymbol       = sdkerrors.Register(ModuleName, 4, "invalid standard denom")
	ErrInvalidInitSupply   = sdkerrors.Register(ModuleName, 5, "invalid token initial supply")
	ErrInvalidMaxSupply    = sdkerrors.Register(ModuleName, 6, "invalid token maximum supply")
	ErrInvalidScale        = sdkerrors.Register(ModuleName, 7, "invalid token scale")
	ErrSymbolAlreadyExists = sdkerrors.Register(ModuleName, 8, "symbol already exists")
	ErrTokenNotExists      = sdkerrors.Register(ModuleName, 10, "token does not exist")
	ErrInvalidOwner        = sdkerrors.Register(ModuleName, 11, "invalid token owner")
	ErrNotMintable         = sdkerrors.Register(ModuleName, 13, "token is not mintable")
	ErrInvalidAmount       = sdkerrors.Register(ModuleName, 15, "invalid amount")
)
