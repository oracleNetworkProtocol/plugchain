# Staking

Staking模块提供了一系列查询staking状态和发送staking交易的命令。

## 可用命令

| 名称                                                                         | 描述                                                         |
| ---------------------------------------------------------------------------- | ------------------------------------------------------------ |
| [validator](#plugchaind-query-staking-validator)                                   | 查询某个验证者                                               |
| [validators](#plugchaind-query-staking-validators)                                 | 查询所有的验证者                                             |
| [delegation](#plugchaind-query-staking-delegation)                                 | 基于委托者地址和验证者地址的委托查询                         |
| [delegations](#plugchaind-query-staking-delegations)                               | 基于委托者地址的所有委托查询                                 |
| [delegations-to](#plugchaind-query-staking-delegations-to)                         | 查询在某个验证人上的所有委托                                 |
| [unbonding-delegation](#plugchaind-query-staking-unbonding-delegation)             | 基于委托者地址和验证者地址的unbonding-delegation记录查询     |
| [unbonding-delegations](#plugchaind-query-staking-unbonding-delegations)           | 基于委托者地址的所有unbonding-delegation记录查询             |
| [unbonding-delegations-from](#plugchaind-query-staking-unbonding-delegations-from) | 基于验证者地址的所有unbonding-delegation记录查询             |
| [redelegations-from](#plugchaind-query-staking-redelegations-from)                 | 基于某一验证者的所有转委托查询                               |
| [redelegation](#plugchaind-query-staking-redelegation)                             | 基于委托者地址，原验证者地址和目标验证者地址的转委托记录查询 |
| [redelegations](#plugchaind-query-staking-redelegations)                           | 基于委托者地址的所有转委托记录查询                           |
| [pool](#plugchaind-query-staking-pool)                                             | 查询最新的权益池                                             |
| [params](#plugchaind-query-staking-params)                                         | 查询最新的权益参数信息                                       |
| [historical-info](#plugchaind-query-staking-historical-info)                       | 查询给定高度的历史信息                                       |
| [create-validator](#plugchaind-tx-staking-create-validator)                        | 以自委托的方式创建一个新的验证者                             |
| [edit-validator](#plugchaind-tx-staking-edit-validator)                            | 编辑已存在的验证者信息                                       |
| [delegate](#plugchaind-tx-staking-delegate)                                        | 委托一定代币到某个验证者                                     |
| [unbond](#plugchaind-tx-staking-unbond)                                            | 从指定的验证者解绑一定的股份                                 |
| [redelegate](#plugchaind-tx-staking-redelegate)                                    | 转委托一定的token从一个验证者到另一个验证者                  |

## plugchaind query staking validator

### 通过地址查询验证人

```bash
plugchaind query staking validator <gxvaloper...>
```

## plugchaind query staking validators

### 查询所有验证人

```bash
plugchaind query staking validators
```

## plugchaind query staking delegation

通过委托人地址和验证人地址查询委托交易。

```bash
plugchaind query staking delegation [delegator-addr] [validator-addr]
```

### 查询委托交易

```bash
plugchaind query staking delegation <gx...> <gxvaloper...>
```

示例输出：

```bash
Delegation:
  Delegator:  gx13lcwnxpyn2ea3skzmek64vvnp97jsk8qrcezvm
  Validator:  gxvaloper15grv3xg3ekxh9xrf79zd0w077krgv5xfzzunhs
  Shares:     1.0000000000000000000000000000
  Height:     26
```

## plugchaind query staking delegations

查询某个委托人发起的所有委托记录。

```bash
plugchaind query staking delegations [delegator-address] [flags]
```

### 查询某个委托人发起的所有委托记录

```bash
plugchaind query staking delegations <gx...>
```

## plugchaind query staking delegations-to

查询某个验证人接受的所有委托。

```bash
plugchaind query staking delegations-to [validator-address] [flags]
```

### 查询某个验证人接受的所有委托

```bash
plugchaind query staking delegations-to <gxvaloper...>
```

示例输出：

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

通过委托人与验证人地址查询unbonding-delegation记录。

```bash
plugchaind query staking unbonding-delegation [delegator-addr] [validator-addr] [flags]
```

### 查询unbonding-delegation记录

```bash
plugchaind query staking unbonding-delegation <gx...> <gxvaloper...>
```

## plugchaind query staking unbonding-delegations

### 查询委托人的所有未绑定委托记录

```bash
plugchaind query staking unbonding-delegations <gx...>
```

## plugchaind query staking unbonding-delegations-from

### 查询验证人的所有未绑定委托记录

```bash
plugchaind query staking unbonding-delegations-from <gxvaloper...>
```

## plugchaind query staking redelegations-from

查询验证人的所有转委托记录。

```bash
plugchaind query staking redelegations-from [validator-address] [flags]
```

### 查询验证人的所有转委托记录

```bash
plugchaind query staking redelegations-from <gxvaloper...>
```

## plugchaind query staking redelegation

通过委托人地址、原验证人地址、目标验证人地址查询转委托记录。

```bash
plugchaind query staking redelegation [delegator-addr] [src-validator-addr] [dst-validator-addr] [flags]
```

### 查询转委托记录

```bash
plugchaind query staking redelegation <gx...> <gxvaloper...> <gxvaloper...>
```

## plugchaind query staking redelegations

### 查询委托人的所有转委托记录

```bash
plugchaind query staking redelegations <gx...>
```

## plugchaind query staking pool

### 查询当前权益池

```bash
plugchaind query staking pool
```

示例输出：

```bash
Pool:
  Loose Tokens:   1409493892.759816067399143966
  Bonded Tokens:  590526409.65743521209068061
  Token Supply:   2000020302.417251279489824576
  Bonded Ratio:   0.2952602076
```

## plugchaind query staking params

### 查询当前权益参数信息

```bash
plugchaind query staking params
```

## plugchaind query staking historical-info

### 查询给定高度的历史信息

```bash
plugchaind query staking historical-info <height>
```

## plugchaind tx staking create-validator

发送交易申请成为验证人，并委托一定数量的plug到该验证人。

```bash
plugchaind tx staking create-validator [flags]
```

**标志：**

| 名称，速记                   | 类型   | 必须 | 默认  | 描述                                |
| ---------------------------- | ------ | ---- | ----- | ----------------------------------- |
| --amount                     | string | 是   |       | 委托金额                            |
| --commission-rate            | float  | 是   | 0.0   | 初始佣金比例                        |
| --commission-max-rate        | float  |      | 0.0   | 最大佣金比例                        |
| --commission-max-change-rate | float  |      | 0.0   | 最高佣金变更率百分比（每天）        |
| --min-self-delegation        | string |      |       | 验证人要求的最小抵押                |
| --details                    | string |      |       | 验证人节点的详细信息                |
| --genesis-format             | bool   |      | false | 是否以genesis transaction的方式导出 |
| --identity                   | string |      |       | 身份信息的签名                      |
| --ip                         | string |      |       | 验证人节点的IP                      |
| --node-id                    | string |      |       | 节点ID                              |
| --moniker                    | string | 是   |       | 验证人节点名称                      |
| --pubkey                     | string | 是   |       | Amino编码的验证人公钥               |
| --website                    | string |      |       | 验证人节点的网址                    |
| --security-contact           | string |      |       | 验证人（可选）的安全联系电子邮件    |

### 创建验证人

```bash
plugchaind tx staking create-validator --chain-id=plugchain --from=<key-name> --fees=0.3plug --pubkey=<validator-pubKey> --commission-rate=0.1 --amount=100plug --moniker=<validator-name>
```

:::tip
查看 [主网](../get-started/mainnet.md#升级为验证人节点) 说明以了解更多。
:::

## plugchaind tx staking edit-validator

修改验证的的参数，包括佣金比率，验证人节点名称以及其他描述信息。

```bash
plugchaind tx staking edit-validator [flags]
```

**标志：**

| 名称，速记            | 类型   | 必须 | 默认 | 描述                             |
| --------------------- | ------ | ---- | ---- | -------------------------------- |
| --commission-rate     | float  |      | 0.0  | 佣金比率                         |
| --moniker             | string |      |      | 验证人名称                       |
| --identity            | string |      |      | 身份签名                         |
| --website             | string |      |      | 网址                             |
| --details             | string |      |      | 验证人节点详细信息               |
| --security-contact    | string |      |      | 验证人（可选）的安全联系电子邮件 |
| --min-self-delegation | string |      |      | 验证人要求的最小抵押             |

### 编辑验证人信息

```bash
plugchaind tx staking edit-validator --from=<key-name> --chain-id=plugchain --fees=0.3plug --commission-rate=0.10 --moniker=<validator-name>
```

### 上传验证人头像

请参考 [如何将验证人的Logo上传到区块浏览器](../concepts/validator-faq.md#如何将验证人的logo上传到区块浏览器)。

## plugchaind tx staking delegate

向验证人委托通证。

```bash
plugchaind tx staking delegate [validator-addr] [amount] [flags]
```

```bash
plugchaind tx staking delegate <gxvaloper...> <amount> --chain-id=plugchain --from=<key-name> --fees=0.3plug
```

## plugchaind tx staking unbond

从验证人解委托通证。

```bash
plugchaind tx staking unbond [validator-addr] [amount] [flags]
```

### 从验证人中解委托一定数量的代币

```bash
plugchaind tx staking unbond <gxvaloper...> 10plugchaind --from=<key-name> --chain-id=plugchain --fees=0.3plug
```

## plugchaind tx staking redelegate

把某个委托的一部分或者全部从一个验证人转移到另外一个验证人。

:::tip
转委托没有`unbonding time`，所以你不会错过奖励。但是对每个验证人的转委托，在一个`unbonding time`区间内只能操作一次。
:::

```bash
plugchaind tx staking redelegate [src-validator-addr] [dst-validator-addr] [amount] [flags]
```

### 转委托一定数量代币到其他验证人

```bash
plugchaind tx staking redelegate <gxvaloper...> <gxvaloper...> 10plugchaind --chain-id=plugchain --from=<key-name> --fees=0.3plug
```
