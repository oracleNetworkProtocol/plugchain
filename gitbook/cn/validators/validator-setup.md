#### 成为验证者参与主链生态管理

- 以下教程都依赖节点执行,如未安装,请移步到 [安装节点](../installation.md)

#### 添加验证者
* plugchaind tx staking create-validator [flags]
```shell
plugchaind tx staking create-validator --from mywallet \
--amount 1000000plug --pubkey $(plugchaind tendermint show-validator) \
--moniker="my validator" --commission-rate="0.10" --commission-max-rate="0.20" \
--commission-max-change-rate="0.01"  --min-self-delegation="1000000" --fees 20plug --chain-id plugchain
```
#### 修改已经存在的验证者配置
* plugchaind tx staking edit-validator [flags]
```shell
plugchaind tx staking edit-validator --from mywallet \
--moniker="my validator" --commission-rate="0.20" --min-self-delegation="1000000" --fees 20plug --chain-id plugchain
```
**[!warning]** : 

`commission-max-rate`,`commission-max-change-rate` 两个字段在`create-validator`设置完之后就不可以修改,请您慎重设置

`edit-validator` 修改验证者信息时,各字段修改间隔24H,`min-self-delegation` 不可减小,只能增加,没有更改的字段修改时不可添加

`commission-rate`的值必须符合如下的不变量检查：

+ 必须在 0 和 验证人的`commission-max-rate` 之间
+ 不得超过 验证人的`commission-max-change-rate`, 该参数标识**每日**最大的百分点变化数。也就是,一个验证人在`commission-max-change-rate`的界限内每24H一次可调整的最大佣金变化。

`min-self-delegation`的值单位是`10e-6`,所以此值最小为 `1000000`


#### 查询验证者的所有委托人的质押总量,想了解委托人信息,请移步[委托人质押](../delegators/delegator-setup.md)
* plugchaind query staking delegations-to [validator-addr] [flags]
```shell
plugchaind query staking delegations-to gxvaloperxxxxxxxxxxxxxxxxxx  
```

#### 根据验证者查询所有取消质押的记录
*  plugchaind query staking unbonding-delegations-from [validator-addr] [flags]
```shell
plugchaind query staking unbonding-delegations-from gxvaloperxxxxxxxxxxxxxxxxxxxxxxxxxxxx  
```


::: warning 注意
为了能进入验证人集合,你的权重必须超过第100名的验证人。
:::

- 成为验证者之后,搭建哨兵节点保护验证者节点时，需要修改验证者配置信息[validator config](../node/README.md)
- 修改验证者配置，需要注意配置文件里 `double_sign_check_height` 参数，验证者节点断开期间，链出了多少个块，需要在此参数里写上多少个块，表示向上验证多少个区块，防止进入监禁状态。


## 常见问题

### 验证者四种状态

| 状态(status) | 状态码 | 描述  | 状况  | 
| ---- | ------ | ----------- | ---- | 
| BOND_STATUS_UNSPECIFIED | 0 | 定义了无效的验证器状态 | 无效 | 
| BOND_STATUS_UNBONDED | 1 | 定义了一个未绑定的验证器 | 离线 | 
| BOND_STATUS_UNBONDING | 2 | 定义了一个解除绑定的验证器 | 监禁 | 
| BOND_STATUS_BONDED | 3 | 定义了一个绑定的验证器 | 正常 | 

### 验证者信息类似如下：

```
commission:
  commission_rates:
    max_change_rate: "0.010000000000000000"
    max_rate: "0.200000000000000000"
    rate: "0.100000000000000000"
  update_time: "2021-07-27T13:19:37.983601182Z"
consensus_pubkey:
  '@type': /cosmos.crypto.ed25519.PubKey
  key: 5TTh7vhc3RvibPttF29r2iC6qSNfGbJdY9sVsQHBT8w=
delegator_shares: "8813835048395.526196202816719238"
description:
  details: ""
  identity: ""
  moniker: My Validator
  security_contact: ""
  website: ""
jailed: true
min_self_delegation: "1000000"
operator_address: gxvaloper19quyxaq87dpsxt9w2q23c4d8x3fck4fzh4u8es
status: BOND_STATUS_UNBONDING
tokens: "8466534803455"
unbonding_height: "831092"
unbonding_time: "2021-09-02T18:02:44.032627007Z"

```

### 当验证者处于 `jailed:true` 态度时，并且`stautus:BOND_STATUS_UNBONDING`说明验证者节点已经处理惩罚状态，需要解除惩罚状态：

```
plugchaind tx slashing  unjail --from mykey --chain-id plugchain
```
- --from mykey ,`mykey` 是成为验证者时，操作的地址名称
- 解禁时，必须保证当前验证者委托量大于节点配置参数`min_self_delegation`的值，不够的，通过委托币，超过此参数
