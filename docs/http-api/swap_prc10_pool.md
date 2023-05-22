---
order: 16
---

# SWAPTransaction pool(PRC10)

- swapTransaction pool details(prc10) 

### Request method
`GET`

### URL
`v1/swap/prc10/pool`

### Request example

```
v1/swap/prc10/pool?pool_id=1
```



#### Request parameters

- `pool_id int` : Transaction pool ID

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
  "pool": {
    "id": "1",
    "type_id": 1,
    "reserve_coin_denoms": [
      "kitty",
      "uplugcn"
    ],
    "reserve_account_address": "gx1k30gpr8lgxgn6lluu86junfayaknyu74vendx9",
    "pool_coin_denom": "pool747FF80A34B22F1A5AEF8C2255023938B55781B627A4E14CB3F497A51AA415B7"
  }
}
```