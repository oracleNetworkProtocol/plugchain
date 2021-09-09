---
order: 4
---

# Join The Testnet

testnet is activity
:::tip
[isntall plug](install.md)
:::
## Public Endpoints

- [GRPC:8.210.180.240:9090]()
- [RPC:http://8.210.180.240:26657/](http://8.210.180.240:26657/)
- [REST:http://8.210.180.240:1317/](http://8.210.180.240:1317/)



## Run a Full Node

### Start node from genesis
:::tip 
You must use PLUGChain[v0.5.0](https://github.com/oracleNetworkProtocol/plugchain.git) to initialize your node
:::

1. Initialize the node

```bash
plugchaind init <moniker> --chain-id=plugchain-testnet-1
```

2. Download genesis.json and seed information disclosed on the testnet or enter the cloned plugchain directory:
*[ Testnet genesis file and seed information](https://github.com/oracleNetworkProtocol/plugchain/tree/main/testnet/latest)*, the following is achieved by moving the genesis.json in the clone warehouse.

```bash
mv ./testnet/latest/genesis.json ~/.plugchain/
```

:::warning
The following third step can be skipped, if skipping, you need to add the parameter `--p2p.seeds` parameter to the fourth step)
:::

3. After overwriting genesis.json in the data directory, modify the seeds parameter in ~/.plugchain/config/config.toml and add seed information

Modify the seeds provided in the ./testnet/latest/seeds.txt file and modify the `seeds` parameter in ~/.plugchain/config/config.toml to set the seed nodes of the link. The seed information is separated by English commas


4. Start the node service
```bash
# Start the node (you can also use nohup or systemd to run in the background)

# The third step does not modify the seed information. When running start, add the parameter 
# --p2p.seeds="5f81625b69d192d3ef5bf47b83484326e0546491@47.100.161.102:26656"

# If you modify the service port configuration, you need to add parameters where the service is used:
# For example, modify the default tendermint rpc service: tcp://localhost:26657 => tcp://localhost:5000
# When using cli, commands with the `--node` parameter need to point to this parameter as --node=tcp://localhost:5000
# For example: plugchaind q account gx1tulwx2hwz4dv8te6cflhda64dn0984harlzegw --node tcp://localhost:5000


plugchaind start --minimum-gas-prices 0.025line
```


## Faucet

welcome to get coin form [test explorers](http://test.plugchain.network/wallet/receive)

## Explorers

<http://test.plugchain.network/>

