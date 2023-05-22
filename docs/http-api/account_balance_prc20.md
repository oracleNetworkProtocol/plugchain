---
order: 4
---

# PRC20Token balance

### Request method
`GET`

### URL
`/server/v1/account/balance/prc20`

### Request example

```
/server/v1/account/balance/prc20?addr=gx1sqyqeef3ym6q9x7anlgxmxh6u7sjfx6ccc3d4m&contract=gx1ay6eqcf0crujlkkvvzkrzd8e8fayh85yl6qgcl
```


#### Request parameters
- `addr  string Required` :Address
- `contract  string Required`: Contract address

### Return value
- `code int64 `  : 0 or field does not exist is success, others are failures
- `message string` : Response information
- `data float64` :  Balance of


#### Return to example
```json
{
  "code": 0,
  "data": 1999,
  "message": "ok"
}
```