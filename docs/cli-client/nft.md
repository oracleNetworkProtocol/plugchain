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


:::warning

参数规则：
| 名称               | 类型           | 描述   | 规则 |
| ----------------- | -----------   | ----  |  ---- |
| class-id          | string        |  assets class ID  |  only accepts alphanumeric characters, and begin with an english letter. length [3,64] |
| class-name        | string        |  assets class name  | none |
| class-symbol      |  string       |  assets class symbol | none |
| mint-restricted   | bool          |  MintRestricted is true means that only Class owners can issue NFTs under this category, false means anyone can | none |
| edit-restricted   | bool          |  EditRestricted is true means that no one in this category can Edit the NFT, false means that only the owner of this NFT can Edit   | none |
| nft-id            | string        |  nft ID  |  only accepts alphanumeric characters, and begin with an english letter. length [3,64] |
| nft-name          | string        |  nft name | none |
| nft-uri           | string        |  The URI pointing to a JSON object that contains subsequent nftData information off-chain | The maximum length is 256 bytes, starting with `http://`,`https://` |
| nft-data          | string        |  NFT specifications defined under this category   | none |
| nft-recipient     | string        |  owner address | none |

:::



## issue-class

Issue assets

```bash
plugchaind tx nft issue-class [class-id] [class-name] [class-symbol] [mint-restricted] [edit-restricted] [schema-content or path to schema.json] [flags]
```

## issue-nft

Additional issuance (creation) of assets

```bash
plugchaind tx nft issue-nft [class-id] [nft-id] [nft-name] [nft-uri] [nft-data] [nft-recipient] [flags]
```


## edit-nft

Edit assets

```bash
plugchaind tx nft edit-nft [class-id] [nft-id] [flags]
```

**Logo:**

| Name, shorthand | Default | Description | Required |
| ---------- | ---- | ------------------ | ---- |
| --nft-name | | The name of the nft asset | Optional |
| --nft-uri | | Where nft assets are located | Optional |
| --nft-data | |nft asset metadata | Optional |

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