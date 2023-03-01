---
order: 22
---

# PRC20代币详情

### 请求方式
`GET`

### URL
`/server/v1/token/prc20/token`

### 请求示例

```
/server/v1/token/prc20/token?contract=gx1ay6eqcf0crujlkkvvzkrzd8e8fayh85yl6qgcl
```


#### 请求参数
- `contract  string 必填`:合约地址

### 返回值
- `code int64 `  : 0或字段不存在为成功,其他均为失败
- `message string` : 响应信息
- `data object` : 成功返回数据
    - `name`: 币种名
    - `symbol`: 币种符号
    - `total_supply`: 总发行量
    - `decimal`: 小数位数

#### 返回示例
```json
{
  "code": 0,
  "data": {
    "name": "SofiOfficial",
    "symbol": "SOFI",
    "total_supply": "646318148423000",
    "decimal": "6"
  },
  "message": "ok"
}
```