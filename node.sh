# !/bin/bash
make install
rm -rf home1
NODENAME1="home1"

CHAIN_ID=plugchain

DENOM=plug
plugchaind init $NODENAME1 --chain-id $CHAIN_ID --home $NODENAME1

plugchaind keys add mywallet  --home $NODENAME1 
plugchaind keys add mywallet1  --home $NODENAME1 

sed -i '' 's#tcp://127.0.0.1:26657#tcp://0.0.0.0:26657#g' "$(pwd)/"$NODENAME1/config/config.toml
sed -i '' "s/\"stake\"/\"$DENOM\"/g" "$(pwd)/"$NODENAME1/config/genesis.json
sed -i '' 's/pruning = "syncable"/pruning = "nothing"/g' "$(pwd)/"$NODENAME1/config/app.toml

plugchaind add-genesis-account $(plugchaind keys show mywallet -a --home $NODENAME1)  "15869000000000000 plug" --home $NODENAME1 &&\
plugchaind add-genesis-account $(plugchaind keys show mywallet1 -a --home $NODENAME1) "80000000000000 plug" --home $NODENAME1 &&\
plugchaind add-genesis-account $(plugchaind keys show mywallet2 -a --home $NODENAME1) "20000000000000 plug" --home $NODENAME1 &&\
plugchaind add-genesis-account $(plugchaind keys show mywallet3 -a --home $NODENAME1) "20000000000000 plug" --home $NODENAME1 &&\

plugchaind gentx mywallet 80000000000000plug --keyring-dir $NODENAME1 --home $NODENAME1 --chain-id $CHAIN_ID &&\
plugchaind gentx mywallet 80000000000000plug --keyring-dir $NODENAME1 --home $NODENAME1 --chain-id $CHAIN_ID &&\
plugchaind gentx mywallet 20000000000000plug --keyring-dir $NODENAME1 --home $NODENAME1 --chain-id $CHAIN_ID &&\
plugchaind gentx mywallet 20000000000000plug --keyring-dir $NODENAME1 --home $NODENAME1 --chain-id $CHAIN_ID &&\

plugchaind collect-gentxs --home $NODENAME1

# echo $(jq . $(pwd)/node1/config/genesis.json | jq -r '.app_state.genutil.gen_txs[0].body.memo') 
# echo $(jq . $(pwd)/node2/config/genesis.json | jq -r '.app_state.genutil.gen_txs[0].body.memo') 
# echo $(jq . $(pwd)/node3/config/genesis.json | jq -r '.app_state.genutil.gen_txs[0].body.memo') 
# echo $(jq . $(pwd)/$NODENAME1/config/gentx/*.json)
# echo $(jq . $(pwd)/$NODENAME2/config/gentx/*.json)
# echo $(jq . $(pwd)/$NODENAME3/config/gentx/*.json)