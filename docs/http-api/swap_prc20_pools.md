---
order: 19
---


# SWAPTransaction pool list(PRC20)

### Request method
`GET`

### URL
`/server/v1/swap/prc20/pools`

### Request example

```
/server/v1/swap/prc20/pools
```


#### Request parameters

- Not

### Return value
- `code int64 `  : 0 or field does not exist is success, others are failures
- `message string` : Response information
- `data object` : Data returned successfully
    - `Contract string`: Contract address
    - `Token0 string`: TOKEN0 Currency type
    - `Token1 string`: TOKEN1 Currency type




#### Return to example
```json5
{
  "code": 0,
  "data": [
    {
      "Contract": "gx1rmf6rcdfcstl94mscdtgshjzy8q0d2ljjy5a3q",
      "Token0": "MTK",
      "Token1": "WPLUG"
    },
    {
      "Contract": "gx1uj5clzmywjmvldrkfuzap7dw3t6eek2npg5k8d",
      "Token0": "WPLUG",
      "Token1": "SOFI"
    },
    {
      "Contract": "gx1k9gerpcj8ghaz2lldlqtctkkd0078qs34fwsv7",
      "Token0": "ECO",
      "Token1": "PPPT"
    },
    //....
    //....
    //....
  ],
  "message": "ok"
}
```