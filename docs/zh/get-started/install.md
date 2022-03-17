---
order: 2
---

# 安装

以centos为例：

依赖工具
- git
- make
- gcc
- wget

```shell
yum install make gcc git wget
```
## 安装`go`

:::tip
编译安装 Plug Chain Hub 软件依赖 **Go 1.16+**。
:::

按照[官方文档](https://golang.org/doc/install) 安装`go`(>=1.16)。

别忘记设置您的`$GOPATH`，`$GOBIN`和`$PATH`环境变量，例如：

```bash
mkdir -p $HOME/go/bin
echo "export GOPATH=$HOME/go" >> ~/.bashrc
source ~/.bashrc
echo "export GOBIN=$GOPATH/bin" >> ~/.bashrc
source ~/.bashrc
echo "export PATH=$PATH:$GOBIN" >> ~/.bashrc
source ~/.bashrc
```

确认已成功安装`go`

```bash
go version
```

## 安装`plugchaind`

正确配置`go`之后，您应该可以编译并运行`plugchaind`了。

请确保您的服务器可以访问 google.com，因为我们的项目依赖于google提供的某些库（如果您无法访问`google.com`，也可以尝试添加代理：`export GOPROXY=https://goproxy.io`,或者 `export GOPROXY=https://goproxy.cn,direct` ）

```bash
git clone https://github.com/oracleNetworkProtocol/plugchain.git
cd plugchain
git checkout v1.0.0
make install
```
> 如果克隆仓库比较慢或者没反应的情况，可以把克隆地址 `github.com` 换成国内镜像，已知的国内镜像有`github.com.cnpmjs.org`和`git.sdut.me/`

如果环境变量配置无误，则通过运行以上命令即可完成`plugchaind`的安装。现在检查您的`plugchaind`版本是否正确：

```bash
plugchaind version
```
