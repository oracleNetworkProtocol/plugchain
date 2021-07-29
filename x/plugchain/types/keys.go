package types

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
	TokenCountKey  = []byte("tokenCountKey")
	TokenKeyPrefix = []byte{0x1}
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
