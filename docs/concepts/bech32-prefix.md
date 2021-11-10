---
order: 5
---

# Bech32 on Plug Chain Hub

Bech32 is a new Bitcoin address format proposed by Pieter Wuille and Greg Maxwell. Besides Bitcoin addresses, Bech32 can encode any short binary data. In the plugchaind network, keys and addresses may refer to a number of different roles in the network like accounts, validators etc. The plugchaind network is designed to use the Bech32 address format to provide robust integrity checks on data. The human readable part(HRP) makes it more efficient to read and the users could see error messages. More details in [bip-0173](https://github.com/bitcoin/bips/blob/master/bip-0173.mediawiki)

## Human Readable Part Table

| HRP | Definition                              |
| --- | --------------------------------------- |
| gx | Plug Chain Hub Account Address                 |
| gxpub | Plug Chain Hub Account Public Key              |
| gxvaloper | Plug Chain Hub Validator's Operator Address    |
| gxvaloperpub | Plug Chain Hub Validator's Operator Public Key |
| gxvalcons | Tendermint Consensus Address            |
| gxvalconspub | Tendermint Consensus Public Key         |

## Encoding

Not all interfaces to Plug Chain Hub users should be exposed as bech32 interfaces. Many addresses are still in hex or base64 encoded form.

To covert between other binary representation of addresses and keys, it is important to first apply the Amino encoding process before bech32 encoding.

## Account Key Example

Account Key, aka. [Application Key](validator-faq.md#application-key). Once you create a new address, you should see the following:

```bash
NAME:    TYPE:           ADDRESS:                                PUBKEY:
test1    local    gx18ekc4dswwrh2a6lfyev4tr25h5y76jkpclyxkz    gxpub1addwnpepqgxa40ww28uy9q46gg48g6ulqdzwupyjcwfumgfjpvz7krmg5mrnwk5xq9l
```

This means you have created a new address `gx18ekc4dswwrh2a6lfyev4tr25h5y76jkpclyxkz`, with the HRP `gx`. And its public key could be encoded into `gxpub1addwnpepqgxa40ww28uy9q46gg48g6ulqdzwupyjcwfumgfjpvz7krmg5mrnwk5xq9l`, with the HRP `gxpub`.

## Validator Key Example

Validator Key, aka. [Tendermint Key](validator-faq.md#tendermint-key). A Tendermint Consensus Public key is generated when the node is created with  `plugchaind init`.
You can get this value with

```bash
plugchaind tendermint show-validator --home=<plugchaind-home>
```

Example Output:

```bash
gxvalconspub1zcjduepqzuz420weqehs3mq0qny54umfk5r78yup6twtdt7mxafrprms5zqsjeuxvx
```
