# Bank模块

## 简介

该模块主要用于账户之间转账、查询账户余额，同时提供了通用的离线签名与交易广播方法。

## 使用场景

1. 账户查询

    查询指定账户的余额。

    ```bash
    plugchaind query bank balances [address] --denom=[denom]
    ```

2. 账户间转账

    该命令包括了交易“构造，签名，广播”的所有操作。 如从账户A转账10plug给账户B：

    ```bash
    plugchaind tx bank send [A] [B] [10plug] --fees=20uplugcn --chain-id=plugchain_520-1
    ```

    Plug Chain支持多种代币流通，将来Plug Chain可以在一个交易中包含多种代币交换。

3. 交易签名

    为了提高账户安全性，Plug Chain支持交易离线签名保护账户的私钥。在任意交易中，使用参数--generate-only可以构建一个未签名的交易。使用转账交易作为示例：

    ```bash
    plugchaind tx bank send [from_key_or_address] [to_address] [amount] --fees=20uplugcn --generate-only
    ```

    以上命令将构建一未签名交易：

    ```json
    {
        "type": "auth/StdTx",
        "value": {
            "msg": [ "txMsg" ],
            "fee": "fee",
            "signatures": null,
            "memo": ""
        }
    }
    ```

    将结果保存到文件`<file>`。

    对上述的离线交易进行签名：

    ```bash
    plugchaind tx sign <file> --chain-id=plugchain_520-1 --from=<key-name>
    ```

    将返回已签名的交易：

    ```json
    {
        "type": "auth/StdTx",
        "value": {
        "msg": [ "txMsg" ],
        "fee": "fee",
        "signatures": [
            {
            "pub_key": {
                "type": "tendermint/PubKeySecp256k1",
                "value": "A+qXW5isQDb7blT/KwEgQHepji8RfpzIstkHpKoZq0kr"
            },
            "signature": "5hxk/R81SWmKAGi4kTW2OapezQZpp6zEnaJbVcyDiWRfgBm4Uejq8+CDk6uzk0aFSgAZzz06E014UkgGpelU7w==",
            "account_number": "0",
            "sequence": "11"
            }
        ],
        "memo": ""
        }
    }
    ```

    将结果保存到文件`<file>`。

4. 广播交易


    ```bash
    plugchaind tx broadcast <file>
    ```

    该交易将在Plug Chain Hub中广播并执行。
