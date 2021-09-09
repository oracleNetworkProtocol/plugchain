---
order: 4
---

# 加入测试网

测试网现在处于活动状态，以下是一些重要的详细信息：

:::tip
需要先 [安装 plugchaind](install.md)
:::
## 公共端点

- [GRPC:8.210.180.240:9090]()
- [RPC:http://8.210.180.240:26657/](http://8.210.180.240:26657/)
- [REST:http://8.210.180.240:1317/](http://8.210.180.240:1317/)



## 运行全节点

#### 从genesis开始运行节点

:::提示
使用 PLUGChain[v0.5.0](https://github.com/oracleNetworkProtocol/plugchain.git) 初始化你的节点::
:::

1. 初始化节点

```bash
plugchaind init <moniker> --chain-id=plugchain
```

2. 下载测试网公开的 genesis.json和种子信息 或者 进入clone下来的 plugchain 目录里面：
*[测试网创世文件和种子信息](https://github.com/oracleNetworkProtocol/plugchain/blob/main/testnet/latest/)* ,如下通过移动clone的仓库里面的genesis.json 来实现覆盖。

```bash 
mv ./testnet/latest/genesis.json ~/.plugchain/
```

:::warning 
如下第三步可跳过，如跳过需要在第四步加上参数 `--p2p.seeds` 参数）
:::

3. 覆盖数据目录的genesis.json之后,修改 ~/.plugchain/config/config.toml 里面 seeds 参数，添加种子信息

把 ./testnet/latest/seeds.txt 文件中提供的 seeds , 修改 ~/.plugchain/config/config.toml 中的`seeds`参数以设置链接的种子节点,种子信息以英文逗号分隔


4. 启动节点服务
```bash
# 启动节点（也可使用 nohup 或 systemd 等方式后台运行）

# 第三步未修改种子信息，运行start时，添加参数 --p2p.seeds="5f81625b69d192d3ef5bf47b83484326e0546491@47.100.161.102:26656"

# 如果修改服务端口配置，需要在使用服务的地方加上参数:
# 比如修改默认的tendermint rpc服务: tcp://localhost:26657 => tcp://localhost:5000 
# cli使用时，有`--node`参数的命令 都需要指向此参数为 --node=tcp://localhost:5000
# 例如： plugchaind q account gx1tulwx2hwz4dv8te6cflhda64dn0984harlzegw --node tcp://localhost:5000


plugchaind start --minimum-gas-prices 0.025line
```



## 水龙头

欢迎加入我们的水龙头链接[test explorers](http://test.plugchain.network/wallet/receive)

## 浏览器

<http://test.plugchain.network/>



