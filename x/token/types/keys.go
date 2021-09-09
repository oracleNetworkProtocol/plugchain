package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	// ModuleName defines the module name
	ModuleName = "token"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName
)

var (
	PrefixTokenForSymbol  = []byte{0x01}
	PrefixTokenForMinUint = []byte{0x02}
	PrefixTokens          = []byte{0x03}
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func KeySymbol(symbol string) []byte {
	return append(PrefixTokenForSymbol, []byte(symbol)...)
}
func KeyMinUint(minUnit string) []byte {
	return append(PrefixTokenForMinUint, []byte(minUnit)...)
}

func KeyTokens(owner sdk.AccAddress, symbol string) []byte {
	return append(append(PrefixTokens, owner.Bytes()...), []byte(symbol)...)
}
