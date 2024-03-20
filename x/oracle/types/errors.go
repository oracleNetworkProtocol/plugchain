package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/oracle module errors
var (
	ErrInvalidClaim        = sdkerrors.Register(ModuleName, 2, "invalid claim")
	ErrNoClaimExists       = sdkerrors.Register(ModuleName, 3, "no claim exits")
	ErrNoClaimTypeExists   = sdkerrors.Register(ModuleName, 4, "claim type is not registered as part of the oracle params")
	ErrNoPrevote           = sdkerrors.Register(ModuleName, 5, "no prevote exists for this claim")
	ErrIncorrectClaimRound = sdkerrors.Register(ModuleName, 6, "claim must be submitted after the prevote round is over")
)
