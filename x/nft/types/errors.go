package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/nft module sentinel errors
var (
	ErrInvalidDenom  = sdkerrors.Register(ModuleName, 1, "invalid denom")
	ErrInvalidNFTID  = sdkerrors.Register(ModuleName, 2, "invalid nft ID")
	ErrInvalidNFTURL = sdkerrors.Register(ModuleName, 3, "invalid nft url")
)
