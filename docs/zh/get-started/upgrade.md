---
order: 7
---

# Plug Chain Hub 最新版本升级

本文档描述了验证器和全节点操作员成功执行*升级计划*的步骤。由于升级导致的资产丢失，官方概不负责，请升级时备份好自己的财产。
Plug Chain 将在区块高度 `3000000` 停止主链运行，进行正式版本`v1.0`升级。


升级内容如下：
1. 升级主网，修改chain-id 为 `plugchain_520-1`
2. 修改链上币名 `plug` 为:（链上相关币数据是精度0的数据，精度6的币名用于钱包，浏览器，外部APP应用展示）
   - 精度0:  `uplugcn`
   - 精度6:  `plugcn`
3. 修改社区提案模块（gov） `募资` 和 `投票期的时间`，整体修改为14天
4. 调整最大验证者个数为`50`个
5. 接入EVM模块
6. 开通销毁`uplugcn`功能
7. 整体调整`x/liquidity`，`x/token`等模块手续费参数
8. 钱包支持两种密钥签名算法 `secp256k1`，`eth_secp256k1`

注意事项：
1. 不保留交易记录
2. 保留当前块高高度，不保留原先块高数据
3. 保留提案信息
4. 保留liquidity项目数据，保留token数据 





# 备份

*首先需要停止节点运行*

1. 验证者节点
 - 备份整个数据目录，默认 `.plugchain/`
 - 备份 `.plugchain/data/priv_validator_state.json` 文件，`.plugchain/data/`里面的数据就可以删除。
 
2. 其余功能节点
 - 备份钱包目录


# 操作步骤 

1. 获取 `v1.0.0` 版本的二进制文件plugchaind

```bash
# 拉取v1.0.0版本代码 （本地可使用 `git tag` 查看下tag版本，如果有 `v1.0.0`,跳过此步骤）
git fetch origin v1.0.0

# 执行编译二进制文件
make install

```

2. 创建新的数据目录,可使用 --home 指定数据目录

```bash
plugchaind init myNode --chain-id plugchain_520-1
```

3. 下载主网公开的 `genesis.json`,`app.toml`,`config.toml`:


```bash 
curl -o ~/.plugchain/config/genesis.json https://raw.githubusercontent.com/oracleNetworkProtocol/plugchain/main/mainnet/v1/genesis.json
curl -o ~/.plugchain/config/app.toml https://raw.githubusercontent.com/oracleNetworkProtocol/plugchain/main/mainnet/v1/app.toml
curl -o ~/.plugchain/config/config.toml https://raw.githubusercontent.com/oracleNetworkProtocol/plugchain/main/mainnet/v1/config.toml
```

*第四步只需要验证者节点操作，其余节点跳过*

4. 验证者节点需要把 原来节点config目录的 `node_key.json`,`priv_validator_key.json` 文件覆盖到新数据目录的config里面作为识别验证者，其余节点不需要。

5. *根据原来节点的配置修改当下配置满足自己的需求，需要操作钱包的，需要把备份的钱包目录导入到新数据目录才可以操作。*


6. 启动节点

```bash
plugchaind start
```

