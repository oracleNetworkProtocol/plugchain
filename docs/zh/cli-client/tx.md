# Tx

签名、广播、查询交易

## 可用命令

| 名称                            | 描述                     |
| ------------------------------- | ------------------------ |
| [sign](#plugchaind-tx-sign)           | 签名生成的离线交易文件   |
| [broadcast](#plugchaind-tx-broadcast) | 广播一个已签名交易到网络 |
| [multisig](#plugchaind-tx-multisign)  | 用多个账户为同一交易签名 |
| [tx](#plugchaind-query-tx)            | 使用交易hash查询交易     |
| [txs](#plugchaind-query-txs)          | 使用Tag查询交易          |

## plugchaind tx sign

签名生成的离线交易文件。该文件由`generate-only`标志生成。

```bash
plugchaind tx sign <file> [flags]
```

### 标志

| 命令，速记       | 类型   | 必须 | 默认  | 描述                             |
| ---------------- | ------ | ---- | ----- | -------------------------------- |
| -h，--help       |        |      |       | 打印帮助                         |
| --append         | bool   |      | true  | 将签名附加到现有签名             |
| --multisig       | string |      | true  | 代表交易签名的multisig帐户的地址 |
| --from           | string | 是   |       | 用于签名的私钥名称               |
| --offline        | bool   |      | false | 离线模式                         |
| --signature-only | bool   |      | false | 仅打印生成的签名，然后退出       |

### 生成离线交易

:::tip
任意类型的离线交易都可以通过追加`--generate-only`标志来生成
:::

下面示例中使用Transfer交易：

```bash
plugchaind tx bank send gx1w9lvhwlvkwqvg08q84n2k4nn896u9pqx93velx gx15uys54epmd2xzhcn32szps56wvev40tt908h62 10plugchaind --chain-id=plugchain --generate-only
```

`unsigned.json` 看起来是这样的：

```json
{
    "type": "cosmos-sdk/StdTx",
    "value": {
        "msg": [
            {
                "type": "cosmos-sdk/MsgSend",
                "value": {
                    "from_address": "gx1w9lvhwlvkwqvg08q84n2k4nn896u9pqx93velx",
                    "to_address": "gx15uys54epmd2xzhcn32szps56wvev40tt908h62",
                    "amount": [
                        {
                            "denom": "plug",
                            "amount": "10"
                        }
                    ]
                }
            }
        ],
        "fee": {
            "amount": [],
            "gas": "200000"
        },
        "signatures": null,
        "memo": ""
    }
}
```

### 签名离线交易

```bash
plugchaind tx sign unsigned.json --name=<key-name> > signed.tx
```

`signed.json` 看起来是这样的：

```json
{
    "type": "auth/StdTx",
    "value": {
        "msg": [
            {
                "type": "cosmos-sdk/Send",
                "value": {
                    "inputs": [
                        {
                            "address": "gx106nhdckyf996q69v3qdxwe6y7408pvyvyxzhxh",
                            "coins": [
                                {
                                    "denom": "plug",
                                    "amount": "10000000"
                                }
                            ]
                        }
                    ],
                    "outputs": [
                        {
                            "address": "gx1893x4l2rdshytfzvfpduecpswz7qtpstevr742",
                            "coins": [
                                {
                                    "denom": "plug",
                                    "amount": "10000000"
                                }
                            ]
                        }
                    ]
                }
            }
        ],
        "fee": {
            "amount": [
                {
                    "denom": "plug",
                    "amount": "40000000"
                }
            ],
            "gas": "200000"
        },
        "signatures": [
            {
                "pub_key": {
                    "type": "tendermint/PubKeySecp256k1",
                    "value": "Auouudrg0P86v2kq2lykdr97AJYGHyD6BJXAQtjR1gzd"
                },
                "signature": "sJewd6lKjma49rAiGVfdT+V0YYerKNx6ZksdumVCvuItqGm24bEN9msh7IJ12Sil1lYjqQjdAcjVCX/77FKlIQ==",
                "account_number": "0",
                "sequence": "3"
            }
        ],
        "memo": "test"
    }
}
```

签名之后，`signed.json`中的`signature`字段将不再为空。

现在准备[广播这个已签名交易](#plugchaind-tx-broadcast)到PLUGChain Hub。

## plugchaind tx broadcast

这个命令用于广播已离线签名的交易到网络。

### 广播离线签名的交易

```bash
plugchaind tx broadcast signed.json --chain-id=plugchain
```

## plugchaind tx multisign

用多个账户为一个交易签名。这个交易只有在签名数满足multisig-threshold时才可以广播。

```bash
plugchaind tx multisign <file> <key-name> <[signature]...> [flags]
```

### 用多签密钥创建离线交易

:::tip
没有多签密钥？[创建一个](keys.md#创建多签密钥)
:::

```bash
plugchaind tx bank send <from> <to> 10plugchaind --fees=20plug --chain-id=plugchain --from=<multisig-keyname> --generate-only > unsigned.json
```

### 签名多签交易

#### 查询多签地址

```bash
plugchaind keys show <multisig-keyname>
```

#### 签名`unsigned.json`

假定multisig-threshold是2，我们使用2个签名者签名`unsigned.json`

用signer-1签名：

```bash
plugchaind tx sign unsigned.json --from=<signer-keyname-1> --chain-id=plugchain --multisig=<multisig-address> --signature-only > signed-1.json
```

用signer-2签名：

```bash
plugchaind tx sign unsigned.json --from=<signer-keyname-2> --chain-id=plugchain --multisig=<multisig-address> --signature-only > signed-2.json
```

#### 合并签名

合并所有签名到 `signed.json`

```bash
plugchaind tx multisign --chain-id=plugchain unsigned.json <multisig-keyname> signed-1.json signed-2.json > signed.json
```

现在可以[广播](#plugchaind-tx-broadcast)这个已签名交易了。

## plugchaind query tx

```bash
plugchaind query tx [hash] [flags]
```

## plugchaind query txs

```bash
plugchaind query txs --events 'message.sender=<gx...>&message.action=xxxx' --page 1 --limit 30
```

其中`message.action`可取值：

| module       | Msg                                       | action               |
| ------------ | ----------------------------------------- | -------------------- |
| bank         | cosmos-sdk/MsgSend                        | transfer             |
|              | cosmos-sdk/MsgMultiSend                   | transfer             |
| distribution | cosmos-sdk/MsgModifyWithdrawAddress       | set_withdraw_address |
|              | cosmos-sdk/MsgWithdrawValidatorCommission | withdraw_commission  |
|              | cosmos-sdk/MsgWithdrawDelegatorReward     | withdraw_rewards     |
| gov          | cosmos-sdk/MsgSubmitProposal              | submit_proposal      |
|              | cosmos-sdk/MsgDeposit                     | proposal_deposit     |
|              | cosmos-sdk/MsgVote                        | proposal_vote        |
| stake        | cosmos-sdk/MsgCreateValidator             | create_validator     |
|              | cosmos-sdk/MsgEditValidator               | edit_validator       |
|              | cosmos-sdk/MsgDelegate                    | delegate             |
|              | cosmos-sdk/MsgBeginRedelegate             | redelegate           |
|              | cosmos-sdk/MsgUndelegate                  | unbond               |
| slashing     | cosmos-sdk/MsgUnjail                      | unjail               |