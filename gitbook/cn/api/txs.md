<!--
order: 4
-->
# 生成、签署和广播交易

本文档描述了如何生成（未签名）交易、对其进行签名（使用一个或多个密钥）并将其广播到网络。

### 服务端发起交易，以golang作为服务端语言开发:

在生成交易之前，需要创建一个“TxBuilder”的新实例。 由于 SDK 支持 Amino 和 Protobuf 事务，第一步是决定使用哪种编码方案。 所有后续步骤都保持不变，无论您使用的是 Amino 还是 Protobuf，因为“TxBuilder”抽象了编码机制。 在以下代码段中，我们将使用 Protobuf。

### 1. 初始化需要的 `txBuilder` 和 `codec`
```go
import (
	"github.com/oracleNetworkProtocol/plugchain/app"
)

func sendTx() error {
    // 选择您的编解码器：Amino 或 Protobuf
    encCfg := app.MakeTestEncodingConfig()

    // 创建一个新的 TxBuilder。
    txBuilder := encCfg.TxConfig.NewTxBuilder()

    // --剪断--
}
```

### 2. 根据自己操作的不同功能，需要准备不同的数据，以下以转账为例，需要准备：

* 发送方地址 `gx1s65azh0yj7n8yn4u0q449wt50eqr4qtyjzmhed`  节点上name是 `walletName`
* 发起者地址的私钥  可以通过 `plugchaind keys export walletName --unarmored-hex --unsafe` 得到
* 发起者地址的 `accountNumber` 和 `Sequence` 可以通过查询 `节点ip:1317/cosmos/auth/v1beta1/accounts/{address}`
* 接收者地址

```go
import (
    "encoding/hex"
	sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"    
)
    chainID := "plugchain"
    addr1, _ := sdk.AccAddressFromBech32("gx1s65azh0yj7n8yn4u0q449wt50eqr4qtyjzmhed")
	addr2, _ := sdk.AccAddressFromBech32("gx1d0ug2e7ehy6prw6msrtqwt55mydmxdsx4em9ds")
	priv := "55e2413b83e590944c6a4bcb443374c60bba847fc079788bd97ea455cb555bf0"
	privB, _ := hex.DecodeString(priv)
	priv1 := secp256k1.PrivKey{Key: privB}
	accountSeq := uint64(1)
	acountNumber := uint64(0)
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

```go
import (
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func sendTx() error {
    // --剪断--

    //发送一笔转账:
    //地址addr1 到 addr2
    //交易需要addr1 签名
    msg1 := banktypes.NewMsgSend(addr1, addr2, types.NewCoins(types.NewInt64Coin("line", 5000000)))

    err := txBuilder.SetMsgs(msg1, msg2)
    if err != nil {
        return err
    }

    txBuilder.SetGasLimit(200000)
    // txBuilder.SetFeeAmount(...)
    // txBuilder.SetMemo(...)
    // txBuilder.SetTimeoutHeight(...)
}
```

此时，`TxBuilder` 的底层交易已准备好进行签名。

### 3. 签署交易

我们选择我们的编码配置来使用 Protobuf，默认情况下它将使用`SIGN_MODE_DIRECT`。 根据 [ADR-020](https://github.com/cosmos/cosmos-sdk/blob/v0.40.0-rc6/docs/architecture/adr-020-protobuf-transaction-encoding.md)，每个签名者需要 签署所有其他签名者的`SignerInfo`。 这意味着我们需要依次执行两个步骤：

- 对于每个签名者，在`TxBuilder`中填充签名者的`SignerInfo`，
- 一旦所有`SignerInfo`s 被填充，对于每个签名者，签署`SignDoc`（要签名的有效负载）。

在当前的 `TxBuilder` 的 API 中，这两个步骤都使用相同的方法完成：`SetSignatures()`。 当前的 API 要求我们首先执行一轮 `SetSignatures()` _带有空签名_，仅填充 `SignerInfo`s，以及第二轮 `SetSignatures()` 以实际签署正确的有效负载。

```go
import (
    cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
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

`TxBuilder`现在已正确填充。 要打印它，您可以使用初始编码配置 `encCfg` 中的 `TxConfig` 接口：
```go
func sendTx() error {
    // --剪断--

    // 生成的 Protobuf 编码字节。
    txBytes, err := encCfg.TxConfig.TxEncoder()(txBuilder.GetTx())
    if err != nil {
        return err
    }
}
```

### 广播交易

广播交易的首选方式是使用 gRPC，但也可以使用 REST（通过 `gRPC-gateway`）或 Tendermint RPC。

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
    grpcConn := grpc.Dial(
        "127.0.0.1:9090", // 你的 gRPC 服务器地址。
        grpc.WithInsecure(), // SDK 不支持任何传输安全机制。
    )
    defer grpcConn.Close()

    // 通过 gRPC 广播 tx。 我们为 Protobuf Tx 服务创建了一个新客户端。
    txClient := tx.NewServiceClient(grpcConn)
    //然后我们在这个客户端上调用 BroadcastTx 方法。
    grpcRes, err := txClient.BroadcastTx(
        context.Background(),
        &tx.BroadcastTxRequest{
            Mode:    tx.BroadcastMode_BROADCAST_MODE_SYNC,
            TxBytes: txBytes,
        },
    )
    if err != nil {
        return err
    }

    fmt.Println(grpcRes.TxResponse.Code) // 如果 tx 成功，则应为 0

    return nil
}
```

## 使用REST API 广播交易

*不能使用 REST API 生成或签署交易，只能广播。*

使用 REST 端点（由 `gRPC-gateway` 提供服务）广播交易可以通过如下发送 POST 请求来完成，其中 `txBytes` 是签名交易的 protobuf 编码字节：

下文使用的`NewTxBytes` ,需要通过上面获取的 `txBytes` 执行以下操作
```golang
import (
    "fmt"
    "encoding/base64"
)

func sendTx() error {
    //--剪断--
    // base64 encode the encoded tx bytes
    NewTxBytes := base64.StdEncoding.EncodeToString(txBytes)
    fmt.Println(NewTxBytes)
}
```

然后通过下面方式广播：
```bash
curl -X POST \
    -H "Content-Type: application/json"
    -d'{"tx_bytes":"{{NewTxBytes}}","mode":"BROADCAST_MODE_SYNC"}'
    localhost:1317/cosmos/tx/v1beta1/txs
```

## 使用 CosmJS（JavaScript 和 TypeScript）操作

CosmJS 旨在用 JavaScript 构建可以嵌入 Web 应用程序的客户端库。 请参阅 [https://github.com/oracleNetworkProtocol/cosmjs](https://github.com/oracleNetworkProtocol/cosmjs) 了解更多信息。
