package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/nft module sentinel errors
var (
	ErrInvalidClass      = sdkerrors.Register(ModuleName, 2, "invalid denom")
	ErrInvalidNFTID      = sdkerrors.Register(ModuleName, 3, "invalid nft ID")
	ErrInvalidNFTURL     = sdkerrors.Register(ModuleName, 4, "invalid nft url")
	ErrNFTAreadyExists   = sdkerrors.Register(ModuleName, 5, "nft already exists")
	ErrUnknownCollection = sdkerrors.Register(ModuleName, 6, "unknown nft collection")
	ErrUnauthorized      = sdkerrors.Register(ModuleName, 7, "unauthorized address")
	ErrKnownNFT          = sdkerrors.Register(ModuleName, 8, "unknown nft")
)
