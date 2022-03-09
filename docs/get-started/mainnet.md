---
order: 3
---

# Join The Mainnet

:::tip
**Requirements:** [install plugchaind](install.md)
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
curl -o ~/.plugchain/config/genesis.json https://raw.githubusercontent.com/oracleNetworkProtocol/plugchain/main/mainnet/v1/genesis.json
curl -o ~/.plugchain/config/app.toml https://raw.githubusercontent.com/oracleNetworkProtocol/plugchain/main/mainnet/v1/app.toml
curl -o ~/.plugchain/config/config.toml https://raw.githubusercontent.com/oracleNetworkProtocol/plugchain/main/mainnet/v1/config.toml
````
3. Before starting, if you want to modify the service port, seed information, peering point, sentinel mode, etc., you can modify the file by yourself, and then start the node.


4. Start the node service
```bash
# Start the node (you can also use nohup or systemd to run in the background)


plugchaind start --minimum-gas-prices 0.0001uplugcn
```

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

If there are no errors, then your node is now a validator or candidate (depending on whether your delegation amount is in the top 100)

Read more:

- Concepts
  - [General Concepts](../concepts/general-concepts.md)
  - [Validator FAQ](../concepts/validator-faq.md)
- Validator Security
  - [Sentry Nodes (DDOS Protection)](../concepts/sentry-nodes.md)
  - [Key Management](../tools/kms.md)