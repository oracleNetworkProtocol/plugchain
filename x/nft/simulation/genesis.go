package simulation

import (
	"math/rand"
	"strings"

	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/oracleNetworkProtocol/plugchain/x/nft/types"
)

// RandomizedGenState generates a random GenesisState for nft
func RandomizedGenState(simState *module.SimulationState) {
	simState.GenState[types.ModuleName] = nil
}

func RandnNFTID(r *rand.Rand, min, max int) string {
	n := simtypes.RandIntBetween(r, min, max)
	id := simtypes.RandStringOfLength(r, n)
	return strings.ToLower(id)
}
