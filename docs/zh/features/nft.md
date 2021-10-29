# NFT

## 简介

`NFT`提供了将资产进行数字化的能力。通过该模块，每个链外资产将被建模为唯一的链上资产。

链上资产用 `ID` 进行标识，借助区块链安全、不可篡改的特性，资产的所有权将得到明确。资产在成员间的交易过程也将被公开地记录，以便于追溯以及争议处理。

资产的元数据（`metadata`）可以直接存储在链上，也可以将其在链外存储源的 `URL` 存储在链上。资产元数据按照特定的 [JSON Schema](https://JSON-Schema.org/) 进行组织。[这里](https://github.com/oracleNetworkProtocol/plugchain/tree/main/docs/zh/features/nft-metadata.json)是一个元数据 JSON Schema 示例。

资产在创建前需要发行，用以声明其抽象属性：

- _Denom_：即全局唯一的资产类别名
  
- _Denom ID_：Demon的全局唯一标识符

- _元数据规范_：资产元数据应遵循的 JSON Schema

每一个具体的资产由以下元素描述：

- _Denom_：该资产的类别

- _ID_：资产的标识符，在此资产类别中唯一

- _元数据_：包含资产具体数据的结构

- _元数据 URL_：当元数据存储在链外时，此 URL 表示其存储位置

## [开放功能和使用方式](../cli-client/nft.md)

