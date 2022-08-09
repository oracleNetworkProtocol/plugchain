# rm -rf $HOME/.relayer
rm -rf ./relayer
rly config init --home relayer
rly chains add --file  ibc/chains/evm_100-1.json --home relayer
# rly chains add --file  ibc/chains/plugchain_520-1.json --home relayer 
rly chains add --file  ibc/chains/jntm_522-1.json --home relayer 

ibcplug="bottom dad own clap trade dice action ripple crater obvious mountain cement penalty first doctor advice keen reform coyote fun leader verb december dry"
ibceth="ticket used error wait wave wait top what fix garlic jewel body frozen tone reject tomorrow voyage album aware shop safe trumpet slush skin"
ibcsyw="area hollow woman enemy work shy reopen sadness opera adjust ensure pupil valley market shoot resource raccoon provide arena old music interest huge library"
# rly keys restore plugchain_520-1 ibcplug "$ibcplug" --home relayer
rly keys restore evm_100-1 ibcevm "$ibceth" --home relayer --coin-type 60
rly keys restore jntm_522-1 val1 "$ibcsyw" --home relayer

# 只有新创建两个链关联需要执行

##创建两个新paths
# rly paths new evm_100-1 plugchain_520-1 evm-plugchain --home relayer
# rly paths new plugchain_520-1 evm_100-1 plugchain-evm --home relayer 
rly paths new evm_100-1 jntm_522-1 aa --home relayer 
## 新创建 channel, client, and connection，
# rly transact link evm-plugchain --home relayer
rly transact link aa --home relayer
###2022-08-04T06:08:09.928113Z     info    Successful transaction  {"provider_type": "cosmos", "chain_id": "plugchain_520-1", "gas_used": 80533, "fees": "107uplugcn", "fee_payer": "gx1kmu6k4s0uqyvpprewq5klvn0urwmesyr9jygdn", "height": 5236049, "msg_types": ["/ibc.core.client.v1.MsgCreateClient"], "tx_hash": "AB6DFD70B5A9C06ECC57580EA3CE711FD935EB0879E231B8B0D74199D6B8A7E9"}
###2022-08-04T06:08:09.928161Z     info    Client Created  {"src_chain_id": "plugchain_520-1", "src_client_id": "07-tendermint-3", "dst_chain_id": "evm_100-1"}

# rly start ibc-plug --home relayer