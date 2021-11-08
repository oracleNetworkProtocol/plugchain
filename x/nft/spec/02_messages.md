<!--
order: 2
-->

# Messages

## MsgIssueClass

| **Field** | **Type** | **Description**                                                                                                                  |
| :-------- | :------- | :------------------------------------------------------------------------------------------------------------------------------- |
| ID      | `string`     | The denomination ID of the NFT, necessary as multiple denominations are able to be represented on each chain.                    |
| Name      | `string` | The denomination name of the NFT, necessary as multiple denominations are able to be represented on each chain.                  |
| Owner    | `string` | The account address of the user sending the NFT. By default it is __not__ required that the sender is also the owner of the NFT. |
| Schema    | `string` | NFT specifications defined under this category                                                                                   |
| Symbol    | `string` | The abbreviated name of a specific NFT type                                                                                 |
| MintRestricted    | `bool` | MintRestricted is true means that only Denom owners can issue NFTs under this category, false means anyone can         |                                                                        |
| EditRestricted    | `bool` | EditRestricted is true means that no one in this category can Edit the NFT, false means that only the owner of this NFT can Edit   |                                                                             |

```go
type MsgIssueClass struct {
    ID     string
    Name   string
    Schema string
    Owner string
    Symbol string
    MintRestricted bool
    EditRestricted bool
}
```

## MsgIssueNFT

This message type is used for issuing new nft. If a new `NFT` is minted under a new `Denom`, a new `Collection` will also be created, otherwise the `NFT` is added to the existing `Collection`. If a new `NFT` is minted by a new account, a new `Owner` is created, otherwise the `NFT` `ID` is added to the existing `Owner`'s `IDCollection`. By default anyone can execute this Message type. **It is highly recommended that a custom handler is made to restrict use of this Message type to prevent unintended use.**

| **Field** | **Type** | **Description**                                                                            |
| :-------- | :------- | :----------------------------------------------------------------------------------------- |
| ID        | `string` | The unique ID of the NFT being minted                                                      |
| ClassID   | `string` | The unique ID of the denomination.                                                         |
| Name      | `string` | The name of the NFT being minted.                                                          |
| URL       | `string` | The URL pointing to a JSON object that contains subsequent nftData information off-chain |
| Data      | `string` | The data of the NFT.                                                                       |
| owner    | `string` | The owner of the Message                                                                  |
| Recipient | `string` | The recipiet of the new NFT                                                                |

```go
// MsgIssueNFT defines an SDK message for creating a new NFT.
type MsgIssueNFT struct {
    ID        string
    ClassID   string
    Name      string
    URL       string
    Data      string
    Owner    string
    Recipient string
}
```


## MsgEditNFT

This message type allows the `NFTURL` to be updated. By default anyone can execute this Message type. **It is highly recommended that a custom handler is made to restrict use of this Message type to prevent unintended use.**

| **Field** | **Type** | **Description**                                                                                                  |
| :-------- | :------- | :--------------------------------------------------------------------------------------------------------------- |
| ID        | `string` | The unique ID of the NFT being edited.                                                                           |
| ClassID   | `string` | The unique ID of the denomination, necessary as multiple denominations are able to be represented on each chain. |
| Name      | `string` | The name of the NFT being edited.                                                                                |
| URL       | `string` | The URL pointing to a JSON object that contains subsequent nftData information off-chain                       |
| Data      | `string` | The data of the NFT                                                                                              |
| Owner    | `string` | The creator of the message                                                                                       |

```go
// MsgEditNFT defines an SDK message for editing a nft.
type MsgEditNFT struct {
    ID      string
    ClassID string
    Name    string
    URL     string
    Data    string
    Owner  string
}
```


### MsgBurnNFT

This message type is used for burning tokens which destroys and deletes them. By default anyone can execute this Message type. **It is highly recommended that a custom handler is made to restrict use of this Message type to prevent unintended use.**

| **Field** | **Type** | **Description**                                    |
| :-------- | :------- | :------------------------------------------------- |
| ID        | `string` | The ID of the nft.                                 |
| ClassID   | `string` | The Denom ID of the Denom.                         |
| Owner     | `string` | The account address of the user burning the denom. |

```go
// MsgBurnNFT defines an SDK message for burning a NFT.
type MsgBurnNFT struct {
    ID      string
    ClassID string
    Owner  string
}
```

## MsgTransferNFT

This is the most commonly expected MsgType to be supported across chains. While each application specific blockchain will have very different adoption of the `MsgMintNFT`, `MsgBurnNFT` and `MsgEditNFT` it should be expected that most chains support the ability to transfer ownership of the NFT asset. The exception to this would be non-transferable NFTs that might be attached to reputation or some asset which should not be transferable. It still makes sense for this to be represented as an NFT because there are common queriers which will remain relevant to the NFT type even if non-transferable. This Message will fail if the NFT does not exist. By default it will not fail if the transfer is executed by someone beside the owner. **It is highly recommended that a custom handler is made to restrict use of this Message type to prevent unintended use.**

| **Field** | **Type** | **Description**                                                                                                                  |
| :-------- | :------- | :------------------------------------------------------------------------------------------------------------------------------- |
| ID        | `string` | The unique ID of the NFT being transferred.                                                                                      |
| ClassID   | `string` | The unique ID of the denomination, necessary as multiple denominations are able to be represented on each chain.                 |
| Owner    | `string` | The account address of the user sending the NFT. By default it is __not__ required that the sender is also the owner of the NFT. |
| Recipient | `string` | The account address who will receive the NFT as a result of the transfer transaction.                                            |

```go
// MsgTransferNFT defines an SDK message for transferring an NFT to recipient.
type MsgTransferNFT struct {
    ID        string
    ClassID   string
    Owner    string
    Recipient string
}
```





## MsgTransferClass
This message is used by the owner of the NFT classification to transfer the ownership of the NFT classification to others

| **Field** | **Type** | **Description**                                                                                                                  |
| :-------- | :------- | :------------------------------------------------------------------------------------------------------------------------------- |
| ID        | `string` | The unique ID of the Denom being transferred.                                                                                      | 
| Owner     | `string` | The account address of the user sending the Denom. |
| Recipient | `string` | The account address who will receive the Denom as a result of the transfer transaction.                                            |

```go
// MsgTransferClass defines an SDK message for transferring an Class to recipient.
type MsgTransferClass struct {
    ID        string
    Owner     string
    Recipient string
}
```