---
title: 加入主网
---
## PlugChain 主网

**[!Info]**: 加入主网步骤

1. 安装`plugchaind`
2. 初始化数据目录
3. 替换数据目录的`genesis.json`创世文件，添加种子连接信息
4. 启动节点,同步区块


#### 加入主网

1. 根据[教程](./installation.md)确认`plugchaind`命令已成功安装

2. 使用 `plugchaind init` 初始化数据目录，chain-id 设置为 `plugchain`

```bash
plugchaind init myNode --chain-id plugchain
```

*[主网创世文件和种子信息](https://github.com/oracleNetworkProtocol/plugchain/blob/main/mainnet/v1)*

3. 下载`https://github.com/oracleNetworkProtocol/plugchain/blob/main/mainnet/v1/genesis.json`替换 --home 目录下的`config/genesis.json` //--home 默认目录为`~/.plugchain`

4. 复制 `https://github.com/oracleNetworkProtocol/plugchain/blob/main/mainnet/v1/seeds.txt` 文件中提供的 `seeds` 修改--home目录下的 `config/config.toml` 中的`seeds`参数以设置链接的种子节点 

5. 修改完之后，使用 `plugchaind start --minimum-gas-prices 0.0001plug` 启动节点


6. 查询节点是否同步完区块，返回`false`，表示同步成功。使用：

```shell
plugchaind status 2>&1 | jq -r '.SyncInfo.catching_up'
```

7. 只有节点已完成同步时，才可以将您的节点升级为验证人[成为验证人](./validators/validator-setup.md)

**[!Danger]**:
一定要备份好 home（默认为〜/.plugchain/）目录中的 config 目录！如果您的服务器磁盘损坏或您准备迁移服务器，这是恢复验证人的唯一方法。

