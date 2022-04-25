package rpc

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/params"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/pkg/errors"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	rpctypes "github.com/oracleNetworkProtocol/plugchain/rpc/ethereum/types"
	ethermint "github.com/tharsis/ethermint/types"
	evmtypes "github.com/tharsis/ethermint/x/evm/types"
)

// ClientCtx returns client context
func (e *PublicAPI) ClientCtx() client.Context {
	return e.clientCtx
}

func (e *PublicAPI) QueryClient() *rpctypes.QueryClient {
	return e.queryClient
}

func (e *PublicAPI) Ctx() context.Context {
	return e.ctx
}

// ProtocolVersion returns the supported Ethereum protocol version.
func (e *PublicAPI) ProtocolVersion() hexutil.Uint {
	e.logger.Debug("rpc_protocolVersion")
	return hexutil.Uint(ethermint.ProtocolVersion)
}

// ChainId is the EIP-155 replay-protection chain id for the current ethereum chain config.
func (e *PublicAPI) ChainId() (*big.Int, error) { // nolint
	e.logger.Debug("rpc_chainId")
	// if current block is at or past the EIP-155 replay-protection fork block, return chainID from config
	bn, err := e.backend.BlockNumber()
	if err != nil {
		e.logger.Debug("failed to fetch latest block number", "error", err.Error())
		return e.chainIDEpoch, nil
	}

	if config := e.backend.ChainConfig(); config.IsEIP155(new(big.Int).SetUint64(uint64(bn))) {
		return config.ChainID, nil
	}

	return nil, fmt.Errorf("chain not synced beyond EIP-155 replay-protection fork block")
}

// Syncing returns false in case the node is currently not syncing with the network. It can be up to date or has not
// yet received the latest block headers from its pears. In case it is synchronizing:
// - startingBlock: block number this node started to synchronize from
// - currentBlock:  block number this node is currently importing
// - highestBlock:  block number of the highest block header this node has received from peers
// - pulledStates:  number of state entries processed until now
// - knownStates:   number of known state entries that still need to be pulled
func (e *PublicAPI) Syncing() (interface{}, error) {
	e.logger.Debug("rpc_syncing")

	status, err := e.clientCtx.Client.Status(e.ctx)
	if err != nil {
		return false, err
	}

	if !status.SyncInfo.CatchingUp {
		return false, nil
	}

	return map[string]interface{}{
		"startingBlock": status.SyncInfo.EarliestBlockHeight,
		"currentBlock":  status.SyncInfo.LatestBlockHeight,
		// "highestBlock":  nil, // NA
		// "pulledStates":  nil, // NA
		// "knownStates":   nil, // NA
	}, nil
}

// Coinbase is the address that staking rewards will be send to (alias for Etherbase).
func (e *PublicAPI) Coinbase() (string, error) {
	e.logger.Debug("rpc_coinbase")

	coinbase, err := e.backend.GetCoinbase()
	if err != nil {
		return "", err
	}
	return coinbase.String(), nil
}

// GasPrice returns the current gas price based on Ethermint's gas price oracle.
func (e *PublicAPI) GasPrice() (*big.Int, error) {
	e.logger.Debug("rpc_gasPrice")
	tipcap, err := e.backend.SuggestGasTipCap()
	if err != nil {
		return nil, err
	}

	if head := e.backend.CurrentHeader(); head.BaseFee != nil {
		tipcap.Add(tipcap, head.BaseFee)
	}

	return tipcap, nil
}

// MaxPriorityFeePerGas returns a suggestion for a gas tip cap for dynamic fee transactions.
func (e *PublicAPI) MaxPriorityFeePerGas() (*big.Int, error) {
	e.logger.Debug("rpc_maxPriorityFeePerGas")
	tipcap, err := e.backend.SuggestGasTipCap()
	if err != nil {
		return nil, err
	}

	return tipcap, nil
}

// Accounts returns the list of accounts available to this node.
func (e *PublicAPI) Accounts() ([]sdk.AccAddress, error) {
	e.logger.Debug("rpc_accounts")

	addresses := make([]sdk.AccAddress, 0) // return [] instead of nil if empty

	infos, err := e.clientCtx.Keyring.List()
	if err != nil {
		return addresses, err
	}

	for _, info := range infos {
		addressBytes := info.GetPubKey().Address().Bytes()
		addresses = append(addresses, sdk.AccAddress(addressBytes))
	}

	return addresses, nil
}

