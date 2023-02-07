---
order: 4
---

# 代币标准


## 介绍

Plug Chain 是一个基于 Cosmos SDK 的链，具有完整的 PVM 支持。由于这种架构，网络中的代币和资产可能来自不同的独立来源：bank模块，token模块，evm模块。
`bank模块，token模块代币 属于PRC-10协议`，`pvm模块属于PRC-20协议`。

## PC

PC是Plug Chain 上最主要的加密货币，应用场景广泛 ，网络上的奖励以PC的形式发放，用户只能通过质押PC来获得资源以及投票权。而在DeFi借贷市场中，PC还被用作主要的抵押品，在NFT市场中，PC被用作记账单位，等等。

Plug Chain允许开发人员构建去中心化应用程序，也叫DAPP，这些应用程序共享有限的Plug Chain资源，因此，Plug Chain需要一种机制来防止DAPP意外或恶意地占用所有网络资源。

### 铸造PC
铸造就是在Plug Chain网络上创建新PC的过程。只有Plug  Chain网络协议可以创建新的PC，用户不可能创建出PC。

当一个验证这节点在Plug Chain网络上生产一个区块后，一定量的PC被铸造。对于每个新区块，目前网络大约有130PC出块奖励。出块奖励为Plug Chain网络动态参数，可以通过提案提议来修改。

### 燃烧PC
PC可以通过一个叫做“燃烧”的过程被摧毁。当PC被燃烧时，它会被永久减除。

Plug Chain上的每笔交易都需要燃烧PC来支付交易所需的资源。PC的燃烧不但可以有助于降低PC的通胀，而且还可以防止意外或者恶意的交易占用Plug Chain资源。

### PC面额
由于Plug Chain上的许多交易涉及到的金额较小，因此Plug Chain引入了最小的货币单位`uplugcn`，许多应用的技术实现都是基于`uplugcn`进行计算的，PC和uplugcn的换算如下：

```
1 pc = 1000000 uplugcn
```
### 查询PC余额

- Restful API
```
curl -X GET "http://124.248.67.122:1317/cosmos/bank/v1beta1/balances/gx1v6574fxhha5cwrhkgv6ddlmraue0fc8thr3gpe/by_denom?denom=uplugcn" -H  "accept: application/json"
```

## PRC-10
PRC-10是一种是通过 Plug Chain 公链内置的通证。与PRC-20代币相比，PRC-10代币面临用户 体验灵活性问题。在 Plug Chain 网络中，每个帐户都能够通过[`x/prc10`](../cli-client/token.md)模块发行PRC-10代币。 用户可以单独锁定其代币。 要发放代币，发行者需要指定代币名称、总大小、精度、描述、等信息。

代币 `pc`,`dhw1`,`kingdm`,`joey`等都属于 PRC-10 协议代币

### 查询PRC-10余额

- Restful API
```
curl -X GET "http://124.248.67.122:1317/cosmos/bank/v1beta1/balances/gx1v6574fxhha5cwrhkgv6ddlmraue0fc8thr3gpe/by_denom?denom=dhw1" -H  "accept: application/json"
```

## PRC-20
PRC-20是为发行通证资产而制定的一套合约标准，即遵守这一标准编写的合约都被认为是一个PRC-20合约。当各类钱包、交易所在对接PRC-20合约的资产时，从这套合约标准中就可以知道这个合约定义了哪些函数、事件，从而方便的进行对接。
### 合约标准

```js
contract PRC20 {
    function name() public view returns (string)
    function symbol() public view returns (string)
    function decimals() public view returns (uint8)
    function totalSupply() constant returns (uint theTotalSupply);
    function balanceOf(address _owner) constant returns (uint balance);
    function transfer(address _to, uint _value) returns (bool success);
    function transferFrom(address _from, address _to, uint _value) returns (bool success);
    function approve(address _spender, uint _value) returns (bool success);
    function allowance(address _owner, address _spender) constant returns (uint remaining);
    event Transfer(address indexed _from, address indexed _to, uint _value);
    event Approval(address indexed _owner, address indexed _spender, uint _value);
}

```

name()
这个方法返回发布合约时设置的通证名称

symbol()
这个方法返回通证缩写

decimals()
这个方法返回通证精度

totalSupply()
这个方法返回通证总的发行量。

balanceOf(address _owner)
这个方法返回查询账户的通证余额。

transfer(address _to, uint _value)
这个方法用来从智能合约地址里转账通证到指定账户。

approve(address _spender, uint _value)
这个方法用来授权第三方（例如DAPP合约）从通证拥有者账户转账通证。

transferFrom(address _from, address _to, uint _value)
这个方法可供第三方从通证拥有者账户转账通证。需要配合approve()方法使用。

