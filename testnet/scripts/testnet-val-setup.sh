#!/bin/bash
. ./check-go.sh

echo "-- Clear old plugchain testnet data and install plugchain and setup the node --"

rm -rf ~/.plugchain

YOUR_KEY_NAME=$1
YOUR_NAME=$2
DAEMON=plugchaind
DENOM=line
CHAIN_ID=plugchain-testnet-1
SEEDS=""
APPNAME="~/.plugchain"

echo "install plugchain"
git clone https://github.com/oracleNetworkProtocol/plugchain $GOPATH/src/github.com/oracleNetworkProtocol/plugchain
cd $GOPATH/src/github.com/oracleNetworkProtocol/plugchain
git fetch
git checkout v0.2.0
make install

echo "Creating keys"
$DAEMON keys add $YOUR_KEY_NAME

echo ""
echo "After you have copied the mnemonic phrase in a safe place,"
echo "press the space bar to continue."
read -s -d ' '
echo ""

echo "Setting up your validator"
$DAEMON init --chain-id $CHAIN_ID $YOUR_NAME
 cp -f $GOPATH/src/github.com/oracleNetworkProtocol/plugchain/testnet/latest/genesis.json $APPNAME/config/

echo "----------Setting config for seed node---------"
sed -i 's#tcp://127.0.0.1:26657#tcp://0.0.0.0:26657#g' $APPNAME/config/config.toml
sed -i '/seeds =/c\seeds = "'"$SEEDS"'"' $APPNAME/config/config.toml

DAEMON_PATH=$(which $DAEMON)

# echo "Installing cosmovisor - an upgrade manager..."

# rm -rf $GOPATH/src/github.com/cosmos/cosmos-sdk
# git clone https://github.com/cosmos/cosmos-sdk $GOPATH/src/github.com/cosmos/cosmos-sdk
# cd $GOPATH/src/github.com/cosmos/cosmos-sdk
# git checkout v0.40.0
# cd cosmovisor
# make cosmovisor
# cp cosmovisor $GOBIN/cosmovisor

# echo "Setting up cosmovisor directories"
# mkdir -p $APPNAME/cosmovisor
# mkdir -p $APPNAME/cosmovisor/genesis/bin
# cp $GOBIN/plugchaind $APPNAME/cosmovisor/genesis/bin

# echo "---------Creating system file---------"

# echo "[Unit]
# Description=Cosmovisor daemon
# After=network-online.target
# [Service]
# Environment="DAEMON_NAME=plugchaind"
# Environment="DAEMON_HOME=${HOME}/."
# Environment="DAEMON_RESTART_AFTER_UPGRADE=on"
# User=${USER}
# ExecStart=${GOBIN}/cosmovisor start
# Restart=always
# RestartSec=3
# LimitNOFILE=4096
# [Install]
# WantedBy=multi-user.target
# " >cosmovisor.service

# sudo mv cosmovisor.service /lib/systemd/system/cosmovisor.service
# sudo -S systemctl daemon-reload
# sudo -S systemctl start cosmovisor

echo
echo "Your account address is :"
$DAEMON keys show $YOUR_KEY_NAME -a
echo "Your node setup is done. You would need some tokens to start your validator. You can get some tokens from the faucet: http://www.plugchain.network/wallet/receive"
echo
echo
# echo "After receiving tokens, you can create your validator by running"
# echo "$DAEMON tx staking create-validator --amount 9000000000$DENOM --commission-max-change-rate \"0.1\" --commission-max-rate \"0.20\" --commission-rate \"0.1\" --details \"Some details about yourvalidator\" --from $YOUR_KEY_NAME   --pubkey=\"$($DAEMON tendermint show-validator)\" --moniker $YOUR_NAME --min-self-delegation \"1000000\" --fees 5000$DENOM --chain-id $CHAIN_ID"