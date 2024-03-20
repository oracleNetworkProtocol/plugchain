package types

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
)

// Keys for oracle store, with <prefix><key> -> <value>
const (
	// ModuleName defines the module name
	ModuleName = "oracle"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_capability"
)

// Keys for oracle store, with <prefix><key> -> <value>
var (
	// 0x00 | claim_hash -> claim
	ClaimKey = []byte{0x00}

	// 0x01 | claimType | roundId -> round
	RoundKey = []byte{0x01}

	// 0x02 | claimType | roundId -> roundId
	PendingRoundKey = []byte{0x02}

	// - 0x03 | prevote_hash -> prevote_hash
	PrevoteKey = []byte{0x03}

	// - 0x04 | del_address -> val_address
	DelValKey = []byte{0x05} // key for validator feed delegation

	// - 0x05 | val_address -> del_address
	ValDelKey = []byte{0x06} // key for validator feed delegation

	// - 0x06 | claimType -> roundId
	FinalizedRoundKey = []byte{0x07} // key for validator feed delegation
)

// KeyPrefix helper
func KeyPrefix(p string) []byte {
	return []byte(p)
}

// GetDelValKey returns the validator for a given delegate key
func GetDelValKey(del sdk.AccAddress) []byte {
	return append(DelValKey, del.Bytes()...)
}

// GetValDelKey returns the validator for a given delegate key
func GetValDelKey(val sdk.AccAddress) []byte {
	return append(ValDelKey, val.Bytes()...)
}

// GetClaimPrevoteKey returns the key for a validators prevote
func GetClaimPrevoteKey(hash tmbytes.HexBytes) []byte {
	return append(PrevoteKey, hash...)
}

// GetRoundKey helper
func GetRoundKey(claimType string, roundID uint64) string {
	return claimType + strconv.FormatUint(roundID, 10)
}

// RoundPrefix helper
func RoundPrefix(claimType string, roundID uint64) []byte {
	return KeyPrefix(GetRoundKey(claimType, roundID))
}
