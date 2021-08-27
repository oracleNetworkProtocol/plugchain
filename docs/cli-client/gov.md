# Gov

This module provides the basic functionalities for [Governance](../features/governance.md).

## Available Commands

| Name                                            | Description                                                       |
| ----------------------------------------------- | ----------------------------------------------------------------- |
| [proposal](#plugchaind-query-gov-proposal)            | Query details of a single proposal                                |
| [proposals](#plugchaind-query-gov-proposals)          | Query proposals with optional filter                              |
| [vote](#plugchaind-query-gov-vote)                    | Query details of a single vote                                    |
| [votes](#plugchaind-query-gov-votes)                  | Query votes on a proposal                                         |
| [deposit](#plugchaind-query-gov-deposit)              | Query details of a deposit                                        |
| [deposits](#plugchaind-query-gov-deposits)            | Query deposits on a proposal                                      |
| [tally](#plugchaind-query-gov-tally)                  | Get the tally of a proposal vote                                  |
| [param](#plugchaind-query-gov-param)                  | Query the parameters (voting                                      |
| [params](#plugchaind-query-gov-params)                | Query the parameters of the governance process                    |
| [proposer](#plugchaind-query-gov-proposer)            | Query which address proposed a proposal with a given ID.          |
| [submit-proposal](#plugchaind-tx-gov-submit-proposal) | Submit a proposal along with an initial deposit                   |
| [deposit](#plugchaind-tx-gov-deposit)                 | Deposit tokens for an active proposal                             |
| [vote](#plugchaind-tx-gov-vote)                       | Vote for an active proposal, options: yes/no/no_with_veto/abstain |

## plugchaind query gov proposal

Query details of a proposal.

```bash
plugchaind query gov proposal [proposal-id] [flags]
```

### Query a proposal

```bash
plugchaind query gov proposal <proposal-id>
```

## plugchaind query gov proposals

Query proposals with optional filter.

```bash
plugchaind query gov proposals [flags]
```

**Flags:**

| Name, shorthand | Type    | Required | Default | Description                                                         |
| --------------- | ------- | -------- | ------- | ------------------------------------------------------------------- |
| --depositor     | Address |          |         | Filter proposals by depositor address                               |
| --limit         | uint    |          |         | Limit to the latest [number] of proposals. Default to all proposals |
| --status        | string  |          |         | Filter proposals by status                                          |
| --voter         | Address |          |         | Filter proposals by voter address                                   |

### Query all proposals

```bash
plugchaind query gov proposals
```

### Query proposals by conditions

```bash
plugchaind query gov proposals --limit=3 --status=Passed --depositor=<gx...>
```

## plugchaind query gov vote

Query details of a single vote.

```bash
plugchaind query gov vote [proposal-id] [voter-addr] [flags]
```

### Query a vote

```bash
plugchaind query gov vote <proposal-id> <gx...>
```

## plugchaind query gov votes

Query votes on a proposal.

```bash
plugchaind query gov votes [proposal-id] [flags]
```

### Query all votes of a proposal

```bash
plugchaind query gov votes <proposal-id>
```

## plugchaind query gov deposit

Query details for a single proposal deposit on a proposal by its identifier.

```bash
plugchaind query gov deposit [proposal-id] [depositer-addr] [flags]
```

### Query a deposit of a proposal

```bash
plugchaind query gov deposit <proposal-id> <gx...>
```

## plugchaind query gov deposits

Query details for all deposits on a proposal.

```bash
plugchaind query gov deposits [proposal-id] [flags]
```

### Query all deposits of a proposal

```bash
plugchaind query gov deposits <proposal-id>
```

## plugchaind query gov tally

Query tally of votes on a proposal. You can find the proposal-id by running "plugchaind query gov proposals".

```bash
plugchaind query gov tally [proposal-id] [flags]
```

### Query the statistics of a proposal

```bash
plugchaind query gov tally <proposal-id>
```

## plugchaind query gov param

Query the parameters (voting|tallying|deposit) of the governance process.

```bash
plugchaind query gov param [param-type] [flags]
```

Example:

```bash
> plugchaind query gov param voting
> plugchaind query gov param tallying
> plugchaind query gov param deposit
```

## plugchaind query gov params

Query the all the parameters for the governance process.

```bash
plugchaind query gov params [flags]
```

## plugchaind query gov proposer

Query which address proposed a proposal with a given ID.

```bash
plugchaind query gov proposer [proposal-id] [flags]
```

## plugchaind tx gov submit-proposal

Submit a proposal along with an initial deposit. Proposal title, description, type and deposit can be given directly or through a proposal JSON file.
Available Commands:  `community-pool-spend`, `param-change`, `software-upgrade`, `cancel-software-upgrade` .

### plugchaind tx gov submit-proposal community-pool-spend

Submit a community pool spend proposal along with an initial deposit.
The proposal details must be supplied via a JSON file.

```bash
plugchaind tx gov submit-proposal community-pool-spend <path/to/proposal.json> --from=<key_or_address>
```

Where proposal.json contains:

```json
{
    "title": "Community Pool Spend",
    "description": "Pay me some Atoms!",
    "recipient": "gx1mjk4p68mmulwla3x5uzlgjwsc3zrms448rel3q",
    "amount": "1000plug",
    "deposit": "1000plug"
}
```

### plugchaind tx gov submit-proposal param-change

Submit a parameter proposal along with an initial deposit.
The proposal details must be supplied via a JSON file. For values that contains objects, only non-empty fields will be updated.

IMPORTANT: Currently parameter changes are evaluated but not validated, so it is very important that any "value" change is valid (ie. correct type and within bounds) for its respective parameter, eg. "MaxValidators" should be an integer and not a decimal.

Proper vetting of a parameter change proposal should prevent this from happening (no deposits should occur during the governance process), but it should be noted regardless.

```bash
plugchaind tx gov submit-proposal param-change <path/to/proposal.json> --from=<key_or_address>
```

Where proposal.json contains:

```json
{
    "title": "Staking Param Change",
    "description": "Update max validators",
    "changes": [
        {
        "subspace": "staking",
        "key": "MaxValidators",
        "value": 105
        }
    ],
    "deposit": "1000plug"
}
```

### plugchaind tx gov submit-proposal software-upgrade

Submit a software upgrade along with an initial deposit.
Please specify a unique name and height OR time for the upgrade to take effect.

```bash
plugchaind tx gov submit-proposal software-upgrade [name] (--upgrade-height [height] | --upgrade-time [time]) (--upgrade-info [info]) [flags]
```

**Flags:**

| Name, shorthand  | Type   | Required | Default | Description                                                                               |
| ---------------- | ------ | -------- | ------- | ----------------------------------------------------------------------------------------- |
| --deposit        | Coin   | Yes      |         | Deposit of the proposal                                                                   |
| --title          | string | Yes      |         | Title of proposal                                                                         |
| --description    | string | Yes      |         | Description of proposal                                                                   |
| --upgrade-height | int64  |          |         | The height at which the upgrade must happen (not to be used together with --upgrade-time) |
| --time           | string |          |         | The time at which the upgrade must happen (not to be used together with --upgrade-height) |
| --info           | string |          |         | Optional info for the planned upgrade such as commit hash, etc.                           |

### plugchaind tx gov submit-proposal cancel-software-upgrade

Cancel a software upgrade along with an initial deposit.

```bash
plugchaind tx gov submit-proposal cancel-software-upgrade [flags]
```

**Flags:**

| Name, shorthand | Type   | Required | Default | Description             |
| --------------- | ------ | -------- | ------- | ----------------------- |
| --deposit       | Coin   | Yes      |         | Deposit of the proposal |
| --title         | string | Yes      |         | Title of proposal       |
| --description   | string | Yes      |         | Description of proposal |

## plugchaind tx gov deposit

Submit a deposit for an active proposal. You can find the proposal-id by running "plugchaind query gov proposals".

```bash
plugchaind tx gov deposit [proposal-id] [deposit] [flags]
```

### Deposit for an active proposal

```bash
plugchaind tx gov deposit [proposal-id] [deposit]
```

## plugchaind tx gov vote

Submit a vote for an active proposal. You can find the proposal-id by running "plugchaind query gov proposals".
Vote for an active proposal, options: yes/no/no_with_veto/abstain.

```bash
plugchaind tx gov vote [proposal-id] [option] [flags]
```

### Vote for an active proposal

```bash
plugchaind tx gov vote <proposal-id> <option> --from=<key-name> --fees=0.3plug
```
