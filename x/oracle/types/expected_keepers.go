package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

type (
	// StakingKeeper defines the staking module interface contract needed by the
	// claim module.
	StakingKeeper interface {
		// ValidatorByConsAddr(sdk.Context, sdk.ConsAddress) stakingexported.ValidatorI
		Validator(ctx sdk.Context, address sdk.ValAddress) stakingtypes.ValidatorI // get validator by operator address; nil when validator not found

		TotalBondedTokens(sdk.Context) sdk.Int // total bonded tokens within the validator set
	}

	// ParamSubspace defines the expected Subspace interfacace
	ParamSubspace interface {
		HasKeyTable() bool
		WithKeyTable(table paramtypes.KeyTable) paramtypes.Subspace
		Get(ctx sdk.Context, key []byte, ptr interface{})
		GetParamSet(ctx sdk.Context, ps paramtypes.ParamSet)
		SetParamSet(ctx sdk.Context, ps paramtypes.ParamSet)
	}
)
