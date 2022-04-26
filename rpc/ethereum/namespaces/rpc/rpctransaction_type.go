package rpc

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	rpctypes "github.com/oracleNetworkProtocol/plugchain/rpc/ethereum/types"
)

type RPCTransactionResult struct {
	*rpctypes.RPCTransaction
}

// MarshalJSON marshals as JSON.
func (r RPCTransactionResult) MarshalJSON() ([]byte, error) {
	type RPCTransactionResult struct {
		BlockHash        *common.Hash    `json:"blockHash"`
		BlockNumber      *hexutil.Big    `json:"blockNumber"`
		From             sdk.AccAddress  `json:"from"`
		Gas              hexutil.Uint64  `json:"gas"`
		GasPrice         *hexutil.Big    `json:"gasPrice"`
		GasFeeCap        *hexutil.Big    `json:"maxFeePerGas,omitempty"`
		GasTipCap        *hexutil.Big    `json:"maxPriorityFeePerGas,omitempty"`
		Hash             common.Hash     `json:"hash"`
		Input            hexutil.Bytes   `json:"input"`
		Nonce            hexutil.Uint64  `json:"nonce"`
		To               sdk.AccAddress  `json:"to"`
		TransactionIndex *hexutil.Uint64 `json:"transactionIndex"`
		Value            *hexutil.Big    `json:"value"`
		Type             hexutil.Uint64  `json:"type"`
		Accesses         AccessList      `json:"accessList,omitempty"`
		ChainID          *hexutil.Big    `json:"chainId,omitempty"`
		V                *hexutil.Big    `json:"v"`
		R                *hexutil.Big    `json:"r"`
		S                *hexutil.Big    `json:"s"`
	}
	var enc RPCTransactionResult
	enc.BlockHash = r.BlockHash
	enc.BlockNumber = r.BlockNumber
	enc.From = sdk.AccAddress(r.From[:])
	enc.Gas = r.Gas
	enc.GasPrice = r.GasPrice
	enc.GasFeeCap = r.GasFeeCap
	enc.GasTipCap = r.GasTipCap
	enc.Hash = r.Hash
	enc.Input = r.Input
	enc.Nonce = r.Nonce
	if r.To != nil {
		to := sdk.AccAddress(r.To[:])
		enc.To = to
	}
	enc.TransactionIndex = r.TransactionIndex
	enc.Value = r.Value
	enc.Type = r.Type
	if r.Accesses != nil {
		enc.Accesses = AccessListReplace(*r.Accesses)
	}
	enc.ChainID = r.ChainID
	enc.V = r.V
	enc.R = r.R
	enc.S = r.S
	return json.Marshal(&enc)
}
