---
order: 21
---

# PRC10代币列表

### 请求方式
`GET`

### URL
`/server/v1/token/prc10/tokens`

### 请求示例

```
/server/v1/token/prc10/tokens
```


#### 请求参数
- 无

### 返回值
- `Tokens array` : 成功返回数据
    - `object` : 查询币种
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
  "Tokens": [
    {
      "@type": "/plugchain.prc10.Token",
      "symbol": "aubrey",
      "name": "Aubrey Token",
      "scale": 6,
      "min_unit": "aubrey",
      "initial_supply": "1000000",
      "max_supply": "100000000",
      "mintable": true,
      "owner": "gx1cmpv09zh5xzm8wt69pgnklqdy6ysk63jq0n667"
    },
    {
      "@type": "/plugchain.prc10.Token",
      "symbol": "bhw1",
      "name": "DHW1",
      "scale": 6,
      "min_unit": "bhw1",
      "initial_supply": "3900000",
      "max_supply": "3900000",
      "mintable": false,
      "owner": "gx124g6keqt6etyuw2zhmzz78kz6halnmsvqztzmk"
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "2"
  }
}
```