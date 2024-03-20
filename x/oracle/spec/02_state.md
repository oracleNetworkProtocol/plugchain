<!--
order: 2
-->

# State

## Claim

A `Claim` is an abstract type that the result of an off-chain worker process must implement. A `Claim` must have a `Type`. You must specify a list of `ClaimTypes` and their associate the module must proceess via the oracle module's `ClaimParams` param map.

- `0x00 | claimHash -> ProtocolBuffer(claim)`

```go
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
```

## Round

A round contains all the validator votes for the given round and claim type.

- `0x01 | claimType | roundID -> round`

```go
// Round contains all claim votes for a given round
message Round {
  uint64 roundId = 1;
  string type = 2; // namespace so we can have multiple oracles
  repeated Vote votes = 3 [(gogoproto.nullable) = false];
}

// Vote is a vote for a given claim by a validator
message Vote {
  uint64 roundId = 1; // use int in case we need to enforce sequential ordering
  bytes claimHash = 2 [(gogoproto.casttype) = "github.com/tendermint/tendermint/libs/bytes.HexBytes"];
  string consensusId = 3; //  used to compare claims when tallying votes
  string type = 4;
  bytes validator = 5 [(gogoproto.casttype) = "github.com/cosmos/cosmos-sdk/types.ValAddress"];
}
```

## PendingRounds

An array of rounds that have not yet reached consensus

- `0x02 | claimType | roundId -> roundID`
- `0x02 | claimType -> []roundID`

## Prevote

A hash of the prevote

- `0x03 | prevote_hash -> prevote_hash`

## Delegate

Delegation mapping

- `0x04 | validator_address -> delegate_address`
- `0x05 | delegate_address -> validator_address`

## LastFinalizedRound

The most recent round that has reached consensus. Votes with RoundID lower than LastFinalizedRound are not processed

- `0x06 | claimType -> roundID`
