# Plug Chain 文档

## Plug Chain 简介

互联网早已成为了人们生活中必不可少的工具，许多人早已经习惯于从互联网获取所需要的知识，不论是准确的或是虚假的。而且大部分情况下我们不会去判别互联网带给我们的知识的真实性，也不会去主动去梳理杂乱的知识灌输。全世界的数据传输已经非常的方便，与此同时虚假的消息也保持了极高的传播速度。如今，人们亟需在纷乱的互联网世界中快速的获取真实数据，同样的在区块链的世界亦是如此。

区块链行业已经经过多年的发展，如今从金融、产品溯源、政务、民生扩展到智慧城市建设等方面，技术应用场景不断深入化和多元化，已经和大众息息相关了。但是现有所有应用链都存在着一个致命的问题——无法和真实世界进行可信的数据互通。而将应用链与外部系统打通，推动区块链技术应用场景真实落地的工具，就是Oracle Network Protocol预言机。

在区块链的悬浮半空、预言机的数据跨链交互、公链的实体商业场景应用多重需求下，PLUG 研发团队致力于构建一条能够在区块链与区块链、区块链与现实世界之间，实现无感交互的高性能公链，也就是 Plug Chain。

Plug Chain 是 Cosmos Hub 的 Cosmos SDK 应用程序的名称。 它带有 1 个主要入口点：

- `plugchaind` 守护进程和命令行界面 (CLI)。 运行 Plug Chain 应用程序的全节点。

Plug Chain 使用以下模块构建在 Cosmos SDK 上：

- x/auth：帐户和签名。
- x/bank：代币转账。
- x/staking：Staking 逻辑。
- x/mint：通货膨胀逻辑。
- x/distribution：费用分配逻辑。
- x/slashing：削减逻辑。
- x/gov：治理逻辑。
- ibc-go/modules：区块链间通信。 托管在 [cosmos/ibc-go](https://github.com/cosmos/ibc-go) 存储库中。
- x/params：处理应用级参数。
- liquidity:  自动化做市商。 托管在 [oracleNetworkProtocol/liquidity](https://github.com/oracleNetworkProtocol/liquidity) 存储库中。
- pvm: 智能合约虚拟机. 托管在 [ethermint](https://github.com/evmos/ethermint) 存储库中.

以下模块构建在 OracleNetworkProtocol 仓库上：

- x/prc10: 代币逻辑。
- x/nft: 链下资产上链。