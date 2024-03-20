<!--
order: 5
-->

# Parameters

Each `claimType` your application needs to process, must be registered under in the Oracle Module params.
`ClaimParams` is a map that contains the parameters for each `claimType`.

```protobuf
// Params represents the parameters used by the oracle module.
message Params {
  map<string, ClaimParams> claim_params  = 1 [(gogoproto.nullable) = false, (gogoproto.moretags) = "yaml:\"claim_params\""];
}

// ClaimParams is the parameters set for each oracle claim type
message ClaimParams {
  uint64 vote_period = 1 [(gogoproto.moretags) = "yaml:\"vote_period\""];
  string claim_type = 2  [(gogoproto.moretags) = "yaml:\"claim_type\""];
  bool prevote = 3  [(gogoproto.moretags) = "yaml:\"prevote\""];
  bytes vote_threshold = 2 [
    (gogoproto.moretags)   = "yaml:\"vote_threshold\"",
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec",
    (gogoproto.nullable)   = false
  ];
}

```

Each `OracleType` supports the following parameters:

| Key            | Type          | Example                | Description                       |
| -------------- | ------------- | ---------------------- | --------------------------------- |
| vote_period    | string (int)  | "5"                    | duration of voting round          |
| vote_threshold | string (dec)  | "0.500000000000000000" | threshold necessary for consensus |
| prevote        | boolean (dec) | true                   | requires prevote round            |
| claim_type     | string        | "myOracleClaimType"    | claim type                        |
