---
title: Joining Testnets
---
## Plug Chain Testnet

This repo collects the genesis and configuration files for the various Plug Chain testnets. It exists so the Plug Chain repo does not get bogged down with large genesis files and status updates.


## Joining

`plugchain-test` is active now and here are some important details:

Faucet: 

Here are the instructions to run a validator for plugchain-test:

1. Modify the `/scripts/testnet-val-setup.sh` file according to the `chain-id` and `persistent_peers` provided in the `./latest/` directory
2. Run the latest setup script
```
cd scripts && chmod +x testnet-val-setup.sh
./testnet-val-setup.sh
```