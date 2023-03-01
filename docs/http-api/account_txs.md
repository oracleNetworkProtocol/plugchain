---
order: 7
---

# Account transactions

### Request method
`GET`

### URL
`/server/v1/account/txs`

### Request example

```
/server/v1/account/txs?addr=gx1sqyqeef3ym6q9x7anlgxmxh6u7sjfx6ccc3d4m
```


#### Request parameters
-  `code  string `Optional "0=All 1=Success 2=Failure"
-  `addr string ` Required "Address"
-  `coins  string `Optional "Currency"
-  `begin_time  string `Optional "Start Time"
-  `end_time  string `Optional "End Time"
-  `page  string `Optional "Current number of pages, minimum 1"

### Return value
- `code int64 `  : 0 or field does not exist is success, others are failures
- `message string` : Response information
- `data object` : Data returned successfully
    - `count int` : Total number
    - `txs array`: Transaction array
        - `tx object`: Transaction data
            - `height string` : Block height,
            - `hash string` : Transaction hash,
            - `module string` : Transaction module,
            - `action string` : Transaction module branch,
            - `type string` : Transaction arrival type,
            - `signer string` : Transaction signer,
            - `fee_amount string` : Transaction fee,
            - `fee_coin string` : Transaction fee currency,
            - `code string` : Response code,
            - `code_err string` : Error information,
            - `status string` : Transaction status,
            - `status_err string` : Status error information,
            - `coins string` : Currency included in the transaction,
            - `memo string` : Remarks,
            - `create_time string` : Creation time
        - `msg string` : Transaction json data


#### Return to example
```json
{
  "code": 0,
  "data": {
    "count": 4600,
    "txs": [
      {
        "tx": {
          "height": 8160857,
          "hash": "1214C80429CA33776ACF5F1AA1900EA9EFC20119CFB1C7E5AD22733E89FAA8A4",
          "module": "bank",
          "action": "MsgSend",
          "type": "Transfer",
          "signer": "gx1sqyqeef3ym6q9x7anlgxmxh6u7sjfx6ccc3d4m",
          "fee_amount": 200,
          "fee_coin": "uplugcn",
          "code": 0,
          "code_err": "",
          "status": "success",
          "status_err": "",
          "coins": "uplugcn",
          "memo": "",
          "create_time": "2023-02-10 16:04:44"
        },
        "msg": "{\"@type\":\"/cosmos.bank.v1beta1.MsgSend\",\"amount\":[{\"amount\":\"200000000\",\"denom\":\"uplugcn\"}],\"from_address\":\"gx1sqyqeef3ym6q9x7anlgxmxh6u7sjfx6ccc3d4m\",\"to_address\":\"gx1pu006rclnmkhs3y26v7euxstpr4ftgp5cfkctc\"}"
      }
    ]
  },
  "message": "ok"
}
```