package types

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	evmtypes "github.com/tharsis/ethermint/x/evm/types"
)

// Copied the Account and StorageResult types since they are registered under an
// internal pkg on geth.

type TransactionArgs struct {
	From                 sdk.AccAddress `json:"from"`
	To                   sdk.AccAddress `json:"to"`
	Gas                  uint64         `json:"gas"`
	GasPrice             uint64         `json:"gasPrice"`
	MaxFeePerGas         uint64         `json:"maxFeePerGas"`
	MaxPriorityFeePerGas uint64         `json:"maxPriorityFeePerGas"`
	Value                uint64         `json:"value"`
	Nonce                uint64         `json:"nonce"`
	evmtypes.TransactionArgs
}

func (trans TransactionArgs) New() evmtypes.TransactionArgs {
	var args evmtypes.TransactionArgs
	args = trans.TransactionArgs
	if trans.From != nil {
		from := common.BytesToAddress(trans.From)
		args.From = &from
	}
	if trans.To != nil {
		to := common.BytesToAddress(trans.To)
		args.To = &to
	}
	if trans.Gas > 0 {
		args.Gas = (*hexutil.Uint64)(&trans.Gas)
	}
	if trans.GasPrice > 0 {
		args.GasPrice = (*hexutil.Big)(big.NewInt(int64(trans.GasPrice)))
	}
	if trans.MaxFeePerGas > 0 {
		args.MaxFeePerGas = (*hexutil.Big)(big.NewInt(int64(trans.MaxFeePerGas)))
	}
	if trans.MaxPriorityFeePerGas > 0 {
		args.MaxPriorityFeePerGas = (*hexutil.Big)(big.NewInt(int64(trans.MaxPriorityFeePerGas)))
	}
	if trans.Value > 0 {
		args.Value = (*hexutil.Big)(big.NewInt(int64(trans.Value)))
	}
	if trans.Nonce > 0 {
		args.Nonce = (*hexutil.Uint64)(&trans.Nonce)
	}
	return args
}

// AccountResult struct for account proof
type AccountResult struct {
	Address      common.Address  `json:"address"`
	AccountProof []string        `json:"accountProof"`
	Balance      *hexutil.Big    `json:"balance"`
	CodeHash     common.Hash     `json:"codeHash"`
	Nonce        hexutil.Uint64  `json:"nonce"`
	StorageHash  common.Hash     `json:"storageHash"`
	StorageProof []StorageResult `json:"storageProof"`
}

// StorageResult defines the format for storage proof return
type StorageResult struct {
	Key   string       `json:"key"`
	Value *hexutil.Big `json:"value"`
	Proof []string     `json:"proof"`
}

// RPCTransaction represents a transaction that will serialize to the RPC representation of a transaction
type RPCTransaction struct {
	BlockHash        *common.Hash         `json:"blockHash"`
	BlockNumber      *hexutil.Big         `json:"blockNumber"`
	From             common.Address       `json:"from"`
	Gas              hexutil.Uint64       `json:"gas"`
	GasPrice         *hexutil.Big         `json:"gasPrice"`
	GasFeeCap        *hexutil.Big         `json:"maxFeePerGas,omitempty"`
	GasTipCap        *hexutil.Big         `json:"maxPriorityFeePerGas,omitempty"`
	Hash             common.Hash          `json:"hash"`
	Input            hexutil.Bytes        `json:"input"`
	Nonce            hexutil.Uint64       `json:"nonce"`
	To               *common.Address      `json:"to"`
	TransactionIndex *hexutil.Uint64      `json:"transactionIndex"`
	Value            *hexutil.Big         `json:"value"`
	Type             hexutil.Uint64       `json:"type"`
	Accesses         *ethtypes.AccessList `json:"accessList,omitempty"`
	ChainID          *hexutil.Big         `json:"chainId,omitempty"`
	V                *hexutil.Big         `json:"v"`
	R                *hexutil.Big         `json:"r"`
	S                *hexutil.Big         `json:"s"`
}

// StateOverride is the collection of overridden accounts.
type StateOverride map[common.Address]OverrideAccount

// OverrideAccount indicates the overriding fields of account during the execution of
// a message call.
// Note, state and stateDiff can't be specified at the same time. If state is
// set, message execution will only use the data in the given state. Otherwise
// if statDiff is set, all diff will be applied first and then execute the call
// message.
type OverrideAccount struct {
	Nonce     *hexutil.Uint64              `json:"nonce"`
	Code      *hexutil.Bytes               `json:"code"`
	Balance   **hexutil.Big                `json:"balance"`
	State     *map[common.Hash]common.Hash `json:"state"`
	StateDiff *map[common.Hash]common.Hash `json:"stateDiff"`
}

type FeeHistoryResult struct {
	OldestBlock  *hexutil.Big     `json:"oldestBlock"`
	Reward       [][]*hexutil.Big `json:"reward,omitempty"`
	BaseFee      []*hexutil.Big   `json:"baseFeePerGas,omitempty"`
	GasUsedRatio []float64        `json:"gasUsedRatio"`
}

// SignTransactionResult represents a RLP encoded signed transaction.
type SignTransactionResult struct {
	Raw hexutil.Bytes         `json:"raw"`
	Tx  *ethtypes.Transaction `json:"tx"`
}

type OneFeeHistory struct {
	BaseFee      *big.Int   // base fee  for each block
	Reward       []*big.Int // each element of the array will have the tip provided to miners for the percentile given
	GasUsedRatio float64    // the ratio of gas used to the gas limit for each block
}