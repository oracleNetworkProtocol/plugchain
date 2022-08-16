package rpc

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	rpctypes "github.com/oracleNetworkProtocol/plugchain/rpc/types"
)

//AccountResult refactored structure address type to sdk.AccAddress
type ProofAccountResult struct {
	rpctypes.AccountResult
}

// MarshalJSON marshals as JSON.
func (p ProofAccountResult) MarshalJSON() ([]byte, error) {
	type ProofAccountResult struct {
		Address      sdk.AccAddress           `json:"address"`
		AccountProof []string                 `json:"accountProof"`
		Balance      *hexutil.Big             `json:"balance"`
		CodeHash     common.Hash              `json:"codeHash"`
		Nonce        hexutil.Uint64           `json:"nonce"`
		StorageHash  common.Hash              `json:"storageHash"`
		StorageProof []rpctypes.StorageResult `json:"storageProof"`
	}
	var enc ProofAccountResult
	enc.Address = sdk.AccAddress(p.Address[:])
	enc.AccountProof = p.AccountProof
	enc.Balance = p.Balance
	enc.CodeHash = p.CodeHash
	enc.Nonce = p.Nonce
	enc.StorageHash = p.StorageHash
	enc.StorageProof = p.StorageProof

	return json.Marshal(&enc)
}
