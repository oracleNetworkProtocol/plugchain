package keeper

import (
	"bytes"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/oracleNetworkProtocol/plugchain/x/oracle/exported"
	"github.com/oracleNetworkProtocol/plugchain/x/oracle/types"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
)

////////////////////
/// VOTE
////////////////////

// CreateVote casts a vote for a given claim.
func (k Keeper) CreateVote(ctx sdk.Context, claim exported.Claim, validator sdk.ValAddress) {
	k.CreateClaim(ctx, claim)
	roundID := claim.GetRoundID()
	claimType := claim.Type()
	vote := types.NewVote(roundID, claim, validator, claimType)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RoundKey)

	var votes types.Round
	bz := store.Get(types.RoundPrefix(claimType, roundID))
	if len(bz) == 0 {
		votes = types.Round{
			Votes:     []types.Vote{*vote},
			RoundId:   roundID,
			ClaimType: claimType,
		}
	} else {
		k.cdc.MustUnmarshal(bz, &votes)
		votes.Votes = append(votes.Votes, *vote)
	}

	k.AddPendingRound(ctx, vote.ClaimType, vote.RoundId)
	store.Set(types.RoundPrefix(claimType, roundID), k.cdc.MustMarshal(&votes))
}

// DeleteVotesForRound deletes all votes and claims for a given round and claimType
func (k Keeper) DeleteVotesForRound(ctx sdk.Context, claimType string, roundID uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RoundKey)
	roundKey := types.GetRoundKey(claimType, roundID)

	round := k.GetRound(ctx, claimType, roundID)
	if round == nil {
		return
	}

	for _, vote := range round.Votes {
		k.DeleteClaim(ctx, []byte(vote.ClaimHash))
	}
	store.Delete(types.KeyPrefix(roundKey))
}

// TallyVotes tallies up the votes for a given Claim and returns the result with the maximum claim
// vote.ClaimHash
func (k Keeper) TallyVotes(ctx sdk.Context, claimType string, roundID uint64) *types.RoundResult {
	votes := k.GetRound(ctx, claimType, roundID)

	voteMap := make(map[string]*types.RoundResult)
	var maxVotePower int64
	var maxVoteKey string
	for _, vote := range votes.Votes {
		validator := k.StakingKeeper.Validator(ctx, vote.Validator)
		if validator == nil || !validator.IsBonded() || validator.IsJailed() {
			continue
		}
		weight := validator.GetConsensusPower(sdk.Int{})

		key := vote.GetConsensusId()

		resultObject := voteMap[key]
		if resultObject == nil {
			resultObject = &types.RoundResult{
				ClaimType: vote.GetClaimType(),
			}
		}

		resultObject.UpsertClaim(vote.ClaimHash, weight)
		resultObject.VotePower += weight

		voteMap[key] = resultObject

		if resultObject.VotePower > maxVotePower {
			maxVoteKey = key
		}
	}

	result := voteMap[maxVoteKey]
	if result == nil {
		return nil
	}

	passesThreshold, totalBondedPower := k.VotePassedThreshold(ctx, result)

	if passesThreshold {
		result.TotalPower = totalBondedPower
		return result
	}
	return nil
}

// VotePassedThreshold checks if a given claim has cleared the vote threshold
func (k Keeper) VotePassedThreshold(ctx sdk.Context, roundResult *types.RoundResult) (bool, int64) {
	totalBondedPower := sdk.TokensToConsensusPower(k.StakingKeeper.TotalBondedTokens(ctx), sdk.Int{})
	claimParams := k.ClaimParamsForType(ctx, roundResult.ClaimType)
	voteThreshold := claimParams.VoteThreshold
	thresholdVotes := voteThreshold.MulInt64(totalBondedPower).RoundInt()
	votePower := sdk.NewInt(roundResult.VotePower)
	return !votePower.IsZero() && votePower.GT(thresholdVotes), totalBondedPower
}

////////////////////
/// PRE-VOTE
////////////////////

// CreatePrevote sets the prevote for a given validator
func (k Keeper) CreatePrevote(ctx sdk.Context, hash []byte) {
	ctx.KVStore(k.storeKey).Set(types.GetClaimPrevoteKey(hash), hash)
}

// GetPrevote gets the prevote for a given validator
func (k Keeper) GetPrevote(ctx sdk.Context, hash tmbytes.HexBytes) tmbytes.HexBytes {
	return ctx.KVStore(k.storeKey).Get(hash)
}

// DeletePrevote deletes the prevote for a given validator
func (k Keeper) DeletePrevote(ctx sdk.Context, hash tmbytes.HexBytes) {
	ctx.KVStore(k.storeKey).Delete(types.GetClaimPrevoteKey(hash))
}

// HasPrevote gets the prevote for a given hash
func (k Keeper) HasPrevote(ctx sdk.Context, hash tmbytes.HexBytes) bool {
	return ctx.KVStore(k.storeKey).Has(types.GetClaimPrevoteKey(hash))
}

// GetAllPrevotes returns all prevotes
func (k Keeper) GetAllPrevotes(ctx sdk.Context) [][]byte {
	prevotes := [][]byte{}
	k.IteratePrevotes(ctx, func(hash []byte) (stop bool) {
		prevotes = append(prevotes, hash)
		return
	})
	return prevotes
}

// IteratePrevotes iterates over all prevotes in the store
func (k Keeper) IteratePrevotes(ctx sdk.Context, handler func(hash []byte) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.PrevoteKey)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		hash := bytes.TrimPrefix(iter.Key(), types.PrevoteKey)
		if handler(hash) {
			break
		}
	}
}
