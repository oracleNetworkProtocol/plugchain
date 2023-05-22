---
order: 13
---

# Get proposal list

- Get all proposal information

### Request method
`GET`

### URL
`/server/v1/gov/proposals`

### Request example

```
/server/v1/gov/proposals
```


#### Request parameters

- Not

### Return value
- `code int64 `: 0 or field does not exist is success, others are failures
- `message string` : Response information
- `proposals array` : Verifier array
    - `proposal_id string `: Address of verifier
    - `content object`: Proposal details
    - `status string`: Proposal status
    - `final_tally_result object`: Final voting result
    - `submit_time string `: Submission time
    - `deposit_end_time string` : Deposit end time
    - `total_deposit object`: Deposits
    - `voting_start_time string`: Voting start time
    - `voting_end_time string`: Voting end time

#### Return to example
```json5
{
  "proposals": [
    {
      "proposal_id": "11",
      "content": {
        "@type": "/cosmos.params.v1beta1.ParameterChangeProposal",
        "title": "Slashing Param Change",
        "description": "Adjust the slash_fraction_downtime from 0.001 to 0.0001",
        "changes": [
          {
            "subspace": "slashing",
            "key": "SlashFractionDowntime",
            "value": "\"0.000100000000000000\""
          }
        ]
      },
      "status": "PROPOSAL_STATUS_VOTING_PERIOD",
      "final_tally_result": {
        "yes": "0",
        "abstain": "0",
        "no": "0",
        "no_with_veto": "0"
      },
      "submit_time": "2023-02-16T02:41:22.830539778Z",
      "deposit_end_time": "2023-02-18T02:41:22.830539778Z",
      "total_deposit": [
        {
          "denom": "uplugcn",
          "amount": "10000000"
        }
      ],
      "voting_start_time": "2023-02-16T02:41:22.830539778Z",
      "voting_end_time": "2023-02-21T02:41:22.830539778Z"
    },
    {
      "proposal_id": "10",
      "content": {
        "@type": "/cosmos.upgrade.v1beta1.SoftwareUpgradeProposal",
        "title": "v1.7",
        "description": "This on-chain upgrade governance proposal is to adopt plugchaind v1.7+ which includes a number of updates and fixes. Adjust the proposal fundraising time to 2 days and the voting time to 5 days. Rename plugcn to pc,dump cosmos-sdk v0.45+ ethermint v0.18+ tendermint v0.34+ ibc-go v3+ . del nft module and add authz module ...",
        "plan": {
          "name": "v1.7",
          "time": "0001-01-01T00:00:00Z",
          "height": "5633000",
          "info": "{\"binaries\":{\"linux/amd64\":\"https://github.com/oracleNetworkProtocol/plugchain/releases/download/v1.7.0/plugchain_1.7.0_Linux_x86_64.tar.gz\",\"linux/arm64\":\"https://github.com/oracleNetworkProtocol/plugchain/releases/download/v1.7.0/plugchain_1.7.0_Linux_arm64.tar.gz\",\"darwin/amd64\":\"https://github.com/oracleNetworkProtocol/plugchain/releases/download/v1.7.0/plugchain_1.7.0_Darwin_x86_64.tar.gz\"}}",
          "upgraded_client_state": null
        }
      },
      "status": "PROPOSAL_STATUS_PASSED",
      "final_tally_result": {
        "yes": "4033515558113959",
        "abstain": "0",
        "no": "0",
        "no_with_veto": "0"
      },
      "submit_time": "2022-08-16T07:55:21.890959838Z",
      "deposit_end_time": "2022-08-19T07:55:21.890959838Z",
      "total_deposit": [
        {
          "denom": "uplugcn",
          "amount": "1000000000"
        }
      ],
      "voting_start_time": "2022-08-16T07:55:21.890959838Z",
      "voting_end_time": "2022-08-27T07:55:21.890959838Z"
    },
    //....
    //....
    //....
  ],
  "pagination": {
    "next_key": null,
    "total": "11"
  }
}
```