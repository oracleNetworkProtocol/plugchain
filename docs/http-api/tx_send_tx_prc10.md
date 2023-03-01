---
order: 26
---

# Send transaction(PRC10)
 
- Need to provide mnemonics, please use with caution!
- Return the hash to the browser for query

### Request method
`POST`

### URL
`/server/v1/tx/send_tx_prc10`

### Request example

```
/server/v1/tx/send_tx_prc10
```


#### Request parameters
- `mnemonic string` : The mnemonic of the sending address
    - example:`clomp arain oblige fall switch pass import some wise color embrace foot language evil section steel purpose dynamic fork draw genre amazing squeeze hybrid`
- `to_addr string` : Receiving address
    - example: `gx1q4ymmvhnp276ls9j6gn6kszy27quf4j65q4pgf`
- `amount float64` : Send amount
    - example:`0.01`
- `contract string` : Currency name
    - example:`pc`

### Return value
- `code int64 `  : 0 or field does not exist is success, others are failures
- `message string` : Response information
- `hash object` : Transaction hash returned successfully

#### Return to example
```json
{
  "code": 0,
  "data": "89F333B7119721E72D105E626FF8BBF792B78E54A0195BB28A0F8586BD8FE1C7",
  "message": "ok"
}
```