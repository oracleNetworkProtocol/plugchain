# 账户交易

### 一. 查看**tx**参数

```sh
Transactions subcommands

Usage:
  plugchaind tx [flags]
  plugchaind tx [command]

Available Commands:
  bank                Bank transaction subcommands
  broadcast           Broadcast transactions generated offline
  crisis              Crisis transactions subcommands
  decode              Decode an binary encoded transaction string.
  distribution        Distribution transactions subcommands
  encode              Encode transactions generated offline
  evidence            Evidence transaction subcommands
  gov                 Governance transactions subcommands
  ibc                 IBC transaction subcommands
  ibc-transfer        IBC fungible token transfer transaction subcommands
  multisign           Generate multisig signatures for transactions generated offline
  sign                Sign a transaction generated offline
  sign-batch          Sign transaction batch files
  slashing            Slashing transaction subcommands
  staking             Staking transaction subcommands
  token               token transactions subcommands
  validate-signatures Validate transactions signatures
  vesting             Vesting transaction subcommands
  vesting             Vesting transaction subcommands

Flags:
      --chain-id string   The network chain ID (default "token")
  -h, --help              help for tx

Global Flags:
      --home string         directory for config and data (default "/root/.token")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors

Additional help topics:
  plugchaind tx upgrade      Upgrade transaction subcommands

Use "plugchaind tx [command] --help" for more information about a command.
```
##### - 其中我们经常用到的有
|模块| 概述|
|-|-|
|bank|转账模块|
|ibc|跨链通信模块|
|staking|质押模块|
### 二. bank-send 转账参数
- 使用 `plugchaind tx bank send -h` 查看用法和参数

```sh
Usage:
  tokend tx bank send [from_key_or_address] [to_address] [amount] [flags]

Flags:
  -a, --account-number uint      The account number of the signing account (offline mode only)
  -b, --broadcast-mode string    Transaction broadcasting mode (sync|async|block) (default "sync")
      --dry-run                  ignore the --gas flag and perform a simulation of a transaction, but don't broadcast it
      --fees string              Fees to pay along with transaction; eg: 10uatom
      --from string              Name or address of private key with which to sign
      --gas string               gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default 200000)
      --gas-adjustment float     adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default 1)
      --gas-prices string        Gas prices in decimal format to determine the transaction fee (e.g. 0.1uatom)
      --generate-only            Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)
  -h, --help                     help for send
      --keyring-backend string   Select keyring's backend (os|file|kwallet|pass|test|memory) (default "test")
      --keyring-dir string       The client Keyring directory; if omitted, the default 'home' directory will be used
      --ledger                   Use a connected Ledger device
      --memo string              Memo to send along with transaction
      --node string              <host>:<port> to tendermint rpc interface for this chain (default "tcp://localhost:26657")
      --offline                  Offline mode (does not allow any online functionality
  -s, --sequence uint            The sequence number of the signing account (offline mode only)
      --sign-mode string         Choose sign mode (direct|amino-json), this is an advanced feature
      --timeout-height uint      Set a block timeout height to prevent the tx from being committed past a certain height
  -y, --yes                      Skip tx broadcasting prompt confirmation

Global Flags:
      --chain-id string     The network chain ID (default "token")
      --home string         directory for config and data (default "/root/.token")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```
##### 交易参数名词解释
- ** tokend tx bank send  [from_key_or_address] [to_address] [amount] [flags] **

|参数|解释|
|-|-|
|**from_key_or_address**|发送方地址|
|**to_address**|接收方地址|
|**amount**|金额|
|**flags**|其他参数(--home --chain-id等)|
|**--fees**|手续费|
|**--gas**|手续费limit|
|**--gas-price**|手续费价格|

- 其中发起交易的时候需要额外设置手续费参数
- `plug`的单位有6位小数 结果要`/1000000`才是正常的单位 即命令行中的`10000plug`实际为`0.01plug`)

### 三. 实现账户之间的转账

##### 1. 直接指定本次交易手续费(10000plug) `--fees` 
```sh
plugchaind tx bank send $from $to 10plug --fees 10000plug --chain-id testnet --home home1
```
#####  2. 分开设置,那么总用量就是 gas*gas-pirce 结果是(20000plug) 
   - `--gas` 手续费用量 默认 200000
   - `--gas-prices` gas价格 比如 0.1plug
```sh
plugchaind tx bank send $from $to 100line --gas 200000 --gas-prices 0.0001line --chain-id plugchain-testnet-1 --home home1
```

#### 注意事项
1. 在设置手续费的时候要注意各个节点设置的最小手续费价格 不能小于最小价格 `minimum-gas-prices = "0.1plug"`