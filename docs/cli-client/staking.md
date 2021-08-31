# Staking

Staking module provides a set of subcommands to query staking state and send staking transactions.

## Available Commands

| Name                                                                         | Description                                                                                   |
| ---------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| [validator](#plugchaind-query-staking-validator)                                   | Query a validator                                                                             |
| [validators](#plugchaind-query-staking-validators)                                 | Query for all validators                                                                      |
| [delegation](#plugchaind-query-staking-delegation)                                 | Query a delegation based on address and validator address                                     |
| [delegations](#plugchaind-query-staking-delegations)                               | Query all delegations made from one delegator                                                 |
| [delegations-to](#plugchaind-query-staking-delegations-to)                         | Query all delegations to one validator                                                        |
| [unbonding-delegation](#plugchaind-query-staking-unbonding-delegation)             | Query an unbonding-delegation record based on delegator and validator address                 |
| [unbonding-delegations](#plugchaind-query-staking-unbonding-delegations)           | Query all unbonding-delegations records for one delegator                                     |
| [unbonding-delegations-from](#plugchaind-query-staking-unbonding-delegations-from) | Query all unbonding delegatations from a validator                                            |
| [redelegations-from](#plugchaind-query-staking-redelegations-from)                 | Query all outgoing redelegatations from a validator                                           |
| [redelegation](#plugchaind-query-staking-redelegation)                             | Query a redelegation record based on delegator and a source and destination validator address |
| [redelegations](#plugchaind-query-staking-redelegations)                           | Query all redelegations records for one delegator                                             |
| [pool](#plugchaind-query-staking-pool)                                             | Query the current staking pool values                                                         |
| [params](#plugchaind-query-staking-params)                                         | Query the current staking parameters information                                              |
| [historical-info](#plugchaind-query-staking-historical-info)                       | Query historical info at given height                                                         |
| [create-validator](#plugchaind-tx-staking-create-validator)                        | Create new validator initialized with a self-delegation to it                                 |
| [edit-validator](#plugchaind-tx-staking-edit-validator)                            | Edit existing validator account                                                               |
| [delegate](#plugchaind-tx-staking-delegate)                                        | Delegate liquid tokens to an validator                                                        |
| [unbond](#plugchaind-tx-staking-unbond)                                            | Unbond shares from a validator                                                                |
| [redelegate](#plugchaind-tx-staking-redelegate)                                    | Redelegate illiquid tokens from one validator to another                                      |

## plugchaind query staking validator

### Query a validator by validator address

```bash
plugchaind query staking validator <gxvaloper...>
```

## plugchaind query staking validators

### Query all validators

```bash
plugchaind query staking validators
```

## plugchaind query staking delegation

Query a delegation based on delegator address and validator address.

```bash
plugchaind query staking delegation [delegator-addr] [validator-addr]
```

### Query a delegation

```bash
plugchaind query staking delegation <gx...> <gxvaloper...>
```

Example Output:

```bash
Delegation:
  Delegator:  gx13lcwnxpyn2ea3skzmek64vvnp97jsk8qrcezvm
  Validator:  gxvaloper15grv3xg3ekxh9xrf79zd0w077krgv5xfzzunhs
  Shares:     1.0000000000000000000000000000
  Height:     26
```

## plugchaind query staking delegations

Query all delegations delegated from one delegator.

```bash
plugchaind query staking delegations [delegator-address] [flags]
```

### Query all delegations of a delegator

```bash
plugchaind query staking delegations <gx...>
```

## plugchaind query staking delegations-to

Query all delegations to one validator.

```bash
plugchaind query staking delegations-to [validator-address] [flags]
```

### Query all delegations to one validator

```bash
plugchaind query staking delegations-to <gxvaloper...>
```

Example Output:

```bash
Delegation:
  Delegator:  gx13lcwnxpyn2ea3skzmek64vvnp97jsk8qrcezvm
  Validator:  gxvaloper1yclscskdtqu9rgufgws293wxp3njsesxxlnhmh
  Shares:     100.0000000000000000000000000000
  Height:     0
Delegation:
  Delegator:  gx1td4xnefkthfs6jg469x33shzf578fed6n7k7ua
  Validator:  gxvaloper1yclscskdtqu9rgufgws293wxp3njsesxxlnhmh
  Shares:     1.0000000000000000000000000000
  Height:     26
```

## plugchaind query staking unbonding-delegation

Query an unbonding-delegation record based on delegator and validator address.

```bash
plugchaind query staking unbonding-delegation [delegator-addr] [validator-addr] [flags]
```

### Query an unbonding delegation record

```bash
plugchaind query staking unbonding-delegation <gx...> <gxvaloper...>
```

## plugchaind query staking unbonding-delegations

### Query all unbonding delegations records of a delegator

```bash
plugchaind query staking unbonding-delegations <gx...>
```

## plugchaind query staking unbonding-delegations-from

### Query all unbonding delegations from a validator

```bash
plugchaind query staking unbonding-delegations-from <gxvaloper...>
```

## plugchaind query staking redelegations-from

Query all outgoing redelegations of a validator

```bash
plugchaind query staking redelegations-from [validator-address] [flags]
```

### Query all outgoing redelegatations of a validator

```bash
plugchaind query staking redelegations-from <gxvaloper...>
```

## plugchaind query staking redelegation

Query a redelegation record based on delegator and source validator address and destination validator address.

```bash
plugchaind query staking redelegation [delegator-addr] [src-validator-addr] [dst-validator-addr] [flags]
```

### Query a redelegation record

```bash
plugchaind query staking redelegation <gx...> <gxvaloper...> <gxvaloper...>
```

## plugchaind query staking redelegations

### Query all redelegations records of a delegator

```bash
plugchaind query staking redelegations <gx...>
```

## plugchaind query staking pool

### Query the current staking pool values

```bash
plugchaind query staking pool
```

Example Output:

```bash
Pool:
  Loose Tokens:   1409493892.759816067399143966
  Bonded Tokens:  590526409.65743521209068061
  Token Supply:   2000020302.417251279489824576
  Bonded Ratio:   0.2952602076
```

## plugchaind query staking params

### Query the current staking parameters information

```bash
plugchaind query staking params
```

## plugchaind query staking historical-info

### Query historical info at given height

```bash
plugchaind query staking historical-info <height>
```

## plugchaind tx staking create-validator

Send a transaction to apply to be a validator and delegate a certain amount of plugchaind to it.

```bash
plugchaind tx staking create-validator [flags]
```

**Flags:**

| Name, shorthand              | type   | Required | Default | Description                                                                                      |
| ---------------------------- | ------ | -------- | ------- | ------------------------------------------------------------------------------------------------ |
| --amount                     | string | Yes      |         | Amount of coins to bond                                                                          |
| --commission-rate            | float  | Yes      | 0.0     | The initial commission rate percentage                                                           |
| --commission-max-rate        | float  |          | 0.0     | The maximum commission rate percentage                                                           |
| --commission-max-change-rate | float  |          | 0.0     | The maximum commission change rate percentage (per day)                                          |
| --min-self-delegation        | string |          |         | The minimum self delegation required on the validator                                            |
| --details                    | string |          |         | Optional details                                                                                 |
| --genesis-format             | bool   |          | false   | Export the transaction in gen-tx format; it implies --generate-only                              |
| --identity                   | string |          |         | Optional identity signature (ex. UPort or Keybase)                                               |
| --ip                         | string |          |         | Node's public IP. It takes effect only when used in combination with                             |
| --node-id                    | string |          |         | The node's ID                                                                                    |
| --moniker                    | string | Yes      |         | Validator name                                                                                   |
| --pubkey                     | string | Yes      |         | Go-Amino encoded hex PubKey of the validator. For Ed25519 the go-amino prepend hex is 1624de6220 |
| --website                    | string |          |         | Optional website                                                                                 |
| --security-contact           | string |          |         | The validator's (optional) security contact email                                                |

### Create a validator

```bash
plugchaind tx staking create-validator --chain-id=plugchain --from=<key-name> --fees=20plug --pubkey=<validator-pubKey> --commission-rate=0.1 --amount=100plug --moniker=<validator-name>
```

:::tip
Follow the [Mainnet](../get-started/mainnet.md#create-validator) instructions to learn more.
:::

## plugchaind tx staking edit-validator

Edit an existing validator's settings, such as commission rate, name, etc.

```bash
plugchaind tx staking edit-validator [flags]
```

**Flags:**

| Name, shorthand       | type   | Required | Default | Description                                           |
| --------------------- | ------ | -------- | ------- | ----------------------------------------------------- |
| --commission-rate     | float  |          | 0.0     | Commission rate percentage                            |
| --moniker             | string |          |         | Validator name                                        |
| --identity            | string |          |         | Optional identity signature (ex. UPort or Keybase)    |
| --website             | string |          |         | Optional website                                      |
| --details             | string |          |         | Optional details                                      |
| --security-contact    | string |          |         | The validator's (optional) security contact email     |
| --min-self-delegation | string |          |         | The minimum self delegation required on the validator |

### Edit validator information

```bash
plugchaind tx staking edit-validator --from=<key-name> --chain-id=plugchain --fees=20plug --commission-rate=0.10 --moniker=<validator-name>
```

### Upload validator avatar

Please refer to [How to upload my validator's logo to the Explorers](../concepts/validator-faq.md#how-to-upload-my-validator-s-logo-to-the-explorers)

## plugchaind tx staking delegate

Delegate tokens to a validator.

```bash
plugchaind tx staking delegate [validator-addr] [amount] [flags]
```

```bash
plugchaind tx staking delegate <gxvaloper...> <amount> --chain-id=plugchain --from=<key-name> --fees=20plug
```

## plugchaind tx staking unbond

Unbond tokens from a validator.

```bash
plugchaind tx staking unbond [validator-addr] [amount] [flags]
```

### Unbond some tokens from a validator

```bash
plugchaind tx staking unbond <gxvaloper...> 10plugchaind --from=<key-name> --chain-id=plugchain --fees=20plug
```

## plugchaind tx staking redelegate

Transfer delegation from one validator to another.

:::tip
There is no `unbonding time` during the redelegation, so you will not miss the rewards. But you can only redelegate once per validator, until a period (= `unbonding time`) exceed.
:::

```bash
plugchaind tx staking redelegate [src-validator-addr] [dst-validator-addr] [amount] [flags]
```

### Redelegate some tokens to another validator

```bash
plugchaind tx staking redelegate <gxvaloper...> <gxvaloper...> 10plugchaind --chain-id=plugchain --from=<key-name> --fees=20plug
```
