---
order: 3
---

# 本地测试网

出于测试或开发目的，您可能需要运行本地测试网。

## 单节点测试网

**需求**

- [安装plugchaind](../get-started/install.md)

:::tip
对于以下示例，我们全部使用默认的[主目录](intro.md#主目录)
:::

### plugchaind init

初始化 genesis.json 文件，它将帮助你启动网络

```bash
plugchaind init testing --chain-id=plugchain_521-1
```

### 创建一个钱包

创建一个钱包作为您的验证人帐户

```bash
plugchaind keys add MyValidator
```

### plugchaind add-genesis-account

将该钱包地址添加到 genesis 文件中的 genesis.app_state.accounts 数组中

:::tip
此命令使您可以设置通证数量。确保此帐户有 plug，这是 Plug Chain Hub 上唯一的质押通证
:::

```bash
plugchaind add-genesis-account $(plugchaind keys show MyValidator --address) 150000000uplugcn
```

### plugchaind gentx

生成创建验证人的交易。gentx 存储在 `~/.plugchain/config/` 中

```bash
plugchaind gentx MyValidator 100000000uplugcn --chain-id=plugchain_521-1 
```

### plugchaind collect-gentxs

将生成的质押交易添加到创世文件

```bash
plugchaind collect-gentxs
```

### plugchaind start

修改默认token为 `plug`

```bash
sed -i 's/stake/plug/g' $HOME/.plugchain/config/genesis.json
```

现在可以启动 `plugchaind` 了

```bash
plugchaind start
```

### plugchaind unsafe-reset-all

可以使用此命令来重置节点，包括本地区块链数据库，地址簿文件，并将 priv_validator.json 重置为创世状态。

当本地区块链数据库以某种方式中断和无法同步或参与共识时，这是有用的。

```bash
plugchaind unsafe-reset-all
```

### plugchaind tendermint

查询可以在 p2p 连接中使用的唯一节点 ID，例如在 [config.toml](intro.md#cnofig-toml) 中 `seeds` 和 `persistent_peers` 的格式 `<node-id>@ip:26656`。

节点 ID 存储在 [node_key.json](intro.md#node_key-json) 中。

```bash
plugchaind tendermint show-node-id
```

查询 [Tendermint Pubkey](../concepts/validator-faq.md#tendermint-密钥)，用于 [identify your validator](../cli-client/staking.md#plugchaind-tx-staking-create-validator)，并将用于在共识过程中签署 Pre-vote/Pre-commit。

[Tendermint Key](../concepts/validator-faq.md#tendermint-密钥) 存储在 [priv_validator.json](intro.md#priv_validator-json) 中，创建验证人后，请一定要记得[备份](../concepts/validator-faq.md#如何备份验证人节点)。

```bash
plugchaind tendermint show-validator
```

查询bech32前缀验证人地址

```bash
plugchaind tendermint show-address
```

### plugchaind export

请参阅[导出区块状态](export.md)。

<!-- ## 多节点测试网

**前提**

- [安装 plug](../get-started/install.md)
- [安装 jq](https://stedolan.github.io/jq/download/)
- [安装 docker](https://docs.docker.com/engine/installation/)
- [安装 docker-compose](https://docs.docker.com/compose/install/)

### 构建和初始化

```bash
# Work from the Plug Chain Hub repo
cd [your-Plug Chain Hub-repo]

# Build the linux binary in ./build
make build-linux

# Quick init a 4-node testnet configs
make testnet-init
```

`make testnet-init` 将调用 `plugchaind testnet` 命令在 `build/nodecluster` 目录下生成4个节点的测试网配置文件。

```bash
$ tree -L 3 build/nodecluster/
build/nodecluster/
├── gentxs
│   ├── node0.json
│   ├── node1.json
│   ├── node2.json
│   └── node3.json
├── node0
│   ├── plug
│   │   ├── config
│   │   └── data
│   └── plugcli
│       ├── key_seed.json
│       └── keys
├── node1
│   ├── plug
│   │   ├── config
│   │   └── data
│   └── plugcli
│       └── key_seed.json
├── node2
│   ├── plug
│   │   ├── config
│   │   └── data
│   └── plugcli
│       └── key_seed.json
└── node3
    ├── plug
    │   ├── config
    │   └── data
    └── plugcli
        └── key_seed.json
```

### 启动

```bash
make testnet-start
```

该命令将使用 ubuntu:16.04 的 docker 镜像创建4个节点的测试网。下表列出了每个节点的端口：

| Node      | P2P Port | RPC Port |
| --------- | -------- | -------- |
| plugnode0 | 26656    | 26657    |
| plugnode1 | 26659    | 26660    |
| plugnode2 | 26661    | 26662    |
| plugnode3 | 26663    | 26664    |

要更新二进制文件，只需重新构建它并重新启动节点即可：

```bash
make build-linux testnet-start
```

### 停止

停止所有正在运行的节点：

```bash
make testnet-stop
```

### 清理

要停止所有正在运行的节点并删除 `build/` 目录中的所有文件：

```bash
make testnet-clean
``` -->
