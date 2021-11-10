#!/bin/bash

APPNAME=node/node1
CHAINID=chaintest-1
DEMON=plug
DAEMON=plugchaind
a=`uname  -a`
b="Darwin"
if [[ $a =~ $b ]]
then
SedP="''"
else
SedP=""
fi
rm -rf $APPNAME
$DAEMON init node1 --home $APPNAME --chain-id $CHAINID
sed -i $SedP 's#stake#plug#g' $APPNAME/config/genesis.json
sed -i $SedP 's#enable = false#enable = true#g' $APPNAME/config/app.toml
sed -i $SedP 's#swagger = false#swagger = true#g' $APPNAME/config/app.toml
sed -i $SedP 's/172800s/600s/g' $APPNAME/config/genesis.json
$DAEMON keys add validator --home $APPNAME --keyring-backend=test
$DAEMON keys add validator1 --home $APPNAME --keyring-backend=test
$DAEMON add-genesis-account validator "20000000000000 $DEMON" --home $APPNAME --keyring-backend=test
$DAEMON add-genesis-account validator1 "20000000000000 $DEMON" --home $APPNAME --keyring-backend=test
$DAEMON gentx validator "20000000000$DEMON" --min-self-delegation "1000000" --home $APPNAME --chain-id $CHAINID --keyring-backend=test
$DAEMON collect-gentxs --home $APPNAME
$DAEMON start --home $APPNAME --minimum-gas-prices 0.0001plug