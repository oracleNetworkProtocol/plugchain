package rpc

import (
	"fmt"
	"math"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	ethtypes "github.com/ethereum/go-ethereum/core/types"
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
	return unRes[0], nil
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
		res, err := e.Call(trans, blockNrOrHash, nil)
		if err != nil {
			continue
		}

		unRes, err := pvmtypes.ERC20Contract.ABI.Unpack(v, res)
		if err != nil {
			continue
		}
		result[v] = unRes[0]
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
	if err != nil || res == nil {
		return &RPCTransactionResult{}, err
	}
	return &RPCTransactionResult{
		RPCTransaction: res,
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
func (e *PublicAPI) EstimateGas(trans rpctypes.TransactionArgs, blockNrOptional *rpctypes.BlockNumber) (hexutil.Uint64, error) {
	e.logger.Debug("rpc_estimateGas")
	args := trans.New()
	return e.backend.EstimateGas(args, blockNrOptional)
}

// GetTransactionReceipt returns the transaction receipt identified by hash.
func (e *PublicAPI) GetTransactionReceipt(hash common.Hash) (map[string]interface{}, error) {
	hexTx := hash.Hex()
	e.logger.Debug("rpc_getTransactionReceipt", "hash", hexTx)

	res, err := e.backend.GetTxByEthHash(hash)
	if err != nil {
		e.logger.Debug("tx not found", "hash", hexTx, "error", err.Error())
		return nil, nil
	}

	resBlock, err := e.clientCtx.Client.Block(e.ctx, &res.Height)
	if err != nil {
		e.logger.Debug("block not found", "height", res.Height, "error", err.Error())
		return nil, nil
	}

	tx, err := e.clientCtx.TxConfig.TxDecoder()(res.Tx)
	if err != nil {
		e.logger.Debug("decoding failed", "error", err.Error())
		return nil, fmt.Errorf("failed to decode tx: %w", err)
	}

	msg, err := evmtypes.UnwrapEthereumMsg(&tx)
	if err != nil {
		e.logger.Debug("invalid tx", "error", err.Error())
		return nil, err
	}

	txData, err := evmtypes.UnpackTxData(msg.Data)
	if err != nil {
		e.logger.Error("failed to unpack tx data", "error", err.Error())
		return nil, err
	}

	cumulativeGasUsed := uint64(0)
	blockRes, err := e.clientCtx.Client.BlockResults(e.ctx, &res.Height)
	if err != nil {
		e.logger.Debug("failed to retrieve block results", "height", res.Height, "error", err.Error())
		return nil, nil
	}

	for i := 0; i <= int(res.Index) && i < len(blockRes.TxsResults); i++ {
		cumulativeGasUsed += uint64(blockRes.TxsResults[i].GasUsed)
	}

	// Get the transaction result from the log
	var status hexutil.Uint
	if strings.Contains(res.TxResult.GetLog(), evmtypes.AttributeKeyEthereumTxFailed) {
		status = hexutil.Uint(ethtypes.ReceiptStatusFailed)
	} else {
		status = hexutil.Uint(ethtypes.ReceiptStatusSuccessful)
	}

	from, err := msg.GetSender(e.chainIDEpoch)
	if err != nil {
		return nil, err
	}
	to := txData.GetTo()

	logs, err := e.backend.GetTransactionLogs(hash)
	if err != nil {
		e.logger.Debug("logs not found", "hash", hexTx, "error", err.Error())
	}

	// get eth index based on block's txs
	var txIndex uint64
	msgs := e.backend.GetEthereumMsgsFromTendermintBlock(resBlock)
	for i := range msgs {
		if msgs[i].Hash == hexTx {
			txIndex = uint64(i)
			break
		}
	}

	receipt := map[string]interface{}{
		// Consensus fields: These fields are defined by the Yellow Paper
		"status":            status,
		"cumulativeGasUsed": hexutil.Uint64(cumulativeGasUsed),
		"logsBloom":         ethtypes.BytesToBloom(ethtypes.LogsBloom(logs)),
		"logs":              logs,

		// Implementation fields: These fields are added by geth when processing a transaction.
		// They are stored in the chain database.
		"transactionHash": hash,
		"contractAddress": nil,
		"gasUsed":         hexutil.Uint64(res.TxResult.GasUsed),
		"type":            hexutil.Uint(txData.TxType()),

		// Inclusion information: These fields provide information about the inclusion of the
		// transaction corresponding to this receipt.
		"blockHash":        common.BytesToHash(resBlock.Block.Header.Hash()).Hex(),
		"blockNumber":      hexutil.Uint64(res.Height),
		"transactionIndex": hexutil.Uint64(txIndex),

		// sender and receiver (contract or EOA) addreses
		"from": sdk.AccAddress(from.Bytes()),
		"to":   nil,
	}
	if to != nil {
		receipt["to"] = sdk.AccAddress(to.Bytes())
	}
	if logs == nil {
		receipt["logs"] = [][]EthLog{}
	} else {
		var newRes []EthLog
		for _, v := range logs {
			var result = EthLog{v}
			newRes = append(newRes, result)
		}
		receipt["logs"] = newRes
	}

	// If the ContractAddress is 20 0x0 bytes, assume it is not a contract creation
	if txData.GetTo() == nil {
		contractAddr := crypto.CreateAddress(from, txData.GetNonce())
		receipt["contractAddress"] = sdk.AccAddress(contractAddr.Bytes())
	}
	return receipt, nil
}
