<!--
order: 3
-->

# Events

The token module emits the following events:

## Handlers

### MsgIssueToken

| Type        | Attribute Key | Attribute Value  |
| :---------- | :------------ | :--------------- |
| issue_token | symbol        | {symbol}         |
| issue_token | creator       | {creatorAddress} |
| message     | module        | token            |
| message     | sender        | {ownerAddress}   |

### MsgEditToken

| Type       | Attribute Key | Attribute Value |
| :--------- | :------------ | :-------------- |
| edit_token | symbol        | {symbol}        |
| edit_token | owner         | {ownerAddress}  |
| message    | module        | token           |
| message    | sender        | {ownerAddress}  |

### MsgTransferTokenOwner

| Type                 | Attribute Key | Attribute Value   |
| :------------------- | :------------ | :---------------- |
| transfer_owner_token | symbol        | {symbol}          |
| transfer_owner_token | owner         | {ownerAddress}    |
| transfer_owner_token | to            | {toAddress}  |
| message              | module        | token             |
| message              | sender        | {ownerAddress}    |

### MsgMintToken

| Type       | Attribute Key | Attribute Value    |
| :--------- | :------------ | :----------------- |
| mint_token | symbol        | {symbol}           |
| mint_token | amount        | {amount}           |
| mint_token | to            | {toAddress}        |
| message    | module        | token              |
| message    | sender        | {ownerAddress}     |

### MsgBurnToken

| Type       | Attribute Key | Attribute Value |
| :--------- | :------------ | :-------------- |
| burn_token | symbol        | {symbol}        |
| burn_token | amount        | {amount}        |
| message    | module        | token           |
| message    | sender        | {ownerAddress}  |
