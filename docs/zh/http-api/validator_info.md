---
order: 29
---

# 通过地址获取验证者信息

- 通过地址获取验证者

### 请求方式
`GET`

### URL
`/server/v1/validator/info`

### 请求示例

```
/server/v1/validator/info?addr=gxvaloper1rkecwy9pjsfd058w0pwa2perquc8xe638h4u5p
```


#### 请求参数

- 无

### 返回值
- `code int64 `: 0或字段不存在为成功,其他均为失败
- `message string` : 响应信息
- `validators array` : 验证者数组
    - `operator_address`: 验证者地址
    - `consensus_pubkey`: 共识公钥
    - `jailed`: 监禁状态
    - `status`: 验证者状态
    - `tokens`: token
    - `delegator_shares`: 质押股份
    - `description`: 验证者描述
    - `unbonding_height`: 锁定块高
    - `unbonding_time`: 锁定时间
    - `commission`: 共识信息
    - `min_self_delegation`: 最小质押量,单位uplugcn

#### 返回示例
```json5
{
  "validator": {
    "operator_address": "gxvaloper1rkecwy9pjsfd058w0pwa2perquc8xe638h4u5p",
    "consensus_pubkey": {
      "@type": "/cosmos.crypto.ed25519.PubKey",
      "key": "i4unhp2x1nry+iBrSBa4eag4e9rOU//MA4vcwHVJws0="
    },
    "jailed": false,
    "status": "BOND_STATUS_BONDED",
    "tokens": "80216413085275",
    "delegator_shares": "80296709794859.478689896810257297",
    "description": {
      "moniker": "Plug Storm Foundation-1",
      "identity": "",
      "website": "",
      "security_contact": "",
      "details": ""
    },
    "unbonding_height": "3288539",
    "unbonding_time": "2022-04-19T08:50:41.926091423Z",
    "commission": {
      "commission_rates": {
        "rate": "0.100000000000000000",
        "max_rate": "0.200000000000000000",
        "max_change_rate": "0.010000000000000000"
      },
      "update_time": "2021-08-19T08:56:38.706170801Z"
    },
    "min_self_delegation": "1000000"
  }
}
```