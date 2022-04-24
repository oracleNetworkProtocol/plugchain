package rpc

import (
	"context"
	"math/big"

	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"

	ethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/oracleNetworkProtocol/plugchain/rpc/ethereum/backend"
	rpctypes "github.com/oracleNetworkProtocol/plugchain/rpc/ethereum/types"
	"github.com/tharsis/ethermint/crypto/hd"
	ethermint "github.com/tharsis/ethermint/types"
	evmtypes "github.com/tharsis/ethermint/x/evm/types"
)

// PublicAPI is the eth_ prefixed set of APIs in the Web3 JSON-RPC spec.
type PublicAPI struct {
	ctx          context.Context
	logger       log.Logger
	backend      backend.Backend
	clientCtx    client.Context
	queryClient  *rpctypes.QueryClient
	chainIDEpoch *big.Int
	nonceLock    *rpctypes.AddrLocker
	signer       ethtypes.Signer
}

// NewPublicAPI creates an instance of the public ETH Web3 API.
func NewPublicAPI(
	logger log.Logger,
	clientCtx client.Context,
	backend backend.Backend,
	nonceLock *rpctypes.AddrLocker,
) *PublicAPI {
	eip155ChainID, err := ethermint.ParseChainID(clientCtx.ChainID)
	if err != nil {
		panic(err)
	}

	algos, _ := clientCtx.Keyring.SupportedAlgorithms()

	if !algos.Contains(hd.EthSecp256k1) {
		kr, err := keyring.New(
			sdk.KeyringServiceName(),
			viper.GetString(flags.FlagKeyringBackend),
			clientCtx.KeyringDir,
			clientCtx.Input,
			hd.EthSecp256k1Option(),
		)
		if err != nil {
			panic(err)
		}

		clientCtx = clientCtx.WithKeyring(kr)
	}

	// The signer used by the API should always be the 'latest' known one because we expect
	// signers to be backwards-compatible with old transactions.
	cfg := backend.ChainConfig()
	if cfg == nil {
		cfg = evmtypes.DefaultChainConfig().EthereumConfig(eip155ChainID)
	}

	signer := ethtypes.LatestSigner(cfg)

	api := &PublicAPI{
		ctx:          context.Background(),
		clientCtx:    clientCtx,
		queryClient:  rpctypes.NewQueryClient(clientCtx),
		chainIDEpoch: eip155ChainID,
		logger:       logger.With("module", "rpc"),
		backend:      backend,
		nonceLock:    nonceLock,
		signer:       signer,
	}

	return api
}
