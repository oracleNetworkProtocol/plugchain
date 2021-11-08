# NFT

`NFT`提供了将资产进行数字化的能力。通过该模块，每个链外资产将被建模为唯一的链上资产。

## 可用命令

| 名称                                   | 描述           |
| ------------------------------------  | -------------  |
| [issue-class](#issue-class)           | 创建资产分类       |
| [issue-nft](#issue-nft)               | 发行资产       |
| [edit-nft](#edit-nft)                 | 编辑资产       |
| [transfer-class](#transfer-class)     | 转让资产分类       |
| [transfer-nft](#transfer-nft)         | 转让指定nft资产 |
| [burn-nft](#burn-nft)                 | 销毁资产        |
| [supply](#supply)                     | 查询supply     |
| [owner](#owner)                       | 通过owner查询   |
| [collection](#collection)             | 查询collection |
| [class](#class)                       | 查询denom      |
| [classes](#classes)                   | 查询denoms     |
| [nft](#nft)                           | 查询指定nft     |

:::tip
以上命令都可以使用 `-h` 参数，来查看功能的描述和参数的含义
:::

:::warning 
以上`issue-class`,`issue-nft`,`edit-nft`,`transfer-class`,`transfer-nft`,`burn-nft` 命令都需要参数 `--from`,`--chain-id`,`--fees`
:::

**标志：**

| 名称，速记  | 可选值 | 描述               | 必须 | 
| ----------- | ---- | ------------------ | ---- |
| --from       |      | 执行功能发交易用于签名的地址 |   true   |
| --chain-id |   plugchain   | 链的ID          |    true  |
| --fees      |   200plug   | 支付交易的费用            |    true  |
| --home      |   ~/.plugchain   | 链数据所在目录            |    可选  |


## issue-class

发行资产

```bash
plugchaind tx nft issue-class [class-id] [class-name] [class-symbol] [mint-restricted] [edit-restricted] [schema-content or path to schema.json]  [flags]
```

## issue-nft

增发(创建)资产

```bash
plugchaind tx nft issue-nft [class-id] [nft-id] [nft-name] [nft-url] [nft-data] [nft-recipient] [flags]
```


## edit-nft

编辑资产

```bash
plugchaind tx nft edit-nft [class-id] [nft-id] [flags]
```

**标志：**

| 名称，速记 | 默认 | 描述               | 必须 |
| ---------- | ---- | ------------------ | ---- |
| --nft-name      |      | nft资产的名字 |      |
| --nft-url     |      | nft资产所在的地方           |      |
| --nft-data     |      |nft资产的元数据       |      |

## transfer-class

转让资产

```bash
plugchaind tx nft transfer-class <class-id> <recipient-address> [flags]
```

## transfer-nft

转让指定nft资产

```bash
plugchaind tx nft transfer-nft <class-id> <nft-id> <recipient-address> [flags]
```

## burn-nft

销毁资产

```bash
plugchaind tx nft burn-nft [class-id] [nft-id] [flags]
```


## supply

根据 Denom ID查询资产总量；接受可选的 --owner 参数。


```bash
plugchaind q nft supply [class-id] [flags]
plugchaind q nft supply [class-id] --owner=<myAddress> [flags]
```

## owner

查询某一账户所拥有的全部资产；可以指定 Denom ID参数。


```bash
plugchaind query nft owner [address] [class-id] [flags]
```

## collection

根据 Denom ID查询所有资产。


```bash
plugchaind q nft collection [class-id] [flags]
```

## class

根据 Denom ID查询资产类别信息。


```bash
plugchaind q nft class [class-id] [flags]
```

## classes

查询已发行的所有资产类别信息。

```bash
plugchaind q nft classes [flags]
```

## nft

根据 Denom ID以及 ID 查询具体资产。

```bash
plugchaind q nft nft [class-id] [nft-id] [flags]
```
