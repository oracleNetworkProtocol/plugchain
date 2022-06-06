---
order: 3
---

# account

This document describes Plug Chain's built-in account system.

## prerequisites

- [Cosmos SDK Accounts](https://docs.cosmos.network/master/basics/accounts.html)
- [Ethereum Accounts](https://ethereum.org/en/whitepaper/#ethereum-accounts)


## address and public key

There are 3 main addresses provided by default on Plug Chain:
- The address and key of the account, which is used to identify the user (eg: `sender` for the transaction). They are exported using the `secp256k1` or `eth_secp256k1` curves.
- The address and key of the validator operator, used to identify the validator's operator, they are derived using the `secp256k1` curve. [v1.0 proposal](https://www.plugchain.network/v2/communityDetail?id=7) Support `eth_secp256k1` after upgrade.
- The address and key of the consensus node, used to identify the validator nodes participating in the consensus. They are exported using the `ed25519` curve.

|                    | Address bech32 Prefix | Pubkey bech32 Prefix | Curve           | Address byte length | Pubkey byte length |
|--------------------|-----------------------|----------------------|-----------------|---------------------|--------------------|
| Accounts           | `gx`               | `gxpub`           | `secp256k1`,`eth_secp256k1` | `20`                | `33` (compressed)  |
| Validator Operator | `gxvaloper`        | `gxvaloperpub`    | `secp256k1`,`eth_secp256k1` | `20`                | `33` (compressed)  |
| Consensus Nodes    | `gxvalcons`        | `gxvalconspub`    | `ed25519`       | `20`                | `32`               |

## account address

Plug Chain uses `BaseAccount`, `EthAccount` two account types.


### BaseAccount

```go
// BaseAccount defines a base account type. It contains all the necessary fields
// for basic account functionality. Any custom account type should extend this
// type for additional functionality (e.g. vesting).
type BaseAccount struct {
Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
PubKey *types.Any `protobuf:"bytes,2,opt,name=pub_key,json=pubKey,proto3" json:"public_key,omitempty"`
AccountNumber uint64 `protobuf:"varint,3,opt,name=account_number,json=accountNumber,proto3" json:"account_number,omitempty"`
Sequence uint64 `protobuf:"varint,4,opt,name=sequence,proto3" json:"sequence,omitempty"`
}
```


```bash
# plugchaind q account gx1rpyxd0yqfkqcm8pmp0nejpeacd5t7usk26d2h2
# http://8.210.180.240:1317/cosmos/auth/v1beta1/accounts/gx1rpyxd0yqfkqcm8pmp0nejpeacd5t7usk26d2h2

'@type': /cosmos.auth.v1beta1.BaseAccount
account_number: "1"
address: gx1rpyxd0yqfkqcm8pmp0nejpeacd5t7usk26d2h2
pub_key:
  '@type': /cosmos.crypto.secp256k1.PubKey
  key: AlcccAL+NKRmkvu1Hvmt5uSDzXQEMmOhu7YPy1RwnaXU
sequence: "0"
```


- Hierarchical deterministic wallet based on [BIP44](https://github.com/bitcoin/bips/blob/master/bip-0044.mediawiki).
- BIP44 defines a logical hierarchy of deterministic wallets based on the algorithm described in [BIP32](https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki), which allows the processing of multiple tokens , multiple accounts, external and internal chains per account, and millions of addresses per chain, such as Bitcoin and Ethereum.
- pubkey type `secp256k1` .
- The root HD path of the account is `m/44'/118'/0'/0/0` .
- coin-type is 118 .
- Those that support the [PRC-10](./token.md#prc-10) protocol do not support the [PRC-20](./token.md#prc-20) protocol
- bech32 identified as (gx...)



### EthAccount

```go
// EthAccount implements the authtypes. AccountI interface and embeds an
// authtypes.BaseAccount type. It is compatible with the auth AccountKeeper.
type EthAccount struct {
*types.BaseAccount `protobuf:"bytes,1,opt,name=base_account,json=baseAccount,proto3,embedded=base_account" json:"base_account,omitempty" yaml:"base_account"`
CodeHash string `protobuf:"bytes,2,opt,name=code_hash,json=codeHash,proto3" json:"code_hash,omitempty" yaml:"code_hash"`
}
```

```bash
# plugchaind q account gx1tr5gxpl3c78qp4xkkmhw5p9tmwruvte773ync5
# http://8.210.180.240:1317/cosmos/auth/v1beta1/accounts/gx1tr5gxpl3c78qp4xkkmhw5p9tmwruvte773ync5

'@type': /ethermint.types.v1.EthAccount
base_account:
  account_number: "0"
  address: gx1tr5gxpl3c78qp4xkkmhw5p9tmwruvte773ync5
  pub_key:
    '@type': /ethermint.crypto.v1.ethsecp256k1.PubKey
    key: AmIFRfAboGW0P1GSG+8b9m8aMM1ikl4da4vEakglaLep
  sequence: "1"
code_hash:0xc5d2440186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470
```


- Use ETH's ECDSA secp256k1 curve as the key, which fully satisfies [EIP84](https://github.com/ethereum/EIPs/issues/84) [BIP44](https://github.com/bitcoin/bips/ blob/master/bip-0044.mediawiki) path.
- pubkey type `eth_secp256k1` .
- The root HD path of the account is `m/44'/60'/0'/0`.
- coin-type is 60 .
- Support [PRC-10](./token.md#prc-10), [PRC-20](./token.md#prc-20) protocol
- bech32 is identified as (gx...)
- Hex identifier (0x...) for compatibility with EVM's Web3 tools



## Account address format

Bech32 format is the default format for Cosmos-SDK queries and transactions via CLI and REST
client. On the other hand, the hex format is Ethereum's `common.Address`,
`sdk.AccAddress` for Cosmos.

- **Address (Bech32)**: `gx1z3t55m0l9h0eupuz3dp5t5cypyv674jj7mz2jw`
- **Address ([EIP55](https://eips.ethereum.org/EIPS/eip-55) Hex)**: `0x91defC7fE5603DFA8CC9B655cF5772459BF10c6f`
- **Compressed Public Key**: `{"@type":"/ethermint.crypto.v1.ethsecp256k1.PubKey","key":"AsV5oddeB+hkByIJo/4lZiVUgXTzNfBPKC73cZ4K1YD2"}` or `{"@type": "/cosmos.crypto.secp256k1.PubKey","key":"ApdvZ5Mwsb+sNbOFrZC4MOkvnpSe71Nci2kLaRiovZdv"}`

## address translation

`plugchaind debug addr <address>` can be used to convert addresses between hex and bech32 formats. E.g:



```bash
plugchaind debug addr gx1z3t55m0l9h0eupuz3dp5t5cypyv674jj7mz2jw
  Address: [20 87 74 109 255 45 223 158 7 130 139 67 69 211 4 9 25 175 86 82]
  Address (hex): 14574A6DFF2DDF9E07828B4345D3040919AF5652
  Bech32 Acc: gx1z3t55m0l9h0eupuz3dp5t5cypyv674jj7mz2jw
  Bech32 Val: gxvaloper1z3t55m0l9h0eupuz3dp5t5cypyv674jjn4d6nn
```

```bash
plugchaind debug addr 14574A6DFF2DDF9E07828B4345D3040919AF5652
  Address: [20 87 74 109 255 45 223 158 7 130 139 67 69 211 4 9 25 175 86 82]
  Address (hex): 14574A6DFF2DDF9E07828B4345D3040919AF5652
  Bech32 Acc: gx1z3t55m0l9h0eupuz3dp5t5cypyv674jj7mz2jw
  Bech32 Val: gxvaloper1z3t55m0l9h0eupuz3dp5t5cypyv674jjn4d6nn
```

## CLI create address

```bash
plugchaind keys add ethacc1 --algo eth_secp256k1 --coin-type 60 --home node/node1
- name: ethacc1
  type: local
  address: gx1t97d7kc3wph4wymg3c64qg4ltxhv89kdkty3es
  pubkey: '{"@type":"/ethermint.crypto.v1.ethsecp256k1.PubKey","key":"AxFvcgrK0vnSXUuNIDnT9FcGqczH6VQlMman1OY/wMw7"}'
  mnemonic: ""
```

## Import EthAccount type account

```bash
plugchaind keys add exportethacc1 --algo eth_secp256k1 --coin-type 60 --recover --home node/node1
> Enter your bip39 mnemonic

- name: exportethacc1
  type: local
  address: gx1t97d7kc3wph4wymg3c64qg4ltxhv89kdkty3es
  pubkey: '{"@type":"/ethermint.crypto.v1.ethsecp256k1.PubKey","key":"AxFvcgrK0vnSXUuNIDnT9FcGqczH6VQlMman1OY/wMw7"}'
  mnemonic: ""
```

## Export the private key of EthAccount type

Export needs to indicate `--keyring-backend` account storage type
```bash
plugchaind keys unsafe-export-eth-key ethacc1 --home node/node1 --keyring-backend os
2DBECA8D722F93C8DDD2CAAC9FA26121B88958569F1078BD12789AE61970811A
```