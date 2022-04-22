#!/bin/bash

KEY="val"
CHAINID=$([ "$CHAINID" != "" ] && echo "$CHAINID" || echo "plugchain_522-1")
ID=$([ "$ID" != "" ] && echo "$ID" || echo "00" )
GENESIS=$([ "$GENESIS" == "false" ] && echo "$GENESIS" || echo "true" )
CONNECTIP=$([ "$CONNECTIP" != "" ] && echo "$CONNECTIP" || echo "" )
MONIKER="localtestnet$ID"
KEYRING="test"
LOGLEVEL="info"
NODECOMMON="node"
NODEDIR="$NODECOMMON/node$ID"
# to trace evm
#TRACE="--trace"
TRACE=""

# validate dependencies are installed
command -v jq > /dev/null 2>&1 || { echo >&2 "jq not installed. More info: https://stedolan.github.io/jq/download/"; exit 1; }

# remove existing daemon
rm -rf $NODEDIR

plugchaind config keyring-backend $KEYRING --home $NODEDIR
plugchaind config chain-id $CHAINID --home $NODEDIR

# if $KEY exists it should be deleted
plugchaind keys add $KEY --keyring-backend $KEYRING --home $NODEDIR

# Set moniker and chain-id for Evmos (Moniker can be anything, chain-id must be an integer)
plugchaind init $MONIKER --chain-id $CHAINID --home $NODEDIR
if [ $GENESIS == "true" ];then
  rm -rf "$NODECOMMON/genesis.json"
  # Change parameter token denominations to uplugcn
  cat $NODEDIR/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="uplugcn"' > $NODEDIR/config/tmp_genesis.json && mv $NODEDIR/config/tmp_genesis.json $NODEDIR/config/genesis.json
  cat $NODEDIR/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="uplugcn"' > $NODEDIR/config/tmp_genesis.json && mv $NODEDIR/config/tmp_genesis.json $NODEDIR/config/genesis.json
  cat $NODEDIR/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="uplugcn"' > $NODEDIR/config/tmp_genesis.json && mv $NODEDIR/config/tmp_genesis.json $NODEDIR/config/genesis.json
  cat $NODEDIR/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["max_deposit_period"]="180s"' > $NODEDIR/config/tmp_genesis.json && mv $NODEDIR/config/tmp_genesis.json $NODEDIR/config/genesis.json
  cat $NODEDIR/config/genesis.json | jq '.app_state["gov"]["voting_params"]["voting_period"]="300s"' > $NODEDIR/config/tmp_genesis.json && mv $NODEDIR/config/tmp_genesis.json $NODEDIR/config/genesis.json
  cat $NODEDIR/config/genesis.json | jq '.app_state["mint"]["params"]["mint_denom"]="uplugcn"' > $NODEDIR/config/tmp_genesis.json && mv $NODEDIR/config/tmp_genesis.json $NODEDIR/config/genesis.json
  cat $NODEDIR/config/genesis.json | jq '.app_state["evm"]["params"]["evm_denom"]="uplugcn"' > $NODEDIR/config/tmp_genesis.json && mv $NODEDIR/config/tmp_genesis.json $NODEDIR/config/genesis.json
  cat $NODEDIR/config/genesis.json | jq '.app_state["liquidity"]["params"]["pool_creation_fee"][0]["denom"]="uplugcn"' > $NODEDIR/config/tmp_genesis.json && mv $NODEDIR/config/tmp_genesis.json $NODEDIR/config/genesis.json

  # increase block time (?)
  cat $NODEDIR/config/genesis.json | jq '.consensus_params["block"]["time_iota_ms"]="30000"' > $NODEDIR/config/tmp_genesis.json && mv $NODEDIR/config/tmp_genesis.json $NODEDIR/config/genesis.json

  # Set gas limit in genesis
  cat $NODEDIR/config/genesis.json | jq '.consensus_params["block"]["max_gas"]="10000000"' > $NODEDIR/config/tmp_genesis.json && mv $NODEDIR/config/tmp_genesis.json $NODEDIR/config/genesis.json


  # Allocate genesis accounts (cosmos formatted addresses)
  plugchaind add-genesis-account $KEY 100000000000000000uplugcn --keyring-backend $KEYRING --home $NODEDIR

  # Sign genesis transaction
  plugchaind gentx $KEY 10000000000000uplugcn --keyring-backend $KEYRING --chain-id $CHAINID --home $NODEDIR

  # Collect genesis tx
  plugchaind collect-gentxs --home $NODEDIR

  # Run this to ensure everything worked and that the genesis file is setup correctly
  plugchaind validate-genesis --home $NODEDIR

  if [[ $1 == "pending" ]]; then
    echo "pending mode is on, please wait for the first block committed."
  fi
  cp -f "$NODEDIR/config/genesis.json" "$NODECOMMON/"
else
  sleep 5s
  while :
  do
    if [ ! -f "$NODECOMMON/genesis.json" ];then
      echo "validate genesis.json does not exist! sleep 2s";
      sleep 2s
      continue
    fi
    break
  done
  cp -f "$NODECOMMON/genesis.json" "$NODEDIR/config/"
  PERSISTENTPEERS="$(plugchaind tendermint show-node-id --home node/node0/)@${CONNECTIP}"
  sed -i "s/persistent_peers = \"\"/persistent_peers = \"${PERSISTENTPEERS}\"/g" $NODEDIR/config/config.toml
fi

sed -i "s/enable = true/enable = false/g" $NODEDIR/config/app.toml

export DAEMON_NAME=plugchaind
export DAEMON_HOME=$(pwd)/$NODEDIR
export DAEMON_RESTART_AFTER_UPGRADE=true
export DAEMON_ALLOW_DOWNLOAD_BINARIES=true
mkdir -p $DAEMON_HOME/cosmovisor/genesis/bin
cp -f $GOPATH/bin/plugchaind $DAEMON_HOME/cosmovisor/genesis/bin

# Start the node (remove the --pruning=nothing flag if historical queries are not needed)
cosmovisor start --pruning=nothing $TRACE --log_level $LOGLEVEL --minimum-gas-prices=0.0001uplugcn --home $NODEDIR
