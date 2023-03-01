---
order: 6
---

# 生成新地址       

### 请求方式
`GET`    

### URL
`/server/v1/account/new_address`   


### 请求示例

- /server/v1/account/new_address 


### 返回值
- `code int64 `  : 0或字段不存在为成功,其他均为失败
- `message string` : 响应信息
- `data object`    : 返回数据
    - `address string ` : 返回地址
    - `mnemonic string  `: 助记词


#### 返回示例
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