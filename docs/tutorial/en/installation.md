## Install plugchain
This guide will install the plugchaind entry point on the system, you can join as a full node or validator
#### Install build requirements
- On Centos this can be done with the following:
```
yum install make gcc git
```
#### install Golang
-  verison (>=1.16)
- [golang version](https://studygolang.com/dl)
- The v1.16.4 version downloaded in this tutorial
```sh
wget https://studygolang.com/dl/golang/go1.16.4.linux-amd64.tar.gz
```
```sh
rm -rf /usr/local/go && tar -C /usr/local -zxf go1.16.4.linux-amd64.tar.gz 
```
- Add `/usr/local/go/bin` to the PATH environment variable. 
    You can do this by adding the following line to your `/etc/profile` or `~/.bash_profile` (for a system-wide installation):
```
echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bash_profile
```
**Note**
Changes made to a profile file may not apply until the next time you log into your computer. To apply the changes immediately, just run the shell commands directly or execute them from the profile using a command such as source `~/.bash_profile`.

```
source ~/.bash_profile
go version
```
1.  Install the latest version of plugchain and clone the code from the warehouse
```
git clone https://github.com/oracleNetworkProtocol/plugchain.git
```
> If cloning the warehouse is slow or unresponsive, you can replace the clone address `github.com` with a domestic mirror. The known domestic mirrors are `github.com.cnpmjs.org` and `git.sdut.me/`

2.  Open gomodule and set terminal GOPROXY
```
go env -w GO111MODULE=on
export GOPROXY=https://goproxy.io,direct
```
#### Compile the binary file
```shell
cd plugchain && make install
```
##### The main branch code is compiled by default, and fixed version compilation can be selected, for example:
> `cd plugchaind && git checkout v0.2.0 && make install`

3. After the installation is complete, there will be a plugchaind binary file, you can verify whether it takes effect
```
plugchaind version
```
Confirm that the command prints the installed version of Go.


#### To run as a full node or a validator node, please copy the genesis file to your own data directory first
- Initialize your own node
```
plugchaind init mycustomMoniker --chain-id plugchain
```

- Set the data directory and initialize the default directory to ~/.plugchain. If you want to change it, please delete the default and reinitialize your node. Point to --home, `plugchaind init mycustomMoniker --home node --chain-id plugchain`

```shell
APPHOME="~/.plugchain"
```
- Generate your own account
```
plugchaind keys add mywallet
```
- Get the testnet genesis file
::: warning
The mainnet is being prepared. Join [Testnet](../../../testnet/README.md), the plugchain team warmly welcomes like-minded comrades...
:::
```
wget https://github.com/oracleNetworkProtocol/plugchain/tree/main/genesis.json
```
- Overwrite genesis file
```
mv -f genesis.json $APPHOME/config/
```
- Need to add the normal seed node to $APPHOME/config/config.toml, otherwise the node will not work normally 
- #TODO


- Your full node saves unconfirmed transactions in its memory pool. To prevent spam, it is best to set a minimum gas price that a transaction must meet in order to be accepted in your node's memory pool. This parameter `min-gas-prices` can be set in the following file $APPHOME/config/app.toml.

- Run `plugchaind start` to participate in the project