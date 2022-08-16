package rpc

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/pkg/errors"

	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"

	evmtypes "github.com/evmos/ethermint/x/evm/types"
	"github.com/oracleNetworkProtocol/plugchain/rpc/backend"
	rpctypes "github.com/oracleNetworkProtocol/plugchain/rpc/types"
)

// Sha3 returns the keccak-256 hash of the passed-in input.
func (e *PublicAPI) Sha3(input hexutil.Bytes) hexutil.Bytes {
	e.logger.Debug("rpc_sha3")

	return crypto.Keccak256(input)
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
func (e *PublicAPI) GasPrice() (*hexutil.Big, error) {
	e.logger.Debug("rpc_gasPrice")
	var (
		result *big.Int
		err    error
	)
	if head := e.backend.CurrentHeader(); head.BaseFee != nil {
		result, err = e.backend.SuggestGasTipCap(head.BaseFee)
		if err != nil {
			return nil, err
		}
		result = result.Add(result, head.BaseFee)
	} else {
		result = big.NewInt(e.backend.RPCMinGasPrice())
	}

	return (*hexutil.Big)(result), nil
}

// MaxPriorityFeePerGas returns a suggestion for a gas tip cap for dynamic fee transactions.
func (e *PublicAPI) MaxPriorityFeePerGas() (*hexutil.Big, error) {
	e.logger.Debug("rpc_maxPriorityFeePerGas")
	head := e.backend.CurrentHeader()
	tipcap, err := e.backend.SuggestGasTipCap(head.BaseFee)
	if err != nil {
		return nil, err
	}
	return (*hexutil.Big)(tipcap), nil
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
func (e *PublicAPI) BlockNumber() (hexutil.Uint64, error) {
	e.logger.Debug("rpc_blockNumber")
	return e.backend.BlockNumber()
}

// GetBalance returns the provided account's balance up to the provided block number.
func (e *PublicAPI) GetBalance(addr sdk.AccAddress, blockNrOrHash rpctypes.BlockNumberOrHash) (*hexutil.Big, error) {
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

	return (*hexutil.Big)(val.BigInt()), nil
}

// GetStorageAt returns the contract storage at the given address, block number, and key.
func (e *PublicAPI) GetStorageAt(addr sdk.AccAddress, key string, blockNrOrHash rpctypes.BlockNumberOrHash) (hexutil.Bytes, error) {
	address := common.BytesToAddress(addr)
	e.logger.Debug("rpc_getStorageAt", "address", address.Hex(), "key", key, "block number or hash", blockNrOrHash)

	blockNum, err := e.getBlockNumber(blockNrOrHash)
	if err != nil {
		return nil, err
	}

	req := &evmtypes.QueryStorageRequest{
		Address: address.String(),
		Key:     key,
	}

	res, err := e.queryClient.Storage(rpctypes.ContextWithHeight(blockNum.Int64()), req)
	if err != nil {
		return nil, err
	}

	value := common.HexToHash(res.Value)
	return value.Bytes(), nil
}

// GetTransactionCount returns the number of transactions at the given address up to the given block number.
func (e *PublicAPI) GetTransactionCount(addr sdk.AccAddress, blockNrOrHash rpctypes.BlockNumberOrHash) (*hexutil.Uint64, error) {
	address := common.BytesToAddress(addr)
	e.logger.Debug("rpc_getTransactionCount", "address", address.Hex(), "block number or hash", blockNrOrHash)
	blockNum, err := e.getBlockNumber(blockNrOrHash)
	if err != nil {
		return nil, err
	}
	return e.backend.GetTransactionCount(address, blockNum)
}

// GetBlockTransactionCountByHash returns the number of transactions in the block identified by hash.
func (e *PublicAPI) GetBlockTransactionCountByHash(hash common.Hash) *hexutil.Uint {
	e.logger.Debug("rpc_getBlockTransactionCountByHash", "hash", hash.Hex())

	block, err := e.clientCtx.Client.BlockByHash(e.ctx, hash.Bytes())
	if err != nil {
		e.logger.Debug("block not found", "hash", hash.Hex(), "error", err.Error())
		return nil
	}

	if block.Block == nil {
		e.logger.Debug("block not found", "hash", hash.Hex())
		return nil
	}

	blockRes, err := e.clientCtx.Client.BlockResults(e.ctx, &block.Block.Height)
	if err != nil {
		return nil
	}

	ethMsgs := e.backend.GetEthereumMsgsFromTendermintBlock(block, blockRes)
	n := hexutil.Uint(len(ethMsgs))
	return &n
}

// GetBlockTransactionCountByNumber returns the number of transactions in the block identified by number.
func (e *PublicAPI) GetBlockTransactionCountByNumber(blockNum rpctypes.BlockNumber) *hexutil.Uint {
	e.logger.Debug("rpc_getBlockTransactionCountByNumber", "height", blockNum.Int64())
	block, err := e.clientCtx.Client.Block(e.ctx, blockNum.TmHeight())
	if err != nil {
		e.logger.Debug("block not found", "height", blockNum.Int64(), "error", err.Error())
		return nil
	}

	if block.Block == nil {
		e.logger.Debug("block not found", "height", blockNum.Int64())
		return nil
	}

	blockRes, err := e.clientCtx.Client.BlockResults(e.ctx, &block.Block.Height)
	if err != nil {
		return nil
	}

	ethMsgs := e.backend.GetEthereumMsgsFromTendermintBlock(block, blockRes)
	n := hexutil.Uint(len(ethMsgs))
	return &n
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

// GetTransactionLogs returns the logs given a transaction hash.
func (e *PublicAPI) GetTransactionLogs(txHash common.Hash) ([]*EthLog, error) {
	e.logger.Debug("rpc_getTransactionLogs", "hash", txHash)

	hexTx := txHash.Hex()
	res, err := e.backend.GetTxByEthHash(txHash)
	if err != nil {
		e.logger.Debug("tx not found", "hash", hexTx, "error", err.Error())
		return nil, nil
	}

	msgIndex, _ := rpctypes.FindTxAttributes(res.TxResult.Events, hexTx)
	if msgIndex < 0 {
		return nil, fmt.Errorf("ethereum tx not found in msgs: %s", hexTx)
	}
	// parse tx logs from events
	logs, err := backend.TxLogsFromEvents(res.TxResult.Events, msgIndex)

	if err != nil {
		return nil, err
	}
	results := make([]*EthLog, 0)
	for _, log := range logs {
		results = append(results, &EthLog{log})
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
		p, err := evmtypes.UnwrapEthereumMsg(tx, common.Hash{})
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
