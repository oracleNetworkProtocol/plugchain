# Events

The nft module emits the following events:

## MsgIssueClass

| Type        | Attribute Key | Attribute Value  |
| :---------- | :------------ | :--------------- |
| issue_class | class_id      | {nftClassID}     |
| issue_class | owner         | {ownerAddress}   |
| message     | module        | nft              |
| message     | sender        | {senderAddress}  |


## MsgIssueNFT

| Type      | Attribute Key | Attribute Value    |
| :-------  | :------------ | :----------------- |
| issue_nft | nft_id        | {nftID}            |
| issue_nft | class_id      | {nftClassID}       |
| issue_nft | recipient     | {recipientAddress} |
| message   | module        | nft                |
| message   | sender        | {senderAddress}    |

## MsgEditNFT

| Type     | Attribute Key | Attribute Value |
| :------- | :------------ | :-------------- |
| edit_nft | nft_id        | {nftID}         |
| edit_nft | class_id      | {nftClassID}    |
| edit_nft | owner         | {ownerAddress}  |
| message  | module        | nft             |
| message  | sender        | {ownerAddress} |

## MsgBurnNFT

| Type     | Attribute Key | Attribute Value |
| :------- | :------------ | :-------------- |
| burn_nft | nft_id        | {nftID}         |
| burn_nft | class_id      | {nftClassID}    |
| burn_nft | owner         | {ownerAddress}  |
| message  | module        | nft             |
| message  | sender        | {ownerAddress} |

## MsgTransferNFT

| Type                  | Attribute Key | Attribute Value   |
| :---------------      | :------------ | :--------------   |
| transfer_nft          | nft_id        | {nftID}           |
| transfer_nft          | class_id      | {ClassID}         |
| transfer_nft          | owner         | {ownerAddress}    |
| transfer_nft          | recipient     | {recipientAddress}|
| message               | module        | nft               |
| message               | sender        | {ownerAddress}    |


## MsgTransferClass

| Type                  | Attribute Key | Attribute Value   |
| :---------------      | :------------ | :--------------   |
| transfer_class        | id            | {ID}              |
| transfer_class        | owner         | {ownerAddress}    |
| transfer_nft          | recipient     | {recipientAddress}|
| message               | module        | nft               |
| message               | sender        | {ownerAddress}    |

