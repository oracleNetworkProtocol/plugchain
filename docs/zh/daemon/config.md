---
order: 3
---

# 设置client配置

## 简介

plugchain升级到v1.0+之后，终端操作可使用`config`命令来配置好当前数据目录的client配置，或者直接修改 $HOME/.plugchain/config/client.toml 文件，来满足操作命令的一些参数设置。无需每次操作命令指定参数。

## 用法

```bash
 plugchaind config <key> [value] [flags]
```


| 名称，速记        | 类型   | 必须 | 默认值       | 描述                                                                               |
| ----------------- | ------ | ---- | ------------ | ---------------------------------------------------------------------------------- |
| --home            | string |      | $HOME/.plugchain  | 指定存储配置和区块链数据的目录                                                     |


- 设置当前目录主链ID

```bash
 plugchaind config chain-id "plugchain_520-1" --home=<path-to-your-home>
```

- 查询当前目录的主链ID
```bash
 plugchaind config chain-id --home=<path-to-your-home>
```

- 可操作性的字段如下：
```bash
# The network chain ID
chain-id = "plugchain_520-1"
# The keyring's backend, where the keys are stored (os|file|kwallet|pass|test|memory)
keyring-backend = "os"
# CLI output format (text|json)
output = "text"
# <host>:<port> to Tendermint RPC interface for this chain
node = "tcp://localhost:26657"
# Transaction broadcasting mode (sync|async|block)
broadcast-mode = "sync"
```

- 设置完之后，操作命令涉及到需要`chain-id`,`keyring-backend`等参数时，不需要额外指定。以数据目录为 `node` 为例 如下：

```bash
plugchaind tx bank send adr adr1 20000uplugcn --fees 200uplugcn --home node
```