---
order: 4
---

# Join the testnet

The testnet is now active, here are some important details:

:::tip
Need to [install plugchaind](install.md) first
:::
## public endpoint

- [Cosmos-GRPC:47.243.255.151:9090]()
- [Cosmos-REST:http://47.243.255.151:1317/](http://47.243.255.151:1317/)
- [Tendermint-RPC:http://47.243.255.151:26657/](http://47.243.255.151:26657/)
- [JSON-RPC:http://47.243.255.151:8545/](http://47.243.255.151:8545/)


## run a full node

#### Running the node from genesis


1. Initialize the node

```bash
plugchaind init <moniker> --chain-id=plugchain_521-1
```

2. Download the [genesis.json](https://github.com/oracleNetworkProtocol/testnet) and seed information published by the testnet:

3. Modify the configuration config.toml file, and add peers to the `private_peer_ids` assignment.

4. Start the node service
```bash
# Start the node (you can also use nohup or systemd to run in the background)

# If you modify the service port configuration, you need to add parameters where the service is used:
# For example, modify the default tendermint rpc service: tcp://localhost:26657 => tcp://localhost:5000
# When cli is used, commands with `--node` parameter need to point to this parameter as --node=tcp://localhost:5000
# Example: plugchaind q account gx1tulwx2hwz4dv8te6cflhda64dn0984harlzegw --node tcp://localhost:5000


plugchaind start
```

## Browser && Faucet

<http://test.plugchain.network/>


## Gxswap

```
Router:
bech32: gx14ewpfal7g758zjv0vu9sthcfllagnehh8s28np
eip55: 0xae5C14F7FE47a871498F670B05dF09fffa89e6F7

Fac:
bech32: gx1anezwnqmssja3phgtz6rjmx5utkry3mptxh0qw
eip55: 0xecf2274c1B8425D886E858b4396cd4E2EC324761

WPLUG:
bech32: gx1nqtlcdxd0yq5ydtqv95qguptpzx9cmpa8ga2r5
eip55: 0x9817fC34CD7901423560616804702b088C5C6C3D
```