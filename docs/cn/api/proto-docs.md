<!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [cosmos/base/v1beta1/coin.proto](#cosmos/base/v1beta1/coin.proto)
    - [Coin](#cosmos.base.v1beta1.Coin)
    - [DecCoin](#cosmos.base.v1beta1.DecCoin)
  
- [cosmos/bank/v1beta1/bank.proto](#cosmos/bank/v1beta1/bank.proto)
    - [Input](#cosmos.bank.v1beta1.Input)
    - [Output](#cosmos.bank.v1beta1.Output)
  
  
- [cosmos/bank/v1beta1/tx.proto](#cosmos/bank/v1beta1/tx.proto)
    - [MsgMultiSend](#cosmos.bank.v1beta1.MsgMultiSend)
    - [MsgSend](#cosmos.bank.v1beta1.MsgSend)
  

- [cosmos/crypto/ed25519/keys.proto](#cosmos/crypto/ed25519/keys.proto)
    - [PrivKey](#cosmos.crypto.ed25519.PrivKey)
    - [PubKey](#cosmos.crypto.ed25519.PubKey)
  
- [cosmos/crypto/secp256k1/keys.proto](#cosmos/crypto/secp256k1/keys.proto)
    - [PrivKey](#cosmos.crypto.secp256k1.PrivKey)
    - [PubKey](#cosmos.crypto.secp256k1.PubKey)

- [cosmos/staking/v1beta1/tx.proto](#cosmos/staking/v1beta1/tx.proto)
    - [MsgBeginRedelegate](#cosmos.staking.v1beta1.MsgBeginRedelegate)
    - [MsgCreateValidator](#cosmos.staking.v1beta1.MsgCreateValidator)
    - [MsgDelegate](#cosmos.staking.v1beta1.MsgDelegate)
    - [MsgEditValidator](#cosmos.staking.v1beta1.MsgEditValidator)
    - [MsgUndelegate](#cosmos.staking.v1beta1.MsgUndelegate)
  
    - [Msg](#cosmos.staking.v1beta1.Msg)
  
- [cosmos/tx/v1beta1/tx.proto](#cosmos/tx/v1beta1/tx.proto)
    - [AuthInfo](#cosmos.tx.v1beta1.AuthInfo)
    - [Fee](#cosmos.tx.v1beta1.Fee)
    - [ModeInfo](#cosmos.tx.v1beta1.ModeInfo)
    - [SignerInfo](#cosmos.tx.v1beta1.SignerInfo)
    - [Tx](#cosmos.tx.v1beta1.Tx)
    - [TxBody](#cosmos.tx.v1beta1.TxBody)
  
- [cosmos/tx/v1beta1/service.proto](#cosmos/tx/v1beta1/service.proto)
    - [BroadcastTxRequest](#cosmos.tx.v1beta1.BroadcastTxRequest)
    - [BroadcastMode](#cosmos.tx.v1beta1.BroadcastMode)
    - [Service](#cosmos.tx.v1beta1.Service)


<a name="cosmos/base/v1beta1/coin.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/base/v1beta1/coin.proto


<a name="cosmos.base.v1beta1.Coin"></a>

### Coin
Coin defines a token with a denomination and an amount.

NOTE: The amount field is an Int which implements the custom method
signatures required by gogoproto.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `amount` | [string](#string) |  |  |






<a name="cosmos.base.v1beta1.DecCoin"></a>

### DecCoin
DecCoin defines a token with a denomination and a decimal amount.

NOTE: The amount field is an Dec which implements the custom method
signatures required by gogoproto.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `amount` | [string](#string) |  |  |




<a name="cosmos/bank/v1beta1/bank.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/bank/v1beta1/bank.proto


<a name="cosmos.bank.v1beta1.Input"></a>

### Input
Input models transaction input.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |
| `coins` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |

<a name="cosmos.bank.v1beta1.Output"></a>

### Output
Output models transaction outputs.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |
| `coins` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |



<a name="cosmos/base/query/v1beta1/pagination.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/base/query/v1beta1/pagination.proto



<a name="cosmos.base.query.v1beta1.PageRequest"></a>

### PageRequest
PageRequest is to be embedded in gRPC request messages for efficient
pagination. Ex:

 message SomeRequest {
         Foo some_parameter = 1;
         PageRequest pagination = 2;
 }


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [bytes](#bytes) |  | key is a value returned in PageResponse.next_key to begin querying the next page most efficiently. Only one of offset or key should be set. |
| `offset` | [uint64](#uint64) |  | offset is a numeric offset that can be used when key is unavailable. It is less efficient than using key. Only one of offset or key should be set. |
| `limit` | [uint64](#uint64) |  | limit is the total number of results to be returned in the result page. If left empty it will default to a value to be set by each app. |
| `count_total` | [bool](#bool) |  | count_total is set to true to indicate that the result set should include a count of the total number of items available for pagination in UIs. count_total is only respected when offset is used. It is ignored when key is set. |


<a name="cosmos.base.query.v1beta1.PageResponse"></a>

### PageResponse
PageResponse is to be embedded in gRPC response messages where the
corresponding request message has used PageRequest.

 message SomeResponse {
         repeated Bar results = 1;
         PageResponse page = 2;
 }


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `next_key` | [bytes](#bytes) |  | next_key is the key to be passed to PageRequest.key to query the next page most efficiently |
| `total` | [uint64](#uint64) |  | total is total number of results available if PageRequest.count_total was set, its value is undefined otherwise |



<a name="cosmos/bank/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/bank/v1beta1/tx.proto



<a name="cosmos.bank.v1beta1.MsgMultiSend"></a>

### MsgMultiSend
MsgMultiSend represents an arbitrary multi-in, multi-out send message.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `inputs` | [Input](#cosmos.bank.v1beta1.Input) | repeated |  |
| `outputs` | [Output](#cosmos.bank.v1beta1.Output) | repeated |  |


<a name="cosmos.bank.v1beta1.MsgSend"></a>

### MsgSend
MsgSend represents a message to send coins from one account to another.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `from_address` | [string](#string) |  |  |
| `to_address` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |


<a name="cosmos/crypto/ed25519/keys.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/crypto/ed25519/keys.proto



<a name="cosmos.crypto.ed25519.PrivKey"></a>

### PrivKey
PrivKey defines a ed25519 private key.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [bytes](#bytes) |  |  |






<a name="cosmos.crypto.ed25519.PubKey"></a>

### PubKey
PubKey defines a ed25519 public key
Key is the compressed form of the pubkey. The first byte depends is a 0x02 byte
if the y-coordinate is the lexicographically largest of the two associated with
the x-coordinate. Otherwise the first byte is a 0x03.
This prefix is followed with the x-coordinate.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [bytes](#bytes) |  |  |


<a name="cosmos/crypto/secp256k1/keys.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/crypto/secp256k1/keys.proto



<a name="cosmos.crypto.secp256k1.PrivKey"></a>

### PrivKey
PrivKey defines a secp256k1 private key.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [bytes](#bytes) |  |  |






<a name="cosmos.crypto.secp256k1.PubKey"></a>

### PubKey
PubKey defines a secp256k1 public key
Key is the compressed form of the pubkey. The first byte depends is a 0x02 byte
if the y-coordinate is the lexicographically largest of the two associated with
the x-coordinate. Otherwise the first byte is a 0x03.
This prefix is followed with the x-coordinate.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [bytes](#bytes) |  |  |


<a name="cosmos/staking/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/staking/v1beta1/tx.proto



<a name="cosmos.staking.v1beta1.MsgBeginRedelegate"></a>

### MsgBeginRedelegate
MsgBeginRedelegate defines a SDK message for performing a redelegation
of coins from a delegator and source validator to a destination validator.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator_address` | [string](#string) |  |  |
| `validator_src_address` | [string](#string) |  |  |
| `validator_dst_address` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |



<a name="cosmos.staking.v1beta1.MsgCreateValidator"></a>

### MsgCreateValidator
MsgCreateValidator defines a SDK message for creating a new validator.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `description` | [Description](#cosmos.staking.v1beta1.Description) |  |  |
| `commission` | [CommissionRates](#cosmos.staking.v1beta1.CommissionRates) |  |  |
| `min_self_delegation` | [string](#string) |  |  |
| `delegator_address` | [string](#string) |  |  |
| `validator_address` | [string](#string) |  |  |
| `pubkey` | [google.protobuf.Any](#google.protobuf.Any) |  |  |
| `value` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |




<a name="cosmos.staking.v1beta1.MsgDelegate"></a>

### MsgDelegate
MsgDelegate defines a SDK message for performing a delegation of coins
from a delegator to a validator.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator_address` | [string](#string) |  |  |
| `validator_address` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |




<a name="cosmos.staking.v1beta1.MsgEditValidator"></a>

### MsgEditValidator
MsgEditValidator defines a SDK message for editing an existing validator.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `description` | [Description](#cosmos.staking.v1beta1.Description) |  |  |
| `validator_address` | [string](#string) |  |  |
| `commission_rate` | [string](#string) |  | We pass a reference to the new commission rate and min self delegation as it's not mandatory to update. If not updated, the deserialized rate will be zero with no way to distinguish if an update was intended. REF: #2373 |
| `min_self_delegation` | [string](#string) |  |  |



<a name="cosmos.staking.v1beta1.MsgUndelegate"></a>

### MsgUndelegate
MsgUndelegate defines a SDK message for performing an undelegation from a
delegate and a validator.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator_address` | [string](#string) |  |  |
| `validator_address` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |


<a name="cosmos.staking.v1beta1.Msg"></a>

### Msg
Msg defines the staking Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `CreateValidator` | [MsgCreateValidator](#cosmos.staking.v1beta1.MsgCreateValidator) | [MsgCreateValidatorResponse](#cosmos.staking.v1beta1.MsgCreateValidatorResponse) | CreateValidator defines a method for creating a new validator. | |
| `EditValidator` | [MsgEditValidator](#cosmos.staking.v1beta1.MsgEditValidator) | [MsgEditValidatorResponse](#cosmos.staking.v1beta1.MsgEditValidatorResponse) | EditValidator defines a method for editing an existing validator. | |
| `Delegate` | [MsgDelegate](#cosmos.staking.v1beta1.MsgDelegate) | [MsgDelegateResponse](#cosmos.staking.v1beta1.MsgDelegateResponse) | Delegate defines a method for performing a delegation of coins from a delegator to a validator. | |
| `BeginRedelegate` | [MsgBeginRedelegate](#cosmos.staking.v1beta1.MsgBeginRedelegate) | [MsgBeginRedelegateResponse](#cosmos.staking.v1beta1.MsgBeginRedelegateResponse) | BeginRedelegate defines a method for performing a redelegation of coins from a delegator and source validator to a destination validator. | |
| `Undelegate` | [MsgUndelegate](#cosmos.staking.v1beta1.MsgUndelegate) | [MsgUndelegateResponse](#cosmos.staking.v1beta1.MsgUndelegateResponse) | Undelegate defines a method for performing an undelegation from a delegate and a validator. | |

 <!-- end services -->



<a name="cosmos/tx/signing/v1beta1/signing.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/tx/signing/v1beta1/signing.proto



<a name="cosmos.tx.signing.v1beta1.SignatureDescriptor"></a>

### SignatureDescriptor
SignatureDescriptor is a convenience type which represents the full data for
a signature including the public key of the signer, signing modes and the
signature itself. It is primarily used for coordinating signatures between
clients.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `public_key` | [google.protobuf.Any](#google.protobuf.Any) |  | public_key is the public key of the signer |
| `data` | [SignatureDescriptor.Data](#cosmos.tx.signing.v1beta1.SignatureDescriptor.Data) |  |  |
| `sequence` | [uint64](#uint64) |  | sequence is the sequence of the account, which describes the number of committed transactions signed by a given address. It is used to prevent replay attacks. |


<a name="cosmos.tx.signing.v1beta1.SignMode"></a>

### SignMode
SignMode represents a signing mode with its own security guarantees.

| Name | Number | Description |
| ---- | ------ | ----------- |
| SIGN_MODE_UNSPECIFIED | 0 | SIGN_MODE_UNSPECIFIED specifies an unknown signing mode and will be rejected |
| SIGN_MODE_DIRECT | 1 | SIGN_MODE_DIRECT specifies a signing mode which uses SignDoc and is verified with raw bytes from Tx |
| SIGN_MODE_TEXTUAL | 2 | SIGN_MODE_TEXTUAL is a future signing mode that will verify some human-readable textual representation on top of the binary representation from SIGN_MODE_DIRECT |
| SIGN_MODE_LEGACY_AMINO_JSON | 127 | SIGN_MODE_LEGACY_AMINO_JSON is a backwards compatibility mode which uses Amino JSON and will be removed in the future |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/tx/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/tx/v1beta1/tx.proto



<a name="cosmos.tx.v1beta1.AuthInfo"></a>

### AuthInfo
AuthInfo describes the fee and signer modes that are used to sign a
transaction.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `signer_infos` | [SignerInfo](#cosmos.tx.v1beta1.SignerInfo) | repeated | signer_infos defines the signing modes for the required signers. The number and order of elements must match the required signers from TxBody's messages. The first element is the primary signer and the one which pays the fee. |
| `fee` | [Fee](#cosmos.tx.v1beta1.Fee) |  | Fee is the fee and gas limit for the transaction. The first signer is the primary signer and the one which pays the fee. The fee can be calculated based on the cost of evaluating the body and doing signature verification of the signers. This can be estimated via simulation. |






<a name="cosmos.tx.v1beta1.Fee"></a>

### Fee
Fee includes the amount of coins paid in fees and the maximum
gas to be used by the transaction. The ratio yields an effective "gasprice",
which must be above some miminum to be accepted into the mempool.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | amount is the amount of coins to be paid as a fee |
| `gas_limit` | [uint64](#uint64) |  | gas_limit is the maximum gas that can be used in transaction processing before an out of gas error occurs |
| `payer` | [string](#string) |  | if unset, the first signer is responsible for paying the fees. If set, the specified account must pay the fees. the payer must be a tx signer (and thus have signed this field in AuthInfo). setting this field does *not* change the ordering of required signers for the transaction. |
| `granter` | [string](#string) |  | if set, the fee payer (either the first signer or the value of the payer field) requests that a fee grant be used to pay fees instead of the fee payer's own balance. If an appropriate fee grant does not exist or the chain does not support fee grants, this will fail |



<a name="cosmos.tx.v1beta1.SignerInfo"></a>

### SignerInfo
SignerInfo describes the public key and signing mode of a single top-level
signer.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `public_key` | [google.protobuf.Any](#google.protobuf.Any) |  | public_key is the public key of the signer. It is optional for accounts that already exist in state. If unset, the verifier can use the required \ signer address for this position and lookup the public key. |
| `mode_info` | [ModeInfo](#cosmos.tx.v1beta1.ModeInfo) |  | mode_info describes the signing mode of the signer and is a nested structure to support nested multisig pubkey's |
| `sequence` | [uint64](#uint64) |  | sequence is the sequence of the account, which describes the number of committed transactions signed by a given address. It is used to prevent replay attacks. |






<a name="cosmos.tx.v1beta1.Tx"></a>

### Tx
Tx is the standard type used for broadcasting transactions.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `body` | [TxBody](#cosmos.tx.v1beta1.TxBody) |  | body is the processable content of the transaction |
| `auth_info` | [AuthInfo](#cosmos.tx.v1beta1.AuthInfo) |  | auth_info is the authorization related content of the transaction, specifically signers, signer modes and fee |
| `signatures` | [bytes](#bytes) | repeated | signatures is a list of signatures that matches the length and order of AuthInfo's signer_infos to allow connecting signature meta information like public key and signing mode by position. |






<a name="cosmos.tx.v1beta1.TxBody"></a>

### TxBody
TxBody is the body of a transaction that all signers sign over.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `messages` | [google.protobuf.Any](#google.protobuf.Any) | repeated | messages is a list of messages to be executed. The required signers of those messages define the number and order of elements in AuthInfo's signer_infos and Tx's signatures. Each required signer address is added to the list only the first time it occurs. By convention, the first required signer (usually from the first message) is referred to as the primary signer and pays the fee for the whole transaction. |
| `memo` | [string](#string) |  | memo is any arbitrary memo to be added to the transaction |
| `timeout_height` | [uint64](#uint64) |  | timeout is the block height after which this transaction will not be processed by the chain |
| `extension_options` | [google.protobuf.Any](#google.protobuf.Any) | repeated | extension_options are arbitrary options that can be added by chains when the default options are not sufficient. If any of these are present and can't be handled, the transaction will be rejected |
| `non_critical_extension_options` | [google.protobuf.Any](#google.protobuf.Any) | repeated | extension_options are arbitrary options that can be added by chains when the default options are not sufficient. If any of these are present and can't be handled, they will be ignored |



<a name="cosmos/tx/v1beta1/service.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/tx/v1beta1/service.proto



<a name="cosmos.tx.v1beta1.BroadcastTxRequest"></a>

### BroadcastTxRequest
BroadcastTxRequest is the request type for the Service.BroadcastTxRequest
RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tx_bytes` | [bytes](#bytes) |  | tx_bytes is the raw transaction. |
| `mode` | [BroadcastMode](#cosmos.tx.v1beta1.BroadcastMode) |  |  |




<a name="cosmos.tx.v1beta1.BroadcastMode"></a>

### BroadcastMode
BroadcastMode specifies the broadcast mode for the TxService.Broadcast RPC method.

| Name | Number | Description |
| ---- | ------ | ----------- |
| BROADCAST_MODE_UNSPECIFIED | 0 | zero-value for mode ordering |
| BROADCAST_MODE_BLOCK | 1 | BROADCAST_MODE_BLOCK defines a tx broadcasting mode where the client waits for the tx to be committed in a block. |
| BROADCAST_MODE_SYNC | 2 | BROADCAST_MODE_SYNC defines a tx broadcasting mode where the client waits for a CheckTx execution response only. |
| BROADCAST_MODE_ASYNC | 3 | BROADCAST_MODE_ASYNC defines a tx broadcasting mode where the client returns immediately. |


## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

