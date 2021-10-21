package types

import (
	"bytes"
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

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
	PrefixCollection = []byte{0x03}
	PrefixOwners     = []byte{0x04}

	delimiter = []byte("-")
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

func GetKeyOwner(adr sdk.AccAddress, denomID, nftID string) []byte {

	key := append(PrefixOwners, delimiter...)
	if adr != nil {
		key = append(key, []byte(adr.String())...)
		key = append(key, delimiter...)
	}

	if adr != nil && len(denomID) > 0 {
		key = append(key, []byte(denomID)...)
		key = append(key, delimiter...)
	}

	if adr != nil && len(denomID) > 0 && len(nftID) > 0 {
		key = append(key, []byte(nftID)...)
	}
	return key
}

func SplitKeyDenom(key []byte) (denomID, nftID string, err error) {
	keys := bytes.Split(key, delimiter)
	if len(keys) != 2 {
		return denomID, nftID, errors.New("wrong keyOwner")
	}
	denomID = string(keys[0])
	nftID = string(keys[1])
	return
}
