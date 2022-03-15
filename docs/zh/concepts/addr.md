---
order: 5
---

# 账户

本文档描述了 Plug Chain 的内置账户系统。

## 先决条件

- [Cosmos SDK Accounts](https://docs.cosmos.network/master/basics/accounts.html)
- [Ethereum Accounts](https://ethereum.org/en/whitepaper/#ethereum-accounts) 


## 地址和公钥

Plug Chain 上默认提供了3种主要的地址：
- 账户的地址和密钥，用于识别用户（例如:交易的发起者 `sender`）。它们是使用`secp256k1`或`eth_secp256k1`曲线导出的。
- 验证者操作员的地址和密钥，用于标识验证者的操作员，他们使用的是 `secp256k1`曲线导出的。 [v1.0提案](https://www.plugchain.network/v2/communityDetail?id=7)升级之后支持`eth_secp256k1`。
- 共识节点的地址和密钥，用于标识参与共识的验证节点。它们是使用`ed25519`曲线导出的。

|                    | 地址bech32前缀 | 公钥bech32前缀 | 公钥类型           | 地址字节长度 | 公钥字节长度 |
|--------------------|-----------------------|----------------------|-----------------|---------------------|--------------------|
| 账户           | `gx`               | `gxpub`           | `secp256k1`,`eth_secp256k1` | `20`                | `33` (compressed)  |
| 验证者操作员 | `gxvaloper`        | `gxvaloperpub`    | `secp256k1`,`eth_secp256k1` | `20`                | `33` (compressed)  |
| 共识节点    | `gxvalcons`        | `gxvalconspub`    | `ed25519`       | `20`                | `32`               |


## 账户地址

Plug Chain 使用 `BaseAccount`,`EthAccount` 两种账户类型。


### BaseAccount 

```go 
// BaseAccount defines a base account type. It contains all the necessary fields
// for basic account functionality. Any custom account type should extend this
// type for additional functionality (e.g. vesting).
type BaseAccount struct {
	Address       string     `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	PubKey        *types.Any `protobuf:"bytes,2,opt,name=pub_key,json=pubKey,proto3" json:"public_key,omitempty"`
	AccountNumber uint64     `protobuf:"varint,3,opt,name=account_number,json=accountNumber,proto3" json:"account_number,omitempty"`
	Sequence      uint64     `protobuf:"varint,4,opt,name=sequence,proto3" json:"sequence,omitempty"`
}
```

- 基于[BIP44](https://github.com/bitcoin/bips/blob/master/bip-0044.mediawiki)的分层确定性钱包。
- BIP44基于[BIP32](https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki)中描述的算法定义确定性钱包的逻辑层次结构，该算法允许处理多种代币，多个帐户，每个帐户的外部和内部链以及每个链的数百万个地址，例如比特币和以太坊。
- pubkey类型 `secp256k1` 。
- 账户的根 HD 路径是 `44'/118'/0'/0/0` 。 
- coin-type是 118 。
- 支持 [PRC-10](./token.md#prc-10)协议的，不支持[PRC-20](./token.md#prc-20)协议
- bech32标识为（gx...）

:::warning
后续会不支持创建此类型账户
:::


### EthAccount 

```go 
// EthAccount implements the authtypes.AccountI interface and embeds an
// authtypes.BaseAccount type. It is compatible with the auth AccountKeeper.
type EthAccount struct {
	*types.BaseAccount `protobuf:"bytes,1,opt,name=base_account,json=baseAccount,proto3,embedded=base_account" json:"base_account,omitempty" yaml:"base_account"`
	CodeHash           string `protobuf:"bytes,2,opt,name=code_hash,json=codeHash,proto3" json:"code_hash,omitempty" yaml:"code_hash"`
}
```

- 使用ETH的ECDSA secp256k1曲线作为密钥，这完全满足 [EIP84](https://github.com/ethereum/EIPs/issues/84) [BIP44](https://github.com/bitcoin/bips/blob/master/bip-0044.mediawiki) 路径。
- pubkey类型 `eth_secp256k1` 。
- 帐户的根 HD 路径是 `m/44'/60'/0'/0`。 
- coin-type是 60 。
- 支持 [PRC-10](./token.md#prc-10)，[PRC-20](./token.md#prc-20)协议
- bech32标识为 （gx...）
- 16进制标识为（0x...），用于EVM的Web3工具的兼容性



## 账户地址格式

Bech32 格式是 Cosmos-SDK 通过 CLI 和 REST 查询和交易的默认格式
客户。 另一方面，十六进制格式是以太坊的`common.Address`,
Cosmos的 `sdk.AccAddress`。

- **Address (Bech32)**: `gx1z3t55m0l9h0eupuz3dp5t5cypyv674jj7mz2jw`
- **Address ([EIP55](https://eips.ethereum.org/EIPS/eip-55) Hex)**: `0x91defC7fE5603DFA8CC9B655cF5772459BF10c6f`
- **Compressed Public Key**: `{"@type":"/ethermint.crypto.v1.ethsecp256k1.PubKey","key":"AsV5oddeB+hkByIJo/4lZiVUgXTzNfBPKC73cZ4K1YD2"}` 或 `{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"ApdvZ5Mwsb+sNbOFrZC4MOkvnpSe71Nci2kLaRiovZdv"}`

## 地址转换

可使用 `plugchaind debug addr <address>` 用于在 hex 和 bech32 格式之间转换地址。例如：


:::: tabs
::: tab Bech32

```bash
plugchaind debug addr gx1z3t55m0l9h0eupuz3dp5t5cypyv674jj7mz2jw
  Address: [20 87 74 109 255 45 223 158 7 130 139 67 69 211 4 9 25 175 86 82]
  Address (hex): 14574A6DFF2DDF9E07828B4345D3040919AF5652
  Bech32 Acc: gx1z3t55m0l9h0eupuz3dp5t5cypyv674jj7mz2jw
  Bech32 Val: gxvaloper1z3t55m0l9h0eupuz3dp5t5cypyv674jjn4d6nn
```
:::
::: tab Hex

```bash
plugchaind debug addr 14574A6DFF2DDF9E07828B4345D3040919AF5652
  Address: [20 87 74 109 255 45 223 158 7 130 139 67 69 211 4 9 25 175 86 82]
  Address (hex): 14574A6DFF2DDF9E07828B4345D3040919AF5652
  Bech32 Acc: gx1z3t55m0l9h0eupuz3dp5t5cypyv674jj7mz2jw
  Bech32 Val: gxvaloper1z3t55m0l9h0eupuz3dp5t5cypyv674jjn4d6nn
```

## CLI 创建地址

```bash
plugchaind keys add ethacc1 --algo eth_secp256k1 --coin-type 60 --home node/node1                                                              
- name: ethacc1
  type: local
  address: gx1t97d7kc3wph4wymg3c64qg4ltxhv89kdkty3es
  pubkey: '{"@type":"/ethermint.crypto.v1.ethsecp256k1.PubKey","key":"AxFvcgrK0vnSXUuNIDnT9FcGqczH6VQlMman1OY/wMw7"}'
  mnemonic: ""
```

## 导入EthAccount类型账户

```bash
plugchaind keys add exportethacc1 --algo eth_secp256k1 --coin-type 60 --recover  --home node/node1                                                                                                  
> Enter your bip39 mnemonic

- name: exportethacc1
  type: local
  address: gx1t97d7kc3wph4wymg3c64qg4ltxhv89kdkty3es
  pubkey: '{"@type":"/ethermint.crypto.v1.ethsecp256k1.PubKey","key":"AxFvcgrK0vnSXUuNIDnT9FcGqczH6VQlMman1OY/wMw7"}'
  mnemonic: ""
```

## 导出EthAccount类型私钥

导出需要指出 `--keyring-backend` 账户存储类型
```bash
plugchaind keys unsafe-export-eth-key ethacc1 --home node/node1 --keyring-backend os
2DBECA8D722F93C8DDD2CAAC9FA26121B88958569F1078BD12789AE61970811A
```