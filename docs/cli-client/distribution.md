# Distribution

The distribution module allows you to manage your [Staking Rewards](../concepts/general-concepts.md#staking-rewards).

## Available Subcommands

| Name                                                                                    | Description                                                                                                                                           |
| --------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- |
| [commission](#plugchaind-query-distribution-commission)                                       | Query distribution validator commission                                                                                                               |
| [community-pool](#plugchaind-query-distribution-community-pool)                               | Query the amount of coins in the community pool                                                                                                       |
| [params](#plugchaind-query-distribution-params)                                               | Query distribution params                                                                                                                             |
| [rewards](#plugchaind-query-distribution-rewards)                                             | Query all distribution delegator rewards or rewards from a particular validator                                                                       |
| [slashes](#plugchaind-query-distribution-slashes)                                             | Query distribution validator slashes.                                                                                                                 |
| [validator-outstanding-rewards](#plugchaind-query-distribution-validator-outstanding-rewards) | Query distribution outstanding (un-withdrawn) rewards for a validator and all their delegations                                                       |
| [fund-community-pool](#plugchaind-tx-distribution-fund-community-pool)                        | Funds the community pool with the specified amount                                                                                                    |
| [set-withdraw-addr](#plugchaind-tx-distribution-set-withdraw-addr)                            | Set the withdraw address for rewards associated with a delegator address                                                                              |
| [withdraw-all-rewards](#plugchaind-tx-distribution-withdraw-all-rewards)                      | Withdraw all rewards for a single delegator                                                                                                           |
| [withdraw-rewards](#plugchaind-tx-distribution-withdraw-rewards)                              | Withdraw rewards from a given delegation address,and optionally withdraw validator commission if the delegation address given is a validator operator |

## plugchaind query distribution commission

Query validator commission rewards from delegators to that validator.

```bash
plugchaind query distribution commission [validator] [flags]
```

## plugchaind query distribution community-pool

Query all coins in the community pool which is under Governance control.

```bash
plugchaind query distribution community-pool [flags]
```

## plugchaind query distribution params

Query distribution params.

```bash
 plugchaind query distribution params [flags]
```

## plugchaind query distribution rewards

Query all rewards earned by a delegator, optionally restrict to rewards from a single validator.

```bash
plugchaind query distribution rewards [delegator-addr] [validator-addr] [flags]
```

## plugchaind query distribution slashes

Query all slashes of a validator for a given block range.

```bash
plugchaind query distribution slashes [validator] [start-height] [end-height] [flags]
```

## plugchaind query distribution validator-outstanding-rewards

Query distribution outstanding (un-withdrawn) rewards for a validator and all their delegations.

```bash
plugchaind query distribution validator-outstanding-rewards [validator] [flags]
```

## plugchaind tx distribution fund-community-pool

Funds the community pool with the specified amount.

```bash
plugchaind tx distribution fund-community-pool [amount] [flags]
```

## plugchaind tx distribution set-withdraw-addr

Set the withdraw address for rewards associated with a delegator address.

```bash
plugchaind tx distribution set-withdraw-addr [withdraw-addr] [flags]
```

## plugchaind tx distribution withdraw-all-rewards

Withdraw all rewards for a single delegator.

```bash
plugchaind tx distribution withdraw-all-rewards [flags]
```

## plugchaind tx distribution withdraw-rewards

Withdraw rewards from a given delegation address, and optionally withdraw validator commission if the delegation address given is a validator operator.

```bash
plugchaind tx distribution withdraw-rewards [validator-addr] [flags]
```
