#### 委托人质押
- 委托人质押需要选择验证者,查询所有验证者
    * plugchaind q staking validators [flags]
```
plugchaind q staking validators --chain-id plugchain
```
- 找到自己满意的验证者，执行质押
    * plugchaind tx staking delegate [validator-addr] [amount] --from mykey [flags]
```
plugchaind tx staking delegate gxvaloperxxxxxxxxxxxxxxxxxxxxx 20000plug --from gxxxxxxxxxxxxxxxxxxxxx --chain-id plugchain
```
- 更改自己的委托验证者
    *   plugchaind tx staking redelegate [src-validator-addr] [dst-validator-addr] [amount] --from mykey [flags]
```
plugchaind tx staking redelegate gxvaloperxxxxxxxxxxxxxxxxxxxx gxvaloperxxxxxxxxxxxxxxxxxxx 100000plug --from gxxxxxxxxxxxxxxxxx --chain-id plugchain
```