allowance(address _owner, address _spender)
这个方法用来查询可供第三方转账的查询账户的通证余额。

事件函数
当通证被成功转账后，会触发Transfer转账事件。

event Transfer(address indexed _from, address indexed _to, uint256 _value)
当approve()方法被成功调用后，会触发Approval事件。

event Approval(address indexed _owner, address indexed _spender, uint256 _value)

## PRC-721

PRC-721是在Plug Chain公链上发行非同质化代币（non-fungible token, NFT）一套标准接口，与ERC-721完全兼容。由于 PRC-721 中的每个代币都是唯一的，因此 PRC-20 标准不足以处理 NFT。

### 必须实现的接口

每个符合PRC-721标准的智能合约都必须实现PRC721与PRC165接口。

```js
   pragma solidity ^0.4.20;

  interface PRC721 {
    event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId);
    event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId);
    event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved);

    function balanceOf(address _owner) external view returns (uint256);
    function ownerOf(uint256 _tokenId) external view returns (address);
    function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes data) external payable;
    function safeTransferFrom(address _from, address _to, uint256 _tokenId) external payable;
    function transferFrom(address _from, address _to, uint256 _tokenId) external payable;
    function approve(address _approved, uint256 _tokenId) external payable;
    function setApprovalForAll(address _operator, bool _approved) external;
    function getApproved(uint256 _tokenId) external view returns (address);
    function isApprovedForAll(address _owner, address _operator) external view returns (bool);
  }
  interface PRC165 {
      function supportsInterface(bytes4 interfaceID) external view returns (bool);
  }
```
balanceOf(address _owner)
返回指定账户拥有的 NFT 数量

ownerOf(uint256 _tokenId)
返回指定 NFT 的所有者

safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes data)
转让 NFT 的所有权

safeTransferFrom(address _from, address _to, uint256 _tokenId)
转让 NFT 的所有权

transferFrom(address _from, address _to, uint256 _tokenId)
转让一个NFT的所有权（调用者必须确认_to地址是否可以正常接收NFT，否则NFT会丢失）

approve(address _approved, uint256 _tokenId)
授予其他人 NFT 的控制权

setApprovalForAll(address _operator, bool _approved)
由第三方 (_operator) 授予/恢复对所有 NFT 的控制权

getApproved(uint256 _tokenId)
查询某个 NFT 的授权

isApprovedForAll(address _owner, address _operator)
查询operator是否为owner的授权地址

supportsInterface(bytes4 interfaceID)
查询是否支持某个接口（interfaceID）

event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId)
Approve 成功后会触发 Approval 事件

event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
成功的 transferFrom 和 safeTransferFrom 将触发 Transfer 事件

event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
setApprovalForAll 成功后会触发 ApprovalForAll 事件

合约如果需要接受安全转账，必须实现PRC721TokenReceiver接口:

```js
interface PRC721TokenReceiver {
     //This method will be triggered when the ‘_to’ is the contract address during the ‘safeTransferFrom’ execution, and the return value must be checked, If the return value is not  bytes4(keccak256("onPRC721Received(address,address,uint256,bytes)")) throws an exception. The smart contract which can receive NFT must implement the PRC721TokenReceiver interface.
       function onPRC721Received(address _operator, address _from, uint256 _tokenId, bytes _data) external                returns(bytes4);
   }
```
onPRC721Received(address _operator, address _from, uint256 _tokenId, bytes _data)

与safeTransferFrom方法配合使用，当_to为合约地址时，需要调用该方法并检查返回值。 如果返回值不是 bytes4(keccak256("onPRC721Received(address,address,uint256,bytes)")) 将抛出异常。 可以接收 NFT 的智能合约必须实现 PRC721TokenReceiver 接口。

:::warning
bytes4(keccak256("onPRC721Received(address,address,uint256,bytes))) 的哈希值与以太坊版本 bytes4(keccak256("onERC721Received(address,address,uint256,bytes))) 不同。 对于函数onPRC721Received的返回值，请使用0x14ce6b03而不是0x150b7a02。
:::


### Metadata Extensions接口（可选）

metadata extension接口对于PRC-721智能合约来说是可选的，用户可以查询智能合约的名称以及NFT代表的资产的详细信息。

```js
interface PRC721Metadata {
     function name() external view returns (string _name);
     function symbol() external view returns (string _symbol);
     function tokenURI(uint256 _tokenId) external view returns (string);
  }
```
name()
返回合约名称

symbol()
返回合约代码

tokenURI(uint256 _tokenId)
返回 _tokenId 对应的外部文件的 URI。 外部资源文件需要包括名称、描述和图片。

### enumeration extension接口（可选）

enumeration extension对于PRC-721智能合约是可选的，允许用户的智能合约发布其NFT的完整列表并使其可见。