// BlockNumber returns the current block number.
func (e *PublicAPI) BlockNumber() (uint64, error) {
	e.logger.Debug("rpc_blockNumber")
	result, err := e.backend.BlockNumber()
	if err != nil {
		return 0, nil
	}
	return uint64(result), nil
}

// GetBalance returns the provided account's balance up to the provided block number.
func (e *PublicAPI) GetBalance(addr sdk.AccAddress, blockNrOrHash rpctypes.BlockNumberOrHash) (*big.Int, error) {
	address := common.BytesToAddress(addr)
	e.logger.Debug("rpc_getBalance", "address", address.String(), "block number or hash", blockNrOrHash)

	blockNum, err := e.getBlockNumber(blockNrOrHash)
	if err != nil {
		return nil, err
	}

	req := &evmtypes.QueryBalanceRequest{
		Address: address.String(),
	}

	res, err := e.queryClient.Balance(rpctypes.ContextWithHeight(blockNum.Int64()), req)
	if err != nil {
		return nil, err
	}

	val, ok := sdk.NewIntFromString(res.Balance)
	if !ok {
		return nil, errors.New("invalid balance")
	}

	return val.BigInt(), nil
}

// GetStorageAt returns the contract storage at the given address, block number, and key.
func (e *PublicAPI) GetStorageAt(addr sdk.AccAddress, key string, blockNrOrHash rpctypes.BlockNumberOrHash) (string, error) {
	address := common.BytesToAddress(addr)
	e.logger.Debug("rpc_getStorageAt", "address", address.Hex(), "key", key, "block number or hash", blockNrOrHash)

	blockNum, err := e.getBlockNumber(blockNrOrHash)
	if err != nil {
		return "", err
	}

	req := &evmtypes.QueryStorageRequest{
		Address: address.String(),
		Key:     key,
	}

	res, err := e.queryClient.Storage(rpctypes.ContextWithHeight(blockNum.Int64()), req)
	if err != nil {
		return "", err
	}

	value := common.HexToHash(res.Value)
	return value.String(), nil
}

// GetTransactionCount returns the number of transactions at the given address up to the given block number.
func (e *PublicAPI) GetTransactionCount(addr sdk.AccAddress, blockNrOrHash rpctypes.BlockNumberOrHash) (uint64, error) {
	address := common.BytesToAddress(addr)
	e.logger.Debug("rpc_getTransactionCount", "address", address.Hex(), "block number or hash", blockNrOrHash)
	blockNum, err := e.getBlockNumber(blockNrOrHash)
	if err != nil {
		return 0, err
	}
	result, err := e.backend.GetTransactionCount(address, blockNum)
	return uint64(*result), err
}

// GetBlockTransactionCountByHash returns the number of transactions in the block identified by hash.
func (e *PublicAPI) GetBlockTransactionCountByHash(hash common.Hash) uint {
	e.logger.Debug("rpc_getBlockTransactionCountByHash", "hash", hash.Hex())

	resBlock, err := e.clientCtx.Client.BlockByHash(e.ctx, hash.Bytes())
	if err != nil {
		e.logger.Debug("block not found", "hash", hash.Hex(), "error", err.Error())
		return 0
	}

	if resBlock.Block == nil {
		e.logger.Debug("block not found", "hash", hash.Hex())
		return 0
	}

	ethMsgs := e.backend.GetEthereumMsgsFromTendermintBlock(resBlock)
	return uint(len(ethMsgs))
}

// GetBlockTransactionCountByNumber returns the number of transactions in the block identified by number.
func (e *PublicAPI) GetBlockTransactionCountByNumber(blockNum rpctypes.BlockNumber) uint {
	e.logger.Debug("rpc_getBlockTransactionCountByNumber", "height", blockNum.Int64())
	resBlock, err := e.clientCtx.Client.Block(e.ctx, blockNum.TmHeight())
	if err != nil {
		e.logger.Debug("block not found", "height", blockNum.Int64(), "error", err.Error())
		return 0
	}

	if resBlock.Block == nil {
		e.logger.Debug("block not found", "height", blockNum.Int64())
		return 0
	}

	ethMsgs := e.backend.GetEthereumMsgsFromTendermintBlock(resBlock)
	return uint(len(ethMsgs))
}

