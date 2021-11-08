package nft

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/oracleNetworkProtocol/plugchain/x/nft/keeper"
	"github.com/oracleNetworkProtocol/plugchain/x/nft/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	if err := genState.Validate(); err != nil {
		panic(err.Error())
	}
	for _, v := range genState.Collections {
		if err := k.SetClass(ctx, v.Class); err != nil {
			panic(err)
		}
		if err := k.SetCollection(ctx, v); err != nil {
			panic(err)
		}
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {

	return types.NewGenesisState(k.GetCollections(ctx))
}

func DefaultGenesisState() *types.GenesisState {
	return types.NewGenesisState([]types.Collection{})
}
