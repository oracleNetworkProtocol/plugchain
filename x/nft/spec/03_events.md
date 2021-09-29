# Events

The nft module emits the following events:

## MsgIssueDenom

| Type        | Attribute Key | Attribute Value  |
| :---------- | :------------ | :--------------- |
| issue_denom | denom_id      | {nftDenomID}     |
| issue_denom | owner         | {ownerAddress}   |
| message     | module        | nft              |
| message     | sender        | {senderAddress}  |


## MsgIssueNFT

| Type      | Attribute Key | Attribute Value    |
| :-------  | :------------ | :----------------- |
| issue_nft | nft_id        | {nftID}            |
| issue_nft | denom_id      | {nftDenomID}       |
| issue_nft | recipient     | {recipientAddress} |
| message   | module        | nft                |
| message   | sender        | {senderAddress}    |

## MsgEditNFT

| Type     | Attribute Key | Attribute Value |
| :------- | :------------ | :-------------- |
| edit_nft | nft_id        | {nftID}         |
| edit_nft | denom_id      | {nftDenomID}    |
| edit_nft | owner         | {ownerAddress}  |
| message  | module        | nft             |
| message  | sender        | {ownerAddress} |