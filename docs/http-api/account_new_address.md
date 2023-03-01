---
order: 6
---

# Generate new address       

### Request method
`GET`    

### URL
`/server/v1/account/new_address`   


### Request example
- /server/v1/account/new_address 
- 
### Return value
- `code int64 `  : 0 or field does not exist is success, others are failures
- `message string` : Response information
- `data object`    : Return to data
    - `address string ` : Return to address
    - `mnemonic string  `: Mnemonic words


#### Return to example
```json
{
  "data": {
      "address": "gx1zsvp5cfe4q2uve48yvsa4s4mr385utjqw3jefl",
      "mnemonic": "absurd coil front control brain kitten bunker beach remember limit hour else enact bone region divert biology balcony vacant regular motor prevent female kit"
      },
  "message": "响应完成",
  "status": 200
  }
```