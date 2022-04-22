<!--
order: 1
-->

# State

## Token

Definition of data structure of Token

```go
type Token struct {
    Symbol        string
    Name          string
    Scale         uint32
    MinUnit       string
    InitialSupply uint64
    MaxSupply     uint64
    Mintable      bool
    Owner         string
}
```

## Params

Params is a module-wide configuration structure that stores system parameters and defines overall functioning of the token module.

```go
type Params struct {
    IssueTokenBaseFee sdk.Coin
    OperateTokenFeeRatio sdk.Dec
}
```
