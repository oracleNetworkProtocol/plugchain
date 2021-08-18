### 常见问题

##### genesis文件

- 请注意，所有持续时间均以纳秒为单位。
- mint模块参数 ，出块造币（mint_denom:onp）数量 = total（mint_denom:onp币种的发行总量） / blocks_per_year（年出块数量） * inflation_rate_change(膨胀率)

##### 开发问题

- 节点
  * 在线节点总质押量超过区块总质押量的2/3才出块
  * 自己的节点想共享出去，需要把公网ip写入external_address 参数里

- 手续费
  * 区块链小数位数为1000000，在链上都是乘1000000
  * app.toml 节点运行需要设置最低手续费防止垃圾邮件 ，minimum-gas-prices
  * 手续费= (gas-prices  * gas )或者 fees
  * gas 默认值 200000
  * fees/1000000 为真正扣除手续费

- 验证者
  * config.toml文件， 验证者：double_sign_check_height = 0 参数：表示验证者断开区块链重新加入之后向上验证的块数，0表示不验证断开期间的块，其余正整数表示向上验证几个区块，当此数字大于期间出的块数时，会宕机连不上区块，等待断开期间出块数大于等于此值才可连上
  * create-validator， 字段 `--commission-max-change-rate` `--commission-max-rate` 设置之后不准修改，谨慎设置
  * edit-validator， 修改验证者信息时，各字段修改间隔24小时，min-self-delegation 不可减小，只能增加

- 质押问题

  + 质押,取消质押:
    * 取消质押币需要冻结21天后到账
  + 转质押:
    * 转质押后新节点对用户的转质押操作锁仓21天(质押,取消质押不受影响)
    * 转质押一月可以操作7次
  + 收益:
    * 收益需要提取才能到可操作账户
    * 进行质押,取消质押,转质押操作时自动提取收益