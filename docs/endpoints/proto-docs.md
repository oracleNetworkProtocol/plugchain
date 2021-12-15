---
order: 6
---
# Protobuf Documentation
<!-- This file is auto-generated. Please do not modify it yourself. -->
<a name="top"></a>

## Table of Contents

- [nft/nft.proto](#nft/nft.proto)
    - [Class](#plugchain.nft.Class)
    - [Collection](#plugchain.nft.Collection)
    - [CollectionID](#plugchain.nft.CollectionID)
    - [NFT](#plugchain.nft.NFT)
    - [Owner](#plugchain.nft.Owner)
  
- [nft/genesis.proto](#nft/genesis.proto)
    - [GenesisState](#plugchain.nft.GenesisState)
  
- [nft/query.proto](#nft/query.proto)
    - [QueryClassRequest](#plugchain.nft.QueryClassRequest)
    - [QueryClassResponse](#plugchain.nft.QueryClassResponse)
    - [QueryClassesRequest](#plugchain.nft.QueryClassesRequest)
    - [QueryClassesResponse](#plugchain.nft.QueryClassesResponse)
    - [QueryCollectionRequest](#plugchain.nft.QueryCollectionRequest)
    - [QueryCollectionResponse](#plugchain.nft.QueryCollectionResponse)
    - [QueryNFTRequest](#plugchain.nft.QueryNFTRequest)
    - [QueryNFTResponse](#plugchain.nft.QueryNFTResponse)
    - [QueryOwnerRequest](#plugchain.nft.QueryOwnerRequest)
    - [QueryOwnerResponse](#plugchain.nft.QueryOwnerResponse)
    - [QuerySupplyRequest](#plugchain.nft.QuerySupplyRequest)
    - [QuerySupplyResponse](#plugchain.nft.QuerySupplyResponse)
  
    - [Query](#plugchain.nft.Query)
  
- [nft/tx.proto](#nft/tx.proto)
    - [MsgBurnNFT](#plugchain.nft.MsgBurnNFT)
    - [MsgBurnNFTResponse](#plugchain.nft.MsgBurnNFTResponse)
    - [MsgEditNFT](#plugchain.nft.MsgEditNFT)
    - [MsgEditNFTResponse](#plugchain.nft.MsgEditNFTResponse)
    - [MsgIssueClass](#plugchain.nft.MsgIssueClass)
    - [MsgIssueClassResponse](#plugchain.nft.MsgIssueClassResponse)
    - [MsgIssueNFT](#plugchain.nft.MsgIssueNFT)
    - [MsgIssueNFTResponse](#plugchain.nft.MsgIssueNFTResponse)
    - [MsgTransferClass](#plugchain.nft.MsgTransferClass)
    - [MsgTransferClassResponse](#plugchain.nft.MsgTransferClassResponse)
    - [MsgTransferNFT](#plugchain.nft.MsgTransferNFT)
    - [MsgTransferNFTResponse](#plugchain.nft.MsgTransferNFTResponse)
  
    - [Msg](#plugchain.nft.Msg)
  
- [token/token.proto](#token/token.proto)
    - [Params](#plugchain.token.Params)
    - [Token](#plugchain.token.Token)
  
- [token/genesis.proto](#token/genesis.proto)
    - [GenesisState](#plugchain.token.GenesisState)
  
- [token/query.proto](#token/query.proto)
    - [QueryFeesRequest](#plugchain.token.QueryFeesRequest)
    - [QueryFeesResponse](#plugchain.token.QueryFeesResponse)
    - [QueryParamsRequest](#plugchain.token.QueryParamsRequest)
    - [QueryParamsResponse](#plugchain.token.QueryParamsResponse)
    - [QueryTokenRequest](#plugchain.token.QueryTokenRequest)
    - [QueryTokenResponse](#plugchain.token.QueryTokenResponse)
    - [QueryTokensRequest](#plugchain.token.QueryTokensRequest)
    - [QueryTokensResponse](#plugchain.token.QueryTokensResponse)
  
    - [Query](#plugchain.token.Query)
  
- [token/tx.proto](#token/tx.proto)
    - [MsgBurnToken](#plugchain.token.MsgBurnToken)
    - [MsgBurnTokenResponse](#plugchain.token.MsgBurnTokenResponse)
    - [MsgEditToken](#plugchain.token.MsgEditToken)
    - [MsgEditTokenResponse](#plugchain.token.MsgEditTokenResponse)
    - [MsgIssueToken](#plugchain.token.MsgIssueToken)
    - [MsgIssueTokenResponse](#plugchain.token.MsgIssueTokenResponse)
    - [MsgMintToken](#plugchain.token.MsgMintToken)
    - [MsgMintTokenResponse](#plugchain.token.MsgMintTokenResponse)
    - [MsgTransferOwnerToken](#plugchain.token.MsgTransferOwnerToken)
    - [MsgTransferOwnerTokenResponse](#plugchain.token.MsgTransferOwnerTokenResponse)
  
    - [Msg](#plugchain.token.Msg)
  
- [Scalar Value Types](#scalar-value-types)



<a name="nft/nft.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nft/nft.proto



<a name="plugchain.nft.Class"></a>

### Class



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [string](#string) |  |  |
| `name` | [string](#string) |  |  |
| `schema` | [string](#string) |  |  |
| `symbol` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |
| `mint_restricted` | [bool](#bool) |  |  |
| `edit_restricted` | [bool](#bool) |  |  |






<a name="plugchain.nft.Collection"></a>

### Collection



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class` | [Class](#plugchain.nft.Class) |  |  |
| `nfts` | [NFT](#plugchain.nft.NFT) | repeated |  |






<a name="plugchain.nft.CollectionID"></a>

### CollectionID



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `Class_id` | [string](#string) |  |  |
| `nft_ids` | [string](#string) | repeated |  |






<a name="plugchain.nft.NFT"></a>

### NFT



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [string](#string) |  |  |
| `name` | [string](#string) |  |  |
| `uri` | [string](#string) |  |  |
| `data` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |






<a name="plugchain.nft.Owner"></a>

### Owner



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |
| `collection_ids` | [CollectionID](#plugchain.nft.CollectionID) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="nft/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nft/genesis.proto



<a name="plugchain.nft.GenesisState"></a>

### GenesisState
GenesisState defines the nft module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collections` | [Collection](#plugchain.nft.Collection) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="nft/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nft/query.proto



<a name="plugchain.nft.QueryClassRequest"></a>

### QueryClassRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class_id` | [string](#string) |  |  |






<a name="plugchain.nft.QueryClassResponse"></a>

### QueryClassResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class` | [Class](#plugchain.nft.Class) |  |  |






<a name="plugchain.nft.QueryClassesRequest"></a>

### QueryClassesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="plugchain.nft.QueryClassesResponse"></a>

### QueryClassesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `classes` | [Class](#plugchain.nft.Class) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="plugchain.nft.QueryCollectionRequest"></a>

### QueryCollectionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class_id` | [string](#string) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="plugchain.nft.QueryCollectionResponse"></a>

### QueryCollectionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collection` | [Collection](#plugchain.nft.Collection) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="plugchain.nft.QueryNFTRequest"></a>

### QueryNFTRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class_id` | [string](#string) |  |  |
| `nft_id` | [string](#string) |  |  |






<a name="plugchain.nft.QueryNFTResponse"></a>

### QueryNFTResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `nft` | [NFT](#plugchain.nft.NFT) |  |  |






<a name="plugchain.nft.QueryOwnerRequest"></a>

### QueryOwnerRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class_id` | [string](#string) |  |  |
| `address` | [string](#string) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="plugchain.nft.QueryOwnerResponse"></a>

### QueryOwnerResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [Owner](#plugchain.nft.Owner) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="plugchain.nft.QuerySupplyRequest"></a>

### QuerySupplyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class_id` | [string](#string) |  |  |






<a name="plugchain.nft.QuerySupplyResponse"></a>

### QuerySupplyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `amount` | [uint64](#uint64) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="plugchain.nft.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Class` | [QueryClassRequest](#plugchain.nft.QueryClassRequest) | [QueryClassResponse](#plugchain.nft.QueryClassResponse) |  | GET|/nft/classes/{class_id}|
| `Classes` | [QueryClassesRequest](#plugchain.nft.QueryClassesRequest) | [QueryClassesResponse](#plugchain.nft.QueryClassesResponse) |  | GET|/nft/classes|
| `NFT` | [QueryNFTRequest](#plugchain.nft.QueryNFTRequest) | [QueryNFTResponse](#plugchain.nft.QueryNFTResponse) |  | GET|/nft/nfts/{class_id}/{nft_id}|
| `Collection` | [QueryCollectionRequest](#plugchain.nft.QueryCollectionRequest) | [QueryCollectionResponse](#plugchain.nft.QueryCollectionResponse) |  | GET|/nft/collections/{class_id}|
| `Supply` | [QuerySupplyRequest](#plugchain.nft.QuerySupplyRequest) | [QuerySupplyResponse](#plugchain.nft.QuerySupplyResponse) |  | GET|/nft/collections/{class_id}/supply|
| `Owner` | [QueryOwnerRequest](#plugchain.nft.QueryOwnerRequest) | [QueryOwnerResponse](#plugchain.nft.QueryOwnerResponse) |  | GET|/nft/nfts/{address}/{class_id}|

 <!-- end services -->



<a name="nft/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nft/tx.proto



<a name="plugchain.nft.MsgBurnNFT"></a>

### MsgBurnNFT



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [string](#string) |  |  |
| `class_id` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |






<a name="plugchain.nft.MsgBurnNFTResponse"></a>

### MsgBurnNFTResponse







<a name="plugchain.nft.MsgEditNFT"></a>

### MsgEditNFT



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [string](#string) |  |  |
| `class_id` | [string](#string) |  |  |
| `name` | [string](#string) |  |  |
| `uri` | [string](#string) |  |  |
| `data` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |






<a name="plugchain.nft.MsgEditNFTResponse"></a>

### MsgEditNFTResponse







<a name="plugchain.nft.MsgIssueClass"></a>

### MsgIssueClass



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [string](#string) |  |  |
| `name` | [string](#string) |  |  |
| `schema` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |
| `symbol` | [string](#string) |  |  |
| `mint_restricted` | [bool](#bool) |  |  |
| `edit_restricted` | [bool](#bool) |  |  |






<a name="plugchain.nft.MsgIssueClassResponse"></a>

### MsgIssueClassResponse







<a name="plugchain.nft.MsgIssueNFT"></a>

### MsgIssueNFT



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [string](#string) |  |  |
| `class_id` | [string](#string) |  |  |
| `name` | [string](#string) |  |  |
| `uri` | [string](#string) |  |  |
| `data` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |
| `recipient` | [string](#string) |  |  |






<a name="plugchain.nft.MsgIssueNFTResponse"></a>

### MsgIssueNFTResponse







<a name="plugchain.nft.MsgTransferClass"></a>

### MsgTransferClass



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |
| `recipient` | [string](#string) |  |  |






<a name="plugchain.nft.MsgTransferClassResponse"></a>

### MsgTransferClassResponse







<a name="plugchain.nft.MsgTransferNFT"></a>

### MsgTransferNFT



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [string](#string) |  |  |
| `class_id` | [string](#string) |  |  |
| `recipient` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |






<a name="plugchain.nft.MsgTransferNFTResponse"></a>

### MsgTransferNFTResponse






 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="plugchain.nft.Msg"></a>

### Msg
Msg defines the Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `IssueClass` | [MsgIssueClass](#plugchain.nft.MsgIssueClass) | [MsgIssueClassResponse](#plugchain.nft.MsgIssueClassResponse) | IssueClass defines a method for issue a class | |
| `IssueNFT` | [MsgIssueNFT](#plugchain.nft.MsgIssueNFT) | [MsgIssueNFTResponse](#plugchain.nft.MsgIssueNFTResponse) | IssueNFT defines a method for Issue a new nft | |
| `EditNFT` | [MsgEditNFT](#plugchain.nft.MsgEditNFT) | [MsgEditNFTResponse](#plugchain.nft.MsgEditNFTResponse) | EditNFT defines a method for edit a nft | |
| `BurnNFT` | [MsgBurnNFT](#plugchain.nft.MsgBurnNFT) | [MsgBurnNFTResponse](#plugchain.nft.MsgBurnNFTResponse) | BurnNFT define a method for burning a nft | |
| `TransferNFT` | [MsgTransferNFT](#plugchain.nft.MsgTransferNFT) | [MsgTransferNFTResponse](#plugchain.nft.MsgTransferNFTResponse) | TransferNFT define a method for transferring a nft | |
| `TransferClass` | [MsgTransferClass](#plugchain.nft.MsgTransferClass) | [MsgTransferClassResponse](#plugchain.nft.MsgTransferClassResponse) | TransferClass define a method for transferring a class | |

 <!-- end services -->



<a name="token/token.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## token/token.proto



<a name="plugchain.token.Params"></a>

### Params



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `issue_token_base_fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `operate_token_fee_ratio` | [string](#string) |  |  |






<a name="plugchain.token.Token"></a>

### Token



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `symbol` | [string](#string) |  |  |
| `name` | [string](#string) |  |  |
| `scale` | [uint32](#uint32) |  |  |
| `min_unit` | [string](#string) |  |  |
| `initial_supply` | [uint64](#uint64) |  |  |
| `max_supply` | [uint64](#uint64) |  |  |
| `mintable` | [bool](#bool) |  |  |
| `owner` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="token/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## token/genesis.proto



<a name="plugchain.token.GenesisState"></a>

### GenesisState
GenesisState defines the token module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#plugchain.token.Params) |  |  |
| `tokens` | [Token](#plugchain.token.Token) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="token/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## token/query.proto



<a name="plugchain.token.QueryFeesRequest"></a>

### QueryFeesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `symbol` | [string](#string) |  |  |






<a name="plugchain.token.QueryFeesResponse"></a>

### QueryFeesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `exist` | [bool](#bool) |  |  |
| `issue_fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `operate_fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="plugchain.token.QueryParamsRequest"></a>

### QueryParamsRequest







<a name="plugchain.token.QueryParamsResponse"></a>

### QueryParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#plugchain.token.Params) |  |  |
| `res` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="plugchain.token.QueryTokenRequest"></a>

### QueryTokenRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |






<a name="plugchain.token.QueryTokenResponse"></a>

### QueryTokenResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `Token` | [google.protobuf.Any](#google.protobuf.Any) |  |  |






<a name="plugchain.token.QueryTokensRequest"></a>

### QueryTokensRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [string](#string) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="plugchain.token.QueryTokensResponse"></a>

### QueryTokensResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `Tokens` | [google.protobuf.Any](#google.protobuf.Any) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="plugchain.token.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Token` | [QueryTokenRequest](#plugchain.token.QueryTokenRequest) | [QueryTokenResponse](#plugchain.token.QueryTokenResponse) |  | GET|/token/tokens/{denom}|
| `Fees` | [QueryFeesRequest](#plugchain.token.QueryFeesRequest) | [QueryFeesResponse](#plugchain.token.QueryFeesResponse) |  | GET|/token/fee/{symbol}|
| `Params` | [QueryParamsRequest](#plugchain.token.QueryParamsRequest) | [QueryParamsResponse](#plugchain.token.QueryParamsResponse) |  | GET|/token/params|
| `Tokens` | [QueryTokensRequest](#plugchain.token.QueryTokensRequest) | [QueryTokensResponse](#plugchain.token.QueryTokensResponse) |  | GET|/token/tokens|

 <!-- end services -->



<a name="token/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## token/tx.proto



<a name="plugchain.token.MsgBurnToken"></a>

### MsgBurnToken
MsgBurnToken defines an SDK message for burning some tokens


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `symbol` | [string](#string) |  |  |
| `amount` | [uint64](#uint64) |  |  |
| `owner` | [string](#string) |  |  |






<a name="plugchain.token.MsgBurnTokenResponse"></a>

### MsgBurnTokenResponse







<a name="plugchain.token.MsgEditToken"></a>

### MsgEditToken
MsgEditToken defines an SDK message for editing a new token


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `symbol` | [string](#string) |  |  |
| `name` | [string](#string) |  |  |
| `max_supply` | [uint64](#uint64) |  |  |
| `owner` | [string](#string) |  |  |






<a name="plugchain.token.MsgEditTokenResponse"></a>

### MsgEditTokenResponse







<a name="plugchain.token.MsgIssueToken"></a>

### MsgIssueToken
MsgIssueToken defines an SDK message for issuing a new token


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `symbol` | [string](#string) |  |  |
| `name` | [string](#string) |  |  |
| `scale` | [uint32](#uint32) |  |  |
| `min_unit` | [string](#string) |  |  |
| `initial_supply` | [uint64](#uint64) |  |  |
| `max_supply` | [uint64](#uint64) |  |  |
| `mintable` | [bool](#bool) |  |  |
| `owner` | [string](#string) |  |  |






<a name="plugchain.token.MsgIssueTokenResponse"></a>

### MsgIssueTokenResponse







<a name="plugchain.token.MsgMintToken"></a>

### MsgMintToken
MsgMintTokenResponse defines the Msg/MintToken response type


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `symbol` | [string](#string) |  |  |
| `to` | [string](#string) |  |  |
| `amount` | [uint64](#uint64) |  |  |
| `owner` | [string](#string) |  |  |






<a name="plugchain.token.MsgMintTokenResponse"></a>

### MsgMintTokenResponse







<a name="plugchain.token.MsgTransferOwnerToken"></a>

### MsgTransferOwnerToken
MsgTransferOwnerToken defines an SDK message for transferring the token owner


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `symbol` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |
| `to` | [string](#string) |  |  |






<a name="plugchain.token.MsgTransferOwnerTokenResponse"></a>

### MsgTransferOwnerTokenResponse






 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="plugchain.token.Msg"></a>

### Msg
Msg defines the Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `IssueToken` | [MsgIssueToken](#plugchain.token.MsgIssueToken) | [MsgIssueTokenResponse](#plugchain.token.MsgIssueTokenResponse) |  | |
| `MintToken` | [MsgMintToken](#plugchain.token.MsgMintToken) | [MsgMintTokenResponse](#plugchain.token.MsgMintTokenResponse) |  | |
| `EditToken` | [MsgEditToken](#plugchain.token.MsgEditToken) | [MsgEditTokenResponse](#plugchain.token.MsgEditTokenResponse) |  | |
| `BurnToken` | [MsgBurnToken](#plugchain.token.MsgBurnToken) | [MsgBurnTokenResponse](#plugchain.token.MsgBurnTokenResponse) |  | |
| `TransferOwnerToken` | [MsgTransferOwnerToken](#plugchain.token.MsgTransferOwnerToken) | [MsgTransferOwnerTokenResponse](#plugchain.token.MsgTransferOwnerTokenResponse) |  | |

 <!-- end services -->



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

