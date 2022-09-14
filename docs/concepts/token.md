---
order: 4
---

# Token Standard


## introduce

Plug Chain is a Cosmos SDK based chain with full PVM support. Due to this architecture, tokens and assets in the network may come from different independent sources: bank module, token module, evm module.
`bank module, token module tokens belong to the PRC-10 protocol`, `pvm module belongs to the PRC-20 protocol`.

## PC

The `pc` token belongs to the native token of the bank module, which can be used for pledge, IBC transfer, community governance, handling fee, etc.

## PRC-10
PRC-10 is a token built into the Plug Chain public chain. PRC-10 is a technical token standard supported by the Plug Chain blockchain itself. It does not use a PVM virtual machine. In the Plug Chain network, each account can pass [`x/prc10`](../cli-client/token.md) module issues PRC-10 tokens. Users can lock their tokens individually. To issue tokens, the issuer needs to specify the token name, total size, precision, description, etc.

Tokens `pc`, `dhw1`, `kingdm`, `joey`, etc. belong to PRC-10 protocol tokens


## PRC-20

PRC-20 is a set of standards for issuing assets by deploying smart contracts on the Plug Chain blockchain through `Pvm`. It is compatible with [ERC-20](https://github.com/ethereum/EIPs/blob/master/EIPS/eip-20.md).


## PRC-721

PRC-721 is a set of standards for issuing non-fungible asset NFTs on the Plug Chain blockchain by deploying smart contracts through `Pvm`. It is consistent with [ERC-721](https://github.com/ethereum/EIPs/blob/master/EIPS/eip-721.md).