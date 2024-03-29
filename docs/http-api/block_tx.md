---
order: 11
---

# transaction in the block
 
- Get transaction data in the block

### Request method
`GET`

### URL
`/server/v1/tx/block`

### Request example

```
/server/v1/tx/block?height=8236948
```


#### Request parameters
- `height int ` : Block height

### Return value
- `code int64 `  : 0 or field does not exist is success, others are failures
- `message string` : Response information
- `txs array`: transaction data
  - `body object`
    - `message array`
  - `auth_info object`
    - `signer_infos array`: Signer public key information
    - `fee array` : Service charge array
  - `signatures array` :Signature data 
- `tx_responses object`: Transaction response data Main information
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
  "txs": [
    {
      "body": {
        "messages": [
          {
            "@type": "/cosmos.bank.v1beta1.MsgSend",
            "from_address": "gx1la5af7u4yq6mfmfeap0w9x0a57vgsqhvqm9ywm",
            "to_address": "gx17tdqp6neqjqw985e5wuez84j85d6phrz4d8xag",
            "amount": [
              {
                "denom": "uplugcn",
                "amount": "20000000000"
              }
            ]
          }
        ],
        "memo": "pledge-1676444023863",
        "timeout_height": "0",
        "extension_options": [],
        "non_critical_extension_options": []
      },
      "auth_info": {
        "signer_infos": [
          {
            "public_key": {
              "@type": "/cosmos.crypto.secp256k1.PubKey",
              "key": "AwoZJ8fSF7g1UWjv+gNHjpWSkl5r3XJ7ErMwsXguERkl"
            },
            "mode_info": {
              "single": {
                "mode": "SIGN_MODE_DIRECT"
              }
            },
            "sequence": "65"
          }
        ],
        "fee": {
          "amount": [
            {
              "denom": "uplugcn",
              "amount": "2000"
            }
          ],
          "gas_limit": "200000",
          "payer": "",
          "granter": ""
        }
      },
      "signatures": [
        "JI58ex4mxXm1Qjy/yfSTIXvHU5sv9MSH8Z4lmlCKGlt6q8u1wu9dvIqh04+OcTClCrCBvkiZPvI0mjdUE69xdQ=="
      ]
    }
  ],
  "tx_responses": [
    {
      "height": "8236948",
      "txhash": "3136EE671FF4B5F580661EB252C1D1CC8C5B0514E12D07EE8E1891E0ED8A6EDC",
      "codespace": "",
      "code": 0,
      "data": "0A1E0A1C2F636F736D6F732E62616E6B2E763162657461312E4D736753656E64",
      "raw_log": "[{\"events\":[{\"type\":\"coin_received\",\"attributes\":[{\"key\":\"receiver\",\"value\":\"gx17tdqp6neqjqw985e5wuez84j85d6phrz4d8xag\"},{\"key\":\"amount\",\"value\":\"20000000000uplugcn\"}]},{\"type\":\"coin_spent\",\"attributes\":[{\"key\":\"spender\",\"value\":\"gx1la5af7u4yq6mfmfeap0w9x0a57vgsqhvqm9ywm\"},{\"key\":\"amount\",\"value\":\"20000000000uplugcn\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/cosmos.bank.v1beta1.MsgSend\"},{\"key\":\"sender\",\"value\":\"gx1la5af7u4yq6mfmfeap0w9x0a57vgsqhvqm9ywm\"},{\"key\":\"module\",\"value\":\"bank\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"gx17tdqp6neqjqw985e5wuez84j85d6phrz4d8xag\"},{\"key\":\"sender\",\"value\":\"gx1la5af7u4yq6mfmfeap0w9x0a57vgsqhvqm9ywm\"},{\"key\":\"amount\",\"value\":\"20000000000uplugcn\"}]}]}]",
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
                  "value": "gx17tdqp6neqjqw985e5wuez84j85d6phrz4d8xag"
                },
                {
                  "key": "amount",
                  "value": "20000000000uplugcn"
                }
              ]
            },
            {
              "type": "coin_spent",
              "attributes": [
                {
                  "key": "spender",
                  "value": "gx1la5af7u4yq6mfmfeap0w9x0a57vgsqhvqm9ywm"
                },
                {
                  "key": "amount",
                  "value": "20000000000uplugcn"
                }
              ]
            },
            {
              "type": "message",
              "attributes": [
                {
                  "key": "action",
                  "value": "/cosmos.bank.v1beta1.MsgSend"
                },
                {
                  "key": "sender",
                  "value": "gx1la5af7u4yq6mfmfeap0w9x0a57vgsqhvqm9ywm"
                },
                {
                  "key": "module",
                  "value": "bank"
                }
              ]
            },
            {
              "type": "transfer",
              "attributes": [
                {
                  "key": "recipient",
                  "value": "gx17tdqp6neqjqw985e5wuez84j85d6phrz4d8xag"
                },
                {
                  "key": "sender",
                  "value": "gx1la5af7u4yq6mfmfeap0w9x0a57vgsqhvqm9ywm"
                },
                {
                  "key": "amount",
                  "value": "20000000000uplugcn"
                }
              ]
            }
          ]
        }
      ],
      "info": "",
      "gas_wanted": "200000",
      "gas_used": "91073",
      "tx": {
        "@type": "/cosmos.tx.v1beta1.Tx",
        "body": {
          "messages": [
            {
              "@type": "/cosmos.bank.v1beta1.MsgSend",
              "from_address": "gx1la5af7u4yq6mfmfeap0w9x0a57vgsqhvqm9ywm",
              "to_address": "gx17tdqp6neqjqw985e5wuez84j85d6phrz4d8xag",
              "amount": [
                {
                  "denom": "uplugcn",
                  "amount": "20000000000"
                }
              ]
            }
          ],
          "memo": "pledge-1676444023863",
          "timeout_height": "0",
          "extension_options": [],
          "non_critical_extension_options": []
        },
        "auth_info": {
          "signer_infos": [
            {
              "public_key": {
                "@type": "/cosmos.crypto.secp256k1.PubKey",
                "key": "AwoZJ8fSF7g1UWjv+gNHjpWSkl5r3XJ7ErMwsXguERkl"
              },
              "mode_info": {
                "single": {
                  "mode": "SIGN_MODE_DIRECT"
                }
              },
              "sequence": "65"
            }
          ],
          "fee": {
            "amount": [
              {
                "denom": "uplugcn",
                "amount": "2000"
              }
            ],
            "gas_limit": "200000",
            "payer": "",
            "granter": ""
          }
        },
        "signatures": [
          "JI58ex4mxXm1Qjy/yfSTIXvHU5sv9MSH8Z4lmlCKGlt6q8u1wu9dvIqh04+OcTClCrCBvkiZPvI0mjdUE69xdQ=="
        ]
      },
      "timestamp": "2023-02-15T06:53:46Z",
      "events": [
        {
          "type": "coin_spent",
          "attributes": [
            {
              "key": "c3BlbmRlcg==",
              "value": "Z3gxbGE1YWY3dTR5cTZtZm1mZWFwMHc5eDBhNTd2Z3NxaHZxbTl5d20=",
              "index": true
            },
            {
              "key": "YW1vdW50",
              "value": "MjAwMHVwbHVnY24=",
              "index": true
            }
          ]
        },
        {
          "type": "coin_received",
          "attributes": [
            {
              "key": "cmVjZWl2ZXI=",
              "value": "Z3gxN3hwZnZha20yYW1nOTYyeWxzNmY4NHoza2VsbDhjNWx5ZWd6YXc=",
              "index": true
            },
            {
              "key": "YW1vdW50",
              "value": "MjAwMHVwbHVnY24=",
              "index": true
            }
          ]
        },
        {
          "type": "transfer",
          "attributes": [
            {
              "key": "cmVjaXBpZW50",
              "value": "Z3gxN3hwZnZha20yYW1nOTYyeWxzNmY4NHoza2VsbDhjNWx5ZWd6YXc=",
              "index": true
            },
            {
              "key": "c2VuZGVy",
              "value": "Z3gxbGE1YWY3dTR5cTZtZm1mZWFwMHc5eDBhNTd2Z3NxaHZxbTl5d20=",
              "index": true
            },
            {
              "key": "YW1vdW50",
              "value": "MjAwMHVwbHVnY24=",
              "index": true
            }
          ]
        },
        {
          "type": "message",
          "attributes": [
            {
              "key": "c2VuZGVy",
              "value": "Z3gxbGE1YWY3dTR5cTZtZm1mZWFwMHc5eDBhNTd2Z3NxaHZxbTl5d20=",
              "index": true
            }
          ]
        },
        {
          "type": "tx",
          "attributes": [
            {
              "key": "ZmVl",
              "value": "MjAwMHVwbHVnY24=",
              "index": true
            }
          ]
        },
        {
          "type": "tx",
          "attributes": [
            {
              "key": "YWNjX3NlcQ==",
              "value": "Z3gxbGE1YWY3dTR5cTZtZm1mZWFwMHc5eDBhNTd2Z3NxaHZxbTl5d20vNjU=",
              "index": true
            }
          ]
        },
        {
          "type": "tx",
          "attributes": [
            {
              "key": "c2lnbmF0dXJl",
              "value": "Skk1OGV4NG14WG0xUWp5L3lmU1RJWHZIVTVzdjlNU0g4WjRsbWxDS0dsdDZxOHUxd3U5ZHZJcWgwNCtPY1RDbENyQ0J2a2laUHZJMG1qZFVFNjl4ZFE9PQ==",
              "index": true
            }
          ]
        },
        {
          "type": "message",
          "attributes": [
            {
              "key": "YWN0aW9u",
              "value": "L2Nvc21vcy5iYW5rLnYxYmV0YTEuTXNnU2VuZA==",
              "index": true
            }
          ]
        },
        {
          "type": "coin_spent",
          "attributes": [
            {
              "key": "c3BlbmRlcg==",
              "value": "Z3gxbGE1YWY3dTR5cTZtZm1mZWFwMHc5eDBhNTd2Z3NxaHZxbTl5d20=",
              "index": true
            },
            {
              "key": "YW1vdW50",
              "value": "MjAwMDAwMDAwMDB1cGx1Z2Nu",
              "index": true
            }
          ]
        },
        {
          "type": "coin_received",
          "attributes": [
            {
              "key": "cmVjZWl2ZXI=",
              "value": "Z3gxN3RkcXA2bmVxanF3OTg1ZTV3dWV6ODRqODVkNnBocno0ZDh4YWc=",
              "index": true
            },
            {
              "key": "YW1vdW50",
              "value": "MjAwMDAwMDAwMDB1cGx1Z2Nu",
              "index": true
            }
          ]
        },
        {
          "type": "transfer",
          "attributes": [
            {
              "key": "cmVjaXBpZW50",
              "value": "Z3gxN3RkcXA2bmVxanF3OTg1ZTV3dWV6ODRqODVkNnBocno0ZDh4YWc=",
              "index": true
            },
            {
              "key": "c2VuZGVy",
              "value": "Z3gxbGE1YWY3dTR5cTZtZm1mZWFwMHc5eDBhNTd2Z3NxaHZxbTl5d20=",
              "index": true
            },
            {
              "key": "YW1vdW50",
              "value": "MjAwMDAwMDAwMDB1cGx1Z2Nu",
              "index": true
            }
          ]
        },
        {
          "type": "message",
          "attributes": [
            {
              "key": "c2VuZGVy",
              "value": "Z3gxbGE1YWY3dTR5cTZtZm1mZWFwMHc5eDBhNTd2Z3NxaHZxbTl5d20=",
              "index": true
            }
          ]
        },
        {
          "type": "message",
          "attributes": [
            {
              "key": "bW9kdWxl",
              "value": "YmFuaw==",
              "index": true
            }
          ]
        }
      ]
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "1"
  }
}
```