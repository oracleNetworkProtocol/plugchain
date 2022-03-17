---
order: 2
---

# Commands

## Introduction

plugchaind Daemon Commands allow you to init, start, reset a node, or generate a genesis file, etc.

You can get familiar with these commands by creating a [Local Testnet](local-testnet.md).

## Usage

```bash
plugchaind <command>
```

## Available Commands

| Name                                                             | Description                                                                                                     |
| ---------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- |
| [init](local-testnet.md#plugchaind-init)                               | Initialize private validator, p2p, genesis, and application configuration files                                 |
| [add-genesis-account](local-testnet.md#plugchaind-add-genesis-account) | Add genesis account to genesis.json                                                                             |
| [gentx](local-testnet.md#plugchaind-gentx)                             | Generate a genesis tx carrying a self delegation                                                                |
| [collect-gentxs](local-testnet.md#plugchaind-collect-gentxs)           | Collect genesis txs and output a genesis.json file                                                              |
| [start](local-testnet.md#plugchaind-start)                             | Run the full node                                                                                               |
| [unsafe-reset-all](local-testnet.md#plugchaind-unsafe-reset-all)       | Resets the blockchain database, removes address book files, and resets priv_validator.json to the genesis state |
| [tendermint](local-testnet.md#plugchaind-tendermint)                   | Tendermint subcommands                                                                                          |
| [testnet](local-testnet.md#build-and-init)                       | Initialize files for a Plug Chain Hub testnet                                                                          |
| [reset](local-testnet.md#plugchaind-reset)                             | Reset app state to the specified height                                                                         |
| [export](export.md)                                              | Export state to JSON                                                                                            |
| [config](config.md)                                              | Query and modify the client configuration of the specified data directory                                                                                            |
| version                                                          | Show executable binary version                                                                                  |

## Global Flags

| Name,shorthand | Default      | Description                                        | Required | Type   |
| -------------- | ------------ | -------------------------------------------------- | -------- | ------ |
| -h, --help     |              | Help for plugchaind                                      |          |        |
| --home         | /$HOME/.plugchain | Directory for config and data                      |          | String |
| --log_level    | \*:info      | Log level (default "main:info,state:info,*:error") |          | String |
| --trace        |              | Print out full stack trace on errors               |          |        |
