# NFT

## Introduction

`NFT` provides the ability to digitize assets. Through this module, each off-chain asset will be modeled as a unique on-chain asset.

Assets on the chain are identified by `ID`. With the help of the secure and non-tamperable features of the blockchain, the ownership of the assets will be clarified. The transaction process of assets among members will also be publicly recorded to facilitate traceability and dispute settlement.

Asset metadata (`metadata`) can be stored directly on the chain, or it can be stored on the chain with the `URL` of its storage source outside the chain. Asset metadata is organized according to a specific [JSON Schema](https://JSON-Schema.org/). [Here](https://github.com/oracleNetworkProtocol/plugchain/tree/main/docs/zh/features/nft-metadata.json) is an example of metadata JSON Schema.

Assets need to be issued before they are created to declare their abstract properties:

-_Class_: the globally unique asset class name
  
-_Class ID_: Class’s globally unique identifier

-_Metadata Specification_: The JSON Schema that asset metadata should follow

Each specific asset is described by the following elements:

-_Class_: the category of the asset

-_ID_: the identifier of the asset, unique in this asset class

-_Metadata_: The structure that contains asset-specific data

-_Metadata URL_: When metadata is stored off-chain, this URL indicates its storage location

## [Open functions and usage](../cli-client/nft.md)