package rpc

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

type AccessList []AccessTuple

// StorageKeys returns the total number of storage keys in the access list.
func (al AccessList) StorageKeys() int {
	sum := 0
	for _, tuple := range al {
		sum += len(tuple.StorageKeys)
	}
	return sum
}

// AccessTuple is the element type of an access list.
type AccessTuple struct {
	Address     sdk.AccAddress `json:"address"        gencodec:"required"`
	StorageKeys []common.Hash  `json:"storageKeys"    gencodec:"required"`
}

// MarshalJSON marshals as JSON.
func (a AccessTuple) MarshalJSON() ([]byte, error) {
	type AccessTuple struct {
		Address     sdk.AccAddress `json:"address"        gencodec:"required"`
		StorageKeys []common.Hash  `json:"storageKeys"    gencodec:"required"`
	}
	var enc AccessTuple
	enc.Address = a.Address
	enc.StorageKeys = a.StorageKeys
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (a *AccessTuple) UnmarshalJSON(input []byte) error {
	type AccessTuple struct {
		Address     *sdk.AccAddress `json:"address"        gencodec:"required"`
		StorageKeys []common.Hash   `json:"storageKeys"    gencodec:"required"`
	}
	var dec AccessTuple
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Address == nil {
		return errors.New("missing required field 'address' for AccessTuple")
	}
	a.Address = *dec.Address
	if dec.StorageKeys == nil {
		return errors.New("missing required field 'storageKeys' for AccessTuple")
	}
	a.StorageKeys = dec.StorageKeys
	return nil
}

//AccessListReplace
func AccessListReplace(eal ethtypes.AccessList) AccessList {
	len := len(eal)
	if len == 0 {
		return nil
	}
	var access = make(AccessList, len)

	for k, v := range eal {
		aT := AccessTuple{
			Address:     sdk.AccAddress(v.Address[:]),
			StorageKeys: v.StorageKeys,
		}
		access[k] = aT
	}
	return access
}
