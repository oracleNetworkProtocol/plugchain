# Accounts

本文档介绍了 Cosmos SDK 的内置账号和公钥系统。

## 开始前读物

[Cosmos SDK Accounts](https://docs.cosmos.network/master/basics/accounts.html)

## Keys, accounts, addresses, and signatures

对用户进行身份验证的主要方式是使用 [数字签名](https://en.wikipedia.org/wiki/Digital_signature)。 用户使用自己的私钥签署交易。 签名验证是使用关联的公钥完成的。 出于链上签名验证的目的，我们将公钥存储在`Account`对象中（以及正确交易验证所需的其他数据）。

在节点中，所有数据都使用 Protocol Buffers 序列化存储。

Cosmos SDK 支持以下用于创建数字签名的数字密钥方案：

- `secp256k1`, 在 [SDK's `crypto/keys/secp256k1` package](https://github.com/cosmos/cosmos-sdk/blob/v0.42.1/crypto/keys/secp256k1/secp256k1.go) 包中实现.
- `secp256r1`, 在 [SDK's `crypto/keys/secp256r1` package](https://github.com/cosmos/cosmos-sdk/blob/master/crypto/keys/secp256r1/pubkey.go) 包中实现.
- `tm-ed25519`, 在 [SDK `crypto/keys/ed25519` package](https://github.com/cosmos/cosmos-sdk/blob/v0.42.1/crypto/keys/ed25519/ed25519.go) 包中实现. 该方案仅支持共识验证。

|              | 地址长度（以字节为单位） | 公钥长度(以字节为单位) | 用于交易认证 |用于共识认证 |
|:------------:|:-----------------------:|:--------------------------:|:-----------------------------------:|:-------------------------------:|
| `secp256k1`  | 20                      |                         33 | yes                                 | no                              |
| `secp256r1`  | 32                      |                         33 | yes                                 | no                              |
| `tm-ed25519` | -- not used --          |                         32 | no                                  | yes                             |

## 地址

`Addresses` 和 `PubKey`s 都是识别应用程序中参与者的公共信息。 `Account` 用于存储认证信息。 基本帐户实现由`BaseAccount`对象提供。

每个帐户都使用`address`标识，地址是从公钥派生的字节序列。 在 SDK 中，我们定义了 3 种类型的地址，用于指定使用帐户的上下文：

- `AccAddress` 标识用户（`message` 的发送者）。
- `ValAddress` 标识验证器操作符。
- `ConsAddress` 标识参与共识的验证器节点。 验证器节点是使用 **`ed25519`** 曲线导出的。

这些类型实现了 `Address` 接口：

+++ https://github.com/cosmos/cosmos-sdk/blob/v0.42.1/types/address.go#L71-L90

`Address` 构造算法定义在 [ADR-28](https://github.com/cosmos/cosmos-sdk/blob/master/docs/architecture/adr-028-public-key-addresses.md).
以下是从`pub`公钥获取帐户地址的标准方法:

```go
sdk.AccAddress(pub.Address().Bytes())
```

对于用户交互，地址使用 [Bech32](https://en.bitcoin.it/wiki/Bech32) 格式化并通过 `String` 方法实现。 Bech32 方法是与区块链交互时唯一支持使用的格式。 Bech32 人类可读部分（Bech32 前缀）用于表示地址类型。 例子：

+++ https://github.com/cosmos/cosmos-sdk/blob/v0.42.1/types/address.go#L230-L244

|                    | Address Bech32 Prefix |
| ------------------ | --------------------- |
| Accounts           | gx                |
| Validator Operator | gxvaloper         |
| Consensus Nodes    | gxvalcons         |

### Public Keys

Cosmos SDK 中的公钥由 `cryptotypes.PubKey` 接口定义。 由于公钥保存在store中，`cryptotypes.PubKey` 扩展了 `proto.Message` 接口：

+++ https://github.com/cosmos/cosmos-sdk/blob/v0.42.1/crypto/types/types.go#L8-L17

压缩格式用于`secp256k1`和`secp256r1`序列化。


公钥不用于引用帐户（或用户），通常在编写交易消息时不使用（除了少数例外：`MsgCreateValidator`、`Validator` 和 `Multisig` 消息）。
对于用户交互，`PubKey` 使用 Protobufs JSON ([ProtoMarshalJSON](https://github.com/cosmos/cosmos-sdk/blob/release/v0.42.x/codec/json.go#L12) 函数格式化 ）。 例子：

+++ https://github.com/cosmos/cosmos-sdk/blob/7568b66/crypto/keyring/output.go#L23-L39

## Keyring

`Keyring` 是一个存储和管理帐户的对象。 在 Cosmos SDK 中，Keyring 实现遵循 Keyring 接口：

+++ https://github.com/cosmos/cosmos-sdk/blob/v0.42.1/crypto/keyring/keyring.go#L51-L89
