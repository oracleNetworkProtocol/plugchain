---
order: 5
---

# Bech32 前缀

Bech32是由Pieter Wuille和Greg Maxwel提出的新比特币地址格式。除了比特币之外，bech32可以编码任何短二进制数据。在PLUGChain里，私钥和地址可能指的是一些在网络中不同的角色，例如普通账户和验证人账户等。PLUGChain设计使用Bech32地址格式来提供对数据鲁棒的完整性检查。用户可读部分（human readable part）可帮助用户有效理解地址和阅读错误信息。Bech32更多细节见 [bip-0173](https://github.com/bitcoin/bips/blob/master/bip-0173.mediawiki)。

## 用户可读部分表

| HRP | Definition                              |
| --- | :-------------------------------------- |
| gx | PLUGChain Hub 账户地址前缀                 |
| gxpub | PLUGChain Hub 账户地址公钥前缀              |
| gxvaloper | PLUGChain Hub 验证者操作员地址前缀    |
| gxvaloperpub | PLUGChain Hub 验证者操作员地址公钥前缀 |
| gxvalcons | Tendermint 共识地址前缀            |
| gxvalconspub | Tendermint 共识地址公钥前缀          |

## 编码

不是所有PLUGChain的用户地址都会以bech32的格式暴露出来。许多地址仍然是hex编码或者base64编码。 在bech32编码之前，首先要使用amino对其他地址私钥二进制表示进行编码。

## 账户密钥示例

一旦创建一个新的账户，你将会看到以下信息：

```bash
NAME:    TYPE:           ADDRESS:                                PUBKEY:
test1    local    gx18ekc4dswwrh2a6lfyev4tr25h5y76jkpclyxkz    gxpub1addwnpepqgxa40ww28uy9q46gg48g6ulqdzwupyjcwfumgfjpvz7krmg5mrnwk5xq9l
```

这意味着你创建了一个新账户地址 `gx18ekc4dswwrh2a6lfyev4tr25h5y76jkpclyxkz`， 它的用户可读部分是 `gx`。他的公钥被编码成  `gxpub1addwnpepqgxa40ww28uy9q46gg48g6ulqdzwupyjcwfumgfjpvz7krmg5mrnwk5xq9l`， 它的用户可读部分是 `gxpub`。

## 验证人密钥示例

在执行 `plugchaind init`命令时回自动产生一个Tendermint的共识密钥给该节点。你可以通过以下命令查询：

```bash
plugchaind tendermint show-validator --home=<plugchaind-home>
```

示例输出：

```bash
gxvalconspub1zcjduepqzuz420weqehs3mq0qny54umfk5r78yup6twtdt7mxafrprms5zqsjeuxvx
```
