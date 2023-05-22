---
order: 3
---

# PRC10代币余额

### 请求方式
`GET`

### URL
`/server/v1/account/balance/prc10`

### 请求示例

```
/server/v1/account/balance/prc10?addr=gx1sqyqeef3ym6q9x7anlgxmxh6u7sjfx6ccc3d4m&denom=uplugcn
```


#### 请求参数
- `addr  string 必填` :地址
- `denom  string 必填`:币种,请填写代币信息中的min_unit进行查询,例如pc的min_unit为uplugcn

### 返回值
- `code int64 `  : 0或字段不存在为成功,其他均为失败
- `message string` : 响应信息
- `balances object` : 成功返回数据
    - `denom` : 查询币种
    - `amount`: 余额,需要除小数位后才是正常数值

#### 返回示例
```json
{
  "balances": [
    {
      "denom": "uplugcn",
      "amount": "22919755"
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "1"
  }
}
```