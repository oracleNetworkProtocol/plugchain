## INSTALL plugchain
本指南将安装plugchaind入口点到系统上，您可以作为完整节点或验证者加入
#### 安装依赖
- centos系统示例如下:
```
yum install make gcc git
```
#### install Golang

- golang 安装(>=1.16)

- [golang各种版本下载](https://studygolang.com/dl)

- 本教程下载的v1.16.4版本,官方安装golang教程[docs](https://golang.org/doc/install)

```sh
wget https://studygolang.com/dl/golang/go1.16.4.linux-amd64.tar.gz
```

```sh
rm -rf /usr/local/go && tar -C /usr/local -zxf go1.16.4.linux-amd64.tar.gz
```

```
export PATH=$PATH:/usr/local/go/bin 
```

```
mkdir -p $HOME/go/bin
echo "export PATH=$PATH:$(go env GOPATH)/bin" >> ~/.bash_profile
echo "export GOPATH=$HOME/go" >> ~/.bash_profile
```

```
source ~/.bash_profile
go version
```

- 在Windows下，你可以通过"系统"控制面板的"高级"选项卡上的"环境变量"按钮设置环境变量(HOME或GO111MODULE)。一些版本的Windows通过"系统"控制面板中的"高级系统设置"选项提供了这个控制面板。

1.  安装最新版的plugchain,从仓库clone代码

```
git clone https://github.com/oracleNetworkProtocol/plugchain.git
```
> 如果克隆仓库比较慢或者没反应的情况，可以把克隆地址 `github.com` 换成国内镜像，已知的国内镜像有`github.com.cnpmjs.org`和`git.sdut.me/`

2.  开启gomodule和设置终端GOPROXY

```
go env -w GO111MODULE=on
export GOPROXY=https://goproxy.io,direct
```

> 下一步 `make install` 如果包下载超时，可以选择更换下代理，`export GOPROXY=https://goproxy.cn,direct`
#### 编译二进制文件
```shell
cd plugchain && make install
```
##### (可选)默认编译主分支代码，可选择固定版本编译，例如: 
> `cd plugchaind && git checkout v0.3.0 && make install`

3. 安装完成之后，会有plugchaind二进制文件，可以验证下是否生效

```
plugchaind version
```

输出内容示例如下

```
v0.3.0
```
####  完成 `plugchaind` 安装，就可以选择加入主网或者测试网了。
- [加入主网](./mainnet.md)
- [加入测试网](./testnet.md)

<!-- - 初始化节点，默认数据目录

```
plugchaind init mycustomMoniker --chain-id plugchain
```

- （可选）初始化节点，自定义数据目录 ，默认数据目录为 ~/.plugchain，如要更改，请把默认的删除，重新初始化自己的节点 指向--home, `plugchaind init mycustomMoniker --home node --chain-id plugchain`

```shell
APPHOME="~/.plugchain"
```

- 初始化完成之后,获取主网创世文件,下载 `https://github.com/oracleNetworkProtocol/plugchain/blob/main/mainnet/v1/genesis.json` 文件

- 把下载的`genesis.json`文件，覆盖掉数据目录的 `$APPHOME/config/genesis.json`

- 需要将正常的种子节点添加到$APPHOME/config/config.toml里，不然节点无法正常工作 ，种子信息在 `https://github.com/oracleNetworkProtocol/plugchain/blob/main/mainnet/v1/seeds.txt` 文件里面

- 修改--home目录下的 `$APPHOME/config/config.toml` 中的`seeds`参数,把seeds.txt文件中的种子信息，填充到`seeds`字段里面

- 您的全节点将未确认的事务保存在其内存池中。为了防止垃圾邮件，最好设置一个交易必须满足的最小gas价格，以便在您的节点的内存池中被接受。这个参数 `min-gas-prices` 可以在下面的文件$APPHOME/config/app.toml中设置,也可以直接启动时设置。如下：

- 运行 `plugchaind start --minimum-gas-prices 0.0001plug` 加入到主网中。。。 -->