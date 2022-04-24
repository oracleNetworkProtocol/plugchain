package rpc

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"

	rpctypes "github.com/oracleNetworkProtocol/plugchain/rpc/ethereum/types"
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
