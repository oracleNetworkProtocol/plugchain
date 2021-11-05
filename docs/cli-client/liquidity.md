# Liquidity

liquidity模块用于管理你在PLUGCHAN Hub上发行的资产进行交易。

## 可用命令

| Name                                       |Description                    |
| ------------------------------------------ | ------------------------ |
| [create-pool](#plugchaind-tx-liquidity-create-pool)        | Create liquidity pool and deposit coins   |
| [deposit](#plugchaind-tx-liquidity-deposit)             | Deposit coins to a liquidity pool   |
| [swap](#plugchaind-tx-liquidity-swap)        | Swap offer coin with demand coin from the liquidity pool with the given order price  |
| [withdraw](#plugchaind-tx-liquidity-withdraw)                | Withdraw pool coin from the specified liquidity pool  |
| [batch](#plugchaind-q-liquidity-batch)                | Query details of a liquidity pool batch       |
| [deposit](#plugchaind-q-liquidity-deposit)                | Query the deposit messages on the liquidity pool batch       |
| [deposits](#plugchaind-q-liquidity-deposits)                | Query all deposit messages of the liquidity pool batch       |
| [params](#plugchaind-q-liquidity-params)                | Query the values set as liquidity parameters      |
| [pool](#plugchaind-q-liquidity-pool)                | Query details of a liquidity pool   |
| [pools](#plugchaind-q-liquidity-pools)                | Query for all liquidity pools  |
| [swap](#plugchaind-q-liquidity-swap)  |Query for the swap message on the batch of the liquidity pool specified pool-id and msg-index|
| [swaps](#plugchaind-q-liquidity-swaps)                | Query all swap messages in the liquidity pool batch |
| [withdraw](#plugchaind-q-liquidity-withdraw)                | Query the withdraw messages in the liquidity pool batch       |
| [withdraws](#plugchaind-q-liquidity-withdraws)                | Query for all withdraw messages on the liquidity pool batch |


## plugchaind tx liquidity create-pool

Create liquidity pool and deposit coins

```bash
plugchaind tx liquidity create-pool [pool-type] [deposit-coins] [flags]
```

**Flags：**
| Name, shorthand | type  | Required |Description                                      |
| ---------- | ------ | ----  | ------------------------------------------ |
| pool-type     | uint64 | yes   |The id of the liquidity pool-type. The only supported pool type is 1|
| deposit-coins |string|yes|The amount of coins to deposit to the liquidity pool. The number of deposit coins must be 2 in pool type 1.|

### Create liquidity pool and deposit coins

```bash
plugchaind tx liquidity create-pool 1 1000000000uatom,50000000000uusd  --from=<key-name> --chain-id=<chain-id> --fees=<fee>
```


## plugchaind tx liquidity deposit

Deposit coins to a liquidity pool

```bash
plugchaind tx liquidity deposit [pool-id] [deposit-coins] [flags]
```

**Flags：**
| Name, shorthand | type  | Required |Description                                      |
| ---------- | ------ | ----  | ------------------------------------------ |
| pool-id     | uint64 | yes   | The pool id of the liquidity pool|
| deposit-coins  | string |yes  |  The amount of coins to deposit to the liquidity pool  |

### Deposit coins to a liquidity pool

```bash
plugchaind tx liquidity deposit 1 100000000uatom,5000000000uusd  --from=<key-name> --chain-id=<chain-id> --fees=<fee>
```

## plugchaind tx liquidity swap

Swap offer coin with demand coin from the liquidity pool with the given order price

```bash
plugchaind tx liquidity swap [pool-id] [swap-type] [offer-coin] [demand-coin-denom] [order-price] [swap-fee-rate] [flags]
```

**Flags：**
| Name, shorthand | type  | Required |Description                                      |
| ---------- | ------ | ----  | ------------------------------------------ |
| pool-id     | uint64 | yes   | The pool id of the liquidity pool|
| swap-type   | uint64 |yes  | The swap type of the swap message. The only supported swap type is 1 (instant swap). |
| offer-coin   | string |yes  |  The amount of offer coin to swap |
| demand-coin-denom   | string |yes  | The denomination of the coin to exchange with offer coin   |
| order-price   | float |yes  |The limit order price for the swap order. The price is the exchange ratio of X/Y where X is the amount of the first coin and Y is the amount of the second coin when their denoms are sorted alphabetically|
| swap-fee-rate   | float |yes|The swap fee rate to pay for swap that is proportional to swap amount. The swap fee rate must be the value that set as liquidity parameter in the current network.|

### swap

```bash
plugchaind tx liquidity swap 1 1 50000000uusd uatom 0.019 0.003  --from=<key-name> --chain-id=<chain-id> --fees=<fee>
```

## plugchaind tx liquidity withdraw

Withdraw pool coin from the specified liquidity pool

```bash
plugchaind tx liquidity withdraw [pool-id] [pool-coin] [flags]
```

**Flags：**
| Name, shorthand | type  | Required |Description                                      |
| ---------- | ------ | ----  | ------------------------------------------ |
| pool-id     | uint64 | yes   |  The pool id of the liquidity pool|
| pool-coin  | string |yes  |  The amount of pool coin to withdraw from the liquidity pool |

### Withdraw pool coin from the specified liquidity pool

```bash
plugchaind tx liquidity withdraw 1 10000pool96EF6EA6E5AC828ED87E8D07E7AE2A8180570ADD212117B2DA6F0B75D17A6295 --from=<key-name> --chain-id=<chain-id> --fees=<fee>
```

## plugchaind  q liquidity batch

Query details of a liquidity pool batch

```bash
plugchaind query liquidity batch [pool-id] [flags]
```
### Query details of a liquidity pool batch

```bash
plugchaind q liquidity batch 1 
```

## plugchaind q liquidity deposit

Query the deposit messages on the liquidity pool batch

```bash
plugchaind query liquidity deposit [pool-id] [msg-index] [flags]
```

### Query the deposit messages on the liquidity pool batch

```bash
plugchaind query liquidity deposit 1 20
```

## plugchaind q liquidity deposits

Query all deposit messages of the liquidity pool batch

```bash
plugchaind query liquidity deposits [pool-id] [flags]
```

### Query all deposit messages of the liquidity pool batch

```bash
plugchaind query liquidity deposits 1
```

## plugchaind q liquidity params

 Query the values set as liquidity parameters

```bash
plugchaind query liquidity params [flags]
```

## plugchaind q liquidity pool

Query details of a liquidity pool

```bash
plugchaind query liquidity pool [pool-id] [flags]
```

**Flags：**
| Name, shorthand | type  | Required |Description                                      |
| ---------- | ------ | ----  | ------------------------------------------ |
| --pool-coin-denom    | string | no   | The denomination of the pool coin |
| --reserve-acc  | string |no  | The Bech32 address of the reserve account  |

### Query details of a liquidity pool

```bash
plugchaind query liquidity pool 1
```

## plugchaind q liquidity pools

Query for all liquidity pools

```bash
plugchaind query liquidity pools [flags]
```


## plugchaind q liquidity swap

Query for the swap message on the batch of the liquidity pool specified pool-id and msg-index

```bash
plugchaind query liquidity swap [pool-id] [msg-index] [flags]
```

### Query for the swap message on the batch of the liquidity pool specified pool-id and msg-index

```bash
plugchaind query liquidity swap 1 20
```


## plugchaind q liquidity swaps

Query all swap messages in the liquidity pool batch

```bash
plugchaind query liquidity swaps [pool-id] [flags]
```

### Query all swap messages in the liquidity pool batch

```bash
plugchaind query liquidity swaps 1
```


## plugchaind q liquidity withdraw

Query the withdraw messages in the liquidity pool batch

```bash
plugchaind query liquidity withdraw [pool-id] [msg-index] [flags]
```

### Query the withdraw messages in the liquidity pool batch

```bash
plugchaind query liquidity withdraw 1 20
```


## plugchaind q liquidity withdraws

Query for all withdraw messages on the liquidity pool batch

```bash
plugchaind query liquidity withdraws [pool-id] [flags]
```

### Query for all withdraw messages on the liquidity pool batch

```bash
plugchaind query liquidity withdraws 1
```