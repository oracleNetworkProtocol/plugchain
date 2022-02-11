---
order: 7
---

# Plug Chain Hub latest version upgrade

This document describes the steps for validators and full node operators to successfully execute the *upgrade plan*. The official is not responsible for the loss of assets due to the upgrade. Please back up your assets when upgrading.
Plug Chain will stop running the main chain at block height `3000000` and upgrade to the official version `v1.0`.


The upgrade contents are as follows:
1. Upgrade the mainnet and change the chain-id to `plugchain_520-1`
2. Modify the coin name `plug` on the chain to: (the data of the relevant coins on the chain are data with a precision of 0, and the coin name with a precision of 6 is used for wallet, browser, and external APP application display)
   - precision 0: `uplugcn`
   - Precision 6: `plugcn`
3. Modify the community proposal module (gov) `fundraising` and `voting period`, the overall revision is 14 days
4. Adjust the maximum number of validators to `50`
5. Access the EVM module
6. Enable the function of destroying `plugcn`
7. Overall adjustment of `x/liquidity`, `x/token` and other module fee parameters
8. The wallet supports two key signature algorithms `secp256k1`, `eth_secp256k1`

Precautions:
1. No transaction records are kept
2. The current block height is retained, and the original block height data is not retained
3. Retain proposal information
4. Keep the liquidity project data and keep the token data




# backup

*First you need to stop the node running*

1. Validator Node
 - Backup entire data directory, default `.plugchain/`
 - Backup `.plugchain/data/priv_validator_state.json` file, the data in `.plugchain/data/` can be deleted.
 
2. Other functional nodes
 - Backup wallet directory


# Steps 

1. Get the `v1.0.0` binary file plugchaind

```bash
# Pull the v1.0.0 version code (you can use `git tag` to view the next tag version locally, if there is `v1.0.0`, skip this step)
git fetch origin v1.0.0

# Execute the compiled binary
make install

````

2. Create a new data directory, you can use --home to specify the data directory

```bash
plugchaind init myNode --chain-id plugchain_520-1
````

3. Download the `genesis.json`, `app.toml`, `config.toml` public on the mainnet:


```bash
curl -o ~/.plugchain/config/genesis.json https://raw.githubusercontent.com/oracleNetworkProtocol/plugchain/main/mainnet/v1/genesis.json
curl -o ~/.plugchain/config/app.toml https://raw.githubusercontent.com/oracleNetworkProtocol/plugchain/main/mainnet/v1/app.toml
curl -o ~/.plugchain/config/config.toml https://raw.githubusercontent.com/oracleNetworkProtocol/plugchain/main/mainnet/v1/config.toml
````

*The fourth step only requires the operation of the validator node, and the rest of the nodes are skipped*

4. The validator node needs to overwrite the `node_key.json` and `priv_validator_key.json` files in the config directory of the original node into the config of the new data directory to identify the validator, and other nodes do not need it.

5. * Modify the current configuration according to the original node configuration to meet your own needs. If you need to operate the wallet, you need to import the backup wallet directory into the new data directory before you can operate. *


6. Start the node

```bash
plugchaind start
````