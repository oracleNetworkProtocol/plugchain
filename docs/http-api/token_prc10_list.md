---
order: 21
---

# PRC10Token List

### Request method
`GET`

### URL
`/server/v1/token/prc10/tokens`

### Request example

```
/server/v1/token/prc10/tokens
```


#### Request parameters
- Not

### Return value
- `Tokens array` : Data returned successfully
    - `object` : Query currency
      - `@type`: token type
      - `symbol`: Currency name
      - `name`: Currency introduction
      - `scale`: Decimal places
      - `min_unit`: Unique token name
      - `initial_supply`: Initial circulation
      - `max_supply`: Maximum circulation
      - `mintable`: Whether additional issuance is possible
      - `owner`: Publisher

#### Return to exampleReturn to example
```json
{
  "Tokens": [
    {
      "@type": "/plugchain.prc10.Token",
      "symbol": "aubrey",
      "name": "Aubrey Token",
      "scale": 6,
      "min_unit": "aubrey",
      "initial_supply": "1000000",
      "max_supply": "100000000",
      "mintable": true,
      "owner": "gx1cmpv09zh5xzm8wt69pgnklqdy6ysk63jq0n667"
    },
    {
      "@type": "/plugchain.prc10.Token",
      "symbol": "bhw1",
      "name": "DHW1",
      "scale": 6,
      "min_unit": "bhw1",
      "initial_supply": "3900000",
      "max_supply": "3900000",
      "mintable": false,
      "owner": "gx124g6keqt6etyuw2zhmzz78kz6halnmsvqztzmk"
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "2"
  }
}
```