---
order: 2
---

# Address translation
 
- Conversion between addresses gx<==>0x<==>gxvloper

### Request method
`GET`

### URL
`/server/v1/account/addr_convert`

### Request example

```
/server/v1/account/addr_convert?addr=gx1sqyqeef3ym6q9x7anlgxmxh6u7sjfx6ccc3d4m
```


#### Request parameters
 
- `addr string ` Required "gx Address | gxvaloper Address of verifier | 0x Address"

### Return value
- `data`
  - `Byte string`: Byte data
  - `Hex string`: Hexadecimal
  - `EIP55 string`:EIP55 0x address,EVM address
  - `BechAcc string`:gx user address
  - `BechVal string`:gxvaloper address of verifier


#### Return to examples
```json
{
  "code": 0,
  "data": {
    "Byte": "/2nU+5UgNbTtOehe4pn9p5iIAuw=",
    "Hex": "FF69D4FB952035B4ED39E85EE299FDA7988802EC",
    "EIP55": "0xFf69D4fb952035B4ed39E85eE299fDA7988802Ec",
    "BechAcc": "gx1la5af7u4yq6mfmfeap0w9x0a57vgsqhvqm9ywm",
    "BechVal": "gxvaloper1la5af7u4yq6mfmfeap0w9x0a57vgsqhvzl4jp2"
  },
  "message": "ok"
}
```