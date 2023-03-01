---
order: 27
---

# Send transaction(PVM)
 
- Need to provide mnemonics, please use with caution!
- Return the hash to the browser for query


### Request method
`POST`

### URL
`/server/v1/tx/send_tx_pvm`

### Request example

```
/server/v1/tx/send_tx_pvm
```


#### Request parameters
- `mnemonic string` : The mnemonic of the sending address 
  - example:`clomp arain oblige fall switch pass import some wise color embrace foot language evil section steel purpose dynamic fork draw genre amazing squeeze hybrid`
- `to_addr string` : Receiving address 
  - example: `gx1q4ymmvhnp276ls9j6gn6kszy27quf4j65q4pgf`
- `amount float64` : Send amount
  - example:`0.01`
- `contract string` : Currency contract address 
  - example:`gx1eahrjsw0hu2f3ey4d796rvp9074z7qx9mv9r58`

### Return value
- `code int64 `  : 0 or field does not exist is success, others are failures
- `message string` : Response information
- `hash object` : Transaction hash returned successfully

#### Return to example
```json
{
  "code": 0,
  "data": "0x8960c2a20e009dbb9bb0bc7c714b3be2e5ccb0f035629cc37f0170f9ee2abbdf",
  "message": "ok"
}
```