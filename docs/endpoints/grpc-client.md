---
order: 4
---

# gRPC

Plug Chain Hub v1.0.0 (depends on Cosmos-SDK v0.42) introduced Protobuf as the main [encoding](https://github.com/cosmos/cosmos-sdk/blob/master/docs/core/encoding.md) library, and this brings a wide range of Protobuf-based tools that can be plugged into the SDK. One such tool is [gRPC](https://grpc.io), a modern open source high performance RPC framework that has decent client support in several languages.

## gRPC Server Port, Activation and Configuration

The `grpc.Server` is a concrete gRPC server, which spawns and serves any gRPC requests. This server can be configured inside `~/.plugchain/config/app.toml`:

- `grpc.enable = true|false` field defines if the gRPC server should be enabled. Defaults to `true`.
- `grpc.address = {string}` field defines the address (really, the port, since the host should be kept at `0.0.0.0`) the server should bind to. Defaults to `0.0.0.0:9090`.
```
###############################################################################
###                           gRPC Configuration                            ###
###############################################################################

[grpc]

# Enable defines if the gRPC server should be enabled.
enable = true

# Address defines the gRPC server address to bind to.
address = "0.0.0.0:9090"
```
Once the gRPC server is started, you can send requests to it using a gRPC client.

## gRPC Endpoints

An overview of all available gRPC endpoints shipped with the Plug Chain Hub is [Protobuf documention](./proto-docs.md).

## Generating, Signing and Broadcasting Transactions

It is possible to manipulate transactions programmatically via Go using the Cosmos SDK's `TxBuilder` interface.

### Generating a Transaction

Before generating a transaction, a new instance of a `TxBuilder` needs to be created. Since the SDK supports both Amino and Protobuf transactions, the first step would be to decide which encoding scheme to use. All the subsequent steps remain unchanged, whether you're using Amino or Protobuf, as `TxBuilder` abstracts the encoding mechanisms. In the following snippet, we will use Protobuf.

### gomod add
```
	require (
		github.com/cosmos/cosmos-sdk v0.45.9
		github.com/oracleNetworkProtocol/plugchain v1.7.2	
	)

	replace (
	github.com/99designs/keyring => github.com/cosmos/keyring v1.1.7-0.20210622111912-ef00f8ac3d76
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	google.golang.org/grpc => google.golang.org/grpc v1.33.2
)
```

```go
import (
    sdk "github.com/cosmos/cosmos-sdk/types"
    plugchainapp "github.com/oracleNetworkProtocol/plugchain/app"
	ethencoding "github.com/evmos/ethermint/encoding"
)

func sendTx() error {
    // Choose your codec: Amino or Protobuf. Here, we use Protobuf, given by the following function.
    encCfg := ethencoding.MakeConfig(plugchainapp.ModuleBasics)

	config := sdk.GetConfig()
	plugchainapp.SetBech32Prefixes(config)


    // Create a new TxBuilder.
    txBuilder := encCfg.TxConfig.NewTxBuilder()

    // --snip--
}
```

// TxBuilder defines an application-defined interface for specific transactions
// The type must be implemented. That is, it must be able to set the message, generate
// Sign and provide canonical bytes for signing. Transaction must
// Also know how to code yourself.

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


According to the different functions you operate, you need to prepare different data. The following is an example of transfer, which needs to be prepared:

* The sender's address is `gx1s65azh0yj7n8yn4u0q449wt50eqr4qtyjzmhed` The name on the node is `walletName`
* The private key of the initiator address can be obtained by `plugchaind keys export walletName --unarmored-hex --unsafe`
* The `accountNumber` and `Sequence` of the initiator address can be queried by `node ip:1317/cosmos/auth/v1beta1/accounts/{address}`
* Recipient address

```go
import (
    "encoding/hex"
	sdk "github.com/cosmos/cosmos-sdk/types"
    _ "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1" 
    "github.com/evmos/ethermint/crypto/ethsecp256k1"    
)
    chainID := "plugchain_520-1"
    addr1, _ := sdk.AccAddressFromBech32("gx14z0773hnqq0qvwpr95stvn5tmtgvp8033e7aky")
	addr2, _ := sdk.AccAddressFromBech32("gx1d0ug2e7ehy6prw6msrtqwt55mydmxdsx4em9ds")
    //Initiator private key
	priv := "55e2413b83e590944c6a4bcb443374c60bba847fc079788bd97ea455cb555bf0"
	privB, _ := hex.DecodeString(priv)
	// To query the address information as follows, use account_number, sequence, you need to lock the private key type according to the `@type` type, the EthAcount type is `eth_secp256k1`, and the BaseAccount is `secp256k1`
	//curl -X GET "http://8.210.180.240:1317/cosmos/auth/v1beta1/accounts/gx13udxpqpmq6herxqk9yqa3agln8a0va9whjuqe7" -H  "accept: application/json"
	accountSeq := uint64(2)
	acountNumber := uint64(73991)
	//EthAccount type, using package "github.com/evmos/ethermint/crypto/ethsecp256k1"
	//BaseAccount type using package  "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	priv1 := ethsecp256k1.PrivKey{Key: priva}
```


```go
import (
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func sendTx() error {
    // --snip--

    // send
    //addr1 to addr2
    //addr1 to addr3
    //The transaction is signed by addr1
	msg1 := banktypes.NewMsgSend(addr1, addr2, types.NewCoins(types.NewInt64Coin("uplugcn", 1000000)))

	err := txBuilder.SetMsgs(msg1)
	if err != nil {
		return err
	}
	txBuilder.SetGasLimit(200000)
	txBuilder.SetFeeAmount(types.NewCoins(types.NewInt64Coin("uplugcn", 20)))
	txBuilder.SetMemo("give your my friend to LiLei")
    // txBuilder.SetTimeoutHeight(...)
}
```

At this point, `TxBuilder`'s underlying transaction is ready to be signed.

### Signing a Transaction

We set encoding config to use Protobuf, which will use `SIGN_MODE_DIRECT` by default. As per [ADR-020](https://github.com/cosmos/cosmos-sdk/blob/v0.41.0/docs/architecture/adr-020-protobuf-transaction-encoding.md), each signer needs to sign the `SignerInfo`s of all other signers. This means that we need to perform two steps sequentially:

- for each signer, populate the signer's `SignerInfo` inside `TxBuilder`,
- once all `SignerInfo`s are populated, for each signer, sign the `SignDoc` (the payload to be signed).

In the current `TxBuilder`'s API, both steps are done using the same method: `SetSignatures()`. The current API requires us to first perform a round of `SetSignatures()` _with empty signatures_, only to populate `SignerInfo`s, and a second round of `SetSignatures()` to actually sign the correct payload.

```go
import (
    cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
    cliTx "github.com/cosmos/cosmos-sdk/client/tx"
)

func sendTx() error {
    // --snip--
    
    //The first round: We collect all signer information. We use the "set empty signature" technique to do this
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
		return err
	}


    //Second round: Set all signer information, so every signer can sign.
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
		return err
	}
	err = txBuilder.SetSignatures(sign)
	if err != nil {
		return err
	}
}
```

The `TxBuilder` is now correctly populated. To print it, you can use the `TxConfig` interface from the initial encoding config `encCfg`:

```go
func sendTx() error {
    // --snip--

    // Generated Protobuf-encoded bytes.
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

### Broadcasting a Transaction

The preferred way to broadcast a transaction is to use gRPC, though using REST (via `gRPC-gateway`) or the Tendermint RPC is also posible. For this tutorial, we will only describe the gRPC method.

```go
import (
    "context"
    "fmt"

	"google.golang.org/grpc"

	"github.com/cosmos/cosmos-sdk/types/tx"
)

func sendTx(ctx context.Context) error {
    // --snip--

    // Create a connection to the gRPC server.
    grpcConn,err := grpc.Dial(
        "127.0.0.1:9090", // Or your gRPC server address.
        grpc.WithInsecure(), // The SDK doesn't support any transport security mechanism.
    )
    if err != nil {
		return err
	}
    defer grpcConn.Close()

    // Broadcast the tx via gRPC. We create a new client for the Protobuf Tx
    // service.
    txClient := tx.NewServiceClient(grpcConn)
    // We then call the BroadcastTx method on this client.
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

	fmt.Println(grpcRes.GetTxResponse())

	return nil
}
```

## 完整代码
```go
package main

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/cosmos/cosmos-sdk/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc"

	ethencoding "github.com/evmos/ethermint/encoding"
	plugchainapp "github.com/oracleNetworkProtocol/plugchain/app"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	_ "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/evmos/ethermint/crypto/ethsecp256k1"

	cliTx "github.com/cosmos/cosmos-sdk/client/tx"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
)

func main() {
	err := sendTx()
	fmt.Println(err)
}

func sendTx() error {

	// 选择您的编解码器：Amino 或 Protobuf
	encCfg := ethencoding.MakeConfig(plugchainapp.ModuleBasics)

	config := sdk.GetConfig()
	plugchainapp.SetBech32Prefixes(config)

	// 创建一个新的 TxBuilder。
	txBuilder := encCfg.TxConfig.NewTxBuilder()

	// --剪断--
	chainID := "plugchain_520-1"
	addr1, _ := sdk.AccAddressFromBech32("gx14z0773hnqq0qvwpr95stvn5tmtgvp8033e7aky")
	addr2, _ := sdk.AccAddressFromBech32("gx1d0ug2e7ehy6prw6msrtqwt55mydmxdsx4em9ds")
	//发起者私钥
	priv := "55e2413b83e590944c6a4bcb443374c60bba847fc079788bd97ea455cb555bf0"
	privB, _ := hex.DecodeString(priv)

	accountSeq := uint64(1)
	acountNumber := uint64(73991)
	//EthAccount 类型， 使用包 "github.com/evmos/ethermint/crypto/ethsecp256k1"
	//BaseAccount 类型 ，使用包 "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	priv1 := ethsecp256k1.PrivKey{Key: privB}

	msg1 := banktypes.NewMsgSend(addr1, addr2, types.NewCoins(types.NewInt64Coin("uplugcn", 1000000)))

	err := txBuilder.SetMsgs(msg1)
	if err != nil {
		return err
	}
	txBuilder.SetGasLimit(200000)
	txBuilder.SetFeeAmount(types.NewCoins(types.NewInt64Coin("uplugcn", 20)))
	txBuilder.SetMemo("give your my friend to LiLei")

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
		return err
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
		return err
	}
	err = txBuilder.SetSignatures(sign)
	if err != nil {
		return err
	}

	// 生成的 Protobuf 编码字节。
	txBytes, err := encCfg.TxConfig.TxEncoder()(txBuilder.GetTx())
	if err != nil {
		return err
	}

	// Generate a JSON string.
	txJSONBytes, err := encCfg.TxConfig.TxJSONEncoder()(txBuilder.GetTx())
	if err != nil {
		return err
	}
	txJSON := string(txJSONBytes)

	fmt.Println(txJSON)

	// 创建一个grpc服务
	grpcConn, err := grpc.Dial(
		"127.0.0.1:9090", // 你的 gRPC 服务器地址。
		grpc.WithInsecure(),   // SDK 不支持任何传输安全机制。
	)
	if err != nil {
		return err
	}

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

	fmt.Println(grpcRes.GetTxResponse())

	return nil
}
```