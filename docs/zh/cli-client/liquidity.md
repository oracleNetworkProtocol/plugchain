# Liquidity

流动性模块通过提供流动性活动和硬币交换功能，为自动做市商 (AMM) 式的去中心化流动性提供服务。

该模块使用户能够创建流动性池，进行存款和取款，并从流动性池中请求硬币掉期。

## 可用命令

| 名称                                       | 描述                     |
| ------------------------------------------ | ------------------------ |
| [create-pool](#plugchaind-tx-liquidity-create-pool)        | 创建流动性池                 |
| [deposit](#plugchaind-tx-liquidity-deposit)             | 将硬币存入流动性池                 |
| [swap](#plugchaind-tx-liquidity-swap)        | 用流动性池中的活期硬币与给定订单价格进行交换           |
| [withdraw](#plugchaind-tx-liquidity-withdraw)                | 从指定的流动性池中提取池硬币       |
| [batch](#plugchaind-q-liquidity-batch)                | 查询流动性池批次的详细信息       |
| [deposit](#plugchaind-q-liquidity-deposit)                | 查询流动性池批次上的存款消息       |
| [deposits](#plugchaind-q-liquidity-deposits)                | 查询流动性池批次的所有存款消息       |
| [params](#plugchaind-q-liquidity-params)                | 查询设置为流动性参数的值       |
| [pool](#plugchaind-q-liquidity-pool)                | 查询流动资金池的详细信息       |
| [pools](#plugchaind-q-liquidity-pools)                | 查询所有流动性池       |
| [swap](#plugchaind-q-liquidity-swap)                | 查询流动性池指定池id和消息索引批次上的掉期消息       |
| [swaps](#plugchaind-q-liquidity-swaps)                | 查询流动性池批处理中的所有掉期消息       |
| [withdraw](#plugchaind-q-liquidity-withdraw)                | 查询流动性池批次中的提款消息       |
| [withdraws](#plugchaind-q-liquidity-withdraws)                | 查询流动性池批次上的所有提款消息       |


## plugchaind tx liquidity create-pool

创建一个新流动性池。

```bash
plugchaind tx liquidity create-pool [pool-type] [deposit-coins] [flags]
```

**标识：**
| 名称,速记 | 类型   | 必须  | 描述                                       |
| ---------- | ------ | ----  | ------------------------------------------ |
| pool-type     | uint64 |  是    |流动性池类型的id。唯一支持的池类型是1(即时交换）|
| deposit-coins  | string | 是   | 存入流动性池的硬币数量。在池类型1中，存款硬币的数量必须为2。|

### 创建新流动性池

```bash
plugchaind tx liquidity create-pool 1 1000000000uatom,50000000000uusd  --from=<key-name> --chain-id=<chain-id> --fees=<fee>
```


## plugchaind tx liquidity deposit

将硬币存入流动资金池。

```bash
plugchaind tx liquidity deposit [pool-id] [deposit-coins] [flags]
```

**标识：**
| 名称,速记 | 类型   | 必须  | 描述                                       |
| ---------- | ------ | ----  | ------------------------------------------ |
| pool-id     | uint64 |  是    |       流动性池的池id |
| deposit-coins  | string | 是   |      存入流动性池的硬币数量          |

### 将硬币存入流动资金池

```bash
plugchaind tx liquidity deposit 1 100000000uatom,5000000000uusd  --from=<key-name> --chain-id=<chain-id> --fees=<fee>
```

## plugchaind tx liquidity swap

用流动性池中的活期硬币与给定订单价格进行交换

```bash
plugchaind tx liquidity swap [pool-id] [swap-type] [offer-coin] [demand-coin-denom] [order-price] [swap-fee-rate] [flags]
```

**标识：**
| 名称,速记 | 类型   | 必须  | 描述                                       |
| ---------- | ------ | ----  | ------------------------------------------ |
| pool-id     | uint64 |  是    |       流动性池的池id |
| swap-type   | uint64 | 是   |      交换消息的交换类型。唯一支持的交换类型是1（即时交换）。            |
| offer-coin   | string | 是   |      要交换的硬币数量                                |
| demand-coin-denom   | string | 是   |      与报价硬币交换的硬币的面额                 |
| order-price   | float | 是   |订单的限制订单价格。价格是X/Y的交换比率，其中X是第一枚硬币的数量，Y是第二枚硬币的数量，当其名称按字母顺序排序时  |
| swap-fee-rate   | float | 是 |与掉期金额成比例的掉期费率。掉期费率必须是当前网络中设置为流动性参数的值。|

### 发起交换

```bash
plugchaind tx liquidity swap 1 1 50000000uusd uatom 0.019 0.003  --from=<key-name> --chain-id=<chain-id> --fees=<fee>
```

## plugchaind tx liquidity withdraw

从指定的流动性池中提取池硬币。

```bash
plugchaind tx liquidity withdraw [pool-id] [pool-coin] [flags]
```

**标识：**
| 名称,速记 | 类型   | 必须  | 描述                                       |
| ---------- | ------ | ----  | ------------------------------------------ |
| pool-id     | uint64 |  是    |       流动性池的池id |
| pool-coin  | string | 是   |    从流动性池中提取的池币金额          |

### 提取池硬币

```bash
plugchaind tx liquidity withdraw 1 10000pool96EF6EA6E5AC828ED87E8D07E7AE2A8180570ADD212117B2DA6F0B75D17A6295 --from=<key-name> --chain-id=<chain-id> --fees=<fee>
```

## plugchaind  q liquidity batch

查询流动性池批次的详细信息

```bash
plugchaind query liquidity batch [pool-id] [flags]
```
### 查询流动性池批次的详细信息

```bash
plugchaind q liquidity batch 1 
```

## plugchaind q liquidity deposit

查询流动性池指定池id和消息索引的批次存款消息

```bash
plugchaind query liquidity deposit [pool-id] [msg-index] [flags]
```

### 查询流动性池批次上指定池id和消息索引的存款消息

```bash
plugchaind query liquidity deposit 1 20
```

## plugchaind q liquidity deposits

查询指定池上流动性池的所有批次存款消息

```bash
plugchaind query liquidity deposits [pool-id] [flags]
```

### 查询指定池上流动性池的所有批次存款消息

```bash
plugchaind query liquidity deposits 1
```

## plugchaind q liquidity params

设置为流动性参数的查询值。

```bash
plugchaind query liquidity params [flags]
```

## plugchaind q liquidity pool

查询流动资金池的详细信息

```bash
plugchaind query liquidity pool [pool-id] [flags]
```

**标识：**
| 名称,速记 | 类型   | 必须  | 描述                                       |
| ---------- | ------ | ----  | ------------------------------------------ |
| --pool-coin-denom    | string |  否    |   币的面额 |
| --reserve-acc  | string | 否   |    储备账户的Bech32地址          |

### 查询流动资金池的详细信息

```bash
plugchaind query liquidity pool 1
```

## plugchaind q liquidity pools

查询网络上所有流动性池的详细信息。

```bash
plugchaind query liquidity pools [flags]
```


## plugchaind q liquidity swap

查询流动性池指定池id和消息索引批次的掉期消息

```bash
plugchaind query liquidity swap [pool-id] [msg-index] [flags]
```

### 查询流动性池指定池id和消息索引批次的掉期消息

```bash
plugchaind query liquidity swap 1 20
```


## plugchaind q liquidity swaps

查询流动性池批次中指定池id的所有掉期消息

```bash
plugchaind query liquidity swaps [pool-id] [flags]
```

### 查询流动性池批次中指定池id的所有掉期消息

```bash
plugchaind query liquidity swaps 1
```


## plugchaind q liquidity withdraw

在流动性池批次中查询指定池id和消息索引的提款消息

```bash
plugchaind query liquidity withdraw [pool-id] [msg-index] [flags]
```

### 在流动性池批次中查询指定池id和消息索引的提款消息

```bash
plugchaind query liquidity withdraw 1 20
```


## plugchaind q liquidity withdraws

查询流动性池批次上指定池id的所有提款消息

```bash
plugchaind query liquidity withdraws [pool-id] [flags]
```

### 查询流动性池批次上指定池id的所有提款消息

```bash
plugchaind query liquidity withdraws 1
```