// GetCode returns the contract code at the given address and block number.
func (e *PublicAPI) GetCode(addr sdk.AccAddress, blockNrOrHash rpctypes.BlockNumberOrHash) (hexutil.Bytes, error) {
	address := common.BytesToAddress(addr)
	e.logger.Debug("rpc_getCode", "address", address.Hex(), "block number or hash", blockNrOrHash)

	blockNum, err := e.getBlockNumber(blockNrOrHash)
	if err != nil {
		return nil, err
	}

	req := &evmtypes.QueryCodeRequest{
		Address: address.String(),
	}

	res, err := e.queryClient.Code(rpctypes.ContextWithHeight(blockNum.Int64()), req)
	if err != nil {
		return nil, err
	}

	return res.Code, nil
}

//TODO:
// GetTransactionLogs returns the logs given a transaction hash.
func (e *PublicAPI) GetTransactionLogs(txHash common.Hash) ([]*EthLog, error) {
	e.logger.Debug("rpc_getTransactionLogs", "hash", txHash)
	logs, err := e.backend.GetTransactionLogs(txHash)
	if err != nil {
		return nil, err
	}
	results := make([]*EthLog, 0)
	for _, log := range logs {
		var result = EthLog{log}
		results = append(results, &result)
	}
	return results, err
}

//TODO:
// Sign signs the provided data using the private key of address via Geth's signature standard.
func (e *PublicAPI) Sign(addr sdk.AccAddress, data hexutil.Bytes) (hexutil.Bytes, error) {
	address := common.BytesToAddress(addr)
	e.logger.Debug("rpc_sign", "address", address.Hex(), "data", common.Bytes2Hex(data))

	from := sdk.AccAddress(address.Bytes())

	_, err := e.clientCtx.Keyring.KeyByAddress(from)
	if err != nil {
		e.logger.Error("failed to find key in keyring", "address", address.String())
		return nil, fmt.Errorf("%s; %s", keystore.ErrNoMatch, err.Error())
	}

	// Sign the requested hash with the wallet
	signature, _, err := e.clientCtx.Keyring.SignByAddress(from, data)
	if err != nil {
		e.logger.Error("keyring.SignByAddress failed", "address", address.Hex())
		return nil, err
	}

	signature[64] += 27 // Transform V from 0/1 to 27/28 according to the yellow paper
	return signature, nil
}

// SendTransaction sends an Ethereum transaction.
func (e *PublicAPI) SendTransaction(trans rpctypes.TransactionArgs) (common.Hash, error) {
	args := trans.New()

	e.logger.Debug("rpc_sendTransaction", "args", args.String())
	return e.backend.SendTransaction(args)
}

// FillTransaction fills the defaults (nonce, gas, gasPrice or 1559 fields)
// on a given unsigned transaction, and returns it to the caller for further
// processing (signing + broadcast).
func (e *PublicAPI) FillTransaction(trans rpctypes.TransactionArgs) (*rpctypes.SignTransactionResult, error) {
	args := trans.New()
	// Set some sanity defaults and terminate on failure
	args, err := e.backend.SetTxDefaults(args)
	if err != nil {
		return nil, err
	}

	// Assemble the transaction and obtain rlp
	tx := args.ToTransaction().AsTransaction()

	data, err := tx.MarshalBinary()
	if err != nil {
		return nil, err
	}

	return &rpctypes.SignTransactionResult{
		Raw: data,
		Tx:  tx,
	}, nil
}

//TODO:
// SendRawTransaction send a raw Ethereum transaction.
func (e *PublicAPI) SendRawTransaction(data hexutil.Bytes) (common.Hash, error) {
	e.logger.Debug("rpc_sendRawTransaction", "length", len(data))

	// RLP decode raw transaction bytes
	tx := &ethtypes.Transaction{}
	if err := tx.UnmarshalBinary(data); err != nil {
		e.logger.Error("transaction decoding failed", "error", err.Error())
		return common.Hash{}, err
	}

	ethereumTx := &evmtypes.MsgEthereumTx{}
	if err := ethereumTx.FromEthereumTx(tx); err != nil {
		e.logger.Error("transaction converting failed", "error", err.Error())
		return common.Hash{}, err
	}

	if err := ethereumTx.ValidateBasic(); err != nil {
		e.logger.Debug("tx failed basic validation", "error", err.Error())
		return common.Hash{}, err
	}

	// Query params to use the EVM denomination
	res, err := e.queryClient.QueryClient.Params(e.ctx, &evmtypes.QueryParamsRequest{})
	if err != nil {
		e.logger.Error("failed to query evm params", "error", err.Error())
		return common.Hash{}, err
	}

	cosmosTx, err := ethereumTx.BuildTx(e.clientCtx.TxConfig.NewTxBuilder(), res.Params.EvmDenom)
	if err != nil {
		e.logger.Error("failed to build cosmos tx", "error", err.Error())
		return common.Hash{}, err
	}

	// Encode transaction by default Tx encoder
	txBytes, err := e.clientCtx.TxConfig.TxEncoder()(cosmosTx)
	if err != nil {
		e.logger.Error("failed to encode eth tx using default encoder", "error", err.Error())
		return common.Hash{}, err
	}

	txHash := ethereumTx.AsTransaction().Hash()

	syncCtx := e.clientCtx.WithBroadcastMode(flags.BroadcastSync)
	rsp, err := syncCtx.BroadcastTx(txBytes)
	if rsp != nil && rsp.Code != 0 {
		err = sdkerrors.ABCIError(rsp.Codespace, rsp.Code, rsp.RawLog)
	}
	if err != nil {
		e.logger.Error("failed to broadcast tx", "error", err.Error())
		return txHash, err
	}

	return txHash, nil
}

