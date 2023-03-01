---
order: 5
---

# Address information

### Request method
`GET`

### URL
`/server/v1/account/account_info`

### Request example

```
/server/v1/account/account_info?addr=gx1sqyqeef3ym6q9x7anlgxmxh6u7sjfx6ccc3d4m
```


#### Request parameters
- `addr string Required` gx Address

### Return value
- `code int64 `  : 0 or field does not exist is success, others are failures
- `message string` : Response information
- `account object` : Data returned successfully
    - `@type` : Address type 
    - `account_number`:  Address serial number
    - `sequence` : Address transaction serial number, similar to the nonce of ETH


#### prc20 address return example
```json
{
    "account": {
        "@type": "/ethermint.types.v1.EthAccount",
        "base_account": {
            "address": "gx1sqyqeef3ym6q9x7anlgxmxh6u7sjfx6ccc3d4m",
            "pub_key": {
                "@type": "/ethermint.crypto.v1.ethsecp256k1.PubKey",
                "key": "AmKig8ab3m88TtimAIfrUu/neBxk1axFBdl231FuOqfa"
            },
            "account_number": "25237",
            "sequence": "4556"
        },
        "code_hash": "0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470"
    }
}
```


#### prc10 address return example
```json
{
  "account": {
    "@type": "/cosmos.auth.v1beta1.BaseAccount",
    "address": "gx14x8qtfg8u075dmhunvqymy0v2a5zrwzfg5y2fj",
    "pub_key": {
      "@type": "/cosmos.crypto.secp256k1.PubKey",
      "key": "AouLtAHVcdYNL1WPl0Tpb7wPZtsPYqwK5Ietr8jfsCnt"
    },
    "account_number": "8538",
    "sequence": "362"
  }
}
```