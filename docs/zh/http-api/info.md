---
order: 1
---

# API接口

## 请求URL

- 主域名 :`https://api.plugchain.network`

### API列表

| NAME                                     | METHOD | PATH                               |
|------------------------------------------|--------|------------------------------------|
| [创建地址](./account_new_address.md)         | GET    | `/server/v1/account/new_address`   |
| [地址转换](./account_addr_convert.md)        | GET    | `/server/v1/account/addr_convert`  |
| [账户信息](./account_info.md)                | GET    | `/server/v1/account/info`          |
| [账户余额 prc10](./account_balance_prc10.md) | GET    | `/server/v1/account/balance/prc10` |
| [账户余额 prc20](./account_balance_prc20.md) | GET    | `/server/v1/account/balance/prc20` |
| [账户交易记录](./account_txs.md)               | GET    | `/server/v1/account/txs`           |
| [prc10 代币列表](./token_prc10_list.md)      | GET    | `/server/v1/token/prc10/tokens`    |
| [prc10 代币详情](./token_prc20_info.md)      | GET    | `/server/v1/token/prc10/token`     |
| [prc20 代币详情](./token_prc20_info.md)      | GET    | `/server/v1/token/prc20/token`     |
| [获取行情](./market_pc.md)                   | GET    | `/server/v1/market/pc`             |
| [广播交易](./tx_broadcast.md)                | POST   | `/server/v1/tx/broadcast`          |
| [发送交易prc10](./tx_send_tx_prc10.md)       | POST   | `/server/v1/tx/send_tx_prc10`      |
| [发送交易PVM](./tx_send_tx_pvm.md)           | POST   | `/server/v1/tx/send_tx_pvm`        |
| [最新交易](./tx_latest.md)                   | GET    | `/server/v1/tx/latest`             |
| [区块内交易](./block_tx.md)                   | GET    | `/server/v1/tx/block`              |
| [交易详情](./tx_by_hash.md)                  | GET    | `/server/v1/tx/by_hash`            |
| [通过hash获取块](./block_by_hash.md)          | GET    | `/server/v1/block/by_hash`         |
| [通过块高获取块](./block_by_height.md)          | GET    | `/server/v1/block/by_height`       |
| [最新区块](./block_last.md)                  | GET    | `/server/v1/block/last`            |
| [验证者列表 ](./validator_list.md)            | GET    | `/server/v1/validator/validators`  |
| [验证者详情](./validator_info.md)             | GET    | `/server/v1/validator/info`        |
| [提案列表](./gov_list.md)                    | GET    | `/server/v1/gov/proposals`         |
| [提案详情](./gov_info.md)                    | GET    | `/server/v1/gov/detail`            |
| [节点信息](./node_info.md)                   | GET    | `/server/v1/node_info`             |
| [交易池prc_10列表](./swap_prc10_pools.md)     | GET    | `/server/v1/swap/prc10/pools`      |
| [交易池prc_10](./swap_prc10_pool.md)        | GET    | `/server/v1/swap/prc10/pool`       |
| [交易池prc_20列表](./swap_prc10_pools.md)     | GET    | `/server/v1/swap/prc20/pools`      |
| [交易池prc_20](./swap_prc10_pool.md)        | GET    | `/server/v1/swap/prc20/pool`       |
             
