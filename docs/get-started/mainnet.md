---
order: 3
---

# Join mainnet

:::tip

**Requirements:** [install plugchaind](install.md) or download the corresponding version of the binary file [releases](https://github.com/oracleNetworkProtocol/plugchain/releases)

You need to [install plugchaind](install.md) first, or download the corresponding version of the binary file [releases](https://github.com/oracleNetworkProtocol/plugchain/releases)
:::

## run a full node

### Using the genesis file

:::tip
Your node must be initialized with Plug Chain [v1.1.0](https://github.com/oracleNetworkProtocol/plugchain.git)
:::

1. Initialize the node

```bash
plugchaind init <moniker> --chain-id=plugchain_520-1
```

2. Download the `genesis.json`, `app.toml`, `config.toml` public on the mainnet:

```bash
curl -o ~/.plugchain/config/genesis.json https://raw.githubusercontent.com/oracleNetworkProtocol/mainnet/main/version/v1/genesis.json
curl -o ~/.plugchain/config/app.toml https://raw.githubusercontent.com/oracleNetworkProtocol/mainnet/main/version/v1/app.toml
curl -o ~/.plugchain/config/config.toml https://raw.githubusercontent.com/oracleNetworkProtocol/mainnet/main/version/v1/config.toml
```
3. Before starting, if you want to modify the service port, seed information, peering point, sentinel mode, etc., you can modify the file yourself, and then synchronize the block.


### Sync block

Sync mainnet block data
#### One snapshot synchronization

Depending on the snapshot height, lock the `plugchaind` binary version to use


| Block height | Database | plugchaind version | Download address | Desc |
| ---- | --------- | -------- | ----| ----- |
| 8021377 | goleveldb (default） | [v1.7](https://github.com/oracleNetworkProtocol/plugchain/releases/tag/v1.7.0) | [mainnet-8021377-20230201-goleveldb.zip](https://snapshot-node-mainnet.oss-cn-hangzhou.aliyuncs.com/mainnet-8021377-20230201-goleveldb.zip) | （130.654GB）crop data |

1. Download snapshot data
```bash
wget -b -c https://snapshot-node-mainnet.oss-cn-hangzhou.aliyuncs.com/mainnet-8021377-20230201-goleveldb.zip
```

2.Unzip Data override `~/.plugchain/data/` directory

3. Use the corresponding plugchaind version to start `plugchaind start`



#### 2 Gradual upgrade synchronization
start node service

```bash
# Start the node (you can also use nohup or systemd to run in the background)

plugchaind start
```


Next, your node will perform all chain upgrade procedures. Between each upgrade, you must sync blocks with a specific version. Don't worry about using an older version at an upgrade height, the node will stop automatically.

| Proposal | Gov Height | Upgrade Height | plugchaind Version |
| -------- | ------------ | -------------- | ----- |
| [v1.0](https://www.plugchain.network/v2/communityDetail?id=7)  |  3000000     |    | [v1.1.0](https://github.com/oracleNetworkProtocol/plugchain/releases/tag/v1.1.0) |
| [v1.2.1](https://www.plugchain.network/v2/communityDetail?id=8)  |  3349542     |  3576853  | [v1.2.1](https://github.com/oracleNetworkProtocol/plugchain/releases/tag/v1.2.1) |
| [v1.5.0](https://www.plugchain.network/v2/communityDetail?id=9)  |  3935641     |  4152263  | [v1.5.0](https://github.com/oracleNetworkProtocol/plugchain/releases/tag/v1.5.0) |
| [v1.7](https://www.plugchain.network/v2/communityDetail?id=10)  |  5420512     |  5633000  | [v1.7.0](https://github.com/oracleNetworkProtocol/plugchain/releases/tag/v1.7.0) |




:::tip
You may see some connection errors, that's okay, the P2P network is trying to find an available connection


:::


## Upgrade to validator node

### Create wallet

You can [create a new wallet](../cli-client/keys.md#create a key) or [import an existing wallet](../cli-client/keys.md#recover the key with the mnemonic phrase key), then transfer some plugs from the exchange or anywhere into the wallet you just created:

```bash
# Create a new wallet
plugchaind keys add <key-name>
```

:::warning
Back up your mnemonic phrases in a safe place! If you forget your password, this is the only way to recover your account.
:::

### Confirm node synchronization status

```bash
# You can install jq with this command
# apt-get update && apt-get install -y jq

# If the output is false, your node has finished syncing
plugchaind status 2>&1 | jq -r '.SyncInfo.catching_up'
```

### Create validator

You can run the following command to promote your node to a validator only if the node has finished syncing:

```bash
plugchaind tx staking create-validator --from mywallet \
--amount 1000000uplugcn \
--pubkey $(plugchaind tendermint show-validator) \
--moniker="my validator" \
--commission-rate="0.10" \
--commission-max-rate="0.20" \
--commission-max-change-rate="0.01" \
--min-self-delegation="1000000" \
--fees 20uplugcn --chain-id plugchain_520-1
```


:::warning
**important**

Be sure to back up the `config` directory in your home (~/.plugchain/ by default) directory! This is the only way to recover validators if your server disk is damaged or if you are migrating your server.
:::

If the above command gives no errors, your node is already a validator or candidate (depending on whether your Voting Power is in the top 50)