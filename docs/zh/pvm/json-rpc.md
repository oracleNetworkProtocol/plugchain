---
order: 3
---

# JSON-RPC 方法

检查 Plug Chain 上支持的 JSON-RPC 方法和开关。
## json-rpc服务端口、激活方式和配置

可以在 `~/.plugchain/config/app.toml` 中配置：

- `enable = true|false` 字段定义了 json-rpc 服务是否可用，默认为 `true`。
- `address = {string}` 字段定义了服务绑定的地址（实际上是端口，因为主机必须保持为 `0.0.0.0`），默认为 `0.0.0.0:8545`。
<<<<<<< HEAD
- `api = {string}` 字段定义服务的开放api前缀
=======
- `api = {string}` 字段定义服务的开放api命名空间



## 方法命名空间

| Namespace                                     | Description                                                                                                                                                                                                                  | Supported | Enabled by Default |
| --------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------- | ------------------ |
| [`rpc`](#rpc-methods)           | The rpc namespace that interacts with the `gx` prefix address, the prefix will successively implement all JSON-RPC of ETH . | ✔        | ✔                 |    
| [`eth`](#eth-methods)           | Evmos provides several extensions to the standard `eth` JSON-RPC namespace.                                                                                                                                                  | ✔        | ✔                 |
| [`web3`](#web3-methods)         | The `web3` API provides utility functions for the web3 client.                                                                                                                                                               | ✔        | ✔                 |
| [`net`](#net-methods)           | The `net` API provides access to network information of the node                                                                                                                                                             | ✔        | ✔                 |
| `clique`                                      | The `clique` API provides access to the state of the clique consensus engine. You can use this API to manage signer votes and to check the health of a private network.                                                      | ❌        |                    |
| `debug`                                       | The `debug` API gives you access to several non-standard RPC methods, which will allow you to inspect, debug and set certain debugging flags during runtime.                                                                 | ✔        |                    |
| `les`                                         | The `les` API allows you to manage LES server settings, including client parameters and payment settings for prioritized clients. It also provides functions to query checkpoint information in both server and client mode. | ❌        |                    |
| [`miner`](#miner-methods)       | The `miner` API allows you to remote control the node’s mining operation and set various mining specific settings.                                                                                                           | ✔        | ❌                 |
| [`txpool`](#txpool-methods)     | The `txpool` API gives you access to several non-standard RPC methods to inspect the contents of the transaction pool containing all the currently pending transactions as well as the ones queued for future processing.    | ✔        | ❌                 |
| `admin`                                       | The `admin` API gives you access to several non-standard RPC methods, which will allow you to have a fine grained control over your nodeinstance, including but not limited to network peer and RPC endpoint management.     | ❌        |                    |
| [`personal`](#personal-methods) | The `personal` API manages private keys in the key store.                                                                                                                                                                    | ✔        | ❌                 |

>>>>>>> d79323fb0f7dd4dd7b9baf4a27f51c8e5ac75472

## 端点

| Method                                                                            | Namespace | Implemented | Public | Notes              |
|-----------------------------------------------------------------------------------|-----------|-------------|--------|--------------------|
| [`rpc_syncing`](#rpc-syncing)                                                     | Rpc       | ✔           | ✔      |                    |
| [`rpc_chainId`](#rpc-chainid)                                                     | Rpc       | ✔           | ✔      |                    |
| [`rpc_gasPrice`](#rpc-gasprice)                                                   | Rpc       | ✔           | ✔      |                    |
| [`rpc_accounts`](#rpc-accounts)                                                   | Rpc       | ✔           | ✔      |                    |
| [`rpc_blockNumber`](#rpc-blocknumber)                                             | Rpc       | ✔           | ✔      |                    |
| [`rpc_getBalance`](#rpc-getbalance)                                               | Rpc       | ✔           | ✔      |                    |
| [`rpc_balanceOf`](#rpc-balanceof)                                                 | Rpc       | ✔           | ✔      |                    |
| [`rpc_prc20TokenInfo`](#rpc-prc20tokeninfo)                                       | Rpc       | ✔           | ✔      |                    |
| [`rpc_addressTranslation`](#rpc-addresstranslation)                               | Rpc       | ✔           | ✔      |                    |
| [`rpc_getStorageAt`](#rpc-getstorageat)                                           | Rpc       | ✔           | ✔      |                    |
| [`rpc_getTransactionCount`](#rpc-gettransactioncount)                             | Rpc       | ✔           | ✔      |                    |
| [`rpc_getBlockTransactionCountByNumber`](#rpc-getblocktransactioncountbynumber)   | Rpc       | ✔           | ✔      |                    |
| [`rpc_getBlockTransactionCountByHash`](#rpc-getblocktransactioncountbyhash)       | Rpc       | ✔           | ✔      |                    |
| [`rpc_getCode`](#rpc-getcode)                                                     | Rpc       | ✔           | ✔      |                    |
| [`rpc_getTransactionLogs`](#rpc-gettransactionlogs)                               | Rpc       | ✔           | ✔      |                    |
| [`rpc_sign`](#rpc-sign)                                                           | Rpc       | ✔           | ✔      |                    |
| [`rpc_sendTransaction`](#rpc-sendtransaction)                                     | Rpc       | ✔           | ✔      |                    |
| [`rpc_sendRawTransaction`](#rpc-sendrawtransaction)                               | Rpc       | ✔           | ✔      |                    |
| [`rpc_call`](#rpc-call)                                                           | Rpc       | ✔           | ✔      |                    |
| [`rpc_estimateGas`](#rpc-estimategas)                                             | Rpc       | ✔           | ✔      |                    |
| [`rpc_getBlockByNumber`](#rpc-getblockbynumber)                                   | Rpc       | ✔           | ✔      |                    |
| [`rpc_getBlockByHash`](#rpc-getblockbyhash)                                       | Rpc       | ✔           | ✔      |                    |
| [`rpc_getTransactionByHash`](#rpc-gettransactionbyhash)                           | Rpc       | ✔           | ✔      |                    |
| [`rpc_getTransactionReceipt`](#rpc-gettransactionreceipt)                         | Rpc       | ✔           | ✔      |                    |
| [`rpc_getProof`](#rpc-getproof)                                                   | Rpc       | ✔           |        |                    |
<<<<<<< HEAD
=======
| [`web3_clientVersion`](#web3-clientversion)                                       | Web3      | ✔           | ✔      |                    |
| [`web3_sha3`](#web3-sha3)                                                         | Web3      | ✔           | ✔      |                    |
| [`net_version`](#net-version)                                                     | Net       | ✔           | ✔      |                    |
| [`net_peerCount`](#net-peerCount)                                                 | Net       | ✔           | ✔      |                    |
| [`net_listening`](#net-listening)                                                 | Net       | ✔           | ✔      |                    |
| [`eth_protocolVersion`](#eth-protocolversion)                                     | Eth       | ✔           | ✔      |                    |
| [`eth_syncing`](#eth-syncing)                                                     | Eth       | ✔           | ✔      |                    |
| [`eth_gasPrice`](#eth-gasprice)                                                   | Eth       | ✔           | ✔      |                    |
| [`eth_accounts`](#eth-accounts)                                                   | Eth       | ✔           | ✔      |                    |
| [`eth_blockNumber`](#eth-blocknumber)                                             | Eth       | ✔           | ✔      |                    |
| [`eth_getBalance`](#eth-getbalance)                                               | Eth       | ✔           | ✔      |                    |
| [`eth_getStorageAt`](#eth-getstorageat)                                           | Eth       | ✔           | ✔      |                    |
| [`eth_getTransactionCount`](#eth-gettransactioncount)                             | Eth       | ✔           | ✔      |                    |
| [`eth_getBlockTransactionCountByNumber`](#eth-getblocktransactioncountbynumber)   | Eth       | ✔           | ✔      |                    |
| [`eth_getBlockTransactionCountByHash`](#eth-getblocktransactioncountbyhash)       | Eth       | ✔           | ✔      |                    |
| [`eth_getCode`](#eth-getcode)                                                     | Eth       | ✔           | ✔      |                    |
| [`eth_sign`](#eth-sign)                                                           | Eth       | ✔           | ✔      |                    |
| [`eth_sendTransaction`](#eth-sendtransaction)                                     | Eth       | ✔           | ✔      |                    |
| [`eth_sendRawTransaction`](#eth-sendrawtransaction)                               | Eth       | ✔           | ✔      |                    |
| [`eth_call`](#eth-call)                                                           | Eth       | ✔           | ✔      |                    |
| [`eth_estimateGas`](#eth-estimategas)                                             | Eth       | ✔           | ✔      |                    |
| [`eth_getBlockByNumber`](#eth-getblockbynumber)                                   | Eth       | ✔           | ✔      |                    |
| [`eth_getBlockByHash`](#eth-getblockbyhash)                                       | Eth       | ✔           | ✔      |                    |
| [`eth_getTransactionByHash`](#eth-gettransactionbyhash)                           | Eth       | ✔           | ✔      |                    |
| [`eth_getTransactionByBlockHashAndIndex`](#eth-gettransactionbyblockhashandindex) | Eth       | ✔           | ✔      |                    |
| [`eth_getTransactionReceipt`](#eth-gettransactionreceipt)                         | Eth       | ✔           | ✔      |                    |
| [`eth_newFilter`](#eth-newfilter)                                                 | Eth       | ✔           | ✔      |                    |
| [`eth_newBlockFilter`](#eth-newblockfilter)                                       | Eth       | ✔           | ✔      |                    |
| [`eth_newPendingTransactionFilter`](#eth-newpendingtransactionfilter)             | Eth       | ✔           | ✔      |                    |
| [`eth_uninstallFilter`](#eth-uninstallfilter)                                     | Eth       | ✔           | ✔      |                    |
| [`eth_getFilterChanges`](#eth-getfilterchanges)                                   | Eth       | ✔           | ✔      |                    |
| [`eth_getFilterLogs`](#eth-getfilterlogs)                                         | Eth       | ✔           | ✔      |                    |
| [`eth_getLogs`](#eth-getlogs)                                                     | Eth       | ✔           | ✔      |                    |
| `eth_getTransactionbyBlockNumberAndIndex`                                         | Eth       |             | ✔      |                    |
| `eth_getWork`                                                                     | Eth       | N/A         | ✔      | PoW-only           |
| `eth_submitWork`                                                                  | Eth       | N/A         | ✔      | PoW-only           |
| `eth_submitHashrate`                                                              | Eth       |             |        |                    |
| `eth_getCompilers`                                                                | Eth       |             |        |                    |
| `eth_compileLLL`                                                                  | Eth       |             |        |                    |
| `eth_compileSolidity`                                                             | Eth       |             |        |                    |
| `eth_compileSerpent`                                                              | Eth       |             |        |                    |
| `eth_signTransaction`                                                             | Eth       |             |        |                    |
| `eth_mining`                                                                      | Eth       |             | ❌      |                    |
| [`eth_coinbase`](#eth-coinbase)                                                   | Eth       | ✔           |        |                    |
| `eth_hashrate`                                                                    | Eth       | N/A         | ❌      | PoW-only           |
| `eth_getUncleCountByBlockHash`                                                    | Eth       | N/A         |        | PoW-only           |
| `eth_getUncleCountByBlockNumber`                                                  | Eth       | N/A         |        | PoW-only           |
| `eth_getUncleByBlockHashAndIndex`                                                 | Eth       | N/A         |        | PoW-only           |
| `eth_getUncleByBlockNumberAndIndex`                                               | Eth       | N/A         |        | PoW-only           |
| [`eth_getProof`](#eth-getProof)                                                   | Eth       | ✔           |        |                    |
| [`eth_subscribe`](#eth-subscribe)                                                 | Websocket | ✔           |        |                    |
| [`eth_unsubscribe`](#eth-unsubscribe)                                             | Websocket | ✔           |        |                    |
| [`personal_importRawKey`](#personal-importrawkey)                                 | Personal  | ✔           | ❌      |                    |
| [`personal_listAccounts`](#personal-listaccounts)                                 | Personal  | ✔           | ❌      |                    |
| [`personal_lockAccount`](#personal-lockaccount)                                   | Personal  | ✔           | ❌      |                    |
| [`personal_newAccount`](#personal-newaccount)                                     | Personal  | ✔           | ❌      |                    |
| [`personal_unlockAccount`](#personal-unlockaccount)                               | Personal  | ✔           | ❌      |                    |
| [`personal_sendTransaction`](#personal-sendtransaction)                           | Personal  | ✔           | ❌      |                    |
| [`personal_sign`](#personal-sign)                                                 | Personal  | ✔           | ❌      |                    |
| [`personal_ecRecover`](#personal-ecrecover)                                       | Personal  | ✔           | ❌      |                    |
| [`personal_initializeWallet`](#personal_initializewallet)                                       | Personal  | ✔           | ❌      ||
| [`personal_unpair`](#personal_unpair)                                       | Personal  | ✔           | ❌      |                    |
| `db_putString`                                                                    | DB        |             |        |                    |
| `db_getString`                                                                    | DB        |             |        |                    |
| `db_putHex`                                                                       | DB        |             |        |                    |
| `db_getHex`                                                                       | DB        |             |        |                    |
| `shh_post`                                                                        | SSH       |             |        |                    |
| `shh_version`                                                                     | SSH       |             |        |                    |
| `shh_newIdentity`                                                                 | SSH       |             |        |                    |
| `shh_hasIdentity`                                                                 | SSH       |             |        |                    |
| `shh_newGroup`                                                                    | SSH       |             |        |                    |
| `shh_addToGroup`                                                                  | SSH       |             |        |                    |
| `shh_newFilter`                                                                   | SSH       |             |        |                    |
| `shh_uninstallFilter`                                                             | SSH       |             |        |                    |
| `shh_getFilterChanges`                                                            | SSH       |             |        |                    |
| `shh_getMessages`                                                                 | SSH       |             |        |                    |
| `admin_addPeer`                                                                   | Admin     |             | ❌      |                    |
| `admin_datadir`                                                                   | Admin     |             | ❌      |                    |
| `admin_nodeInfo`                                                                  | Admin     |             | ❌      |                    |
| `admin_peers`                                                                     | Admin     |             | ❌      |                    |
| `admin_startRPC`                                                                  | Admin     |             | ❌      |                    |
| `admin_startWS`                                                                   | Admin     |             | ❌      |                    |
| `admin_stopRPC`                                                                   | Admin     |             | ❌      |                    |
| `admin_stopWS`                                                                    | Admin     |             | ❌      |                    |
| `clique_getSnapshot`                                                              | Clique    |             |        |                    |
| `clique_getSnapshotAtHash`                                                        | Clique    |             |        |                    |
| `clique_getSigners`                                                               | Clique    |             |        |                    |
| `clique_proposals`                                                                | Clique    |             |        |                    |
| `clique_propose`                                                                  | Clique    |             |        |                    |
| `clique_discard`                                                                  | Clique    |             |        |                    |
| `clique_status`                                                                   | Clique    |             |        |                    |
| `debug_backtraceAt`                                                               | Debug     |             |        |                    |
| `debug_blockProfile`                                                              | Debug     | ✔           |        |                    |
| `debug_cpuProfile`                                                                | Debug     | ✔           |        |                    |
| `debug_dumpBlock`                                                                 | Debug     |             |        |                    |
| `debug_gcStats`                                                                   | Debug     | ✔           |        |                    |
| `debug_getBlockRlp`                                                               | Debug     |             |        |                    |
| `debug_goTrace`                                                                   | Debug     | ✔           |        |                    |
| `debug_freeOSMemory`                                                              | Debug     | ✔           |        |                    |
| `debug_memStats`                                                                  | Debug     | ✔           |        |                    |
| `debug_mutexProfile`                                                              | Debug     | ✔           |        |                    |
| `debug_seedHash`                                                                  | Debug     |             |        |                    |
| `debug_setHead`                                                                   | Debug     |             |        |                    |
| `debug_setBlockProfileRate`                                                       | Debug     | ✔           |        |                    |
| `debug_setGCPercent`                                                              | Debug     | ✔           |        |                    |
| `debug_setMutexProfileFraction`                                                   | Debug     | ✔           |        |                    |
| `debug_stacks`                                                                    | Debug     | ✔           |        |                    |
| `debug_startCPUProfile`                                                           | Debug     | ✔           |        |                    |
| `debug_startGoTrace`                                                              | Debug     | ✔           |        |                    |
| `debug_stopCPUProfile`                                                            | Debug     | ✔           |        |                    |
| `debug_stopGoTrace`                                                               | Debug     | ✔           |        |                    |
| [`debug_traceBlock`](#debug-traceblock)                                                                | Debug     | ✔           |        |                    |
| [`debug_traceBlockByNumber`](#debug-traceblockbynumber)                                                        | Debug     | ✔ |        |                    |
| [`debug_traceBlockByHash`](#debug-traceblockbyhash)                                                          | Debug     | ✔ |        |                    |
| `debug_traceBlockFromFile`                                                        | Debug     |             |        |                    |
| `debug_standardTraceBlockToFile`                                                  | Debug     |             |        |                    |
| `debug_standardTraceBadBlockToFile`                                               | Debug     |             |        |                    |
| [`debug_traceTransaction`](#debug-tracetransaction)                                                          | Debug     | ✔           |        |                    |
| `debug_verbosity`                                                                 | Debug     |             |        |                    |
| `debug_vmodule`                                                                   | Debug     |             |        |                    |
| `debug_writeBlockProfile`                                                         | Debug     | ✔           |        |                    |
| `debug_writeMemProfile`                                                           | Debug     | ✔           |        |                    |
| `debug_writeMutexProfile`                                                         | Debug     | ✔           |        |                    |
| `les_serverInfo`                                                                  | Les       |             |        |                    |
| `les_clientInfo`                                                                  | Les       |             |        |                    |
| `les_priorityClientInfo`                                                          | Les       |             |        |                    |
| `les_addBalance`                                                                  | Les       |             |        |                    |
| `les_setClientParams`                                                             | Les       |             |        |                    |
| `les_setDefaultParams`                                                            | Les       |             |        |                    |
| `les_latestCheckpoint`                                                            | Les       |             |        |                    |
| `les_getCheckpoint`                                                               | Les       |             |        |                    |
| `les_getCheckpointContractAddress`                                                | Les       |             |        |                    |
| [`miner_getHashrate`](#miner-gethashrate)                                         | Miner     | ✔           | ❌      | No-op              |
| [`miner_setExtra`](#miner-setextra)                                               | Miner     | ✔           | ❌      | No-op              |
| [`miner_setGasPrice`](#miner-setgasprice)                                         | Miner     | ✔           | ❌      | Needs node restart |
| [`miner_start`](#miner-start)                                                     | Miner     | ✔           | ❌      | No-op              |
| [`miner_stop`](#miner-stop)                                                       | Miner     | ✔           | ❌      | No-op              |
| [`miner_setGasLimit`](#miner-setgaslimit)                                         | Miner     | ✔           | ❌      | No-op              |
| [`miner_setEtherbase`](#miner-setetherbase)                                       | Miner     | ✔           | ❌      |                    |
| [`txpool_content`](#txpool-content)                                               | TxPool    | ✔           |        |                    |
| [`txpool_inspect`](#txpool-inspect)                                               | TxPool    | ✔           |        |                    |
| [`txpool_status`](#txpool-status)                                                 | TxPool    | ✔           |        |                    |
>>>>>>> d79323fb0f7dd4dd7b9baf4a27f51c8e5ac75472

::: warning
块号可以输入为十六进制字符串，`"earliest"`、`"latest"` 或 `"pending"`。
:::

<<<<<<< HEAD
下面是 RPC 方法、参数和来自命名空间的示例响应的列表。

=======
下面是 JSON_RPC 方法、参数和来自命名空间的示例响应的列表。
>>>>>>> d79323fb0f7dd4dd7b9baf4a27f51c8e5ac75472

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

<<<<<<< HEAD
- Block Number or Block Hash 
=======
- Block Number or 块hash（0x...........） 
>>>>>>> d79323fb0f7dd4dd7b9baf4a27f51c8e5ac75472

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

<<<<<<< HEAD
- Integer of the position in the storage

- Block Number  or Block Hash 
=======
- 存储位置的整数

- Block Number  or 块hash（0x...........） 
>>>>>>> d79323fb0f7dd4dd7b9baf4a27f51c8e5ac75472

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

<<<<<<< HEAD
- Block Number or Block Hash 
=======
- Block Number or 块hash（0x...........） 
>>>>>>> d79323fb0f7dd4dd7b9baf4a27f51c8e5ac75472

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"rpc_getTransactionCount","params":["gx1vt7svcst6982ka4c3ufxmqrs05a9e9eq676tyv", "latest"],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":"0x2"}
```

### `rpc_getBlockTransactionCountByNumber`

返回给定块号的总交易计数。

#### 参数

<<<<<<< HEAD
- Block number
=======
- 块高（0x0）
>>>>>>> d79323fb0f7dd4dd7b9baf4a27f51c8e5ac75472

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"rpc_getBlockTransactionCountByNumber","params":["0x30"],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
 {"jsonrpc":"2.0","id":1,"result":"0x1"}
```

### `rpc_getBlockTransactionCountByHash`

返回给定块哈希的总交易计数。

#### 参数

<<<<<<< HEAD
- Block Hash
=======
- 块hash（0x...........）
>>>>>>> d79323fb0f7dd4dd7b9baf4a27f51c8e5ac75472

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

<<<<<<< HEAD
- Block Number  or Block Hash 
=======
- Block Number  or 块hash（0x...........） 
>>>>>>> d79323fb0f7dd4dd7b9baf4a27f51c8e5ac75472

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

<<<<<<< HEAD
- Block number  or Block Hash 
=======
- 块高（0x0）  or 块hash（0x...........） 
>>>>>>> d79323fb0f7dd4dd7b9baf4a27f51c8e5ac75472

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

<<<<<<< HEAD
-  账户或者合约地址

- 存储位置的整数

- Block Number  or Block Hash 
=======
-  bech32地址或者合约地址

- 存储位置的整数

- Block Number  or 块hash（0x...........） 
>>>>>>> d79323fb0f7dd4dd7b9baf4a27f51c8e5ac75472

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"rpc_getProof","params":["gx1n4m7gafr6cxtj8q4w9mm3480ar4gq3t95k6he9",["0x0000000000000000000000000000000000000000000000000000000000000000","0x0000000000000000000000000000000000000000000000000000000000000001"],"latest"],"id":1}' -H "Content-type:application/json" http://localhost:8545

// Result
{"jsonrpc": "2.0", "id": 1, "result": {"address": "gx1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq8rta3w", "accountProof": ["0xf90211a090dcaf88c40c7bbc95a912cbdde67c175767b31173df9ee4b0d733bfdd511c43a0babe369f6b12092f49181ae04ca173fb68d1a5456f18d20fa32cba73954052bda0473ecf8a7e36a829e75039a3b055e51b8332cbf03324ab4af2066bbd6fbf0021a0bbda34753d7aa6c38e603f360244e8f59611921d9e1f128372fec0d586d4f9e0a04e44caecff45c9891f74f6a2156735886eedf6f1a733628ebc802ec79d844648a0a5f3f2f7542148c973977c8a1e154c4300fec92f755f7846f1b734d3ab1d90e7a0e823850f50bf72baae9d1733a36a444ab65d0a6faaba404f0583ce0ca4dad92da0f7a00cbe7d4b30b11faea3ae61b7f1f2b315b61d9f6bd68bfe587ad0eeceb721a07117ef9fc932f1a88e908eaead8565c19b5645dc9e5b1b6e841c5edbdfd71681a069eb2de283f32c11f859d7bcf93da23990d3e662935ed4d6b39ce3673ec84472a0203d26456312bbc4da5cd293b75b840fc5045e493d6f904d180823ec22bfed8ea09287b5c21f2254af4e64fca76acc5cd87399c7f1ede818db4326c98ce2dc2208a06fc2d754e304c48ce6a517753c62b1a9c1d5925b89707486d7fc08919e0a94eca07b1c54f15e299bd58bdfef9741538c7828b5d7d11a489f9c20d052b3471df475a051f9dd3739a927c89e357580a4c97b40234aa01ed3d5e0390dc982a7975880a0a089d613f26159af43616fd9455bb461f4869bfede26f2130835ed067a8b967bfb80", "0xf90211a0395d87a95873cd98c21cf1df9421af03f7247880a2554e20738eec2c7507a494a0bcf6546339a1e7e14eb8fb572a968d217d2a0d1f3bc4257b22ef5333e9e4433ca012ae12498af8b2752c99efce07f3feef8ec910493be749acd63822c3558e6671a0dbf51303afdc36fc0c2d68a9bb05dab4f4917e7531e4a37ab0a153472d1b86e2a0ae90b50f067d9a2244e3d975233c0a0558c39ee152969f6678790abf773a9621a01d65cd682cc1be7c5e38d8da5c942e0a73eeaef10f387340a40a106699d494c3a06163b53d956c55544390c13634ea9aa75309f4fd866f312586942daf0f60fb37a058a52c1e858b1382a8893eb9c1f111f266eb9e21e6137aff0dddea243a567000a037b4b100761e02de63ea5f1fcfcf43e81a372dafb4419d126342136d329b7a7ba032472415864b08f808ba4374092003c8d7c40a9f7f9fe9cc8291f62538e1cc14a074e238ff5ec96b810364515551344100138916594d6af966170ff326a092fab0a0d31ac4eef14a79845200a496662e92186ca8b55e29ed0f9f59dbc6b521b116fea090607784fe738458b63c1942bba7c0321ae77e18df4961b2bc66727ea996464ea078f757653c1b63f72aff3dcc3f2a2e4c8cb4a9d36d1117c742833c84e20de994a0f78407de07f4b4cb4f899dfb95eedeb4049aeb5fc1635d65cf2f2f4dfd25d1d7a0862037513ba9d45354dd3e36264aceb2b862ac79d2050f14c95657e43a51b85c80", "0xf90171a04ad705ea7bf04339fa36b124fa221379bd5a38ffe9a6112cb2d94be3a437b879a08e45b5f72e8149c01efcb71429841d6a8879d4bbe27335604a5bff8dfdf85dcea00313d9b2f7c03733d6549ea3b810e5262ed844ea12f70993d87d3e0f04e3979ea0b59e3cdd6750fa8b15164612a5cb6567cdfb386d4e0137fccee5f35ab55d0efda0fe6db56e42f2057a071c980a778d9a0b61038f269dd74a0e90155b3f40f14364a08538587f2378a0849f9608942cf481da4120c360f8391bbcc225d811823c6432a026eac94e755534e16f9552e73025d6d9c30d1d7682a4cb5bd7741ddabfd48c50a041557da9a74ca68da793e743e81e2029b2835e1cc16e9e25bd0c1e89d4ccad6980a041dda0a40a21ade3a20fcd1a4abb2a42b74e9a32b02424ff8db4ea708a5e0fb9a09aaf8326a51f613607a8685f57458329b41e938bb761131a5747e066b81a0a16808080a022e6cef138e16d2272ef58434ddf49260dc1de1f8ad6dfca3da5d2a92aaaadc58080", "0xf851808080a009833150c367df138f1538689984b8a84fc55692d3d41fe4d1e5720ff5483a6980808080808080808080a0a319c1c415b271afc0adcb664e67738d103ac168e0bc0b7bd2da7966165cb9518080"], "balance": "0x0", "codeHash": "0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470", "nonce": "0x0", "storageHash": "0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421", "storageProof": [{"key": "0x0000000000000000000000000000000000000000000000000000000000000000", "value": "0x0", "proof": []}, {"key": "0x0000000000000000000000000000000000000000000000000000000000000001", "value": "0x0", "proof": []}]}}
```

<<<<<<< HEAD
=======
## Web3 Methods

### `web3_clientVersion`

获取 web3 客户端版本。

```shell
curl -X POST -H "Content-Type: application/json" http://localhost:8545 --data '{"jsonrpc": "2.0", "id": 42, "method": "web3_clientVersion", "params": []}'

# result
 {"jsonrpc":"2.0","id":1,"result":"Version dev ()\nCompiled at  using Go go1.17.1 (amd64)"}
```


### `web3_sha3`

返回给定数据的 Keccak-256（不是标准化的 SHA3-256）。

#### 参数

1: input `hexutil.Bytes`

- 必填: ✓ 是


```shell
curl -X POST -H "Content-Type: application/json" http://localhost:8545 --data '{"jsonrpc": "2.0", "id": 42, "method": "web3_sha3", "params": [<input>]}'

# result
{"jsonrpc":"2.0","id":1,"result":"0x1b84adea42d5b7d192fd8a61a85b25abe0757e9a65cab1da470258914053823f"}
```


## Net Methods

### `net_version`

返回当前网络 ID。

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"net_version","params":[],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":"520"}
```

### `net_peerCount`

返回当前连接到客户端的对等点数。

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"net_peerCount","params":[],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":23}
```

### `net_listening`

如果客户端正在主动侦听网络连接，则返回。

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"net_listening","params":[],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":true}
```

## Eth Methods

### `eth_protocolVersion`

返回当前的Plug Chain协议版本。

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_protocolVersion","params":[],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":"0x3f"}
```

### `eth_syncing`

同步状态对象可能需要根据 Tendermint 同步协议的详细信息而有所不同。然而，同步结果只是一个布尔值，可以很容易地从 Tendermint 的内部同步状态中导出。

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_syncing","params":[],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":false}
```

### `eth_gasPrice`

返回默认 EVM 面额参数中的当前 gas 价格。

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_gasPrice","params":[],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":"0x0"}
```

### `eth_accounts`

返回所有 eth 帐户的数组。

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_accounts","params":[],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":["0x3b7252d007059ffc82d16d022da3cbf9992d2f70","0xddd64b4712f7c8f1ace3c145c950339eddaf221d","0x0f54f47bf9b8e317b214ccd6a7c3e38b893cd7f0"]}
```

### `eth_blockNumber`

返回当前块高度。

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":"0x66"}
```

### `eth_getBalance`

返回给定帐户地址和块号的帐户余额。

#### 参数

- EIP-55 地址（0x..........）

- 块高或者块hash（0x）

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_getBalance","params":["0x0f54f47bf9b8e317b214ccd6a7c3e38b893cd7f0", "0x0"],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":"0x36354d5575577c8000"}
```

### `eth_getStorageAt`

返回给定帐户地址的存储地址。

#### 参数

- EIP-55 地址（0x..........）

- 存储位置的整数

- 块高或者块hash

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_getStorageAt","params":["0x0f54f47bf9b8e317b214ccd6a7c3e38b893cd7f0", "0", `"latest"`],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":"0x0000000000000000000000000000000000000000000000000000000000000000"}
```

### `eth_getTransactionCount`

返回给定帐户地址和块号的总交易量。

#### 参数

- EIP-55 地址（0x..........）

- 块高或者块hash

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_getTransactionCount","params":["0x7bf7b17da59880d9bcca24915679668db75f9397", "0x0"],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":"0x8"}
```

### `eth_getBlockTransactionCountByNumber`

返回给定块号的总交易计数。

#### 参数

- 块高（0x0）

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_getBlockTransactionCountByNumber","params":["0x1"],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
 {"jsonrpc":"2.0","id":1,"result":{"difficulty":null,"extraData":"0x0","gasLimit":"0xffffffff","gasUsed":"0x0","hash":"0x8101cc04aea3341a6d4b3ced715e3f38de1e72867d6c0db5f5247d1a42fbb085","logsBloom":"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","miner":"0x0000000000000000000000000000000000000000","nonce":null,"number":"0x17d","parentHash":"0x70445488069d2584fea7d18c829e179322e2b2185b25430850deced481ca2e77","sha3Uncles":null,"size":"0x1df","stateRoot":"0x269bb17fe7adb8dd5f15f57b717979f82078d6b7a675c1ba1b0da2d27e415fcc","timestamp":"0x5f5ba97c","totalDifficulty":null,"transactions":[],"transactionsRoot":"0x","uncles":[]}}
```

### `eth_getBlockTransactionCountByHash`

返回给定块哈希的总交易计数。

#### 参数

- 块hash（0x...........）

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_getBlockTransactionCountByHash","params":["0x8101cc04aea3341a6d4b3ced715e3f38de1e72867d6c0db5f5247d1a42fbb085"],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":"0x3"}
```

### `eth_getCode`

返回给定帐户地址和块号的代码。

#### 参数

- EIP-55 地址（0x..........）

- 块高或者块hash

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_getCode","params":["0x7bf7b17da59880d9bcca24915679668db75f9397", "0x0"],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":"0xef616c92f3cfc9e92dc270d6acff9cea213cecc7020a76ee4395af09bdceb4837a1ebdb5735e11e7d3adb6104e0c3ac55180b4ddf5e54d022cc5e8837f6a4f971b"}
```

### `eth_sign`

`sign` 方法计算以太坊特定的签名：`sign(keccak256("\x19Ethereum Signed Message:\n" + len(message) + message)))`。

通过在消息中添加前缀，可以将计算出的签名识别为以太坊特定的签名。这可以防止恶意 DApp 可以签署任意数据（例如交易）并使用签名来冒充受害者的滥用行为。

::: warning
用于签名的地址必须解锁。
:::

#### 参数

- 账户地址

- 要签名的信息

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_sign","params":["0x3b7252d007059ffc82d16d022da3cbf9992d2f70", "0xdeadbeaf"],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":"0x909809c76ed2a5d38733de39207d0f411222b9b49c64a192bf649cb13f63f37b45acb4f6939facb4f1c277bc70fb00407564140c0f18600ac44388f2c1dfd1dc1b"}
```

### `eth_sendTransaction`

将交易从给定帐户发送到给定帐户。

#### 参数

- 交易对象:

    `from`: `DATA`, 20 Bytes - 发送交易的地址。

    `to`: `DATA`, 20 Bytes - （创建新合约时可选）交易指向的地址。

    `gas`: QUANTITY - （可选，默认值：90000）为交易执行提供的气体的整数。它将返回未使用的气体。

    `gasPrice`: QUANTITY - (可选，默认值：待确定）气体的整数用于每个付费气体的价格

    `value`: QUANTITY - 与此交易一起发送的值

    `data`: `DATA` - 合约的编译代码或调用的方法签名和编码参数的哈希值。详见以太坊合约 ABI

    `nonce`: QUANTITY - （可选）nonce 的整数。这允许覆盖您自己的使用相同 nonce 的待处理事务。

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_sendTransaction","params":[{"from":"0x3b7252d007059ffc82d16d022da3cbf9992d2f70", "to":"0x0f54f47bf9b8e317b214ccd6a7c3e38b893cd7f0", "value":"0x16345785d8a0000", "gasLimit":"0x5208", "gasPrice":"0x55ae82600"}],"id":1}'  -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":"0x33653249db68ebe5c7ae36d93c9b2abc10745c80a72f591e296f598e2d4709f6"}
```

### `eth_sendRawTransaction`

为已签名的交易创建新的消息调用交易或合同创建。
您可以使用 [`personal_sign`](#personal-sign) 方法获取签名的交易数据。

#### 参数

- 签署的交易数据

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_sendRawTransaction","params":["0xf9ff74c86aefeb5f6019d77280bbb44fb695b4d45cfe97e6eed7acd62905f4a85034d5c68ed25a2e7a8eeb9baf1b8401e4f865d92ec48c1763bf649e354d900b1c"],"id":1}'  -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":"0x0000000000000000000000000000000000000000000000000000000000000000"}
```

### `eth_call`

立即执行新的消息调用，而不在区块链上创建交易。

#### 参数

- 交易对象:

    `from`: `DATA`, 20 Bytes - （可选）发送交易的地址。

    `to`: `DATA`, 20 Bytes - 交易被定向到的地址。

    `gas`: QUANTITY - 为交易执行提供的气体。 eth_call 消耗零气体，但某些执行可能需要此参数。

    `gasPrice`: QUANTITY - 用于每个付费gas的gasPrice

    `value`: QUANTITY - 与此交易一起发送的值

    `data`: `DATA` - 可选）方法签名和编码参数的哈希。有关详细信息，请参阅 Solidity 文档中的以太坊合约 ABI

- 块高（0x0）  or 块hash（0x...........）

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_call","params":[{"from":"0x3b7252d007059ffc82d16d022da3cbf9992d2f70", "to":"0xddd64b4712f7c8f1ace3c145c950339eddaf221d", "gas":"0x5208", "gasPrice":"0x55ae82600", "value":"0x16345785d8a0000", "data": "0xd46e8dd67c5d32be8d46e8dd67c5d32be8058bb8eb970870f072445675058bb8eb970870f072445675"}, "0x0"],"id":1}'  -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":"0x"}
```

### `eth_estimateGas`

返回发送交易所需气体的估计值。

#### 参数

- 交易对象:

    `from`: `DATA`, 20 Bytes - 发送交易的地址。.

    `to`: `DATA`, 20 Bytes - （创建新合约时可选）交易指向的地址。

    `value`: `QUANTITY` - 与此交易一起发送的值

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_estimateGas","params":[{"from":"0x0f54f47bf9b8e317b214ccd6a7c3e38b893cd7f0", "to":"0x3b7252d007059ffc82d16d022da3cbf9992d2f70", "value":"0x16345785d8a00000"}],"id":1}'  -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":"0x1199b"}
```

### `eth_getBlockByNumber`

按块号返回有关块的信息。

#### 参数

- 块号

-  true|false  如果为 true，则返回完整的交易对象，如果为 false，则仅返回交易的哈希值。

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["0x1", false],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":{"difficulty":null,"extraData":"0x0","gasLimit":"0xffffffff","gasUsed":null,"hash":"0xabac6416f737a0eb54f47495b60246d405d138a6a64946458cf6cbeae0d48465","logsBloom":"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","miner":"0x0000000000000000000000000000000000000000","nonce":null,"number":"0x1","parentHash":"0x","sha3Uncles":null,"size":"0x9b","stateRoot":"0x","timestamp":"0x5f5bd3e5","totalDifficulty":null,"transactions":[],"transactionsRoot":"0x","uncles":[]}}
```

### `eth_getBlockByHash`

返回给定上面命令中的哈希值和一个布尔值，锁定信息。

#### 参数 

- 块hash.

- true|false 如果为 true，则返回完整的交易对象，如果为 false，则仅返回交易的哈希值。

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_getBlockByHash","params":["0x1b9911f57c13e5160d567ea6cf5b545413f96b95e43ec6e02787043351fb2cc4", false],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":{"difficulty":null,"extraData":"0x0","gasLimit":"0xffffffff","gasUsed":null,"hash":"0x1b9911f57c13e5160d567ea6cf5b545413f96b95e43ec6e02787043351fb2cc4","logsBloom":"0x00000000100000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000040000000000000000000000000200000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000000000000000000000004000000000000002000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000000000000000","miner":"0x0000000000000000000000000000000000000000","nonce":null,"number":"0xc","parentHash":"0x404e58f31a9ede1b614b98701d6b0fbf1450f186842dbcf6426dd16811a5ca0d","sha3Uncles":null,"size":"0x307","stateRoot":"0x599ccdb111fc62c6398dc39be957df8e97bf8ab72ce6c06ff10641a92b754627","timestamp":"0x5f5fdbbd","totalDifficulty":null,"transactions":["0xae64961cb206a9773a6e5efeb337773a6fd0a2085ce480a174135a029afea615"],"transactionsRoot":"0x4764dba431128836fa919b83d314ba9cc000e75f38e1c31a60484409acea777b","uncles":[]}}
```

### `eth_getTransactionByHash`

返回给定以太坊 tx 的交易细节。

#### 参数

- 交易的哈希

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_getTransactionByHash","params":["0xec5fa15e1368d6ac314f9f64118c5794f076f63c02e66f97ea5fe1de761a8973"],"id":1}' -H "Content-Type: application/json" http://localhost:8545
 
// Result
{"jsonrpc":"2.0","id":1,"result":{"blockHash":"0x7a7398cc11d9c4c8e6f53e0c73824297aceafdab62db9e4b867a0da694384864","blockNumber":"0x188","from":"0x3b7252d007059ffc82d16d022da3cbf9992d2f70","gas":"0x147ee","gasPrice":"0x3b9aca00","hash":"0xec5fa15e1368d6ac314f9f64118c5794f076f63c02e66f97ea5fe1de761a8973","input":"0x6dba746c","nonce":"0x18","to":"0xa655256f589060437e5ffe2246dec385d040f148","transactionIndex":"0x0","value":"0x0","v":"0xa96","r":"0x6db399d694a452fb4106419140a6e5dbbe6817743a0f6f695a651e6576e59a5e","s":"0x25dd6ab1f936d0280d2fed0caeb0ebe5b9a46de6d8cb08ad8fd2c88deb55fc31"}}
```

### `eth_getTransactionByBlockHashAndIndex`

返回给定块哈希和交易索引的交易详细信息。

#### 参数

- 块哈希.

- 交易索引位置.

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_getTransactionByBlockHashAndIndex","params":["0x1b9911f57c13e5160d567ea6cf5b545413f96b95e43ec6e02787043351fb2cc4", "0x0"],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":{"blockHash":"0x1b9911f57c13e5160d567ea6cf5b545413f96b95e43ec6e02787043351fb2cc4","blockNumber":"0xc","from":"0xddd64b4712f7c8f1ace3c145c950339eddaf221d","gas":"0x4c4b40","gasPrice":"0x3b9aca00","hash":"0xae64961cb206a9773a6e5efeb337773a6fd0a2085ce480a174135a029afea615","input":"0x4f2be91f","nonce":"0x0","to":"0x439c697e0742a0ddb124a376efd62a72a94ac35a","transactionIndex":"0x0","value":"0x0","v":"0xa96","r":"0xced57d973e58b0f634f776d57daf41d3d3387ceb347a3a72ca0746e5ec2b709e","s":"0x384e89e209a5eb147a2bac3a4e399507400ac7b29cd155531f9d6203a89db3f2"}}
```

### `eth_getTransactionReceipt`

通过交易哈希返回交易的收据。

注意: Tendermint 的 Tx Code 和 Ethereum 收据状态被切换：
|         | Tendermint | Ethereum |
|---------|------------|----------|
| Success | 0          | 1        |
| Fail    | 1          | 0        |

#### 参数

- 交易的哈希

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_getTransactionReceipt","params":["0xae64961cb206a9773a6e5efeb337773a6fd0a2085ce480a174135a029afea614"],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":{"blockHash":"0x1b9911f57c13e5160d567ea6cf5b545413f96b95e43ec6e02787043351fb2cc4","blockNumber":"0xc","contractAddress":"0x0000000000000000000000000000000000000000","cumulativeGasUsed":null,"from":"0xddd64b4712f7c8f1ace3c145c950339eddaf221d","gasUsed":"0x5289","logs":[{"address":"0x439c697e0742a0ddb124a376efd62a72a94ac35a","topics":["0x64a55044d1f2eddebe1b90e8e2853e8e96931cefadbfa0b2ceb34bee36061941"],"data":"0x0000000000000000000000000000000000000000000000000000000000000002","blockNumber":"0xc","transactionHash":"0xae64961cb206a9773a6e5efeb337773a6fd0a2085ce480a174135a029afea615","transactionIndex":"0x0","blockHash":"0x0000000000000000000000000000000000000000000000000000000000000000","logIndex":"0x0","removed":false},{"address":"0x439c697e0742a0ddb124a376efd62a72a94ac35a","topics":["0x938d2ee5be9cfb0f7270ee2eff90507e94b37625d9d2b3a61c97d30a4560b829"],"data":"0x0000000000000000000000000000000000000000000000000000000000000002","blockNumber":"0xc","transactionHash":"0xae64961cb206a9773a6e5efeb337773a6fd0a2085ce480a174135a029afea615","transactionIndex":"0x0","blockHash":"0x0000000000000000000000000000000000000000000000000000000000000000","logIndex":"0x1","removed":false}],"logsBloom":"0x00000000100000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000040000000000000000000000000200000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000000000000000000000004000000000000002000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000000000000000","status":"0x1","to":"0x439c697e0742a0ddb124a376efd62a72a94ac35a","transactionHash":"0xae64961cb206a9773a6e5efeb337773a6fd0a2085ce480a174135a029afea615","transactionIndex":"0x0"}}
```

### `eth_newFilter`

使用某种主题创建新过滤器。

#### 参数

- 交易哈希

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_newFilter","params":[{"topics":["0x0000000000000000000000000000000000000000000000000000000012341234"]}],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":"0xdc714a4a2e3c39dc0b0b84d66a3ccb00"}
```

### `eth_newBlockFilter`

在节点中创建一个过滤器，以在新块到达时通知。

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_newBlockFilter","params":[],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":"0x3503de5f0c766c68f78a03a3b05036a5"}
```

### `eth_newPendingTransactionFilter`

在节点中创建一个过滤器，以在新的待处理事务到达时发出通知。

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_newPendingTransactionFilter","params":[],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":"0x9daacfb5893d946997d3801ea18e9902"}
```

### `eth_uninstallFilter`

删除具有给定过滤器 ID 的过滤器。如果过滤器已成功卸载，则返回 true，否则返回 false。

#### 参数

- 过滤器id

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_uninstallFilter","params":["0xb91b6608b61bf56288a661a1bd5eb34a"],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":true}
```

### `eth_getFilterChanges`

过滤器的轮询方法，它返回自上次轮询以来发生的日志数组。

#### 参数

- 过滤器id

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_getFilterChanges","params":["0x127e9eca4f7751fb4e5cb5291ad8b455"],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":["0xc6f08d183a81e149896fc5317c872f9092068e88e956ca1864e9bd4c81c09b44","0x3ca6dfb5be15549d721d1b3d10c1bec50ed6217c9ac7b61df361fac9692a27e5","0x776fffac134171acb1ebf2e59856625501ad5ccc5c4c8fe0359e0d4dff8919f2","0x84123103704dbd738c089276ab2b04b5936330b24f6e78453c4ba8bf4848aaf9","0xffddbe5bd8e8aa41e44002daa9ea89ade9e6980a0d83f51d104cf16498827eca","0x53430e49963e8ae32605d8f22dec2e757a691e6436d593854ca4d9383eeab86a","0x975948058c9351a91fbec332ca00dda39d1a919f5f16b996a4c7e30c38ba423b","0x619e37e32024c8efef7f7220e6caff4ee1d682ea78b2ac91e0a6b30850dc0677","0x31a5d985a40d08303ac68000ce008df512bcd1a911c497415c97f0624b4a271a","0x91dcf1fce4503a8dbb3e6fb61073f25cd31d69c766ecba639fefde4436e59d07","0x606d9e0143cfdb410a6812c590a8135b5c6b5c59eec26d760d5cd930aa47257d","0xd3c00b859b29b20ba654415eef648ef58251389c73a138580db87675b0d5465f","0x954391f0eb50888be90489898016ebb54f750f612f3adec2a00854955d5e52d8","0x698905f06aff921a9e9fcef39b8b0d107747c3e6204d2ea79cf4c12debf8d253","0x9fcafec5721938a06eb8e2951ede4b6ef8fae54a8c8f85f3166ec9782a0032b5","0xaec6d3364e47a5716ba69e4705f3c705d017f81298859589591183bfea87be7a","0x91bf2ee13319b6eaca96ed89c126437b66c4df1b13560c6a9bb18556ee3b7e1f","0x4f426dc1fc0ea8149052033065b237892d2d34927b2d558ab50c5a7fb98d6e79","0xdd809fb07e5aab638fef5311371b4e2b27c9c9a6183fde0cdd2b7724f6d2a89b","0x7e12fc92ab953e233a304959a2a8474d96195e71efd9388fdceb1326a577811a","0x30618ef6b490c3cc9979c47163459db37c1a1e0aa5793c56accd417f9d89973b","0x614609f06ee24bae7408e45895b1a25e6b19a8159aeea7a95c9d1339d9ba286f","0x115ddc6d533620040791d241f01f1c5ae3d9d1a8f64b15af5e9793e4d9096e22","0xb7458c9323beeca2cd54f32a6af5671f3cd5a7a251aed9d82bdd6ebe5f56305b","0x573dd48a5ba7bf4cc3d49597cd7419f75ecc9897258f1ebadebd670446d0d358","0xcb6670918439f9698413b53f3b5336d82ca4be152fdefaacf45e052fff6262fc","0xf3fe2a8945abafd269ab97bfdc80b3dbff2202ffdce59a227f952874b966b230","0x989980707007533cc0840a079f77f261a2e818abae1a1ffd3af02f3fff1d35fd","0x886b6ae365fec996be8a9a2c31cf4cda97ff8352908be2c83f17abd66ef1591e","0xfd90df68706ef95a62b317de93d6899a9bd6c80416e42d007f5c30fcdedfce24","0x7af8491fbb0373886d9032bb74e0ef52ed9e100f260b79bd15f46126b38cbede","0x91d1e2cd55533cf7dd5de86c9aa73295e811b1279be193d429bbd6ba83810e16","0x6b65b3128c2104005a04923288fe2aa33a2477a4962bef70532f94cab582f2a7"]}
```

### `eth_getFilterLogs`

返回具有给定 id 的所有与过滤器匹配的日志的数组。

#### 参数

- `QUANTITY` - 过滤器id

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_getFilterLogs","params":["0x127e9eca4f7751fb4e5cb5291ad8b455"],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"error":{"code":-32000,"message":"filter 0x35b64c227ce30e84fc5c7bd347be380e doesn't have a LogsSubscription type: got 5"}} 
```

### `eth_getLogs`

返回与给定过滤器对象匹配的所有日志的数组。

#### 参数

- 交易对象:

    `fromBlock`: `QUANTITY|TAG` - （可选，默认值：`"latest"`）整数块编号，或 `"latest"` 表示最后一个开采的区块或 `"pending"`，`"earliest"` 表示未开采尚未开采的交易。
    `toBlock`: `QUANTITY|TAG` - （可选，默认值：`"latest"`）整数块编号，或 `"latest"` 表示最后一个开采的区块或 `"pending"`，`"earliest"` 表示未开采尚未开采的交易。
    `address`: `DATA|Array`, 20 Bytes - （可选）合约地址或日志应源自的地址列表。

    `topics`: Array of `DATA`, - （可选）32字节`DATA`主题数组。主题是顺序相关的。每个主题也可以是带有“或”选项的“DATA”数组。

    `blockhash`: 可选，未来）随着 [EIP-234](https://eips.ethereum.org/EIPS/eip-234) 的加入，`blockHash` 将成为一个新的过滤器选项，用于限制返回的日志到具有 32 字节哈希 `blockHash` 的单个块。使用 `blockHash` 等价于 `fromBlock` = `toBlock` = 哈希值为 `blockHash` 的区块号。如果 `blockHash` 出现在过滤条件中，则 `fromBlock` 和 `toBlock` 都不允许。
```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_getLogs","params":[{"topics":["0x775a94827b8fd9b519d36cd827093c664f93347070a554f65e4a6f56cd738898","0x0000000000000000000000000000000000000000000000000000000000000011"], "fromBlock":`"latest"`}],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":[]}
```

### eth_coinbase

返回将发送挖矿奖励的账户。

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_coinbase","params":[],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":"0x7cB61D4117AE31a12E393a1Cfa3BaC666481D02E"}
```

### eth_getProof

返回指定帐户的帐户值和存储值，包括默克尔证明。

#### 参数

- EIP-55 地址或者合约地址

- 存储位置的整数

- 块高或者块hash

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_getProof","params":["0x1234567890123456789012345678901234567890",["0x0000000000000000000000000000000000000000000000000000000000000000","0x0000000000000000000000000000000000000000000000000000000000000001"],`"latest"`],"id":1}' -H "Content-type:application/json" http://localhost:8545

// Result
{"jsonrpc": "2.0", "id": 1, "result": {"address": "0x1234567890123456789012345678901234567890", "accountProof": ["0xf90211a090dcaf88c40c7bbc95a912cbdde67c175767b31173df9ee4b0d733bfdd511c43a0babe369f6b12092f49181ae04ca173fb68d1a5456f18d20fa32cba73954052bda0473ecf8a7e36a829e75039a3b055e51b8332cbf03324ab4af2066bbd6fbf0021a0bbda34753d7aa6c38e603f360244e8f59611921d9e1f128372fec0d586d4f9e0a04e44caecff45c9891f74f6a2156735886eedf6f1a733628ebc802ec79d844648a0a5f3f2f7542148c973977c8a1e154c4300fec92f755f7846f1b734d3ab1d90e7a0e823850f50bf72baae9d1733a36a444ab65d0a6faaba404f0583ce0ca4dad92da0f7a00cbe7d4b30b11faea3ae61b7f1f2b315b61d9f6bd68bfe587ad0eeceb721a07117ef9fc932f1a88e908eaead8565c19b5645dc9e5b1b6e841c5edbdfd71681a069eb2de283f32c11f859d7bcf93da23990d3e662935ed4d6b39ce3673ec84472a0203d26456312bbc4da5cd293b75b840fc5045e493d6f904d180823ec22bfed8ea09287b5c21f2254af4e64fca76acc5cd87399c7f1ede818db4326c98ce2dc2208a06fc2d754e304c48ce6a517753c62b1a9c1d5925b89707486d7fc08919e0a94eca07b1c54f15e299bd58bdfef9741538c7828b5d7d11a489f9c20d052b3471df475a051f9dd3739a927c89e357580a4c97b40234aa01ed3d5e0390dc982a7975880a0a089d613f26159af43616fd9455bb461f4869bfede26f2130835ed067a8b967bfb80", "0xf90211a0395d87a95873cd98c21cf1df9421af03f7247880a2554e20738eec2c7507a494a0bcf6546339a1e7e14eb8fb572a968d217d2a0d1f3bc4257b22ef5333e9e4433ca012ae12498af8b2752c99efce07f3feef8ec910493be749acd63822c3558e6671a0dbf51303afdc36fc0c2d68a9bb05dab4f4917e7531e4a37ab0a153472d1b86e2a0ae90b50f067d9a2244e3d975233c0a0558c39ee152969f6678790abf773a9621a01d65cd682cc1be7c5e38d8da5c942e0a73eeaef10f387340a40a106699d494c3a06163b53d956c55544390c13634ea9aa75309f4fd866f312586942daf0f60fb37a058a52c1e858b1382a8893eb9c1f111f266eb9e21e6137aff0dddea243a567000a037b4b100761e02de63ea5f1fcfcf43e81a372dafb4419d126342136d329b7a7ba032472415864b08f808ba4374092003c8d7c40a9f7f9fe9cc8291f62538e1cc14a074e238ff5ec96b810364515551344100138916594d6af966170ff326a092fab0a0d31ac4eef14a79845200a496662e92186ca8b55e29ed0f9f59dbc6b521b116fea090607784fe738458b63c1942bba7c0321ae77e18df4961b2bc66727ea996464ea078f757653c1b63f72aff3dcc3f2a2e4c8cb4a9d36d1117c742833c84e20de994a0f78407de07f4b4cb4f899dfb95eedeb4049aeb5fc1635d65cf2f2f4dfd25d1d7a0862037513ba9d45354dd3e36264aceb2b862ac79d2050f14c95657e43a51b85c80", "0xf90171a04ad705ea7bf04339fa36b124fa221379bd5a38ffe9a6112cb2d94be3a437b879a08e45b5f72e8149c01efcb71429841d6a8879d4bbe27335604a5bff8dfdf85dcea00313d9b2f7c03733d6549ea3b810e5262ed844ea12f70993d87d3e0f04e3979ea0b59e3cdd6750fa8b15164612a5cb6567cdfb386d4e0137fccee5f35ab55d0efda0fe6db56e42f2057a071c980a778d9a0b61038f269dd74a0e90155b3f40f14364a08538587f2378a0849f9608942cf481da4120c360f8391bbcc225d811823c6432a026eac94e755534e16f9552e73025d6d9c30d1d7682a4cb5bd7741ddabfd48c50a041557da9a74ca68da793e743e81e2029b2835e1cc16e9e25bd0c1e89d4ccad6980a041dda0a40a21ade3a20fcd1a4abb2a42b74e9a32b02424ff8db4ea708a5e0fb9a09aaf8326a51f613607a8685f57458329b41e938bb761131a5747e066b81a0a16808080a022e6cef138e16d2272ef58434ddf49260dc1de1f8ad6dfca3da5d2a92aaaadc58080", "0xf851808080a009833150c367df138f1538689984b8a84fc55692d3d41fe4d1e5720ff5483a6980808080808080808080a0a319c1c415b271afc0adcb664e67738d103ac168e0bc0b7bd2da7966165cb9518080"], "balance": "0x0", "codeHash": "0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470", "nonce": "0x0", "storageHash": "0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421", "storageProof": [{"key": "0x0000000000000000000000000000000000000000000000000000000000000000", "value": "0x0", "proof": []}, {"key": "0x0000000000000000000000000000000000000000000000000000000000000001", "value": "0x0", "proof": []}]}}
```


## WebSocket Methods

websockets相关信息 [events](./events.md)

### `eth_subscribe`

使用 JSON-RPC 通知订阅。这允许客户端等待事件而不是轮询它们。

它通过订阅特定事件来工作。该节点将返回一个订阅 ID。对于与订阅匹配的每个事件，带有相关数据的通知与订阅 ID 一起发送。

#### 参数

- 订阅名称

- 可选参数

```json
// Request
{"id": 1, "method": "eth_subscribe", "params": ["newHeads", {"includeTransactions": true}]}

// Result
< {"jsonrpc":"2.0","result":"0x34da6f29e3e953af4d0c7c58658fd525","id":1}
```

### `eth_unsubscribe`

使用订阅 ID 取消订阅事件

#### 参数

- Subscription ID

```json
// Request
{"id": 1, "method": "eth_unsubscribe", "params": ["0x34da6f29e3e953af4d0c7c58658fd525"]}

// Result
{"jsonrpc":"2.0","result":true,"id":1}
```

## Personal Methods

### `personal_importRawKey`

::: tip
**Private**: 需要身份验证。
:::

将给定的未加密私钥（十六进制编码字符串）导入密钥存储，并使用密码对其进行加密。

#### 参数 (2)

**1:** privkey `string`

- Required: ✓ Yes

**2:** password `string`

- Required: ✓ Yes



```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"personal_importRawKey","params":["c5bd76cd0cd948de17a31261567d219576e992d9066fe1a6bca97496dec634e2c8e06f8949773b300b9f73fabbbc7710d5d6691e96bcf3c9145e15daf6fe07b9", "the key is this"],"id":1}' -H "Content-Type: application/json" http://localhost:8545

```

### `personal_listAccounts`

::: tip
**Private**: 需要身份验证。
:::

返回此节点管理的帐户的地址列表。

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"personal_listAccounts","params":[],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":["0x3b7252d007059ffc82d16d022da3cbf9992d2f70","0xddd64b4712f7c8f1ace3c145c950339eddaf221d","0x0f54f47bf9b8e317b214ccd6a7c3e38b893cd7f0"]}
```

### `personal_lockAccount`

::: tip
**Private**: 需要身份验证。
:::

从内存中删除具有给定地址的私钥。该帐户不能再用于发送交易。

#### 参数

- EIP-55 地址（0x..........）

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"personal_lockAccount","params":["0x0f54f47bf9b8e317b214ccd6a7c3e38b893cd7f0"],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":true}
```

### `personal_newAccount`

::: tip
**Private**: 需要身份验证。
:::

生成新的私钥并将其存储在密钥存储目录中。密钥文件使用给定的密码进行加密。返回新帐户的地址。

#### 参数

- 密码

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"personal_newAccount","params":["This is the passphrase"],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":"0xf0e4086ad1c6aab5d42161d5baaae2f9ad0571c0"}
```

### `personal_unlockAccount`

::: tip
**Private**: 需要身份验证。
:::

从密钥库中使用给定地址解密密钥。

使用 JavaScript 控制台时，密码和解锁持续时间都是可选的。未加密的密钥将保存在内存中，直到解锁持续时间到期。如果解锁持续时间默认为 300 秒。零秒的显式持续时间会解锁密钥，直到 geth 退出。

该帐户在解锁时可与 [`eth_sign`](#eth-sign) 和 [`eth_sendTransaction`](#eth-sendtransaction) 一起使用。

#### 参数

- 账户

- 密码

- 期间

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"personal_unlockAccount","params":["0x0f54f47bf9b8e317b214ccd6a7c3e38b893cd7f0", "secret passphrase", 30],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":true}
```

### `personal_sendTransaction`

::: tip
**Private**: 需要身份验证。
:::

验证给定的密码并提交交易。

该事务与 [`eth_sendTransaction`](#eth-sendtransaction) 的参数相同，并且包含 `from` 地址。如果密码可用于解密属于 tx.from 的私钥，则交易被验证、签名并发送到网络上。

:::warning
该账户在节点中未全局解锁，不能用于其他RPC调用。
:::

#### 参数

- 交易对象:

    `from`: `DATA`, 20 Bytes - 发送交易的地址。.

    `to`: `DATA`, 20 Bytes - （创建新合约时可选）交易指向的地址。

    `value`: QUANTITY - 与此交易一起发送的值

- Passphrase

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"personal_sendTransaction","params":[{"from":"0x3b7252d007059ffc82d16d022da3cbf9992d2f70","to":"0xddd64b4712f7c8f1ace3c145c950339eddaf221d", "value":"0x16345785d8a0000"}, "passphrase"],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":"0xd2a31ec1b89615c8d1f4d08fe4e4182efa4a9c0d5758ace6676f485ea60e154c"}
```

### `personal_sign`

::: tip
**Private**: 需要身份验证。
:::

sign 方法计算以太坊特定的签名：`sign(keccack256("\x19Ethereum Signed Message:\n" + len(message) + message)))`，

#### 参数

- 信息

- 账户

- 密码

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"personal_sign","params":["0xdeadbeaf", "0x3b7252d007059ffc82d16d022da3cbf9992d2f70", "password"],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":"0xf9ff74c86aefeb5f6019d77280bbb44fb695b4d45cfe97e6eed7acd62905f4a85034d5c68ed25a2e7a8eeb9baf1b8401e4f865d92ec48c1763bf649e354d900b1c"}
```

### `personal_ecRecover`

::: tip
**Private**: 需要身份验证。
:::

`ecRecover` 返回与用于计算 [`personal_sign`](#personal-sign) 中签名的私钥关联的地址。

#### 参数

- 信息

- 从 [`personal_sign`](#personal-sign) 返回的签名

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"personal_ecRecover","params":["0xdeadbeaf", "0xf9ff74c86aefeb5f6019d77280bbb44fb695b4d45cfe97e6eed7acd62905f4a85034d5c68ed25a2e7a8eeb9baf1b8401e4f865d92ec48c1763bf649e354d900b1c"],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":"0x3b7252d007059ffc82d16d022da3cbf9992d2f70"}
```

### `personal_initializeWallet`

::: tip
**Private**: 需要身份验证。
:::

通过生成并返回一个新的私钥，在提供的 URL 处初始化一个新的钱包。

#### 参数 (1)

参数必须由位置给出。

1: url `string`

-  Required: ✓ Yes


```shell
curl -X POST -H "Content-Type: application/json" http://localhost:8545 --data '{"jsonrpc": "2.0", "id": 42, "method": "personal_initializeWallet", "params": [<url>]}'
```

### `personal_unpair`

::: tip
**Private**: 需要身份验证。
:::

Unpair 删除钱包和节点之间的配对。

#### 参数 (2)

- URL

- 配对密码


```shell
curl -X POST -H "Content-Type: application/json" http://localhost:8545 --data '{"jsonrpc": "2.0", "id": 42, "method": "personal_unpair", "params": [<url>, <pin>]}'
```


## Debug Methods

### `debug_traceTransaction`

`traceTransaction` 调试方法将尝试以与在网络上执行事务完全相同的方式运行事务。 在最终尝试执行与给定哈希相对应的事务之前，它将重放任何可能在此之前执行的事务。


#### 参数

- 跟踪配置

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"debug_traceTransaction","params":["0xddecdb13226339681372b44e01df0fbc0f446fca6f834b2de5ecb1e569022ec8", {"tracer": "{data: [], fault: function(log) {}, step: function(log) { if(log.op.toString() == \"CALL\") this.data.push(log.stack.peek(0)); }, result: function() { return this.data; }}"}],"id":1}' -H "Content-Type: application/json" http://localhost:8545

//Result
["68410", "51470"]
```

### `debug_traceBlockByNumber`

`traceBlockByNumber` 端点接受一个块号，并将重播数据库中已经存在的块。

#### 参数

- 跟踪配置

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"debug_traceBlockByNumber","params":["0xe", {"tracer": "{data: [], fault: function(log) {}, step: function(log) { if(log.op.toString() == \"CALL\") this.data.push(log.stack.peek(0)); }, result: function() { return this.data; }}"}],"id":1}' -H "Content-Type: application/json" http://localhost:8545

//Result
{"jsonrpc":"2.0","id":1,"result":[{"result":["68410", "51470"]}]}
```


## Miner Methods

### `miner_getHashrate`

::: tip
**Private**: 需要身份验证。
:::

获取以 H/s 为单位的哈希率（每秒哈希操作数）。

::: warning
工作证明特定。此端点始终返回"0"。
:::

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"miner_setGasPrice","params":[],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":0}
```

### `miner_setExtra`

::: tip
**Private**: 需要身份验证。
:::

设置验证者在提议区块时可以包含的额外数据。上限为 32 个字节。

::: warning
不受支持。此端点始终返回错误
:::

#### 参数

- Data

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"miner_setExtra","params":["data"],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":false}
```

### `miner_setGasPrice`

::: tip
**Private**: 需要身份验证。
:::

设置用于接受交易的最低汽油价格。低于此限制的任何交易都将被排除在验证者区块提议过程之外。

此方法在调用后需要重新启动"节点"，因为它会更改配置文件。

确保您的 `plugchaind start` 调用未使用标志 `minimum-gas-prices`，因为将使用此值而不是配置文件中设置的值。

#### 参数

- Hex Gas Price

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"miner_setGasPrice","params":["0x0"],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":true}
```

### `miner_start`

::: tip
**Private**: 需要身份验证。
:::

使用给定的线程数启动 CPU 验证过程。

::: warning
不受支持。此端点始终返回错误
:::

#### 参数

- Hex Number of threads

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"miner_start","params":["0x1"],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":false}
```

### `miner_stop`

::: tip
**Private**: 需要身份验证。
:::

停止验证操作。

::: warning
不受支持。此端点始终执行无操作。
:::

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"miner_stop","params":[],"id":1}' -H "Content-Type: application/json" http://localhost:8545
```

### `miner_setGasLimit`

::: tip
**Private**: 需要身份验证。
:::

设置矿工在采矿时将瞄准的气体限制。注意：在激活 [EIP-1559](https://eips.ethereum.org/EIPS/eip-1559) 的网络上，这应该设置为您想要的气体目标的两倍（即平均使用的有效气体每块）是。

::: warning
不受支持。此端点始终返回 `false`
:::

#### 参数

- Hex gas limit

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"miner_start","params":["0x10000"],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":false}
```

#### 参数

- Hex Number of threads

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"miner_start","params":["0x1"],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":false}
```

### `miner_setEtherbase`

::: tip
**Private**: 需要身份验证。
:::

设置 etherbase。它会更改将存放验证者奖励的钱包。

#### 参数

- 账户地址

```json
// Request
curl -X POST --data '{"jsonrpc":"2.0","method":"miner_setEtherbase","params":["0x3b7252d007059ffc82d16d022da3cbf9992d2f70"],"id":1}' -H "Content-Type: application/json" http://localhost:8545

// Result
{"jsonrpc":"2.0","id":1,"result":true}
```

## TxPool Methods
>>>>>>> d79323fb0f7dd4dd7b9baf4a27f51c8e5ac75472
