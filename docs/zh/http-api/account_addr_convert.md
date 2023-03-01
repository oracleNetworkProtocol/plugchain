---
order: 2
---

# 地址转换
 
- 地址之间互相转化 gx<==>0x<==>gxvloper

### 请求方式
`GET`

### URL
`/server/v1/account/addr_convert`

### 请求示例

```
/server/v1/account/addr_convert?addr=gx1sqyqeef3ym6q9x7anlgxmxh6u7sjfx6ccc3d4m
```


#### 请求参数
-  `addr string ` 必填 "gx地址 | gxvaloper验证者地址 | 0x地址"
### 返回值
- `data`
  - `Byte string`: 字节数据
  - `Hex string`: 16进制
  - `EIP55 string`:EIP55 0x地址,EVM地址
  - `BechAcc string`:gx用户地址
  - `BechVal string`:gxvaloper验证者地址

#### 返回示例
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