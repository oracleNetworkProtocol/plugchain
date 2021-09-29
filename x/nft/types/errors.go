package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/nft module sentinel errors
var (
	ErrInvalidDenom      = sdkerrors.Register(ModuleName, 1, "invalid denom")
	ErrInvalidNFTID      = sdkerrors.Register(ModuleName, 2, "invalid nft ID")
	ErrInvalidNFTURL     = sdkerrors.Register(ModuleName, 3, "invalid nft url")
	ErrNFTAreadyExists   = sdkerrors.Register(ModuleName, 4, "nft already exists")
	ErrUnknownCollection = sdkerrors.Register(ModuleName, 5, "unknown nft collection")
	ErrUnauthorized      = sdkerrors.Register(ModuleName, 6, "unauthorized address")
)
