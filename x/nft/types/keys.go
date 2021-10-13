package types

const (
	// ModuleName defines the module name
	ModuleName = "nft"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName
)

var (
	PrefixDenom      = []byte{0x01}
	PrefixNFT        = []byte{0x02}
	delimiter        = []byte("-")
	PrefixCollection = []byte{0x03}
)

func GetKeyDenomID(id string) []byte {
	key := append(PrefixDenom, delimiter...)
	return append(key, []byte(id)...)
}

// KeyNFT gets the key of nft stored by an denom and id
func GetKeyNFT(denomID, nftID string) []byte {
	baseKey := append(PrefixNFT, delimiter...)
	if len(denomID) > 0 {
		baseKey = append(baseKey, []byte(denomID)...)
		baseKey = append(baseKey, delimiter...)
	}
	if len(denomID) > 0 && len(nftID) > 0 {
		baseKey = append(baseKey, []byte(nftID)...)
	}
	return baseKey
}

func KeyCollectionByDenomID(denomID string) []byte {
	key := append(PrefixCollection, delimiter...)
	return append(key, []byte(denomID)...)
}
