---
order: 23
---

# Broadcast transaction
 
- Broadcast transaction data to the chain

### Request method
`POST`

### URL
`/server/v1/tx/broadcast`

### Request example

```
/server/v1/tx/broadcast
```


#### Request parameters
- body 
    ```
    {"tx_bytes": "string","mode": "BROADCAST_MODE_UNSPECIFIED"}
    ```
    - tx_bytes : Original transaction data
    - mode : Broadcast mode
      - BROADCAST_MODE_UNSPECIFIED
      - BROADCAST_MODE_BLOCK
      - BROADCAST_MODE_SYNC
      - BROADCAST_MODE_ASYNC
### Return value
- `code int64 `  : 0 or field does not exist is success, others are failures
- `message string` : Response information
- `tx_responses object` : Data returned successfully
    - `height string` : Block height
    - `txhash string`: Transaction hash
    - `code string` : Error code, 0 succeeded, others failed,
    - `data string` : data,
    - `raw_log string` : Raw log,
    - `logs string` : Formatted log,
    - `info string` : "",
    - `gas_wanted string` : Service charge required,
    - `gas_used string` : Service charge for use,
    - `tx string` : Transaction information,
    - `timestamp string` : Time stamp,
    - `events string` : Transaction event,

#### Return to example
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