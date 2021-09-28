<!--
order: 2
-->

# Messages

## MsgIssueDenom

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
type MsgIssueDenom struct {
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
| DenomID   | `string` | The unique ID of the denomination.                                                         |
| Name      | `string` | The name of the NFT being minted.                                                          |
| URL       | `string` | The URL pointing to a JSON object that contains subsequent nftData information off-chain |
| Data      | `string` | The data of the NFT.                                                                       |
| owner    | `string` | The owner of the Message                                                                  |
| Recipient | `string` | The recipiet of the new NFT                                                                |

```go
// MsgIssueNFT defines an SDK message for creating a new NFT.
type MsgIssueNFT struct {
    ID        string
    DenomID   string
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
| DenomID   | `string` | The unique ID of the denomination, necessary as multiple denominations are able to be represented on each chain. |
| Name      | `string` | The name of the NFT being edited.                                                                                |
| URL       | `string` | The URL pointing to a JSON object that contains subsequent nftData information off-chain                       |
| Data      | `string` | The data of the NFT                                                                                              |
| Owner    | `string` | The creator of the message                                                                                       |

```go
// MsgEditNFT defines an SDK message for editing a nft.
type MsgEditNFT struct {
    ID      string
    DenomID string
    Name    string
    URL     string
    Data    string
    Owner  string
}
```