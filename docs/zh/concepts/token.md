---
order: 4
---

# 代币标准


## 介绍

Plug Chain 是一个基于 Cosmos SDK 的链，具有完整的 PVM 支持。由于这种架构，网络中的代币和资产可能来自不同的独立来源：bank模块，token模块，evm模块。
`bank模块，token模块代币 属于PRC-10协议`，`pvm模块属于PRC-20协议`。

## PC

在 Plug Chain 上用于质押、治理和 gas 消耗的面额是PC. PC提供以下效用：保护权益证明链、用于治理提案的代币、费用分配以及作为在 PVM 上运行智能合约的气体手段。

$$ 1 pc = 1 ~ * ~ 10^{6} uplugcn $$

## PRC-10
PRC-10是一种是通过 Plug Chain 公链内置的通证。 PRC-10是 Plug Chain 区块链本身支持的技术代币标准,没有使用PVM虚拟机, 在 Plug Chain 网络中，每个帐户都能够通过[`x/prc10`](../cli-client/token.md)模块发行PRC-10代币。 用户可以单独锁定其代币。 要发放代币，发行者需要指定代币名称、总大小、精度、描述、等信息。

代币 `pc`,`dhw1`,`kingdm`,`joey`等都属于 PRC-10 协议代币


## PRC-20

PRC-20是在 Plug Chain 区块链上通过`Pvm`部署智能合约的方式来发行资产的一套标准，兼容[ERC-20](https://github.com/ethereum/EIPs/blob/master/EIPS/eip-20.md) 。


## PRC-721

PRC-721是在 Plug Chain 区块链上通过`Pvm`部署智能合约的方式来发行非同质化资产NFT的一套标准，兼容[ERC-721](https://github.com/ethereum/EIPs/blob/master/EIPS/eip-721.md) 。



