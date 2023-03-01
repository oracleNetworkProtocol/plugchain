---
order: 19
---

# SWAP交易池列表(PRC20)

### 请求方式
`GET`

### URL
`/server/v1/swap/prc20/pools`

### 请求示例

```
/server/v1/swap/prc20/pools
```


#### 请求参数

- 无

### 返回值
- `code int64 `  : 0或字段不存在为成功,其他均为失败
- `message string` : 响应信息
- `data object` : 成功返回数据
    - `Contract string`: 合约地址
    - `Token0 string`: TOKEN0 币种
    - `Token1 string`: TOKEN1 币种




#### 返回示例
```json5
{
  "code": 0,
  "data": [
    {
      "Contract": "gx1rmf6rcdfcstl94mscdtgshjzy8q0d2ljjy5a3q",
      "Token0": "MTK",
      "Token1": "WPLUG"
    },
    {
      "Contract": "gx1uj5clzmywjmvldrkfuzap7dw3t6eek2npg5k8d",
      "Token0": "WPLUG",
      "Token1": "SOFI"
    },
    {
      "Contract": "gx1k9gerpcj8ghaz2lldlqtctkkd0078qs34fwsv7",
      "Token0": "ECO",
      "Token1": "PPPT"
    },
    //....
    //....
    //....
  ],
  "message": "ok"
}
```