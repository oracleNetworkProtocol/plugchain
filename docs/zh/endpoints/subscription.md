---
order: 6
---

# 通过 Websocket 订阅事件

Tendermint 发出不同的事件，您可以通过以下方式订阅
[Websocket](https://en.wikipedia.org/wiki/WebSocket)。这很有用
用于第三方应用程序（用于分析）或用于检查状态。

[params-活动列表](https://godoc.org/github.com/tendermint/tendermint/types#pkg-constants)

要从 CLI 通过 websocket 连接到节点，您可以使用诸如
[wscat](https://github.com/websockets/wscat) 并运行：

```sh
wscat ws://127.0.0.1:26657/websocket
```

您可以通过调用 `subscribe` RPC 来订阅上述任何事件
方法通过 Websocket 以及有效的查询。

```json
{
    "jsonrpc": "2.0",
    "method": "subscribe",
    "id": 0,
    "params": {
        "query": "tm.event='NewBlock'"
    }
}
```

## 验证器集更新

当验证器集更改时，将发布 ValidatorSetUpdates 事件。这
事件携带公钥/权力对的列表。名单是一样的
Tendermint 从 ABCI 应用程序接收（参见 [EndBlock部分](https://github.com/tendermint/spec/blob/master/spec/abci/abci.md#endblock)在
ABCI 规范）。

回复：

```json
{
    "jsonrpc": "2.0",
    "id": 0,
    "result": {
        "query": "tm.event='ValidatorSetUpdates'",
        "data": {
            "type": "tendermint/event/ValidatorSetUpdates",
            "value": {
              "validator_updates": [
                {
                  "address": "09EAD022FD25DE3A02E64B0FE9610B1417183EE4",
                  "pub_key": {
                    "type": "tendermint/PubKeyEd25519",
                    "value": "ww0z4WaZ0Xg+YI10w43wTWbBmM3dpVza4mmSQYsd0ck="
                  },
                  "voting_power": "10",
                  "proposer_priority": "0"
                }
              ]
            }
        }
    }
}
```