package keeper

import (
	"fmt"
	"sort"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/oracleNetworkProtocol/plugchain/x/atom/types"
)

func (k Keeper) SetAtomUsd(ctx sdk.Context, atomUsd types.AtomUsd) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&atomUsd)
	store.Set(types.AtomUsdKey, b)
}

func (k Keeper) GetAtomUsd(ctx sdk.Context) *types.AtomUsd {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.AtomUsdKey)
	if len(bz) == 0 {
		return nil
	}

	var atomUsd types.AtomUsd
	k.cdc.MustUnmarshal(bz, &atomUsd)
	return &atomUsd
}

func (k Keeper) UpdateAtomUsd(ctx sdk.Context) {
	claimType := types.AtomClaim
	pending := k.oracleKeeper.GetPendingRounds(ctx, claimType)

	// sort pending rounds in ascending order
	sort.SliceStable(pending, func(i, j int) bool { return pending[i] < pending[j] })

	for _, roundID := range pending {
		result := k.oracleKeeper.TallyVotes(ctx, claimType, roundID)

		if result == nil || result.Claims[0] == nil {
			continue
		}

		// take an average of all claims and commit to chain
		avgAtomUsd := sdk.NewDec(0)
		var blockHeight int64
		var totalVotePower int64
		for _, claimResult := range result.Claims {
			claimHash := claimResult.ClaimHash
			atomClaim, ok := k.oracleKeeper.GetClaim(ctx, claimHash).(*types.AtomUsd)

			if ok == false {
				fmt.Printf("Error retrieving claim")
				continue
			}
			weightedAvg := avgAtomUsd.Mul(sdk.NewDec(totalVotePower))
			weightedVote := atomClaim.Price.Mul(sdk.NewDec(result.VotePower))
			totalVotePower += result.VotePower
			avgAtomUsd = weightedAvg.Add(weightedVote).Quo(sdk.NewDec(totalVotePower))
			blockHeight = atomClaim.BlockHeight
		}

		atomUsd := &types.AtomUsd{
			Price:       avgAtomUsd,
			BlockHeight: blockHeight,
		}
		k.SetAtomUsd(ctx, *atomUsd)

		// TODO delete the any earlier pending rounds
		k.oracleKeeper.FinalizeRound(ctx, claimType, roundID)
	}
}
