---
order: 22
---


# PRC20Token Details

### Request method
`GET`

### URL
`/server/v1/token/prc20/token`

### Request example

```
/server/v1/token/prc20/token?contract=gx1ay6eqcf0crujlkkvvzkrzd8e8fayh85yl6qgcl
```


#### Request parameters
- `contract  string Required`:Contract address

### Return value
- `code int64 `  : 0 or field does not exist is success, others are failures
- `message string` : Response information
- `data object` : Data returned successfully
    - `name`: Currency name
    - `symbol`: Currency symbol
    - `total_supply`: Total circulation
    - `decimal`: Decimals digit

#### Return to example
```json
{
  "code": 0,
  "data": {
    "name": "SofiOfficial",
    "symbol": "SOFI",
    "total_supply": "646318148423000",
    "decimal": "6"
  },
  "message": "ok"
}
```