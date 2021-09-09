APPNAME=node1
CHAINID=chaintest-1
DEMON=plug
DAEMON=plugchaind
$DAEMON init node1 --home $APPNAME --chain-id $CHAINID
# mac os 运行需要使用 sed -i "" 
sed -i "" 's#stake#plug#g' $APPNAME/config/genesis.json
sed -i "" 's#enable = false#enable = true#g' $APPNAME/config/app.toml
sed -i "" 's#swagger = false#swagger = true#g' $APPNAME/config/app.toml
$DAEMON keys add validator --home $APPNAME
$DAEMON keys add validator1 --home $APPNAME
$DAEMON add-genesis-account validator "20000000000000 $DEMON" --home $APPNAME
$DAEMON add-genesis-account validator1 "20000000000000 $DEMON" --home $APPNAME
$DAEMON gentx validator "20000000000$DEMON" --min-self-delegation "1000000" --home $APPNAME --chain-id $CHAINID
$DAEMON collect-gentxs --home $APPNAME
$DAEMON start --home $APPNAME --minimum-gas-prices 0.0001plug