```js
interface PRC721Enumerable  {
    function totalSupply() external view returns (uint256);
    function tokenByIndex(uint256 _index) external view returns (uint256);
    function tokenOfOwnerByIndex(address _owner, uint256 _index) external view returns (uint256);
  }
```

totalSupply()
返回 NFT 的总量

tokenByIndex(uint256 _index)
通过_index返回对应的tokenId

tokenOfOwnerByIndex(address _owner, uint256 _index)
返回所有者拥有的 NFT 列表中索引对应的 tokenId


## PRC1155

用于多种代币管理的合约标准接口。 单个部署的合约可以包括同质化代币、非同质化代币或其他配置（如半同质化代币）的任何组合。

### 必须实现的接口

```js
pragma solidity ^0.5.9;

interface PRC1155 {
    event TransferSingle(address indexed _operator, address indexed _from, address indexed _to, uint256 _id, uint256 _value);
    event TransferBatch(address indexed _operator, address indexed _from, address indexed _to, uint256[] _ids, uint256[] _values);
    event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved);
    event URI(string _value, uint256 indexed _id);
   
    function safeTransferFrom(address _from, address _to, uint256 _id, uint256 _value, bytes calldata _data) external;
    function safeBatchTransferFrom(address _from, address _to, uint256[] calldata _ids, uint256[] calldata _values, bytes calldata _data) external;
    function balanceOf(address _owner, uint256 _id) external view returns (uint256);
    function balanceOfBatch(address[] calldata _owners, uint256[] calldata _ids) external view returns (uint256[] memory);
    function setApprovalForAll(address _operator, bool _approved) external;
    function isApprovedForAll(address _owner, address _operator) external view returns (bool);
}
```

function safeTransferFrom(address _from, address _to, uint256 _id, uint256 _value, bytes calldata _data)
安全单币转账，将amount单位id种类的代币从from地址转账给to地址。如果to地址是合约，则会验证是否实现了onERC1155Received()接收函数。

function safeBatchTransferFrom(address _from, address _to, uint256[] calldata _ids, uint256[] calldata _values, bytes calldata _data)
安全多币转账，与单币转账类似，只不过转账数量amounts和代币种类ids变为数组，且长度相等。如果to地址是合约，则会验证是否实现了onERC1155BatchReceived()接收函数。

function balanceOf(address _owner, uint256 _id)
单币种余额查询，返回account拥有的id种类的代币的持仓量。

function balanceOfBatch(address[] calldata _owners, uint256[] calldata _ids)
多币种余额查询，查询的地址accounts数组和代币种类ids数组的长度要相等。

function setApprovalForAll(address _operator, bool _approved)
批量授权，将调用者的代币授权给operator地址。。

function isApprovedForAll(address _owner, address _operator)
查询批量授权信息，如果授权地址operator被account授权，则返回true。

event TransferSingle(address indexed _operator, address indexed _from, address indexed _to, uint256 _id, uint256 _value)
单类代币转账事件，在单币种转账时释放。

event TransferBatch(address indexed _operator, address indexed _from, address indexed _to, uint256[] _ids, uint256[] _values)
批量代币转账事件，在多币种转账时释放。

event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
批量授权事件，在批量授权时释放。

event URI(string _value, uint256 indexed _id)
元数据地址变更事件，在uri变化时释放。



合约如果需要接受安全转账，必须实现PRC1155TokenReceiver接口:
```js
interface PRC1155TokenReceiver {
    function onPRC1155Received(address _operator, address _from, uint256 _id, uint256 _value, bytes calldata _data) external returns(bytes4);
    function onPRC1155BatchReceived(address _operator, address _from, uint256[] calldata _ids, uint256[] calldata _values, bytes calldata _data) external returns(bytes4);       
}
```
function onPRC1155Received(address _operator, address _from, uint256 _id, uint256 _value, bytes calldata _data)
单币转账接收函数，接受PRC1155安全转账safeTransferFrom 需要实现并返回自己的选择器`0xca78c6ac`。

function onPRC1155BatchReceived(address _operator, address _from, uint256[] calldata _ids, uint256[] calldata _values, bytes calldata _data)
多币转账接收函数，接受PRC1155安全多币转账safeBatchTransferFrom 需要实现并返回自己的选择器`0xd12c9bf8`。


###  Metadata Extensions接口（可选）
```js
interface PRC1155Metadata_URI {
    function uri(uint256 _id) external view returns (string memory);
}
```
function uri(uint256 _id)
返回第`id`种类代币的URI




## 合约示例
- [模板](https://github.com/oracleNetworkProtocol/PVM-Contract-Template/)
## 部署合约

- 请移步到[合约部署](../pvm/prc-20-contract-zh.md)

