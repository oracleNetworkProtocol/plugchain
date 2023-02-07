---
order: 5
---


# Gxswap 

相关合约信息：
```
Router:
bech32: gx1xulgpeuajthdc52eyqhfpsrf8w3thu97lhguxs
eip55: 0x373E80e79d92eEdC5159202e90c0693bA2Bbf0Be

Factory:
bech32: gx18g7wv6uq6p08mkupr8j2cze8hhhz5twu0ml2cz
eip55: 0x3A3Ce66B80D05E7dDB8119E4ac0b27Bdee2a2dDC

WPLUG 合约地址：
bech32: gx1d2wdkrvdu4y8l9k8pv0hs4cyrc03emtda8zepz
eip55: 0x6A9CdB0d8De5487F96C70b1f7857041e1f1CEd6d
```

## addLiquidityPLUG 方法
```js
addLiquidityPLUG(address token,uint256 amountTokenDesired,uint256 amountTokenMin,uint256 amountPLUGMin,address to,uint256 deadline)

```
功能描述：

创建流动池，如果存在池子，提供流动性。

添加流动性之前需要先进行approve操作. 

参数描述：
| 参数 | 类型 | 描述 |
| --- | ---| --- |
| token | address | 代币合约地址 |
| amountTokenDesired | uint256 | 注入池子token代币数量 |
| amountTokenMin | uint256 | 最低提出池子token代币数量，最低可填写0 |
| amountPLUGMin | uint256 | 最低提出池子`pc`数量，最低可填写0 |
| to | address | LP输出地址，一般为from地址 |
| deadline | uint256 | UTC时区的时间戳  必须大于当前时间 |


## swapExactPLUGForTokens 方法
```js
swapExactPLUGForTokens(uint256 amountOutMin,address[] path,address to,uint256 deadline)

```
功能描述：

pc兑换token

参数描述：
| 参数 | 类型 | 描述 |
| --- | ---| --- |
| amountOutMin | uint256 | 最低要兑换的token数量 |
| path | address[] | 交易对合约地址(0:wplug,1:token) |
| to | address |  token输出地址 |
| deadline | uint256 | UTC时区的时间戳  必须大于当前时间 |


## swapExactTokensForPLUG 方法
```js
swapExactTokensForPLUG(uint256 amountIn,uint256 amountOutMin,address[] path,address to,uint256 deadline)

```
功能描述：

token兑换pc

参数描述：
| 参数 | 类型 | 描述 |
| --- | ---| --- |
| amountIn | uint256 | 注入的token数量 |
| amountOutMin | uint256 | 最低要兑换的pc数量 |
| path | address[] | 交易对合约地址(0:token,1:wplug) |
| to | address |  pc输出地址 |
| deadline | uint256 | UTC时区的时间戳  必须大于当前时间 |


## removeLiquidityPLUG 方法
```js
removeLiquidityPLUG(address token,uint256 liquidity,uint256 amountTokenMin,uint256 amountPLUGMin,address to,uint256 deadline)

```
功能描述：

撤资
撤资之前需要先对lp合约进行approve操作. 

参数描述：
| 参数 | 类型 | 描述 |
| --- | ---| --- |
| token | address | 代币合约地址 |
| liquidity | uint256 | 撤资lp数量 |
| amountTokenMin | uint256 | 最低提出池子token代币数量，最低可填写0 |
| amountPLUGMin | uint256 | 最低提出池子`uplugcn`数量，最低可填写0 |
| to | address |  token和pc输出地址 |
| deadline | uint256 | UTC时区的时间戳  必须大于当前时间 |
 ## 完整代码

 使用golang实现
