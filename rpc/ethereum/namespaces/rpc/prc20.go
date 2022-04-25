package rpc

import (
	"fmt"
	"math"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	pvmtypes "github.com/oracleNetworkProtocol/plugchain/types"
	"github.com/tendermint/tendermint/libs/bytes"

	"github.com/pkg/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	rpctypes "github.com/oracleNetworkProtocol/plugchain/rpc/ethereum/types"
	evmtypes "github.com/tharsis/ethermint/x/evm/types"
)

// BalanceOf returns prc20 balance
func (e *PublicAPI) BalanceOf(trans rpctypes.TransactionArgs, blockNrOrHash rpctypes.BlockNumberOrHash) (interface{}, error) {

	e.logger.Debug("rpc_balanceOf")

	if trans.Data == nil {
		data, err := pvmtypes.ERC20Contract.ABI.Pack("balanceOf", common.BytesToAddress(trans.From.Bytes()))
		if err != nil {
			return nil, err
		}
		trans.Data = (*hexutil.Bytes)(&data)
	}
	res, err := e.Call(trans, blockNrOrHash, nil)
	if err != nil {
		return nil, err
	}
	unRes, err := pvmtypes.ERC20Contract.ABI.Unpack("balanceOf", res)
	if err != nil {
		return nil, err
	}
	return unRes, nil
}

// Prc20TokenInfo returns to the basic information of prc20 token contract
func (e *PublicAPI) Prc20TokenInfo(trans rpctypes.TransactionArgs, blockNrOrHash rpctypes.BlockNumberOrHash) map[string]interface{} {
	e.logger.Debug("rpc_prc20TokenInfo")

	method := []string{"name", "symbol", "decimals", "totalSupply"}
	var result = make(map[string]interface{})
	for _, v := range method {
		data, err := pvmtypes.ERC20Contract.ABI.Pack(v)
		if err != nil {
			continue
		}
		trans.Data = (*hexutil.Bytes)(&data)
		res, _ := e.Call(trans, blockNrOrHash, nil)

		unRes, _ := pvmtypes.ERC20Contract.ABI.Unpack(v, res)
		result[v] = unRes
	}

	return result
}

// AddressTranslation returns the byte, hex, eip-55, bech32 type information of the incoming address
func (e *PublicAPI) AddressTranslation(arg string) (map[string]interface{}, error) {
	e.logger.Debug("rpc_addressTranslation")
	var (
		addr []byte
		err  error
	)
	switch {
	case common.IsHexAddress(arg):
		addr = common.HexToAddress(arg).Bytes()
	case strings.HasPrefix(arg, sdk.GetConfig().GetBech32ValidatorAddrPrefix()):
		addr, err = sdk.ValAddressFromBech32(arg)
	case strings.HasPrefix(arg, sdk.GetConfig().GetBech32AccountAddrPrefix()):
		addr, err = sdk.AccAddressFromBech32(arg)
	default:
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"bytes":  fmt.Sprint(addr),
		"bech32": sdk.AccAddress(addr),
		"hex":    bytes.HexBytes(addr),
		"EIP-55": common.BytesToAddress(addr),
	}, nil
}

// GetProof returns an account object with proof and any storage proofs
func (e *PublicAPI) GetProof(addr sdk.AccAddress, storageKeys []string, blockNrOrHash rpctypes.BlockNumberOrHash) (*ProofAccountResult, error) {

	e.logger.Debug("rpc_getProof", "address", addr.String(), "keys", storageKeys, "block number or hash", blockNrOrHash)

	blockNum, err := e.getBlockNumber(blockNrOrHash)
	if err != nil {
		return nil, err
	}

	height := blockNum.Int64()
	ctx := rpctypes.ContextWithHeight(height)

	// if the height is equal to zero, meaning the query condition of the block is either "pending" or "latest"
	if height == 0 {
		bn, err := e.backend.BlockNumber()
		if err != nil {
			return nil, err
		}

		if bn > math.MaxInt64 {
			return nil, fmt.Errorf("not able to query block number greater than MaxInt64")
		}

		height = int64(bn)
	}

	clientCtx := e.clientCtx.WithHeight(height)

	// query storage proofs
	storageProofs := make([]rpctypes.StorageResult, len(storageKeys))
	address := common.BytesToAddress(addr)

	for i, key := range storageKeys {
		hexKey := common.HexToHash(key)
		valueBz, proof, err := e.queryClient.GetProof(clientCtx, evmtypes.StoreKey, evmtypes.StateKey(address, hexKey.Bytes()))
		if err != nil {
			return nil, err
		}

		// check for proof
		var proofStr string
		if proof != nil {
			proofStr = proof.String()
		}

		storageProofs[i] = rpctypes.StorageResult{
			Key:   key,
			Value: (*hexutil.Big)(new(big.Int).SetBytes(valueBz)),
			Proof: []string{proofStr},
		}
	}

	// query EVM account
	req := &evmtypes.QueryAccountRequest{
		Address: address.String(),
	}

	res, err := e.queryClient.Account(ctx, req)
	if err != nil {
		return nil, err
	}

	// query account proofs
	accountKey := authtypes.AddressStoreKey(addr)
	_, proof, err := e.queryClient.GetProof(clientCtx, authtypes.StoreKey, accountKey)
	if err != nil {
		return nil, err
	}

	// check for proof
	var accProofStr string
	if proof != nil {
		accProofStr = proof.String()
	}

	balance, ok := sdk.NewIntFromString(res.Balance)
	if !ok {
		return nil, errors.New("invalid balance")
	}
	return &ProofAccountResult{
		Address: addr,
		AccountResult: rpctypes.AccountResult{
			AccountProof: []string{accProofStr},
			Balance:      (*hexutil.Big)(balance.BigInt()),
			CodeHash:     common.HexToHash(res.CodeHash),
			Nonce:        hexutil.Uint64(res.Nonce),
			StorageHash:  common.Hash{}, // NOTE: Ethermint doesn't have a storage hash. TODO: implement?
			StorageProof: storageProofs,
		},
	}, nil

}

