---
order: 3
---

# JSON-RPC 方法

检查 Plug Chain 上支持的 JSON-RPC 方法和开关。
## json-rpc服务端口、激活方式和配置

可以在 `~/.plugchain/config/app.toml` 中配置：

- `enable = true|false` 字段定义了 json-rpc 服务是否可用，默认为 `true`。
- `address = {string}` 字段定义了服务绑定的地址（实际上是端口，因为主机必须保持为 `0.0.0.0`），默认为 `0.0.0.0:8545`。
- `api = {string}` 字段定义服务的开放api前缀

## 端点

| Method                                                                            | Namespace | Implemented | Public | Notes              |
|-----------------------------------------------------------------------------------|-----------|-------------|--------|--------------------|
| [`rpc_syncing`](#rpc-syncing)                                                     | Rpc       | ✔           | ✔      |                    |
| [`rpc_chainId`](#rpc-chainId)                                                     | Rpc       | ✔           | ✔      |                    |
| [`rpc_gasPrice`](#rpc-gasprice)                                                   | Rpc       | ✔           | ✔      |                    |
| [`rpc_accounts`](#rpc-accounts)                                                   | Rpc       | ✔           | ✔      |                    |
| [`rpc_blockNumber`](#rpc-blocknumber)                                             | Rpc       | ✔           | ✔      |                    |
| [`rpc_getBalance`](#rpc-getbalance)                                               | Rpc       | ✔           | ✔      |                    |
| [`rpc_balanceOf`](#rpc-balanceOf)                                                 | Rpc       | ✔           | ✔      |                    |
| [`rpc_prc20TokenInfo`](#rpc-prc20TokenInfo)                                       | Rpc       | ✔           | ✔      |                    |
| [`rpc_addressTranslation`](#rpc-addressTranslation)                               | Rpc       | ✔           | ✔      |                    |
| [`rpc_getStorageAt`](#rpc-getstorageat)                                           | Rpc       | ✔           | ✔      |                    |
| [`rpc_getTransactionCount`](#rpc-gettransactioncount)                             | Rpc       | ✔           | ✔      |                    |
| [`rpc_getBlockTransactionCountByNumber`](#rpc-getblocktransactioncountbynumber)   | Rpc       | ✔           | ✔      |                    |
| [`rpc_getBlockTransactionCountByHash`](#rpc-getblocktransactioncountbyhash)       | Rpc       | ✔           | ✔      |                    |
| [`rpc_getCode`](#rpc-getcode)                                                     | Rpc       | ✔           | ✔      |                    |
| [`rpc_getTransactionLogs`](#rpc-getTransactionLogs)                               | Rpc       | ✔           | ✔      |                    |
| [`rpc_sign`](#rpc-sign)                                                           | Rpc       | ✔           | ✔      |                    |
| [`rpc_sendTransaction`](#rpc-sendtransaction)                                     | Rpc       | ✔           | ✔      |                    |
| [`rpc_sendRawTransaction`](#rpc-sendrawtransaction)                               | Rpc       | ✔           | ✔      |                    |
| [`rpc_call`](#rpc-call)                                                           | Rpc       | ✔           | ✔      |                    |
| [`rpc_estimateGas`](#rpc-estimategas)                                             | Rpc       | ✔           | ✔      |                    |
| [`rpc_getBlockByNumber`](#rpc-getblockbynumber)                                   | Rpc       | ✔           | ✔      |                    |
| [`rpc_getBlockByHash`](#rpc-getblockbyhash)                                       | Rpc       | ✔           | ✔      |                    |
| [`rpc_getTransactionByHash`](#rpc-gettransactionbyhash)                           | Rpc       | ✔           | ✔      |                    |
| [`rpc_getTransactionReceipt`](#rpc-gettransactionreceipt)                         | Rpc       | ✔           | ✔      |                    |
| [`rpc_getProof`](#rpc-getProof)                                                   | Rpc       | ✔           |        |                    |

::: warning
块号可以输入为十六进制字符串，`"earliest"`、`"latest"` 或 `"pending"`。
:::

下面是 RPC 方法、参数和来自命名空间的示例响应的列表。


## Rpc Methods


### `rpc_syncing`

同步状态对象可能需要根据 Tendermint 同步协议的详细信息而有所不同。然而，同步结果只是一个布尔值，可以很容易地从 Tendermint 的内部同步状态中导出。

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"rpc_syncing","params":[],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":false}
```

### `rpc_chainId`

返回EIP155规则的链ID

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"rpc_chainId","params":[],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":521}
```


### `rpc_gasPrice`

返回默认 PVM 面额参数中的当前 gas 价格。

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"rpc_gasPrice","params":[],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":"0x1b"}
```

### `rpc_accounts`

返回所有帐户的数组。

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"rpc_accounts","params":[],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":["gx1vt7svcst6982ka4c3ufxmqrs05a9e9eq676tyv","gx13pcg58n0gnk68ttsupcmejvtd8fdss9lj4g0ds"]}
```

### `rpc_blockNumber`

返回当前块高度。

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"rpc_blockNumber","params":[],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":"0x5fce"}
```

### `rpc_getBalance`

返回给定帐户地址和块号的帐户余额。

#### 参数

- bech32 Address

- Block Number or Block Hash 

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"rpc_getBalance","params":["gx1vt7svcst6982ka4c3ufxmqrs05a9e9eq676tyv", "latest"],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":"0x2b9e819506dc"}
```

### `rpc_balanceOf`

返回给定帐户地址的帐户合约代币余额。

#### 参数

- bech32 Address

- contract Address

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"rpc_balanceOf","params":[{"from":"gx1vt7svcst6982ka4c3ufxmqrs05a9e9eq676tyv", "to":"gx1n4m7gafr6cxtj8q4w9mm3480ar4gq3t95k6he9"}, "latest"],"id":1}'  -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":1000000000000000000000000}
```


### `rpc_prc20TokenInfo`

返回给定合约地址的prc20标准合约的基础信息

#### 参数

- bech32 contract address


```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"rpc_prc20TokenInfo","params":[{"to":"gx1n4m7gafr6cxtj8q4w9mm3480ar4gq3t95k6he9"}, "latest"],"id":1}'  -H "Content-Type: application/json" http://localhost:8545
// Result
{"jsonrpc":"2.0","id":1,"result":{"decimals":18,"name":"test name","symbol":"tee","totalSupply":1000000000000000000000000}}
```


### `rpc_addressTranslation`

返回给定bech32，hex，eip55 地址的转换类型

#### 参数

- bech32 or hex or eip55 address string


```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"rpc_addressTranslation","params":["0x9D77e47523D60CB91c157177B8d4EFe8Ea804565"],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":{"EIP-55":"0x9d77e47523d60cb91c157177b8d4efe8ea804565","bech32":"gx1n4m7gafr6cxtj8q4w9mm3480ar4gq3t95k6he9","bytes":"[157 119 228 117 35 214 12 185 28 21 113 119 184 212 239 232 234 128 69 101]","hex":"9D77E47523D60CB91C157177B8D4EFE8EA804565"}}
```



### `rpc_getStorageAt`

返回给定帐户地址的存储地址。

#### 参数

- bech32 Address

- Integer of the position in the storage

- Block Number  or Block Hash 

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"rpc_getStorageAt","params":["gx1vt7svcst6982ka4c3ufxmqrs05a9e9eq676tyv", "0", "latest"],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":"0x0000000000000000000000000000000000000000000000000000000000000000"}
```

### `rpc_getTransactionCount`

返回给定帐户地址和块号的总交易量。

#### 参数

- bech32 Address

- Block Number or Block Hash 

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"rpc_getTransactionCount","params":["gx1vt7svcst6982ka4c3ufxmqrs05a9e9eq676tyv", "latest"],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":"0x2"}
```

### `rpc_getBlockTransactionCountByNumber`

返回给定块号的总交易计数。

#### 参数

- Block number

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"rpc_getBlockTransactionCountByNumber","params":["0x30"],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
 {"jsonrpc":"2.0","id":1,"result":"0x1"}
```

### `rpc_getBlockTransactionCountByHash`

返回给定块哈希的总交易计数。

#### 参数

- Block Hash

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"rpc_getBlockTransactionCountByHash","params":["0xe169606a2628cfba6a46227ba7edf2a18e1bb10f32b7790d5137473d3c9777d1"],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":"0x3"}
```

### `rpc_getCode`

返回给定帐户地址和块号的代码。

#### 参数

- bech32 Address

- Block Number  or Block Hash 

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"rpc_getCode","params":["gx1vt7svcst6982ka4c3ufxmqrs05a9e9eq676tyv", "0x30"],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":"0xef616c92f3cfc9e92dc270d6acff9cea213cecc7020a76ee4395af09bdceb4837a1ebdb5735e11e7d3adb6104e0c3ac55180b4ddf5e54d022cc5e8837f6a4f971b"}
```


### `rpc_getTransactionLogs`

返回给定hash的交易log日志

#### 参数

- tx hash

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"rpc_getTransactionLogs","params":["0xaa90d5fe351a053dc17241da083857842010a3e2b344ada381cca73d8b3965e9"],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":[{"address":"gx1n4m7gafr6cxtj8q4w9mm3480ar4gq3t95k6he9","topics":["0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef","0x00000000000000000000000062fd06620bd14eab76b88f126d80707d3a5c9720","0x00000000000000000000000062fd06620bd14eab76b88f126d80707d3a5c9720"],"data":"0x00000000000000000000000000000000000000000000003635c9adc5dea00000","blockNumber":"0x38c0","transactionHash":"0xb846a21fa354644e5e60d20516b898fd73290c48479df382abfe5089822a0708","transactionIndex":"0x0","blockHash":"0x955dec631d904454010530a3c4a548300b5d103d6656e6dbaef0137b628d85d7","logIndex":"0x0","removed":false}]}
```


### `rpc_sign`

`sign` 方法计算以太坊特定的签名：`sign(keccak256("\x19Ethereum Signed Message:\n" + len(message) + message)))`。

通过在消息中添加前缀，可以将计算出的签名识别为以太坊特定的签名。这可以防止恶意 DApp 可以签署任意数据（例如交易）并使用签名来冒充受害者的滥用行为。

::: warning
用于签名的地址必须解锁。
:::

#### 参数

- bech32 Address

- 要签名的信息 data

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"rpc_sign","params":["gx1vt7svcst6982ka4c3ufxmqrs05a9e9eq676tyv", "0xdeadbeaf"],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":"0x909809c76ed2a5d38733de39207d0f411222b9b49c64a192bf649cb13f63f37b45acb4f6939facb4f1c277bc70fb00407564140c0f18600ac44388f2c1dfd1dc1b"}
```

### `rpc_sendTransaction`

将交易从给定帐户发送到给定帐户。

#### 参数

- 交易对象:

    `from`: `DATA`, 20 Bytes - 发送交易的地址。

    `to`: `DATA`, 20 Bytes - （创建新合约时可选）交易指向的地址。

    `gas`: QUANTITY - （可选，默认值：20000）为交易执行提供的气体的整数。它将返回未使用的气体。

    `gasPrice`: QUANTITY - (可选，默认值：待确定）气体的整数用于每个付费气体的价格

    `value`: QUANTITY - 与此交易一起发送的值

    `data`: `DATA` - 合约的编译代码或调用的方法签名和编码参数的哈希值。详见合约 ABI

    `nonce`: QUANTITY - （可选）nonce 的整数。这允许覆盖您自己的使用相同 nonce 的待处理事务。

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"rpc_sendTransaction","params":[{"from":"gx1vt7svcst6982ka4c3ufxmqrs05a9e9eq676tyv", "to":"gx1n4m7gafr6cxtj8q4w9mm3480ar4gq3t95k6he9", "value":"0x16345", "gasLimit":"0x5208", "gasPrice":"0x55ae82600"}],"id":1}'  -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":"0x33653249db68ebe5c7ae36d93c9b2abc10745c80a72f591e296f598e2d4709f6"}
```

### `rpc_sendRawTransaction`

为已签名的交易创建新的消息调用交易或合同创建。

#### 参数

- 签署的交易数据

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"rpc_sendRawTransaction","params":["0xf9ff74c86aefeb5f6019d77280bbb44fb695b4d45cfe97e6eed7acd62905f4a85034d5c68ed25a2e7a8eeb9baf1b8401e4f865d92ec48c1763bf649e354d900b1c"],"id":1}'  -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":"0x0000000000000000000000000000000000000000000000000000000000000000"}
```

### `rpc_call`

立即执行新的消息调用，而不在区块链上创建交易。

#### 参数

- 交易对象:

    `from`: `DATA`, 20 Bytes - （可选）发送交易的地址。

    `to`: `DATA`, 20 Bytes - 交易被定向到的地址。

    `gas`: QUANTITY - 为交易执行提供的气体。 rpc_call 消耗零气体，但某些执行可能需要此参数。

    `gasPrice`: QUANTITY - 用于每个付费gas的gasPrice

    `value`: QUANTITY - 与此交易一起发送的值

    `data`: `DATA` - 可选）方法签名和编码参数的哈希。有关详细信息，请参阅 Solidity 文档中的以太坊合约 ABI

- Block number  or Block Hash 

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"rpc_call","params":[{"from":"gx1n4m7gafr6cxtj8q4w9mm3480ar4gq3t95k6he9", "to":"gx1n4m7gafr6cxtj8q4w9mm3480ar4gq3t95k6he9", "gas":"0x5208", "gasPrice":"0x55ae82600", "value":"0x16345785d8a0000", "data": "0xd46e8dd67c5d32be8d46e8dd67c5d32be8058bb8eb970870f072445675058bb8eb970870f072445675"}, "latest"],"id":1}'  -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":"0x"}
```

### `rpc_estimateGas`

返回发送交易所需气体的估计值。

#### 参数

- 交易对象:

    `from`: `DATA`, 20 Bytes - 发送交易的地址。

    `to`: `DATA`, 20 Bytes - （创建新合约时可选）交易指向的地址。

    `value`: `QUANTITY` - 与此交易一起发送的值

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"rpc_estimateGas","params":[{"from":"gx1vt7svcst6982ka4c3ufxmqrs05a9e9eq676tyv", "to":"gx1n4m7gafr6cxtj8q4w9mm3480ar4gq3t95k6he9", "value":"0x30"}],"id":1}'  -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":"0x1199b"}
```

### `rpc_getBlockByNumber`

按块号返回有关块的信息。

#### 参数

- 块号

-  true|false  如果为 true，则返回完整的交易对象，如果为 false，则仅返回交易的哈希值。

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"rpc_getBlockByNumber","params":["0x30",false],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":{"baseFeePerGas":"0x191d10","difficulty":"0x0","extraData":"0x","gasLimit":"0x989680","gasUsed":"0x115c9c","hash":"0xe169606a2628cfba6a46227ba7edf2a18e1bb10f32b7790d5137473d3c9777d1","logsBloom":"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","miner":"gx13pcg58n0gnk68ttsupcmejvtd8fdss9lj4g0ds","mixHash":"0x0000000000000000000000000000000000000000000000000000000000000000","nonce":"0x0000000000000000","number":"0x30","parentHash":"0x0b6593fb90b0c063236ba6c658030f47fe1fbcfc9098b36ed6a4b462077ea9bf","receiptsRoot":"0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421","sha3Uncles":"0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347","size":"0x1714","stateRoot":"0xbfbfe1aa039288795aa33ac4bd5281927712fa9e1232cbb5a179ed0a9aa1a430","timestamp":"0x6262243f","totalDifficulty":"0x0","transactions":["0xaa90d5fe351a053dc17241da083857842010a3e2b344ada381cca73d8b3965e9"],"transactionsRoot":"0x2b7666f8dde9d3fafa4573f751d1e084765f66aca3f1d78bb106ff526eab335c","uncles":[]}}
```

### `rpc_getBlockByHash`

返回给定上面命令中的哈希值和一个布尔值，锁定信息。

#### 参数 

- 块hash.

- true|false 如果为 true，则返回完整的交易对象，如果为 false，则仅返回交易的哈希值。

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"rpc_getBlockByHash","params":["0xe169606a2628cfba6a46227ba7edf2a18e1bb10f32b7790d5137473d3c9777d1",false],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":{"baseFeePerGas":"0x191d10","difficulty":"0x0","extraData":"0x","gasLimit":"0x989680","gasUsed":"0x115c9c","hash":"0xe169606a2628cfba6a46227ba7edf2a18e1bb10f32b7790d5137473d3c9777d1","logsBloom":"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","miner":"gx13pcg58n0gnk68ttsupcmejvtd8fdss9lj4g0ds","mixHash":"0x0000000000000000000000000000000000000000000000000000000000000000","nonce":"0x0000000000000000","number":"0x30","parentHash":"0x0b6593fb90b0c063236ba6c658030f47fe1fbcfc9098b36ed6a4b462077ea9bf","receiptsRoot":"0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421","sha3Uncles":"0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347","size":"0x1714","stateRoot":"0xbfbfe1aa039288795aa33ac4bd5281927712fa9e1232cbb5a179ed0a9aa1a430","timestamp":"0x6262243f","totalDifficulty":"0x0","transactions":["0xaa90d5fe351a053dc17241da083857842010a3e2b344ada381cca73d8b3965e9"],"transactionsRoot":"0x2b7666f8dde9d3fafa4573f751d1e084765f66aca3f1d78bb106ff526eab335c","uncles":[]}}
```

### `rpc_getTransactionByHash`

返回给定 tx hash 的交易细节。

#### 参数

- 交易的哈希

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"rpc_getTransactionByHash","params":["0xb846a21fa354644e5e60d20516b898fd73290c48479df382abfe5089822a0708"],"id":1}' -H "Content-Type: application/json" http://localhost:8545
 
// Result
{"jsonrpc":"2.0","id":1,"result":{"blockHash":"0x955dec631d904454010530a3c4a548300b5d103d6656e6dbaef0137b628d85d7","blockNumber":"0x38c0","from":"gx1vt7svcst6982ka4c3ufxmqrs05a9e9eq676tyv","gas":"0x33450","gasPrice":"0xc8","hash":"0xb846a21fa354644e5e60d20516b898fd73290c48479df382abfe5089822a0708","input":"0xa9059cbb00000000000000000000000062fd06620bd14eab76b88f126d80707d3a5c972000000000000000000000000000000000000000000000003635c9adc5dea00000","nonce":"0x1","to":"gx1n4m7gafr6cxtj8q4w9mm3480ar4gq3t95k6he9","transactionIndex":"0x0","value":"0x0","type":"0x0","v":"0x436","r":"0xe6072ed57727a81087ff8e4e1df8a168038a6295d8ff3835dfab1bd7a387b9f","s":"0xf634ebcad4215e8417572c31ec6f0dc990981b193d3159fdfbfd81b04ea8be8"}}
```


### `rpc_getTransactionReceipt`

通过交易哈希返回交易的收据。

注意: Tendermint 的 Tx Code 和 Ethereum 收据状态被切换：
|         | Tendermint | Rpcereum |
|---------|------------|----------|
| Success | 0          | 1        |
| Fail    | 1          | 0        |

#### 参数

- tx hash

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"rpc_getTransactionReceipt","params":["0xb846a21fa354644e5e60d20516b898fd73290c48479df382abfe5089822a0708"],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":{"blockHash":"0x955dec631d904454010530a3c4a548300b5d103d6656e6dbaef0137b628d85d7","blockNumber":"0x38c0","contractAddress":null,"cumulativeGasUsed":"0x6dc2","from":"gx1vt7svcst6982ka4c3ufxmqrs05a9e9eq676tyv","gasUsed":"0x6dc2","logs":[{"address":"gx1n4m7gafr6cxtj8q4w9mm3480ar4gq3t95k6he9","topics":["0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef","0x00000000000000000000000062fd06620bd14eab76b88f126d80707d3a5c9720","0x00000000000000000000000062fd06620bd14eab76b88f126d80707d3a5c9720"],"data":"0x00000000000000000000000000000000000000000000003635c9adc5dea00000","blockNumber":"0x38c0","transactionHash":"0xb846a21fa354644e5e60d20516b898fd73290c48479df382abfe5089822a0708","transactionIndex":"0x0","blockHash":"0x955dec631d904454010530a3c4a548300b5d103d6656e6dbaef0137b628d85d7","logIndex":"0x0","removed":false}],"logsBloom":"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000008008000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000110000000000000000000000200000000000000020000000010000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","status":"0x1","to":"gx1n4m7gafr6cxtj8q4w9mm3480ar4gq3t95k6he9","transactionHash":"0xb846a21fa354644e5e60d20516b898fd73290c48479df382abfe5089822a0708","transactionIndex":"0x0","type":"0x0"}}
```

### rpc_getProof

返回指定帐户的帐户值和存储值，包括默克尔证明。

#### 参数

-  账户或者合约地址

- 存储位置的整数

- Block Number  or Block Hash 

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"rpc_getProof","params":["gx1n4m7gafr6cxtj8q4w9mm3480ar4gq3t95k6he9",["0x0000000000000000000000000000000000000000000000000000000000000000","0x0000000000000000000000000000000000000000000000000000000000000001"],"latest"],"id":1}' -H "Content-type:application/json" http://localhost:8545

// Result
{"jsonrpc": "2.0", "id": 1, "result": {"address": "gx1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq8rta3w", "accountProof": ["0xf90211a090dcaf88c40c7bbc95a912cbdde67c175767b31173df9ee4b0d733bfdd511c43a0babe369f6b12092f49181ae04ca173fb68d1a5456f18d20fa32cba73954052bda0473ecf8a7e36a829e75039a3b055e51b8332cbf03324ab4af2066bbd6fbf0021a0bbda34753d7aa6c38e603f360244e8f59611921d9e1f128372fec0d586d4f9e0a04e44caecff45c9891f74f6a2156735886eedf6f1a733628ebc802ec79d844648a0a5f3f2f7542148c973977c8a1e154c4300fec92f755f7846f1b734d3ab1d90e7a0e823850f50bf72baae9d1733a36a444ab65d0a6faaba404f0583ce0ca4dad92da0f7a00cbe7d4b30b11faea3ae61b7f1f2b315b61d9f6bd68bfe587ad0eeceb721a07117ef9fc932f1a88e908eaead8565c19b5645dc9e5b1b6e841c5edbdfd71681a069eb2de283f32c11f859d7bcf93da23990d3e662935ed4d6b39ce3673ec84472a0203d26456312bbc4da5cd293b75b840fc5045e493d6f904d180823ec22bfed8ea09287b5c21f2254af4e64fca76acc5cd87399c7f1ede818db4326c98ce2dc2208a06fc2d754e304c48ce6a517753c62b1a9c1d5925b89707486d7fc08919e0a94eca07b1c54f15e299bd58bdfef9741538c7828b5d7d11a489f9c20d052b3471df475a051f9dd3739a927c89e357580a4c97b40234aa01ed3d5e0390dc982a7975880a0a089d613f26159af43616fd9455bb461f4869bfede26f2130835ed067a8b967bfb80", "0xf90211a0395d87a95873cd98c21cf1df9421af03f7247880a2554e20738eec2c7507a494a0bcf6546339a1e7e14eb8fb572a968d217d2a0d1f3bc4257b22ef5333e9e4433ca012ae12498af8b2752c99efce07f3feef8ec910493be749acd63822c3558e6671a0dbf51303afdc36fc0c2d68a9bb05dab4f4917e7531e4a37ab0a153472d1b86e2a0ae90b50f067d9a2244e3d975233c0a0558c39ee152969f6678790abf773a9621a01d65cd682cc1be7c5e38d8da5c942e0a73eeaef10f387340a40a106699d494c3a06163b53d956c55544390c13634ea9aa75309f4fd866f312586942daf0f60fb37a058a52c1e858b1382a8893eb9c1f111f266eb9e21e6137aff0dddea243a567000a037b4b100761e02de63ea5f1fcfcf43e81a372dafb4419d126342136d329b7a7ba032472415864b08f808ba4374092003c8d7c40a9f7f9fe9cc8291f62538e1cc14a074e238ff5ec96b810364515551344100138916594d6af966170ff326a092fab0a0d31ac4eef14a79845200a496662e92186ca8b55e29ed0f9f59dbc6b521b116fea090607784fe738458b63c1942bba7c0321ae77e18df4961b2bc66727ea996464ea078f757653c1b63f72aff3dcc3f2a2e4c8cb4a9d36d1117c742833c84e20de994a0f78407de07f4b4cb4f899dfb95eedeb4049aeb5fc1635d65cf2f2f4dfd25d1d7a0862037513ba9d45354dd3e36264aceb2b862ac79d2050f14c95657e43a51b85c80", "0xf90171a04ad705ea7bf04339fa36b124fa221379bd5a38ffe9a6112cb2d94be3a437b879a08e45b5f72e8149c01efcb71429841d6a8879d4bbe27335604a5bff8dfdf85dcea00313d9b2f7c03733d6549ea3b810e5262ed844ea12f70993d87d3e0f04e3979ea0b59e3cdd6750fa8b15164612a5cb6567cdfb386d4e0137fccee5f35ab55d0efda0fe6db56e42f2057a071c980a778d9a0b61038f269dd74a0e90155b3f40f14364a08538587f2378a0849f9608942cf481da4120c360f8391bbcc225d811823c6432a026eac94e755534e16f9552e73025d6d9c30d1d7682a4cb5bd7741ddabfd48c50a041557da9a74ca68da793e743e81e2029b2835e1cc16e9e25bd0c1e89d4ccad6980a041dda0a40a21ade3a20fcd1a4abb2a42b74e9a32b02424ff8db4ea708a5e0fb9a09aaf8326a51f613607a8685f57458329b41e938bb761131a5747e066b81a0a16808080a022e6cef138e16d2272ef58434ddf49260dc1de1f8ad6dfca3da5d2a92aaaadc58080", "0xf851808080a009833150c367df138f1538689984b8a84fc55692d3d41fe4d1e5720ff5483a6980808080808080808080a0a319c1c415b271afc0adcb664e67738d103ac168e0bc0b7bd2da7966165cb9518080"], "balance": "0x0", "codeHash": "0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470", "nonce": "0x0", "storageHash": "0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421", "storageProof": [{"key": "0x0000000000000000000000000000000000000000000000000000000000000000", "value": "0x0", "proof": []}, {"key": "0x0000000000000000000000000000000000000000000000000000000000000001", "value": "0x0", "proof": []}]}}
```

