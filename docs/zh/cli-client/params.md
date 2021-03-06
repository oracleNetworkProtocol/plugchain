# Params

Params模块允许查询系统里预设的参数,查询结果中除了Gov模块的参数,其他都可以通过[Gov模块](./gov.md)发起提议来修改。

```bash
plugchaind query params subspace [subspace] [key] [flags]
```

`subspace`目前支持：`auth`,`bank`,`staking`,`mint`,`distribution`,`slashing`,`gov`,`crisis`,`token`。
其中,可用于每个`subspace`查询的参数如下：

## auth

| key                      | description                    | default |
| ------------------------ | ------------------------------ | ------- |
| `MaxMemoCharacters`      | 交易字段中备注的最大字符数     | 256     |
| `TxSigLimit`             | 每笔交易的最大签名数量         | 7       |
| `TxSizeCostPerByte`      | 交易的每个字节消耗的Gas        | 10      |
| `SigVerifyCostED25519`   | edd2519算法签名验证消耗的Gas   | 590     |
| `SigVerifyCostSecp256k1` | secp256k1算法签名验证消耗的Gas | 1000    |

## bank

| key                  | description      | default |
| -------------------- | ---------------- | ------- |
| `SendEnabled`        | 支持转账的代币   | {}      |
| `DefaultSendEnabled` | 是否开启转账功能 | true    |

## staking

| key                 | description            | default   |
| ------------------- | ---------------------- | --------- |
| `UnbondingTime`     | 抵押解绑时间           | 3w(weeks) |
| `MaxValidators`     | 最大验证人数量         | 100       |
| `MaxEntries`        | 解绑、转委托的最大数量 | 7         |
| `BondDenom`         | 可抵押的代币           | uplugcn     |
| `HistoricalEntries` | 历史条目数             | 100       |

## mint

| key         | description    | default |
| ----------- | -------------- | ------- |
| `Inflation` | 代币增发频率   | 0.04    |
| `MintDenom` | 增发的代币名称 | uplugcn   |

## distribution

| key                   | description            | default |
| --------------------- | ---------------------- | ------- |
| `communitytax`        | 提现收取的手续费率     | 0.02    |
| `baseproposerreward`  | 区块提议者的基础奖励率 | 0.01    |
| `bonusproposerreward` | 区块提议者的奖励率     | 0.04    |
| `withdrawaddrenabled` | 是否支持设置提现地址   | true    |

## slashing

| key                       | description              | default |
| ------------------------- | ------------------------ | ------- |
| `SignedBlocksWindow`      | 验证人下线的滑动窗口大小 | 100     |
| `MinSignedPerWindow`      | 每个窗口最低投票率       | 0.5     |
| `DowntimeJailDuration`    | 验证人最大的下线时间     | 10m     |
| `SlashFractionDoubleSign` | 双重签名的惩罚系数       | 1/20    |
| `SlashFractionDowntime`   | 验证人下线的惩罚系数     | 1/100   |

## gov

| key             | description            | default                                                      |
| --------------- | ---------------------- | ------------------------------------------------------------ |
| `depositparams` | 提议抵押阶段的相关参数 | `min_deposit`: 10000000plug; `max_deposit_period`: 2d(days) |
| `votingparams`  | 提议投票阶段的相关参数 | `voting_period`: 2d(days)                                    |
| `tallyparams`   | 投票统计阶段的相关参数 | `quorum`: 0.334; `threshold`: 0.5; `veto_threshold`: 0.334   |

## crisis

| key           | description | default   |
| ------------- | ----------- | --------- |
| `ConstantFee` | 固定费用    | 1000uplugcn |


## token

| key           | description | default   |
| ------------- | ----------- | --------- |
| `IssueTokenBaseFee` | 固定费用    | 10000uplugcn |
| `OperateTokenFeeRatio` | 操作费用比例 | 0.1 |
