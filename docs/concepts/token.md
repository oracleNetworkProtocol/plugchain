---
order: 4
---

# token standard


## introduce

Plug Chain is a Cosmos SDK based chain with full PVM support. Due to this architecture, tokens and assets in the network may come from different independent sources: bank module, token module, evm module.
`bank module, token module token belongs to the PRC-10 protocol`, `pvm module belongs to the PRC-20 protocol`.

##PC

PC is the most important cryptocurrency on Plug Chain, with a wide range of application scenarios. The rewards on the network are issued in the form of PC. Users can only obtain resources and voting rights by pledging PC. In the DeFi lending market, PC is also used as the main collateral, in the NFT market, PC is used as a unit of account, and so on.

Plug Chain allows developers to build decentralized applications, also called DAPP, which share limited Plug Chain resources. Therefore, Plug Chain needs a mechanism to prevent DAPP from accidentally or maliciously occupying all network resources.

### Cast PC
Minting is the process of creating new PCs on the Plug Chain network. Only the Plug Chain network protocol can create new PCs, and users cannot create PCs.

When a validating node produces a block on the Plug Chain network, a certain amount of PC is minted. For each new block, the current network has about 130PC block rewards. The block reward is a dynamic parameter of the Plug Chain network, which can be modified through proposals.

### Burn PC
PCs can be destroyed through a process called "burning". When a PC is burned, it is permanently subtracted.

Every transaction on the Plug Chain needs to burn PCs to pay for the resources required for the transaction. PC burning can not only help reduce PC inflation, but also prevent accidental or malicious transactions from occupying Plug Chain resources.

### PC denomination
Since many transactions on the Plug Chain involve relatively small amounts, the Plug Chain introduces the smallest currency unit `uplugcn`. The technical implementation of many applications is calculated based on `uplugcn`. The conversion between PC and uplugcn is as follows:

```js
1 pc = 1000000 uplugcn
```
### Query PC balance

- Restful API
```
curl -X GET "http://124.248.67.122:1317/cosmos/bank/v1beta1/balances/gx1v6574fxhha5cwrhkgv6ddlmraue0fc8thr3gpe/by_denom?denom=uplugcn" -H "accept: application/json"
```


## PRC-10
PRC-10 is a token that is built into the Plug Chain public chain. Compared with PRC-20 tokens, PRC-10 tokens face the problem of user experience flexibility. In the Plug Chain network, each account can issue PRC-10 tokens through the [`x/prc10`](../cli-client/token.md) module. Users can lock their tokens individually. To issue tokens, the issuer needs to specify the token name, total size, accuracy, description, etc.

Tokens `pc`, `dhw1`, `kingdm`, `joey`, etc. belong to PRC-10 protocol tokens

### Query PRC-10 balance

- Restful API
```
curl -X GET "http://124.248.67.122:1317/cosmos/bank/v1beta1/balances/gx1v6574fxhha5cwrhkgv6ddlmraue0fc8thr3gpe/by_denom?denom=dhw1" -H "accept: application/json"
```


## PRC-20
PRC-20 is a set of contract standards formulated for the issuance of token assets, that is, contracts written in compliance with this standard are considered to be a PRC-20 contract. When all kinds of wallets and exchanges are docking the assets of the PRC-20 contract, from this set of contract standards, you can know which functions and events are defined in this contract, so as to facilitate the docking.
### Contract Standard

```js
contract PRC20 {
     function name() public view returns (string)
     function symbol() public view returns (string)
     function decimals() public view returns (uint8)
     function totalSupply() constant returns (uint theTotalSupply);
     function balanceOf(address_owner) constant returns (uint balance);
     function transfer(address_to, uint_value) returns (bool success);
     function transferFrom(address _from, address _to, uint _value) returns (bool success);
     function approve(address _spender, uint _value) returns (bool success);
     function allowance(address _owner, address _spender) constant returns (uint remaining);
     event Transfer(address indexed _from, address indexed _to, uint _value);
     event Approval(address indexed _owner, address indexed _spender, uint _value);
}

```
## PRC-721

