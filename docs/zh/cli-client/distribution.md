# Distribution

distribution模块用于管理自己的 [Staking 收益](../concepts/general-concepts.md#staking-收益)。

## 可用命令

| 名称                                                                                    | 描述                                                                                           |
| --------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| [commission](#plugchaind-query-distribution-commission)                                       | 查询分配的验证人佣金                                                                           |
| [community-pool](#plugchaind-query-distribution-community-pool)                               | 查询社区池总币数                                                                               |
| [params](#plugchaind-query-distribution-params)                                               | 查询分配参数                                                                                   |
| [rewards](#plugchaind-query-distribution-rewards)                                             | 查询所有分销委托人收益或来自指定验证人的收益                                                   |
| [slashes](#plugchaind-query-distribution-slashes)                                             | 查询验证人指定块范围内的分割                                                                   |
| [validator-outstanding-rewards](#plugchaind-query-distribution-validator-outstanding-rewards) | 查询验证人的未付奖励分配及其所有授权                                                           |
| [fund-community-pool](#plugchaind-tx-distribution-fund-community-pool)                        | 为社区基金池提供指定数额的资金                                                                 |
| [set-withdraw-addr](#plugchaind-tx-distribution-set-withdraw-addr)                            | 设置提现地址                                                                                   |
| [withdraw-all-rewards](#plugchaind-tx-distribution-withdraw-all-rewards)                      | 取回委托人所有收益                                                                             |
| [withdraw-rewards](#plugchaind-tx-distribution-withdraw-rewards)                              | 取回收益，有以下几种模式：取回所有奖励、从指定的验证者取回委派奖励、验证人取回所有奖励以及佣金 |

## plugchaind query distribution commission

查询分配的验证人佣金。

```bash
plugchaind query distribution commission [validator] [flags]
```

## plugchaind query distribution community-pool

查询社区池总币数。

```bash
plugchaind query distribution community-pool [flags]
```

## plugchaind query distribution params

查询分配参数。

```bash
 plugchaind query distribution params [flags]
```

## plugchaind query distribution rewards

查询所有分销委托人收益或来自指定验证人的收益。

```bash
plugchaind query distribution rewards [delegator-addr] [validator-addr] [flags]
```

## plugchaind query distribution slashes

查询验证人指定块范围内的的惩罚。

```bash
plugchaind query distribution slashes [validator] [start-height] [end-height] [flags]
```

## plugchaind query distribution validator-outstanding-rewards

查询验证人的未付奖励分配及其所有授权。

```bash
plugchaind query distribution validator-outstanding-rewards [validator] [flags]
```

## plugchaind tx distribution fund-community-pool

为社区基金池提供指定数额的资金。

```bash
plugchaind tx distribution fund-community-pool [amount] [flags]
```

## plugchaind tx distribution set-withdraw-addr

设置提现地址。

```bash
plugchaind tx distribution set-withdraw-addr [withdraw-addr] [flags]
```

## plugchaind tx distribution withdraw-all-rewards

取回委托人所有收益。

```bash
plugchaind tx distribution withdraw-all-rewards [flags]
```

## plugchaind tx distribution withdraw-rewards

取回收益，有以下几种模式：取回所有奖励、从指定的验证者取回委派奖励、验证人取回所有奖励以及佣金。

```bash
plugchaind tx distribution withdraw-rewards [validator-addr] [flags]
```
