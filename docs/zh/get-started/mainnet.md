---
order: 3
---

# 加入主网

:::tip

需要先 [安装 plugchaind](install.md),或者下载对应版本的二进制文件 [releases](https://github.com/oracleNetworkProtocol/plugchain/releases)
:::

## 运行全节点

### 使用创世文件

:::tip
必须使用 Plug Chain [v1.1.0](https://github.com/oracleNetworkProtocol/plugchain.git) 初始化你的节点
:::

1. 初始化节点

```bash
plugchaind init <moniker> --chain-id=plugchain_520-1
```

2. 下载主网公开的 `genesis.json`,`app.toml`,`config.toml`:

```bash 
curl -o ~/.plugchain/config/genesis.json https://raw.githubusercontent.com/oracleNetworkProtocol/mainnet/main/version/v1/genesis.json
curl -o ~/.plugchain/config/app.toml https://raw.githubusercontent.com/oracleNetworkProtocol/mainnet/main/version/v1/app.toml
curl -o ~/.plugchain/config/config.toml https://raw.githubusercontent.com/oracleNetworkProtocol/mainnet/main/version/v1/config.toml
```
3. 启动之前如果想修改服务端口,种子信息，对等点，哨兵模式等，可自行修改文件，然后再同步区块。


###  同步区块

同步主网块数据
#### 一 快照同步

根据快照高度，锁定`plugchaind`二进制版本使用


| 块高 | 数据库  | plugchaind 版本 | 下载地址 | 描述 |
| ---- | --------- | -------- | ----| ----|
| 7171838 | goleveldb (default） | [v1.7](https://github.com/oracleNetworkProtocol/plugchain/releases/tag/v1.7.0) | [mainnet-7171838-20221208-goleveldb.zip](https://snapshot-node-mainnet.oss-cn-hangzhou.aliyuncs.com/mainnet-7171838-20221208-goleveldb.zip) | （107.915GB） 裁剪数据|

1. 下载快照数据

2. 数据覆盖 `~/.plugchain/data/` 目录

3. 使用 对应plugchaind版本启动 `plugchaind start`



#### 二 逐步升级同步
启动节点服务

```bash
# 启动节点（也可使用 nohup 或 systemd 等方式后台运行）

plugchaind start
```


接下来，你的节点将执行所有链升级过程。在每次升级之间，你必须使用特定的版本同步区块。不用担心在升级高度使用旧版本，节点会自动停止。

| 提案 | 起始高度 | 升级高度 | plugchaind 版本 |
| -------- | ------------ | -------------- | ----- |
| [v1.0](https://www.plugchain.network/v2/communityDetail?id=7)  |  3000000     |    | [v1.1.0](https://github.com/oracleNetworkProtocol/plugchain/releases/tag/v1.1.0) |
| [v1.2.1](https://www.plugchain.network/v2/communityDetail?id=8)  |  3349542     |  3576853  | [v1.2.1](https://github.com/oracleNetworkProtocol/plugchain/releases/tag/v1.2.1) |
| [v1.5.0](https://www.plugchain.network/v2/communityDetail?id=9)  |  3935641     |  4152263  | [v1.5.0](https://github.com/oracleNetworkProtocol/plugchain/releases/tag/v1.5.0) |
| [v1.7](https://www.plugchain.network/v2/communityDetail?id=10)  |  5420512     |  5633000  | [v1.7.0](https://github.com/oracleNetworkProtocol/plugchain/releases/tag/v1.7.0) |






:::tip
您可能会看到一些连接错误，这没关系，P2P网络正在尝试查找可用的连接


:::


## 升级为验证人节点

### 创建钱包

您可以[创建新的钱包](../cli-client/keys.md#创建密钥)或[导入现有的钱包](../cli-client/keys.md#通过助记词恢复密钥)，然后从交易所或其他任何地方转入一些plug到您刚刚创建的钱包中：

```bash
# 创建一个新钱包
plugchaind keys add <key-name>
```

:::warning
在安全的地方备份好助记词！如果您忘记密码，这是恢复帐户的唯一方法。
:::

### 确认节点同步状态

```bash
# 可以使用此命令安装 jq
# apt-get update && apt-get install -y jq

# 如果输出为 false，则表明您的节点已经完成同步
plugchaind status 2>&1 | jq -r '.SyncInfo.catching_up'
```

### 创建验证人

只有节点已完成同步时，才可以运行以下命令将您的节点升级为验证人：

```bash
plugchaind tx staking create-validator --from mywallet \
--amount 1000000uplugcn \
--pubkey $(plugchaind tendermint show-validator) \
--moniker="my validator" \
--commission-rate="0.10" \
--commission-max-rate="0.20" \
--commission-max-change-rate="0.01" \
--min-self-delegation="1000000" \
--fees 20uplugcn --chain-id plugchain_520-1
```


:::warning
**重要**

一定要备份好 home（默认为〜/.plugchain/）目录中的 `config` 目录！如果您的服务器磁盘损坏或您准备迁移服务器，这是恢复验证人的唯一方法。
:::

如果以上命令没有出现错误，则您的节点已经是验证人或候选人了（取决于您的Voting Power是否在前50名中）