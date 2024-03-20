// Package exported defines a Claim interface that all oracle claims must implement
package exported

import (
	"github.com/gogo/protobuf/proto"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Claim defines the methods that all oracle claims must implement
type Claim interface {
	proto.Message

	Type() string
	Hash() tmbytes.HexBytes
	GetRoundID() uint64
	// GetConcensusKey returns a key the oracle will use fore voting consensus
	// for deterministic results it should be the the hash of the content
	// (this is because we expect all validators to submit claims with the same exact content)
	// however for nondeterministic content it should be a constant because voting doesn't
	// depend on the content of the claim (developers will need to use the results of the voting
	// to reconcile the various claims)
	GetConcensusKey() string
	ValidateBasic() error
}

// MsgVoteI defines the specific interface a concrete message must
// implement in order to process an oracle claim. The concrete MsgSubmitClaim
// must be defined at the application-level.
type MsgVoteI interface {
	sdk.Msg

	GetClaim() Claim
	GetSigner() sdk.AccAddress
}
