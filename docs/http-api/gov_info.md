---
order: 12
---
# Get proposals

- Get proposal information through proposal ID

### Request method
`GET`

### URL
`/server/v1/gov/detail`

### Request example

```
/server/v1/gov/detail?proposal_id=11
```


#### Request parameters

- Not

### Return value
- `code int64 `: 0 or field does not exist is success, others are failures
- `message string` : Response information
- `proposals array` : Verifier array
  - `proposal_id string `: Address of verifier
  - `content object`: Proposal details
    - `@type`
    - `title`
    - `description`
    - `changes`
  - `status string`: Proposal status
  - `final_tally_result object`: Final voting result
    - `yes`:Number of affirmative votes
    - `abstain`: Number of abstained votes
    - `no`: Number of negative votes
    - `no_with_veto`: Number of negative votes to exercise the veto
  - `submit_time string `: Submission time
  - `deposit_end_time string` : Deposit end time
  - `total_deposit object`: Deposits
  - `voting_start_time string`: Voting start time
  - `voting_end_time string`: Voting end time

#### Return to example
```json5
{
  "proposal": {
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
  }
}
```