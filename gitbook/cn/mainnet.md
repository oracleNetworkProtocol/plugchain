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
