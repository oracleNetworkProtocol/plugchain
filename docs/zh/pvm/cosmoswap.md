---
order: 5
---


# Cosmoswap 

相关合约信息：
```
cosmoSwap 池子合约地址:
bech32: gx1xulgpeuajthdc52eyqhfpsrf8w3thu97lhguxs
eip55: 0x373E80e79d92eEdC5159202e90c0693bA2Bbf0Be


WPLUG 合约地址：（对标主网plugcn， 类似于以太坊主网的 WETH，波场主网的 WTRX）
bech32: gx1d2wdkrvdu4y8l9k8pv0hs4cyrc03emtda8zepz
eip55: 0x6A9CdB0d8De5487F96C70b1f7857041e1f1CEd6d
```

## addLiquidityPLUG 方法
```js
addLiquidityPLUG(address token,uint amountTokenDesired,uint amountTokenMin,uint amountPLUGMin,address to,uint deadline)

```
功能描述：

创建流动池，如果存在池子，提供流动性。

添加流动性之前需要先进行approve操作. 

参数描述：
| 参数 | 类型 | 描述 |
| --- | ---| --- |
| token | address | 代币合约地址 |
| amountTokenDesired | uint | 注入池子token代币数量 |
| amountTokenMin | uint | 最低提出池子token代币数量，最低可填写0 |
| amountPLUGMin | uint | 最低提出池子`plugcn`数量，最低可填写0 |
| to | address | LP输出地址，一般为from地址 |
| deadline | uint | UTC时区的时间戳  必须大于当前时间 |
