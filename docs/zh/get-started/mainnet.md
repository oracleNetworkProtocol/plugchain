---
order: 3
---

# 加入主网

:::tip
需要先 [安装 plugchaind](install.md)
:::

## 运行全节点

### Start node from genesis

:::tip
必须使用 PLUGChain [v0.2.0](https://github.com/oracleNetworkProtocol/plugchain.git) 初始化你的节点
:::

1. 初始化节点

```bash
plugchaind init <moniker> --chain-id=plugchain
```

2. 下载主网公开的 genesis.json和种子信息 或者 进入clone下来的 plugchain 目录里面：
*[主网创世文件和种子信息](https://github.com/oracleNetworkProtocol/plugchain/blob/main/mainnet/v1)* ,如下通过移动clone的仓库里面的genesis.json 来实现覆盖。

```bash 
mv ./mainnet/v1/genesis.json ~/.plugchain/
```

:::warning 
如下第三步可跳过，如跳过需要在第四步加上参数 `--p2p.seeds` 参数）
:::

3. 覆盖数据目录的genesis.json之后,修改 ~/.plugchain/config/config.toml 里面 seeds 参数，添加种子信息

把 ./mainnet/v1/seeds.txt 文件中提供的 seeds , 修改 ~/.plugchain/config/config.toml 中的`seeds`参数以设置链接的种子节点,种子信息以英文逗号分隔


4. 启动节点服务
```bash
# 启动节点（也可使用 nohup 或 systemd 等方式后台运行）

# 第三步未修改种子信息，运行start时，添加参数 --p2p.seeds="7488f044132cec94e72c0eb5cdd267fb5607f5d1@47.102.107.120:26656,60fde7a070938367ede8943ee45bee622424753a@47.102.126.234:26656"

plugchaind start --minimum-gas-prices 0.0001plug 
```


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
--amount 1000000plug \
--pubkey $(plugchaind tendermint show-validator) \
--moniker="my validator" \
--commission-rate="0.10" \
--commission-max-rate="0.20" \
--commission-max-change-rate="0.01" \
--min-self-delegation="1000000" \
--fees 20plug --chain-id plugchain
```


:::warning
**重要**

一定要备份好 home（默认为〜/.plugchain/）目录中的 `config` 目录！如果您的服务器磁盘损坏或您准备迁移服务器，这是恢复验证人的唯一方法。
:::

如果以上命令没有出现错误，则您的节点已经是验证人或候选人了（取决于您的Voting Power是否在前100名中）

阅读更多：

- 概念
  - [基础概念](../concepts/general-concepts.md)
  - [验证人问答](../concepts/validator-faq.md)
- 验证人安全
  - [哨兵节点 (DDOS 防护)](../concepts/sentry-nodes.md)
  - [密钥管理](../tools/kms.md)