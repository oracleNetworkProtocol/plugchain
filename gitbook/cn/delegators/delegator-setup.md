#### 委托人质押
* 委托人质押需要选择验证者,查询所有验证者
* plugchaind q staking validators [flags]
```shell
plugchaind q staking validators --chain-id plugchain
```
#### 找到自己满意的验证者,执行质押
* plugchaind tx staking delegate [validator-addr] [amount] --from mykey [flags]
```shell
plugchaind tx staking delegate gxvaloperxxxxxxxxxxxxxxxxxxxxx 20000plug \
 --from mywallet --fees 20plug --chain-id plugchain
```

#### 更改自己的委托验证者
* plugchaind tx staking redelegate [src-validator-addr] [dst-validator-addr] [amount] --from mykey [flags]
```shell
plugchaind tx staking redelegate gxvaloperxxxxxxxxxxxxxxxxxxxx gxvaloperxxxxxxxxxxxxxxxxxxx 100000plug \
  --from mywallet --fees 20plug --chain-id plugchain
```

#### 根据验证者查询未提取收益
* plugchaind query distribution rewards [delegator-addr] [validator-addr] [flags]
* delegator-addr 是验证者地址
* validator-addr 为 `plugchaind q staking validators  ` 命令根据moniker找到验证者信息,对应的validator-addr
```shell
plugchaind query distribution rewards $(plugchaind keys show mywallet -a) [validator-addr]  
```

#### 提取所有未提取奖励
* plugchaind tx distribution withdraw-all-rewards --from mykey [flags]
```
plugchaind tx distribution withdraw-all-rewards --from mywallet  
```

#### <span id="unbond">取消质押-锁仓21天 </span>
* plugchaind tx staking unbond [validator-addr] [amount] --from mykey
```shell
plugchaind tx staking unbond gxvaloperxxxxxxxxxxxxxxxxxx 10000000plug --from mywallet --chain-id plugchain
```

#### 质押限制

    1.质押,取消质押:
        取消质押币需要冻结21天后到账
    
    2.转质押:
        转质押后新节点对用户的转质押操作锁仓21天(质押,取消质押不受影响)
        转质押一月可以操作7次

    3.收益:
        收益需要提取才能到可操作账户
        进行质押,取消质押,转质押操作时自动提取收益