### [完整代码](https://github.com/oracleNetworkProtocol/gxswap)

 ```go
 package main

import (
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/oracleNetworkProtocol/gxswap/contracts/factory"
	"github.com/oracleNetworkProtocol/gxswap/contracts/router"
	"github.com/oracleNetworkProtocol/gxswap/contracts/token"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	//连接的8545服务
	blockchain, _ = ethclient.Dial("http://127.0.0.1:8545")

	//地址的私钥
	myKey, _ = crypto.HexToECDSA(strings.TrimPrefix("私钥", "0x"))
	auth, _  = bind.NewKeyedTransactorWithChainID(myKey, big.NewInt(520))

	//自己的地址
	myAddress = common.HexToAddress("EIP-55地址")

	//需要操作的合约
	myTokenAddress = common.HexToAddress("EIP-55合约地址")
	myToken, _     = token.NewToken(myTokenAddress, blockchain)

	//wPlug地址
	wPlugAddress = common.HexToAddress("0x6A9CdB0d8De5487F96C70b1f7857041e1f1CEd6d")

	//路由合约地址
	routerContractAddress = common.HexToAddress("0x373E80e79d92eEdC5159202e90c0693bA2Bbf0Be")
	routerContract, _     = contracts.NewCosmoswapRouter02(routerContractAddress, blockchain)

    //工厂合约地址
	factoryContractAddress = common.HexToAddress("0x3A3Ce66B80D05E7dDB8119E4ac0b27Bdee2a2dDC")
	factoryContract, _     = factory.NewICosmoswapFactory(factoryContractAddress, blockchain)
)

func main() {
    // removeLiquidityPLUG()
	// addLiquidityPLUG()
	// sell()
	// buy()
}

//addLiquidity 创建/注资token兑pc资金池
func addLiquidityPLUG() {
	//进行代币授权
	result, err := myToken.Approve(
		&bind.TransactOpts{
			From:   myAddress,
			Signer: auth.Signer,
			Value:  nil,
		},
		routerContractAddress,   //路由合约地址
		big.NewInt(60000000000), //授权数量
	)

	if err != nil {
		fmt.Println("Approve:", err)
		return
	}
	fmt.Println(result.Hash().Hex())

	//添加代币兑换uplugcn池
	trans, err := routerContract.AddLiquidityPLUG(
		&bind.TransactOpts{
			From:     myAddress,
			Signer:   auth.Signer,
			Value:    big.NewInt(60000000), //注入uplugcn数量
			GasPrice: big.NewInt(7),
		},
		myTokenAddress,          // mytoken 合约地址
		big.NewInt(60000000000), //注入代币数量
		big.NewInt(0),
		big.NewInt(0),
		myAddress,              //获得lp地址
		big.NewInt(time.Now().Unix()+1200), //超时时间
	)
	if err != nil {
		fmt.Println("add liquidity err:", err)
		return
	}
	fmt.Println(trans.Hash().Hex())
}

//removeLiquidityPLUG 撤资
func removeLiquidityPLUG() {
	lpTokenAddress := getPair()
	lpToken, _ := token.NewToken(lpTokenAddress, blockchain)
	//授权给lp合约操作权限
	lpToken.Approve(
		&bind.TransactOpts{
			From:   myAddress,
			Signer: auth.Signer,
			Value:  nil,
		},
		routerContractAddress, //获取lp地址
		big.NewInt(1000),      //授权数量
	)

	// lpBalances, err := lpToken.BalanceOf(
	// 	&bind.CallOpts{
	// 		From: myAddress,
	// 	},
	// 	myAddress,
	// )

	//如果mytoken的transfer方法有额外的销毁和转账操作需要使用routerContract.RemoveLiquidityPLUGSupportingFeeOnTransferTokens方法
	trans, err := routerContract.RemoveLiquidityPLUG(
		&bind.TransactOpts{
			From:     myAddress,
			Signer:   auth.Signer,
			Value:    nil,
			GasPrice: big.NewInt(7),
		},
		myTokenAddress,   // mytoken 合约地址
		big.NewInt(1000), //lp数量
		big.NewInt(0),    //最低撤资token数量
		big.NewInt(0),    //最低撤资pc数量
		myAddress,        //输出地址
		big.NewInt(time.Now().Add(30*time.Minute).Unix()), //超时时间
	)
	if err != nil {
		fmt.Println("remove liquidity err:", err)
		return
	}
	fmt.Println(trans.Hash().Hex())
}

 //获取lp地址
func getPair() common.Address {
	addr, err := factoryContract.GetPair(
		&bind.CallOpts{
			From: myAddress,
		},
		wPlugAddress,
		myTokenAddress,
	)
	if err != nil {
		log.Fatalf("getPair err: %v \n", err)
	}
	return addr
}


//buy pc兑换token
func buy() {
	trans, err := routerContract.SwapExactPLUGForTokens(
		&bind.TransactOpts{
			From:     myAddress,
			Signer:   auth.Signer,
			Value:    big.NewInt(10000000), //使用的uplugcn数量
			GasPrice: big.NewInt(7),
		},
		big.NewInt(100000000), //最低要兑换的token数量
		[]common.Address{   //交易对合约地址(0:wplug,1:token)
			wPlugAddress,   //Wrapped Plugcn (WPLUG) 合约地址
			myTokenAddress, //mytoken 合约地址
		},
		myAddress, //token输出地址
		big.NewInt(time.Now().Unix()+1200), //超时时间
	)
	if err != nil {
		log.Fatalf("TransferFrom err: %v \n", err)
	}
	fmt.Println("tx sent: ", trans.Hash().Hex())
}

//sell token兑换pc
func sell() {
	result, err := myToken.Approve(
		&bind.TransactOpts{
			From:   myAddress,
			Signer: auth.Signer,
			Value:  nil,
		},
		routerContractAddress,
		big.NewInt(9500000000),
	)

	if err != nil {
		fmt.Println("Approve:", err)
		return
	}
	fmt.Println(result.Hash().Hex())

	//如果mytoken的transfer方法有额外的销毁和转账操作需要使用routerContract.SwapExactTokensForPLUGSupportingFeeOnTransferTokens()
	trans, err := routerContract.SwapExactTokensForPLUG(
		&bind.TransactOpts{
			From:     myAddress,
			Signer:   auth.Signer,
			Value:    nil,
			GasPrice: big.NewInt(7),
		},
		big.NewInt(9500000000), //使用的token数量
		big.NewInt(1000000),    //最低需要兑换的uplugcn数量
		[]common.Address{
			myTokenAddress, //mytoken 合约地址
			wPlugAddress,   //Wrapped Plugcn (WPLUG) 合约地址
		},
		myAddress,
		big.NewInt(time.Now().Unix()+1200), //超时时间
	)
	if err != nil {
		fmt.Println("sell err:", err)
		return
	}
	fmt.Println(trans.Hash().Hex())
}
 ```