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

## MsgBurnNFT

| Type     | Attribute Key | Attribute Value |
| :------- | :------------ | :-------------- |
| burn_nft | nft_id        | {nftID}         |
| burn_nft | denom_id      | {nftDenomID}    |
| burn_nft | owner         | {ownerAddress}  |
| message  | module        | nft             |
| message  | sender        | {ownerAddress} |

## MsgTransferNFT

| Type                  | Attribute Key | Attribute Value   |
| :---------------      | :------------ | :--------------   |
| transfer_nft          | nft_id        | {nftID}           |
| transfer_nft          | denom_id      | {DenomID}         |
| transfer_nft          | owner         | {ownerAddress}    |
| transfer_nft          | recipient     | {recipientAddress}|
| message               | module        | nft               |
| message               | sender        | {ownerAddress}    |


## MsgTransferDenom

| Type                  | Attribute Key | Attribute Value   |
| :---------------      | :------------ | :--------------   |
| transfer_denom        | id            | {ID}              |
| transfer_denom        | owner         | {ownerAddress}    |
| transfer_nft          | recipient     | {recipientAddress}|
| message               | module        | nft               |
| message               | sender        | {ownerAddress}    |

