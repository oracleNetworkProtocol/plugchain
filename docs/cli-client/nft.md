# NFT

`NFT` provides the ability to digitize assets. Through this module, each off-chain asset will be modeled as a unique on-chain asset.

## Available commands

| Name                                     | Description |
| ---------------------------------------- | ----------- |
| [issue-denom](#issue-denom)              | Issue assets |
| [issue-nft](#issue-nft)                  | Issue additional assets |
| [edit-nft](#edit-nft)                    | Edit Assets |
| [transfer-denom](#transfer-denom)        | Transfer assets |
| [transfer-nft](#transfer-nft)            | Transfer specified nft assets |
| [burn-nft](#burn-nft)                    | Destroy Assets |
| [supply](#supply)                        | Query supply |
| [owner](#owner)                          | Query by owner |
| [collection](#collection)                | Query collection |
| [denom](#denom)                          | Query denom |
| [denoms](#denoms)                        | Query denoms |
| [nft](#nft)                              | Query specified nft |

:::tip
The above commands can use the `-h` parameter to view the description of the function and the meaning of the parameters
:::

:::warning
The above `issue-denom`,`issue-nft`,`edit-nft`,`transfer-denom`,`transfer-nft`,`burn-nft` commands all require parameters `--from`,`--chain -id`,`--fees`
:::

**Logo:**

| Name, shorthand | Optional value | Description | Required |
| ----------- | ---- | ------------------ | ---- |
| --from | | The address used to sign the transaction sent by the execution function | true |
| --chain-id | plugchain | ID of the chain | true |
| --fees | 200plug | Pay transaction fees | true |
| --home | ~/.plugchain | The directory where the chain data is located | Optional |


## issue-denom

Issue assets

```bash
plugchaind tx nft issue-denom [denom-id] [denom-name] [denom-symbol] [mint-restricted] [edit-restricted] [schema-content or path to schema.json] [flags]
```

## issue-nft

Additional issuance (creation) of assets

```bash
plugchaind tx nft issue-nft [denom-id] [nft-id] [nft-name] [nft-url] [nft-data] [nft-recipient] [flags]
```


## edit-nft

Edit assets

```bash
plugchaind tx nft edit-nft [denom-id] [nft-id] [flags]
```

**Logo:**

| Name, shorthand | Default | Description | Required |
| ---------- | ---- | ------------------ | ---- |
| --nft-name | | The name of the nft asset | |
| --nft-url | | Where nft assets are located | |
| --nft-data | |nft asset metadata | |

## transfer-denom

Transfer assets

```bash
plugchaind tx nft transfer-denom <denom-id> <recipient-address> [flags]
```

## transfer-nft

Transfer of designated nft assets

```bash
plugchaind tx nft transfer-nft <denom-id> <nft-id> <recipient-address> [flags]
```

## burn-nft

Destroy assets

```bash
plugchaind tx nft burn-nft [denom-id] [nft-id] [flags]
```


## supply

Query the total amount of assets based on Denom ID; accept the optional --owner parameter.


```bash
plugchaind q nft supply [denom-id] [flags]
plugchaind q nft supply [denom-id] --owner=<myAddress> [flags]
```

## owner

Query all assets owned by an account; you can specify the Denom ID parameter.


```bash
plugchaind query nft owner [address] [denom-id] [flags]
```

## collection

Query all assets based on Denom ID.


```bash
plugchaind q nft collection [denom-id] [flags]
```

## denom

Query asset category information based on Denom ID.


```bash
plugchaind q nft denom [denom-id] [flags]
```

## denoms

Query all asset class information that has been issued.

```bash
plugchaind q nft denoms [flags]
```

## nft

Query specific assets based on Denom ID and ID.

```bash
plugchaind q nft nft [denom-id] [nft-id] [flags]
```