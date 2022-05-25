---
order: 6
---

# 治理参数

在Plug Chain中，存在一些特殊的参数，它们可通过链上治理被修改。持有plug通证的用户都可以参与到参数修改的链上治理。
如果社区对某些可修改的参数不满意，可以发起[参数修改提案](../features/governance.md#usage-scenario-of-parameter-change)，社区投票通过后即可在线自动完成修改。

## Auth 模块可治理参数

| 字段                          | 描述                               | 有效范围                  | 当前值 |
| ----------------------------- | ---------------------------------- | ------------------------- | ------ |
| `auth/MaxMemoCharacters`      | 交易的memo字段的最大字符数         | (0, 18446744073709551615] | 256    |
| `auth/TxSigLimit`             | 每个交易的最大签名数               | (0, 18446744073709551615] | 7      |
| `auth/TxSizeCostPerByte`      | 交易每个字节消耗的gas量            | (0, 18446744073709551615] | 10     |
| `auth/SigVerifyCostED25519`   | 在ED25519算法签名验证上花费的gas   | (0, 18446744073709551615] | 590    |
| `auth/SigVerifyCostSecp256k1` | 在Secp256k1算法签名验证上花费的gas | (0, 18446744073709551615] | 1000   |

## Bank 模块可治理参数

| 字段                      | 描述                    | 有效范围     | 当前值 |
| ------------------------- | ----------------------- | ------------ | ------ |
| `bank/SendEnabled`        | 支持transfer的token集合 |              | []     |
| `bank/DefaultSendEnabled` | 默认是否支持转账功能    | {true,false} | true   |

详见 [Bank](../features/bank.md)

## Distribution 模块可治理参数

| 字段                               | 描述                   | 有效范围     | 当前值 |
| ---------------------------------- | ---------------------- | ------------ | ------ |
| `distribution/communitytax`        | 提现收取的手续费率     | [0, 1]       | 0.02   |
| `distribution/baseproposerreward`  | 区块提议者的基础奖励率 | [0, 1]       | 0.01   |
| `distribution/bonusproposerreward` | 区块提议者的奖励率     | [0, 1]       | 0.04   |
| `distribution/withdrawaddrenabled` | 是否支持设置提现地址   | {true,false} | true   |

详见 [Distribution](../features/distribution.md)

## Gov 模块可治理参数

| 字段                | 描述                   | 有效范围                                                 | 当前值                                                                                                         |
| ------------------- | ---------------------- | -------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `gov/depositparams` | 提议抵押阶段的相关参数 | max_deposit_period:(0, 9223372036854775807]              | {"min_deposit": [{"denom": "uplugcn", "amount": "10000000"}], "max_deposit_period": "604800s" },                |
| `gov/votingparams`  | 提议投票阶段的相关参数 | voting_period:(0, 9223372036854775807]                   | {"voting_period": "1209600s"}                                                                                   |
| `gov/tallyparams`   | 投票统计阶段的相关参数 | quorum:[0,1]<br>threshold:(0,1]<br/>veto_threshold:(0,1] | {"quorum":"0.334000000000000000","threshold": "0.500000000000000000","veto_threshold": "0.334000000000000000"} |

详见 [Governance](../features/governance.md)

## IBC 模块可治理参数

| 字段                      | 描述             | 有效范围     | 当前值                             |
| ------------------------- | ---------------- | ------------ | ---------------------------------- |
| `ibc/AllowedClients`      | 支持ibc的客户端  |              | ["06-solomachine","07-tendermint"] |
| `transfer/SendEnabled`    | 是否开启交易功能  | {true,false} | true                              |
| `transfer/ReceiveEnabled` | 是否开启接收功能  | {true,false} | true                              |

## Mint 模块可治理参数

| 字段             | 描述           | 有效范围 | 当前值 |
| ---------------- | -------------- | -------- | ------ |
| `mint/Inflation` | 代币增发频率   | [0, 0.2] | 0.13   |
| `mint/MintDenom` | 增发的代币名称 |          | uplugcn  |
| `mint/params` | 增发参数 |    []      | {"mint_denom": "uplugcn","inflation_rate_change": "0.130000000000000000","inflation_max": "0.200000000000000000","inflation_min": "0.070000000000000000","goal_bonded": "0.670000000000000000","blocks_per_year": "6311520"} |

详见 [Mint](../features/mint.md)

## Slashing 模块可治理参数

| 字段                               | 描述                     | 有效范围                  | 当前值 |
| ---------------------------------- | ------------------------ | ------------------------- | ------ |
| `slashing/DowntimeJailDuration`    | 验证人最大的下线时间     | (0, 9223372036854775807]  | 600s  |
| `slashing/MinSignedPerWindow`      | 每个窗口最低投票率       | [0, 1]                    | 0.5    |
| `slashing/SignedBlocksWindow`      | 验证人下线的滑动窗口大小 | (0, 18446744073709551615] | 100  |
| `slashing/SlashFractionDoubleSign` | 双重签名的惩罚系数       | [0, 1]                    | 0.05   |
| `slashing/SlashFractionDowntime`   | 验证人下线的惩罚系数     | [0, 1]                    | 0.01 |

详见 [Slashing](../features/slashing.md)

## Staking 模块可治理参数

| 字段                        | 描述                   | 有效范围                 | 当前值   |
| --------------------------- | ---------------------- | ------------------------ | -------- |
| `staking/UnbondingTime`     | 抵押解绑时间           | (0, 9223372036854775807] | 1814400s |
| `staking/MaxValidators`     | 最大验证人数量         | (0, 4294967295]          | 100      |
| `staking/MaxEntries`        | 解绑、转委托的最大数量 | (0, 4294967295]          | 7        |
| `staking/BondDenom`         | 可抵押的代币           |                          | uplugcn    |
| `staking/HistoricalEntries` | 历史条目               | [0, 4294967295]          | 10000    |

详见 [Staking](../features/staking.md)
