# Bank

## Summary

This module is mainly used to transfer coins between accounts, query account balances, and provide common offline transaction signing and broadcasting methods. In addition, the available units of tokens in the Plug Chain Hub system are defined using [coin-type](../concepts/coin-type.md).

## Usage Scenario

1. Query the balances of an account

    Query the total balance of an account or of a specific denomination.

    ```bash
    plugchaind query bank balances [address] --denom=[denom]
    ```

2. Transfer between accounts

    For example, transfer 10plugchaind from account A to account B:

    ```bash
    plugchaind tx bank send [A] [B] [10plug] --fees=20plug --chain-id=plugchain
    ```

    Plug Chain Hub supports multiple tokens in circulation, and in the future Plug Chain Hub will be able to include multiple tokens in one transaction.

3. Sign transactions generated offline

    To improve account security, Plug Chain Hub supports offline signing of transactions to protect the account's private key. In any transaction, you can build an unsigned transaction using the flag --generate-only. Take transfer transaction as an example:

    ```bash
    plugchaind tx bank send [from_key_or_address] [to_address] [amount]  --fees=20plug --generate-only
    ```

    Return the built transaction with empty signatures:

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

    Save the result to a file.

    Sign the transaction:

    ```bash
    plugchaind tx sign <file> --chain-id=plugchain --from=<key-name>
    ```

    Return signed transactions:

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

    Save the result to a file.

4. Broadcast transactions


    ```bash
    plugchaind tx broadcast <file>
    ```

    The transaction will be broadcast and executed in Plug Chain Hub.