// checkTxFee is an internal function used to check whether the fee of
// the given transaction is _reasonable_(under the cap).
func checkTxFee(gasPrice *big.Int, gas uint64, cap float64) error {
	// Short circuit if there is no cap for transaction fee at all.
	if cap == 0 {
		return nil
	}
	totalfee := new(big.Float).SetInt(new(big.Int).Mul(gasPrice, new(big.Int).SetUint64(gas)))
	// 1 photon in 10^18 aphoton
	oneToken := new(big.Float).SetInt(big.NewInt(params.Ether))
	// quo = rounded(x/y)
	feeEth := new(big.Float).Quo(totalfee, oneToken)
	// no need to check error from parsing
	feeFloat, _ := feeEth.Float64()
	if feeFloat > cap {
		return fmt.Errorf("tx fee (%.2f ether) exceeds the configured cap (%.2f ether)", feeFloat, cap)
	}
	return nil
}

// Resend accepts an existing transaction and a new gas price and limit. It will remove
// the given transaction from the pool and reinsert it with the new gas price and limit.
func (e *PublicAPI) Resend(ctx context.Context, trans rpctypes.TransactionArgs, gasPrice *hexutil.Big, gasLimit *hexutil.Uint64) (common.Hash, error) {
	args := trans.New()
	e.logger.Debug("rpc_resend", "args", args.String())
	if args.Nonce == nil {
		return common.Hash{}, fmt.Errorf("missing transaction nonce in transaction spec")
	}

	args, err := e.backend.SetTxDefaults(args)
	if err != nil {
		return common.Hash{}, err
	}

	matchTx := args.ToTransaction().AsTransaction()

	// Before replacing the old transaction, ensure the _new_ transaction fee is reasonable.
	price := matchTx.GasPrice()
	if gasPrice != nil {
		price = gasPrice.ToInt()
	}
	gas := matchTx.Gas()
	if gasLimit != nil {
		gas = uint64(*gasLimit)
	}
	if err := checkTxFee(price, gas, e.backend.RPCTxFeeCap()); err != nil {
		return common.Hash{}, err
	}

	pending, err := e.backend.PendingTransactions()
	if err != nil {
		return common.Hash{}, err
	}

	for _, tx := range pending {
		p, err := evmtypes.UnwrapEthereumMsg(tx)
		if err != nil {
			// not valid ethereum tx
			continue
		}

		pTx := p.AsTransaction()

		wantSigHash := e.signer.Hash(matchTx)
		pFrom, err := ethtypes.Sender(e.signer, pTx)
		if err != nil {
			continue
		}

		if pFrom == *args.From && e.signer.Hash(pTx) == wantSigHash {
			// Match. Re-sign and send the transaction.
			if gasPrice != nil && (*big.Int)(gasPrice).Sign() != 0 {
				args.GasPrice = gasPrice
			}
			if gasLimit != nil && *gasLimit != 0 {
				args.Gas = gasLimit
			}

			return e.backend.SendTransaction(args) // TODO: this calls SetTxDefaults again, refactor to avoid calling it twice
		}
	}

	return common.Hash{}, fmt.Errorf("transaction %#x not found", matchTx.Hash())
}

