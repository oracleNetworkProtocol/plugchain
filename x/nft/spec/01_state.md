<!--
order: 1
-->

# State

## NFT

Nft defines the tokenData of non-fungible tokens

```go
type NFTI interface {
    GetOwner()  sdk.AccAddress
    GetData()   string
}

//NFT defines a non-fungible token
type NFT struct {
    ID      string
    Name    string
    URI     string
    Data    string
    Owner   string
}
```


## Collections 

As all NFTs belong to a specific `Collection`

```go

//collection of nft
type Collection struct {
    Class class `json:"class"` // class of the collection; not exported
    NFTs  []NFT `json:"nfts"` // nfts that belongs to a collection
}
```

## Class
Class defines a type of NFT
```go
type Class struct {
    ID      string
    Name    string
    Schema  string
    Symbol  string
    Owner   string
    MintRestricted bool // Is true mint nft
    EditRestricted bool // Is true edit nft
}
```