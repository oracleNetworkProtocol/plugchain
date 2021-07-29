package keeper

import (
	"github.com/oracleNetworkProtocol/plugchain/x/plugchain/types"
)

type Querier struct {
	Keeper
}

var _ types.QueryServer = Querier{}

// func NewQuerierServer(keeper Keeper) types.QueryServer {
// 	return &Querier{Keeper: keeper}
// }