// Call performs a raw contract call.
func (e *PublicAPI) Call(trans rpctypes.TransactionArgs, blockNrOrHash rpctypes.BlockNumberOrHash, _ *rpctypes.StateOverride) (hexutil.Bytes, error) {
	args := trans.New()
	e.logger.Debug("rpc_call", "args", args.String(), "block number or hash", blockNrOrHash)

	blockNum, err := e.getBlockNumber(blockNrOrHash)
	if err != nil {
		return nil, err
	}
	data, err := e.doCall(args, blockNum)
	if err != nil {
		return []byte{}, err
	}

	return (hexutil.Bytes)(data.Ret), nil
}

// DoCall performs a simulated call operation through the evmtypes. It returns the
// estimated gas used on the operation or an error if fails.
func (e *PublicAPI) doCall(
	args evmtypes.TransactionArgs, blockNr rpctypes.BlockNumber,
) (*evmtypes.MsgEthereumTxResponse, error) {
	bz, err := json.Marshal(&args)
	if err != nil {
		return nil, err
	}

	req := evmtypes.EthCallRequest{
		Args:   bz,
		GasCap: e.backend.RPCGasCap(),
	}

	// From ContextWithHeight: if the provided height is 0,
	// it will return an empty context and the gRPC query will use
	// the latest block height for querying.
	ctx := rpctypes.ContextWithHeight(blockNr.Int64())
	timeout := e.backend.RPCEVMTimeout()

	// Setup context so it may be canceled the call has completed
	// or, in case of unmetered gas, setup a context with a timeout.
	var cancel context.CancelFunc
	if timeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, timeout)
	} else {
		ctx, cancel = context.WithCancel(ctx)
	}

	// Make sure the context is canceled when the call has completed
	// this makes sure resources are cleaned up.
	defer cancel()

	res, err := e.queryClient.EthCall(ctx, &req)
	if err != nil {
		return nil, err
	}

	if res.Failed() {
		if res.VmError != vm.ErrExecutionReverted.Error() {
			return nil, status.Error(codes.Internal, res.VmError)
		}
		return nil, evmtypes.NewExecErrorWithReason(res.Ret)
	}

	return res, nil
}

// EstimateGas returns an estimate of gas usage for the given smart contract call.
func (e *PublicAPI) EstimateGas(trans rpctypes.TransactionArgs, blockNrOptional *rpctypes.BlockNumber) (hexutil.Uint64, error) {
	var args evmtypes.TransactionArgs
	args = trans.TransactionArgs
	from, to := common.BytesToAddress(trans.From), common.BytesToAddress(trans.To)
	args.From, args.To = &from, &to
	e.logger.Debug("rpc_estimateGas")
	return e.backend.EstimateGas(args, blockNrOptional)
}

// GetBlockByHash returns the block identified by hash.
func (e *PublicAPI) GetBlockByHash(hash common.Hash, fullTx bool) (map[string]interface{}, error) {
	e.logger.Debug("rpc_getBlockByHash", "hash", hash.Hex(), "full", fullTx)
	return e.backend.GetBlockByHash(hash, fullTx)
}

// GetBlockByNumber returns the block identified by number.
func (e *PublicAPI) GetBlockByNumber(ethBlockNum rpctypes.BlockNumber, fullTx bool) (map[string]interface{}, error) {
	e.logger.Debug("rpc_getBlockByNumber", "number", ethBlockNum, "full", fullTx)
	return e.backend.GetBlockByNumber(ethBlockNum, fullTx)
}

// GetTransactionByHash returns the transaction identified by hash.
func (e *PublicAPI) GetTransactionByHash(hash common.Hash) (*rpctypes.RPCTransaction, error) {
	e.logger.Debug("rpc_getTransactionByHash", "hash", hash.Hex())
	return e.backend.GetTransactionByHash(hash)
}

// GetTransactionByBlockHashAndIndex returns the transaction identified by hash and index.
func (e *PublicAPI) GetTransactionByBlockHashAndIndex(hash common.Hash, idx hexutil.Uint) (*rpctypes.RPCTransaction, error) {
	e.logger.Debug("rpc_getTransactionByBlockHashAndIndex", "hash", hash.Hex(), "index", idx)

	resBlock, err := e.clientCtx.Client.BlockByHash(e.ctx, hash.Bytes())
	if err != nil {
		e.logger.Debug("block not found", "hash", hash.Hex(), "error", err.Error())
		return nil, nil
	}

	if resBlock.Block == nil {
		e.logger.Debug("block not found", "hash", hash.Hex())
		return nil, nil
	}

	i := int(idx)
	ethMsgs := e.backend.GetEthereumMsgsFromTendermintBlock(resBlock)
	if i >= len(ethMsgs) {
		e.logger.Debug("block txs index out of bound", "index", i)
		return nil, nil
	}

	msg := ethMsgs[i]

	baseFee, err := e.backend.BaseFee(resBlock.Block.Height)
	if err != nil {
		return nil, err
	}

	return rpctypes.NewTransactionFromMsg(
		msg,
		hash,
		uint64(resBlock.Block.Height),
		uint64(idx),
		baseFee,
	)
}

