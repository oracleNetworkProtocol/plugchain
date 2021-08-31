---
order: 3
---

# Local Testnet

For testing or developing purpose, you may want to setup a local testnet.

## Single Node Testnet

**Requirements:**

- [Install plug](../get-started/install.md)

:::tip
We use the default [home directory](intro.md#home-directory) for all the following examples
:::

### plugchaind init

Initialize the genesis.json file that will help you to bootstrap the network

```bash
plugchaind init testing --chain-id=testing
```

### create a key

Create a key to hold your validator account

```bash
plugchaind keys add MyValidator
```

### plugchaind add-genesis-account

Add that key into the genesis.app_state.accounts array in the genesis file

:::tip
this command lets you set the number of coins. Make sure this account has some plugchaind which is the only staking coin on PLUGChain Hub
:::

```bash
plugchaind add-genesis-account $(plugchaind keys show MyValidator --address) 150000000plug
```

### plugchaind gentx

Generate the transaction that creates your validator. The gentxs are stored in `~/.plugchain/config/gentx/`

```bash
plugchaind gentx MyValidator 100000000plugchaind --chain-id=testing 
```

### plugchaind collect-gentxs

Add the generated staking transactions to the genesis file

```bash
plugchaind collect-gentxs
```

### plugchaind start

Change the default token denom to `plug`

```bash
sed -i 's/stake/plug/g' $HOME/.plugchain/config/genesis.json
```

Now it‘s ready to start `plugchaind`

```bash
plugchaind start
```

### plugchaind unsafe-reset-all

You can use this command to reset your node, including the local blockchain database, address book file, and resets priv_validator.json to the genesis state.

This is useful when your local blockchain database somehow breaks and you are not able to sync or participate in the consensus.

```bash
plugchaind unsafe-reset-all
```

### plugchaind tendermint

Query the unique node id which can be used in p2p connection, e.g. the `seeds` and `persistent_peers` in the [config.toml](intro.md#cnofig-toml) are formatted as `<node-id>@ip:26656`.

The node id is stored in the [node_key.json](intro.md#node_key-json).

```bash
plugchaind tendermint show-node-id
```

Query the [Tendermint Pubkey](../concepts/validator-faq.md#tendermint-key) which is used to [identify your validator](../cli-client/stake/create-validator.md), and the corresponding private key will be used to sign the Pre-vote/Pre-commit in the consensus.

The [Tendermint Key](../concepts/validator-faq.md#tendermint-key) is stored in the [priv_validator.json](intro.md#priv_validator-json) which is [required to be backed up](../concepts/validator-faq.md#how-to-backup-the-validator) once you become a validator.

```bash
plugchaind tendermint show-validator
```

Query the bech32 prefixed validator address

```bash
plugchaind tendermint show-address
```

### plugchaind export

Please refer to [Export Blockchain State](export.md)

<!-- ## Multiple Nodes Testnet

**Requirements:**

- [Install plug](../get-started/install.md)
- [Install jq](https://stedolan.github.io/jq/download/)
- [Install docker](https://docs.docker.com/engine/installation/)
- [Install docker-compose](https://docs.docker.com/compose/install/)

### Build and Init

```bash
# Work from the PLUGChain Hub repo
cd [your-PLUGChain Hub-repo]

# Build the linux binary in ./build
make build-linux

# Quick init a 4-node testnet configs
make testnet-init
```

The `make testnet-init` generates config files for a 4-node testnet in the `./build/nodecluster` directory by calling the `plugchaind testnet` command:

```bash
$ tree -L 3 build/nodecluster/
build/nodecluster/
├── gentxs
│   ├── node0.json
│   ├── node1.json
│   ├── node2.json
│   └── node3.json
├── node0
│   ├── plug
│   │   ├── config
│   │   └── data
│   └── plugcli
│       ├── key_seed.json
│       └── keys
├── node1
│   ├── plug
│   │   ├── config
│   │   └── data
│   └── plugcli
│       └── key_seed.json
├── node2
│   ├── plug
│   │   ├── config
│   │   └── data
│   └── plugcli
│       └── key_seed.json
└── node3
    ├── plug
    │   ├── config
    │   └── data
    └── plugcli
        └── key_seed.json
```

### Start

```bash
make testnet-start
```

This command creates a 4-node network using the ubuntu:16.04 docker image. The ports for each node are found in this table:

| Node      | P2P Port | RPC Port |
| --------- | -------- | -------- |
| plugnode0 | 26656    | 26657    |
| plugnode1 | 26659    | 26660    |
| plugnode2 | 26661    | 26662    |
| plugnode3 | 26663    | 26664    |

To update the binary, just rebuild it and restart the nodes:

```bash
make build-linux testnet-start
```

### Stop

To stop all the running nodes:

```bash
make testnet-stop
```

### Clean

To stop all the running nodes and delete all the files in the `build/` directory:

```bash
make testnet-clean
``` -->
