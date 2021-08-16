  ### 测试网目前需要操作步骤 
1. 节点搭建，同步测试网 [搭建文档](https://oraclenetworkprotocol.github.io/plugchain/cn/testnet.html)
2. 成为[验证者](https://oraclenetworkprotocol.github.io/plugchain/cn/validators/validator-setup.html)，会在[验证者](http://www.plugchain.network/wallet/pledge)页面出现自己的节点
3. 在终端或者[浏览器](http://www.plugchain.network/)委托line币到自己的节点，查看节点信息等操作
4. 查看终端各个命令的用处和使用方式
5. 开通自己节点的对外api服务,[文档](https://oraclenetworkprotocol.github.io/plugchain/cn/api/swagger-api.html) 
6. 关于不同节点的配置，请移步[节点知识](https://oraclenetworkprotocol.github.io/plugchain/cn/node/)，配置修改：验证者的配置需要成为验证者之后修改，种子和哨兵的配置可以直接修改，重启即可成为此身份

**[!warning]** 
  切记，保存好自己创建的地址的助记词，如下信息中的：`pride cricket area future day trust pattern exhaust orange mouse chronic game make tobacco paddle float tuition vocal shove tag gas cargo idea label`一串字母

```text
  - name: testwallet
  type: local
  address: gx1fp5qd7ztst597qzfyhcmxf0jlrdddex53s2837
  pubkey: gxpub1addwnpepqvmancaq7l52sshz4k09y9v7ftlt2cdaqgwpml4x2kga7xvspcrjq8z496t
  mnemonic: ""
  threshold: 0
  pubkeys: []


**Important** write this mnemonic phrase in a safe place.
It is the only way to recover your account if you ever forget your password.

pride cricket area future day trust pattern exhaust orange mouse chronic game make tobacco paddle float tuition vocal shove tag gas cargo idea label
```


##### 相关知识点
- 节点
  * 在线节点总质押量超过区块总质押量的2/3才出块
  * 种子节点，需要修改参数seed_mode 为true ，当作为种子发现节点时，内网通信需要修改（config.toml） p2p的 laddr 为本机当时ip
  * 自己的节点想共享出去，需要把公网ip写入external_address 参数里
- 手续费
  * 区块链小数位数为1000000，在链上都是乘1000000
  * app.toml 节点运行需要设置最低手续费防止垃圾邮件 ，minimum-gas-prices = "0.0001line"
  * 手续费= (gas-prices   * gas)或者 fees
  * fees/1000000 为真正扣除手续费
- 验证者
  * config.toml文件， 验证者：double_sign_check_height = 0 参数：表示验证者断开区块链重新加入之后向上验证的块数，0表示不验证断开期间的块，其余正整数表示向上验证几个区块，当此数字大于期间出的块数时，会宕机连不上区块，等待断开期间出块数大于等于此值才可连上
  * create-validator， 字段 `--commission-max-change-rate` `--commission-max-rate` 设置之后不准修改，谨慎设置
  * edit-validator， 修改验证者信息时，各字段修改间隔24小时，min-self-delegation 不可减小，只能增加
- 质押限制
  + 质押,取消质押:
    * 取消质押币需要冻结21天后到账
  + 转质押:
    * 转质押后新节点对用户的转质押操作锁仓21天(质押,取消质押不受影响)
    * 转质押一月可以操作7次
  + 收益:
    * 收益需要提取才能到可操作账户
    * 进行质押,取消质押,转质押操作时自动提取收益

**测试期间有问题可以总结起来,提交给技术部，共同维护plugchain**