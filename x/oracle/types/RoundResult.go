package types

import tmbytes "github.com/tendermint/tendermint/libs/bytes"

// ClaimVoteResult is a record of votes for each claim
type ClaimVoteResult struct {
	ClaimHash tmbytes.HexBytes
	VotePower int64
}

// RoundResult is is a record of vote tallies for a given round
type RoundResult struct {
	VotePower  int64
	TotalPower int64
	ClaimType  string
	Claims     []*ClaimVoteResult
}

// UpsertClaim upserts a claim
func (r *RoundResult) UpsertClaim(claimHash tmbytes.HexBytes, votePower int64) {
	var existingClaim *ClaimVoteResult
	for _, claim := range r.Claims {
		if claim.ClaimHash.String() == claimHash.String() {
			existingClaim = claim
		}
	}

	if existingClaim != nil {
		existingClaim.VotePower += votePower
		return
	}
	newClaim := newClaimVoteResult(claimHash, votePower)
	r.Claims = append(r.Claims, newClaim)
	return
}

func newClaimVoteResult(claimHash tmbytes.HexBytes, votePower int64) *ClaimVoteResult {
	return &ClaimVoteResult{
		ClaimHash: claimHash,
		VotePower: votePower,
	}
}
