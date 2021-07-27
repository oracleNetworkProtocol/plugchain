APPNAME=node/token
CHAINID=plugchain-token-1
DEMON=line
DAEMON=plugchaind
rm -rf $APPNAME
$DAEMON init node1 --home $APPNAME --chain-id $CHAINID
sed -i "" 's#stake#line#g' $APPNAME/config/genesis.json
sed -i "" 's#enable = false#enable = true#g' $APPNAME/config/app.toml
sed -i "" 's#swagger = false#swagger = true#g' $APPNAME/config/app.toml
$DAEMON keys add jcl --home $APPNAME
$DAEMON keys add jcl1 --home $APPNAME
$DAEMON keys add jcl2 --home $APPNAME
$DAEMON keys add jcl3 --home $APPNAME
$DAEMON keys add jcl4 --home $APPNAME
$DAEMON add-genesis-account jcl "20000000000000 $DEMON" --home $APPNAME
$DAEMON add-genesis-account jcl1 "20000000000000 $DEMON" --home $APPNAME
$DAEMON add-genesis-account jcl2 "20000000000000 $DEMON" --home $APPNAME
$DAEMON gentx jcl "20000000000$DEMON" --min-self-delegation "1000000" --home $APPNAME --chain-id $CHAINID
$DAEMON collect-gentxs --home $APPNAME
$DAEMON start --home $APPNAME