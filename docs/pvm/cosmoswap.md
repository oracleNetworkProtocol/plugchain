---
order: 5
---


# Cosmoswap

Related contract information:
```
cosmoSwap pool contract address:
bech32: gx1xulgpeuajthdc52eyqhfpsrf8w3thu97lhguxs
eip55: 0x373E80e79d92eEdC5159202e90c0693bA2Bbf0Be

WPLUG contract address:
bech32: gx1d2wdkrvdu4y8l9k8pv0hs4cyrc03emtda8zepz
eip55: 0x6A9CdB0d8De5487F96C70b1f7857041e1f1CEd6d
```

## addLiquidityPLUG method
```js
addLiquidityPLUG(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountPLUGMin, address to, uint256 deadline)

```
Function description:

Create a liquidity pool, and provide liquidity if a pool exists.

Approve operation is required before adding liquidity.

Parameter Description:
| Parameters | Type | Description |
| --- | ---| --- |
| token | address | Token contract address |
| amountTokenDesired | uint256 | Amount of tokens injected into the pool |
| amountTokenMin | uint256 | The minimum number of tokens to withdraw from the pool, the minimum value can be 0 |
| amountPLUGMin | uint256 | The minimum amount of `pc` in the pool, the minimum can be filled in 0 |
| to | address | LP output address, usually from address |
| deadline | uint256 | UTC time zone timestamp must be greater than current time |


## swapExactPLUGForTokens method
```js
swapExactPLUGForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline)

```
Function description:

plug exchange token

Parameter Description:
| Parameters | Type | Description |
| --- | ---| --- |
| amountOutMin | uint256 | Minimum amount of tokens to be exchanged |
| path | address[] | trading pair contract address (0:wplug,1:token) |
| to | address | token output address |
| deadline | uint256 | UTC time zone timestamp must be greater than current time |


## swapExactTokensForPLUG method
```js
swapExactTokensForPLUG(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline)

```
Function description:

token exchange plug

Parameter Description:
| Parameters | Type | Description |
| --- | ---| --- |
| amountIn | uint256 | Amount of tokens injected |
| amountOutMin | uint256 | Minimum amount of plugs to be redeemed |
| path | address[] | trading pair contract address (0:token,1:wplug) |
| to | address | plug output address |
| deadline | uint256 | UTC time zone timestamp must be greater than current time |


## removeLiquidityPLUG method
```js
removeLiquidityPLUG(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountPLUGMin, address to, uint256 deadline)

```
Function description:

Divestment
Before withdrawing funds, you need to perform an approve operation on the lp contract.

Parameter Description:
| Parameters | Type | Description |
| --- | ---| --- |
| token | address | Token contract address |
| liquidity | uint256 | Divestment lp number |
| amountTokenMin | uint256 | The minimum number of tokens to withdraw from the pool, the minimum value can be 0 |
| amountPLUGMin | uint256 | The minimum amount of pool `uplugcn` to be proposed, the minimum can be filled in 0 |
| to | address | token and plug output address |
| deadline | uint256 | UTC time zone timestamp must be greater than current time |
## Full code

 Implemented using golang
