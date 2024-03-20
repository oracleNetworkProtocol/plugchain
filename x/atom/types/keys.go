package types

const (
	// ModuleName defines the module name
	ModuleName = "atom"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_capability"

	// Atom Price Claim
	AtomClaim = "AtomClaim"
)

// Keys for oracle store, with <prefix><key> -> <value>
var (
	// 0x00 -> Dec
	AtomUsdKey = []byte{0x00}
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
