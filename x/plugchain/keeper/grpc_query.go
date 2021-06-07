package keeper

import (
	"github.com/oracleNetworkProtocol/plugchain/x/plugchain/types"
)

var _ types.QueryServer = Keeper{}
