# NFT

`NFT`提供了将资产进行数字化的能力。通过该模块，每个链外资产将被建模为唯一的链上资产。

## 可用命令

| 名称                                   | 描述           |
| ------------------------------------  | -------------  |
| [issue-class](#issue-class)           | 创建资产分类     |
| [issue-nft](#issue-nft)               | 发行资产        |
| [edit-nft](#edit-nft)                 | 编辑资产        |
| [transfer-class](#transfer-class)     | 转让资产分类     |
| [transfer-nft](#transfer-nft)         | 转让指定nft资产  |
| [burn-nft](#burn-nft)                 | 销毁资产        |
| [supply](#supply)                     | 查询supply     |
| [owner](#owner)                       | 通过owner查询   |
| [nft](#nft)                           | 查询指定nft      |
| [nfts](#nfts)                         | 查询nfts  |
| [class](#class)                       | 查询class       |
| [classes](#classes)                   | 查询classes     |

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
| --chain-id |   plugchain_520-1   | 链的ID          |    true  |
| --fees      |   200uplugcn   | 支付交易的费用            |    true  |
| --home      |   ~/.plugchain   | 链数据所在目录            |    可选  |


:::warning

参数规则：
| 名称               | 类型           | 描述   | 规则 |
| ----------------- | -----------   | ----  |  ---- |
| class-id          | string        |  资产类别序号  | [a-zA-Z][a-zA-Z0-9/:-]{2,100} |
| nft-id            | string        |  资产序号  | [a-zA-Z][a-zA-Z0-9/:-]{2,100} |
| class-name        | string        |  资产类别名称  | 无 |
| class-symbol      |  string       |  资产类别缩写  | 无 |
| mint-restricted   | bool          |  true 表示只有类别拥有者才可以在此类别下创建资产，false 表示任何人都可以| 无 |
| edit-restricted   | bool          |  为 true 表示此类别中没有人可以编辑 NFT，false 表示只有此 NFT 的所有者可以编辑  | 无 |

| nft-name          | string        |  资产名称 | 无 |
| nft-uri           | string        |  资产链外信息的 JSON 对象的 URI | 最大长度256字节，以`http://`,`https://`开头 |
| nft-data          | string        |  资产元数据  | 无 |
| nft-recipient     | string        |  资产创建者 | 无 |

:::


## issue-class

发行资产

```bash
plugchaind tx nft issue-class [class-id] [class-name] [class-symbol] [mint-restricted] [edit-restricted] [schema-content or path to schema.json]  [flags]
```

## issue-nft

增发(创建)资产

```bash
plugchaind tx nft issue-nft [class-id] [nft-id] [nft-name] [nft-uri] [nft-data] [nft-recipient] [flags]
```


## edit-nft

编辑资产

```bash
plugchaind tx nft edit-nft [class-id] [nft-id] [flags]
```

**标志：**

| 名称，速记  | 默认  | 描述                   |  必须 |
| ---------- | ---- | ------------------   |  ---- |
| --nft-name |      | nft资产的名字         | 可选  | 
| --nft-uri  |      | nft资产链外元数据地址  | 可选  |
| --nft-data |      |nft资产的元数据        |  可选  |

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

根据 class-id 查询资产总量；接受可选的 --owner 参数。


```bash
plugchaind q nft supply [class-id] [flags]
plugchaind q nft supply [class-id] --owner=<myAddress> [flags]
```

## owner

查询某一账户所拥有的全部资产；可以指定 class-id 参数。


```bash
plugchaind query nft owner [address] [class-id] [flags]
```

## nft

根据 class-id 以及 ID 查询具体资产。

```bash
plugchaind q nft nft [class-id] [nft-id] [flags]
```

## nfts

根据 class-id 查询所有资产。接受可选的 --owner 参数。

```bash
plugchaind q nft nfts [class-id] [flags]
```

## class

根据 class-id 查询资产类别信息。


```bash
plugchaind q nft class [class-id] [flags]
```

## classes

查询已发行的所有资产类别信息。

```bash
plugchaind q nft classes [flags]
```
