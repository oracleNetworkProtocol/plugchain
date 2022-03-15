---
order: 9
---

# Token Standard


## introduce

Plug Chain is a Cosmos SDK based chain with full EVM support. Due to this architecture, tokens and assets in the network may come from different independent sources: bank module, token module, evm module.
`bank module, token module tokens belong to the PRC-10 protocol`, `evm module belongs to the PRC-20 protocol`.

##PLUGCN

The `plugcn` token belongs to the native token of the bank module, which can be used for pledge, IBC transfer, community governance, handling fee, etc.

## PRC-10
PRC-10 is a token built into the Plug Chain public chain. PRC-10 is a technical token standard supported by the Plug Chain blockchain itself, without the use of EVM virtual machine, in the Plug Chain network, each account can issue PRC-10 tokens through the `x/token` module. Users can lock their tokens individually. To issue tokens, the issuer needs to specify the token name, total size, precision, description, etc.

Tokens `plugcn`, `dhw1`, `kingdm`, `joey`, etc. are all PRC-10 protocol tokens


## PRC-20

PRC-20 is a set of standards for issuing assets by deploying smart contracts on the Plug Chain blockchain, and it is fully compatible with ERC-20.