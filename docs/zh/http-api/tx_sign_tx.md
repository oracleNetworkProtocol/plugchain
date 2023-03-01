---
order: 28
---

# 签名交易
 
- 通过用户给予的交易数据,签名广播交易

### 请求方式
`POST`

### URL
`/server/v1/tx/broadcast`

### 请求示例

```
/server/v1/tx/broadcast
```


#### 请求参数
- body 
    ```
    {"tx_bytes": "string","mode": "BROADCAST_MODE_UNSPECIFIED"}
    ```
    - tx_bytes : 交易原始数据
    - mode : 广播模式
      - BROADCAST_MODE_UNSPECIFIED
      - BROADCAST_MODE_BLOCK
      - BROADCAST_MODE_SYNC
      - BROADCAST_MODE_ASYNC
### 返回值
- `code int64 `  : 0或字段不存在为成功,其他均为失败
- `message string` : 响应信息
- `tx_responses object` : 成功返回数据
    - `height string` : 块高
    - `txhash string`: 交易hash
    - `code string` : 错误代码,0成功,其他均为失败,
    - `data string` : data,
    - `raw_log string` : 原始日志,
    - `logs string` : 格式化后的日志,
    - `info string` : "",
    - `gas_wanted string` : 手续费需要,
    - `gas_used string` : 使用手续费,
    - `tx string` : 交易信息,
    - `timestamp string` : 时间戳,
    - `events string` : 交易事件,

#### 返回示例
```json
{
 "tx_responses" : 
   {
      "height": "3864696",
      "txhash": "6A9956E8F4B7A0485930AB7301849C66FFDF4CB6052D95EF1026F19AD8763206",
      "codespace": "",
      "code": 0,
      "data": "0A390A372F636F736D6F732E646973747269627574696F6E2E763162657461312E4D7367576974686472617744656C656761746F72526577617264",
      "raw_log": "[{\"events\":[{\"type\":\"coin_received\",\"attributes\":[{\"key\":\"receiver\",\"value\":\"gx12hm4uyz9tv846ehg8zk8a7499hy84qrx4y0dnl\"},{\"key\":\"amount\",\"value\":\"7149580uplugcn\"}]},{\"type\":\"coin_spent\",\"attributes\":[{\"key\":\"spender\",\"value\":\"gx1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8nl07lv\"},{\"key\":\"amount\",\"value\":\"7149580uplugcn\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward\"},{\"key\":\"sender\",\"value\":\"gx1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8nl07lv\"},{\"key\":\"module\",\"value\":\"distribution\"},{\"key\":\"sender\",\"value\":\"gx12hm4uyz9tv846ehg8zk8a7499hy84qrx4y0dnl\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"gx12hm4uyz9tv846ehg8zk8a7499hy84qrx4y0dnl\"},{\"key\":\"sender\",\"value\":\"gx1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8nl07lv\"},{\"key\":\"amount\",\"value\":\"7149580uplugcn\"}]},{\"type\":\"withdraw_rewards\",\"attributes\":[{\"key\":\"amount\",\"value\":\"7149580uplugcn\"},{\"key\":\"validator\",\"value\":\"gxvaloper1vlpgtspv5f0yacrm2dpn5yzafz633du2c3cr8h\"}]}]}]",
      "logs": [
        {
          "msg_index": 0,
          "log": "",
          "events": [
            {
              "type": "coin_received",
              "attributes": [
                {
                  "key": "receiver",
                  "value": "gx12hm4uyz9tv846ehg8zk8a7499hy84qrx4y0dnl"
                },
                {
                  "key": "amount",
                  "value": "7149580uplugcn"
                }
              ]
            },
            {
              "type": "coin_spent",
              "attributes": [
                {
                  "key": "spender",
                  "value": "gx1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8nl07lv"
                },
                {
                  "key": "amount",
                  "value": "7149580uplugcn"
                }
              ]
            },
            {
              "type": "message",
              "attributes": [
                {
                  "key": "action",
                  "value": "/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward"
                },
                {
                  "key": "sender",
                  "value": "gx1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8nl07lv"
                },
                {
                  "key": "module",
                  "value": "distribution"
                },
                {
                  "key": "sender",
                  "value": "gx12hm4uyz9tv846ehg8zk8a7499hy84qrx4y0dnl"
                }
              ]
            },
            {
              "type": "transfer",
              "attributes": [
                {
                  "key": "recipient",
                  "value": "gx12hm4uyz9tv846ehg8zk8a7499hy84qrx4y0dnl"
                },
                {
                  "key": "sender",
                  "value": "gx1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8nl07lv"
                },
                {
                  "key": "amount",
                  "value": "7149580uplugcn"
                }
              ]
            },
            {
              "type": "withdraw_rewards",
              "attributes": [
                {
                  "key": "amount",
                  "value": "7149580uplugcn"
                },
                {
                  "key": "validator",
                  "value": "gxvaloper1vlpgtspv5f0yacrm2dpn5yzafz633du2c3cr8h"
                }
              ]
            }
          ]
        }
      ],
      "info": "",
      "gas_wanted": "200000",
      "gas_used": "106199",
      "tx": {
        "@type": "/cosmos.tx.v1beta1.Tx",
        "body": {
          "messages": [
            {
              "@type": "/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward",
              "delegator_address": "gx12hm4uyz9tv846ehg8zk8a7499hy84qrx4y0dnl",
              "validator_address": "gxvaloper1vlpgtspv5f0yacrm2dpn5yzafz633du2c3cr8h"
            }
          ],
          "memo": "",
          "timeout_height": "0",
          "extension_options": [],
          "non_critical_extension_options": []
        },
        "auth_info": {
          "signer_infos": [
            {
              "public_key": {
                "@type": "/cosmos.crypto.secp256k1.PubKey",
                "key": "AzwK895WIBnk+3ZP0yJtOcZGPoTz6zDBGkL4Om7JSDRZ"
              },
              "mode_info": {
                "single": {
                  "mode": "SIGN_MODE_DIRECT"
                }
              },
              "sequence": "11362"
            }
          ],
          "fee": {
            "amount": [
              {
                "denom": "uplugcn",
                "amount": "200"
              }
            ],
            "gas_limit": "200000",
            "payer": "",
            "granter": ""
          }
        },
        "signatures": [
          "u/CVR0Dg76Y7esmRBdXuIEUQmwObV33Qkui6Ixgu8/4I0fSw6XZHYhtp7Nb/fZzxpH0XQuqbqXaemDCqhcr4nQ=="
        ]
      },
      "timestamp": "2022-05-06T11:30:02Z",
      "events": [
      ]
    }
}
```