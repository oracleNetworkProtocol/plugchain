---
order: 4
---

# gRPC 

Plug Chain Hub v1.0.0（依赖Cosmos-SDK v0.42）引入了 Protobuf 作为主要的[编码](https://github.com/cosmos/cosmos-sdk/blob/master/docs/core/encoding.md)库，这带来了可插入 SDK 的各种基于 Protobuf 的工具。一种这样的工具是 [gRPC](https://grpc.io)，这是一种现代的开源高性能 RPC 框架，具有多语言客户端支持。

## gRPC 服务端口、激活方式和配置

`grpc.Server` 是一个具体的 gRPC 服务，它产生并服务任何gRPC请求。可以在 `~/.plugchain/config/app.toml` 中配置：

- `grpc.enable = true|false` 字段定义了 gRPC 服务是否可用，默认为 `true`。
- `grpc.address = {string}` 字段定义了服务绑定的地址（实际上是端口，因为主机必须保持为 `0.0.0.0`），默认为 `0.0.0.0:9000`。

gRPC 服务启动后，您可以使用 gRPC 客户端向其发送请求。

## gRPC 端点

Plug Chain Hub 附带的所有可用 gRPC 端点的概述见[Protobuf 文档](./proto-docs.md)。

## 构造、签名和广播交易

可以使用 Cosmos SDK 的 `TxBuilder` 接口，通过 Golang 以编程方式处理交易。

### 构造一个交易

在生成交易之前，需要创建一个新的 `TxBuilder` 实例。 由于 SDK 支持 Amino 和 Protobuf 交易，因此第一步将是确定要使用哪种编码方案。无论您使用的是 Amino 还是 Protobuf，所有后续步骤均保持不变，因为 `TxBuilder` 抽象了编码机制。在以下代码段中，我们将使用 Protobuf。

```go
import (
    "github.com/oracleNetworkProtocol/plugchain/app"
    sdk "github.com/cosmos/cosmos-sdk/types"
    plugchainapp "github.com/oracleNetworkProtocol/plugchain/app"
)

func sendTx() error {
    // 选择您的编解码器：Amino 或 Protobuf
    encCfg := ethencoding.MakeConfig(plugchainapp.ModuleBasics)

	config := sdk.GetConfig()
	plugchainapp.SetBech32Prefixes(config)

    // 创建一个新的 TxBuilder。
    txBuilder := encCfg.TxConfig.NewTxBuilder()

    // --剪断--
}
```

// TxBuilder 定义了一个应用程序定义的具体事务的接口
// 类型必须实现。 即，它必须能够设置消息，生成
// 签名，并提供规范字节进行签名。 交易必须
// 也知道如何编码自己。
```go
TxBuilder interface {
    GetTx() signing.Tx

    SetMsgs(msgs ...sdk.Msg) error
    SetSignatures(signatures ...signingtypes.SignatureV2) error
    SetMemo(memo string)
    SetFeeAmount(amount sdk.Coins)
    SetGasLimit(limit uint64)
    SetTimeoutHeight(height uint64)
}

```



根据自己操作的不同功能，需要准备不同的数据，以下以转账为例，需要准备：

* 发送方地址 `gx1s65azh0yj7n8yn4u0q449wt50eqr4qtyjzmhed`  节点上name是 `walletName`
* 发起者地址的私钥  可以通过 `plugchaind keys export walletName --unarmored-hex --unsafe` 得到
* 发起者地址的 `accountNumber` 和 `Sequence` 可以通过查询 `节点ip:1317/cosmos/auth/v1beta1/accounts/{address}`
* 接收者地址

```go
import (
    "encoding/hex"
	sdk "github.com/cosmos/cosmos-sdk/types"
    _ "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1" 
    "github.com/tharsis/ethermint/crypto/ethsecp256k1"   
)
    chainID := "plugchain_520-1"
    addr1, _ := sdk.AccAddressFromBech32("gx1s65azh0yj7n8yn4u0q449wt50eqr4qtyjzmhed")
	addr2, _ := sdk.AccAddressFromBech32("gx1d0ug2e7ehy6prw6msrtqwt55mydmxdsx4em9ds")
	addr3, _ := sdk.AccAddressFromBech32("gx1pq9yjvqwpmd5r6gpjs8cathhcljmymvp66sjjp")
    //发起者私钥
	priv := "55e2413b83e590944c6a4bcb443374c60bba847fc079788bd97ea455cb555bf0"
	privB, _ := hex.DecodeString(priv)
	// 如下查询地址的信息,使用account_number,sequence,需要根据 `@type` 类型 来锁定私钥类型，EthAcount 类型为`eth_secp256k1`,BaseAccount为`secp256k1`
	//curl -X GET "http://8.210.180.240:1317/cosmos/auth/v1beta1/accounts/gx13udxpqpmq6herxqk9yqa3agln8a0va9whjuqe7" -H  "accept: application/json"
	accountSeq := uint64(0)
	acountNumber := uint64(8)
	//EthAccount 类型， 使用包 "github.com/tharsis/ethermint/crypto/ethsecp256k1"
	//BaseAccount 类型 ，使用包 "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	priv1 := ethsecp256k1.PrivKey{Key: priva}
```


```go
import (
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func sendTx() error {
    // --剪断--

    //发送一笔转账:
    //地址addr1 到 addr2
    //地址addr1 到 addr3
    //交易需要 addr1 签名
    msg1 := banktypes.NewMsgSend(addr1, addr2, types.NewCoins(types.NewInt64Coin("uplugcn", 5000000)))
    msg2 := banktypes.NewMsgSend(addr1, addr3, types.NewCoins(types.NewInt64Coin("uplugcn", 4000000)))
    err := txBuilder.SetMsgs(msg1, msg2)
    if err != nil {
        return err
    }

    txBuilder.SetGasLimit(200000)
    txBuilder.SetFeeAmount(types.NewCoins(types.NewInt64Coin("uplugcn", 20)))
    txBuilder.SetMemo("give your my friend to LiLei")
    // txBuilder.SetTimeoutHeight(...)
}
```

至此，以 `TxBuilder` 为基础的交易已经准备好进行签名。

### 签名一个交易

我们将编码设置为使用 Protobuf，默认情况下将使用 `SIGN_MODE_DIRECT`。 根据[ADR-020](https://github.com/cosmos/cosmos-sdk/blob/v0.41.0/docs/architecture/adr-020-protobuf-transaction-encoding.md)，每个签名者都需要对所有其他签名者的 `SignerInfo`s 进行签名。这意味着我们需要依次执行两个步骤：

- 对于每个签名者，在 `TxBuilder` 中设置签名者的 `SignerInfo`，
- 设置所有 `SignerInfo` 之后，每个签名者对 `SignDoc`（要签名的有效数据）进行签名。

在当前的 `TxBuilder` API中，两个步骤都使用相同的方法 `SetSignatures()` 完成。当前的 API 要求我们首先循环执行带不带签名的 `SetSignatures()`，仅设置 `SignerInfo`s，然后进行第二轮 `SetSignatures()` 来对正确的有效数据进行签名。

```go
import (
    cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
    cliTx "github.com/cosmos/cosmos-sdk/client/tx"
)

func sendTx() error {
    // --剪断--
    
    //第一轮：我们收集所有签名者信息。 我们使用“设置空签名”技巧来做到这一点
    sign := signing.SignatureV2{
		PubKey: priv1.PubKey(),
		Data: &signing.SingleSignatureData{
			SignMode:  encCfg.TxConfig.SignModeHandler().DefaultMode(),
			Signature: nil,
		},

		Sequence: accountSeq,
	}

	err = txBuilder.SetSignatures(sign)
	if err != nil {
		panic(err)
	}


    //第二轮： 设置所有签名者信息，因此每个签名者都可以签名。
    sign = signing.SignatureV2{}
	signerD := xauthsigning.SignerData{
		ChainID:       chainID,
		AccountNumber: acountNumber,
		Sequence:      accountSeq,
	}
	sign, err = cliTx.SignWithPrivKey(
		encCfg.TxConfig.SignModeHandler().DefaultMode(), signerD,
		txBuilder, cryptotypes.PrivKey(&priv1), encCfg.TxConfig, accountSeq)

	if err != nil {
		panic(err)
	}
	err = txBuilder.SetSignatures(sign)
	if err != nil {
		panic(err)
	}
}
```


现在已经正确配置了 `TxBuilder`。 要打印它，您可以使用初始编码配置 `encCfg` 中的 `TxConfig` 接口：

```go
func sendTx() error {
    // --剪断--

    // 生成的 Protobuf 编码字节。
    txBytes, err := encCfg.TxConfig.TxEncoder()(txBuilder.GetTx())
    if err != nil {
        return err
    }

    // Generate a JSON string.
    // txJSONBytes, err := encCfg.TxConfig.TxJSONEncoder()(txBuilder.GetTx())
    // if err != nil {
    //     return err
    // }
    // txJSON := string(txJSONBytes)
}
```

### 广播一个交易

广播交易的首选方法是使用 gRPC，尽管也可以使用 REST（通过 `gRPC-gateway`）或 Tendermint RPC。 本教程中，我们仅介绍 gRPC 方法。

```go
import (
    "context"
    "fmt"

	"google.golang.org/grpc"

	"github.com/cosmos/cosmos-sdk/types/tx"
)

func sendTx() error {
    // --剪断--

    // 创建一个grpc服务
    grpcConn,ger := grpc.Dial(
        "127.0.0.1:9090", // 你的 gRPC 服务器地址。
        grpc.WithInsecure(), // SDK 不支持任何传输安全机制。
    )
    if ger != nil {
		panic(ger)
	}

    defer grpcConn.Close()

    // 通过 gRPC 广播 tx。 我们为 Protobuf Tx 服务创建了一个新客户端。
    txClient := tx.NewServiceClient(grpcConn)
    //然后我们在这个客户端上调用 BroadcastTx 方法。
    grpcRes, err := txClient.BroadcastTx(
        context.Background(),
        &tx.BroadcastTxRequest{
            Mode:    tx.BroadcastMode_BROADCAST_MODE_ASYNC,
            TxBytes: txBytes,
        },
    )
    if err != nil {
        return err
    }

    fmt.Println(grpcRes.GetTxResponse())

    return nil
}
```


## 完整代码
```go
package main

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	cliTx "github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	plugchainapp "github.com/oracleNetworkProtocol/plugchain/app"
	"github.com/tharsis/ethermint/crypto/ethsecp256k1"
	ethencoding "github.com/tharsis/ethermint/encoding"
	"google.golang.org/grpc"
)

func main() {
    var chainID = "plugchain_520-1"
    
	encCfg := ethencoding.MakeConfig(plugchainapp.ModuleBasics)
	config := sdk.GetConfig()
	plugchainapp.SetBech32Prefixes(config)

	txBuilder := encCfg.TxConfig.NewTxBuilder()

	addr1, _ := sdk.AccAddressFromBech32("gx13udxpqpmq6herxqk9yqa3agln8a0va9whjuqe7")
	addr2, _ := sdk.AccAddressFromBech32("gx1r4ffuq72pmny390lkk4l4x8wzamkcnru24wxaq")
	priv := "30116C43525488477CA78BB9C5653FB421334C7117070DFDD80240C895234505"
	priva, _ := hex.DecodeString(priv)
	// 如下查询地址的信息,使用account_number,sequence,需要根据 `@type` 类型 来锁定私钥类型，EthAcount 类型为`eth_secp256k1`,BaseAccount为`secp256k1`
	//curl -X GET "http://8.210.180.240:1317/cosmos/auth/v1beta1/accounts/gx13udxpqpmq6herxqk9yqa3agln8a0va9whjuqe7" -H  "accept: application/json"
	accountSeq := uint64(0)
	acountNumber := uint64(8)
	//EthAccount 类型， 使用包 "github.com/tharsis/ethermint/crypto/ethsecp256k1"
	//BaseAccount 类型 ，使用包 "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	priv1 := ethsecp256k1.PrivKey{Key: priva}

	msg1 := banktypes.NewMsgSend(addr1, addr2, sdk.NewCoins(sdk.NewInt64Coin("uplugcn", 111)))
	// msg2 := banktypes.NewMsgSend(addr2, addr1, sdk.NewCoins(sdk.NewInt64Coin("uplugcn", 222)))
	err := txBuilder.SetMsgs(msg1)
	if err != nil {
		panic(err)
	}
	txBuilder.SetGasLimit(200000)
	txBuilder.SetFeeAmount(sdk.NewCoins(sdk.NewInt64Coin("uplugcn", 20)))
	txBuilder.SetMemo("test send")
	// txBuilder.SetTimeoutHeight(61111800)

	sign := signing.SignatureV2{
		PubKey: priv1.PubKey(),
		Data: &signing.SingleSignatureData{
			SignMode:  encCfg.TxConfig.SignModeHandler().DefaultMode(),
			Signature: nil,
		},

		Sequence: accountSeq,
	}
	err = txBuilder.SetSignatures(sign)
	if err != nil {
		panic(err)
	}

	sign = signing.SignatureV2{}
	signerD := xauthsigning.SignerData{
		ChainID:       chainID,
		AccountNumber: acountNumber,
		Sequence:      accountSeq,
	}
	sign, err = cliTx.SignWithPrivKey(
		encCfg.TxConfig.SignModeHandler().DefaultMode(), signerD,
		txBuilder, cryptotypes.PrivKey(&priv1), encCfg.TxConfig, accountSeq)
	if err != nil {
		panic(err)
	}
	err = txBuilder.SetSignatures(sign)
	if err != nil {
		panic(err)

	}

	txBytes, er := encCfg.TxConfig.TxEncoder()(txBuilder.GetTx())

	if er != nil {
		panic(err)

	}

	grpcConn, ger := grpc.Dial(
		"localhost:9090",
		grpc.WithInsecure(),
	)

	if ger != nil {
		panic(ger)

	}
	defer grpcConn.Close()
	// base64 encode the encoded tx bytes
	// txBytes1 := base64.StdEncoding.EncodeToString(txBytes)
	txClient := tx.NewServiceClient(grpcConn)

	grpcRes, gerr := txClient.BroadcastTx(
		context.Background(),
		&tx.BroadcastTxRequest{
			Mode:    tx.BroadcastMode_BROADCAST_MODE_ASYNC,
			TxBytes: txBytes,
		},
	)
	log.Println("broadcast-tx::::", gerr, grpcRes)
	if gerr != nil {
		panic(err)
	}
	fmt.Println(grpcRes.GetTxResponse())

}
```