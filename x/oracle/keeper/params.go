package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/oracleNetworkProtocol/plugchain/x/oracle/types"
)

// ClaimParams returns all claim params.
func (k Keeper) ClaimParams(ctx sdk.Context) (res map[string](types.ClaimParams)) {
	k.paramspace.Get(ctx, types.KeyClaimParams, &res)
	return
}

// ClaimParamsForType returns claim params for a given claimType.
func (k Keeper) ClaimParamsForType(ctx sdk.Context, claimType string) (res types.ClaimParams) {
	res = k.ClaimParams(ctx)[claimType]
	return
}

// GetParams returns the total set of oracle parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramspace.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the total set of oracle parameters.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramspace.SetParamSet(ctx, &params)
}
