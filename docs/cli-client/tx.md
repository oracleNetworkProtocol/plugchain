# Tx

Tx module allows you to sign or broadcast transactions

## Available Commands

| Name                            | Description                                                                           |
| ------------------------------- | ------------------------------------------------------------------------------------- |
| [sign](#plugchaind-tx-sign)           | Sign transactions generated offline                                                   |
| [broadcast](#plugchaind-tx-broadcast) | Broadcast a signed transaction to the network                                         |
| [multisign](#plugchaind-tx-multisign) | Sign the same transaction by multiple accounts                                        |
| [tx](#plugchaind-query-tx)            | Query for a transaction by hash in a committed block                                  |
| [txs](#plugchaind-query-txs)          | Search for transactions that match the exact given events where results are paginated |

## plugchaind tx sign

Sign transactions in generated offline file. The file created with the --generate-only flag.

```bash
plugchaind tx sign <file> [flags]
```

### Flags

| Name, shorthand  | Type   | Required | Default | Description                                                                        |
| ---------------- | ------ | -------- | ------- | ---------------------------------------------------------------------------------- |
| --append         | bool   | true     | true    | Attach a signature to an existing signature.                                       |
| --from           | string | true     |         | Key name for signature                                                             |
| --offline        | bool   | true     |         | Offline mode.                                                                      |
| --signature-only | bool   | true     |         | Print only the generated signature, then exit                                      |
| --multisig       | string |          | true    | Address of the multisig account on behalf of which the transaction shall be signed |

### Generate an offline tx

:::tip
You can generate any type of txs offline by appending the flag `--generate-only`
:::

We use a transfer tx in the following examples:

```bash
plugchaind tx bank send gx1l7422f2ljgmrnzwxgd9rxr8cw3yz63wx89tnqs gx1rkecwy9pjsfd058w0pwa2perquc8xe639n92ms 900000uplugcn --fees 20uplugcn --generate-only > unsigned.json
```

The `unsigned.json` should look like:

```json
{
    "body": {
        "messages": [
            {
                "@type": "/cosmos.bank.v1beta1.MsgSend",
                "from_address": "gx1l7422f2ljgmrnzwxgd9rxr8cw3yz63wx89tnqs",
                "to_address": "gx1rkecwy9pjsfd058w0pwa2perquc8xe639n92ms",
                "amount": [
                    {
                        "denom": "uplugcn",
                        "amount": "900000"
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
                    "denom": "uplugcn",
                    "amount": "20"
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

### Sign tx offline

```bash
plugchaind tx sign unsigned.json  --from gx1l7422f2ljgmrnzwxgd9rxr8cw3yz63wx89tnqs --chain-id plugchain_520-1 > signed.json
```

The `signed.json` should look like:

```json
{
    "body": {
        "messages": [
            {
                "@type": "/cosmos.bank.v1beta1.MsgSend",
                "from_address": "gx1l7422f2ljgmrnzwxgd9rxr8cw3yz63wx89tnqs",
                "to_address": "gx1rkecwy9pjsfd058w0pwa2perquc8xe639n92ms",
                "amount": [
                    {
                        "denom": "uplugcn",
                        "amount": "900000"
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
                    "key": "Auouudrg0P86v2kq2lykdr97AJYGHyD6BJXAQtjR1gzd"
                },
                "mode_info": {
                    "single": {
                        "mode": "SIGN_MODE_DIRECT"
                    }
                },
                "sequence": "1"
            }
        ],
        "fee": {
            "amount": [
                {
                    "denom": "uplugcn",
                    "amount": "20"
                }
            ],
            "gas_limit": "200000",
            "payer": "",
            "granter": ""
        }
    },
    "signatures": [
        "sJewd6lKjma49rAiGVfdT+V0YYerKNx6ZksdumVCvuItqGm24bEN9msh7IJ12Sil1lYjqQjdAcjVCX/77FKlIQ=="
    ]
}
```

Note the `signature` in the `signed.json` should no longer be empty after signing.

Now it's ready to [broadcast the signed tx](#plugchaind-tx-broadcast) to the Plug Chain Hub.

## plugchaind tx broadcast

This command is used to broadcast an offline signed transaction to the network.

### Broadcast offline signed transaction

```bash
plugchaind tx broadcast signed.json --chain-id=plugchain_520-1
```

## plugchaind tx multisign

Sign a transaction by multiple accounts. The tx could be broadcasted only when the number of signatures meets the multisig-threshold.

```bash
plugchaind tx multisign <file> <key-name> <[signature]...> [flags]
```

### Generate an offline tx by multisig key

:::tip
No multisig key? [Create one](keys.md#create-a-multisig-key)
:::

```bash
plugchaind tx bank send <from> <to> 10uplugcn --fees=20uplugcn --chain-id=plugchain_520-1 --from=<multisig-keyname> --generate-only > unsigned.json
```

### Sign the multisig tx

#### Query the multisig address

```bash
plugchaind keys show <multisig-keyname>
```

#### Sign the `unsigned.json`

Assume the multisig-threshold is 2, here we sign the `unsigned.json` by 2 of the signers

Sign the tx by signer-1:

```bash
plugchaind tx sign unsigned.json --from=<signer-keyname-1> --chain-id=plugchain_520-1 --multisig=<multisig-address> --signature-only > signed-1.json
```

Sign the tx by signer-2:

```bash
plugchaind tx sign unsigned.json --from=<signer-keyname-2> --chain-id=plugchain_520-1 --multisig=<multisig-address> --signature-only > signed-2.json
```

#### Merge the signatures

Merge all the signatures into `signed.json`

```bash
plugchaind tx multisign --chain-id=plugchain_520-1 unsigned.json <multisig-keyname> signed-1.json signed-2.json > signed.json
```

Now you can [broadcast the signed tx](#plugchaind-tx-broadcast).

## plugchaind query tx

```bash
plugchaind query tx [hash] [flags]
```

## plugchaind query txs

```bash
plugchaind query txs --events 'message.sender=<gx...>&message.action=xxxx' --page 1 --limit 30
```

Among the possible values of `message.action`:

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
