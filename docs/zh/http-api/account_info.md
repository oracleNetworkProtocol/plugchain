---
order: 5
---

# 地址信息

### 请求方式
`GET`

### URL
`/server/v1/account/account_info`

### 请求示例

```
/server/v1/account/account_info?addr=gx1sqyqeef3ym6q9x7anlgxmxh6u7sjfx6ccc3d4m
```


#### 请求参数
- `addr string 必填` gx地址

### 返回值
- `code int64 `  : 0或字段不存在为成功,其他均为失败
- `message string` : 响应信息
- `account object` : 成功返回数据
    - `@type` : 地址类型 
    - `account_number`:  地址序号
    - `sequence` : 地址交易序号,类似ETH的nonce


#### prc20地址返回示例
```json
{
    "account": {
        "@type": "/ethermint.types.v1.EthAccount",
        "base_account": {
            "address": "gx1sqyqeef3ym6q9x7anlgxmxh6u7sjfx6ccc3d4m",
            "pub_key": {
                "@type": "/ethermint.crypto.v1.ethsecp256k1.PubKey",
                "key": "AmKig8ab3m88TtimAIfrUu/neBxk1axFBdl231FuOqfa"
            },
            "account_number": "25237",
            "sequence": "4556"
        },
        "code_hash": "0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470"
    }
}
```


#### prc10地址返回示例
```json
{
  "account": {
    "@type": "/cosmos.auth.v1beta1.BaseAccount",
    "address": "gx14x8qtfg8u075dmhunvqymy0v2a5zrwzfg5y2fj",
    "pub_key": {
      "@type": "/cosmos.crypto.secp256k1.PubKey",
      "key": "AouLtAHVcdYNL1WPl0Tpb7wPZtsPYqwK5Ietr8jfsCnt"
    },
    "account_number": "8538",
    "sequence": "362"
  }
}
```