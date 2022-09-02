---
order: 6
---

# Gas and Fees

了解 PVM 和 Cosmos 中“Gas”和“Fees”之间的区别。

## 介绍 `Gas` 和 `Fees`

在 Cosmos SDK 中，`gas` 是一个特殊的单位，用于跟踪执行期间的资源消耗。 `gas` 通常在对存储进行读写时消耗，但如果需要进行昂贵的计算，也可以消耗它。 它有两个主要目的：

* 确保区块不会消耗太多资源并将最终确定。
* 防止最终用户的垃圾邮件和滥用。 为此，在 `Msg` 执行期间消耗的 `gas` 通常会被定价，从而产生 `fee`（`fees = gas * gas-prices`）。 `fees`一般必须由执行`Msg`的发送者支付。 请注意，Cosmos SDK 默认不强制执行“gas”定价，因为可能有其他方法可以防止垃圾邮件（例如带宽方案）。 尽管如此，大多数应用程序将实施“费用”机制来防止垃圾邮件。


## Cosmos SDK `Fees`

* 验证交易提供的 gas 价格是否高于当地的 `min-gas-prices`。 `min-gas-prices` 是每个全节点的本地参数，在 `CheckTx` 期间用于丢弃不提供最低费用的交易。 这确保了内存池不会被垃圾事务发送到垃圾邮件中。
* 公式: `fees = gas * gas-price`
* 例如: `gas=200000`（默认）、`gas-price=0.0001`、`fees=20uplugcn`


## PVM `Fees`


PVM 和 Cosmos 状态转换之间的主要区别在于，PVM 对每个 OPCODE 使用 [gas table](https://github.com/ethereum/go-ethereum/blob/master/params/protocol_params.go)，而 Cosmos 使用“GasConfig”，通过设置访问数据库的单位和每字节成本来为每个 CRUD 操作收取 gas。

PVM和EVM的Gas使用计算方式相同， 但PVM `gas`使用 与 EVM `gas`使用流程相同，当gas足够支付交易时，多余的gas会返回。

* PVM交易首先要满足`fees`大于Cosmos SDK`要求的最低手续费（目前最低手续费为0.0001*200000=20uplugcn）`
* 交易提供的gas应大于或等于根据[gas table](https://github.com/ethereum/go-ethereum/blob/master/params)交易字节计算的gas总和
* `gasPrice` 最低为 `7`
* 公式: `fees = gas * gasPrice / 1000000` (提案[v1.7.0](https://www.plugchain.network/v2/communityDetail?id=10) 升级之后的算法)