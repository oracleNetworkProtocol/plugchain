---
order: 4
---

# Events

`Event`s 是包含有关应用程序执行信息的对象。他们是主要由区块浏览器和钱包等服务提供商用于跟踪各种执行情况消息和索引事务。

## 先决条件

- [Cosmos SDK Events](https://docs.cosmos.network/master/core/events.html) 
- [Ethereum's PubSub JSON-RPC API](https://geth.ethereum.org/docs/rpc/pubsub) 

## 订阅事件

### SDK 和 Tendermint 事件

可以通过 Tendermint 的 [Websocket](https://tendermint.com/docs/app-dev/subscribing-to-events-via-websocket.html#subscribing-to-events-via-websocket) 订阅`events` ）。
这是通过 Websocket 调用 `subscribe` RPC 方法来完成的：

```json
{
    "jsonrpc": "2.0",
    "method": "subscribe",
    "id": "0",
    "params": {
        "query": "tm.event='eventCategory' AND eventType.eventAttribute='attributeValue'"
    }
}
```

您可以订阅的主要 `eventCategory` 是：

- `NewBlock`：包含在`BeginBlock`和`EndBlock`期间触发的`events`。
- `Tx`：包含在`DeliverTx`（即事务处理）期间触发的`events`。
- `ValidatorSetUpdates`：包含块的验证器集更新。

这些事件是在提交块后从`state`包中触发的。你可以获得完整的`events`类别列表。[这里]（https://godoc.org/github.com/tendermint/tendermint/types#pkg-constants）。

`query` 的 `type` 和 `attribute` 值允许您过滤您所在的特定 `event`
寻找。例如，`MsgEthereumTx`交易会触发`ethermint`类型的`event`，并且
具有`sender`和`recipient`作为`attributes`。订阅此`event`将按如下方式完成：
```json
{
    "jsonrpc": "2.0",
    "method": "subscribe",
    "id": "0",
    "params": {
        "query": "tm.event='Tx' AND ethereum.recipient='hexAddress'"
    }
}
```

其中“hexAddress”是一个以太坊十六进制地址（例如：“0x1122334455667788990011223344556677889900”）。

### Ethereum JSON-RPC Events

Plug Chain 还支持 Ethereum [JSON-RPC](https://eth.wiki/json-rpc/API) 过滤器调用订阅 [状态日志](https://eth.wiki/json-rpc/API#eth_newfilter),[区块](https://eth.wiki/json-rpc/API#eth_newblockfilter) 或 [待定交易]（https://eth.wiki/json-rpc/API#eth_newpendingtransactionfilter）更改。

在底层，它使用 Tendermint RPC 客户端的事件系统来处理订阅
然后格式化为与以太坊兼容的事件。

```bash
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_newBlockFilter","params":[],"id":1}' -H "Content-Type: application/json" http://localhost:8545

{"jsonrpc":"2.0","id":1,"result":"0x3503de5f0c766c68f78a03a3b05036a5"}
```

然后您可以使用 [`eth_getFilterChanges`](https://eth.wiki/json-rpc/API#eth_getfilterchanges) 调用检查状态是否发生变化：

```bash
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_getFilterChanges","params":["0x3503de5f0c766c68f78a03a3b05036a5"],"id":1}' -H "Content-Type: application/json" http://localhost:8545

{"jsonrpc":"2.0","id":1,"result":["0x7d44dceff05d5963b5bc81df7e9f79b27e777b0a03a6feca09f3447b99c6fa71","0x3961e4050c27ce0145d375255b3cb829a5b4e795ac475c05a219b3733723d376","0xd7a497f95167d63e6feca70f344d9f6e843d097b62729b8f43bdcd5febf142ab","0x55d80a4ba6ef54f2a8c0b99589d017b810ed13a1fda6a111e1b87725bc8ceb0e","0x9e8b92c17280dd05f2562af6eea3285181c562ebf41fc758527d4c30364bcbc4","0x7353a4b9d6b35c9eafeccaf9722dd293c46ae2ffd4093b2367165c3620a0c7c9","0x026d91bda61c8789c59632c349b38fd7e7557e6b598b94879654a644cfa75f30","0x73e3245d4ddc3bba48fa67633f9993c6e11728a36401fa1206437f8be94ef1d3"]}
```

## Websocket Connection

### Tendermint Websocket

要开始与 Tendermint websocket 的连接，您需要使用 `--rpc.laddr` 定义地址
启动节点时的标志（默认`tcp://127.0.0.1:26657`）：

```bash
plugchaind start --rpc.laddr="tcp://127.0.0.1:26657"
```

然后，使用 [ws](https://github.com/hashrocket/ws) 开始一个 websocket 订阅

```bash
# connect to tendermint websocket at port 8080 as defined above
ws ws://localhost:8080/websocket

# subscribe to new Tendermint block headers
> { "jsonrpc": "2.0", "method": "subscribe", "params": ["tm.event='NewBlockHeader'"], "id": 1 }
```

### Ethereum Websocket

由于 Plug Chain 运行使用 Tendermint Core 作为它的共识引擎，并且它是使用 Cosmos 构建的
SDK 框架，它继承了它们的事件格式。但是为了支持原生的Web3
[Ethereum's websockets的兼容性
PubSubAPI](https://geth.ethereum.org/docs/rpc/pubsub)，Plug Chain 需要转换 Tendermint
检索到以太坊类型的响应。

您可以在启动时使用 `--json-rpc.ws-address` 标志启动与以太坊 websocket 的连接
节点（默认“0.0.0.0:8546”）：
```bash
plugchaind start  --json-rpc.address"0.0.0.0:8545" --json-rpc.ws-address="0.0.0.0:8546" --evm.rpc.api="eth,web3,net,txpool,debug" --json-rpc.enable
```

然后，使用 [`ws`](https://github.com/hashrocket/ws) 开始一个 websocket 订阅

```bash
# connect to tendermint websocet at port 8546 as defined above
ws ws://localhost:8546/

# subscribe to new Ethereum-formatted block Headers
> {"id": 1, "method": "eth_subscribe", "params": ["newHeads", {}]}
< {"jsonrpc":"2.0","result":"0x44e010cb2c3161e9c02207ff172166ef","id":1}
```
