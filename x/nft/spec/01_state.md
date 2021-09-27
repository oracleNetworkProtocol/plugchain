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
    URL     string
    Data    string
    Owner   string
}
```


## Collections 

As all NFTs belong to a specific `Collection`

```go

//collection of nft
type Collection struct {
    Denom Denom `json:"denom"` // denom of the collection; not exported
    NFTs  []NFT `json:"nfts"` // nfts that belongs to a collection
}
```

## Denom
Denom defines a type of NFT
```go
type Denom struct {
    ID      string
    Name    string
    Schema  string
    Symbol  string
    Owner   string
    MintRestricted bool // Is true mint nft
    EditRestricted bool // Is true edit nft
}
```