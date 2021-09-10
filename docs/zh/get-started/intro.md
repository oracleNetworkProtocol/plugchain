---
order: 1
---

# 简介

## plugchain 网络

*plugchain* 网络是一个区块链互联网，它旨在提供便于构建分布式商业应用程序的技术基础。


*plugchain* 网络是更大的Cosmos网络的一部分 - 网络中所有分区都能够通过标准IBC协议与Cosmos网络中的任何其他分区进行交互。通过在网络中引入一层服务语义，我们将提供一种创新的解决方案，支持全新的业务场景，从而增加Cosmos网络的规模和多样性。

## plugchain 

在*plugchain*网络的“中心”是一个称为 *plugchain* 的区块链，它是一个基于Cosmos SDK和Tendermint构建的Proof-of-Stake（PoS）区块链。它将成为第一个连接Cosmos枢纽的区域性枢纽。*plugchain* 枢纽配备了服务协议，该协议将链上的交易处理与链下的数据处理和业务逻辑执行进行协调。我们还增强了IBC协议，以促进那些链下服务在有需要的情况下被跨链调用。

服务协议和增强的IBC协议最终可以回馈到Cosmos SDK中，允许SDK用户开发与*plugchain*网络兼容的区块链。

我们也会提供面向客户的、针对特定编程语言的SDK，方便在plug网络内轻松提供和使用链下服务。

## plugchain 通证

*plugchain* 枢纽有自己的原生通证，称为 *plugchain*，在网络中有三个作用。

* **抵押。** 与Cosmos Hub中的ATOM通证类似，*plugchain* 通证将用作抵押通证以保护PoS区块链的安全运行。

* **交易费用。** *plugchain* 通证也将用于支付*plugchain* 网络中所有交易的费用。

* **服务费。** *plugchain* 网络中的服务提供者需要以*plugchain* 通证为单位收取服务费。

*plugchain* 网络最终将支持来自Cosmos网络的所有列入白名单的费用通证，它们可用于支付交易费用和服务费用。
