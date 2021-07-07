APPNAME=node/node2
CHAINID=chain-2
DEMON=bbb
DAEMON=plugchaind
rm -rf $APPNAME
$DAEMON init node1 --home $APPNAME --chain-id $CHAINID
sed -i 's#tcp://127.0.0.1:26657#tcp://0.0.0.0:26657#g' $APPNAME/config/config.toml
sed -i 's#stake#bbb#g' $APPNAME/config/genesis.json
sed -i 's/pruning = "default"/pruning = "nothing"/g' $APPNAME/config/app.toml
$DAEMON keys add jcl --home $APPNAME
$DAEMON add-genesis-account jcl "20000000000000 $DEMON" --home $APPNAME
$DAEMON gentx jcl "20000000000$DEMON" --min-self-delegation "1000000" --home $APPNAME --chain-id $CHAINID
$DAEMON collect-gentxs --home $APPNAME
$DAEMON start --home $APPNAME