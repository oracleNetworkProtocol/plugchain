package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/nft module sentinel errors
var (
	ErrInvalidDenom = sdkerrors.Register(ModuleName, 1, "invalid denom")
)
