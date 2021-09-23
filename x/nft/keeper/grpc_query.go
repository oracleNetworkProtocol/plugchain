package keeper

import (
	"github.com/oracleNetworkProtocol/plugchain/x/nft/types"
)

var _ types.QueryServer = Keeper{}
