---
order: 25
---

# 获取最近交易记录
 
- 获取最近的10条交易记录

### 请求方式
`GET`

### URL
`/server/v1/tx/latest`

### 请求示例

```
/server/v1/tx/latest
```


#### 请求参数
- 无
### 返回值
- `code int64 `  : 0或字段不存在为成功,其他均为失败
- `message string` : 响应信息
- `data object` : 成功返回数据
    - `count int` : 总条数
    - `txs array`: 交易数组
      - `tx object`: 交易数据
         - `height string` : 块高,
         - `hash string` : 交易hash,
         - `module string` : 交易模块,
         - `action string` : 交易模块分支,
         - `type string` : 交易到达类型,
         - `signer string` : 交易签名者,
         - `fee_amount string` : 交易手续费,
         - `fee_coin string` : 交易手续费币种,
         - `code string` : 响应代码,
         - `code_err string` : 错误信息,
         - `status string` : 交易状态,
         - `status_err string` : 状态错误信息,
         - `coins string` : 交易包含币种,
         - `memo string` : 备注,
         - `create_time string` : 创建时间
      - `msg string` : 交易json数据

#### 返回示例
```json
{
  "code": 0,
  "data": {
    "count": 4600,
    "txs": [
      {
        "tx": {
          "height": 8160857,
          "hash": "1214C80429CA33776ACF5F1AA1900EA9EFC20119CFB1C7E5AD22733E89FAA8A4",
          "module": "bank",
          "action": "MsgSend",
          "type": "Transfer",
          "signer": "gx1sqyqeef3ym6q9x7anlgxmxh6u7sjfx6ccc3d4m",
          "fee_amount": 200,
          "fee_coin": "uplugcn",
          "code": 0,
          "code_err": "",
          "status": "success",
          "status_err": "",
          "coins": "uplugcn",
          "memo": "",
          "create_time": "2023-02-10 16:04:44"
        },
        "msg": "{\"@type\":\"/cosmos.bank.v1beta1.MsgSend\",\"amount\":[{\"amount\":\"200000000\",\"denom\":\"uplugcn\"}],\"from_address\":\"gx1sqyqeef3ym6q9x7anlgxmxh6u7sjfx6ccc3d4m\",\"to_address\":\"gx1pu006rclnmkhs3y26v7euxstpr4ftgp5cfkctc\"}"
      }
    ]
  },
  "message": "ok"
}
```