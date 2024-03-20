package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto/tmhash"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
)

// AtomUsd creates a new AtomUsd claim
func NewAtomUsd(blockHeight int64, price string) *AtomUsd {
	decPrice, err := sdk.NewDecFromStr(price)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &AtomUsd{
		Price:       decPrice,
		BlockHeight: blockHeight,
	}
}

// Claim methods that implement the abstract Claim interface of the oracle module

// ValidateBasic performs basic validation of the claim
func (c *AtomUsd) ValidateBasic() error {
	if c.BlockHeight < 1 {
		return fmt.Errorf("invalid claim height: %d", c.BlockHeight)
	}
	return nil
}

// Type return type of oracle Claim
func (c *AtomUsd) Type() string {
	return AtomClaim
}

// Hash returns the hash of an Claim Content.
func (c *AtomUsd) Hash() tmbytes.HexBytes {
	bz, err := c.Marshal()
	if err != nil {
		panic(err)
	}
	return tmhash.Sum(bz)
}

// GetRoundID returns the block height for when the data was used.
func (c *AtomUsd) GetRoundID() uint64 {
	return uint64(c.BlockHeight)
}

// GetConcensusKey returns a key the oracle will use of vote consensus
// for deterministic results it should be the same as the hash of the content
// for nondeterministic content it should be a constant
func (c *AtomUsd) GetConcensusKey() string {
	return ""
}