// GetTransactionByBlockNumberAndIndex returns the transaction identified by number and index.
func (e *PublicAPI) GetTransactionByBlockNumberAndIndex(blockNum rpctypes.BlockNumber, idx hexutil.Uint) (*rpctypes.RPCTransaction, error) {
	e.logger.Debug("rpc_getTransactionByBlockNumberAndIndex", "number", blockNum, "index", idx)

	resBlock, err := e.clientCtx.Client.Block(e.ctx, blockNum.TmHeight())
	if err != nil {
		e.logger.Debug("block not found", "height", blockNum.Int64(), "error", err.Error())
		return nil, nil
	}

	if resBlock.Block == nil {
		e.logger.Debug("block not found", "height", blockNum.Int64())
		return nil, nil
	}

	i := int(idx)
	ethMsgs := e.backend.GetEthereumMsgsFromTendermintBlock(resBlock)
	if i >= len(ethMsgs) {
		e.logger.Debug("block txs index out of bound", "index", i)
		return nil, nil
	}

	msg := ethMsgs[i]

	return rpctypes.NewTransactionFromMsg(
		msg,
		common.BytesToHash(resBlock.Block.Hash()),
		uint64(resBlock.Block.Height),
		uint64(idx),
		e.chainIDEpoch,
	)
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
		"from": from,
		"to":   txData.GetTo(),
	}

	if logs == nil {
		receipt["logs"] = [][]*ethtypes.Log{}
	}

	// If the ContractAddress is 20 0x0 bytes, assume it is not a contract creation
	if txData.GetTo() == nil {
		receipt["contractAddress"] = crypto.CreateAddress(from, txData.GetNonce())
	}

	return receipt, nil
}

// GetPendingTransactions returns the transactions that are in the transaction pool
// and have a from address that is one of the accounts this node manages.
func (e *PublicAPI) GetPendingTransactions() ([]*rpctypes.RPCTransaction, error) {
	e.logger.Debug("rpc_getPendingTransactions")

	txs, err := e.backend.PendingTransactions()
	if err != nil {
		return nil, err
	}

	result := make([]*rpctypes.RPCTransaction, 0, len(txs))
	for _, tx := range txs {
		msg, err := evmtypes.UnwrapEthereumMsg(tx)
		if err != nil {
			// not valid ethereum tx
			continue
		}

		rpctx, err := rpctypes.NewTransactionFromMsg(
			msg,
			common.Hash{},
			uint64(0),
			uint64(0),
			e.chainIDEpoch,
		)
		if err != nil {
			return nil, err
		}

		result = append(result, rpctx)
	}

	return result, nil
}

// GetUncleByBlockHashAndIndex returns the uncle identified by hash and index. Always returns nil.
func (e *PublicAPI) GetUncleByBlockHashAndIndex(hash common.Hash, idx hexutil.Uint) map[string]interface{} {
	return nil
}

// GetUncleByBlockNumberAndIndex returns the uncle identified by number and index. Always returns nil.
func (e *PublicAPI) GetUncleByBlockNumberAndIndex(number, idx hexutil.Uint) map[string]interface{} {
	return nil
}

// GetProof returns an account object with proof and any storage proofs
func (e *PublicAPI) GetProof(addr sdk.AccAddress, storageKeys []string, blockNrOrHash rpctypes.BlockNumberOrHash) (*rpctypes.AccountResult, error) {
	address := common.BytesToAddress(addr)
	e.logger.Debug("rpc_getProof", "address", address.Hex(), "keys", storageKeys, "block number or hash", blockNrOrHash)

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
	accountKey := authtypes.AddressStoreKey(sdk.AccAddress(address.Bytes()))
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

	return &rpctypes.AccountResult{
		Address:      address,
		AccountProof: []string{accProofStr},
		Balance:      (*hexutil.Big)(balance.BigInt()),
		CodeHash:     common.HexToHash(res.CodeHash),
		Nonce:        hexutil.Uint64(res.Nonce),
		StorageHash:  common.Hash{}, // NOTE: Ethermint doesn't have a storage hash. TODO: implement?
		StorageProof: storageProofs,
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
