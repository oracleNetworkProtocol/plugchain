---
order: 4
---

# PRC20代币余额

### 请求方式
`GET`

### URL
`/server/v1/account/balance/prc20`

### 请求示例

```
/server/v1/account/balance/prc20?addr=gx1sqyqeef3ym6q9x7anlgxmxh6u7sjfx6ccc3d4m&contract=gx1ay6eqcf0crujlkkvvzkrzd8e8fayh85yl6qgcl
```


#### 请求参数
- `addr  string 必填` :地址
- `contract  string 必填`:合约地址

### 返回值
- `code int64 `  : 0或字段不存在为成功,其他均为失败
- `message string` : 响应信息
- `balances float64` : 余额

#### 返回示例
```json
{
  "code": 0,
  "data": 1999,
  "message": "ok"
}
```