APPNAME=node/node1
CHAINID=test-plugchain
DEMON=line
DAEMON=plugchaind
rm -rf $APPNAME
$DAEMON init node1 --home $APPNAME --chain-id $CHAINID
sed -i "" 's#stake#line#g' $APPNAME/config/genesis.json
$DAEMON keys add jcl --home $APPNAME
$DAEMON add-genesis-account jcl "20000000000000 $DEMON" --home $APPNAME
$DAEMON gentx jcl "20000000000$DEMON" --min-self-delegation "1000000" --home $APPNAME --chain-id $CHAINID
$DAEMON collect-gentxs --home $APPNAME
$DAEMON start --home $APPNAME