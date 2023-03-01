---
order: 17
---

# SWAP交易池列表(PRC10)

### 请求方式
`GET`

### URL
`v1/swap/prc10/pools`

### 请求示例

```
v1/swap/prc10/pools
```


#### 请求参数

- 无

### 返回值
- `code int64 `  : 0或字段不存在为成功,其他均为失败
- `message string` : 响应信息
- `pools array` : 成功返回数据
  `id string`: pool_id
  `type_id int `: 池类型
  `reserve_coin_denoms array`: 交易池包含币种
  `reserve_account_address string`: 交易池收款地址
  `pool_coin_denom string`: 交易池币唯一名


#### 返回示例
```json5
{
  "pools": [
    {
      "id": "1",
      "type_id": 1,
      "reserve_coin_denoms": [
        "kitty",
        "uplugcn"
      ],
      "reserve_account_address": "gx1k30gpr8lgxgn6lluu86junfayaknyu74vendx9",
      "pool_coin_denom": "pool747FF80A34B22F1A5AEF8C2255023938B55781B627A4E14CB3F497A51AA415B7"
    },
    {
      "id": "2",
      "type_id": 1,
      "reserve_coin_denoms": [
        "dhw",
        "uplugcn"
      ],
      "reserve_account_address": "gx1kr00950kcg8asn4ls5cehmwyspe4zsl786g8lv",
      "pool_coin_denom": "pool81567FB090A1E8390E083B7D8B10ED459F27FC138D0AF114234BC209A78083B3"
    }
    //....
    //....
    //....
  ],
  "pagination": {
    "next_key": null,
    "total": "10"
  }
}
```