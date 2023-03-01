---
order: 20
---


# PRC10Token Details

### Request method
`GET`

### URL
`/server/v1/token/prc10/token`

### Request example

```
/server/v1/token/prc10/token?denom=uplugcn
```


#### Request parameters
- `denom  string Required`:Currency type

### Return value
- `code int64 `  : 0 or field does not exist is success, others are failures
- `message string` : Response information
- `Token object` : Data returned successfully
  - `@type`: token type
  - `symbol`: Currency name
  - `name`: Currency introduction
  - `scale`: Decimal places
  - `min_unit`: Unique token name
  - `initial_supply`: Initial circulation
  - `max_supply`: Maximum circulation
  - `mintable`: Whether additional issuance is possible
  - `owner`: Publisher

#### Return to example
```json
{
  "Token": {
    "@type": "/plugchain.prc10.Token",
    "symbol": "pc",
    "name": "plughub staking token",
    "scale": 6,
    "min_unit": "uplugcn",
    "initial_supply": "15989000000",
    "max_supply": "100000000000",
    "mintable": false,
    "owner": "gx1fjljkcf5f9ceh9cu54z7pp9wtmm586r2fm5gde"
  }
}
```