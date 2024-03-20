package types

import (
	"crypto/sha256"
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/oracleNetworkProtocol/plugchain/x/oracle/exported"
)

// NewVote creates a MsgVote instance
func NewVote(roundID uint64, claim exported.Claim, validator sdk.ValAddress, claimType string) *Vote {
	return &Vote{
		RoundId:     roundID,
		ClaimHash:   claim.Hash(),
		ConsensusId: claim.GetConcensusKey(),
		Validator:   validator,
		ClaimType:   claimType,
	}
}

// VoteHash returns the hash for a precommit given the proper args
func VoteHash(salt string, claimHash string, signer sdk.AccAddress) []byte {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%s:%s:%s", salt, claimHash, signer.String())))
	return h.Sum(nil)
}
