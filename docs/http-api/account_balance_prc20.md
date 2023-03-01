---
order: 4
---

# PRC20Token balance

### Request method
`GET`

### URL
`/server/v1/account/balance/prc10`

### Request example

```
/server/v1/account/balance/prc10?addr=gx1sqyqeef3ym6q9x7anlgxmxh6u7sjfx6ccc3d4m&denom=uplugcn
```


#### Request parameters
- `addr  string Required` :Address
- `denom  string Required`:Currency, please fill in min_ unit in the token information to query, such as min_ unit of PC is uplugcn

### Return value
- `code int64 `  : 0 or field does not exist is success, others are failures
- `message string` : Response information
- `balances object` : Data returned successfully
    - `denom` : Query currency
    - `amount`: Balance, which is normal only after dividing decimal places


#### Return to example
```json
{
  "balances": [
    {
      "denom": "uplugcn",
      "amount": "22919755"
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "1"
  }
}
```