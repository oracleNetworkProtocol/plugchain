---
order: 17
---

# SWAPTransaction pool list(PRC10)

### Request method
`GET`

### URL
`v1/swap/prc10/pools`

### Request example

```
v1/swap/prc10/pools
```


#### Request parameters

- Not

### Return value
- `code int64 `  : 0 or field does not exist is success, others are failures
- `message string` : Response information
- `pools array` : Data returned successfully
  `id string`: pool_id
  `type_id int `: Pool type
  `reserve_coin_denoms array`: Currency included in the transaction pool
  `reserve_account_address string`: Receiving address of transaction pool
  `pool_coin_denom string`: Unique name of transaction pool currency


#### Return to example
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