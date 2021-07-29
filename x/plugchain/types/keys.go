package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	// ModuleName defines the module name
	ModuleName = "plugchain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_plugchain"

	//token name abbreviation prefix
	TokenNamePrefix = "token"

	//Maximum length of token shorthand
	SymbolMaxLen int = 4
	//Default decimal places
	DefaultDecimal uint32 = 6

	DescriptionMaxLen = 256
)

var (
	TokenCountKey         = []byte("tokenCountKey")
	TokenKeyPrefix        = []byte{0x1}
	TokenAccAddressPrefix = []byte{0x2}
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

// GetPoolKey returns kv indexing key of the pool

func GetTokenKey(symbol string) []byte {
	// 10 = len([]byte(symbol))
	key := make([]byte, 10)
	key[0] = TokenKeyPrefix[0]
	copy(key[1:], KeyPrefix(symbol))
	return key
}

func GetTokenAccAddressPrefix(accAddr sdk.AccAddress) []byte {
	adrLen := len(accAddr)
	key := make([]byte, adrLen+1)
	key[0] = TokenAccAddressPrefix[0]
	copy(key[1:], accAddr)
	return key
}