// getBlockNumber returns the BlockNumber from BlockNumberOrHash
func (e *PublicAPI) getBlockNumber(blockNrOrHash rpctypes.BlockNumberOrHash) (rpctypes.BlockNumber, error) {
	switch {
	case blockNrOrHash.BlockHash == nil && blockNrOrHash.BlockNumber == nil:
		return rpctypes.EthEarliestBlockNumber, fmt.Errorf("types BlockHash and BlockNumber cannot be both nil")
	case blockNrOrHash.BlockHash != nil:
		blockHeader, err := e.backend.HeaderByHash(*blockNrOrHash.BlockHash)
		if err != nil {
			return rpctypes.EthEarliestBlockNumber, err
		}
		return rpctypes.NewBlockNumber(blockHeader.Number), nil
	case blockNrOrHash.BlockNumber != nil:
		return *blockNrOrHash.BlockNumber, nil
	default:
		return rpctypes.EthEarliestBlockNumber, nil
	}
}

// GetTransactionByHash returns the transaction identified by hash.
func (e *PublicAPI) GetTransactionByHash(hash common.Hash) (*RPCTransactionResult, error) {

	e.logger.Debug("rpc_getTransactionByHash", "hash", hash.Hex())

	res, err := e.backend.GetTransactionByHash(hash)
	if err != nil {
		return &RPCTransactionResult{}, err
	}
	var (
		from sdk.AccAddress
		to   sdk.AccAddress
	)
	if len(res.From) > 2 {
		from = sdk.AccAddress(res.From.Bytes())
	}
	if res.To != nil {
		to = sdk.AccAddress(res.To.Bytes())
	}
	return &RPCTransactionResult{
		From: from,
		To:   to,
		RPCTransaction: rpctypes.RPCTransaction{
			BlockHash:   res.BlockHash,
			BlockNumber: res.BlockNumber,
			Type:        res.Type,
			Gas:         res.Gas,
			GasPrice:    res.GasPrice,
			Hash:        res.Hash,
			Input:       res.Input,
			Nonce:       res.Nonce,
			Value:       res.Value,
			V:           res.V,
			R:           res.R,
			S:           res.S,
		},
	}, nil
}

// GetBlockByHash returns the block identified by hash.
func (e *PublicAPI) GetBlockByHash(hash common.Hash, fullTx bool) (result map[string]interface{}, err error) {
	e.logger.Debug("rpc_getBlockByHash", "hash", hash.Hex(), "full", fullTx)

	result, err = e.backend.GetBlockByHash(hash, fullTx)

	if _, ok := result["miner"]; ok {
		result["miner"] = sdk.AccAddress(result["miner"].(common.Address).Bytes())
	}
	return
}

// GetBlockByNumber returns the block identified by number.
func (e *PublicAPI) GetBlockByNumber(ethBlockNum rpctypes.BlockNumber, fullTx bool) (result map[string]interface{}, err error) {
	e.logger.Debug("rpc_getBlockByNumber", "number", ethBlockNum, "full", fullTx)
	result, err = e.backend.GetBlockByNumber(ethBlockNum, fullTx)

	if _, ok := result["miner"]; ok {
		result["miner"] = sdk.AccAddress(result["miner"].(common.Address).Bytes())
	}
	return
}

// EstimateGas returns an estimate of gas usage for the given smart contract call.
func (e *PublicAPI) EstimateGas(args evmtypes.TransactionArgs, blockNrOptional *rpctypes.BlockNumber) (hexutil.Uint64, error) {
	e.logger.Debug("rpc_estimateGas")
	return e.backend.EstimateGas(args, blockNrOptional)
}

func (e *PublicAPI) getTransactionLogs(txHash common.Hash) ([]*evmtypes.Log, error) {
	tx, err := e.backend.GetTxByEthHash(txHash)
	if err != nil {
		return nil, err
	}

	return TxLogsFromEvents(tx.TxResult.Events)
}
