#### Become a validator and participate in the ecological management of the main chain

- The following tutorials all rely on node execution, if not installed, please move to [Install Node](../installation.md)

#### Become a validator
* plugchaind tx staking create-validator [flags]
```
plugchaind tx staking create-validator --from mywallet \
--amount 10000onp --pubkey $(plugchaind tendermint show-validator) \
--moniker="my validator" --commission-rate="0.10" --commission-max-rate="0.20" \
--commission-max-change-rate="0.01"  --min-self-delegation="1000000" --chain-id plugchain
  
```
#### Modify validator configuration
* plugchaind tx staking edit-validator [flags]
```
plugchaind tx staking edit-validator --from mywallet \
--amount 10000onp --pubkey $(plugchaind tendermint show-validator  ) \
--moniker="my validator" --commission-rate="0.20" --commission-max-rate="0.30" \
--commission-max-change-rate="0.02" --min-self-delegation="1000000" --chain-id plugchain
  



```
**warning** : 
`commission-max-rate` and `commission-max-change-rate` cannot be modified after `create validator` is set. Please set them carefully
When `edit-validator` modifies the verifier information, the modification interval of each field is 24h, and the `min self delegation` cannot be reduced, but can only be increased

- `commission-rate`: The commission rate on block rewards and fees charged to delegators.
- `commission-max-rate`: The maximum commission rate which this validator can charge. This parameter cannot be changed after `create-validator` is processed.
- `commission-max-change-rate`: The maximum daily increase of the validator commission. This parameter cannot be changed after `create-validator` is processed.
- `min-self-delegation`: Minimum amount of plug the validator needs to have bonded at all time. If the validator's self-delegated stake falls below this limit, their entire staking pool will unbond.
- The value unit of `min-self-delegation` is `10e-6`, so the minimum value is `1000000`

#### Query validator did not withdraw the reward
* plugchaind query distribution rewards [delegator-addr] [validator-addr] [flags]
* `delegator-addr` is verifier address
* `validator-addr` is `plugchaind q staking validators` Command to find validator information according to moniker, corresponding validator-addr
```
plugchaind query distribution rewards $(plugchaind keys show mywallet -a) [validator-addr]  
```

#### Withdraw all undrawn rewards
* plugchaind tx distribution withdraw-all-rewards --from mykey [flags]
```
plugchaind tx distribution withdraw-all-rewards --from mywallet  
```

#### <span id="unbond">Cancel the pledge-lock the warehouse for 21 days </span>
* plugchaind tx staking unbond [validator-addr] [amount] --from mykey
```
plugchaind tx staking unbond onpvaloperxxxxxxxxxxxxxxxxxx 10000000onp --from mywallet --chain-id plugchain
```


#### Query the total amount of pledges of all the delegators of the validator. If you want to know the information of the delegators, please move to [Pledge of Delegators](../delegators/delegator-setup.md)
* plugchaind query staking delegations-to [validator-addr] [flags]
```
plugchaind query staking delegations-to onpvaloperxxxxxxxxxxxxxxxxxx  
```

#### Query all canceled pledge records according to the verifier
*  plugchaind query staking unbonding-delegations-from [validator-addr] [flags]
```
plugchaind query staking unbonding-delegations-from onpvaloperxxxxxxxxxxxxxxxxxxxxxxxxxxxx  
```


::: warning
In order to enter the validator set, your weight must exceed the 100th validator.
:::

-After becoming a verifier, you need to modify the configuration[validator config](../../images/node_config.png)


## common problem

### Question #1: My validator’s `voting_power: 0`

Your validator is already jailed. If the validator has not voted in more than 500 blocks out of the last 10000 blocks, or is found to be double-signed, it will be jailed.

If you are jailed due to disconnection, you can regain your voting equity to return to the validator team. First, if `plugchaind` is not running, start it again:
```bash
plugchaind start
```

Wait for your full node to catch up with the latest block height. Then, run the following command. Next, you can unjail your validator.

Finally, check with your validators to see if the voting equity is restored:

```bash
plugchaind status
```

You may notice that your voting rights are less than before. This is due to the reduction penalty for your downline!


### Question #2: My `plugchaind` crashed due to `too many open files`

The default number of files (per process) that Linux can open is 1024. It is known that `plugchaind` can open more than 1024 files. This will cause the process to crash. Quick fix Run `ulimit -n 4096` (increase the number of open files allowed) to make a quick fix, and then use `plugchaind start` to restart the process. If you use `systemd` or another process manager to start `plugchaind`, you may need to do some configuration at this level. A sample `systemd` file to solve this problem is as follows:

```toml
# /etc/systemd/system/plugchaind.service
[Unit]
Description=plugchain Node
After=network.target

[Service]
Type=simple
User=ubuntu
WorkingDirectory=/home/ubuntu
ExecStart=/home/ubuntu/go/bin/plugchaind start
Restart=on-failure
RestartSec=3
LimitNOFILE=4096

[Install]
WantedBy=multi-user.target
```
