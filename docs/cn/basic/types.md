

### 操作性功能参数

<a name="cosmos.bank.v1beta1.MsgSend"></a>

### MsgSend
表示将硬币从一个帐户发送到另一个帐户的消息。
- 参数

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `from_address` | [string](#string) |  |  |
| `to_address` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |

- 使用
```golang 
import (
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func sendTx() error {
    // --剪断--
    addr1, _ := sdk.AccAddressFromBech32("gx1s65azh0yj7n8yn4u0q449wt50eqr4qtyjzmhed")
	addr2, _ := sdk.AccAddressFromBech32("gx1d0ug2e7ehy6prw6msrtqwt55mydmxdsx4em9ds")

    msg1 := banktypes.NewMsgSend(addr1, addr2, types.NewCoins(types.NewInt64Coin("plug", 5000000)))

    err := txBuilder.SetMsgs(msg1, msg2)
    if err != nil {
        return err
    }
    // txBuilder.SetGasLimit(200000)
    // txBuilder.SetFeeAmount(...)
    // txBuilder.SetMemo(...)
    // txBuilder.SetTimeoutHeight(...)
}
```

<a name="cosmos.staking.v1beta1.MsgCreateValidator"></a>

### MsgCreateValidator
MsgCreateValidator 定义了用于创建新验证器的 SDK 消息。


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `description` | [Description](#cosmos.staking.v1beta1.Description) |  |  |
| `commission` | [CommissionRates](#cosmos.staking.v1beta1.CommissionRates) |  |  |
| `min_self_delegation` | [string](#string) |  |  |
| `delegator_address` | [string](#string) |  |  |
| `validator_address` | [string](#string) |  |  |
| `pubkey` | [google.protobuf.Any](#google.protobuf.Any) |  |  |
| `value` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |