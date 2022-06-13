---
order: 3
---

# Set client configuration

## Introduction

After the plugchain is upgraded to v1.0+, the terminal operation can use the `config` command to configure the client configuration of the current data directory, or directly modify the `$HOME/.plugchain/config/client.toml` file to meet some parameters of the operation command set up. There is no need to specify parameters for each operation command.

## usage

```bash
 plugchaind config <key> [value] [flags]
```
|  name, shorthand  | type   | required | default value       | description                                                                               |
| ----------------- | ------ | ---- | ------------ | ---------------------------------------------------------------------------------- |
| --home            | string |      | $HOME/.plugchain  | Specifies the directory to store configuration and blockchain data     



- Set the main chain ID of the current directory

```bash
 plugchaind config chain-id "plugchain_520-1" --home=<path-to-your-home>
```

- Query the main chain ID of the current directory
```bash
 plugchaind config chain-id --home=<path-to-your-home>
```

- The fields of operability are as follows:
```bash
# The network chain ID
chain-id = "plugchain_520-1"
# The keyring's backend, where the keys are stored (os|file|kwallet|pass|test|memory)
keyring-backend = "os"
# CLI output format (text|json)
output = "text"
# <host>:<port> to Tendermint RPC interface for this chain
node = "tcp://localhost:26657"
# Transaction broadcasting mode (sync|async|block)
broadcast-mode = "sync"
```

- After setting, when the operation command involves parameters such as `chain-id`, `keyring-backend`, etc., no additional specification is required. Take the data directory as `node` as an example as follows:

```bash
plugchaind tx bank send adr adr1 20000uplugcn --fees 200uplugcn --home node
```