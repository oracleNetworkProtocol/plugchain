<!--
order: 3
-->

# Messages

## MsgDelegate

Validators may also elect to delegate voting rights to another key as to not require the validator operator key to be kept online. To do so, validators must submit a `MsgDelegateFeedConsent`, delegating their oracle voting rights to a `Delegate` that can sign `MsgExchangeRatePrevote` and `MsgExchangeRateVote` on behalf of the validator.

The `Validator` field contains the operator address of the validator (prefixed `cosmos1...`). The `Delegate` field is the account address (prefixed `cosmos1...`) of the delegate account that will be submitting exchange rate related votes and prevotes on behalf of the `Operator`.

```protobuf
// MsgDelegate - sdk.Msg for delegating oracle voting rights from a validator
// to another address, must be signed by an active validator
message MsgDelegate {
    string delegate  = 1;
    string validator = 2;
}
```

## MsgPrevote

`hashe` is the SHA256 hash of a string of the format `{salt}:{claim_hash}:{signer}`. This is a commitment to the actual `MsgVote` a validator will submit in the subsequent `VotePeriod`. You can use the `oracletypes.DataHash(salt string, jsn string, signer sdk.AccAddress)` function to help encode it. Note that since in the subsequent `MsgVote`, the salt for the prevote will have to be revealed. Salt used must be regenerated for each prevote submission.

```protobuf
// MsgPrevote - sdk.Msg for prevoting on an array of oracle data types.
// The purpose of the prevote is to hide vote for data with hashes formatted as hex string:
// SHA256("{salt}:{claim_hash}:{voter}")
message MsgPrevote {
  bytes  hash = 1;
  string signer = 2;
}
```

## MsgVote

The `MsgVote` contains a concrete `Claim`. The `salt` parameter must match the salt used to create the prevote, otherwise the vote cannot be processed.

```protobuf
// MsgVote represents a message that supports submitting a vote for
// an arbitrary oracle Claim.
message MsgVote {
  option (gogoproto.equal)           = false;
  option (gogoproto.goproto_getters) = false;

  string              signer = 1;
  google.protobuf.Any claim  = 2 [(cosmos_proto.accepts_interface) = "Claim"];
  string              salt = 3;
}
```
