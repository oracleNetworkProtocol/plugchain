---
order: 3
---

# Join The Mainnet

:::tip
**Requirements:** [install plugchaind](install.md) or download the corresponding version of the binary file [releases](https://github.com/oracleNetworkProtocol/plugchain/releases)
:::

## Run a Full Node

### Start node from genesis

:::tip
You must use Plug Chain [v1.1.0](https://github.com/oracleNetworkProtocol/plugchain.git) to initialize your node.
:::

1. Initialize the node

```bash
plugchaind init <moniker> --chain-id=plugchain_520-1
```

2. Download the `genesis.json`, `app.toml`, `config.toml` public on the mainnet:

```bash
curl -o ~/.plugchain/config/genesis.json https://raw.githubusercontent.com/oracleNetworkProtocol/mainnet/main/v1/genesis.json
curl -o ~/.plugchain/config/app.toml https://raw.githubusercontent.com/oracleNetworkProtocol/mainnet/main/v1/app.toml
curl -o ~/.plugchain/config/config.toml https://raw.githubusercontent.com/oracleNetworkProtocol/mainnet/main/v1/config.toml
````
3. Before starting, if you want to modify the service port, seed information, peering point, sentinel mode, etc., you can modify the file by yourself, and then start the node.


4. Start the node service
```bash
# Start the node (you can also use nohup or systemd to run in the background)


plugchaind start --minimum-gas-prices 0.0001uplugcn
```

Next, your node will perform all chain upgrade procedures. Between each upgrade, you must sync blocks with a specific version. Don't worry about using an older version at an upgrade height, the node will stop automatically.

| Proposal | Starting Height | Upgrade Height | plugchaind Version |
| -------- | ------------ | -------------- | ----- |
| [v1.0](https://www.plugchain.network/v2/communityDetail?id=7)  |  3000000     |    | [v1.1.0](https://github.com/oracleNetworkProtocol/plugchain/releases/tag/v1.1.0) |
| [v1.2.1](https://www.plugchain.network/v2/communityDetail?id=8)  |  3349542     |  3576853  | [v1.2.1](https://github.com/oracleNetworkProtocol/plugchain/releases/tag/v1.2.1) |
| [v1.5.0](https://www.plugchain.network/v2/communityDetail?id=9)  |  3935641     |  4152263  | [v1.5.0](https://github.com/oracleNetworkProtocol/plugchain/releases/tag/v1.5.0) |

:::tip
You may see some connection errors, that's okay, the P2P network is trying to find an available connection


:::

## Upgrade to Validator Node

### Create a Wallet

You can [create a new wallet](../cli-client/keys.md#create-a-new-key) or [import an existing one](../cli-client/keys.md#recover-an-existing-key-from-seed-phrase), then get some plugchaind from the exchanges or anywhere else into the wallet you just created, .e.g.

```bash
# create a new wallet
plugchaind keys add <key-name>
```

:::warning
**Important**

write the seed phrase in a safe place! It is the only way to recover your account if you ever forget your password.
:::

### Confirm your node has caught-up

```bash
# if you have not installed jq
# apt-get update && apt-get install -y jq

# if the output is false, means your node has caught-up
plugchaind status 2>&1 | jq -r '.SyncInfo.catching_up'
```

### Create Validator

Only if your node has caught-up, you can run the following command to upgrade your node to be a validator.

```bash
plugchaind tx staking create-validator \
--from mywallet \
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
**Important**

Backup the `config` directory located in your plugchaind home (default ~/.plugchain/) carefully! It is the only way to recover your validator.
:::

If there are no errors, then your node is now a validator or candidate (depending on whether your delegation amount is in the top 50)
