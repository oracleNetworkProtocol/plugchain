# Token

Token 模块用于管理你在Plug Chain Hub上发行的[PRC-10](../concepts/token.md#prc-10)资产。

## 可用命令

| 名称                                       | 描述                     |
| ------------------------------------------ | ------------------------ |
| [issue](#plugchaind-tx-token-issue)              | 发行通证                 |
| [edit](#plugchaind-tx-token-edit)                | 编辑通证                 |
| [transfer](#plugchaind-tx-token-transfer)        | 转让通证所有权           |
| [mint](#plugchaind-tx-token-mint)                | 增发通证到指定账户       |
| [burn](#plugchaind-tx-token-burn)                | 销毁通证                 |
| [token](#plugchaind-query-token-token)           | 查询通证                 |
| [tokens](#plugchaind-query-token-tokens)         | 查询指定所有者的或者平台所有的通证集合 |
| [fee](#plugchaind-query-token-fee)               | 查询通证相关费用         |
| [params](#plugchaind-query-token-params)         | 查询通证相关参数         |

## plugchaind tx token issue

发行一个新通证。

```bash
plugchaind tx token issue [flags]
```

**标识：**

| 名称,速记       | 类型    | 必须 | 默认          | 描述                                                               |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------ |
| --name           | string  | 是   |               | 通证的名称,限制为32个unicode字符,例如“Plug Chain Network”              |
| --symbol         | string  | 是   |               | 币名缩写,长度在3到8之间,字母数字字符,以字符开始,不区分大小写 |
| --initial-supply | uint64  | 是   |               | 此通证的初始供应                    |
| --max-supply     | uint64  |      | 100000000000 | 通证上限,总供应不能超过最大供应。 增发前的数量不应超过1000万亿       |
| --scale          | uint8   | 是   |               | 通证最多可以有8位小数                                             |
| --min-unit       | string  |      |               | 币名最小单位缩写（可与 `--symbol`相同）,长度在3到8之间,字母数字字符,以字符开始,不区分大小写 |
| --mintable       | boolean |      | false         | 首次发行后是否可以增发此通证                                       |

### 发行通证

```bash
plugchaind tx token issue \
    --name="Test Token" \
    --symbol="token" \
    --min-unit="token" \
    --scale=6 \
    --initial-supply=1000000000 \
    --max-supply=100000000000 \
    --mintable=true \
    --from=<key-name> \
    --chain-id=<chain-id> \
    --fees=<fee>
```

### 发送通证

您可以像[发送PC](./bank.md#plugchaind-tx-bank-send)一样发送任何通证。

#### 发送通证

```bash
plugchaind tx bank send [from_key_or_address] [to_address] [amount] [flags]
```

## plugchaind tx token edit

编辑通证。

```bash
plugchaind tx token edit [symbol] [flags]
```

**标识：**

| 名称,速记   | 类型   | 必须 | 默认  | 描述                          |
| ------------ | ------ | ---- | ----- | ----------------------------- |
| --name       | string |      |       | 通证名称,例如：Plug Chain Network  |
| --max-supply | uint   |      | 0     | 通证的最大供应量              |

`max-supply` 不得少于当前的总供应量。

### 编辑通证

```bash
plugchaind tx token edit <symbol> --name="Cat Token" --max-supply=100000000000 --mintable=true --from=<key-name> --chain-id=<chain-id> --fees=<fee>
```

## plugchaind tx token transfer

转让通证所有权。

```bash
plugchaind tx token transfer [symbol] [flags]
```

**标识：**

| 名称,速记 | 类型   | 必须 | 默认 | 描述       |
| ---------- | ------ | ---- | ---- | ---------- |
| --to       | string | 是   |      | 接收人地址 |

### 转让通证所有者

```bash
plugchaind tx token transfer <symbol> --to=<to> --from=<key-name> --chain-id=<chain-id> --fees=<fee>
```

## plugchaind tx token mint

增发通证到指定地址。

```bash
plugchaind tx token mint [symbol] [flags]
```

**标识：**

| 名称,速记 | 类型   | 必须 | 默认 | 描述                                       |
| ---------- | ------ | ---- | ---- | ------------------------------------------ |
| --to       | string |      |      | 增发的通证的接收地址,默认为发起该交易地址 |
| --amount   | uint64 | 是   | 0    | 增发的数量                                 |

### 增发通证

```bash
plugchaind tx token mint <symbol> --amount=<amount> --to=<to> --from=<key-name> --chain-id=<chain-id> --fees=<fee>
```

## plugchaind tx token burn

增发通证到指定地址。

```bash
plugchaind tx token burn [symbol] [flags]
```

**标识：**

| 名称,速记 | 类型   | 必须 | 默认 | 描述       |
| ---------- | ------ | ---- | ---- | ---------- |
| --amount   | uint64 | 是   | 0    | 销毁的数量 |

### 销毁通证

```bash
plugchaind tx token burn <symbol> --amount=<amount> --from=<key-name> --chain-id=<chain-id> --fees=<fee>
```

## plugchaind query token token

查询通证。

```bash
plugchaind query token token [denom] [flags]
```

### 查询通证

```bash
plugchaind query token token <denom>
```

## plugchaind query token tokens

查询指定所有者的通证集合。所有者是可选的。

```bash
plugchaind query token tokens [owner] [flags]
```

### 查询所有通证

```bash
plugchaind query token tokens
```

### 查询指定所有者的通证

```bash
plugchaind query token tokens <owner>
```

## plugchaind query token fee

查询与通证相关的费用,包括通证发行和增发。

```bash
plugchaind query token fee [symbol] [flags]
```

### 查询发行和增发通证的费用

```bash
plugchaind query token fee kitty
```

## plugchaind query token params

查询通证模块参数。

```bash
plugchaind query token params [flags]
```

### 查询通证模块参数

```bash
plugchaind query token params
```
