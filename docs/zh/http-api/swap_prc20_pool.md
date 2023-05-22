---
order: 18
---

# SWAP交易池(PRC20)

- swap交易池详情(prc20)

### 请求方式
`GET`

### URL
`/server/v1/swap/prc20/pool`

### 请求示例

```
/server/v1/swap/prc20/pool?contract=gx1rmf6rcdfcstl94mscdtgshjzy8q0d2ljjy5a3q
```


#### 请求参数

- `contract int` : 交易池合约地址

### 返回值
- `code int64 `  : 0或字段不存在为成功,其他均为失败
- `message string` : 响应信息
- `data object` : 成功返回数据
  - `token0 string`: pool_id
    - `contract string`: 合约地址
    - `pool_amount string`: 交易池币数
    - `symbol string`: 币种
  - `token1 int `: 池类型
    - `contract string`: 合约地址
    - `pool_amount string`: 交易池币数
    - `symbol string`: 币种


#### 返回示例
```json5
{
  "code": 0,
  "data": {
    "token0": {
      "pool_amount": "8341.747545",
      "symbol": "MTK"
    },
    "token1": {
      "pool_amount": "120.198645",
      "symbol": "WPLUG"
    }
  },
  "message": "ok"
}
```