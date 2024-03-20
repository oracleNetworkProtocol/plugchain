package types

import (
	fmt "fmt"

	"github.com/tendermint/tendermint/crypto/tmhash"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
)

// NewTestClaim creates a new TestClaim object
func NewTestClaim(blockHeight int64, content string, claimType string) *TestClaim {
	return &TestClaim{
		BlockHeight: blockHeight,
		Content:     content,
		ClaimType:   claimType,
	}
}

// Claim types needed for oracle

// ValidateBasic performs basic validation of the claim
func (c *TestClaim) ValidateBasic() error {
	if len(c.ClaimType) == 0 {
		return fmt.Errorf("claim type should not be empty")
	}
	if len(c.Content) == 0 {
		return fmt.Errorf("claim content should not be empty")
	}
	if c.BlockHeight < 1 {
		return fmt.Errorf("invalid claim height: %d", c.BlockHeight)
	}
	return nil
}

// Type return type of oracle Claim
func (c *TestClaim) Type() string {
	return c.ClaimType
}

// Hash returns the hash of an Claim Content.
func (c *TestClaim) Hash() tmbytes.HexBytes {
	bz, err := c.Marshal()
	if err != nil {
		panic(err)
	}
	return tmhash.Sum(bz)
}

// GetRoundID returns the block height for when the data was used.
func (c *TestClaim) GetRoundID() uint64 {
	return uint64(c.BlockHeight)
}

// GetConcensusKey returns a key the oracle will use of vote consensus
// for deterministic results it should be the same as the hash of the content
// for nondeterministic content it should be a constant
func (c *TestClaim) GetConcensusKey() string {
	return c.Hash().String()
}
