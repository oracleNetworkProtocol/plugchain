---
order: 1
---

# Chain ID

Learn about Plug Chain ID format

## Mainnet chain ID


| Name | ChainID | Identifier | EIP155 Number | Version Number |
|---------------------|--------------------|-------------| ---|-----|
| Plug Chain Mainnet | `plugchain_520-1` | `plugchain` | `520` | `1` |
| Plug Chain Testnet | `plugchain_521-1` | `plugchain` | `521` | `1` |


## Chain identifier

Each chain must have a unique identifier or "chain-id". Tendermint requires every application
Define your own `chain-id` in [genesis.json fields](https://docs.tendermint.com/master/spec/core/genesis.html#genesis-fields). However, in order to comply with both EIP155 and Cosmos' chain upgrade standards, Plug Chain-compatible chains must implement a special structure for their chain identifiers.

## Structure

Plug Chain chain ID consists of 3 main components

- **Identifier**: An unstructured string that defines the name of the application.
- **EIP155 ID**: Immutable [EIP155](https://github.com/ethereum/EIPs/blob/master/EIPS/eip-155.md) `CHAIN_ID` defines the replay attack protection ID.
- **version number**: is the version number the chain is currently running on (always positive).
This number must be incremented every time the chain is upgraded or forked to avoid network or consensus errors.

### Format

The format of specifying a chain ID compatible with Plug Chain in genesis is as follows:

``` Bash
{identifier}_{EIP155}-{version}
```

The following table provides an example, where the second row corresponds to the upgrade of the first row:

|Chain ID |Identifier |EIP155 Number |Version Number |
|----------------|------------|---------------|---------|
| `plugchain_520-1` | plugchain | 520 | 1 |
| `plugchain_520-2` | plugchain | 520 | 2 |
| `...` | ... | ... | ... |
| `plugchain_520-N` | plugchain | 520 | N |