### [full code](https://github.com/oracleNetworkProtocol/cosmoswap)

 ```go
 package main

import (
"fmt"
"log"
"math/big"
"strings"

"github.com/oracleNetworkProtocol/cosmoswap/contracts/factory"
"github.com/oracleNetworkProtocol/cosmoswap/contracts/router"
"github.com/oracleNetworkProtocol/cosmoswap/contracts/token"

"github.com/ethereum/go-ethereum/accounts/abi/bind"
"github.com/ethereum/go-ethereum/common"
"github.com/ethereum/go-ethereum/crypto"
"github.com/ethereum/go-ethereum/ethclient"
)

var (
//Connected 8545 service
blockchain, _ = ethclient.Dial("http://127.0.0.1:8545")

//The private key of the address
myKey, _ = crypto.HexToECDSA(strings.TrimPrefix("Private key", "0x"))
auth, _ = bind.NewKeyedTransactorWithChainID(myKey, big.NewInt(520))

// own address
myAddress = common.HexToAddress("EIP-55 address")

//The contract that needs to be operated
myTokenAddress = common.HexToAddress("EIP-55 contract address")
myToken, _ = token.NewToken(myTokenAddress, blockchain)

//wPlug address
wPlugAddress = common.HexToAddress("0x6A9CdB0d8De5487F96C70b1f7857041e1f1CEd6d")

//Route contract address
routerContractAddress = common.HexToAddress("0x373E80e79d92eEdC5159202e90c0693bA2Bbf0Be")
routerContract, _ = contracts.NewCosmoswapRouter02(routerContractAddress, blockchain)

    //Factory contract address
factoryContractAddress = common.HexToAddress("0x3A3Ce66B80D05E7dDB8119E4ac0b27Bdee2a2dDC")
factoryContract, _ = factory.NewICosmoswapFactory(factoryContractAddress, blockchain)
)

func main() {
    // removeLiquidityPLUG()
// addLiquidityPLUG()
// sell()
// buy()
}

//addLiquidity creates/funds token exchange plug pool
func addLiquidityPLUG() {
    //do token authorization
    result, err := myToken.Approve(
        &bind.TransactOpts{
        From: myAddress,
        Signer: auth.Signer,
        Value: nil,
    },
    routerContractAddress, //Route contract address
    big.NewInt(60000000000), //authorized quantity
)

if err != nil {
    fmt.Println("Approve:", err)
    return
}
fmt.Println(result.Hash().Hex())

//Add token exchange uplugcn pool
trans, err := routerContract.AddLiquidityPLUG(
    &bind.TransactOpts{
        From: myAddress,
        Signer: auth.Signer,
        Value: big.NewInt(60000000), //Inject the number of uplugcn
        GasPrice: big.NewInt(7),
    },
    myTokenAddress, // mytoken contract address
    big.NewInt(60000000000), //Number of injected tokens
    big.NewInt(0),
    big.NewInt(0),
    myAddress, //get lp address
    big.NewInt(time.Now().Unix()+1200), //timeout
)
if err != nil {
    fmt.Println("add liquidity err:", err)
    return
}
fmt.Println(trans.Hash().Hex())
}

//removeLiquidityPLUG withdrawal
func removeLiquidityPLUG() {
    lpTokenAddress := getPair()
    lpToken, _ := token.NewToken(lpTokenAddress, blockchain)
    // Authorize the lp contract operation permission
    lpToken.Approve(
        &bind.TransactOpts{
        From: myAddress,
        Signer: auth.Signer,
        Value: nil,
        },
        routerContractAddress, //Get the lp address
        big.NewInt(1000), //authorized quantity
    )

    // lpBalances, err := lpToken.BalanceOf(
    // &bind.CallOpts{
    // From: myAddress,
    // },
    // myAddress,
    // )

//If the transfer method of mytoken has additional destruction and transfer operations, routerContract.RemoveLiquidityPLUGSupportingFeeOnTransferTokens method needs to be used
trans, err := routerContract.RemoveLiquidityPLUG(
    &bind.TransactOpts{
        From: myAddress,
        Signer: auth.Signer,
        Value: nil,
        GasPrice: big.NewInt(7),
    },
    myTokenAddress, // mytoken contract address
    big.NewInt(1000), //lp quantity
    big.NewInt(0), //Minimum withdrawal token amount
    big.NewInt(0), //Minimum withdrawal plug quantity
    myAddress, //output address
    big.NewInt(time.Now().Add(30*time.Minute).Unix()), //timeout
)
if err != nil {
    fmt.Println("remove liquidity err:", err)
    return
}
fmt.Println(trans.Hash().Hex())
}

 //Get lp address
func getPair() common.Address {
    addr, err := factoryContract.GetPair(
    &bind.CallOpts{
        From: myAddress,
    },
    wPlugAddress,
    myTokenAddress,
)
if err != nil {
    log. Fatalf("getPair err: %v \n", err)
}
    return addr
}


//buy plug exchange token
func buy() {
    trans, err := routerContract.SwapExactPLUGForTokens(
        &bind.TransactOpts{
            From: myAddress,
            Signer: auth.Signer,
            Value: big.NewInt(10000000), //Number of uplugcn used
            GasPrice: big.NewInt(7),
        },
        big.NewInt(100000000), //The minimum number of tokens to be exchanged
        []common.Address{ //Transaction pair contract address (0:wplug,1:token)
        wPlugAddress, //Wrapped Plugcn (WPLUG) contract address
        myTokenAddress, //mytoken contract address
        },
        myAddress, //token output address
        big.NewInt(time.Now().Unix()+1200), //timeout
    )
if err != nil {
    log. Fatalf("TransferFrom err: %v \n", err)
}
    fmt.Println("tx sent: ", trans.Hash().Hex())
}

//sell token exchange plug
func sell() {
    result, err := myToken.Approve(
    &bind.TransactOpts{
        From: myAddress,
        Signer: auth.Signer,
        Value: nil,
    },
    routerContractAddress,
    big.NewInt(9500000000),
)

if err != nil {
    fmt.Println("Approve:", err)
    return
}
fmt.Println(result.Hash().Hex())

//If the transfer method of mytoken has additional destruction and transfer operations, you need to use routerContract.SwapExactTokensForPLUGSupportingFeeOnTransferTokens()
trans, err := routerContract.SwapExactTokensForPLUG(
    &bind.TransactOpts{
        From: myAddress,
        Signer: auth.Signer,
        Value: nil,
        GasPrice: big.NewInt(7),
    },
    big.NewInt(9500000000), //Number of tokens used
    big.NewInt(1000000), //minimum amount of uplugcn to be exchanged
    []common.Address{
        myTokenAddress, //mytoken contract address
        wPlugAddress, //Wrapped Plugcn (WPLUG) contract address
    },
    myAddress,
    big.NewInt(time.Now().Unix()+1200), //timeout
)
if err != nil {
    fmt.Println("sell err:", err)
    return
}
fmt.Println(trans.Hash().Hex())
}

```