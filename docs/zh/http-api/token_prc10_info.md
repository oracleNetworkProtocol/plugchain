---
order: 20
---

# PRC10代币详情

### 请求方式
`GET`

### URL
`/server/v1/token/prc10/token`

### 请求示例

```
/server/v1/token/prc10/token?denom=uplugcn
```


#### 请求参数
- `denom  string 必填`:币种

### 返回值
- `code int64 `  : 0或字段不存在为成功,其他均为失败
- `message string` : 响应信息
- `Token object` : 成功返回数据
  - `@type`: token 类型
  - `symbol`: 币种名
  - `name`: 币种介绍
  - `scale`: 小数位
  - `min_unit`: 唯一代币名称
  - `initial_supply`: 初始发行量
  - `max_supply`: 最大发行量
  - `mintable`: 是否能增发
  - `owner`: 发行者

#### 返回示例
```json
{
  "Token": {
    "@type": "/plugchain.prc10.Token",
    "symbol": "pc",
    "name": "plughub staking token",
    "scale": 6,
    "min_unit": "uplugcn",
    "initial_supply": "15989000000",
    "max_supply": "100000000000",
    "mintable": false,
    "owner": "gx1fjljkcf5f9ceh9cu54z7pp9wtmm586r2fm5gde"
  }
}
```