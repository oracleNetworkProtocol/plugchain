---
title: 流程性交易
---
### 本文交易步骤，以转账为例
#### 参数
- 两个地址
```golang 
  - name: wallet1
    type: local
    address: gx1cpq88us4kprcdd4fy4z5halwkr8h5eejrgacsq
    pubkey: gxpub1addwnpepqf0egzqdu83e4yvd3hsw4mck0d7jjkhq76d5etc2rymmkhxtlv4pxw7r8st
    mnemonic: ""
    threshold: 0
    pubkeys: []
  - name: wallet2
    type: local
    address: gx173axkxh93lgtj3x9ddjvhtyy32xrq6q5pk04e9
    pubkey: gxpub1addwnpepqf0du4l04hqquvz7q3ydq75uma6zu3q8ylk2qhydv5ka56edzu0v2y8gznl
    mnemonic: ""
    threshold: 0
    pubkeys: []
```
- 需求：wallet1 向 wallet2 转账一笔

1. 生成无符号tx
```bash
plugchaind tx bank send $(plugchaind keys show wallet1 -a --home node1) \
$(plugchaind keys show wallet2 -a --home node1) 1000000line \
--from wallet1 --fees 5000line --chain-id plugchain --generate-only > tx.json
```
tx.json信息如下:
```json
{
    "body": {
        "messages": [
            {
                "@type": "/cosmos.bank.v1beta1.MsgSend",
                "from_address": "gx1cpq88us4kprcdd4fy4z5halwkr8h5eejrgacsq",
                "to_address": "gx173axkxh93lgtj3x9ddjvhtyy32xrq6q5pk04e9",
                "amount": [
                    {
                        "denom": "line",
                        "amount": "1000000"
                    }
                ]
            }
        ],
        "memo": "",
        "timeout_height": "0",
        "extension_options": [],
        "non_critical_extension_options": []
    },
    "auth_info": {
        "signer_infos": [],
        "fee": {
            "amount": [
                {
                    "denom": "line",
                    "amount": "5000"
                }
            ],
            "gas_limit": "200000",
            "payer": "",
            "granter": ""
        }
    },
    "signatures": []
}
```

2. 对无符号tx签名
```bash
plugchaind tx sign tx.json --from wallet1 --chain-id plugchain -y > tx-sign.json
```
tx-sign.json信息：
```json
{
    "body": {
        "messages": [
            {
                "@type": "/cosmos.bank.v1beta1.MsgSend",
                "from_address": "gx1cpq88us4kprcdd4fy4z5halwkr8h5eejrgacsq",
                "to_address": "gx173axkxh93lgtj3x9ddjvhtyy32xrq6q5pk04e9",
                "amount": [
                    {
                        "denom": "line",
                        "amount": "1000000"
                    }
                ]
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
                    "key": "Al+UCA3h45qRjY3g6u8We30pWuD2m0yvChk3u1zL+yoT"
                },
                "mode_info": {
                    "single": {
                        "mode": "SIGN_MODE_DIRECT"
                    }
                },
                "sequence": "0"
            }
        ],
        "fee": {
            "amount": [
                {
                    "denom": "line",
                    "amount": "5000"
                }
            ],
            "gas_limit": "200000",
            "payer": "",
            "granter": ""
        }
    },
    "signatures": [
        "MCJY9Gxg9pMEjj4Dn+WvD5z7GgR93M8pWSsfgnrQVBRj1rSm5CNG8ZOwpICPfvPuBoyyFJzehDzdIQU4rd29vQ=="
    ]
}
```
3. 对签名交易信息编码
```bash
plugchaind tx encode tx-sign.json
```
输出如下：
```text
CooBCocBChwvY29zbW9zLmJhbmsudjFiZXRhMS5Nc2dTZW5kEmcKKWd4MWNwcTg4dXM0a3ByY2RkNGZ5NHo1aGFsd2tyOGg1ZWVqcmdhY3NxEilneDE3M2F4a3hoOTNsZ3RqM3g5ZGRqdmh0eXkzMnhycTZxNXBrMDRlORoPCgRsaW5lEgcxMDAwMDAwEmQKTgpGCh8vY29zbW9zLmNyeXB0by5zZWNwMjU2azEuUHViS2V5EiMKIQJflAgN4eOakY2N4OrvFnt9KVrg9ptMrwoZN7tcy/sqExIECgIIARISCgwKBGxpbmUSBDUwMDAQwJoMGkAwIlj0bGD2kwSOPgOf5a8PnPsaBH3czylZKx+CetBUFGPWtKbkI0bxk7CkgI9+8+4GjLIUnN6EPN0hBTit3b29
```
4. 广播交易
```bash
curl --header "Content-Type: application/json" --request POST --data '{"tx_bytes":"CooBCocBChwvY29zbW9zLmJhbmsudjFiZXRhMS5Nc2dTZW5kEmcKKWd4MWNwcTg4dXM0a3ByY2RkNGZ5NHo1aGFsd2tyOGg1ZWVqcmdhY3NxEilneDE3M2F4a3hoOTNsZ3RqM3g5ZGRqdmh0eXkzMnhycTZxNXBrMDRlORoPCgRsaW5lEgcxMDAwMDAwEmQKTgpGCh8vY29zbW9zLmNyeXB0by5zZWNwMjU2azEuUHViS2V5EiMKIQJflAgN4eOakY2N4OrvFnt9KVrg9ptMrwoZN7tcy/sqExIECgIIARISCgwKBGxpbmUSBDUwMDAQwJoMGkAwIlj0bGD2kwSOPgOf5a8PnPsaBH3czylZKx+CetBUFGPWtKbkI0bxk7CkgI9+8+4GjLIUnN6EPN0hBTit3b29","mode":1}' localhost:1317/cosmos/tx/v1beta1/txs
```
返回消息如下：
```json
{
  "tx_response": {
    "height": "292",
    "txhash": "D7CD55CDE644952F2DCAEB32DD6E865657210EBCBBAE2914B256292762702C69",
    "codespace": "",
    "code": 0,
    "data": "0A060A0473656E64",
    "raw_log": "[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"send\"},{\"key\":\"sender\",\"value\":\"gx1cpq88us4kprcdd4fy4z5halwkr8h5eejrgacsq\"},{\"key\":\"module\",\"value\":\"bank\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"gx173axkxh93lgtj3x9ddjvhtyy32xrq6q5pk04e9\"},{\"key\":\"sender\",\"value\":\"gx1cpq88us4kprcdd4fy4z5halwkr8h5eejrgacsq\"},{\"key\":\"amount\",\"value\":\"1000000line\"}]}]}]",
    "logs": [
      {
        "msg_index": 0,
        "log": "",
        "events": [
          {
            "type": "message",
            "attributes": [
              {
                "key": "action",
                "value": "send"
              },
              {
                "key": "sender",
                "value": "gx1cpq88us4kprcdd4fy4z5halwkr8h5eejrgacsq"
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
                "value": "gx173axkxh93lgtj3x9ddjvhtyy32xrq6q5pk04e9"
              },
              {
                "key": "sender",
                "value": "gx1cpq88us4kprcdd4fy4z5halwkr8h5eejrgacsq"
              },
              {
                "key": "amount",
                "value": "1000000line"
              }
            ]
          }
        ]
      }
    ],
    "info": "",
    "gas_wanted": "200000",
    "gas_used": "74685",
    "tx": null,
    "timestamp": ""
  }
}
```
5. 转账完成