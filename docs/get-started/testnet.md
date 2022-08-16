---
order: 4
---

# Join the testnet

The testnet is now active, here are some important details:

:::tip
Need to [install plugchaind](install.md) first
:::
## public endpoint

- [Cosmos-GRPC:47.108.75.227:9090]()
- [Cosmos-REST:http://47.108.75.227:1317/](http://47.108.75.227:1317/)
- [Tendermint-RPC:http://47.108.75.227:26657/](http://47.108.75.227:26657/)
- [JSON-RPC:http://47.108.75.227:8545/](http://47.108.75.227:8545/)


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