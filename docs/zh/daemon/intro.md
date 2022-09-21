---
order: 1
---

# 简介

plugchaind可执行程序是运行Plug Chain Hub节点的入口，包括验证人节点和其他全节点都需要通过安装`plugchaind`，并启动守护进程来加入到Plug Chain网络。你也可以使用`plugchaind`在本地启动自己的测试网络。

## 硬件要求

建议在Linux服务器上运行Plug Chain Hub节点。

### 最低硬件要求

- 4 CPU
- 内存：8GB
- 磁盘：500GB SSD
- 系统：Ubuntu 16.04 LTS +
- 带宽：5Mbps
- 允许TCP端口26656和26657的所有传入连接

## 主目录

主目录是plug节点的工作目录。主目录包含所有配置信息和节点运行的所有数据。

在`plugchaind`命令中，可以使用标识`--home`来指定节点的主目录。如果在同一台计算机上运行多个节点，则需要为其指定不同的主目录。如果plug中未指定`--home`标识，则默认使用`$HOME/.plugchain`作为主目录。

`plugchaind init`命令负责初始化指定的`--home`目录并创建默认配置文件。除`plugchaind init`命令外，任何其他`plugchaind`子命令使用的主目录必须初始化，否则将报错。

Plug Chain Hub节点的数据存储在主目录的“data”目录中，包括区块链数据，应用程序层数据和索引数据。

所有配置文件都存储在`<home-dir>/config`目录中：

### genesis.json

genesis.json定义了创世块数据，该数据定义了系统参数，例如chain_id，共识参数，初始帐户通证分配，验证人的创建以及各模块的参数。详细信息参见[genesis-file](../concepts/genesis-file.md)。

### node_key.json

node_key.json用于存储节点的密钥。`plugchaind tendermint show-node-id`查询的节点ID由该密钥派生，该ID是节点的唯一标识。它用于p2p连接。

### priv_validator.json

pri_validator.json是[Tendermint Key](../concepts/validator-faq.md#tendermint-密钥)文件，验证人节点将在每轮共识投票中使用该文件来签署Pre-vote/Pre-commit。随着共识的进行，tendermint共识引擎将不断更新`last_height` /`last_round` /`last_step`值。

### config.toml

config.toml是节点的非共识配置。不同的节点可以根据自己的情况进行配置。常见的修改是`persistent_peers`、`moniker`、`laddr`

### app.toml

app.toml为Plug Chain Hub提供了基础配置、监控配置、API配置、同步状态配置和gRPC配置。
