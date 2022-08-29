---
order: 4
---

# 加入测试网

测试网现在处于活动状态，以下是一些重要的详细信息：

:::tip
需要先 [安装 plugchaind](install.md)
:::
## 公共端点

- [Cosmos-GRPC:47.108.75.227:9090]()
- [Cosmos-REST:http://47.108.75.227:1317/](http://47.108.75.227:1317/)
- [Tendermint-RPC:http://47.108.75.227:26657/](http://47.108.75.227:26657/)
- [JSON-RPC:http://47.108.75.227:8545/](http://47.108.75.227:8545/)


## 运行全节点

#### 从genesis开始运行节点


1. 初始化节点

```bash
plugchaind init <moniker> --chain-id=plugchain_521-1
```

2. 下载测试网公开的 [genesis.json](https://github.com/oracleNetworkProtocol/testnet)和对等点信息:

3. 修改配置config.toml文件，`private_peer_ids` 赋值增加对等点。

4. 启动节点服务
```bash
# 启动节点（也可使用 nohup 或 systemd 等方式后台运行）

# 如果修改服务端口配置，需要在使用服务的地方加上参数:
# 比如修改默认的tendermint rpc服务: tcp://localhost:26657 => tcp://localhost:5000 
# cli使用时，有`--node`参数的命令 都需要指向此参数为 --node=tcp://localhost:5000
# 例如： plugchaind q account gx1tulwx2hwz4dv8te6cflhda64dn0984harlzegw --node tcp://localhost:5000


plugchaind start
```

## 浏览器&&水龙头

<https://test.plugchain.network/v2>



## cosmoswap 相关信息

```
router路由合约:
bech32: gx14ewpfal7g758zjv0vu9sthcfllagnehh8s28np
eip55: 0xae5C14F7FE47a871498F670B05dF09fffa89e6F7

fac工厂合约:
bech32: gx1anezwnqmssja3phgtz6rjmx5utkry3mptxh0qw
eip55: 0xecf2274c1B8425D886E858b4396cd4E2EC324761

WPLUG 合约地址：
bech32: gx1nqtlcdxd0yq5ydtqv95qguptpzx9cmpa8ga2r5
eip55: 0x9817fC34CD7901423560616804702b088C5C6C3D
```
