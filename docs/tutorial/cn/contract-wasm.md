## installation 

###  GO
- version >= v1.15
- 安装 [指定版本](./README.md)

### Rust
- 安装一些rust工具,标准的方法是使用rustup来管理依赖，并处理您将要使用的cargo和rustc的多个版本的更新。
- [安装rust](https://rustup.rs/)
- 安装之后，确保您具有wasm32目标：
```sh
rustup default stable
cargo version
#如果低于1.50.0+，请更新rustup更新稳定
rustup update stable

rustup target list --installed
rustup target add wasm32-unknown-unknown
```

***

## 编译并测试合约 
### 以 cw20-base 为例
```sh
git clone git@github.com:CosmWasm/cosmwasm-plus.git
#编译合约
cd cosmwasm-plus/contracts/cw20-base && RUSTFLAGS='-C link-arg=-s' cargo wasm
#运行测试
RUST_BACKTRACE=1 cargo unit-test


APP_HOME="home1"
CHAIN_ID="localnet"

#查询我们现在有多少代码
plugchaind q wasm list-code 

RES=$(plugchaind tx wasm store ./cosmwasm-plus/target/wasm32-unknown-unknown/release/cw20_base.wasm --from mywallet --home ${APP_HOME} --chain-id ${CHAIN_ID} --gas-prices 0.01onp --gas 1800000 -y)

CODE_ID=$(echo $RES | jq -r '.logs[0].events[0].attributes[-1].value')

#合约实例化内容
INIT=$(jq -n --arg mywallet $(plugchaind keys show -a mywallet --home ${APP_HOME}) --arg mymain $(plugchaind keys show -a mymain --home ${APP_HOME}) '{"name":"my one wasm","symbol":"OHOU","decimals":2,"initial_balances":[{"address":$mywallet,"amount":"100000000"}],"minter":$mywallet}')

#合约实例化
plugchaind tx wasm instantiate $CODE_ID "$INIT" --from mywallet --amount=50000onp  --label "escrow 1" --gas-prices="0.01onp" --gas 200000 --home ${APP_HOME} --chain-id ${CHAIN_ID}

#获取合约地址
CONTRACT=$(plugchaind query wasm list-contract-by-code $CODE_ID --output json --home ${APP_HOME} | jq -r '.contracts[-1]')

#合约查询方法
QUERY='{"balance":{"address":"onp16tmztx92nu3pthghs4k5dwe0hvhq0f57jsc27p"}}'
plugchaind query wasm contract-state smart $CONTRACT "$QUERY" --home ${APP_HOME}

#合约执行方法
TRANSF='{"transfer":{"recipient":"接收地址","amount":"50000"}}'
plugchaind tx wasm execute $CONTRACT "$TRANSF" --from mywallet --gas-prices "0.01onp" --gas 200000 --home ${APP_HOME} --chain-id ${CHAIN_ID} -y

```
