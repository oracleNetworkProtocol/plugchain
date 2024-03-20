<!--
order: 4
-->

# Events

The oracle module emits the following events:

### MsgDelegate

| Type     | Attribute Key | Attribute Value    |
| -------- | ------------- | ------------------ |
| delegate | validator     | {validatorAddress} |
| delegate | delegate      | {delegateAddress}  |
| message  | module        | oracle             |
| message  | action        | delegate           |
| message  | sender        | {senderAddress}    |

### MsgPrevote

| Type    | Attribute Key | Attribute Value    |
| ------- | ------------- | ------------------ |
| prevote | prevote_hash  | {prevoteHash}      |
| prevote | validator     | {validatorAddress} |
| message | module        | oracle             |
| message | action        | prevote            |
| message | sender        | {senderAddress}    |

### MsgVote

| Type    | Attribute Key | Attribute Value    |
| ------- | ------------- | ------------------ |
| vote    | claim_hash    | {claimHash}        |
| vote    | validator     | {validatorAddress} |
| message | module        | oracle             |
| message | action        | vote               |
| message | sender        | {senderAddress}    |
