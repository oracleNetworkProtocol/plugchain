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
You must use Plug Chain [v0.5.0](https://github.com/oracleNetworkProtocol/plugchain.git) to initialize your node.
:::

1. Initialize the node

```bash
plugchaind init <moniker> --chain-id=plugchain
```

2. Download genesis.json and seed information disclosed on the mainnet or enter the cloned plugchain directory:
*[Mainnet genesis file and seed information](https://github.com/oracleNetworkProtocol/plugchain/blob/main/mainnet/v1)*, the following is achieved by moving the genesis.json in the clone warehouse.

```bash
mv ./mainnet/v1/genesis.json ~/.plugchain/
```

:::warning
The following third step can be skipped, if skipping, you need to add the parameter `--p2p.seeds` parameter to the fourth step)
:::

3. After overwriting genesis.json in the data directory, modify the seeds parameter in ~/.plugchain/config/config.toml and add seed information

Modify the seeds provided in the ./mainnet/v1/seeds.txt file and modify the `seeds` parameter in ~/.plugchain/config/config.toml to set the seed nodes of the link. The seed information is separated by English commas


4. Start the node service
```bash
# Start the node (you can also use nohup or systemd to run in the background)

# The third step does not modify the seed information. When running startadd the parameter
# --p2p.seeds="7488f044132cec94e72c0eb5cdd267fb5607f5d1@47.102.107.120:26656,60fde7a070938367ede8943ee45bee622424753a@47.102.126.234:26656"

# If you modify the service port configuration, you need to add parameters where the service is used:
# For example, modify the default tendermint rpc service: tcp://localhost:26657 => tcp://localhost:5000
# When using cli, commands with the `--node` parameter need to point to this parameter as --node=tcp://localhost:5000
# For example: plugchaind q account gx1tulwx2hwz4dv8te6cflhda64dn0984harlzegw --node tcp://localhost:5000

plugchaind start --minimum-gas-prices 0.0001plug
```

When the block reaches the upgrade height (762880)
```
ERR UPGRADE "x/token" NEEDED at height: 762880:
2:04AM ERR CONSENSUS FAILURE!!! err="UPGRADE \"x/token\" NEEDED at height: 762880: "
```
You can [Upgrade Info](./upgrade-process.md)

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
--amount 1000000plug \
--pubkey $(plugchaind tendermint show-validator) \
--moniker="my validator" \
--commission-rate="0.10" \
--commission-max-rate="0.20" \
--commission-max-change-rate="0.01" \
--min-self-delegation="1000000" \
--fees 20plug --chain-id plugchain
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