PRC-721 is a standard interface for issuing non-fungible tokens (NFT) on the Plug Chain public chain, which is fully compatible with ERC-721. Since each token in PRC-721 is unique, the PRC-20 standard is insufficient to handle NFTs.

### Interfaces that must be implemented

Every smart contract conforming to the PRC-721 standard must implement the PRC721 and PRC165 interfaces.

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

If the contract needs to accept secure transfers, it must implement the PRC721TokenReceiver interface:

```js
interface PRC721TokenReceiver {
      //This method will be triggered when the '_to' is the contract address during the 'safeTransferFrom' execution, and the return value must be checked, If the return value is not bytes4(keccak256("onPRC721Received(address,address,uint256, bytes)")") throws an exception. The smart contract which can receive NFT must implement the PRC721TokenReceiver interface.
        function onPRC721Received(address _operator, address _from, uint256 _tokenId, bytes _data) external returns(bytes4);
    }
```

### Metadata Extensions interface (optional)

The metadata extension interface is optional for the PRC-721 smart contract, and users can query the name of the smart contract and the details of the assets represented by the NFT.

```js
interface PRC721Metadata {
      function name() external view returns (string _name);
      function symbol() external view returns (string _symbol);
      function tokenURI(uint256 _tokenId) external view returns (string);
   }
```
### enumeration extension interface (optional)

The enumeration extension is optional for PRC-721 smart contracts, allowing a user's smart contract to publish the full listing of its NFT and make it visible.

```js
interface PRC721Enumerable {
     function totalSupply() external view returns (uint256);
     function tokenByIndex(uint256 _index) external view returns (uint256);
     function tokenOfOwnerByIndex(address _owner, uint256 _index) external view returns (uint256);
   }
```
## PRC1155

Contract standard interface for multiple token management. A single deployed contract can include any combination of fungible tokens, non-fungible tokens, or other configurations such as semi-fungible tokens.

### Interfaces that must be implemented

```js
pragma solidity ^0.5.9;

interface PRC1155 {
     event TransferSingle(address indexed _operator, address indexed _from, address indexed _to, uint256 _id, uint256 _value);
     event TransferBatch(address indexed _operator, address indexed _from, address indexed _to, uint256[] _ids, uint256[] _values);
     event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved);
     event URI(string_value, uint256 indexed_id);
   
     function safeTransferFrom(address _from, address _to, uint256 _id, uint256 _value, bytes calldata _data) external;
     function safeBatchTransferFrom(address _from, address _to, uint256[] calldata _ids, uint256[] calldata _values, bytes calldata _data) external;
     function balanceOf(address _owner, uint256 _id) external view returns (uint256);
     function balanceOfBatch(address[] calldata _owners, uint256[] calldata _ids) external view returns (uint256[] memory);
     function setApprovalForAll(address _operator, bool _approved) external;
     function isApprovedForAll(address _owner, address _operator) external view returns (bool);
}
```

If the contract needs to accept secure transfers, it must implement the PRC1155TokenReceiver interface:
```js
interface PRC1155TokenReceiver {
     function onPRC1155Received(address _operator, address _from, uint256 _id, uint256 _value, bytes calldata _data) external returns(bytes4);
     function onPRC1155BatchReceived(address _operator, address _from, uint256[] calldata _ids, uint256[] calldata _values, bytes calldata _data) external returns(bytes4);
}
```
### Metadata Extensions interface (optional)
```js
interface PRC1155Metadata_URI {
     function uri(uint256 _id) external view returns (string memory);
}
```

## contract example

- [Contract Template](https://github.com/oracleNetworkProtocol/PVM-Contract-Template/)

## Deploy the contract

- Please move to [Contract Deployment](../pvm/prc-20-contract-zh.md)