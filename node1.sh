# !/bin/bash
NODENAME1="node/node1"
NODENAME2="node/node2"
NODENAME3="node/node3"
CHAIN_ID=onpchain
plugchaind init $NODENAME1 --chain-id $CHAIN_ID --home $NODENAME1
plugchaind init $NODENAME2 --chain-id $CHAIN_ID --home $NODENAME2
plugchaind init $NODENAME3 --chain-id $CHAIN_ID --home $NODENAME3

plugchaind keys add mywallet1 --home $NODENAME1 &&\
plugchaind keys add mywallet2 --home $NODENAME1 &&\
plugchaind keys add mywallet3 --home $NODENAME1 && \
plugchaind keys add mywallet4 --home $NODENAME1

sed -i '' 's/stake/onp/g' "$(pwd)/"$NODENAME1/config/genesis.json
sed -i '' 's/stake/onp/g' "$(pwd)/"$NODENAME2/config/genesis.json
sed -i '' 's/stake/onp/g' "$(pwd)/"$NODENAME3/config/genesis.json
sed -i '' 's/172800s/1000s/g' "$(pwd)/"$NODENAME1/config/genesis.json

plugchaind add-genesis-account $(plugchaind keys show mywallet1 -a --home $NODENAME1)  "10000000000 onp" --home $NODENAME1 &&\
plugchaind add-genesis-account $(plugchaind keys show mywallet2 -a --home $NODENAME1)  "10000000000 onp" --home $NODENAME1 &&\
plugchaind add-genesis-account $(plugchaind keys show mywallet3 -a --home $NODENAME1)  "10000000000 onp" --home $NODENAME1 

plugchaind add-genesis-account $(plugchaind keys show mywallet2 -a --home $NODENAME1) "10000000000 onp" --home $NODENAME2 &&\
plugchaind add-genesis-account $(plugchaind keys show mywallet3 -a --home $NODENAME1) "10000000000 onp" --home $NODENAME3

plugchaind gentx mywallet1 1000000000onp --keyring-dir $NODENAME1 --home $NODENAME1 --chain-id $CHAIN_ID &&\
plugchaind gentx mywallet2 1000000000onp --keyring-dir $NODENAME1 --home $NODENAME2 --chain-id $CHAIN_ID  &&\
plugchaind gentx mywallet3 1000000000onp --keyring-dir $NODENAME1 --home $NODENAME3 --chain-id $CHAIN_ID

plugchaind collect-gentxs --home $NODENAME1
plugchaind collect-gentxs --home $NODENAME2
plugchaind collect-gentxs --home $NODENAME3

echo $(jq . $(pwd)/$NODENAME1/config/genesis.json | jq -r '.app_state.genutil.gen_txs[0].body.memo')
echo $(jq . $(pwd)/$NODENAME2/config/genesis.json | jq -r '.app_state.genutil.gen_txs[0].body.memo')
echo $(jq . $(pwd)/$NODENAME3/config/genesis.json | jq -r '.app_state.genutil.gen_txs[0].body.memo')