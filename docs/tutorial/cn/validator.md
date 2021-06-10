#### 成为验证者参与主链生态管理

- 以下教程都依赖节点执行，如未安装，请移步到 [安装节点](./installation.md)

#### 添加验证者
* plugchaind tx staking create-validator [flags]
```
plugchaind tx staking create-validator --from mywallet \
--amount 10000onp --pubkey $(plugchaind tendermint show-validator) \
--moniker="my validator" --commission-rate="0.10" --commission-max-rate="0.20" \
--commission-max-change-rate="0.01"  --min-self-delegation="1000000" --chain-id plugchain
  
```
#### 修改已经存在的验证者配置
* plugchaind tx staking edit-validator [flags]
```
plugchaind tx staking edit-validator --from mywallet \
--amount 10000onp --pubkey $(plugchaind tendermint show-validator  ) \
--moniker="my validator" --commission-rate="0.20" --commission-max-rate="0.30" \
--commission-max-change-rate="0.02" --min-self-delegation="1000000" --chain-id plugchain
  
```
**注意** : 

`commission-rate`的值必须符合如下的不变量检查：

+ 必须在 0 和 验证人的`commission-max-rate` 之间
+ 不得超过 验证人的`commission-max-change-rate`, 该参数标识**每日**最大的百分点变化数。也就是，一个验证人在`commission-max-change-rate`的界限内每日一次可调整的最大佣金变化。

`min-self-delegation`的值单位是`10e-6`,所以此值最小为 `1000000`

#### 查询验证者未提取收益
* plugchaind query distribution rewards [delegator-addr] [validator-addr] [flags]
* delegator-addr 是验证者地址
* validator-addr 为 `plugchaind q staking validators  ` 命令根据moniker找到验证者信息，对应的validator-addr
```
plugchaind query distribution rewards $(plugchaind keys show mywallet -a) [validator-addr]  
```

#### 提取所有未提取奖励
* plugchaind tx distribution withdraw-all-rewards --from mykey [flags]
```
plugchaind tx distribution withdraw-all-rewards --from mywallet  
```

#### <span id="unbond">取消质押-锁仓21天 </span>
* plugchaind tx staking unbond [validator-addr] [amount] --from mykey
```
plugchaind tx staking unbond onpvaloperxxxxxxxxxxxxxxxxxx 10000000onp --from mywallet --chain-id plugchain
```


#### 查询验证者的所有委托人的质押总量,想了解委托人信息,请移步[委托人质押](delegator-setup.md)
* plugchaind query staking delegations-to [validator-addr] [flags]
```
plugchaind query staking delegations-to onpvaloperxxxxxxxxxxxxxxxxxx  
```

#### 根据验证者查询所有取消质押的记录
*  plugchaind query staking unbonding-delegations-from [validator-addr] [flags]
```
plugchaind query staking unbonding-delegations-from onpvaloperxxxxxxxxxxxxxxxxxxxxxxxxxxxx  
```


::: warning 注意
为了能进入验证人集合，你的权重必须超过第100名的验证人。
:::

- 成为验证者之后，需要修改配置[validator config](../../images/node_config.png)


## 常见问题

### 问题 #1 : 我的验证人的`voting_power: 0`

你的验证人已经是jailed状态。如果验证人在最近`10000`个区块中有超过`500`个区块没有进行投票，或者被发现双签，就会被jail掉。

如果被因为掉线而遭到jail，你可以重获你的投票股权以重回验证人队伍。首先，如果`plugchaind`没有运行，请再次启动：

```bash
plugchaind start
```

等待你的全节点追赶上最新的区块高度。然后，运行如下命令。接着，你可以unjail你的验证人。

最后，检查你的验证人看看投票股权是否恢复：

```bash
plugchaind status
```

你可能会注意到你的投票权比之前要少。这是由于你的下线受到的削减处罚！


### 问题 #2 : 我的`plugchaind`由于`too many open files`而崩溃

Linux可以打开的默认文件数（每个进程）是1024。已知`plugchaind`可以打开超过1024个文件。这会导致进程崩溃。快速修复运行`ulimit -n 4096`（增加允许的打开文件数）来快速修复，然后使用`plugchaind start`重新启动进程。如果你使用`systemd`或其他进程管理器来启动`plugchaind`，则可能需要在该级别进行一些配置。解决此问题的示例`systemd`文件如下：

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
