---
order: 18
---

# SWAPTransaction pool(PRC20)

- swapTransaction pool details(prc20)

### Request method
`GET`

### URL
`/server/v1/swap/prc20/pool`

### Request example

```
/server/v1/swap/prc20/pool?contract=gx1rmf6rcdfcstl94mscdtgshjzy8q0d2ljjy5a3q
```


#### Request parameters

- `contract int` : Transaction pool contract address

### Return value
- `code int64 `  : 0 or field does not exist is success, others are failures
- `message string` : Response information
- `data object` : Data returned successfully
  - `token0 string`: pool_id
    - `contract string`: Contract address
    - `pool_amount string`: Number of transaction pool currencies
    - `symbol string`: Currency type
  - `token1 int `: Pool type
    - `contract string`: Contract address
    - `pool_amount string`: Number of transaction pool currencies
    - `symbol string`: Currency type


#### Return to example
```json5
{
  "code": 0,
  "data": {
    "token0": {
      "pool_amount": "8341.747545",
      "symbol": "MTK"
    },
    "token1": {
      "pool_amount": "120.198645",
      "symbol": "WPLUG"
    }
  },
  "message": "ok"
}
```