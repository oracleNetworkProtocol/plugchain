---
order: 27
---

# 发送交易(PVM)
 
- 需要提供助记词,请慎用!
- 返回hash可到浏览器查询


### 请求方式
`POST`

### URL
`/server/v1/tx/send_tx_pvm`

### 请求示例

```
/server/v1/tx/send_tx_pvm
```


#### 请求参数
- `mnemonic string` : 发送地址的助记词 
  - 例:`clomp arain oblige fall switch pass import some wise color embrace foot language evil section steel purpose dynamic fork draw genre amazing squeeze hybrid`
- `to_addr string` : 收款地址 
  - 例: `gx1q4ymmvhnp276ls9j6gn6kszy27quf4j65q4pgf`
- `amount float64` : 发送金额 
  - 例:`0.01`
- `contract string` : 币种合约地址 
  - 例:`gx1eahrjsw0hu2f3ey4d796rvp9074z7qx9mv9r58`

### 返回值
- `code int64 `  : 0或字段不存在为成功,其他均为失败
- `message string` : 响应信息
- `hash object` : 成功返回交易hash

#### 返回示例
```json
{
  "code": 0,
  "data": "0x8960c2a20e009dbb9bb0bc7c714b3be2e5ccb0f035629cc37f0170f9ee2abbdf",
  "message": "ok"
}
```