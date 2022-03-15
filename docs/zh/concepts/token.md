---
order: 6
---

# 代币标准


## 介绍

Plug Chain 是一个基于 Cosmos SDK 的链，具有完整的 EVM 支持。由于这种架构，网络中的代币和资产可能来自不同的独立来源：bank模块，token模块，evm模块。
`bank模块，token模块代币 属于PRC-10协议`，`evm模块属于PRC-20协议`。

## PLUGCN

`plugcn`代币属于bank模块原生代币，可用于质押，IBC转账，社区治理，手续费等

## PRC-10
PRC-10是一种是通过 Plug Chain 公链内置的通证。 PRC-10是 Plug Chain 区块链本身支持的技术代币标准,没有使用EVM虚拟机, 在 Plug Chain 网络中，每个帐户都能够通过`x/token`模块发行PRC-10代币。 用户可以单独锁定其代币。 要发放代币，发行者需要指定代币名称、总大小、精度、描述、等信息。

代币 `plugcn`,`dhw1`,`kingdm`,`joey`等都属于 PRC-10 协议代币


## PRC-20

PRC-20是在 Plug Chain 区块链上通过部署智能合约的方式来发行资产的一套标准，它与 ERC-20完全兼容。

