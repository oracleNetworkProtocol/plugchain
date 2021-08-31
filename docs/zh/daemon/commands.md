---
order: 2
---

# 命令

## 简介

plugchaind守护程序命令允许您初始化，启动，重置节点或生成创世文件等。

你可以通过创建[Local Testnet](local-testnet.md)来熟悉这些命令。

## 用法

```bash
plugchaind <command>
```

## 可用命令

| Name                                                             | Description                                                               |
| ---------------------------------------------------------------- | ------------------------------------------------------------------------- |
| [init](local-testnet.md#plugchaind-init)                               | 初始化验证人，p2p，创世纪和应用程序配置文件                               |
| [add-genesis-account](local-testnet.md#plugchaind-add-genesis-account) | 将创世帐户添加到genesis.json                                              |
| [gentx](local-testnet.md#plugchaind-gentx)                             | 生成自抵押的创始TX                                                        |
| [collect-gentxs](local-testnet.md#plugchaind-collect-gentxs)           | 收集创世txs并输出到genesis.json文件                                       |
| [start](local-testnet.md#plugchaind-start)                             | 启动全节点                                                                |
| [unsafe-reset-all](local-testnet.md#plugchaind-unsafe-reset-all)       | 重置区块链数据库，删除address book，并将priv_validator.json重置为创始状态 |
| [tendermint](local-testnet.md#plugchaind-tendermint)                   | Tendermint子命令                                                          |
| [testnet](local-testnet.md#build-and-init)                       | 初始化PLUGChain Hub测试网                                                       |
| [reset](local-testnet.md#plugchaind-reset)                             | 将app状态重置到指定的高度                                                 |
| [export](export.md)                                              | 将状态导出为JSON                                                          |
| version                                                          | 显示版本信息                                                              |

## 全局标识

| 名称，速记  | 默认值       | 描述                                             | 必须 | 类型   |
| ----------- | ------------ | ------------------------------------------------ | ---- | ------ |
| -h, --help  |              | 打印帮助信息                                     |      |        |
| --home      | /$HOME/.plugchain| 配置和数据目录                                   |      | String |
| --log_level | \*:info      | 日志级别（默认为"main:info,state:info,*:error"） |      | String |
| --trace     |              | 错误是打印出完整的堆栈信息                       |      |        |
