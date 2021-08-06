---
title: 加入测试网
---
## PlugChain 测试网

测试网现在处于活动状态，以下是一些重要的详细信息：

#### 手动加入测试网
1. 根据教程确认plugchaind命令已成功安装
2. plugchaind init 时 chain-id 设置为 `plugchain-testnet-1`
3. 下载`https://github.com/oracleNetworkProtocol/plugchain/testnet/latest/genesis.json`替换 --home 目录下的`config/genesis.json` //--home 默认目录为`~/.plugchain`
4. 根据 `testnet/latest` 目录中提供的 `seeds` 修改--home目录下的 `config/config.toml` 中的seeds参数以设置链接的种子节点 


#### 脚本安装并加入测试网：

1. 根据 `testnet/latest` 目录中提供的 `seeds` 修改 `testnet/scripts/testnet-val-setup.sh` 脚本中的SEEDS参数
2. 运行`testnet/scripts/testnet-val-setup.sh`脚本

```shell
cd testnet/scripts && chmod +x testnet-val-setup.sh
./testnet-val-setup.sh
```