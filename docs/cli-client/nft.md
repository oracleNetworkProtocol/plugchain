# NFT

`NFT` provides the ability to digitize assets. Through this module, each off-chain asset will be modeled as a unique on-chain asset.

## Available commands

| Name                                     | Description |
| ---------------------------------------- | ----------- |
| [issue-class](#issue-class)              | Create asset classification |
| [issue-nft](#issue-nft)                  | Issue additional assets |
| [edit-nft](#edit-nft)                    | Edit Assets |
| [transfer-class](#transfer-class)        | Transfer asset classification |
| [transfer-nft](#transfer-nft)            | Transfer specified nft assets |
| [burn-nft](#burn-nft)                    | Destroy Assets |
| [supply](#supply)                        | Query supply |
| [owner](#owner)                          | Query by owner |
| [collection](#collection)                | Query collection |
| [class](#class)                          | Query class |
| [classes](#classes)                       | Query classes |
| [nft](#nft)                              | Query specified nft |

:::tip
The above commands can use the `-h` parameter to view the description of the function and the meaning of the parameters
:::

:::warning
The above `issue-class`,`issue-nft`,`edit-nft`,`transfer-class`,`transfer-nft`,`burn-nft` commands all require parameters `--from`,`--chain -id`,`--fees`
:::

**Logo:**

| Name, shorthand | Optional value | Description | Required |
| ----------- | ---- | ------------------ | ---- |
| --from | | The address used to sign the transaction sent by the execution function | true |
| --chain-id | plugchain | ID of the chain | true |
| --fees | 200plug | Pay transaction fees | true |
| --home | ~/.plugchain | The directory where the chain data is located | Optional |


## issue-class

Issue assets

```bash
plugchaind tx nft issue-class [class-id] [class-name] [class-symbol] [mint-restricted] [edit-restricted] [schema-content or path to schema.json] [flags]
```

## issue-nft

Additional issuance (creation) of assets

```bash
plugchaind tx nft issue-nft [class-id] [nft-id] [nft-name] [nft-url] [nft-data] [nft-recipient] [flags]
```


## edit-nft

Edit assets

```bash
plugchaind tx nft edit-nft [class-id] [nft-id] [flags]
```

**Logo:**

| Name, shorthand | Default | Description | Required |
| ---------- | ---- | ------------------ | ---- |
| --nft-name | | The name of the nft asset | |
| --nft-url | | Where nft assets are located | |
| --nft-data | |nft asset metadata | |

## transfer-class

Transfer assets

```bash
plugchaind tx nft transfer-class <class-id> <recipient-address> [flags]
```

## transfer-nft

Transfer of designated nft assets

```bash
plugchaind tx nft transfer-nft <class-id> <nft-id> <recipient-address> [flags]
```

## burn-nft

Destroy assets

```bash
plugchaind tx nft burn-nft [class-id] [nft-id] [flags]
```


## supply

Query the total amount of assets based on Denom ID; accept the optional --owner parameter.


```bash
plugchaind q nft supply [class-id] [flags]
plugchaind q nft supply [class-id] --owner=<myAddress> [flags]
```

## owner

Query all assets owned by an account; you can specify the Denom ID parameter.


```bash
plugchaind query nft owner [address] [class-id] [flags]
```

## collection

Query all assets based on Denom ID.


```bash
plugchaind q nft collection [class-id] [flags]
```

## class

Query asset category information based on Denom ID.


```bash
plugchaind q nft class [class-id] [flags]
```

## classes

Query all asset class information that has been issued.

```bash
plugchaind q nft classes [flags]
```

## nft

Query specific assets based on Denom ID and ID.

```bash
plugchaind q nft nft [class-id] [nft-id] [flags]
```